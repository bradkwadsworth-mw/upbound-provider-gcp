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

type DefaultObjectAccessControlInitParameters struct {

	// The entity holding the permission, in one of the following forms:
	Entity *string `json:"entity,omitempty" tf:"entity,omitempty"`

	// The name of the object, if applied to an object.
	Object *string `json:"object,omitempty" tf:"object,omitempty"`

	// The access permission for the entity.
	// Possible values are: OWNER, READER.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type DefaultObjectAccessControlObservation struct {

	// The name of the bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The domain associated with the entity.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The email address associated with the entity.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The entity holding the permission, in one of the following forms:
	Entity *string `json:"entity,omitempty" tf:"entity,omitempty"`

	// The ID for the entity
	EntityID *string `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// The content generation of the object, if applied to an object.
	Generation *float64 `json:"generation,omitempty" tf:"generation,omitempty"`

	// an identifier for the resource with format {{bucket}}/{{entity}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the object, if applied to an object.
	Object *string `json:"object,omitempty" tf:"object,omitempty"`

	// The project team associated with the entity
	// Structure is documented below.
	ProjectTeam []ProjectTeamObservation `json:"projectTeam,omitempty" tf:"project_team,omitempty"`

	// The access permission for the entity.
	// Possible values are: OWNER, READER.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type DefaultObjectAccessControlParameters struct {

	// The name of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/storage/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in storage to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in storage to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The entity holding the permission, in one of the following forms:
	// +kubebuilder:validation:Optional
	Entity *string `json:"entity,omitempty" tf:"entity,omitempty"`

	// The name of the object, if applied to an object.
	// +kubebuilder:validation:Optional
	Object *string `json:"object,omitempty" tf:"object,omitempty"`

	// The access permission for the entity.
	// Possible values are: OWNER, READER.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type ProjectTeamInitParameters struct {
}

type ProjectTeamObservation struct {

	// The project team associated with the entity
	ProjectNumber *string `json:"projectNumber,omitempty" tf:"project_number,omitempty"`

	// The team.
	// Possible values are: editors, owners, viewers.
	Team *string `json:"team,omitempty" tf:"team,omitempty"`
}

type ProjectTeamParameters struct {
}

// DefaultObjectAccessControlSpec defines the desired state of DefaultObjectAccessControl
type DefaultObjectAccessControlSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DefaultObjectAccessControlParameters `json:"forProvider"`
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
	InitProvider DefaultObjectAccessControlInitParameters `json:"initProvider,omitempty"`
}

// DefaultObjectAccessControlStatus defines the observed state of DefaultObjectAccessControl.
type DefaultObjectAccessControlStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DefaultObjectAccessControlObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultObjectAccessControl is the Schema for the DefaultObjectAccessControls API. The DefaultObjectAccessControls resources represent the Access Control Lists (ACLs) applied to a new object within a Google Cloud Storage bucket when no ACL was provided for that object.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type DefaultObjectAccessControl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.entity) || has(self.initProvider.entity)",message="entity is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || has(self.initProvider.role)",message="role is a required parameter"
	Spec   DefaultObjectAccessControlSpec   `json:"spec"`
	Status DefaultObjectAccessControlStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultObjectAccessControlList contains a list of DefaultObjectAccessControls
type DefaultObjectAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultObjectAccessControl `json:"items"`
}

// Repository type metadata.
var (
	DefaultObjectAccessControl_Kind             = "DefaultObjectAccessControl"
	DefaultObjectAccessControl_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DefaultObjectAccessControl_Kind}.String()
	DefaultObjectAccessControl_KindAPIVersion   = DefaultObjectAccessControl_Kind + "." + CRDGroupVersion.String()
	DefaultObjectAccessControl_GroupVersionKind = CRDGroupVersion.WithKind(DefaultObjectAccessControl_Kind)
)

func init() {
	SchemeBuilder.Register(&DefaultObjectAccessControl{}, &DefaultObjectAccessControlList{})
}
