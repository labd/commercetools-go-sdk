package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type CustomerGroup struct {
	// Unique identifier of the CustomerGroup.
	ID string `json:"id"`
	// Current version of the CustomerGroup.
	Version int `json:"version"`
	// Date and time (UTC) the CustomerGroup was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the CustomerGroup was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier for the CustomerGroup.
	Key *string `json:"key,omitempty"`
	// Unique name of the CustomerGroup.
	Name string `json:"name"`
	// Custom Fields for the CustomerGroup.
	Custom *CustomFields `json:"custom,omitempty"`
}

type CustomerGroupDraft struct {
	// User-defined unique identifier for the CustomerGroup.
	Key *string `json:"key,omitempty"`
	// Unique value which must be different from any value used for `name` in [CustomerGroup](ctp:api:type:CustomerGroup) in the Project.
	// If not, a `DuplicateField` [error](/../api/errors#400-bad-request-1) is thrown.
	GroupName string `json:"groupName"`
	// Custom Fields for the CustomerGroup.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

/**
*	[PagedQueryResult](/general-concepts#pagedqueryresult) with `results` containing an array of [CustomerGroup](ctp:api:type:CustomerGroup).
*
 */
type CustomerGroupPagedQueryResponse struct {
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
	// [CustomerGroups](ctp:api:type:CustomerGroup) matching the query.
	Results []CustomerGroup `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [CustomerGroup](ctp:api:type:CustomerGroup).
*
 */
type CustomerGroupReference struct {
	// Unique identifier of the referenced [CustomerGroup](ctp:api:type:CustomerGroup).
	ID string `json:"id"`
	// Contains the representation of the expanded CustomerGroup. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for CustomerGroups.
	Obj *CustomerGroup `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerGroupReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "customer-group", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [CustomerGroup](ctp:api:type:CustomerGroup).
*
 */
type CustomerGroupResourceIdentifier struct {
	// Unique identifier of the referenced [CustomerGroup](ctp:api:type:CustomerGroup). Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [CustomerGroup](ctp:api:type:CustomerGroup). Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerGroupResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "customer-group", Alias: (*Alias)(&obj)})
}

type CustomerGroupUpdate struct {
	// Expected version of the CustomerGroup on which the changes should be applied.
	// If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the CustomerGroup.
	Actions []CustomerGroupUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerGroupUpdate) UnmarshalJSON(data []byte) error {
	type Alias CustomerGroupUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorCustomerGroupUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type CustomerGroupUpdateAction interface{}

func mapDiscriminatorCustomerGroupUpdateAction(input interface{}) (CustomerGroupUpdateAction, error) {
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
	case "changeName":
		obj := CustomerGroupChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := CustomerGroupSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := CustomerGroupSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := CustomerGroupSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type CustomerGroupChangeNameAction struct {
	// New name of the CustomerGroup.
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerGroupChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type CustomerGroupSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerGroupSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

/**
*	This action sets or removes the custom type for an existing CustomerGroup.
*	If present, this action overwrites any existing [custom](/../api/projects/custom-fields) type and fields.
*
 */
type CustomerGroupSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the CustomerGroup with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the CustomerGroup.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the CustomerGroup.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerGroupSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type CustomerGroupSetKeyAction struct {
	// If `key` is absent or `null`, the existing key, if any, will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomerGroupSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}
