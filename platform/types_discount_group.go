package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type DiscountGroup struct {
	// Unique identifier of the DiscountGroup.
	ID string `json:"id"`
	// Current version of the DiscountGroup.
	Version int `json:"version"`
	// Date and time (UTC) the DiscountGroup was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the DiscountGroup was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Name of the DiscountGroup.
	Name *LocalizedString `json:"name,omitempty"`
	// User-defined unique identifier of the DiscountGroup.
	Key string `json:"key"`
	// Description of the DiscountGroup.
	Description *LocalizedString `json:"description,omitempty"`
	// Value between `0` and `1` that determines the order in which the CartDiscount from the DiscountGroup is applied; a CartDiscount with a higher value is prioritized.
	//
	// The sort order is unique among all DiscountGroups and CartDiscounts.
	SortOrder string `json:"sortOrder"`
	// IDs and references that last modified the DiscountGroup.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the DiscountGroup.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
}

type DiscountGroupDraft struct {
	// Name of the DiscountGroup.
	Name *LocalizedString `json:"name,omitempty"`
	// User-defined unique identifier for the DiscountGroup.
	Key string `json:"key"`
	// Description for the DiscountGroup.
	Description *LocalizedString `json:"description,omitempty"`
	// Value between `0` and `1` that determines the order in which the CartDiscount from the DiscountGroup will be applied; a CartDiscount with a higher value will be prioritized.
	//
	// The sort order must be unique among all DiscountGroups and CartDiscounts.
	SortOrder string `json:"sortOrder"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [DiscountGroup](ctp:api:type:DiscountGroup).
*
 */
type DiscountGroupPagedQueryResponse struct {
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
	// DiscountGroups matching the query.
	Results []DiscountGroup `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [DiscountGroup](ctp:api:type:DiscountGroup).
*
 */
type DiscountGroupReference struct {
	// Unique identifier of the referenced [DiscountGroup](ctp:api:type:DiscountGroup).
	ID string `json:"id"`
	// Contains the representation of the expanded DiscountGroup.
	// Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for DiscountGroups.
	Obj *DiscountGroup `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountGroupReference) MarshalJSON() ([]byte, error) {
	type Alias DiscountGroupReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "discount-group", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [DiscountGroup](ctp:api:type:DiscountGroup). Either `id` or `key` is required.
*	If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type DiscountGroupResourceIdentifier struct {
	// Unique identifier of the referenced [DiscountGroup](ctp:api:type:DiscountGroup).
	// Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [DiscountGroup](ctp:api:type:DiscountGroup).
	// Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountGroupResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias DiscountGroupResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "discount-group", Alias: (*Alias)(&obj)})
}

type DiscountGroupUpdate struct {
	// Expected version of the DiscountGroup on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the DiscountGroup.
	Actions []DiscountGroupUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountGroupUpdate) UnmarshalJSON(data []byte) error {
	type Alias DiscountGroupUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorDiscountGroupUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type DiscountGroupUpdateAction interface{}

func mapDiscriminatorDiscountGroupUpdateAction(input interface{}) (DiscountGroupUpdateAction, error) {
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
	case "setDescription":
		obj := DiscountGroupSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := DiscountGroupSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setName":
		obj := DiscountGroupSetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSortOrder":
		obj := DiscountGroupSetSortOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type DiscountGroupSetDescriptionAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountGroupSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountGroupSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

/**
*	Setting a key generates the [DiscountGroupKeySet](ctp:api:type:DiscountGroupKeySetMessage) Message.
*
 */
type DiscountGroupSetKeyAction struct {
	// New value to set.
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountGroupSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountGroupSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type DiscountGroupSetNameAction struct {
	// New value to set.
	// If empty, any existing value will be removed.
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountGroupSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountGroupSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

/**
*	Setting the sort order generates the [DiscountGroupSortOrderSet](ctp:api:type:DiscountGroupSortOrderSetMessage) Message.
*
 */
type DiscountGroupSetSortOrderAction struct {
	// New value to set (between `0` and `1`).
	// A CartDiscount with a higher value will be prioritized.
	//
	// The sort order must be unique among all DiscountGroups and CartDiscounts.
	SortOrder string `json:"sortOrder"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DiscountGroupSetSortOrderAction) MarshalJSON() ([]byte, error) {
	type Alias DiscountGroupSetSortOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSortOrder", Alias: (*Alias)(&obj)})
}
