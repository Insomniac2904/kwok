//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KwokConfiguration) DeepCopyInto(out *KwokConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Options.DeepCopyInto(&out.Options)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KwokConfiguration.
func (in *KwokConfiguration) DeepCopy() *KwokConfiguration {
	if in == nil {
		return nil
	}
	out := new(KwokConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KwokConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KwokConfigurationOptions) DeepCopyInto(out *KwokConfigurationOptions) {
	*out = *in
	if in.ManageAllNodes != nil {
		in, out := &in.ManageAllNodes, &out.ManageAllNodes
		*out = new(bool)
		**out = **in
	}
	if in.EnableCNI != nil {
		in, out := &in.EnableCNI, &out.EnableCNI
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KwokConfigurationOptions.
func (in *KwokConfigurationOptions) DeepCopy() *KwokConfigurationOptions {
	if in == nil {
		return nil
	}
	out := new(KwokConfigurationOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KwokctlConfiguration) DeepCopyInto(out *KwokctlConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Options.DeepCopyInto(&out.Options)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KwokctlConfiguration.
func (in *KwokctlConfiguration) DeepCopy() *KwokctlConfiguration {
	if in == nil {
		return nil
	}
	out := new(KwokctlConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KwokctlConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KwokctlConfigurationOptions) DeepCopyInto(out *KwokctlConfigurationOptions) {
	*out = *in
	if in.SecurePort != nil {
		in, out := &in.SecurePort, &out.SecurePort
		*out = new(bool)
		**out = **in
	}
	if in.QuietPull != nil {
		in, out := &in.QuietPull, &out.QuietPull
		*out = new(bool)
		**out = **in
	}
	if in.DisableKubeScheduler != nil {
		in, out := &in.DisableKubeScheduler, &out.DisableKubeScheduler
		*out = new(bool)
		**out = **in
	}
	if in.DisableKubeControllerManager != nil {
		in, out := &in.DisableKubeControllerManager, &out.DisableKubeControllerManager
		*out = new(bool)
		**out = **in
	}
	if in.KubeAuthorization != nil {
		in, out := &in.KubeAuthorization, &out.KubeAuthorization
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KwokctlConfigurationOptions.
func (in *KwokctlConfigurationOptions) DeepCopy() *KwokctlConfigurationOptions {
	if in == nil {
		return nil
	}
	out := new(KwokctlConfigurationOptions)
	in.DeepCopyInto(out)
	return out
}