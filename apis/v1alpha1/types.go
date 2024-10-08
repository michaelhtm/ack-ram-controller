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
	"github.com/aws/aws-sdk-go/aws"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

// An object that describes a managed permission associated with a resource
// share.
type AssociatedPermission struct {
	ARN               *string      `json:"arn,omitempty"`
	DefaultVersion    *bool        `json:"defaultVersion,omitempty"`
	LastUpdatedTime   *metav1.Time `json:"lastUpdatedTime,omitempty"`
	PermissionVersion *string      `json:"permissionVersion,omitempty"`
	ResourceShareARN  *string      `json:"resourceShareARN,omitempty"`
	ResourceType      *string      `json:"resourceType,omitempty"`
	Status            *string      `json:"status,omitempty"`
}

// Describes a principal for use with Resource Access Manager.
type Principal struct {
	CreationTime     *metav1.Time `json:"creationTime,omitempty"`
	External         *bool        `json:"external,omitempty"`
	ID               *string      `json:"id,omitempty"`
	LastUpdatedTime  *metav1.Time `json:"lastUpdatedTime,omitempty"`
	ResourceShareARN *string      `json:"resourceShareARN,omitempty"`
}

// A structure that represents the background work that RAM performs when you
// invoke the ReplacePermissionAssociations operation.
type ReplacePermissionAssociationsWork struct {
	CreationTime          *metav1.Time `json:"creationTime,omitempty"`
	FromPermissionARN     *string      `json:"fromPermissionARN,omitempty"`
	FromPermissionVersion *string      `json:"fromPermissionVersion,omitempty"`
	ID                    *string      `json:"id,omitempty"`
	LastUpdatedTime       *metav1.Time `json:"lastUpdatedTime,omitempty"`
	StatusMessage         *string      `json:"statusMessage,omitempty"`
	ToPermissionARN       *string      `json:"toPermissionARN,omitempty"`
	ToPermissionVersion   *string      `json:"toPermissionVersion,omitempty"`
}

// Describes a resource associated with a resource share in RAM.
type Resource struct {
	ARN              *string      `json:"arn,omitempty"`
	CreationTime     *metav1.Time `json:"creationTime,omitempty"`
	LastUpdatedTime  *metav1.Time `json:"lastUpdatedTime,omitempty"`
	ResourceGroupARN *string      `json:"resourceGroupARN,omitempty"`
	ResourceShareARN *string      `json:"resourceShareARN,omitempty"`
	StatusMessage    *string      `json:"statusMessage,omitempty"`
	Type             *string      `json:"type_,omitempty"`
}

// Describes an association between a resource share and either a principal
// or a resource.
type ResourceShareAssociation struct {
	AssociatedEntity  *string      `json:"associatedEntity,omitempty"`
	CreationTime      *metav1.Time `json:"creationTime,omitempty"`
	External          *bool        `json:"external,omitempty"`
	LastUpdatedTime   *metav1.Time `json:"lastUpdatedTime,omitempty"`
	ResourceShareARN  *string      `json:"resourceShareARN,omitempty"`
	ResourceShareName *string      `json:"resourceShareName,omitempty"`
	StatusMessage     *string      `json:"statusMessage,omitempty"`
}

// Describes an invitation for an Amazon Web Services account to join a resource
// share.
type ResourceShareInvitation struct {
	InvitationTimestamp        *metav1.Time `json:"invitationTimestamp,omitempty"`
	ReceiverAccountID          *string      `json:"receiverAccountID,omitempty"`
	ReceiverARN                *string      `json:"receiverARN,omitempty"`
	ResourceShareARN           *string      `json:"resourceShareARN,omitempty"`
	ResourceShareInvitationARN *string      `json:"resourceShareInvitationARN,omitempty"`
	ResourceShareName          *string      `json:"resourceShareName,omitempty"`
	SenderAccountID            *string      `json:"senderAccountID,omitempty"`
}

// Information about a RAM managed permission.
type ResourceSharePermissionDetail struct {
	ARN                   *string      `json:"arn,omitempty"`
	CreationTime          *metav1.Time `json:"creationTime,omitempty"`
	DefaultVersion        *bool        `json:"defaultVersion,omitempty"`
	IsResourceTypeDefault *bool        `json:"isResourceTypeDefault,omitempty"`
	LastUpdatedTime       *metav1.Time `json:"lastUpdatedTime,omitempty"`
	Name                  *string      `json:"name,omitempty"`
	Permission            *string      `json:"permission,omitempty"`
	ResourceType          *string      `json:"resourceType,omitempty"`
	Tags                  []*Tag       `json:"tags,omitempty"`
	Version               *string      `json:"version,omitempty"`
}

// Information about an RAM permission.
type ResourceSharePermissionSummary struct {
	ARN                   *string      `json:"arn,omitempty"`
	CreationTime          *metav1.Time `json:"creationTime,omitempty"`
	DefaultVersion        *bool        `json:"defaultVersion,omitempty"`
	IsResourceTypeDefault *bool        `json:"isResourceTypeDefault,omitempty"`
	LastUpdatedTime       *metav1.Time `json:"lastUpdatedTime,omitempty"`
	Name                  *string      `json:"name,omitempty"`
	ResourceType          *string      `json:"resourceType,omitempty"`
	Status                *string      `json:"status,omitempty"`
	Tags                  []*Tag       `json:"tags,omitempty"`
	Version               *string      `json:"version,omitempty"`
}

// Describes a resource share in RAM.
type ResourceShare_SDK struct {
	AllowExternalPrincipals *bool        `json:"allowExternalPrincipals,omitempty"`
	CreationTime            *metav1.Time `json:"creationTime,omitempty"`
	FeatureSet              *string      `json:"featureSet,omitempty"`
	LastUpdatedTime         *metav1.Time `json:"lastUpdatedTime,omitempty"`
	Name                    *string      `json:"name,omitempty"`
	OwningAccountID         *string      `json:"owningAccountID,omitempty"`
	ResourceShareARN        *string      `json:"resourceShareARN,omitempty"`
	Status                  *string      `json:"status,omitempty"`
	StatusMessage           *string      `json:"statusMessage,omitempty"`
	Tags                    []*Tag       `json:"tags,omitempty"`
}

// Information about a shareable resource type and the Amazon Web Services service
// to which resources of that type belong.
type ServiceNameAndResourceType struct {
	ResourceType *string `json:"resourceType,omitempty"`
	ServiceName  *string `json:"serviceName,omitempty"`
}

// A structure containing a tag. A tag is metadata that you can attach to your
// resources to help organize and categorize them. You can also use them to
// help you secure your resources. For more information, see Controlling access
// to Amazon Web Services resources using tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html).
//
// For more information about tags, see Tagging Amazon Web Services resources
// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the Amazon
// Web Services General Reference Guide.
type Tag struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// A tag key and optional list of possible values that you can use to filter
// results for tagged resources.
type TagFilter struct {
	TagKey    *string   `json:"tagKey,omitempty"`
	TagValues []*string `json:"tagValues,omitempty"`
}
