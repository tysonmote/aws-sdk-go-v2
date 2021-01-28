// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type IpAddressStatus string

// Enum values for IpAddressStatus
const (
	IpAddressStatusCreating               IpAddressStatus = "CREATING"
	IpAddressStatusFailedCreation         IpAddressStatus = "FAILED_CREATION"
	IpAddressStatusAttaching              IpAddressStatus = "ATTACHING"
	IpAddressStatusAttached               IpAddressStatus = "ATTACHED"
	IpAddressStatusRemapDetaching         IpAddressStatus = "REMAP_DETACHING"
	IpAddressStatusRemapAttaching         IpAddressStatus = "REMAP_ATTACHING"
	IpAddressStatusDetaching              IpAddressStatus = "DETACHING"
	IpAddressStatusFailedResourceGone     IpAddressStatus = "FAILED_RESOURCE_GONE"
	IpAddressStatusDeleting               IpAddressStatus = "DELETING"
	IpAddressStatusDeleteFailedFasExpired IpAddressStatus = "DELETE_FAILED_FAS_EXPIRED"
)

// Values returns all known values for IpAddressStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IpAddressStatus) Values() []IpAddressStatus {
	return []IpAddressStatus{
		"CREATING",
		"FAILED_CREATION",
		"ATTACHING",
		"ATTACHED",
		"REMAP_DETACHING",
		"REMAP_ATTACHING",
		"DETACHING",
		"FAILED_RESOURCE_GONE",
		"DELETING",
		"DELETE_FAILED_FAS_EXPIRED",
	}
}

type ResolverEndpointDirection string

// Enum values for ResolverEndpointDirection
const (
	ResolverEndpointDirectionInbound  ResolverEndpointDirection = "INBOUND"
	ResolverEndpointDirectionOutbound ResolverEndpointDirection = "OUTBOUND"
)

// Values returns all known values for ResolverEndpointDirection. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResolverEndpointDirection) Values() []ResolverEndpointDirection {
	return []ResolverEndpointDirection{
		"INBOUND",
		"OUTBOUND",
	}
}

type ResolverEndpointStatus string

// Enum values for ResolverEndpointStatus
const (
	ResolverEndpointStatusCreating       ResolverEndpointStatus = "CREATING"
	ResolverEndpointStatusOperational    ResolverEndpointStatus = "OPERATIONAL"
	ResolverEndpointStatusUpdating       ResolverEndpointStatus = "UPDATING"
	ResolverEndpointStatusAutoRecovering ResolverEndpointStatus = "AUTO_RECOVERING"
	ResolverEndpointStatusActionNeeded   ResolverEndpointStatus = "ACTION_NEEDED"
	ResolverEndpointStatusDeleting       ResolverEndpointStatus = "DELETING"
)

// Values returns all known values for ResolverEndpointStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResolverEndpointStatus) Values() []ResolverEndpointStatus {
	return []ResolverEndpointStatus{
		"CREATING",
		"OPERATIONAL",
		"UPDATING",
		"AUTO_RECOVERING",
		"ACTION_NEEDED",
		"DELETING",
	}
}

type ResolverQueryLogConfigAssociationError string

// Enum values for ResolverQueryLogConfigAssociationError
const (
	ResolverQueryLogConfigAssociationErrorNone                 ResolverQueryLogConfigAssociationError = "NONE"
	ResolverQueryLogConfigAssociationErrorDestinationNotFound  ResolverQueryLogConfigAssociationError = "DESTINATION_NOT_FOUND"
	ResolverQueryLogConfigAssociationErrorAccessDenied         ResolverQueryLogConfigAssociationError = "ACCESS_DENIED"
	ResolverQueryLogConfigAssociationErrorInternalServiceError ResolverQueryLogConfigAssociationError = "INTERNAL_SERVICE_ERROR"
)

// Values returns all known values for ResolverQueryLogConfigAssociationError. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ResolverQueryLogConfigAssociationError) Values() []ResolverQueryLogConfigAssociationError {
	return []ResolverQueryLogConfigAssociationError{
		"NONE",
		"DESTINATION_NOT_FOUND",
		"ACCESS_DENIED",
		"INTERNAL_SERVICE_ERROR",
	}
}

type ResolverQueryLogConfigAssociationStatus string

// Enum values for ResolverQueryLogConfigAssociationStatus
const (
	ResolverQueryLogConfigAssociationStatusCreating     ResolverQueryLogConfigAssociationStatus = "CREATING"
	ResolverQueryLogConfigAssociationStatusActive       ResolverQueryLogConfigAssociationStatus = "ACTIVE"
	ResolverQueryLogConfigAssociationStatusActionNeeded ResolverQueryLogConfigAssociationStatus = "ACTION_NEEDED"
	ResolverQueryLogConfigAssociationStatusDeleting     ResolverQueryLogConfigAssociationStatus = "DELETING"
	ResolverQueryLogConfigAssociationStatusFailed       ResolverQueryLogConfigAssociationStatus = "FAILED"
)

// Values returns all known values for ResolverQueryLogConfigAssociationStatus.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ResolverQueryLogConfigAssociationStatus) Values() []ResolverQueryLogConfigAssociationStatus {
	return []ResolverQueryLogConfigAssociationStatus{
		"CREATING",
		"ACTIVE",
		"ACTION_NEEDED",
		"DELETING",
		"FAILED",
	}
}

type ResolverQueryLogConfigStatus string

// Enum values for ResolverQueryLogConfigStatus
const (
	ResolverQueryLogConfigStatusCreating ResolverQueryLogConfigStatus = "CREATING"
	ResolverQueryLogConfigStatusCreated  ResolverQueryLogConfigStatus = "CREATED"
	ResolverQueryLogConfigStatusDeleting ResolverQueryLogConfigStatus = "DELETING"
	ResolverQueryLogConfigStatusFailed   ResolverQueryLogConfigStatus = "FAILED"
)

// Values returns all known values for ResolverQueryLogConfigStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResolverQueryLogConfigStatus) Values() []ResolverQueryLogConfigStatus {
	return []ResolverQueryLogConfigStatus{
		"CREATING",
		"CREATED",
		"DELETING",
		"FAILED",
	}
}

type ResolverRuleAssociationStatus string

// Enum values for ResolverRuleAssociationStatus
const (
	ResolverRuleAssociationStatusCreating   ResolverRuleAssociationStatus = "CREATING"
	ResolverRuleAssociationStatusComplete   ResolverRuleAssociationStatus = "COMPLETE"
	ResolverRuleAssociationStatusDeleting   ResolverRuleAssociationStatus = "DELETING"
	ResolverRuleAssociationStatusFailed     ResolverRuleAssociationStatus = "FAILED"
	ResolverRuleAssociationStatusOverridden ResolverRuleAssociationStatus = "OVERRIDDEN"
)

// Values returns all known values for ResolverRuleAssociationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ResolverRuleAssociationStatus) Values() []ResolverRuleAssociationStatus {
	return []ResolverRuleAssociationStatus{
		"CREATING",
		"COMPLETE",
		"DELETING",
		"FAILED",
		"OVERRIDDEN",
	}
}

type ResolverRuleStatus string

// Enum values for ResolverRuleStatus
const (
	ResolverRuleStatusComplete ResolverRuleStatus = "COMPLETE"
	ResolverRuleStatusDeleting ResolverRuleStatus = "DELETING"
	ResolverRuleStatusUpdating ResolverRuleStatus = "UPDATING"
	ResolverRuleStatusFailed   ResolverRuleStatus = "FAILED"
)

// Values returns all known values for ResolverRuleStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResolverRuleStatus) Values() []ResolverRuleStatus {
	return []ResolverRuleStatus{
		"COMPLETE",
		"DELETING",
		"UPDATING",
		"FAILED",
	}
}

type RuleTypeOption string

// Enum values for RuleTypeOption
const (
	RuleTypeOptionForward   RuleTypeOption = "FORWARD"
	RuleTypeOptionSystem    RuleTypeOption = "SYSTEM"
	RuleTypeOptionRecursive RuleTypeOption = "RECURSIVE"
)

// Values returns all known values for RuleTypeOption. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RuleTypeOption) Values() []RuleTypeOption {
	return []RuleTypeOption{
		"FORWARD",
		"SYSTEM",
		"RECURSIVE",
	}
}

type ShareStatus string

// Enum values for ShareStatus
const (
	ShareStatusNotShared    ShareStatus = "NOT_SHARED"
	ShareStatusSharedWithMe ShareStatus = "SHARED_WITH_ME"
	ShareStatusSharedByMe   ShareStatus = "SHARED_BY_ME"
)

// Values returns all known values for ShareStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ShareStatus) Values() []ShareStatus {
	return []ShareStatus{
		"NOT_SHARED",
		"SHARED_WITH_ME",
		"SHARED_BY_ME",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "ASCENDING"
	SortOrderDescending SortOrder = "DESCENDING"
)

// Values returns all known values for SortOrder. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASCENDING",
		"DESCENDING",
	}
}