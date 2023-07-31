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

type BackupBackupPlanInitParameters struct {

	// Defines the configuration of Backups created via this BackupPlan.
	// Structure is documented below.
	BackupConfig []BackupConfigInitParameters `json:"backupConfig,omitempty" tf:"backup_config,omitempty"`

	// Defines a schedule for automatic Backup creation via this BackupPlan.
	// Structure is documented below.
	BackupSchedule []BackupScheduleInitParameters `json:"backupSchedule,omitempty" tf:"backup_schedule,omitempty"`

	// This flag indicates whether this BackupPlan has been deactivated.
	// Setting this field to True locks the BackupPlan such that no further updates will be allowed
	// (except deletes), including the deactivated field itself. It also prevents any new Backups
	// from being created via this BackupPlan (including scheduled Backups).
	Deactivated *bool `json:"deactivated,omitempty" tf:"deactivated,omitempty"`

	// User specified descriptive string for this BackupPlan.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Description: A set of custom labels supplied by the user.
	// A list of key->value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// RetentionPolicy governs lifecycle of Backups created under this plan.
	// Structure is documented below.
	RetentionPolicy []RetentionPolicyInitParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`
}

type BackupBackupPlanObservation struct {

	// Defines the configuration of Backups created via this BackupPlan.
	// Structure is documented below.
	BackupConfig []BackupConfigObservation `json:"backupConfig,omitempty" tf:"backup_config,omitempty"`

	// Defines a schedule for automatic Backup creation via this BackupPlan.
	// Structure is documented below.
	BackupSchedule []BackupScheduleObservation `json:"backupSchedule,omitempty" tf:"backup_schedule,omitempty"`

	// The source cluster from which Backups will be created via this BackupPlan.
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// This flag indicates whether this BackupPlan has been deactivated.
	// Setting this field to True locks the BackupPlan such that no further updates will be allowed
	// (except deletes), including the deactivated field itself. It also prevents any new Backups
	// from being created via this BackupPlan (including scheduled Backups).
	Deactivated *bool `json:"deactivated,omitempty" tf:"deactivated,omitempty"`

	// User specified descriptive string for this BackupPlan.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// etag is used for optimistic concurrency control as a way to help prevent simultaneous
	// updates of a backup plan from overwriting each other. It is strongly suggested that
	// systems make use of the 'etag' in the read-modify-write cycle to perform BackupPlan updates
	// in order to avoid race conditions: An etag is returned in the response to backupPlans.get,
	// and systems are expected to put that etag in the request to backupPlans.patch or
	// backupPlans.delete to ensure that their change will be applied to the same version of the resource.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/backupPlans/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Description: A set of custom labels supplied by the user.
	// A list of key->value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The region of the Backup Plan.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The number of Kubernetes Pods backed up in the last successful Backup created via this BackupPlan.
	ProtectedPodCount *float64 `json:"protectedPodCount,omitempty" tf:"protected_pod_count,omitempty"`

	// RetentionPolicy governs lifecycle of Backups created under this plan.
	// Structure is documented below.
	RetentionPolicy []RetentionPolicyObservation `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// Server generated, unique identifier of UUID format.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type BackupBackupPlanParameters struct {

	// Defines the configuration of Backups created via this BackupPlan.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	BackupConfig []BackupConfigParameters `json:"backupConfig,omitempty" tf:"backup_config,omitempty"`

	// Defines a schedule for automatic Backup creation via this BackupPlan.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	BackupSchedule []BackupScheduleParameters `json:"backupSchedule,omitempty" tf:"backup_schedule,omitempty"`

	// The source cluster from which Backups will be created via this BackupPlan.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/container/v1beta1.Cluster
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// Reference to a Cluster in container to populate cluster.
	// +kubebuilder:validation:Optional
	ClusterRef *v1.Reference `json:"clusterRef,omitempty" tf:"-"`

	// Selector for a Cluster in container to populate cluster.
	// +kubebuilder:validation:Optional
	ClusterSelector *v1.Selector `json:"clusterSelector,omitempty" tf:"-"`

	// This flag indicates whether this BackupPlan has been deactivated.
	// Setting this field to True locks the BackupPlan such that no further updates will be allowed
	// (except deletes), including the deactivated field itself. It also prevents any new Backups
	// from being created via this BackupPlan (including scheduled Backups).
	// +kubebuilder:validation:Optional
	Deactivated *bool `json:"deactivated,omitempty" tf:"deactivated,omitempty"`

	// User specified descriptive string for this BackupPlan.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Description: A set of custom labels supplied by the user.
	// A list of key->value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The region of the Backup Plan.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// RetentionPolicy governs lifecycle of Backups created under this plan.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	RetentionPolicy []RetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`
}

type BackupConfigInitParameters struct {

	// If True, include all namespaced resources.
	AllNamespaces *bool `json:"allNamespaces,omitempty" tf:"all_namespaces,omitempty"`

	// This defines a customer managed encryption key that will be used to encrypt the "config"
	// portion (the Kubernetes resources) of Backups created via this plan.
	// Structure is documented below.
	EncryptionKey []EncryptionKeyInitParameters `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`

	// This flag specifies whether Kubernetes Secret resources should be included
	// when they fall into the scope of Backups.
	IncludeSecrets *bool `json:"includeSecrets,omitempty" tf:"include_secrets,omitempty"`

	// This flag specifies whether volume data should be backed up when PVCs are
	// included in the scope of a Backup.
	IncludeVolumeData *bool `json:"includeVolumeData,omitempty" tf:"include_volume_data,omitempty"`

	// A list of namespaced Kubernetes Resources.
	// Structure is documented below.
	SelectedApplications []SelectedApplicationsInitParameters `json:"selectedApplications,omitempty" tf:"selected_applications,omitempty"`

	// If set, include just the resources in the listed namespaces.
	// Structure is documented below.
	SelectedNamespaces []SelectedNamespacesInitParameters `json:"selectedNamespaces,omitempty" tf:"selected_namespaces,omitempty"`
}

type BackupConfigObservation struct {

	// If True, include all namespaced resources.
	AllNamespaces *bool `json:"allNamespaces,omitempty" tf:"all_namespaces,omitempty"`

	// This defines a customer managed encryption key that will be used to encrypt the "config"
	// portion (the Kubernetes resources) of Backups created via this plan.
	// Structure is documented below.
	EncryptionKey []EncryptionKeyObservation `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`

	// This flag specifies whether Kubernetes Secret resources should be included
	// when they fall into the scope of Backups.
	IncludeSecrets *bool `json:"includeSecrets,omitempty" tf:"include_secrets,omitempty"`

	// This flag specifies whether volume data should be backed up when PVCs are
	// included in the scope of a Backup.
	IncludeVolumeData *bool `json:"includeVolumeData,omitempty" tf:"include_volume_data,omitempty"`

	// A list of namespaced Kubernetes Resources.
	// Structure is documented below.
	SelectedApplications []SelectedApplicationsObservation `json:"selectedApplications,omitempty" tf:"selected_applications,omitempty"`

	// If set, include just the resources in the listed namespaces.
	// Structure is documented below.
	SelectedNamespaces []SelectedNamespacesObservation `json:"selectedNamespaces,omitempty" tf:"selected_namespaces,omitempty"`
}

type BackupConfigParameters struct {

	// If True, include all namespaced resources.
	// +kubebuilder:validation:Optional
	AllNamespaces *bool `json:"allNamespaces,omitempty" tf:"all_namespaces,omitempty"`

	// This defines a customer managed encryption key that will be used to encrypt the "config"
	// portion (the Kubernetes resources) of Backups created via this plan.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EncryptionKey []EncryptionKeyParameters `json:"encryptionKey,omitempty" tf:"encryption_key,omitempty"`

	// This flag specifies whether Kubernetes Secret resources should be included
	// when they fall into the scope of Backups.
	// +kubebuilder:validation:Optional
	IncludeSecrets *bool `json:"includeSecrets,omitempty" tf:"include_secrets,omitempty"`

	// This flag specifies whether volume data should be backed up when PVCs are
	// included in the scope of a Backup.
	// +kubebuilder:validation:Optional
	IncludeVolumeData *bool `json:"includeVolumeData,omitempty" tf:"include_volume_data,omitempty"`

	// A list of namespaced Kubernetes Resources.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SelectedApplications []SelectedApplicationsParameters `json:"selectedApplications,omitempty" tf:"selected_applications,omitempty"`

	// If set, include just the resources in the listed namespaces.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SelectedNamespaces []SelectedNamespacesParameters `json:"selectedNamespaces,omitempty" tf:"selected_namespaces,omitempty"`
}

type BackupScheduleInitParameters struct {

	// A standard cron string that defines a repeating schedule for
	// creating Backups via this BackupPlan.
	// If this is defined, then backupRetainDays must also be defined.
	CronSchedule *string `json:"cronSchedule,omitempty" tf:"cron_schedule,omitempty"`

	// This flag denotes whether automatic Backup creation is paused for this BackupPlan.
	Paused *bool `json:"paused,omitempty" tf:"paused,omitempty"`
}

type BackupScheduleObservation struct {

	// A standard cron string that defines a repeating schedule for
	// creating Backups via this BackupPlan.
	// If this is defined, then backupRetainDays must also be defined.
	CronSchedule *string `json:"cronSchedule,omitempty" tf:"cron_schedule,omitempty"`

	// This flag denotes whether automatic Backup creation is paused for this BackupPlan.
	Paused *bool `json:"paused,omitempty" tf:"paused,omitempty"`
}

type BackupScheduleParameters struct {

	// A standard cron string that defines a repeating schedule for
	// creating Backups via this BackupPlan.
	// If this is defined, then backupRetainDays must also be defined.
	// +kubebuilder:validation:Optional
	CronSchedule *string `json:"cronSchedule,omitempty" tf:"cron_schedule,omitempty"`

	// This flag denotes whether automatic Backup creation is paused for this BackupPlan.
	// +kubebuilder:validation:Optional
	Paused *bool `json:"paused,omitempty" tf:"paused,omitempty"`
}

type EncryptionKeyInitParameters struct {
}

type EncryptionKeyObservation struct {

	// Google Cloud KMS encryption key. Format: projects//locations//keyRings//cryptoKeys/
	GCPKMSEncryptionKey *string `json:"gcpKmsEncryptionKey,omitempty" tf:"gcp_kms_encryption_key,omitempty"`
}

type EncryptionKeyParameters struct {

	// Google Cloud KMS encryption key. Format: projects//locations//keyRings//cryptoKeys/
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta1.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GCPKMSEncryptionKey *string `json:"gcpKmsEncryptionKey,omitempty" tf:"gcp_kms_encryption_key,omitempty"`

	// Reference to a CryptoKey in kms to populate gcpKmsEncryptionKey.
	// +kubebuilder:validation:Optional
	GCPKMSEncryptionKeyRef *v1.Reference `json:"gcpKmsEncryptionKeyRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate gcpKmsEncryptionKey.
	// +kubebuilder:validation:Optional
	GCPKMSEncryptionKeySelector *v1.Selector `json:"gcpKmsEncryptionKeySelector,omitempty" tf:"-"`
}

type NamespacedNamesInitParameters struct {

	// The name of a Kubernetes Resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace of a Kubernetes Resource.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type NamespacedNamesObservation struct {

	// The name of a Kubernetes Resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace of a Kubernetes Resource.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type NamespacedNamesParameters struct {

	// The name of a Kubernetes Resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace of a Kubernetes Resource.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type RetentionPolicyInitParameters struct {

	// Minimum age for a Backup created via this BackupPlan (in days).
	// Must be an integer value between 0-90 (inclusive).
	// A Backup created under this BackupPlan will not be deletable
	// until it reaches Backup's (create time + backup_delete_lock_days).
	// Updating this field of a BackupPlan does not affect existing Backups.
	// Backups created after a successful update will inherit this new value.
	BackupDeleteLockDays *float64 `json:"backupDeleteLockDays,omitempty" tf:"backup_delete_lock_days,omitempty"`

	// The default maximum age of a Backup created via this BackupPlan.
	// This field MUST be an integer value >= 0 and <= 365. If specified,
	// a Backup created under this BackupPlan will be automatically deleted
	// after its age reaches (createTime + backupRetainDays).
	// If not specified, Backups created under this BackupPlan will NOT be
	// subject to automatic deletion. Updating this field does NOT affect
	// existing Backups under it. Backups created AFTER a successful update
	// will automatically pick up the new value.
	// NOTE: backupRetainDays must be >= backupDeleteLockDays.
	// If cronSchedule is defined, then this must be <= 360 * the creation interval.]
	BackupRetainDays *float64 `json:"backupRetainDays,omitempty" tf:"backup_retain_days,omitempty"`

	// This flag denotes whether the retention policy of this BackupPlan is locked.
	// If set to True, no further update is allowed on this policy, including
	// the locked field itself.
	Locked *bool `json:"locked,omitempty" tf:"locked,omitempty"`
}

type RetentionPolicyObservation struct {

	// Minimum age for a Backup created via this BackupPlan (in days).
	// Must be an integer value between 0-90 (inclusive).
	// A Backup created under this BackupPlan will not be deletable
	// until it reaches Backup's (create time + backup_delete_lock_days).
	// Updating this field of a BackupPlan does not affect existing Backups.
	// Backups created after a successful update will inherit this new value.
	BackupDeleteLockDays *float64 `json:"backupDeleteLockDays,omitempty" tf:"backup_delete_lock_days,omitempty"`

	// The default maximum age of a Backup created via this BackupPlan.
	// This field MUST be an integer value >= 0 and <= 365. If specified,
	// a Backup created under this BackupPlan will be automatically deleted
	// after its age reaches (createTime + backupRetainDays).
	// If not specified, Backups created under this BackupPlan will NOT be
	// subject to automatic deletion. Updating this field does NOT affect
	// existing Backups under it. Backups created AFTER a successful update
	// will automatically pick up the new value.
	// NOTE: backupRetainDays must be >= backupDeleteLockDays.
	// If cronSchedule is defined, then this must be <= 360 * the creation interval.]
	BackupRetainDays *float64 `json:"backupRetainDays,omitempty" tf:"backup_retain_days,omitempty"`

	// This flag denotes whether the retention policy of this BackupPlan is locked.
	// If set to True, no further update is allowed on this policy, including
	// the locked field itself.
	Locked *bool `json:"locked,omitempty" tf:"locked,omitempty"`
}

type RetentionPolicyParameters struct {

	// Minimum age for a Backup created via this BackupPlan (in days).
	// Must be an integer value between 0-90 (inclusive).
	// A Backup created under this BackupPlan will not be deletable
	// until it reaches Backup's (create time + backup_delete_lock_days).
	// Updating this field of a BackupPlan does not affect existing Backups.
	// Backups created after a successful update will inherit this new value.
	// +kubebuilder:validation:Optional
	BackupDeleteLockDays *float64 `json:"backupDeleteLockDays,omitempty" tf:"backup_delete_lock_days,omitempty"`

	// The default maximum age of a Backup created via this BackupPlan.
	// This field MUST be an integer value >= 0 and <= 365. If specified,
	// a Backup created under this BackupPlan will be automatically deleted
	// after its age reaches (createTime + backupRetainDays).
	// If not specified, Backups created under this BackupPlan will NOT be
	// subject to automatic deletion. Updating this field does NOT affect
	// existing Backups under it. Backups created AFTER a successful update
	// will automatically pick up the new value.
	// NOTE: backupRetainDays must be >= backupDeleteLockDays.
	// If cronSchedule is defined, then this must be <= 360 * the creation interval.]
	// +kubebuilder:validation:Optional
	BackupRetainDays *float64 `json:"backupRetainDays,omitempty" tf:"backup_retain_days,omitempty"`

	// This flag denotes whether the retention policy of this BackupPlan is locked.
	// If set to True, no further update is allowed on this policy, including
	// the locked field itself.
	// +kubebuilder:validation:Optional
	Locked *bool `json:"locked,omitempty" tf:"locked,omitempty"`
}

type SelectedApplicationsInitParameters struct {

	// A list of namespaced Kubernetes resources.
	// Structure is documented below.
	NamespacedNames []NamespacedNamesInitParameters `json:"namespacedNames,omitempty" tf:"namespaced_names,omitempty"`
}

type SelectedApplicationsObservation struct {

	// A list of namespaced Kubernetes resources.
	// Structure is documented below.
	NamespacedNames []NamespacedNamesObservation `json:"namespacedNames,omitempty" tf:"namespaced_names,omitempty"`
}

type SelectedApplicationsParameters struct {

	// A list of namespaced Kubernetes resources.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	NamespacedNames []NamespacedNamesParameters `json:"namespacedNames,omitempty" tf:"namespaced_names,omitempty"`
}

type SelectedNamespacesInitParameters struct {

	// A list of Kubernetes Namespaces.
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`
}

type SelectedNamespacesObservation struct {

	// A list of Kubernetes Namespaces.
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`
}

type SelectedNamespacesParameters struct {

	// A list of Kubernetes Namespaces.
	// +kubebuilder:validation:Optional
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`
}

// BackupBackupPlanSpec defines the desired state of BackupBackupPlan
type BackupBackupPlanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupBackupPlanParameters `json:"forProvider"`
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
	InitProvider BackupBackupPlanInitParameters `json:"initProvider,omitempty"`
}

// BackupBackupPlanStatus defines the observed state of BackupBackupPlan.
type BackupBackupPlanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupBackupPlanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupBackupPlan is the Schema for the BackupBackupPlans API. Represents a Backup Plan instance.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type BackupBackupPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupBackupPlanSpec   `json:"spec"`
	Status            BackupBackupPlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupBackupPlanList contains a list of BackupBackupPlans
type BackupBackupPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupBackupPlan `json:"items"`
}

// Repository type metadata.
var (
	BackupBackupPlan_Kind             = "BackupBackupPlan"
	BackupBackupPlan_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupBackupPlan_Kind}.String()
	BackupBackupPlan_KindAPIVersion   = BackupBackupPlan_Kind + "." + CRDGroupVersion.String()
	BackupBackupPlan_GroupVersionKind = CRDGroupVersion.WithKind(BackupBackupPlan_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupBackupPlan{}, &BackupBackupPlanList{})
}
