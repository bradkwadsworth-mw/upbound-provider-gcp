//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigInitParameters) DeepCopyInto(out *ConfigInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigInitParameters.
func (in *ConfigInitParameters) DeepCopy() *ConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigObservation) DeepCopyInto(out *ConfigObservation) {
	*out = *in
	if in.BinaryData != nil {
		in, out := &in.BinaryData, &out.BinaryData
		*out = new(string)
		**out = **in
	}
	if in.CloudUpdateTime != nil {
		in, out := &in.CloudUpdateTime, &out.CloudUpdateTime
		*out = new(string)
		**out = **in
	}
	if in.DeviceAckTime != nil {
		in, out := &in.DeviceAckTime, &out.DeviceAckTime
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigObservation.
func (in *ConfigObservation) DeepCopy() *ConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigParameters) DeepCopyInto(out *ConfigParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigParameters.
func (in *ConfigParameters) DeepCopy() *ConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsInitParameters) DeepCopyInto(out *CredentialsInitParameters) {
	*out = *in
	if in.ExpirationTime != nil {
		in, out := &in.ExpirationTime, &out.ExpirationTime
		*out = new(string)
		**out = **in
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = make([]PublicKeyInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsInitParameters.
func (in *CredentialsInitParameters) DeepCopy() *CredentialsInitParameters {
	if in == nil {
		return nil
	}
	out := new(CredentialsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsObservation) DeepCopyInto(out *CredentialsObservation) {
	*out = *in
	if in.ExpirationTime != nil {
		in, out := &in.ExpirationTime, &out.ExpirationTime
		*out = new(string)
		**out = **in
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = make([]PublicKeyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsObservation.
func (in *CredentialsObservation) DeepCopy() *CredentialsObservation {
	if in == nil {
		return nil
	}
	out := new(CredentialsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsParameters) DeepCopyInto(out *CredentialsParameters) {
	*out = *in
	if in.ExpirationTime != nil {
		in, out := &in.ExpirationTime, &out.ExpirationTime
		*out = new(string)
		**out = **in
	}
	if in.PublicKey != nil {
		in, out := &in.PublicKey, &out.PublicKey
		*out = make([]PublicKeyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsParameters.
func (in *CredentialsParameters) DeepCopy() *CredentialsParameters {
	if in == nil {
		return nil
	}
	out := new(CredentialsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Device) DeepCopyInto(out *Device) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Device.
func (in *Device) DeepCopy() *Device {
	if in == nil {
		return nil
	}
	out := new(Device)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Device) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceInitParameters) DeepCopyInto(out *DeviceInitParameters) {
	*out = *in
	if in.Blocked != nil {
		in, out := &in.Blocked, &out.Blocked
		*out = new(bool)
		**out = **in
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]CredentialsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GatewayConfig != nil {
		in, out := &in.GatewayConfig, &out.GatewayConfig
		*out = make([]GatewayConfigInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceInitParameters.
func (in *DeviceInitParameters) DeepCopy() *DeviceInitParameters {
	if in == nil {
		return nil
	}
	out := new(DeviceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceList) DeepCopyInto(out *DeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Device, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceList.
func (in *DeviceList) DeepCopy() *DeviceList {
	if in == nil {
		return nil
	}
	out := new(DeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceObservation) DeepCopyInto(out *DeviceObservation) {
	*out = *in
	if in.Blocked != nil {
		in, out := &in.Blocked, &out.Blocked
		*out = new(bool)
		**out = **in
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make([]ConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]CredentialsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GatewayConfig != nil {
		in, out := &in.GatewayConfig, &out.GatewayConfig
		*out = make([]GatewayConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LastConfigAckTime != nil {
		in, out := &in.LastConfigAckTime, &out.LastConfigAckTime
		*out = new(string)
		**out = **in
	}
	if in.LastConfigSendTime != nil {
		in, out := &in.LastConfigSendTime, &out.LastConfigSendTime
		*out = new(string)
		**out = **in
	}
	if in.LastErrorStatus != nil {
		in, out := &in.LastErrorStatus, &out.LastErrorStatus
		*out = make([]LastErrorStatusObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastErrorTime != nil {
		in, out := &in.LastErrorTime, &out.LastErrorTime
		*out = new(string)
		**out = **in
	}
	if in.LastEventTime != nil {
		in, out := &in.LastEventTime, &out.LastEventTime
		*out = new(string)
		**out = **in
	}
	if in.LastHeartbeatTime != nil {
		in, out := &in.LastHeartbeatTime, &out.LastHeartbeatTime
		*out = new(string)
		**out = **in
	}
	if in.LastStateTime != nil {
		in, out := &in.LastStateTime, &out.LastStateTime
		*out = new(string)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.NumID != nil {
		in, out := &in.NumID, &out.NumID
		*out = new(string)
		**out = **in
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = make([]StateObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceObservation.
func (in *DeviceObservation) DeepCopy() *DeviceObservation {
	if in == nil {
		return nil
	}
	out := new(DeviceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceParameters) DeepCopyInto(out *DeviceParameters) {
	*out = *in
	if in.Blocked != nil {
		in, out := &in.Blocked, &out.Blocked
		*out = new(bool)
		**out = **in
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]CredentialsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GatewayConfig != nil {
		in, out := &in.GatewayConfig, &out.GatewayConfig
		*out = make([]GatewayConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(string)
		**out = **in
	}
	if in.RegistryRef != nil {
		in, out := &in.RegistryRef, &out.RegistryRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RegistrySelector != nil {
		in, out := &in.RegistrySelector, &out.RegistrySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceParameters.
func (in *DeviceParameters) DeepCopy() *DeviceParameters {
	if in == nil {
		return nil
	}
	out := new(DeviceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSpec) DeepCopyInto(out *DeviceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSpec.
func (in *DeviceSpec) DeepCopy() *DeviceSpec {
	if in == nil {
		return nil
	}
	out := new(DeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceStatus) DeepCopyInto(out *DeviceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceStatus.
func (in *DeviceStatus) DeepCopy() *DeviceStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventNotificationConfigsInitParameters) DeepCopyInto(out *EventNotificationConfigsInitParameters) {
	*out = *in
	if in.SubfolderMatches != nil {
		in, out := &in.SubfolderMatches, &out.SubfolderMatches
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventNotificationConfigsInitParameters.
func (in *EventNotificationConfigsInitParameters) DeepCopy() *EventNotificationConfigsInitParameters {
	if in == nil {
		return nil
	}
	out := new(EventNotificationConfigsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventNotificationConfigsObservation) DeepCopyInto(out *EventNotificationConfigsObservation) {
	*out = *in
	if in.PubsubTopicName != nil {
		in, out := &in.PubsubTopicName, &out.PubsubTopicName
		*out = new(string)
		**out = **in
	}
	if in.SubfolderMatches != nil {
		in, out := &in.SubfolderMatches, &out.SubfolderMatches
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventNotificationConfigsObservation.
func (in *EventNotificationConfigsObservation) DeepCopy() *EventNotificationConfigsObservation {
	if in == nil {
		return nil
	}
	out := new(EventNotificationConfigsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventNotificationConfigsParameters) DeepCopyInto(out *EventNotificationConfigsParameters) {
	*out = *in
	if in.PubsubTopicName != nil {
		in, out := &in.PubsubTopicName, &out.PubsubTopicName
		*out = new(string)
		**out = **in
	}
	if in.PubsubTopicNameRef != nil {
		in, out := &in.PubsubTopicNameRef, &out.PubsubTopicNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.PubsubTopicNameSelector != nil {
		in, out := &in.PubsubTopicNameSelector, &out.PubsubTopicNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SubfolderMatches != nil {
		in, out := &in.SubfolderMatches, &out.SubfolderMatches
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventNotificationConfigsParameters.
func (in *EventNotificationConfigsParameters) DeepCopy() *EventNotificationConfigsParameters {
	if in == nil {
		return nil
	}
	out := new(EventNotificationConfigsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayConfigInitParameters) DeepCopyInto(out *GatewayConfigInitParameters) {
	*out = *in
	if in.GatewayAuthMethod != nil {
		in, out := &in.GatewayAuthMethod, &out.GatewayAuthMethod
		*out = new(string)
		**out = **in
	}
	if in.GatewayType != nil {
		in, out := &in.GatewayType, &out.GatewayType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayConfigInitParameters.
func (in *GatewayConfigInitParameters) DeepCopy() *GatewayConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(GatewayConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayConfigObservation) DeepCopyInto(out *GatewayConfigObservation) {
	*out = *in
	if in.GatewayAuthMethod != nil {
		in, out := &in.GatewayAuthMethod, &out.GatewayAuthMethod
		*out = new(string)
		**out = **in
	}
	if in.GatewayType != nil {
		in, out := &in.GatewayType, &out.GatewayType
		*out = new(string)
		**out = **in
	}
	if in.LastAccessedGatewayID != nil {
		in, out := &in.LastAccessedGatewayID, &out.LastAccessedGatewayID
		*out = new(string)
		**out = **in
	}
	if in.LastAccessedGatewayTime != nil {
		in, out := &in.LastAccessedGatewayTime, &out.LastAccessedGatewayTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayConfigObservation.
func (in *GatewayConfigObservation) DeepCopy() *GatewayConfigObservation {
	if in == nil {
		return nil
	}
	out := new(GatewayConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayConfigParameters) DeepCopyInto(out *GatewayConfigParameters) {
	*out = *in
	if in.GatewayAuthMethod != nil {
		in, out := &in.GatewayAuthMethod, &out.GatewayAuthMethod
		*out = new(string)
		**out = **in
	}
	if in.GatewayType != nil {
		in, out := &in.GatewayType, &out.GatewayType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayConfigParameters.
func (in *GatewayConfigParameters) DeepCopy() *GatewayConfigParameters {
	if in == nil {
		return nil
	}
	out := new(GatewayConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastErrorStatusInitParameters) DeepCopyInto(out *LastErrorStatusInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastErrorStatusInitParameters.
func (in *LastErrorStatusInitParameters) DeepCopy() *LastErrorStatusInitParameters {
	if in == nil {
		return nil
	}
	out := new(LastErrorStatusInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastErrorStatusObservation) DeepCopyInto(out *LastErrorStatusObservation) {
	*out = *in
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Number != nil {
		in, out := &in.Number, &out.Number
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastErrorStatusObservation.
func (in *LastErrorStatusObservation) DeepCopy() *LastErrorStatusObservation {
	if in == nil {
		return nil
	}
	out := new(LastErrorStatusObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastErrorStatusParameters) DeepCopyInto(out *LastErrorStatusParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastErrorStatusParameters.
func (in *LastErrorStatusParameters) DeepCopy() *LastErrorStatusParameters {
	if in == nil {
		return nil
	}
	out := new(LastErrorStatusParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicKeyInitParameters) DeepCopyInto(out *PublicKeyInitParameters) {
	*out = *in
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicKeyInitParameters.
func (in *PublicKeyInitParameters) DeepCopy() *PublicKeyInitParameters {
	if in == nil {
		return nil
	}
	out := new(PublicKeyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicKeyObservation) DeepCopyInto(out *PublicKeyObservation) {
	*out = *in
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicKeyObservation.
func (in *PublicKeyObservation) DeepCopy() *PublicKeyObservation {
	if in == nil {
		return nil
	}
	out := new(PublicKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicKeyParameters) DeepCopyInto(out *PublicKeyParameters) {
	*out = *in
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicKeyParameters.
func (in *PublicKeyParameters) DeepCopy() *PublicKeyParameters {
	if in == nil {
		return nil
	}
	out := new(PublicKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Registry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryCredentialsInitParameters) DeepCopyInto(out *RegistryCredentialsInitParameters) {
	*out = *in
	if in.PublicKeyCertificate != nil {
		in, out := &in.PublicKeyCertificate, &out.PublicKeyCertificate
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryCredentialsInitParameters.
func (in *RegistryCredentialsInitParameters) DeepCopy() *RegistryCredentialsInitParameters {
	if in == nil {
		return nil
	}
	out := new(RegistryCredentialsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryCredentialsObservation) DeepCopyInto(out *RegistryCredentialsObservation) {
	*out = *in
	if in.PublicKeyCertificate != nil {
		in, out := &in.PublicKeyCertificate, &out.PublicKeyCertificate
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryCredentialsObservation.
func (in *RegistryCredentialsObservation) DeepCopy() *RegistryCredentialsObservation {
	if in == nil {
		return nil
	}
	out := new(RegistryCredentialsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryCredentialsParameters) DeepCopyInto(out *RegistryCredentialsParameters) {
	*out = *in
	if in.PublicKeyCertificate != nil {
		in, out := &in.PublicKeyCertificate, &out.PublicKeyCertificate
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryCredentialsParameters.
func (in *RegistryCredentialsParameters) DeepCopy() *RegistryCredentialsParameters {
	if in == nil {
		return nil
	}
	out := new(RegistryCredentialsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryInitParameters) DeepCopyInto(out *RegistryInitParameters) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]RegistryCredentialsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EventNotificationConfigs != nil {
		in, out := &in.EventNotificationConfigs, &out.EventNotificationConfigs
		*out = make([]EventNotificationConfigsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.MqttConfig != nil {
		in, out := &in.MqttConfig, &out.MqttConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.StateNotificationConfig != nil {
		in, out := &in.StateNotificationConfig, &out.StateNotificationConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryInitParameters.
func (in *RegistryInitParameters) DeepCopy() *RegistryInitParameters {
	if in == nil {
		return nil
	}
	out := new(RegistryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryList) DeepCopyInto(out *RegistryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Registry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryList.
func (in *RegistryList) DeepCopy() *RegistryList {
	if in == nil {
		return nil
	}
	out := new(RegistryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryObservation) DeepCopyInto(out *RegistryObservation) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]RegistryCredentialsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EventNotificationConfigs != nil {
		in, out := &in.EventNotificationConfigs, &out.EventNotificationConfigs
		*out = make([]EventNotificationConfigsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.MqttConfig != nil {
		in, out := &in.MqttConfig, &out.MqttConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.StateNotificationConfig != nil {
		in, out := &in.StateNotificationConfig, &out.StateNotificationConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryObservation.
func (in *RegistryObservation) DeepCopy() *RegistryObservation {
	if in == nil {
		return nil
	}
	out := new(RegistryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryParameters) DeepCopyInto(out *RegistryParameters) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]RegistryCredentialsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EventNotificationConfigs != nil {
		in, out := &in.EventNotificationConfigs, &out.EventNotificationConfigs
		*out = make([]EventNotificationConfigsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HTTPConfig != nil {
		in, out := &in.HTTPConfig, &out.HTTPConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.MqttConfig != nil {
		in, out := &in.MqttConfig, &out.MqttConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.StateNotificationConfig != nil {
		in, out := &in.StateNotificationConfig, &out.StateNotificationConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryParameters.
func (in *RegistryParameters) DeepCopy() *RegistryParameters {
	if in == nil {
		return nil
	}
	out := new(RegistryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistrySpec) DeepCopyInto(out *RegistrySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistrySpec.
func (in *RegistrySpec) DeepCopy() *RegistrySpec {
	if in == nil {
		return nil
	}
	out := new(RegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryStatus) DeepCopyInto(out *RegistryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryStatus.
func (in *RegistryStatus) DeepCopy() *RegistryStatus {
	if in == nil {
		return nil
	}
	out := new(RegistryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateInitParameters) DeepCopyInto(out *StateInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateInitParameters.
func (in *StateInitParameters) DeepCopy() *StateInitParameters {
	if in == nil {
		return nil
	}
	out := new(StateInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateObservation) DeepCopyInto(out *StateObservation) {
	*out = *in
	if in.BinaryData != nil {
		in, out := &in.BinaryData, &out.BinaryData
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateObservation.
func (in *StateObservation) DeepCopy() *StateObservation {
	if in == nil {
		return nil
	}
	out := new(StateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateParameters) DeepCopyInto(out *StateParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateParameters.
func (in *StateParameters) DeepCopy() *StateParameters {
	if in == nil {
		return nil
	}
	out := new(StateParameters)
	in.DeepCopyInto(out)
	return out
}
