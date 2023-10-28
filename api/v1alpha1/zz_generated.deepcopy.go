//go:build !ignore_autogenerated

/*
Copyright 2023.

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
func (in *ScwExternalIP) DeepCopyInto(out *ScwExternalIP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScwExternalIP.
func (in *ScwExternalIP) DeepCopy() *ScwExternalIP {
	if in == nil {
		return nil
	}
	out := new(ScwExternalIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScwExternalIP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScwExternalIPList) DeepCopyInto(out *ScwExternalIPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScwExternalIP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScwExternalIPList.
func (in *ScwExternalIPList) DeepCopy() *ScwExternalIPList {
	if in == nil {
		return nil
	}
	out := new(ScwExternalIPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScwExternalIPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScwExternalIPSpec) DeepCopyInto(out *ScwExternalIPSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScwExternalIPSpec.
func (in *ScwExternalIPSpec) DeepCopy() *ScwExternalIPSpec {
	if in == nil {
		return nil
	}
	out := new(ScwExternalIPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScwExternalIPStatus) DeepCopyInto(out *ScwExternalIPStatus) {
	*out = *in
	if in.AttachedIPs != nil {
		in, out := &in.AttachedIPs, &out.AttachedIPs
		*out = make([]ScwExternalIPStatusAttachedIP, len(*in))
		copy(*out, *in)
	}
	if in.PendingIPs != nil {
		in, out := &in.PendingIPs, &out.PendingIPs
		*out = make([]ScwExternalIPStatusPendingIP, len(*in))
		copy(*out, *in)
	}
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScwExternalIPStatus.
func (in *ScwExternalIPStatus) DeepCopy() *ScwExternalIPStatus {
	if in == nil {
		return nil
	}
	out := new(ScwExternalIPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScwExternalIPStatusAttachedIP) DeepCopyInto(out *ScwExternalIPStatusAttachedIP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScwExternalIPStatusAttachedIP.
func (in *ScwExternalIPStatusAttachedIP) DeepCopy() *ScwExternalIPStatusAttachedIP {
	if in == nil {
		return nil
	}
	out := new(ScwExternalIPStatusAttachedIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScwExternalIPStatusPendingIP) DeepCopyInto(out *ScwExternalIPStatusPendingIP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScwExternalIPStatusPendingIP.
func (in *ScwExternalIPStatusPendingIP) DeepCopy() *ScwExternalIPStatusPendingIP {
	if in == nil {
		return nil
	}
	out := new(ScwExternalIPStatusPendingIP)
	in.DeepCopyInto(out)
	return out
}