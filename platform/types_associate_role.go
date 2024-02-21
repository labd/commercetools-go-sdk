package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type AssociateRole struct {
	// Unique identifier of the AssociateRole.
	ID string `json:"id"`
	// Current version of the AssociateRole.
	Version int `json:"version"`
	// Date and time (UTC) the AssociateRole was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the AssociateRole was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1 February 2019 except for [events not tracked](/../api/general-concepts#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/general-concepts#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the AssociateRole.
	Key string `json:"key"`
	// Whether the AssociateRole can be assigned to an Associate by a [buyer](/../api/associates-overview#buyer). If false, the AssociateRole can only be assigned using the [general endpoint](/../api/associates-overview#through-the-general-endpoints).
	BuyerAssignable bool `json:"buyerAssignable"`
	// Name of the AssociateRole.
	Name *string `json:"name,omitempty"`
	// List of Permissions for the AssociateRole.
	Permissions []Permission `json:"permissions"`
	// Custom Fields for the AssociateRole.
	Custom *CustomFields `json:"custom,omitempty"`
}

type AssociateRoleDraft struct {
	// User-defined unique identifier for the AssociateRole.
	Key string `json:"key"`
	// Name of the AssociateRole.
	Name *string `json:"name,omitempty"`
	// Whether the AssociateRole can be assigned to an Associate by a [buyer](/../api/associates-overview#buyer).
	BuyerAssignable *bool `json:"buyerAssignable,omitempty"`
	// List of Permissions for the AssociateRole.
	Permissions []Permission `json:"permissions"`
	// Custom Fields for the AssociateRole.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AssociateRoleDraft) MarshalJSON() ([]byte, error) {
	type Alias AssociateRoleDraft
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["permissions"] == nil {
		delete(raw, "permissions")
	}

	return json.Marshal(raw)

}

/**
*	[Reference](ctp:api:type:Reference) to an [AssociateRole](ctp:api:type:AssociateRole) by its key.
*
 */
type AssociateRoleKeyReference struct {
	// Unique and immutable key of the referenced [AssociateRole](ctp:api:type:AssociateRole).
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AssociateRoleKeyReference) MarshalJSON() ([]byte, error) {
	type Alias AssociateRoleKeyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "associate-role", Alias: (*Alias)(&obj)})
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [AssociateRole](ctp:api:type:AssociateRole).
*
 */
type AssociateRolePagedQueryResponse struct {
	// Number of requested [results](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of elements [skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [AssociateRoles](ctp:api:type:AssociateRole) matching the query.
	Results []AssociateRole `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to an [AssociateRole](ctp:api:type:AssociateRole).
*
 */
type AssociateRoleReference struct {
	// Unique identifier of the referenced [AssociateRole](ctp:api:type:AssociateRole).
	ID string `json:"id"`
	// Contains the representation of the expanded AssociateRole. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for AssociateRole.
	Obj *AssociateRole `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AssociateRoleReference) MarshalJSON() ([]byte, error) {
	type Alias AssociateRoleReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "associate-role", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:TypeResourceIdentifier) of an [AssociateRole](ctp:api:type:AssociateRole). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type AssociateRoleResourceIdentifier struct {
	// Unique identifier of the referenced [AssociateRole](ctp:api:type:AssociateRole). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced [AssociateRole](ctp:api:type:AssociateRole). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AssociateRoleResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias AssociateRoleResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "associate-role", Alias: (*Alias)(&obj)})
}

type AssociateRoleUpdate struct {
	// Expected version of the AssociateRole on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the AssociateRole.
	Actions []AssociateRoleUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AssociateRoleUpdate) UnmarshalJSON(data []byte) error {
	type Alias AssociateRoleUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorAssociateRoleUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type AssociateRoleUpdateAction interface{}

func mapDiscriminatorAssociateRoleUpdateAction(input interface{}) (AssociateRoleUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "addPermission":
		obj := AssociateRoleAddPermissionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeBuyerAssignable":
		obj := AssociateRoleChangeBuyerAssignableAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removePermission":
		obj := AssociateRoleRemovePermissionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := AssociateRoleSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := AssociateRoleSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setName":
		obj := AssociateRoleSetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setPermissions":
		obj := AssociateRoleSetPermissionsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Permissions grant granular access to [Approval Rules](ctp:api:type:ApprovalRule), [Approval Flows](ctp:api:type:ApprovalFlow), [Business Units](ctp:api:type:BusinessUnit), [Carts](ctp:api:type:Cart), [Orders](ctp:api:type:Order), [Quotes](ctp:api:type:Quote), and [Quote Requests](ctp:api:type:QuoteRequest).
*
 */
type Permission string

const (
	PermissionAddChildUnits                      Permission = "AddChildUnits"
	PermissionUpdateAssociates                   Permission = "UpdateAssociates"
	PermissionUpdateBusinessUnitDetails          Permission = "UpdateBusinessUnitDetails"
	PermissionUpdateParentUnit                   Permission = "UpdateParentUnit"
	PermissionViewMyCarts                        Permission = "ViewMyCarts"
	PermissionViewOthersCarts                    Permission = "ViewOthersCarts"
	PermissionUpdateMyCarts                      Permission = "UpdateMyCarts"
	PermissionUpdateOthersCarts                  Permission = "UpdateOthersCarts"
	PermissionCreateMyCarts                      Permission = "CreateMyCarts"
	PermissionCreateOthersCarts                  Permission = "CreateOthersCarts"
	PermissionDeleteMyCarts                      Permission = "DeleteMyCarts"
	PermissionDeleteOthersCarts                  Permission = "DeleteOthersCarts"
	PermissionViewMyOrders                       Permission = "ViewMyOrders"
	PermissionViewOthersOrders                   Permission = "ViewOthersOrders"
	PermissionUpdateMyOrders                     Permission = "UpdateMyOrders"
	PermissionUpdateOthersOrders                 Permission = "UpdateOthersOrders"
	PermissionCreateMyOrdersFromMyCarts          Permission = "CreateMyOrdersFromMyCarts"
	PermissionCreateMyOrdersFromMyQuotes         Permission = "CreateMyOrdersFromMyQuotes"
	PermissionCreateOrdersFromOthersCarts        Permission = "CreateOrdersFromOthersCarts"
	PermissionCreateOrdersFromOthersQuotes       Permission = "CreateOrdersFromOthersQuotes"
	PermissionViewMyQuotes                       Permission = "ViewMyQuotes"
	PermissionViewOthersQuotes                   Permission = "ViewOthersQuotes"
	PermissionAcceptMyQuotes                     Permission = "AcceptMyQuotes"
	PermissionAcceptOthersQuotes                 Permission = "AcceptOthersQuotes"
	PermissionDeclineMyQuotes                    Permission = "DeclineMyQuotes"
	PermissionDeclineOthersQuotes                Permission = "DeclineOthersQuotes"
	PermissionRenegotiateMyQuotes                Permission = "RenegotiateMyQuotes"
	PermissionRenegotiateOthersQuotes            Permission = "RenegotiateOthersQuotes"
	PermissionReassignMyQuotes                   Permission = "ReassignMyQuotes"
	PermissionReassignOthersQuotes               Permission = "ReassignOthersQuotes"
	PermissionViewMyQuoteRequests                Permission = "ViewMyQuoteRequests"
	PermissionViewOthersQuoteRequests            Permission = "ViewOthersQuoteRequests"
	PermissionUpdateMyQuoteRequests              Permission = "UpdateMyQuoteRequests"
	PermissionUpdateOthersQuoteRequests          Permission = "UpdateOthersQuoteRequests"
	PermissionCreateMyQuoteRequestsFromMyCarts   Permission = "CreateMyQuoteRequestsFromMyCarts"
	PermissionCreateQuoteRequestsFromOthersCarts Permission = "CreateQuoteRequestsFromOthersCarts"
	PermissionCreateApprovalRules                Permission = "CreateApprovalRules"
	PermissionUpdateApprovalRules                Permission = "UpdateApprovalRules"
	PermissionUpdateApprovalFlows                Permission = "UpdateApprovalFlows"
)

/**
*	Adding a Permission to an [AssociateRole](ctp:api:type:AssociateRole) generates an [AssociateRolePermissionAdded](ctp:api:type:AssociateRolePermissionAddedMessage) Message.
*
 */
type AssociateRoleAddPermissionAction struct {
	// Permission to be added to the AssociateRole.
	Permission Permission `json:"permission"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AssociateRoleAddPermissionAction) MarshalJSON() ([]byte, error) {
	type Alias AssociateRoleAddPermissionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPermission", Alias: (*Alias)(&obj)})
}

/**
*	Changing the `buyerAssignable` value of an AssociateRole generates an [AssociateRoleBuyerAssignableChanged](ctp:api:type:AssociateRoleBuyerAssignableChangedMessage) Message.
*
 */
type AssociateRoleChangeBuyerAssignableAction struct {
	// The new value of the `buyerAssignable` field of the AssociateRole.
	BuyerAssignable bool `json:"buyerAssignable"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AssociateRoleChangeBuyerAssignableAction) MarshalJSON() ([]byte, error) {
	type Alias AssociateRoleChangeBuyerAssignableAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeBuyerAssignable", Alias: (*Alias)(&obj)})
}

/**
*	Removing a Permission from an [AssociateRole](ctp:api:type:AssociateRole) generates an [AssociateRolePermissionRemoved](ctp:api:type:AssociateRolePermissionRemovedMessage) Message.
*
 */
type AssociateRoleRemovePermissionAction struct {
	// Permission to be removed from the AssociateRole.
	Permission Permission `json:"permission"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AssociateRoleRemovePermissionAction) MarshalJSON() ([]byte, error) {
	type Alias AssociateRoleRemovePermissionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePermission", Alias: (*Alias)(&obj)})
}

type AssociateRoleSetCustomFieldAction struct {
	// Name of the [Custom Field](ctp:api:type:CustomFields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperationError](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AssociateRoleSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias AssociateRoleSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type AssociateRoleSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the AssociateRole with [Custom Fields](ctp:api:type:CustomFields).
	// If absent, any existing Type and Custom Fields are removed from the AssociateRole.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](ctp:api:type:CustomFields) for the AssociateRole.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AssociateRoleSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias AssociateRoleSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

/**
*	Updating the name of an [AssociateRole](ctp:api:type:AssociateRole) generates an [AssociateRoleNameSet](ctp:api:type:AssociateRoleNameSetMessage) Message.
*
 */
type AssociateRoleSetNameAction struct {
	// New name to set.
	// If `name` is absent or `null`, the existing name, if any, will be removed.
	Name *string `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AssociateRoleSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias AssociateRoleSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

/**
*	Updating the Permissions on an [AssociateRole](ctp:api:type:AssociateRole) generates an [AssociateRolePermissionsSet](ctp:api:type:AssociateRolePermissionsSetMessage) Message.
*
 */
type AssociateRoleSetPermissionsAction struct {
	// Overrides the current list of Permissions for the AssociateRole.
	Permissions []Permission `json:"permissions"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AssociateRoleSetPermissionsAction) MarshalJSON() ([]byte, error) {
	type Alias AssociateRoleSetPermissionsAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPermissions", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["permissions"] == nil {
		delete(raw, "permissions")
	}

	return json.Marshal(raw)

}
