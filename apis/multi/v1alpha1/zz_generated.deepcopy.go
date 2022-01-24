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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnswersObservation) DeepCopyInto(out *AnswersObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnswersObservation.
func (in *AnswersObservation) DeepCopy() *AnswersObservation {
	if in == nil {
		return nil
	}
	out := new(AnswersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnswersParameters) DeepCopyInto(out *AnswersParameters) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnswersParameters.
func (in *AnswersParameters) DeepCopy() *AnswersParameters {
	if in == nil {
		return nil
	}
	out := new(AnswersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterApp) DeepCopyInto(out *ClusterApp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterApp.
func (in *ClusterApp) DeepCopy() *ClusterApp {
	if in == nil {
		return nil
	}
	out := new(ClusterApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterApp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAppList) DeepCopyInto(out *ClusterAppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterApp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAppList.
func (in *ClusterAppList) DeepCopy() *ClusterAppList {
	if in == nil {
		return nil
	}
	out := new(ClusterAppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAppObservation) DeepCopyInto(out *ClusterAppObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.TemplateVersionID != nil {
		in, out := &in.TemplateVersionID, &out.TemplateVersionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAppObservation.
func (in *ClusterAppObservation) DeepCopy() *ClusterAppObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterAppObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAppParameters) DeepCopyInto(out *ClusterAppParameters) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Answers != nil {
		in, out := &in.Answers, &out.Answers
		*out = make([]AnswersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CatalogName != nil {
		in, out := &in.CatalogName, &out.CatalogName
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
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]MembersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int64)
		**out = **in
	}
	if in.RevisionID != nil {
		in, out := &in.RevisionID, &out.RevisionID
		*out = new(string)
		**out = **in
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]TargetsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TemplateName != nil {
		in, out := &in.TemplateName, &out.TemplateName
		*out = new(string)
		**out = **in
	}
	if in.TemplateVersion != nil {
		in, out := &in.TemplateVersion, &out.TemplateVersion
		*out = new(string)
		**out = **in
	}
	if in.UpgradeStrategy != nil {
		in, out := &in.UpgradeStrategy, &out.UpgradeStrategy
		*out = make([]UpgradeStrategyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Wait != nil {
		in, out := &in.Wait, &out.Wait
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAppParameters.
func (in *ClusterAppParameters) DeepCopy() *ClusterAppParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterAppParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAppSpec) DeepCopyInto(out *ClusterAppSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAppSpec.
func (in *ClusterAppSpec) DeepCopy() *ClusterAppSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterAppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAppStatus) DeepCopyInto(out *ClusterAppStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAppStatus.
func (in *ClusterAppStatus) DeepCopy() *ClusterAppStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterAppStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembersObservation) DeepCopyInto(out *MembersObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembersObservation.
func (in *MembersObservation) DeepCopy() *MembersObservation {
	if in == nil {
		return nil
	}
	out := new(MembersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembersParameters) DeepCopyInto(out *MembersParameters) {
	*out = *in
	if in.AccessType != nil {
		in, out := &in.AccessType, &out.AccessType
		*out = new(string)
		**out = **in
	}
	if in.GroupPrincipalID != nil {
		in, out := &in.GroupPrincipalID, &out.GroupPrincipalID
		*out = new(string)
		**out = **in
	}
	if in.UserPrincipalID != nil {
		in, out := &in.UserPrincipalID, &out.UserPrincipalID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembersParameters.
func (in *MembersParameters) DeepCopy() *MembersParameters {
	if in == nil {
		return nil
	}
	out := new(MembersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdateObservation) DeepCopyInto(out *RollingUpdateObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdateObservation.
func (in *RollingUpdateObservation) DeepCopy() *RollingUpdateObservation {
	if in == nil {
		return nil
	}
	out := new(RollingUpdateObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdateParameters) DeepCopyInto(out *RollingUpdateParameters) {
	*out = *in
	if in.BatchSize != nil {
		in, out := &in.BatchSize, &out.BatchSize
		*out = new(int64)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdateParameters.
func (in *RollingUpdateParameters) DeepCopy() *RollingUpdateParameters {
	if in == nil {
		return nil
	}
	out := new(RollingUpdateParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetsObservation) DeepCopyInto(out *TargetsObservation) {
	*out = *in
	if in.AppID != nil {
		in, out := &in.AppID, &out.AppID
		*out = new(string)
		**out = **in
	}
	if in.HealthState != nil {
		in, out := &in.HealthState, &out.HealthState
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetsObservation.
func (in *TargetsObservation) DeepCopy() *TargetsObservation {
	if in == nil {
		return nil
	}
	out := new(TargetsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetsParameters) DeepCopyInto(out *TargetsParameters) {
	*out = *in
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetsParameters.
func (in *TargetsParameters) DeepCopy() *TargetsParameters {
	if in == nil {
		return nil
	}
	out := new(TargetsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeStrategyObservation) DeepCopyInto(out *UpgradeStrategyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeStrategyObservation.
func (in *UpgradeStrategyObservation) DeepCopy() *UpgradeStrategyObservation {
	if in == nil {
		return nil
	}
	out := new(UpgradeStrategyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeStrategyParameters) DeepCopyInto(out *UpgradeStrategyParameters) {
	*out = *in
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		*out = make([]RollingUpdateParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeStrategyParameters.
func (in *UpgradeStrategyParameters) DeepCopy() *UpgradeStrategyParameters {
	if in == nil {
		return nil
	}
	out := new(UpgradeStrategyParameters)
	in.DeepCopyInto(out)
	return out
}
