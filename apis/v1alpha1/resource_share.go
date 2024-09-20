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

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ResourceShareSpec defines the desired state of ResourceShare.
//
// Describes a resource share in RAM.
type ResourceShareSpec struct {

	// Specifies whether principals outside your organization in Organizations can
	// be associated with a resource share. A value of true lets you share with
	// individual Amazon Web Services accounts that are not in your organization.
	// A value of false only has meaning if your account is a member of an Amazon
	// Web Services Organization. The default value is true.
	AllowExternalPrincipals *bool `json:"allowExternalPrincipals,omitempty"`
	// Specifies the name of the resource share.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// Specifies the Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the RAM permission to associate with the resource share. If you do not
	// specify an ARN for the permission, RAM automatically attaches the default
	// version of the permission for each resource type. You can associate only
	// one permission with each resource type included in the resource share.
	PermissionARNs []*string                                  `json:"permissionARNs,omitempty"`
	PermissionRefs []*ackv1alpha1.AWSResourceReferenceWrapper `json:"permissionRefs,omitempty"`
	// Specifies a list of one or more principals to associate with the resource
	// share.
	//
	// You can include the following values:
	//
	//   - An Amazon Web Services account ID, for example: 123456789012
	//
	//   - An Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	//     of an organization in Organizations, for example: organizations::123456789012:organization/o-exampleorgid
	//
	//   - An ARN of an organizational unit (OU) in Organizations, for example:
	//     organizations::123456789012:ou/o-exampleorgid/ou-examplerootid-exampleouid123
	//
	//   - An ARN of an IAM role, for example: iam::123456789012:role/rolename
	//
	//   - An ARN of an IAM user, for example: iam::123456789012user/username
	//
	// Not all resource types can be shared with IAM roles and users. For more information,
	// see Sharing with IAM roles and users (https://docs.aws.amazon.com/ram/latest/userguide/permissions.html#permissions-rbp-supported-resource-types)
	// in the Resource Access Manager User Guide.
	Principals []*string `json:"principals,omitempty"`
	// Specifies a list of one or more ARNs of the resources to associate with the
	// resource share.
	ResourceARNs []*string `json:"resourceARNs,omitempty"`
	// Specifies from which source accounts the service principal has access to
	// the resources in this resource share.
	Sources []*string `json:"sources,omitempty"`
	// A list of one or more tag key and value pairs. The tag key must be present
	// and not be an empty string. The tag value must be present but can be an empty
	// string.
	Tags []*Tag `json:"tags,omitempty"`
}

// ResourceShareStatus defines the observed state of ResourceShare
type ResourceShareStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The date and time when the resource share was created.
	// +kubebuilder:validation:Optional
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
	// Indicates what features are available for this resource share. This parameter
	// can have one of the following values:
	//
	//    * STANDARD – A resource share that supports all functionality. These
	//    resource shares are visible to all principals you share the resource share
	//    with. You can modify these resource shares in RAM using the console or
	//    APIs. This resource share might have been created by RAM, or it might
	//    have been CREATED_FROM_POLICY and then promoted.
	//
	//    * CREATED_FROM_POLICY – The customer manually shared a resource by attaching
	//    a resource-based policy. That policy did not match any existing managed
	//    permissions, so RAM created this customer managed permission automatically
	//    on the customer's behalf based on the attached policy document. This type
	//    of resource share is visible only to the Amazon Web Services account that
	//    created it. You can't modify it in RAM unless you promote it. For more
	//    information, see PromoteResourceShareCreatedFromPolicy.
	//
	//    * PROMOTING_TO_STANDARD – This resource share was originally CREATED_FROM_POLICY,
	//    but the customer ran the PromoteResourceShareCreatedFromPolicy and that
	//    operation is still in progress. This value changes to STANDARD when complete.
	// +kubebuilder:validation:Optional
	FeatureSet *string `json:"featureSet,omitempty"`
	// The date and time when the resource share was last updated.
	// +kubebuilder:validation:Optional
	LastUpdatedTime *metav1.Time `json:"lastUpdatedTime,omitempty"`
	// The ID of the Amazon Web Services account that owns the resource share.
	// +kubebuilder:validation:Optional
	OwningAccountID *string `json:"owningAccountID,omitempty"`
	// The current status of the resource share.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty"`
	// A message about the status of the resource share.
	// +kubebuilder:validation:Optional
	StatusMessage *string `json:"statusMessage,omitempty"`
}

// ResourceShare is the Schema for the ResourceShares API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type ResourceShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceShareSpec   `json:"spec,omitempty"`
	Status            ResourceShareStatus `json:"status,omitempty"`
}

// ResourceShareList contains a list of ResourceShare
// +kubebuilder:object:root=true
type ResourceShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceShare `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ResourceShare{}, &ResourceShareList{})
}
