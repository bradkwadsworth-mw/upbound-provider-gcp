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

type TableIAMMemberConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TableIAMMemberConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TableIAMMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TableIAMMemberInitParameters struct {
	Condition []TableIAMMemberConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type TableIAMMemberObservation struct {
	Condition []TableIAMMemberConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	TableID *string `json:"tableId,omitempty" tf:"table_id,omitempty"`
}

type TableIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []TableIAMMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Dataset
	// +kubebuilder:validation:Optional
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// Reference to a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDRef *v1.Reference `json:"datasetIdRef,omitempty" tf:"-"`

	// Selector for a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDSelector *v1.Selector `json:"datasetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Table
	// +kubebuilder:validation:Optional
	TableID *string `json:"tableId,omitempty" tf:"table_id,omitempty"`

	// Reference to a Table in bigquery to populate tableId.
	// +kubebuilder:validation:Optional
	TableIDRef *v1.Reference `json:"tableIdRef,omitempty" tf:"-"`

	// Selector for a Table in bigquery to populate tableId.
	// +kubebuilder:validation:Optional
	TableIDSelector *v1.Selector `json:"tableIdSelector,omitempty" tf:"-"`
}

// TableIAMMemberSpec defines the desired state of TableIAMMember
type TableIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableIAMMemberParameters `json:"forProvider"`
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
	InitProvider TableIAMMemberInitParameters `json:"initProvider,omitempty"`
}

// TableIAMMemberStatus defines the observed state of TableIAMMember.
type TableIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TableIAMMember is the Schema for the TableIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TableIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || has(self.initProvider.role)",message="role is a required parameter"
	Spec   TableIAMMemberSpec   `json:"spec"`
	Status TableIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableIAMMemberList contains a list of TableIAMMembers
type TableIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TableIAMMember `json:"items"`
}

// Repository type metadata.
var (
	TableIAMMember_Kind             = "TableIAMMember"
	TableIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TableIAMMember_Kind}.String()
	TableIAMMember_KindAPIVersion   = TableIAMMember_Kind + "." + CRDGroupVersion.String()
	TableIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(TableIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&TableIAMMember{}, &TableIAMMemberList{})
}
