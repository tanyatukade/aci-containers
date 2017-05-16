#!/usr/bin/env python

from __future__ import print_function

import argparse
import base64
import collections
import filecmp
import glob
import json
import os
import socket
import struct
import sys
import yaml


from apic import Apic, ApicKubeConfig
from jinja2 import Environment, PackageLoader


def info(msg):
    print("INFO: " + msg, file=sys.stderr)


def warn(msg):
    print("WARN: " + msg, file=sys.stderr)


def err(msg):
    print("ERR:  " + msg, file=sys.stderr)


def json_indent(s):
    return json.dumps(s, indent=4)


def yaml_quote(s):
    return "'%s'" % str(s).replace("'", "''")


def deep_merge(user, default):
    if isinstance(user, dict) and isinstance(default, dict):
        for k, v in default.iteritems():
            if k not in user:
                user[k] = v
            else:
                user[k] = deep_merge(user[k], v)
    return user


def config_default():
    # Default values for configuration
    default_config = {
        "aci_config": {
            "system_id": "kube",
            "vrf": {
                "name": "kube",
                "tenant": "common",
            },
            "l3out": {
                "name": "l3out",
                "external_networks": ["default"],
            },
            "vmm_domain": {
                "encap_type": "vxlan",
                "mcast_fabric": "225.1.2.3",
                "mcast_range": {
                    "start": "225.2.1.1",
                    "end": "225.2.255.255",
                },
            },
            "client_cert": False,
            "client_ssl": True,
        },
        "net_config": {
            "node_subnet": "10.1.0.1/16",
            "pod_subnet": "10.2.0.1/16",
            "extern_dynamic": "10.3.0.1/24",
            "extern_static": "10.4.0.1/24",
            "node_svc_subnet": "10.5.0.1/24",
            "kubeapi_vlan": 4001,
            "service_vlan": 4003,
            "infra_vlan": 4093,
        },
        "kube_config": {
            "controller": "1.1.1.1",
            "use_cluster_role": True,
            "use_ds_rolling_update": True,
        },
        "registry": {
            "image_prefix": "noiro",
        },
        "logging": {
            "controller_log_level": "info",
            "hostagent_log_level": "info",
            "opflexagent_log_level": "info",
            "aim_debug": "False",
        },
    }
    return default_config


def config_user(config_file):
    config = {}
    if config_file:
        if config_file == "-":
            info("Loading configuration from \"STDIN\"")
            config = yaml.load(sys.stdin)
        else:
            info("Loading configuration from \"%s\"" % config_file)
            with open(config_file, 'r') as file:
                config = yaml.load(file)
    if config is None:
        config = {}
    return config


def cidr_split(cidr):
    ip2int = lambda a: struct.unpack("!I", socket.inet_aton(a))[0]
    int2ip = lambda a: socket.inet_ntoa(struct.pack("!I", a))
    rtr, mask = cidr.split('/')
    maskbits = int('1' * (32 - int(mask)), 2)
    rtri = ip2int(rtr)
    starti = rtri + 1
    endi = (rtri | maskbits) - 1
    subi = (rtri & (0xffffffff ^ maskbits))
    return int2ip(starti), int2ip(endi), rtr, int2ip(subi), mask


def config_adjust(config, prov_apic):
    apic = None
    if prov_apic is not None:
        apic = get_apic(config)

    system_id = config["aci_config"]["system_id"]
    infra_vlan = config["net_config"]["infra_vlan"]
    if apic is not None:
        infra_vlan = apic.get_infravlan()

    pod_subnet = config["net_config"]["pod_subnet"]
    extern_dynamic = config["net_config"]["extern_dynamic"]
    extern_static = config["net_config"]["extern_static"]
    node_svc_subnet = config["net_config"]["node_svc_subnet"]
    encap_type = config["aci_config"]["vmm_domain"]["encap_type"]
    tenant = system_id

    adj_config = {
        "aci_config": {
            "cluster_tenant": tenant,
            "physical_domain": {
                "domain": system_id + "-pdom",
                "vlan_pool": system_id + "-pool",
            },
            "vmm_domain": {
                "domain": system_id,
                "controller": system_id,
                "mcast_pool": system_id + "-mpool",
            },
            "aim_login": {
                "username": system_id,
                # Tmp hack, till I generate certificates
                "password": "ToBeFixed!",
                "certfile": None,
            },
        },
        "net_config": {
            "infra_vlan": infra_vlan,
        },
        "node_config": {
            "encap_type": encap_type,
        },
        "kube_config": {
            "default_endpoint_group": {
                "tenant": tenant,
                "app_profile": "kubernetes",
                "group": "kube-default",
            },
            "namespace_default_endpoint_group": {
                "kube-system": {
                    "tenant": tenant,
                    "app_profile": "kubernetes",
                    "group": "kube-system",
                },
            },
            "pod_ip_pool": [
                {
                    "start": cidr_split(pod_subnet)[0],
                    "end": cidr_split(pod_subnet)[1],
                }
            ],
            "pod_network": [
                {
                    "subnet": "%s/%s" % cidr_split(pod_subnet)[3:],
                    "gateway": cidr_split(pod_subnet)[2],
                    "routes": [
                        {
                            "dst": "0.0.0.0/0",
                            "gw": cidr_split(pod_subnet)[2],
                        }
                    ],
                },
            ],
            "service_ip_pool": [
                {
                    "start": cidr_split(extern_dynamic)[0],
                    "end": cidr_split(extern_dynamic)[1],
                },
            ],
            "static_service_ip_pool": [
                {
                    "start": cidr_split(extern_static)[0],
                    "end": cidr_split(extern_static)[1],
                },
            ],
            "node_service_ip_pool": [
                {
                    "start": cidr_split(node_svc_subnet)[0],
                    "end": cidr_split(node_svc_subnet)[1],
                },
            ],
            "node_service_gw_subnets": [
                node_svc_subnet,
            ],
        },
    }
    return adj_config


def config_validate(config):
    required = lambda x: x
    get = lambda t: reduce(lambda x, y: x and x.get(y), t, config)

    checks = {
        "system_id": (get(("aci_config", "system_id")), required),
        "aep": (get(("aci_config", "aep")), required),
        "apic_host": (get(("aci_config", "apic_hosts")), required),
        "apic_username": (get(("aci_config", "apic_login", "username")), required),
        "apic_password": (get(("aci_config", "apic_login", "password")), required),
        "uplink_if": (get(("node_config", "uplink_iface")), required),
        "vxlan_if": (get(("node_config", "vxlan_uplink_iface")), required),
        "kubeapi_vlan": (get(("net_config", "kubeapi_vlan")), required),
        "service_vlan": (get(("net_config", "service_vlan")), required),
    }

    ret = True
    for k in checks:
        value, validator = checks[k]
        try:
            if not validator(value):
                raise Exception(k)
        except Exception as e:
            err("Required configuration not present or not correct: '%s'" % e.message)
            ret = False
    return ret


def config_advise(config, prov_apic):
    try:
        if prov_apic is not None:
            apic = get_apic(config)

            aep_name = config["aci_config"]["aep"]
            aep = apic.get_aep(aep_name)
            if aep is None:
                warn("AEP not defined in the APIC: %s" % aep_name)

            vrf_tenant = config["aci_config"]["vrf"]["tenant"]
            vrf_name = config["aci_config"]["vrf"]["name"]
            l3out_name = config["aci_config"]["l3out"]["name"]
            vrf = apic.get_vrf(vrf_tenant, vrf_name)
            if vrf is None:
                warn("VRF not defined in the APIC: %s/%s" %
                     (vrf_tenant, vrf_name))
            l3out = apic.get_l3out(vrf_tenant, l3out_name)
            if l3out is None:
                warn("L3out not defined in the APIC: %s/%s" %
                     (vrf_tenant, l3out_name))

    except Exception as e:
        warn("Error in validating existence of AEP: '%s'" % e.message)
    return True


def generate_sample(filep):
    with open('provision-config.yaml', 'r') as inp:
        print(inp.read(), file=filep)
    return filep


def generate_kube_yaml(config, output):
    env = Environment(
        loader=PackageLoader('aci-containers-provision', 'templates'),
        trim_blocks=True,
        lstrip_blocks=True,
    )
    env.filters['base64enc'] = base64.b64encode
    env.filters['json'] = json_indent
    env.filters['yaml_quote'] = yaml_quote
    template = env.get_template('aci-containers.yaml')

    if output:
        if output == "-":
            info("Writing kubernetes infrastructure YAML to \"STDOUT\"")
            template.stream(config=config).dump(sys.stdout)
        else:
            info("Writing kubernetes infrastructure YAML to \"%s\"" % output)
            template.stream(config=config).dump(output)
    return config


def generate_apic_config(config, prov_apic, apic_file):
    apic_config = ApicKubeConfig(config).get_config()
    if apic_file:
        if apic_file == "-":
            info("Writing apic configuration to \"STDOUT\"")
            ApicKubeConfig.save_config(apic_config, sys.stdout)
        else:
            info("Writing apic configuration to \"%s\"" % apic_file)
            with open(apic_file, 'w') as outfile:
                ApicKubeConfig.save_config(apic_config, outfile)

    if prov_apic is not None:
        apic = get_apic(config)
        if prov_apic is True:
            apic.provision(apic_config)
        if prov_apic is False:
            apic.unprovision(apic_config)
    return apic_config


def get_apic(config):
    apic_host = config["aci_config"]["apic_hosts"][0]
    apic_username = config["aci_config"]["apic_login"]["username"]
    apic_password = config["aci_config"]["apic_login"]["password"]
    apic = Apic(apic_host, apic_username, apic_password)
    return apic


def parse_args():
    parser = argparse.ArgumentParser(
        description='Provision an ACI kubernetes installation'
    )
    parser.add_argument('-c', '--config', default="-", metavar='',
                        help='Input file with your fabric configuration')
    parser.add_argument('-o', '--output', default="-", metavar='',
                        help='Output file for your kubernetes deployment')
    parser.add_argument('-a', '--apic', action='store_true', default=False,
                        help='Create/Validate the required APIC resources')
    parser.add_argument('-d', '--delete', action='store_true', default=False,
                        help='Delete the APIC resources that would have be created')
    parser.add_argument('-s', '--sample', action='store_true', default=False,
                        help='Print a sample input file with fabric configuration')
    parser.add_argument('-u', '--username', default=None, metavar='',
                        help='APIC admin username to use for APIC API access')
    parser.add_argument('-p', '--password', default=None, metavar='',
                        help='APIC admin password to use for APIC API access')
    return parser.parse_args()


def main(args, apic_file=None):
    config_file = args.config
    output_file = args.output
    prov_apic = None
    if args.apic:
        prov_apic = True
        if args.delete:
            prov_apic = False

    # Print sample, if needed
    if args.sample:
        generate_sample(sys.stdout)
        return True

    # command line config
    config = {
        "aci_config": {
            "apic_login": {
            }
        }
    }
    if args.username:
        config["aci_config"]["apic_login"]["username"] = args.username
    if args.password:
        config["aci_config"]["apic_login"]["password"] = args.password

    # Create config
    default_config = config_default()
    user_config = config_user(config_file)
    deep_merge(config, user_config)
    deep_merge(config, default_config)

    # Validate config
    if not config_validate(config):
        err("Please fix configuration and retry.")
        return False

    # Adjust config based on convention/apic data
    adj_config = config_adjust(config, prov_apic)
    deep_merge(config, adj_config)
    config["net_config"]["infra_vlan"] = \
        adj_config["net_config"]["infra_vlan"]

    # Advisory checks, including apic checks, ignore failures
    if not config_advise(config, prov_apic):
        pass

    # generate output files; and program apic if needed
    generate_apic_config(config, prov_apic, apic_file)
    generate_kube_yaml(config, output_file)
    return True


def test_main():
    arg = {
        "config": None,
        "output": None,
        "apicfile": None,
        "apic": False,
        "delete": False,
        "username": "admin",
        "password": "",
        "sample": False,
    }
    argc = collections.namedtuple('argc', arg.keys())
    args = argc(**arg)

    for inp in glob.glob("testdata/*.inp.yaml"):
        # Exec main
        args = args._replace(config=inp)
        args = args._replace(output=os.tempnam(".", "tmp-kube-"))
        apicfile = os.tempnam(".", "tmp-apic-")
        main(args, apicfile)

        # Verify generated configs
        expectedkube = inp[:-8] + 'out.yaml'
        assert filecmp.cmp(args.output, expectedkube)
        expectedapic = inp[:-8] + 'apic.txt'
        assert filecmp.cmp(apicfile, expectedapic)

        # Cleanup
        os.remove(args.output)
        os.remove(apicfile)

if __name__ == "__main__":
    args = parse_args()
    main(args)
