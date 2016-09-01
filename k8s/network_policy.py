# Copyright (c) 2015-2016 Tigera Inc.  All rights reserved.
# Copyright (c) 2016 Cisco Systems, Inc.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#  http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Controller for syncing kubernetes network policy with ACI using the
# ACI integration module.
#
# This is based on the Calico k8-spolicy controller which can be found at:
# https://github.com/projectcalico/k8s-policy

import logging
import json
import os

from constants.logging import *
# from policy_parser import PolicyParser

_log = logging.getLogger("__main__")
#client = DatastoreClient()


def add_update_network_policy(policy):
    """
    Takes a new network policy from the Kubernetes API and
    creates the corresponding Calico policy configuration.
    """
    # Determine the name for this policy.
    #name = "%s.%s" % (policy["metadata"]["namespace"],
    #                  policy["metadata"]["name"])
    #_log.debug("Adding new network policy: %s", name)
    #
    #try:
    #    parser = PolicyParser(policy)
    #    selector = parser.calculate_pod_selector()
    #    inbound_rules = parser.calculate_inbound_rules()
    #except Exception:
    #    # If the Policy is malformed, log the error and kill the controller.
    #    # Kubernetes will restart us.
    #    _log.exception("Error parsing policy: %s",
    #                   json.dumps(policy, indent=2))
    #    os.exit(1)
    #else:
    #    rules = Rules(inbound_rules=inbound_rules,
    #                  outbound_rules=[Rule(action="allow")])
    #
    #    # Create the network policy using the calculated selector and rules.
    #    client.create_policy(NET_POL_TIER_NAME,
    #                         name,
    #                         selector,
    #                         order=NET_POL_ORDER,
    #                         rules=rules)
    #    _log.debug("Updated policy '%s' for NetworkPolicy", name)
    pass


def delete_network_policy(policy):
    """
    Takes a deleted network policy and removes the corresponding
    configuration from the Calico datastore.
    """
    ## Determine the name for this policy.
    #name = "%s.%s" % (policy["metadata"]["namespace"],
    #                  policy["metadata"]["name"])
    #_log.debug("Deleting network policy: %s", name)
    #
    ## Delete the corresponding Calico policy
    #try:
    #    client.remove_policy(NET_POL_TIER_NAME, name)
    #except KeyError:
    #    _log.info("Unable to find policy '%s' - already deleted", name)
    pass