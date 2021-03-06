// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package firewallrule

import (
	ingressbe "k8s.io/ingress-gce/pkg/backends"
)

// FirewallRuleSyncerInterface is an interface to manage GCP firewall rules.
type FirewallRuleSyncerInterface interface {
	// EnsureFirewallRule ensures that the required firewall rules exist.
	// If a firewall rule already exists and differs, it will not be updated unless forceUpdate is true.
	EnsureFirewallRule(lbName string, ports []ingressbe.ServicePort, igLinks map[string][]string, forceUpdate bool) error
	// DeleteFirewallRules deletes all firewall rules that would have been created by EnsureFirewallRule.
	DeleteFirewallRules() error
}
