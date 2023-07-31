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

type ActionInitParameters struct {

	// The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// The type of the action of this Lifecycle Rule. Supported values include: Delete, SetStorageClass and AbortIncompleteMultipartUpload.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ActionObservation struct {

	// The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// The type of the action of this Lifecycle Rule. Supported values include: Delete, SetStorageClass and AbortIncompleteMultipartUpload.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ActionParameters struct {

	// The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// The type of the action of this Lifecycle Rule. Supported values include: Delete, SetStorageClass and AbortIncompleteMultipartUpload.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AutoclassInitParameters struct {

	// While set to true, autoclass automatically transitions objects in your bucket to appropriate storage classes based on each object's access pattern.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AutoclassObservation struct {

	// While set to true, autoclass automatically transitions objects in your bucket to appropriate storage classes based on each object's access pattern.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AutoclassParameters struct {

	// While set to true, autoclass automatically transitions objects in your bucket to appropriate storage classes based on each object's access pattern.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type BucketInitParameters struct {

	// The bucket's Autoclass configuration.  Structure is documented below.
	Autoclass []AutoclassInitParameters `json:"autoclass,omitempty" tf:"autoclass,omitempty"`

	// The bucket's Cross-Origin Resource Sharing (CORS) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	Cors []CorsInitParameters `json:"cors,omitempty" tf:"cors,omitempty"`

	// The bucket's custom location configuration, which specifies the individual regions that comprise a dual-region bucket. If the bucket is designated a single or multi-region, the parameters are empty. Structure is documented below.
	CustomPlacementConfig []CustomPlacementConfigInitParameters `json:"customPlacementConfig,omitempty" tf:"custom_placement_config,omitempty"`

	// Whether or not to automatically apply an eventBasedHold to new objects added to the bucket.
	DefaultEventBasedHold *bool `json:"defaultEventBasedHold,omitempty" tf:"default_event_based_hold,omitempty"`

	// The bucket's encryption configuration. Structure is documented below.
	Encryption []EncryptionInitParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// When deleting a bucket, this
	// boolean option will delete all contained objects.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// A map of key/value label pairs to assign to the bucket.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The bucket's Lifecycle Rules configuration. Multiple blocks of this type are permitted. Structure is documented below.
	LifecycleRule []LifecycleRuleInitParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// The GCS location.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The bucket's Access & Storage Logs configuration. Structure is documented below.
	Logging []LoggingInitParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Prevents public access to a bucket. Acceptable values are "inherited" or "enforced". If "inherited", the bucket uses public access prevention. only if the bucket is subject to the public access prevention organization policy constraint. Defaults to "inherited".
	PublicAccessPrevention *string `json:"publicAccessPrevention,omitempty" tf:"public_access_prevention,omitempty"`

	// Enables Requester Pays on a storage bucket.
	RequesterPays *bool `json:"requesterPays,omitempty" tf:"requester_pays,omitempty"`

	// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
	RetentionPolicy []RetentionPolicyInitParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// Enables Uniform bucket-level access access to a bucket.
	UniformBucketLevelAccess *bool `json:"uniformBucketLevelAccess,omitempty" tf:"uniform_bucket_level_access,omitempty"`

	// The bucket's Versioning configuration.  Structure is documented below.
	Versioning []VersioningInitParameters `json:"versioning,omitempty" tf:"versioning,omitempty"`

	// Configuration if the bucket acts as a website. Structure is documented below.
	Website []WebsiteInitParameters `json:"website,omitempty" tf:"website,omitempty"`
}

type BucketObservation struct {

	// The bucket's Autoclass configuration.  Structure is documented below.
	Autoclass []AutoclassObservation `json:"autoclass,omitempty" tf:"autoclass,omitempty"`

	// The bucket's Cross-Origin Resource Sharing (CORS) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	Cors []CorsObservation `json:"cors,omitempty" tf:"cors,omitempty"`

	// The bucket's custom location configuration, which specifies the individual regions that comprise a dual-region bucket. If the bucket is designated a single or multi-region, the parameters are empty. Structure is documented below.
	CustomPlacementConfig []CustomPlacementConfigObservation `json:"customPlacementConfig,omitempty" tf:"custom_placement_config,omitempty"`

	// Whether or not to automatically apply an eventBasedHold to new objects added to the bucket.
	DefaultEventBasedHold *bool `json:"defaultEventBasedHold,omitempty" tf:"default_event_based_hold,omitempty"`

	// The bucket's encryption configuration. Structure is documented below.
	Encryption []EncryptionObservation `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// When deleting a bucket, this
	// boolean option will delete all contained objects.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of key/value label pairs to assign to the bucket.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The bucket's Lifecycle Rules configuration. Multiple blocks of this type are permitted. Structure is documented below.
	LifecycleRule []LifecycleRuleObservation `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// The GCS location.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The bucket's Access & Storage Logs configuration. Structure is documented below.
	Logging []LoggingObservation `json:"logging,omitempty" tf:"logging,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Prevents public access to a bucket. Acceptable values are "inherited" or "enforced". If "inherited", the bucket uses public access prevention. only if the bucket is subject to the public access prevention organization policy constraint. Defaults to "inherited".
	PublicAccessPrevention *string `json:"publicAccessPrevention,omitempty" tf:"public_access_prevention,omitempty"`

	// Enables Requester Pays on a storage bucket.
	RequesterPays *bool `json:"requesterPays,omitempty" tf:"requester_pays,omitempty"`

	// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
	RetentionPolicy []RetentionPolicyObservation `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// The base URL of the bucket, in the format gs://<bucket-name>.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Enables Uniform bucket-level access access to a bucket.
	UniformBucketLevelAccess *bool `json:"uniformBucketLevelAccess,omitempty" tf:"uniform_bucket_level_access,omitempty"`

	// The bucket's Versioning configuration.  Structure is documented below.
	Versioning []VersioningObservation `json:"versioning,omitempty" tf:"versioning,omitempty"`

	// Configuration if the bucket acts as a website. Structure is documented below.
	Website []WebsiteObservation `json:"website,omitempty" tf:"website,omitempty"`
}

type BucketParameters struct {

	// The bucket's Autoclass configuration.  Structure is documented below.
	// +kubebuilder:validation:Optional
	Autoclass []AutoclassParameters `json:"autoclass,omitempty" tf:"autoclass,omitempty"`

	// The bucket's Cross-Origin Resource Sharing (CORS) configuration. Multiple blocks of this type are permitted. Structure is documented below.
	// +kubebuilder:validation:Optional
	Cors []CorsParameters `json:"cors,omitempty" tf:"cors,omitempty"`

	// The bucket's custom location configuration, which specifies the individual regions that comprise a dual-region bucket. If the bucket is designated a single or multi-region, the parameters are empty. Structure is documented below.
	// +kubebuilder:validation:Optional
	CustomPlacementConfig []CustomPlacementConfigParameters `json:"customPlacementConfig,omitempty" tf:"custom_placement_config,omitempty"`

	// Whether or not to automatically apply an eventBasedHold to new objects added to the bucket.
	// +kubebuilder:validation:Optional
	DefaultEventBasedHold *bool `json:"defaultEventBasedHold,omitempty" tf:"default_event_based_hold,omitempty"`

	// The bucket's encryption configuration. Structure is documented below.
	// +kubebuilder:validation:Optional
	Encryption []EncryptionParameters `json:"encryption,omitempty" tf:"encryption,omitempty"`

	// When deleting a bucket, this
	// boolean option will delete all contained objects.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// A map of key/value label pairs to assign to the bucket.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The bucket's Lifecycle Rules configuration. Multiple blocks of this type are permitted. Structure is documented below.
	// +kubebuilder:validation:Optional
	LifecycleRule []LifecycleRuleParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// The GCS location.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The bucket's Access & Storage Logs configuration. Structure is documented below.
	// +kubebuilder:validation:Optional
	Logging []LoggingParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Prevents public access to a bucket. Acceptable values are "inherited" or "enforced". If "inherited", the bucket uses public access prevention. only if the bucket is subject to the public access prevention organization policy constraint. Defaults to "inherited".
	// +kubebuilder:validation:Optional
	PublicAccessPrevention *string `json:"publicAccessPrevention,omitempty" tf:"public_access_prevention,omitempty"`

	// Enables Requester Pays on a storage bucket.
	// +kubebuilder:validation:Optional
	RequesterPays *bool `json:"requesterPays,omitempty" tf:"requester_pays,omitempty"`

	// Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.
	// +kubebuilder:validation:Optional
	RetentionPolicy []RetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// Enables Uniform bucket-level access access to a bucket.
	// +kubebuilder:validation:Optional
	UniformBucketLevelAccess *bool `json:"uniformBucketLevelAccess,omitempty" tf:"uniform_bucket_level_access,omitempty"`

	// The bucket's Versioning configuration.  Structure is documented below.
	// +kubebuilder:validation:Optional
	Versioning []VersioningParameters `json:"versioning,omitempty" tf:"versioning,omitempty"`

	// Configuration if the bucket acts as a website. Structure is documented below.
	// +kubebuilder:validation:Optional
	Website []WebsiteParameters `json:"website,omitempty" tf:"website,omitempty"`
}

type ConditionInitParameters struct {

	// Minimum age of an object in days to satisfy this condition.
	Age *float64 `json:"age,omitempty" tf:"age,omitempty"`

	// A date in the RFC 3339 format YYYY-MM-DD. This condition is satisfied when an object is created before midnight of the specified date in UTC.
	CreatedBefore *string `json:"createdBefore,omitempty" tf:"created_before,omitempty"`

	// A date in the RFC 3339 format YYYY-MM-DD. This condition is satisfied when the customTime metadata for the object is set to an earlier date than the date used in this lifecycle condition.
	CustomTimeBefore *string `json:"customTimeBefore,omitempty" tf:"custom_time_before,omitempty"`

	// Days since the date set in the customTime metadata for the object. This condition is satisfied when the current date and time is at least the specified number of days after the customTime.
	DaysSinceCustomTime *float64 `json:"daysSinceCustomTime,omitempty" tf:"days_since_custom_time,omitempty"`

	// Relevant only for versioned objects. Number of days elapsed since the noncurrent timestamp of an object.
	DaysSinceNoncurrentTime *float64 `json:"daysSinceNoncurrentTime,omitempty" tf:"days_since_noncurrent_time,omitempty"`

	// One or more matching name prefixes to satisfy this condition.
	MatchesPrefix []*string `json:"matchesPrefix,omitempty" tf:"matches_prefix,omitempty"`

	// Storage Class of objects to satisfy this condition. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE, DURABLE_REDUCED_AVAILABILITY.
	MatchesStorageClass []*string `json:"matchesStorageClass,omitempty" tf:"matches_storage_class,omitempty"`

	// One or more matching name suffixes to satisfy this condition.
	MatchesSuffix []*string `json:"matchesSuffix,omitempty" tf:"matches_suffix,omitempty"`

	// Relevant only for versioned objects. The date in RFC 3339 (e.g. 2017-06-13) when the object became nonconcurrent.
	NoncurrentTimeBefore *string `json:"noncurrentTimeBefore,omitempty" tf:"noncurrent_time_before,omitempty"`

	// Relevant only for versioned objects. The number of newer versions of an object to satisfy this condition.
	NumNewerVersions *float64 `json:"numNewerVersions,omitempty" tf:"num_newer_versions,omitempty"`

	// Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: "LIVE", "ARCHIVED", "ANY".
	WithState *string `json:"withState,omitempty" tf:"with_state,omitempty"`
}

type ConditionObservation struct {

	// Minimum age of an object in days to satisfy this condition.
	Age *float64 `json:"age,omitempty" tf:"age,omitempty"`

	// A date in the RFC 3339 format YYYY-MM-DD. This condition is satisfied when an object is created before midnight of the specified date in UTC.
	CreatedBefore *string `json:"createdBefore,omitempty" tf:"created_before,omitempty"`

	// A date in the RFC 3339 format YYYY-MM-DD. This condition is satisfied when the customTime metadata for the object is set to an earlier date than the date used in this lifecycle condition.
	CustomTimeBefore *string `json:"customTimeBefore,omitempty" tf:"custom_time_before,omitempty"`

	// Days since the date set in the customTime metadata for the object. This condition is satisfied when the current date and time is at least the specified number of days after the customTime.
	DaysSinceCustomTime *float64 `json:"daysSinceCustomTime,omitempty" tf:"days_since_custom_time,omitempty"`

	// Relevant only for versioned objects. Number of days elapsed since the noncurrent timestamp of an object.
	DaysSinceNoncurrentTime *float64 `json:"daysSinceNoncurrentTime,omitempty" tf:"days_since_noncurrent_time,omitempty"`

	// One or more matching name prefixes to satisfy this condition.
	MatchesPrefix []*string `json:"matchesPrefix,omitempty" tf:"matches_prefix,omitempty"`

	// Storage Class of objects to satisfy this condition. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE, DURABLE_REDUCED_AVAILABILITY.
	MatchesStorageClass []*string `json:"matchesStorageClass,omitempty" tf:"matches_storage_class,omitempty"`

	// One or more matching name suffixes to satisfy this condition.
	MatchesSuffix []*string `json:"matchesSuffix,omitempty" tf:"matches_suffix,omitempty"`

	// Relevant only for versioned objects. The date in RFC 3339 (e.g. 2017-06-13) when the object became nonconcurrent.
	NoncurrentTimeBefore *string `json:"noncurrentTimeBefore,omitempty" tf:"noncurrent_time_before,omitempty"`

	// Relevant only for versioned objects. The number of newer versions of an object to satisfy this condition.
	NumNewerVersions *float64 `json:"numNewerVersions,omitempty" tf:"num_newer_versions,omitempty"`

	// Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: "LIVE", "ARCHIVED", "ANY".
	WithState *string `json:"withState,omitempty" tf:"with_state,omitempty"`
}

type ConditionParameters struct {

	// Minimum age of an object in days to satisfy this condition.
	// +kubebuilder:validation:Optional
	Age *float64 `json:"age,omitempty" tf:"age,omitempty"`

	// A date in the RFC 3339 format YYYY-MM-DD. This condition is satisfied when an object is created before midnight of the specified date in UTC.
	// +kubebuilder:validation:Optional
	CreatedBefore *string `json:"createdBefore,omitempty" tf:"created_before,omitempty"`

	// A date in the RFC 3339 format YYYY-MM-DD. This condition is satisfied when the customTime metadata for the object is set to an earlier date than the date used in this lifecycle condition.
	// +kubebuilder:validation:Optional
	CustomTimeBefore *string `json:"customTimeBefore,omitempty" tf:"custom_time_before,omitempty"`

	// Days since the date set in the customTime metadata for the object. This condition is satisfied when the current date and time is at least the specified number of days after the customTime.
	// +kubebuilder:validation:Optional
	DaysSinceCustomTime *float64 `json:"daysSinceCustomTime,omitempty" tf:"days_since_custom_time,omitempty"`

	// Relevant only for versioned objects. Number of days elapsed since the noncurrent timestamp of an object.
	// +kubebuilder:validation:Optional
	DaysSinceNoncurrentTime *float64 `json:"daysSinceNoncurrentTime,omitempty" tf:"days_since_noncurrent_time,omitempty"`

	// One or more matching name prefixes to satisfy this condition.
	// +kubebuilder:validation:Optional
	MatchesPrefix []*string `json:"matchesPrefix,omitempty" tf:"matches_prefix,omitempty"`

	// Storage Class of objects to satisfy this condition. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE, DURABLE_REDUCED_AVAILABILITY.
	// +kubebuilder:validation:Optional
	MatchesStorageClass []*string `json:"matchesStorageClass,omitempty" tf:"matches_storage_class,omitempty"`

	// One or more matching name suffixes to satisfy this condition.
	// +kubebuilder:validation:Optional
	MatchesSuffix []*string `json:"matchesSuffix,omitempty" tf:"matches_suffix,omitempty"`

	// Relevant only for versioned objects. The date in RFC 3339 (e.g. 2017-06-13) when the object became nonconcurrent.
	// +kubebuilder:validation:Optional
	NoncurrentTimeBefore *string `json:"noncurrentTimeBefore,omitempty" tf:"noncurrent_time_before,omitempty"`

	// Relevant only for versioned objects. The number of newer versions of an object to satisfy this condition.
	// +kubebuilder:validation:Optional
	NumNewerVersions *float64 `json:"numNewerVersions,omitempty" tf:"num_newer_versions,omitempty"`

	// Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: "LIVE", "ARCHIVED", "ANY".
	// +kubebuilder:validation:Optional
	WithState *string `json:"withState,omitempty" tf:"with_state,omitempty"`
}

type CorsInitParameters struct {

	// The value, in seconds, to return in the Access-Control-Max-Age header used in preflight responses.
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`

	// The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: "*" is permitted in the list of methods, and means "any method".
	Method []*string `json:"method,omitempty" tf:"method,omitempty"`

	// The list of Origins eligible to receive CORS response headers. Note: "*" is permitted in the list of origins, and means "any Origin".
	Origin []*string `json:"origin,omitempty" tf:"origin,omitempty"`

	// The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains.
	ResponseHeader []*string `json:"responseHeader,omitempty" tf:"response_header,omitempty"`
}

type CorsObservation struct {

	// The value, in seconds, to return in the Access-Control-Max-Age header used in preflight responses.
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`

	// The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: "*" is permitted in the list of methods, and means "any method".
	Method []*string `json:"method,omitempty" tf:"method,omitempty"`

	// The list of Origins eligible to receive CORS response headers. Note: "*" is permitted in the list of origins, and means "any Origin".
	Origin []*string `json:"origin,omitempty" tf:"origin,omitempty"`

	// The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains.
	ResponseHeader []*string `json:"responseHeader,omitempty" tf:"response_header,omitempty"`
}

type CorsParameters struct {

	// The value, in seconds, to return in the Access-Control-Max-Age header used in preflight responses.
	// +kubebuilder:validation:Optional
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`

	// The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: "*" is permitted in the list of methods, and means "any method".
	// +kubebuilder:validation:Optional
	Method []*string `json:"method,omitempty" tf:"method,omitempty"`

	// The list of Origins eligible to receive CORS response headers. Note: "*" is permitted in the list of origins, and means "any Origin".
	// +kubebuilder:validation:Optional
	Origin []*string `json:"origin,omitempty" tf:"origin,omitempty"`

	// The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains.
	// +kubebuilder:validation:Optional
	ResponseHeader []*string `json:"responseHeader,omitempty" tf:"response_header,omitempty"`
}

type CustomPlacementConfigInitParameters struct {

	// The list of individual regions that comprise a dual-region bucket. See Cloud Storage bucket locations for a list of acceptable regions. Note: If any of the data_locations changes, it will recreate the bucket.
	DataLocations []*string `json:"dataLocations,omitempty" tf:"data_locations,omitempty"`
}

type CustomPlacementConfigObservation struct {

	// The list of individual regions that comprise a dual-region bucket. See Cloud Storage bucket locations for a list of acceptable regions. Note: If any of the data_locations changes, it will recreate the bucket.
	DataLocations []*string `json:"dataLocations,omitempty" tf:"data_locations,omitempty"`
}

type CustomPlacementConfigParameters struct {

	// The list of individual regions that comprise a dual-region bucket. See Cloud Storage bucket locations for a list of acceptable regions. Note: If any of the data_locations changes, it will recreate the bucket.
	// +kubebuilder:validation:Optional
	DataLocations []*string `json:"dataLocations,omitempty" tf:"data_locations,omitempty"`
}

type EncryptionInitParameters struct {

	// : The id of a Cloud KMS key that will be used to encrypt objects inserted into this bucket, if no encryption method is specified.
	// You must pay attention to whether the crypto key is available in the location that this bucket is created in.
	// See the docs for more details.
	DefaultKMSKeyName *string `json:"defaultKmsKeyName,omitempty" tf:"default_kms_key_name,omitempty"`
}

type EncryptionObservation struct {

	// : The id of a Cloud KMS key that will be used to encrypt objects inserted into this bucket, if no encryption method is specified.
	// You must pay attention to whether the crypto key is available in the location that this bucket is created in.
	// See the docs for more details.
	DefaultKMSKeyName *string `json:"defaultKmsKeyName,omitempty" tf:"default_kms_key_name,omitempty"`
}

type EncryptionParameters struct {

	// : The id of a Cloud KMS key that will be used to encrypt objects inserted into this bucket, if no encryption method is specified.
	// You must pay attention to whether the crypto key is available in the location that this bucket is created in.
	// See the docs for more details.
	// +kubebuilder:validation:Optional
	DefaultKMSKeyName *string `json:"defaultKmsKeyName,omitempty" tf:"default_kms_key_name,omitempty"`
}

type LifecycleRuleInitParameters struct {

	// The Lifecycle Rule's action configuration. A single block of this type is supported. Structure is documented below.
	Action []ActionInitParameters `json:"action,omitempty" tf:"action,omitempty"`

	// The Lifecycle Rule's condition configuration. A single block of this type is supported. Structure is documented below.
	Condition []ConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`
}

type LifecycleRuleObservation struct {

	// The Lifecycle Rule's action configuration. A single block of this type is supported. Structure is documented below.
	Action []ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// The Lifecycle Rule's condition configuration. A single block of this type is supported. Structure is documented below.
	Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`
}

type LifecycleRuleParameters struct {

	// The Lifecycle Rule's action configuration. A single block of this type is supported. Structure is documented below.
	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// The Lifecycle Rule's condition configuration. A single block of this type is supported. Structure is documented below.
	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`
}

type LoggingInitParameters struct {

	// The bucket that will receive log objects.
	LogBucket *string `json:"logBucket,omitempty" tf:"log_bucket,omitempty"`

	// The object prefix for log objects. If it's not provided,
	// by default GCS sets this to this bucket's name.
	LogObjectPrefix *string `json:"logObjectPrefix,omitempty" tf:"log_object_prefix,omitempty"`
}

type LoggingObservation struct {

	// The bucket that will receive log objects.
	LogBucket *string `json:"logBucket,omitempty" tf:"log_bucket,omitempty"`

	// The object prefix for log objects. If it's not provided,
	// by default GCS sets this to this bucket's name.
	LogObjectPrefix *string `json:"logObjectPrefix,omitempty" tf:"log_object_prefix,omitempty"`
}

type LoggingParameters struct {

	// The bucket that will receive log objects.
	// +kubebuilder:validation:Optional
	LogBucket *string `json:"logBucket,omitempty" tf:"log_bucket,omitempty"`

	// The object prefix for log objects. If it's not provided,
	// by default GCS sets this to this bucket's name.
	// +kubebuilder:validation:Optional
	LogObjectPrefix *string `json:"logObjectPrefix,omitempty" tf:"log_object_prefix,omitempty"`
}

type RetentionPolicyInitParameters struct {

	// If set to true, the bucket will be locked and permanently restrict edits to the bucket's retention policy.  Caution: Locking a bucket is an irreversible action.
	IsLocked *bool `json:"isLocked,omitempty" tf:"is_locked,omitempty"`

	// The period of time, in seconds, that objects in the bucket must be retained and cannot be deleted, overwritten, or archived. The value must be less than 2,147,483,647 seconds.
	RetentionPeriod *float64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`
}

type RetentionPolicyObservation struct {

	// If set to true, the bucket will be locked and permanently restrict edits to the bucket's retention policy.  Caution: Locking a bucket is an irreversible action.
	IsLocked *bool `json:"isLocked,omitempty" tf:"is_locked,omitempty"`

	// The period of time, in seconds, that objects in the bucket must be retained and cannot be deleted, overwritten, or archived. The value must be less than 2,147,483,647 seconds.
	RetentionPeriod *float64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`
}

type RetentionPolicyParameters struct {

	// If set to true, the bucket will be locked and permanently restrict edits to the bucket's retention policy.  Caution: Locking a bucket is an irreversible action.
	// +kubebuilder:validation:Optional
	IsLocked *bool `json:"isLocked,omitempty" tf:"is_locked,omitempty"`

	// The period of time, in seconds, that objects in the bucket must be retained and cannot be deleted, overwritten, or archived. The value must be less than 2,147,483,647 seconds.
	// +kubebuilder:validation:Optional
	RetentionPeriod *float64 `json:"retentionPeriod,omitempty" tf:"retention_period,omitempty"`
}

type VersioningInitParameters struct {

	// While set to true, versioning is fully enabled for this bucket.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type VersioningObservation struct {

	// While set to true, versioning is fully enabled for this bucket.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type VersioningParameters struct {

	// While set to true, versioning is fully enabled for this bucket.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type WebsiteInitParameters struct {

	// Behaves as the bucket's directory index where
	// missing objects are treated as potential directories.
	MainPageSuffix *string `json:"mainPageSuffix,omitempty" tf:"main_page_suffix,omitempty"`

	// The custom object to return when a requested
	// resource is not found.
	NotFoundPage *string `json:"notFoundPage,omitempty" tf:"not_found_page,omitempty"`
}

type WebsiteObservation struct {

	// Behaves as the bucket's directory index where
	// missing objects are treated as potential directories.
	MainPageSuffix *string `json:"mainPageSuffix,omitempty" tf:"main_page_suffix,omitempty"`

	// The custom object to return when a requested
	// resource is not found.
	NotFoundPage *string `json:"notFoundPage,omitempty" tf:"not_found_page,omitempty"`
}

type WebsiteParameters struct {

	// Behaves as the bucket's directory index where
	// missing objects are treated as potential directories.
	// +kubebuilder:validation:Optional
	MainPageSuffix *string `json:"mainPageSuffix,omitempty" tf:"main_page_suffix,omitempty"`

	// The custom object to return when a requested
	// resource is not found.
	// +kubebuilder:validation:Optional
	NotFoundPage *string `json:"notFoundPage,omitempty" tf:"not_found_page,omitempty"`
}

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketParameters `json:"forProvider"`
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
	InitProvider BucketInitParameters `json:"initProvider,omitempty"`
}

// BucketStatus defines the observed state of Bucket.
type BucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Bucket is the Schema for the Buckets API. Creates a new bucket in Google Cloud Storage.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	Spec   BucketSpec   `json:"spec"`
	Status BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// Repository type metadata.
var (
	Bucket_Kind             = "Bucket"
	Bucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Bucket_Kind}.String()
	Bucket_KindAPIVersion   = Bucket_Kind + "." + CRDGroupVersion.String()
	Bucket_GroupVersionKind = CRDGroupVersion.WithKind(Bucket_Kind)
)

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
