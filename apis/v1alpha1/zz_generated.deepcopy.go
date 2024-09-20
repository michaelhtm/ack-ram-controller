//go:build !ignore_autogenerated

// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssociatedPermission) DeepCopyInto(out *AssociatedPermission) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.DefaultVersion != nil {
		in, out := &in.DefaultVersion, &out.DefaultVersion
		*out = new(bool)
		**out = **in
	}
	if in.FeatureSet != nil {
		in, out := &in.FeatureSet, &out.FeatureSet
		*out = new(string)
		**out = **in
	}
	if in.LastUpdatedTime != nil {
		in, out := &in.LastUpdatedTime, &out.LastUpdatedTime
		*out = (*in).DeepCopy()
	}
	if in.PermissionVersion != nil {
		in, out := &in.PermissionVersion, &out.PermissionVersion
		*out = new(string)
		**out = **in
	}
	if in.ResourceShareARN != nil {
		in, out := &in.ResourceShareARN, &out.ResourceShareARN
		*out = new(string)
		**out = **in
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssociatedPermission.
func (in *AssociatedPermission) DeepCopy() *AssociatedPermission {
	if in == nil {
		return nil
	}
	out := new(AssociatedPermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Permission) DeepCopyInto(out *Permission) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Permission.
func (in *Permission) DeepCopy() *Permission {
	if in == nil {
		return nil
	}
	out := new(Permission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Permission) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionList) DeepCopyInto(out *PermissionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Permission, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionList.
func (in *PermissionList) DeepCopy() *PermissionList {
	if in == nil {
		return nil
	}
	out := new(PermissionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PermissionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionSpec) DeepCopyInto(out *PermissionSpec) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PolicyTemplate != nil {
		in, out := &in.PolicyTemplate, &out.PolicyTemplate
		*out = new(string)
		**out = **in
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionSpec.
func (in *PermissionSpec) DeepCopy() *PermissionSpec {
	if in == nil {
		return nil
	}
	out := new(PermissionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionStatus) DeepCopyInto(out *PermissionStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.DefaultVersion != nil {
		in, out := &in.DefaultVersion, &out.DefaultVersion
		*out = new(bool)
		**out = **in
	}
	if in.FeatureSet != nil {
		in, out := &in.FeatureSet, &out.FeatureSet
		*out = new(string)
		**out = **in
	}
	if in.IsResourceTypeDefault != nil {
		in, out := &in.IsResourceTypeDefault, &out.IsResourceTypeDefault
		*out = new(bool)
		**out = **in
	}
	if in.LastUpdatedTime != nil {
		in, out := &in.LastUpdatedTime, &out.LastUpdatedTime
		*out = (*in).DeepCopy()
	}
	if in.PermissionType != nil {
		in, out := &in.PermissionType, &out.PermissionType
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionStatus.
func (in *PermissionStatus) DeepCopy() *PermissionStatus {
	if in == nil {
		return nil
	}
	out := new(PermissionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Principal) DeepCopyInto(out *Principal) {
	*out = *in
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LastUpdatedTime != nil {
		in, out := &in.LastUpdatedTime, &out.LastUpdatedTime
		*out = (*in).DeepCopy()
	}
	if in.ResourceShareARN != nil {
		in, out := &in.ResourceShareARN, &out.ResourceShareARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Principal.
func (in *Principal) DeepCopy() *Principal {
	if in == nil {
		return nil
	}
	out := new(Principal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplacePermissionAssociationsWork) DeepCopyInto(out *ReplacePermissionAssociationsWork) {
	*out = *in
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.FromPermissionARN != nil {
		in, out := &in.FromPermissionARN, &out.FromPermissionARN
		*out = new(string)
		**out = **in
	}
	if in.FromPermissionVersion != nil {
		in, out := &in.FromPermissionVersion, &out.FromPermissionVersion
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LastUpdatedTime != nil {
		in, out := &in.LastUpdatedTime, &out.LastUpdatedTime
		*out = (*in).DeepCopy()
	}
	if in.StatusMessage != nil {
		in, out := &in.StatusMessage, &out.StatusMessage
		*out = new(string)
		**out = **in
	}
	if in.ToPermissionARN != nil {
		in, out := &in.ToPermissionARN, &out.ToPermissionARN
		*out = new(string)
		**out = **in
	}
	if in.ToPermissionVersion != nil {
		in, out := &in.ToPermissionVersion, &out.ToPermissionVersion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplacePermissionAssociationsWork.
func (in *ReplacePermissionAssociationsWork) DeepCopy() *ReplacePermissionAssociationsWork {
	if in == nil {
		return nil
	}
	out := new(ReplacePermissionAssociationsWork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.LastUpdatedTime != nil {
		in, out := &in.LastUpdatedTime, &out.LastUpdatedTime
		*out = (*in).DeepCopy()
	}
	if in.ResourceGroupARN != nil {
		in, out := &in.ResourceGroupARN, &out.ResourceGroupARN
		*out = new(string)
		**out = **in
	}
	if in.ResourceShareARN != nil {
		in, out := &in.ResourceShareARN, &out.ResourceShareARN
		*out = new(string)
		**out = **in
	}
	if in.StatusMessage != nil {
		in, out := &in.StatusMessage, &out.StatusMessage
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceShare) DeepCopyInto(out *ResourceShare) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceShare.
func (in *ResourceShare) DeepCopy() *ResourceShare {
	if in == nil {
		return nil
	}
	out := new(ResourceShare)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceShare) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceShareAssociation) DeepCopyInto(out *ResourceShareAssociation) {
	*out = *in
	if in.AssociatedEntity != nil {
		in, out := &in.AssociatedEntity, &out.AssociatedEntity
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(bool)
		**out = **in
	}
	if in.LastUpdatedTime != nil {
		in, out := &in.LastUpdatedTime, &out.LastUpdatedTime
		*out = (*in).DeepCopy()
	}
	if in.ResourceShareARN != nil {
		in, out := &in.ResourceShareARN, &out.ResourceShareARN
		*out = new(string)
		**out = **in
	}
	if in.ResourceShareName != nil {
		in, out := &in.ResourceShareName, &out.ResourceShareName
		*out = new(string)
		**out = **in
	}
	if in.StatusMessage != nil {
		in, out := &in.StatusMessage, &out.StatusMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceShareAssociation.
func (in *ResourceShareAssociation) DeepCopy() *ResourceShareAssociation {
	if in == nil {
		return nil
	}
	out := new(ResourceShareAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceShareInvitation) DeepCopyInto(out *ResourceShareInvitation) {
	*out = *in
	if in.InvitationTimestamp != nil {
		in, out := &in.InvitationTimestamp, &out.InvitationTimestamp
		*out = (*in).DeepCopy()
	}
	if in.ReceiverAccountID != nil {
		in, out := &in.ReceiverAccountID, &out.ReceiverAccountID
		*out = new(string)
		**out = **in
	}
	if in.ReceiverARN != nil {
		in, out := &in.ReceiverARN, &out.ReceiverARN
		*out = new(string)
		**out = **in
	}
	if in.ResourceShareARN != nil {
		in, out := &in.ResourceShareARN, &out.ResourceShareARN
		*out = new(string)
		**out = **in
	}
	if in.ResourceShareInvitationARN != nil {
		in, out := &in.ResourceShareInvitationARN, &out.ResourceShareInvitationARN
		*out = new(string)
		**out = **in
	}
	if in.ResourceShareName != nil {
		in, out := &in.ResourceShareName, &out.ResourceShareName
		*out = new(string)
		**out = **in
	}
	if in.SenderAccountID != nil {
		in, out := &in.SenderAccountID, &out.SenderAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceShareInvitation.
func (in *ResourceShareInvitation) DeepCopy() *ResourceShareInvitation {
	if in == nil {
		return nil
	}
	out := new(ResourceShareInvitation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceShareList) DeepCopyInto(out *ResourceShareList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceShare, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceShareList.
func (in *ResourceShareList) DeepCopy() *ResourceShareList {
	if in == nil {
		return nil
	}
	out := new(ResourceShareList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceShareList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSharePermissionDetail) DeepCopyInto(out *ResourceSharePermissionDetail) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.DefaultVersion != nil {
		in, out := &in.DefaultVersion, &out.DefaultVersion
		*out = new(bool)
		**out = **in
	}
	if in.FeatureSet != nil {
		in, out := &in.FeatureSet, &out.FeatureSet
		*out = new(string)
		**out = **in
	}
	if in.IsResourceTypeDefault != nil {
		in, out := &in.IsResourceTypeDefault, &out.IsResourceTypeDefault
		*out = new(bool)
		**out = **in
	}
	if in.LastUpdatedTime != nil {
		in, out := &in.LastUpdatedTime, &out.LastUpdatedTime
		*out = (*in).DeepCopy()
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Permission != nil {
		in, out := &in.Permission, &out.Permission
		*out = new(string)
		**out = **in
	}
	if in.PermissionType != nil {
		in, out := &in.PermissionType, &out.PermissionType
		*out = new(string)
		**out = **in
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSharePermissionDetail.
func (in *ResourceSharePermissionDetail) DeepCopy() *ResourceSharePermissionDetail {
	if in == nil {
		return nil
	}
	out := new(ResourceSharePermissionDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSharePermissionSummary) DeepCopyInto(out *ResourceSharePermissionSummary) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.DefaultVersion != nil {
		in, out := &in.DefaultVersion, &out.DefaultVersion
		*out = new(bool)
		**out = **in
	}
	if in.FeatureSet != nil {
		in, out := &in.FeatureSet, &out.FeatureSet
		*out = new(string)
		**out = **in
	}
	if in.IsResourceTypeDefault != nil {
		in, out := &in.IsResourceTypeDefault, &out.IsResourceTypeDefault
		*out = new(bool)
		**out = **in
	}
	if in.LastUpdatedTime != nil {
		in, out := &in.LastUpdatedTime, &out.LastUpdatedTime
		*out = (*in).DeepCopy()
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PermissionType != nil {
		in, out := &in.PermissionType, &out.PermissionType
		*out = new(string)
		**out = **in
	}
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSharePermissionSummary.
func (in *ResourceSharePermissionSummary) DeepCopy() *ResourceSharePermissionSummary {
	if in == nil {
		return nil
	}
	out := new(ResourceSharePermissionSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceShareSpec) DeepCopyInto(out *ResourceShareSpec) {
	*out = *in
	if in.AllowExternalPrincipals != nil {
		in, out := &in.AllowExternalPrincipals, &out.AllowExternalPrincipals
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PermissionARNs != nil {
		in, out := &in.PermissionARNs, &out.PermissionARNs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PermissionRefs != nil {
		in, out := &in.PermissionRefs, &out.PermissionRefs
		*out = make([]*corev1alpha1.AWSResourceReferenceWrapper, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.AWSResourceReferenceWrapper)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Principals != nil {
		in, out := &in.Principals, &out.Principals
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ResourceARNs != nil {
		in, out := &in.ResourceARNs, &out.ResourceARNs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceShareSpec.
func (in *ResourceShareSpec) DeepCopy() *ResourceShareSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceShareSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceShareStatus) DeepCopyInto(out *ResourceShareStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.FeatureSet != nil {
		in, out := &in.FeatureSet, &out.FeatureSet
		*out = new(string)
		**out = **in
	}
	if in.LastUpdatedTime != nil {
		in, out := &in.LastUpdatedTime, &out.LastUpdatedTime
		*out = (*in).DeepCopy()
	}
	if in.OwningAccountID != nil {
		in, out := &in.OwningAccountID, &out.OwningAccountID
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.StatusMessage != nil {
		in, out := &in.StatusMessage, &out.StatusMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceShareStatus.
func (in *ResourceShareStatus) DeepCopy() *ResourceShareStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceShareStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceShare_SDK) DeepCopyInto(out *ResourceShare_SDK) {
	*out = *in
	if in.AllowExternalPrincipals != nil {
		in, out := &in.AllowExternalPrincipals, &out.AllowExternalPrincipals
		*out = new(bool)
		**out = **in
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
	if in.FeatureSet != nil {
		in, out := &in.FeatureSet, &out.FeatureSet
		*out = new(string)
		**out = **in
	}
	if in.LastUpdatedTime != nil {
		in, out := &in.LastUpdatedTime, &out.LastUpdatedTime
		*out = (*in).DeepCopy()
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OwningAccountID != nil {
		in, out := &in.OwningAccountID, &out.OwningAccountID
		*out = new(string)
		**out = **in
	}
	if in.ResourceShareARN != nil {
		in, out := &in.ResourceShareARN, &out.ResourceShareARN
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.StatusMessage != nil {
		in, out := &in.StatusMessage, &out.StatusMessage
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceShare_SDK.
func (in *ResourceShare_SDK) DeepCopy() *ResourceShare_SDK {
	if in == nil {
		return nil
	}
	out := new(ResourceShare_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceNameAndResourceType) DeepCopyInto(out *ServiceNameAndResourceType) {
	*out = *in
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceNameAndResourceType.
func (in *ServiceNameAndResourceType) DeepCopy() *ServiceNameAndResourceType {
	if in == nil {
		return nil
	}
	out := new(ServiceNameAndResourceType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagFilter) DeepCopyInto(out *TagFilter) {
	*out = *in
	if in.TagKey != nil {
		in, out := &in.TagKey, &out.TagKey
		*out = new(string)
		**out = **in
	}
	if in.TagValues != nil {
		in, out := &in.TagValues, &out.TagValues
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagFilter.
func (in *TagFilter) DeepCopy() *TagFilter {
	if in == nil {
		return nil
	}
	out := new(TagFilter)
	in.DeepCopyInto(out)
	return out
}
