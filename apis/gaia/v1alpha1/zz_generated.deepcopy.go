// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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
func (in *DeployFeature) DeepCopyInto(out *DeployFeature) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployFeature.
func (in *DeployFeature) DeepCopy() *DeployFeature {
	if in == nil {
		return nil
	}
	out := new(DeployFeature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileConf) DeepCopyInto(out *FileConf) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileConf.
func (in *FileConf) DeepCopy() *FileConf {
	if in == nil {
		return nil
	}
	out := new(FileConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GaiaCluster) DeepCopyInto(out *GaiaCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GaiaCluster.
func (in *GaiaCluster) DeepCopy() *GaiaCluster {
	if in == nil {
		return nil
	}
	out := new(GaiaCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GaiaCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GaiaClusterList) DeepCopyInto(out *GaiaClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GaiaCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GaiaClusterList.
func (in *GaiaClusterList) DeepCopy() *GaiaClusterList {
	if in == nil {
		return nil
	}
	out := new(GaiaClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GaiaClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GaiaClusterSpec) DeepCopyInto(out *GaiaClusterSpec) {
	*out = *in
	out.DeploymentFeature = in.DeploymentFeature
	if in.HostAliases != nil {
		in, out := &in.HostAliases, &out.HostAliases
		*out = make([]HostAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]GaiaNodeSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GaiaClusterSpec.
func (in *GaiaClusterSpec) DeepCopy() *GaiaClusterSpec {
	if in == nil {
		return nil
	}
	out := new(GaiaClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GaiaClusterStatus) DeepCopyInto(out *GaiaClusterStatus) {
	*out = *in
	if in.NodeStates != nil {
		in, out := &in.NodeStates, &out.NodeStates
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GaiaClusterStatus.
func (in *GaiaClusterStatus) DeepCopy() *GaiaClusterStatus {
	if in == nil {
		return nil
	}
	out := new(GaiaClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GaiaNode) DeepCopyInto(out *GaiaNode) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GaiaNode.
func (in *GaiaNode) DeepCopy() *GaiaNode {
	if in == nil {
		return nil
	}
	out := new(GaiaNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GaiaNode) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GaiaNodeList) DeepCopyInto(out *GaiaNodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GaiaNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GaiaNodeList.
func (in *GaiaNodeList) DeepCopy() *GaiaNodeList {
	if in == nil {
		return nil
	}
	out := new(GaiaNodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GaiaNodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GaiaNodeSpec) DeepCopyInto(out *GaiaNodeSpec) {
	*out = *in
	out.DeploymentFeature = in.DeploymentFeature
	in.Resource.DeepCopyInto(&out.Resource)
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make(map[string]Service, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.RemoveAction.DeepCopyInto(&out.RemoveAction)
	if in.NetworkCards != nil {
		in, out := &in.NetworkCards, &out.NetworkCards
		*out = make([]NetworkCardConf, len(*in))
		copy(*out, *in)
	}
	if in.PortForWards != nil {
		in, out := &in.PortForWards, &out.PortForWards
		*out = make([]PortForWardConf, len(*in))
		copy(*out, *in)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]FileConf, len(*in))
		copy(*out, *in)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]PortConf, len(*in))
		copy(*out, *in)
	}
	if in.HostAliases != nil {
		in, out := &in.HostAliases, &out.HostAliases
		*out = make([]HostAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]VolumeConf, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GaiaNodeSpec.
func (in *GaiaNodeSpec) DeepCopy() *GaiaNodeSpec {
	if in == nil {
		return nil
	}
	out := new(GaiaNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GaiaNodeStatus) DeepCopyInto(out *GaiaNodeStatus) {
	*out = *in
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]NetworkInfo, len(*in))
		copy(*out, *in)
	}
	if in.SvcStates != nil {
		in, out := &in.SvcStates, &out.SvcStates
		*out = make(map[string]SvcState, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PortForWardStates != nil {
		in, out := &in.PortForWardStates, &out.PortForWardStates
		*out = make(map[string]PortForWardState, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GaiaNodeStatus.
func (in *GaiaNodeStatus) DeepCopy() *GaiaNodeStatus {
	if in == nil {
		return nil
	}
	out := new(GaiaNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GaiaSet) DeepCopyInto(out *GaiaSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GaiaSet.
func (in *GaiaSet) DeepCopy() *GaiaSet {
	if in == nil {
		return nil
	}
	out := new(GaiaSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GaiaSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GaiaSetList) DeepCopyInto(out *GaiaSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GaiaSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GaiaSetList.
func (in *GaiaSetList) DeepCopy() *GaiaSetList {
	if in == nil {
		return nil
	}
	out := new(GaiaSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GaiaSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GaiaSetSpec) DeepCopyInto(out *GaiaSetSpec) {
	*out = *in
	out.DeploymentFeature = in.DeploymentFeature
	if in.HostAliases != nil {
		in, out := &in.HostAliases, &out.HostAliases
		*out = make([]HostAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GaiaSetSpec.
func (in *GaiaSetSpec) DeepCopy() *GaiaSetSpec {
	if in == nil {
		return nil
	}
	out := new(GaiaSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GaiaSetStatus) DeepCopyInto(out *GaiaSetStatus) {
	*out = *in
	if in.SvcStates != nil {
		in, out := &in.SvcStates, &out.SvcStates
		*out = make(map[string]SvcState, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GaiaSetStatus.
func (in *GaiaSetStatus) DeepCopy() *GaiaSetStatus {
	if in == nil {
		return nil
	}
	out := new(GaiaSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostAlias) DeepCopyInto(out *HostAlias) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostAlias.
func (in *HostAlias) DeepCopy() *HostAlias {
	if in == nil {
		return nil
	}
	out := new(HostAlias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkCardConf) DeepCopyInto(out *NetworkCardConf) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkCardConf.
func (in *NetworkCardConf) DeepCopy() *NetworkCardConf {
	if in == nil {
		return nil
	}
	out := new(NetworkCardConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkInfo) DeepCopyInto(out *NetworkInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkInfo.
func (in *NetworkInfo) DeepCopy() *NetworkInfo {
	if in == nil {
		return nil
	}
	out := new(NetworkInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortConf) DeepCopyInto(out *PortConf) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortConf.
func (in *PortConf) DeepCopy() *PortConf {
	if in == nil {
		return nil
	}
	out := new(PortConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortForWardConf) DeepCopyInto(out *PortForWardConf) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortForWardConf.
func (in *PortForWardConf) DeepCopy() *PortForWardConf {
	if in == nil {
		return nil
	}
	out := new(PortForWardConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortForWardState) DeepCopyInto(out *PortForWardState) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortForWardState.
func (in *PortForWardState) DeepCopy() *PortForWardState {
	if in == nil {
		return nil
	}
	out := new(PortForWardState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Process) DeepCopyInto(out *Process) {
	*out = *in
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Process.
func (in *Process) DeepCopy() *Process {
	if in == nil {
		return nil
	}
	out := new(Process)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	if in.Dependence != nil {
		in, out := &in.Dependence, &out.Dependence
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Srv.DeepCopyInto(&out.Srv)
	in.Init.DeepCopyInto(&out.Init)
	in.Check.DeepCopyInto(&out.Check)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeConf) DeepCopyInto(out *VolumeConf) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeConf.
func (in *VolumeConf) DeepCopy() *VolumeConf {
	if in == nil {
		return nil
	}
	out := new(VolumeConf)
	in.DeepCopyInto(out)
	return out
}