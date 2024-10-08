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

type PermissionFeatureSet string

const (
	PermissionFeatureSet_CREATED_FROM_POLICY   PermissionFeatureSet = "CREATED_FROM_POLICY"
	PermissionFeatureSet_PROMOTING_TO_STANDARD PermissionFeatureSet = "PROMOTING_TO_STANDARD"
	PermissionFeatureSet_STANDARD              PermissionFeatureSet = "STANDARD"
)

type PermissionStatus string

const (
	PermissionStatus_ATTACHABLE   PermissionStatus = "ATTACHABLE"
	PermissionStatus_UNATTACHABLE PermissionStatus = "UNATTACHABLE"
	PermissionStatus_DELETING     PermissionStatus = "DELETING"
	PermissionStatus_DELETED      PermissionStatus = "DELETED"
)

type PermissionType string

const (
	PermissionType_CUSTOMER_MANAGED PermissionType = "CUSTOMER_MANAGED"
	PermissionType_AWS_MANAGED      PermissionType = "AWS_MANAGED"
)

type PermissionTypeFilter string

const (
	PermissionTypeFilter_ALL              PermissionTypeFilter = "ALL"
	PermissionTypeFilter_AWS_MANAGED      PermissionTypeFilter = "AWS_MANAGED"
	PermissionTypeFilter_CUSTOMER_MANAGED PermissionTypeFilter = "CUSTOMER_MANAGED"
)

type ReplacePermissionAssociationsWorkStatus string

const (
	ReplacePermissionAssociationsWorkStatus_IN_PROGRESS ReplacePermissionAssociationsWorkStatus = "IN_PROGRESS"
	ReplacePermissionAssociationsWorkStatus_COMPLETED   ReplacePermissionAssociationsWorkStatus = "COMPLETED"
	ReplacePermissionAssociationsWorkStatus_FAILED      ReplacePermissionAssociationsWorkStatus = "FAILED"
)

type ResourceOwner string

const (
	ResourceOwner_SELF           ResourceOwner = "SELF"
	ResourceOwner_OTHER_ACCOUNTS ResourceOwner = "OTHER-ACCOUNTS"
)

type ResourceRegionScope string

const (
	ResourceRegionScope_REGIONAL ResourceRegionScope = "REGIONAL"
	ResourceRegionScope_GLOBAL   ResourceRegionScope = "GLOBAL"
)

type ResourceRegionScopeFilter string

const (
	ResourceRegionScopeFilter_ALL      ResourceRegionScopeFilter = "ALL"
	ResourceRegionScopeFilter_REGIONAL ResourceRegionScopeFilter = "REGIONAL"
	ResourceRegionScopeFilter_GLOBAL   ResourceRegionScopeFilter = "GLOBAL"
)

type ResourceShareAssociationStatus string

const (
	ResourceShareAssociationStatus_ASSOCIATING    ResourceShareAssociationStatus = "ASSOCIATING"
	ResourceShareAssociationStatus_ASSOCIATED     ResourceShareAssociationStatus = "ASSOCIATED"
	ResourceShareAssociationStatus_FAILED         ResourceShareAssociationStatus = "FAILED"
	ResourceShareAssociationStatus_DISASSOCIATING ResourceShareAssociationStatus = "DISASSOCIATING"
	ResourceShareAssociationStatus_DISASSOCIATED  ResourceShareAssociationStatus = "DISASSOCIATED"
)

type ResourceShareAssociationType string

const (
	ResourceShareAssociationType_PRINCIPAL ResourceShareAssociationType = "PRINCIPAL"
	ResourceShareAssociationType_RESOURCE  ResourceShareAssociationType = "RESOURCE"
)

type ResourceShareFeatureSet string

const (
	ResourceShareFeatureSet_CREATED_FROM_POLICY   ResourceShareFeatureSet = "CREATED_FROM_POLICY"
	ResourceShareFeatureSet_PROMOTING_TO_STANDARD ResourceShareFeatureSet = "PROMOTING_TO_STANDARD"
	ResourceShareFeatureSet_STANDARD              ResourceShareFeatureSet = "STANDARD"
)

type ResourceShareInvitationStatus string

const (
	ResourceShareInvitationStatus_PENDING  ResourceShareInvitationStatus = "PENDING"
	ResourceShareInvitationStatus_ACCEPTED ResourceShareInvitationStatus = "ACCEPTED"
	ResourceShareInvitationStatus_REJECTED ResourceShareInvitationStatus = "REJECTED"
	ResourceShareInvitationStatus_EXPIRED  ResourceShareInvitationStatus = "EXPIRED"
)

type ResourceShareStatus_SDK string

const (
	ResourceShareStatus_SDK_PENDING  ResourceShareStatus_SDK = "PENDING"
	ResourceShareStatus_SDK_ACTIVE   ResourceShareStatus_SDK = "ACTIVE"
	ResourceShareStatus_SDK_FAILED   ResourceShareStatus_SDK = "FAILED"
	ResourceShareStatus_SDK_DELETING ResourceShareStatus_SDK = "DELETING"
	ResourceShareStatus_SDK_DELETED  ResourceShareStatus_SDK = "DELETED"
)

type ResourceStatus string

const (
	ResourceStatus_AVAILABLE                   ResourceStatus = "AVAILABLE"
	ResourceStatus_ZONAL_RESOURCE_INACCESSIBLE ResourceStatus = "ZONAL_RESOURCE_INACCESSIBLE"
	ResourceStatus_LIMIT_EXCEEDED              ResourceStatus = "LIMIT_EXCEEDED"
	ResourceStatus_UNAVAILABLE                 ResourceStatus = "UNAVAILABLE"
	ResourceStatus_PENDING                     ResourceStatus = "PENDING"
)
