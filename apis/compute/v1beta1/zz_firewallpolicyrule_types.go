/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FirewallPolicyRuleInitParameters struct {

	// The Action to perform when the client connection triggers the rule. Can currently be either "allow" or "deny()" where valid values for status are 403, 404, and 502.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// An optional description for this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging,omitempty"`

	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced. Structure is documented below.
	Match []MatchInitParameters `json:"match,omitempty" tf:"match,omitempty"`

	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
	TargetResources []*string `json:"targetResources,omitempty" tf:"target_resources,omitempty"`

	// A list of service accounts indicating the sets of instances that are applied with this rule.
	TargetServiceAccounts []*string `json:"targetServiceAccounts,omitempty" tf:"target_service_accounts,omitempty"`
}

type FirewallPolicyRuleObservation struct {

	// The Action to perform when the client connection triggers the rule. Can currently be either "allow" or "deny()" where valid values for status are 403, 404, and 502.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// An optional description for this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging,omitempty"`

	// The firewall policy of the resource.
	FirewallPolicy *string `json:"firewallPolicy,omitempty" tf:"firewall_policy,omitempty"`

	// an identifier for the resource with format locations/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of the resource. Always compute#firewallPolicyRule for firewall policy rules
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced. Structure is documented below.
	Match []MatchObservation `json:"match,omitempty" tf:"match,omitempty"`

	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Calculation of the complexity of a single firewall policy rule.
	RuleTupleCount *float64 `json:"ruleTupleCount,omitempty" tf:"rule_tuple_count,omitempty"`

	// A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
	TargetResources []*string `json:"targetResources,omitempty" tf:"target_resources,omitempty"`

	// A list of service accounts indicating the sets of instances that are applied with this rule.
	TargetServiceAccounts []*string `json:"targetServiceAccounts,omitempty" tf:"target_service_accounts,omitempty"`
}

type FirewallPolicyRuleParameters struct {

	// The Action to perform when the client connection triggers the rule. Can currently be either "allow" or "deny()" where valid values for status are 403, 404, and 502.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// An optional description for this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
	// +kubebuilder:validation:Optional
	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging,omitempty"`

	// The firewall policy of the resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.FirewallPolicy
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	FirewallPolicy *string `json:"firewallPolicy,omitempty" tf:"firewall_policy,omitempty"`

	// Reference to a FirewallPolicy in compute to populate firewallPolicy.
	// +kubebuilder:validation:Optional
	FirewallPolicyRef *v1.Reference `json:"firewallPolicyRef,omitempty" tf:"-"`

	// Selector for a FirewallPolicy in compute to populate firewallPolicy.
	// +kubebuilder:validation:Optional
	FirewallPolicySelector *v1.Selector `json:"firewallPolicySelector,omitempty" tf:"-"`

	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced. Structure is documented below.
	// +kubebuilder:validation:Optional
	Match []MatchParameters `json:"match,omitempty" tf:"match,omitempty"`

	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
	// +kubebuilder:validation:Optional
	TargetResources []*string `json:"targetResources,omitempty" tf:"target_resources,omitempty"`

	// A list of service accounts indicating the sets of instances that are applied with this rule.
	// +kubebuilder:validation:Optional
	TargetServiceAccounts []*string `json:"targetServiceAccounts,omitempty" tf:"target_service_accounts,omitempty"`
}

type Layer4ConfigsInitParameters struct {

	// The IP protocol to which this rule applies. The protocol type is required when creating a firewall rule. This value can either be one of the following well known protocol strings (tcp, udp, icmp, esp, ah, ipip, sctp), or the IP protocol number.
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// An optional list of ports to which this rule applies. This field is only applicable for UDP or TCP protocol. Each entry must be either an integer or a range. If not specified, this rule applies to connections through any port.
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`
}

type Layer4ConfigsObservation struct {

	// The IP protocol to which this rule applies. The protocol type is required when creating a firewall rule. This value can either be one of the following well known protocol strings (tcp, udp, icmp, esp, ah, ipip, sctp), or the IP protocol number.
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// An optional list of ports to which this rule applies. This field is only applicable for UDP or TCP protocol. Each entry must be either an integer or a range. If not specified, this rule applies to connections through any port.
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`
}

type Layer4ConfigsParameters struct {

	// The IP protocol to which this rule applies. The protocol type is required when creating a firewall rule. This value can either be one of the following well known protocol strings (tcp, udp, icmp, esp, ah, ipip, sctp), or the IP protocol number.
	// +kubebuilder:validation:Optional
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// An optional list of ports to which this rule applies. This field is only applicable for UDP or TCP protocol. Each entry must be either an integer or a range. If not specified, this rule applies to connections through any port.
	// +kubebuilder:validation:Optional
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`
}

type MatchInitParameters struct {
	DestAddressGroups []*string `json:"destAddressGroups,omitempty" tf:"dest_address_groups,omitempty"`

	DestFqdns []*string `json:"destFqdns,omitempty" tf:"dest_fqdns,omitempty"`

	// CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 256.
	DestIPRanges []*string `json:"destIpRanges,omitempty" tf:"dest_ip_ranges,omitempty"`

	DestRegionCodes []*string `json:"destRegionCodes,omitempty" tf:"dest_region_codes,omitempty"`

	DestThreatIntelligences []*string `json:"destThreatIntelligences,omitempty" tf:"dest_threat_intelligences,omitempty"`

	// Pairs of IP protocols and ports that the rule should match. Structure is documented below.
	Layer4Configs []Layer4ConfigsInitParameters `json:"layer4Configs,omitempty" tf:"layer4_configs,omitempty"`

	SrcAddressGroups []*string `json:"srcAddressGroups,omitempty" tf:"src_address_groups,omitempty"`

	SrcFqdns []*string `json:"srcFqdns,omitempty" tf:"src_fqdns,omitempty"`

	// CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 256.
	SrcIPRanges []*string `json:"srcIpRanges,omitempty" tf:"src_ip_ranges,omitempty"`

	SrcRegionCodes []*string `json:"srcRegionCodes,omitempty" tf:"src_region_codes,omitempty"`

	SrcThreatIntelligences []*string `json:"srcThreatIntelligences,omitempty" tf:"src_threat_intelligences,omitempty"`
}

type MatchObservation struct {
	DestAddressGroups []*string `json:"destAddressGroups,omitempty" tf:"dest_address_groups,omitempty"`

	DestFqdns []*string `json:"destFqdns,omitempty" tf:"dest_fqdns,omitempty"`

	// CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 256.
	DestIPRanges []*string `json:"destIpRanges,omitempty" tf:"dest_ip_ranges,omitempty"`

	DestRegionCodes []*string `json:"destRegionCodes,omitempty" tf:"dest_region_codes,omitempty"`

	DestThreatIntelligences []*string `json:"destThreatIntelligences,omitempty" tf:"dest_threat_intelligences,omitempty"`

	// Pairs of IP protocols and ports that the rule should match. Structure is documented below.
	Layer4Configs []Layer4ConfigsObservation `json:"layer4Configs,omitempty" tf:"layer4_configs,omitempty"`

	SrcAddressGroups []*string `json:"srcAddressGroups,omitempty" tf:"src_address_groups,omitempty"`

	SrcFqdns []*string `json:"srcFqdns,omitempty" tf:"src_fqdns,omitempty"`

	// CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 256.
	SrcIPRanges []*string `json:"srcIpRanges,omitempty" tf:"src_ip_ranges,omitempty"`

	SrcRegionCodes []*string `json:"srcRegionCodes,omitempty" tf:"src_region_codes,omitempty"`

	SrcThreatIntelligences []*string `json:"srcThreatIntelligences,omitempty" tf:"src_threat_intelligences,omitempty"`
}

type MatchParameters struct {

	// +kubebuilder:validation:Optional
	DestAddressGroups []*string `json:"destAddressGroups,omitempty" tf:"dest_address_groups,omitempty"`

	// +kubebuilder:validation:Optional
	DestFqdns []*string `json:"destFqdns,omitempty" tf:"dest_fqdns,omitempty"`

	// CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 256.
	// +kubebuilder:validation:Optional
	DestIPRanges []*string `json:"destIpRanges,omitempty" tf:"dest_ip_ranges,omitempty"`

	// +kubebuilder:validation:Optional
	DestRegionCodes []*string `json:"destRegionCodes,omitempty" tf:"dest_region_codes,omitempty"`

	// +kubebuilder:validation:Optional
	DestThreatIntelligences []*string `json:"destThreatIntelligences,omitempty" tf:"dest_threat_intelligences,omitempty"`

	// Pairs of IP protocols and ports that the rule should match. Structure is documented below.
	// +kubebuilder:validation:Optional
	Layer4Configs []Layer4ConfigsParameters `json:"layer4Configs,omitempty" tf:"layer4_configs,omitempty"`

	// +kubebuilder:validation:Optional
	SrcAddressGroups []*string `json:"srcAddressGroups,omitempty" tf:"src_address_groups,omitempty"`

	// +kubebuilder:validation:Optional
	SrcFqdns []*string `json:"srcFqdns,omitempty" tf:"src_fqdns,omitempty"`

	// CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 256.
	// +kubebuilder:validation:Optional
	SrcIPRanges []*string `json:"srcIpRanges,omitempty" tf:"src_ip_ranges,omitempty"`

	// +kubebuilder:validation:Optional
	SrcRegionCodes []*string `json:"srcRegionCodes,omitempty" tf:"src_region_codes,omitempty"`

	// +kubebuilder:validation:Optional
	SrcThreatIntelligences []*string `json:"srcThreatIntelligences,omitempty" tf:"src_threat_intelligences,omitempty"`
}

// FirewallPolicyRuleSpec defines the desired state of FirewallPolicyRule
type FirewallPolicyRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallPolicyRuleParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider FirewallPolicyRuleInitParameters `json:"initProvider,omitempty"`
}

// FirewallPolicyRuleStatus defines the observed state of FirewallPolicyRule.
type FirewallPolicyRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallPolicyRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallPolicyRule is the Schema for the FirewallPolicyRules API. Specific rules to add to a hierarchical firewall policy
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type FirewallPolicyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || has(self.initProvider.action)",message="action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.direction) || has(self.initProvider.direction)",message="direction is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.match) || has(self.initProvider.match)",message="match is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.priority) || has(self.initProvider.priority)",message="priority is a required parameter"
	Spec   FirewallPolicyRuleSpec   `json:"spec"`
	Status FirewallPolicyRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallPolicyRuleList contains a list of FirewallPolicyRules
type FirewallPolicyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallPolicyRule `json:"items"`
}

// Repository type metadata.
var (
	FirewallPolicyRule_Kind             = "FirewallPolicyRule"
	FirewallPolicyRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallPolicyRule_Kind}.String()
	FirewallPolicyRule_KindAPIVersion   = FirewallPolicyRule_Kind + "." + CRDGroupVersion.String()
	FirewallPolicyRule_GroupVersionKind = CRDGroupVersion.WithKind(FirewallPolicyRule_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallPolicyRule{}, &FirewallPolicyRuleList{})
}
