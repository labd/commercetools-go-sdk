package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type DiscountCode struct {
	// Unique identifier of the DiscountCode.
	ID string `json:"id"`
	// Current version of the DiscountCode.
	Version int `json:"version"`
	// Date and time (UTC) the DiscountCode was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the DiscountCode was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Name of the DiscountCode.
	Name *LocalizedString `json:"name,omitempty"`
	// Description of the DiscountCode.
	Description *LocalizedString `json:"description,omitempty"`
	// User-defined unique identifier of the DiscountCode [added to the Cart](/../api/projects/carts#add-discountcode) to apply the related [CartDiscounts](ctp:api:type:CartDiscount).
	Code string `json:"code"`
	// Reference to CartDiscounts that can be applied to the Cart once the DiscountCode is applied.
	CartDiscounts []CartDiscountReference `json:"cartDiscounts"`
	// DiscountCode can only be applied to Carts that match this predicate.
	CartPredicate *string `json:"cartPredicate,omitempty"`
	// Indicates if the DiscountCode is active and can be applied to the Cart.
	IsActive bool `json:"isActive"`
	// Array generated from the Cart predicate.
	// It contains the references of all the resources that are addressed in the predicate.
	References []Reference `json:"references"`
	// Number of times the DiscountCode can be applied.
	// DiscountCode application is counted at the time of Order creation or edit. However, Order cancellation or deletion does not decrement the count.
	MaxApplications *int `json:"maxApplications,omitempty"`
	// Number of times the DiscountCode can be applied per Customer (anonymous Carts are not supported).
	// DiscountCode application is counted at the time of Order creation or edit. However, Order cancellation or deletion does not decrement the count.
	MaxApplicationsPerCustomer *int `json:"maxApplicationsPerCustomer,omitempty"`
	// Custom Fields of the DiscountCode.
	Custom *CustomFields `json:"custom,omitempty"`
	// Groups to which the DiscountCode belongs to.
	Groups []string `json:"groups"`
	// Date and time (UTC) from which the DiscountCode is effective.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Date and time (UTC) until which the DiscountCode is effective.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Used and managed by the API and must not be used in customer logic.
	// The value can change at any time due to internal and external factors.
	ApplicationVersion *int `json:"applicationVersion,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountCode) UnmarshalJSON(data []byte) error {
	type Alias DiscountCode
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.References {
		var err error
		obj.References[i], err = mapDiscriminatorReference(obj.References[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type DiscountCodeDraft struct {
	// Name of the DiscountCode.
	Name *LocalizedString `json:"name,omitempty"`
	// Description of the DiscountCode.
	Description *LocalizedString `json:"description,omitempty"`
	// User-defined unique identifier for the DiscountCode that can be [added to the Cart](/../api/projects/carts#add-discountcode) to apply the related [CartDiscounts](ctp:api:type:CartDiscount).
	// It cannot be modified after the DiscountCode is created.
	Code string `json:"code"`
	// Specify what CartDiscounts the API applies when you add the DiscountCode to the Cart.
	CartDiscounts []CartDiscountResourceIdentifier `json:"cartDiscounts"`
	// DiscountCode can only be applied to Carts that match this predicate.
	CartPredicate *string `json:"cartPredicate,omitempty"`
	// Only active DiscountCodes can be applied to the Cart.
	IsActive *bool `json:"isActive,omitempty"`
	// Number of times the DiscountCode can be applied.
	MaxApplications *int `json:"maxApplications,omitempty"`
	// Number of times the DiscountCode can be applied per Customer.
	MaxApplicationsPerCustomer *int `json:"maxApplicationsPerCustomer,omitempty"`
	// Custom Fields for the DiscountCode.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Groups to which the DiscountCode will belong to.
	Groups []string `json:"groups"`
	// Date and time (UTC) from which the DiscountCode is effective. Must be earlier than `validUntil`.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Date and time (UTC) until which the DiscountCode is effective. Must be later than `validFrom`.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeDraft) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeDraft
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

	if raw["groups"] == nil {
		delete(raw, "groups")
	}

	return json.Marshal(raw)

}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [DiscountCode](ctp:api:type:DiscountCode).
*
 */
type DiscountCodePagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [DiscountCodes](ctp:api:type:DiscountCode) matching the query.
	Results []DiscountCode `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [DiscountCode](ctp:api:type:DiscountCode).
*
 */
type DiscountCodeReference struct {
	// Unique identifier of the referenced [DiscountCode](ctp:api:type:DiscountCode).
	ID string `json:"id"`
	// Contains the representation of the expanded DiscountCode. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for DiscountCodes.
	Obj *DiscountCode `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeReference) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "discount-code", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [DiscountCode](ctp:api:type:DiscountCode). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type DiscountCodeResourceIdentifier struct {
	// Unique identifier of the referenced [DiscountCode](ctp:api:type:DiscountCode). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [DiscountCode](ctp:api:type:DiscountCode). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "discount-code", Alias: (*Alias)(&obj)})
}

type DiscountCodeUpdate struct {
	// Expected version of the DiscountCode on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error is returned.
	Version int `json:"version"`
	// Update actions to be performed on the DiscountCode.
	Actions []DiscountCodeUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountCodeUpdate) UnmarshalJSON(data []byte) error {
	type Alias DiscountCodeUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorDiscountCodeUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type DiscountCodeUpdateAction interface{}

func mapDiscriminatorDiscountCodeUpdateAction(input interface{}) (DiscountCodeUpdateAction, error) {
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
	case "changeCartDiscounts":
		obj := DiscountCodeChangeCartDiscountsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeGroups":
		obj := DiscountCodeChangeGroupsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeIsActive":
		obj := DiscountCodeChangeIsActiveAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCartPredicate":
		obj := DiscountCodeSetCartPredicateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := DiscountCodeSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := DiscountCodeSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := DiscountCodeSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMaxApplications":
		obj := DiscountCodeSetMaxApplicationsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMaxApplicationsPerCustomer":
		obj := DiscountCodeSetMaxApplicationsPerCustomerAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setName":
		obj := DiscountCodeSetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidFrom":
		obj := DiscountCodeSetValidFromAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidFromAndUntil":
		obj := DiscountCodeSetValidFromAndUntilAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidUntil":
		obj := DiscountCodeSetValidUntilAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type DiscountCodeChangeCartDiscountsAction struct {
	// New value to set.
	CartDiscounts []CartDiscountResourceIdentifier `json:"cartDiscounts"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeChangeCartDiscountsAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeChangeCartDiscountsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCartDiscounts", Alias: (*Alias)(&obj)})
}

type DiscountCodeChangeGroupsAction struct {
	// New value to set. An empty array removes the DiscountCode from all groups.
	Groups []string `json:"groups"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeChangeGroupsAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeChangeGroupsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeGroups", Alias: (*Alias)(&obj)})
}

type DiscountCodeChangeIsActiveAction struct {
	// New value to set. Set to `true` to activate the DiscountCode for all matching Discounts.
	IsActive bool `json:"isActive"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeChangeIsActiveAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeChangeIsActiveAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsActive", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetCartPredicateAction struct {
	// Value to set. If empty, any existing value will be removed.
	CartPredicate *string `json:"cartPredicate,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetCartPredicateAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetCartPredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCartPredicate", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the DiscountCode with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the DiscountCode.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the DiscountCode.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetDescriptionAction struct {
	// Value to set. If empty, any existing value will be removed.
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetMaxApplicationsAction struct {
	// Value to set. If empty, any existing value will be removed.
	MaxApplications *int `json:"maxApplications,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetMaxApplicationsAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetMaxApplicationsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMaxApplications", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetMaxApplicationsPerCustomerAction struct {
	// Value to set. If empty, any existing value will be removed.
	MaxApplicationsPerCustomer *int `json:"maxApplicationsPerCustomer,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetMaxApplicationsPerCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetMaxApplicationsPerCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMaxApplicationsPerCustomer", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetNameAction struct {
	// Value to set. If empty, any existing value will be removed.
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetValidFromAction struct {
	// Value to set that must be earlier than `validUntil`. If empty, any existing value will be removed.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetValidFromAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetValidFromAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFrom", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetValidFromAndUntilAction struct {
	// Value to set that must be earlier than `validUntil`. If empty, any existing value will be removed.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Value to set that must be later than `validFrom`. If empty, any existing value will be removed.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetValidFromAndUntilAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetValidFromAndUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFromAndUntil", Alias: (*Alias)(&obj)})
}

type DiscountCodeSetValidUntilAction struct {
	// Value to set that must be later than `validFrom`. If empty, any existing value will be removed.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountCodeSetValidUntilAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountCodeSetValidUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidUntil", Alias: (*Alias)(&obj)})
}
