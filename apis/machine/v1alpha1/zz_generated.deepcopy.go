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

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Amazonec2ConfigObservation) DeepCopyInto(out *Amazonec2ConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Amazonec2ConfigObservation.
func (in *Amazonec2ConfigObservation) DeepCopy() *Amazonec2ConfigObservation {
	if in == nil {
		return nil
	}
	out := new(Amazonec2ConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Amazonec2ConfigParameters) DeepCopyInto(out *Amazonec2ConfigParameters) {
	*out = *in
	if in.AMI != nil {
		in, out := &in.AMI, &out.AMI
		*out = new(string)
		**out = **in
	}
	if in.AccessKeySecretRef != nil {
		in, out := &in.AccessKeySecretRef, &out.AccessKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.BlockDurationMinutes != nil {
		in, out := &in.BlockDurationMinutes, &out.BlockDurationMinutes
		*out = new(string)
		**out = **in
	}
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.EncryptEBSVolume != nil {
		in, out := &in.EncryptEBSVolume, &out.EncryptEBSVolume
		*out = new(bool)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.HTTPEndpoint != nil {
		in, out := &in.HTTPEndpoint, &out.HTTPEndpoint
		*out = new(string)
		**out = **in
	}
	if in.HTTPTokens != nil {
		in, out := &in.HTTPTokens, &out.HTTPTokens
		*out = new(string)
		**out = **in
	}
	if in.IAMInstanceProfile != nil {
		in, out := &in.IAMInstanceProfile, &out.IAMInstanceProfile
		*out = new(string)
		**out = **in
	}
	if in.InsecureTransport != nil {
		in, out := &in.InsecureTransport, &out.InsecureTransport
		*out = new(bool)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.KMSKey != nil {
		in, out := &in.KMSKey, &out.KMSKey
		*out = new(string)
		**out = **in
	}
	if in.KeypairName != nil {
		in, out := &in.KeypairName, &out.KeypairName
		*out = new(string)
		**out = **in
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(bool)
		**out = **in
	}
	if in.OpenPort != nil {
		in, out := &in.OpenPort, &out.OpenPort
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PrivateAddressOnly != nil {
		in, out := &in.PrivateAddressOnly, &out.PrivateAddressOnly
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RequestSpotInstance != nil {
		in, out := &in.RequestSpotInstance, &out.RequestSpotInstance
		*out = new(bool)
		**out = **in
	}
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		*out = new(string)
		**out = **in
	}
	if in.RootSize != nil {
		in, out := &in.RootSize, &out.RootSize
		*out = new(string)
		**out = **in
	}
	if in.SSHKeyContentsSecretRef != nil {
		in, out := &in.SSHKeyContentsSecretRef, &out.SSHKeyContentsSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SSHUser != nil {
		in, out := &in.SSHUser, &out.SSHUser
		*out = new(string)
		**out = **in
	}
	if in.SecretKeySecretRef != nil {
		in, out := &in.SecretKeySecretRef, &out.SecretKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SecurityGroup != nil {
		in, out := &in.SecurityGroup, &out.SecurityGroup
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityGroupReadonly != nil {
		in, out := &in.SecurityGroupReadonly, &out.SecurityGroupReadonly
		*out = new(bool)
		**out = **in
	}
	if in.SessionTokenSecretRef != nil {
		in, out := &in.SessionTokenSecretRef, &out.SessionTokenSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SpotPrice != nil {
		in, out := &in.SpotPrice, &out.SpotPrice
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(string)
		**out = **in
	}
	if in.UseEBSOptimizedInstance != nil {
		in, out := &in.UseEBSOptimizedInstance, &out.UseEBSOptimizedInstance
		*out = new(bool)
		**out = **in
	}
	if in.UsePrivateAddress != nil {
		in, out := &in.UsePrivateAddress, &out.UsePrivateAddress
		*out = new(bool)
		**out = **in
	}
	if in.Userdata != nil {
		in, out := &in.Userdata, &out.Userdata
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Amazonec2ConfigParameters.
func (in *Amazonec2ConfigParameters) DeepCopy() *Amazonec2ConfigParameters {
	if in == nil {
		return nil
	}
	out := new(Amazonec2ConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureConfigObservation) DeepCopyInto(out *AzureConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureConfigObservation.
func (in *AzureConfigObservation) DeepCopy() *AzureConfigObservation {
	if in == nil {
		return nil
	}
	out := new(AzureConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureConfigParameters) DeepCopyInto(out *AzureConfigParameters) {
	*out = *in
	if in.AvailabilitySet != nil {
		in, out := &in.AvailabilitySet, &out.AvailabilitySet
		*out = new(string)
		**out = **in
	}
	if in.ClientIDSecretRef != nil {
		in, out := &in.ClientIDSecretRef, &out.ClientIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ClientSecretSecretRef != nil {
		in, out := &in.ClientSecretSecretRef, &out.ClientSecretSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.CustomData != nil {
		in, out := &in.CustomData, &out.CustomData
		*out = new(string)
		**out = **in
	}
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = new(string)
		**out = **in
	}
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(string)
		**out = **in
	}
	if in.DockerPort != nil {
		in, out := &in.DockerPort, &out.DockerPort
		*out = new(string)
		**out = **in
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = new(string)
		**out = **in
	}
	if in.FaultDomainCount != nil {
		in, out := &in.FaultDomainCount, &out.FaultDomainCount
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ManagedDisks != nil {
		in, out := &in.ManagedDisks, &out.ManagedDisks
		*out = new(bool)
		**out = **in
	}
	if in.NoPublicIP != nil {
		in, out := &in.NoPublicIP, &out.NoPublicIP
		*out = new(bool)
		**out = **in
	}
	if in.Nsg != nil {
		in, out := &in.Nsg, &out.Nsg
		*out = new(string)
		**out = **in
	}
	if in.OpenPort != nil {
		in, out := &in.OpenPort, &out.OpenPort
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PrivateIPAddress != nil {
		in, out := &in.PrivateIPAddress, &out.PrivateIPAddress
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroup != nil {
		in, out := &in.ResourceGroup, &out.ResourceGroup
		*out = new(string)
		**out = **in
	}
	if in.SSHUser != nil {
		in, out := &in.SSHUser, &out.SSHUser
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.StaticPublicIP != nil {
		in, out := &in.StaticPublicIP, &out.StaticPublicIP
		*out = new(bool)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
	if in.Subnet != nil {
		in, out := &in.Subnet, &out.Subnet
		*out = new(string)
		**out = **in
	}
	if in.SubnetPrefix != nil {
		in, out := &in.SubnetPrefix, &out.SubnetPrefix
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionIDSecretRef != nil {
		in, out := &in.SubscriptionIDSecretRef, &out.SubscriptionIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.TenantIDSecretRef != nil {
		in, out := &in.TenantIDSecretRef, &out.TenantIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.UpdateDomainCount != nil {
		in, out := &in.UpdateDomainCount, &out.UpdateDomainCount
		*out = new(string)
		**out = **in
	}
	if in.UsePrivateIP != nil {
		in, out := &in.UsePrivateIP, &out.UsePrivateIP
		*out = new(bool)
		**out = **in
	}
	if in.Vnet != nil {
		in, out := &in.Vnet, &out.Vnet
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureConfigParameters.
func (in *AzureConfigParameters) DeepCopy() *AzureConfigParameters {
	if in == nil {
		return nil
	}
	out := new(AzureConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigV2) DeepCopyInto(out *ConfigV2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigV2.
func (in *ConfigV2) DeepCopy() *ConfigV2 {
	if in == nil {
		return nil
	}
	out := new(ConfigV2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigV2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigV2List) DeepCopyInto(out *ConfigV2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConfigV2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigV2List.
func (in *ConfigV2List) DeepCopy() *ConfigV2List {
	if in == nil {
		return nil
	}
	out := new(ConfigV2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigV2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigV2Observation) DeepCopyInto(out *ConfigV2Observation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.ResourceVersion != nil {
		in, out := &in.ResourceVersion, &out.ResourceVersion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigV2Observation.
func (in *ConfigV2Observation) DeepCopy() *ConfigV2Observation {
	if in == nil {
		return nil
	}
	out := new(ConfigV2Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigV2Parameters) DeepCopyInto(out *ConfigV2Parameters) {
	*out = *in
	if in.Amazonec2Config != nil {
		in, out := &in.Amazonec2Config, &out.Amazonec2Config
		*out = make([]Amazonec2ConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AzureConfig != nil {
		in, out := &in.AzureConfig, &out.AzureConfig
		*out = make([]AzureConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DigitaloceanConfig != nil {
		in, out := &in.DigitaloceanConfig, &out.DigitaloceanConfig
		*out = make([]DigitaloceanConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FleetNamespace != nil {
		in, out := &in.FleetNamespace, &out.FleetNamespace
		*out = new(string)
		**out = **in
	}
	if in.GenerateName != nil {
		in, out := &in.GenerateName, &out.GenerateName
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LinodeConfig != nil {
		in, out := &in.LinodeConfig, &out.LinodeConfig
		*out = make([]LinodeConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OpenstackConfig != nil {
		in, out := &in.OpenstackConfig, &out.OpenstackConfig
		*out = make([]OpenstackConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VsphereConfig != nil {
		in, out := &in.VsphereConfig, &out.VsphereConfig
		*out = make([]VsphereConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigV2Parameters.
func (in *ConfigV2Parameters) DeepCopy() *ConfigV2Parameters {
	if in == nil {
		return nil
	}
	out := new(ConfigV2Parameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigV2Spec) DeepCopyInto(out *ConfigV2Spec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigV2Spec.
func (in *ConfigV2Spec) DeepCopy() *ConfigV2Spec {
	if in == nil {
		return nil
	}
	out := new(ConfigV2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigV2Status) DeepCopyInto(out *ConfigV2Status) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigV2Status.
func (in *ConfigV2Status) DeepCopy() *ConfigV2Status {
	if in == nil {
		return nil
	}
	out := new(ConfigV2Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DigitaloceanConfigObservation) DeepCopyInto(out *DigitaloceanConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DigitaloceanConfigObservation.
func (in *DigitaloceanConfigObservation) DeepCopy() *DigitaloceanConfigObservation {
	if in == nil {
		return nil
	}
	out := new(DigitaloceanConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DigitaloceanConfigParameters) DeepCopyInto(out *DigitaloceanConfigParameters) {
	*out = *in
	if in.AccessTokenSecretRef != nil {
		in, out := &in.AccessTokenSecretRef, &out.AccessTokenSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = new(bool)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(bool)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(bool)
		**out = **in
	}
	if in.PrivateNetworking != nil {
		in, out := &in.PrivateNetworking, &out.PrivateNetworking
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SSHKeyContentsSecretRef != nil {
		in, out := &in.SSHKeyContentsSecretRef, &out.SSHKeyContentsSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SSHKeyFingerprintSecretRef != nil {
		in, out := &in.SSHKeyFingerprintSecretRef, &out.SSHKeyFingerprintSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SSHPort != nil {
		in, out := &in.SSHPort, &out.SSHPort
		*out = new(string)
		**out = **in
	}
	if in.SSHUser != nil {
		in, out := &in.SSHUser, &out.SSHUser
		*out = new(string)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(string)
		**out = **in
	}
	if in.Userdata != nil {
		in, out := &in.Userdata, &out.Userdata
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DigitaloceanConfigParameters.
func (in *DigitaloceanConfigParameters) DeepCopy() *DigitaloceanConfigParameters {
	if in == nil {
		return nil
	}
	out := new(DigitaloceanConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinodeConfigObservation) DeepCopyInto(out *LinodeConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinodeConfigObservation.
func (in *LinodeConfigObservation) DeepCopy() *LinodeConfigObservation {
	if in == nil {
		return nil
	}
	out := new(LinodeConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinodeConfigParameters) DeepCopyInto(out *LinodeConfigParameters) {
	*out = *in
	if in.AuthorizedUsers != nil {
		in, out := &in.AuthorizedUsers, &out.AuthorizedUsers
		*out = new(string)
		**out = **in
	}
	if in.CreatePrivateIP != nil {
		in, out := &in.CreatePrivateIP, &out.CreatePrivateIP
		*out = new(bool)
		**out = **in
	}
	if in.DockerPort != nil {
		in, out := &in.DockerPort, &out.DockerPort
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RootPassSecretRef != nil {
		in, out := &in.RootPassSecretRef, &out.RootPassSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SSHPort != nil {
		in, out := &in.SSHPort, &out.SSHPort
		*out = new(string)
		**out = **in
	}
	if in.SSHUser != nil {
		in, out := &in.SSHUser, &out.SSHUser
		*out = new(string)
		**out = **in
	}
	if in.Stackscript != nil {
		in, out := &in.Stackscript, &out.Stackscript
		*out = new(string)
		**out = **in
	}
	if in.StackscriptData != nil {
		in, out := &in.StackscriptData, &out.StackscriptData
		*out = new(string)
		**out = **in
	}
	if in.SwapSize != nil {
		in, out := &in.SwapSize, &out.SwapSize
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(string)
		**out = **in
	}
	if in.TokenSecretRef != nil {
		in, out := &in.TokenSecretRef, &out.TokenSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.UaPrefix != nil {
		in, out := &in.UaPrefix, &out.UaPrefix
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinodeConfigParameters.
func (in *LinodeConfigParameters) DeepCopy() *LinodeConfigParameters {
	if in == nil {
		return nil
	}
	out := new(LinodeConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackConfigObservation) DeepCopyInto(out *OpenstackConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackConfigObservation.
func (in *OpenstackConfigObservation) DeepCopy() *OpenstackConfigObservation {
	if in == nil {
		return nil
	}
	out := new(OpenstackConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackConfigParameters) DeepCopyInto(out *OpenstackConfigParameters) {
	*out = *in
	if in.ActiveTimeout != nil {
		in, out := &in.ActiveTimeout, &out.ActiveTimeout
		*out = new(string)
		**out = **in
	}
	if in.ApplicationCredentialID != nil {
		in, out := &in.ApplicationCredentialID, &out.ApplicationCredentialID
		*out = new(string)
		**out = **in
	}
	if in.ApplicationCredentialName != nil {
		in, out := &in.ApplicationCredentialName, &out.ApplicationCredentialName
		*out = new(string)
		**out = **in
	}
	if in.ApplicationCredentialSecretSecretRef != nil {
		in, out := &in.ApplicationCredentialSecretSecretRef, &out.ApplicationCredentialSecretSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AuthURL != nil {
		in, out := &in.AuthURL, &out.AuthURL
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.BootFromVolume != nil {
		in, out := &in.BootFromVolume, &out.BootFromVolume
		*out = new(bool)
		**out = **in
	}
	if in.Cacert != nil {
		in, out := &in.Cacert, &out.Cacert
		*out = new(string)
		**out = **in
	}
	if in.ConfigDrive != nil {
		in, out := &in.ConfigDrive, &out.ConfigDrive
		*out = new(bool)
		**out = **in
	}
	if in.DomainID != nil {
		in, out := &in.DomainID, &out.DomainID
		*out = new(string)
		**out = **in
	}
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.EndpointType != nil {
		in, out := &in.EndpointType, &out.EndpointType
		*out = new(string)
		**out = **in
	}
	if in.FlavorID != nil {
		in, out := &in.FlavorID, &out.FlavorID
		*out = new(string)
		**out = **in
	}
	if in.FlavorName != nil {
		in, out := &in.FlavorName, &out.FlavorName
		*out = new(string)
		**out = **in
	}
	if in.FloatingIPPool != nil {
		in, out := &in.FloatingIPPool, &out.FloatingIPPool
		*out = new(string)
		**out = **in
	}
	if in.IPVersion != nil {
		in, out := &in.IPVersion, &out.IPVersion
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.ImageName != nil {
		in, out := &in.ImageName, &out.ImageName
		*out = new(string)
		**out = **in
	}
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		*out = new(bool)
		**out = **in
	}
	if in.KeypairName != nil {
		in, out := &in.KeypairName, &out.KeypairName
		*out = new(string)
		**out = **in
	}
	if in.NetID != nil {
		in, out := &in.NetID, &out.NetID
		*out = new(string)
		**out = **in
	}
	if in.NetName != nil {
		in, out := &in.NetName, &out.NetName
		*out = new(string)
		**out = **in
	}
	if in.NovaNetwork != nil {
		in, out := &in.NovaNetwork, &out.NovaNetwork
		*out = new(bool)
		**out = **in
	}
	if in.PasswordSecretRef != nil {
		in, out := &in.PasswordSecretRef, &out.PasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.PrivateKeyFileSecretRef != nil {
		in, out := &in.PrivateKeyFileSecretRef, &out.PrivateKeyFileSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.SSHPort != nil {
		in, out := &in.SSHPort, &out.SSHPort
		*out = new(string)
		**out = **in
	}
	if in.SSHUser != nil {
		in, out := &in.SSHUser, &out.SSHUser
		*out = new(string)
		**out = **in
	}
	if in.SecGroups != nil {
		in, out := &in.SecGroups, &out.SecGroups
		*out = new(string)
		**out = **in
	}
	if in.TenantDomainID != nil {
		in, out := &in.TenantDomainID, &out.TenantDomainID
		*out = new(string)
		**out = **in
	}
	if in.TenantDomainName != nil {
		in, out := &in.TenantDomainName, &out.TenantDomainName
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.TenantName != nil {
		in, out := &in.TenantName, &out.TenantName
		*out = new(string)
		**out = **in
	}
	if in.UserDataFile != nil {
		in, out := &in.UserDataFile, &out.UserDataFile
		*out = new(string)
		**out = **in
	}
	if in.UserDomainID != nil {
		in, out := &in.UserDomainID, &out.UserDomainID
		*out = new(string)
		**out = **in
	}
	if in.UserDomainName != nil {
		in, out := &in.UserDomainName, &out.UserDomainName
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.VolumeDevicePath != nil {
		in, out := &in.VolumeDevicePath, &out.VolumeDevicePath
		*out = new(string)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(string)
		**out = **in
	}
	if in.VolumeName != nil {
		in, out := &in.VolumeName, &out.VolumeName
		*out = new(string)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(string)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackConfigParameters.
func (in *OpenstackConfigParameters) DeepCopy() *OpenstackConfigParameters {
	if in == nil {
		return nil
	}
	out := new(OpenstackConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereConfigObservation) DeepCopyInto(out *VsphereConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereConfigObservation.
func (in *VsphereConfigObservation) DeepCopy() *VsphereConfigObservation {
	if in == nil {
		return nil
	}
	out := new(VsphereConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VsphereConfigParameters) DeepCopyInto(out *VsphereConfigParameters) {
	*out = *in
	if in.Boot2DockerURL != nil {
		in, out := &in.Boot2DockerURL, &out.Boot2DockerURL
		*out = new(string)
		**out = **in
	}
	if in.CPUCount != nil {
		in, out := &in.CPUCount, &out.CPUCount
		*out = new(string)
		**out = **in
	}
	if in.Cfgparam != nil {
		in, out := &in.Cfgparam, &out.Cfgparam
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CloneFrom != nil {
		in, out := &in.CloneFrom, &out.CloneFrom
		*out = new(string)
		**out = **in
	}
	if in.CloudConfig != nil {
		in, out := &in.CloudConfig, &out.CloudConfig
		*out = new(string)
		**out = **in
	}
	if in.Cloudinit != nil {
		in, out := &in.Cloudinit, &out.Cloudinit
		*out = new(string)
		**out = **in
	}
	if in.ContentLibrary != nil {
		in, out := &in.ContentLibrary, &out.ContentLibrary
		*out = new(string)
		**out = **in
	}
	if in.CreationType != nil {
		in, out := &in.CreationType, &out.CreationType
		*out = new(string)
		**out = **in
	}
	if in.CustomAttributes != nil {
		in, out := &in.CustomAttributes, &out.CustomAttributes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Datacenter != nil {
		in, out := &in.Datacenter, &out.Datacenter
		*out = new(string)
		**out = **in
	}
	if in.Datastore != nil {
		in, out := &in.Datastore, &out.Datastore
		*out = new(string)
		**out = **in
	}
	if in.DatastoreCluster != nil {
		in, out := &in.DatastoreCluster, &out.DatastoreCluster
		*out = new(string)
		**out = **in
	}
	if in.DiskSize != nil {
		in, out := &in.DiskSize, &out.DiskSize
		*out = new(string)
		**out = **in
	}
	if in.Folder != nil {
		in, out := &in.Folder, &out.Folder
		*out = new(string)
		**out = **in
	}
	if in.Hostsystem != nil {
		in, out := &in.Hostsystem, &out.Hostsystem
		*out = new(string)
		**out = **in
	}
	if in.MemorySize != nil {
		in, out := &in.MemorySize, &out.MemorySize
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PasswordSecretRef != nil {
		in, out := &in.PasswordSecretRef, &out.PasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = new(string)
		**out = **in
	}
	if in.SSHPasswordSecretRef != nil {
		in, out := &in.SSHPasswordSecretRef, &out.SSHPasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SSHPort != nil {
		in, out := &in.SSHPort, &out.SSHPort
		*out = new(string)
		**out = **in
	}
	if in.SSHUser != nil {
		in, out := &in.SSHUser, &out.SSHUser
		*out = new(string)
		**out = **in
	}
	if in.SSHUserGroup != nil {
		in, out := &in.SSHUserGroup, &out.SSHUserGroup
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.VappIPAllocationPolicy != nil {
		in, out := &in.VappIPAllocationPolicy, &out.VappIPAllocationPolicy
		*out = new(string)
		**out = **in
	}
	if in.VappIPProtocol != nil {
		in, out := &in.VappIPProtocol, &out.VappIPProtocol
		*out = new(string)
		**out = **in
	}
	if in.VappProperty != nil {
		in, out := &in.VappProperty, &out.VappProperty
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VappTransport != nil {
		in, out := &in.VappTransport, &out.VappTransport
		*out = new(string)
		**out = **in
	}
	if in.Vcenter != nil {
		in, out := &in.Vcenter, &out.Vcenter
		*out = new(string)
		**out = **in
	}
	if in.VcenterPort != nil {
		in, out := &in.VcenterPort, &out.VcenterPort
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VsphereConfigParameters.
func (in *VsphereConfigParameters) DeepCopy() *VsphereConfigParameters {
	if in == nil {
		return nil
	}
	out := new(VsphereConfigParameters)
	in.DeepCopyInto(out)
	return out
}
