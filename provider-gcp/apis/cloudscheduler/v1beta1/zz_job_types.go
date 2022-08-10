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

type AppEngineHTTPTargetObservation struct {
}

type AppEngineHTTPTargetParameters struct {

	// App Engine Routing setting for the job.
	// +kubebuilder:validation:Optional
	AppEngineRouting []AppEngineRoutingParameters `json:"appEngineRouting,omitempty" tf:"app_engine_routing,omitempty"`

	// HTTP request body.
	// A request body is allowed only if the HTTP method is POST or PUT.
	// It will result in invalid argument error to set a body on a job with an incompatible HttpMethod.
	//
	// A base64-encoded string.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Which HTTP method to use for the request.
	// +kubebuilder:validation:Optional
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// HTTP request headers.
	// This map contains the header field names and values.
	// Headers can be set when the job is created.
	// +kubebuilder:validation:Optional
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// The relative URI.
	// The relative URL must begin with "/" and must be a valid HTTP relative URL.
	// It can contain a path, query string arguments, and \# fragments.
	// If the relative URL is empty, then the root path "/" will be used.
	// No spaces are allowed, and the maximum length allowed is 2083 characters
	// +kubebuilder:validation:Required
	RelativeURI *string `json:"relativeUri" tf:"relative_uri,omitempty"`
}

type AppEngineRoutingObservation struct {
}

type AppEngineRoutingParameters struct {

	// App instance.
	// By default, the job is sent to an instance which is available when the job is attempted.
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// App service.
	// By default, the job is sent to the service which is the default service when the job is attempted.
	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// App version.
	// By default, the job is sent to the version which is the default version when the job is attempted.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type HTTPTargetObservation struct {
}

type HTTPTargetParameters struct {

	// HTTP request body.
	// A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
	// It is an error to set body on a job with an incompatible HttpMethod.
	//
	// A base64-encoded string.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Which HTTP method to use for the request.
	// +kubebuilder:validation:Optional
	HTTPMethod *string `json:"httpMethod,omitempty" tf:"http_method,omitempty"`

	// This map contains the header field names and values.
	// Repeated headers are not supported, but a header value can contain commas.
	// +kubebuilder:validation:Optional
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// Contains information needed for generating an OAuth token.
	// This type of authorization should be used when sending requests to a GCP endpoint.
	// +kubebuilder:validation:Optional
	OAuthToken []OAuthTokenParameters `json:"oauthToken,omitempty" tf:"oauth_token,omitempty"`

	// Contains information needed for generating an OpenID Connect token.
	// This type of authorization should be used when sending requests to third party endpoints or Cloud Run.
	// +kubebuilder:validation:Optional
	OidcToken []OidcTokenParameters `json:"oidcToken,omitempty" tf:"oidc_token,omitempty"`

	// The full URI path that the request will be sent to.
	// +kubebuilder:validation:Required
	URI *string `json:"uri" tf:"uri,omitempty"`
}

type JobObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type JobParameters struct {

	// App Engine HTTP target.
	// If the job providers a App Engine HTTP target the cron will
	// send a request to the service instance
	// +kubebuilder:validation:Optional
	AppEngineHTTPTarget []AppEngineHTTPTargetParameters `json:"appEngineHttpTarget,omitempty" tf:"app_engine_http_target,omitempty"`

	// The deadline for job attempts. If the request handler does not respond by this deadline then the request is
	// cancelled and the attempt is marked as a DEADLINE_EXCEEDED failure. The failed attempt can be viewed in
	// execution logs. Cloud Scheduler will retry the job according to the RetryConfig.
	// The allowed duration for this deadline is:
	// * For HTTP targets, between 15 seconds and 30 minutes.
	// * For App Engine HTTP targets, between 15 seconds and 24 hours.
	// * **Note**: For PubSub targets, this field is ignored - setting it will introduce an unresolvable diff.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s"
	// +kubebuilder:validation:Optional
	AttemptDeadline *string `json:"attemptDeadline,omitempty" tf:"attempt_deadline,omitempty"`

	// A human-readable description for the job.
	// This string must not contain more than 500 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// HTTP target.
	// If the job providers a http_target the cron will
	// send a request to the targeted url
	// +kubebuilder:validation:Optional
	HTTPTarget []HTTPTargetParameters `json:"httpTarget,omitempty" tf:"http_target,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/cloudplatform/v1beta1.Project
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// Pub/Sub target
	// If the job providers a Pub/Sub target the cron will publish
	// a message to the provided topic
	// +kubebuilder:validation:Optional
	PubsubTarget []PubsubTargetParameters `json:"pubsubTarget,omitempty" tf:"pubsub_target,omitempty"`

	// Region where the scheduler job resides. If it is not provided, Terraform will use the provider default.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// By default, if a job does not complete successfully,
	// meaning that an acknowledgement is not received from the handler,
	// then it will be retried with exponential backoff according to the settings
	// +kubebuilder:validation:Optional
	RetryConfig []RetryConfigParameters `json:"retryConfig,omitempty" tf:"retry_config,omitempty"`

	// Describes the schedule on which the job will be executed.
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// Specifies the time zone to be used in interpreting schedule.
	// The value of this field must be a time zone name from the tz database.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type OAuthTokenObservation struct {
}

type OAuthTokenParameters struct {

	// OAuth scope to be used for generating OAuth access token. If not specified,
	// "https://www.googleapis.com/auth/cloud-platform" will be used.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// Service account email to be used for generating OAuth token.
	// The service account must be within the same project as the job.
	// +kubebuilder:validation:Required
	ServiceAccountEmail *string `json:"serviceAccountEmail" tf:"service_account_email,omitempty"`
}

type OidcTokenObservation struct {
}

type OidcTokenParameters struct {

	// Audience to be used when generating OIDC token. If not specified,
	// the URI specified in target will be used.
	// +kubebuilder:validation:Optional
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// Service account email to be used for generating OAuth token.
	// The service account must be within the same project as the job.
	// +kubebuilder:validation:Required
	ServiceAccountEmail *string `json:"serviceAccountEmail" tf:"service_account_email,omitempty"`
}

type PubsubTargetObservation struct {
}

type PubsubTargetParameters struct {

	// Attributes for PubsubMessage.
	// Pubsub message must contain either non-empty data, or at least one attribute.
	// +kubebuilder:validation:Optional
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// The message payload for PubsubMessage.
	// Pubsub message must contain either non-empty data, or at least one attribute.
	//
	// A base64-encoded string.
	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data,omitempty"`

	// The full resource name for the Cloud Pub/Sub topic to which
	// messages will be published when a job is delivered. ~>**NOTE:**
	// The topic name must be in the same format as required by PubSub's
	// PublishRequest.name, e.g. 'projects/my-project/topics/my-topic'.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/pubsub/v1beta1.Topic
	// +kubebuilder:validation:Optional
	TopicName *string `json:"topicName,omitempty" tf:"topic_name,omitempty"`

	// +kubebuilder:validation:Optional
	TopicNameRef *v1.Reference `json:"topicNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TopicNameSelector *v1.Selector `json:"topicNameSelector,omitempty" tf:"-"`
}

type RetryConfigObservation struct {
}

type RetryConfigParameters struct {

	// The maximum amount of time to wait before retrying a job after it fails.
	// A duration in seconds with up to nine fractional digits, terminated by 's'.
	// +kubebuilder:validation:Optional
	MaxBackoffDuration *string `json:"maxBackoffDuration,omitempty" tf:"max_backoff_duration,omitempty"`

	// The time between retries will double maxDoublings times.
	// A job's retry interval starts at minBackoffDuration,
	// then doubles maxDoublings times, then increases linearly,
	// and finally retries retries at intervals of maxBackoffDuration up to retryCount times.
	// +kubebuilder:validation:Optional
	MaxDoublings *float64 `json:"maxDoublings,omitempty" tf:"max_doublings,omitempty"`

	// The time limit for retrying a failed job, measured from time when an execution was first attempted.
	// If specified with retryCount, the job will be retried until both limits are reached.
	// A duration in seconds with up to nine fractional digits, terminated by 's'.
	// +kubebuilder:validation:Optional
	MaxRetryDuration *string `json:"maxRetryDuration,omitempty" tf:"max_retry_duration,omitempty"`

	// The minimum amount of time to wait before retrying a job after it fails.
	// A duration in seconds with up to nine fractional digits, terminated by 's'.
	// +kubebuilder:validation:Optional
	MinBackoffDuration *string `json:"minBackoffDuration,omitempty" tf:"min_backoff_duration,omitempty"`

	// The number of attempts that the system will make to run a
	// job using the exponential backoff procedure described by maxDoublings.
	// Values greater than 5 and negative values are not allowed.
	// +kubebuilder:validation:Optional
	RetryCount *float64 `json:"retryCount,omitempty" tf:"retry_count,omitempty"`
}

// JobSpec defines the desired state of Job
type JobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JobParameters `json:"forProvider"`
}

// JobStatus defines the observed state of Job.
type JobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Job is the Schema for the Jobs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Job struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JobSpec   `json:"spec"`
	Status            JobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobList contains a list of Jobs
type JobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Job `json:"items"`
}

// Repository type metadata.
var (
	Job_Kind             = "Job"
	Job_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Job_Kind}.String()
	Job_KindAPIVersion   = Job_Kind + "." + CRDGroupVersion.String()
	Job_GroupVersionKind = CRDGroupVersion.WithKind(Job_Kind)
)

func init() {
	SchemeBuilder.Register(&Job{}, &JobList{})
}
