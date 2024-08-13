package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type AttributeGroup struct {
	// Platform-generated unique identifier of the AttributeGroup.
	ID string `json:"id"`
	// Current version of the AttributeGroup.
	Version int `json:"version"`
	// Date and time (UTC) the AttributeGroup was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the AttributeGroup was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// IDs and references that last modified the AttributeGroup.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the AttributeGroup.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Name of the AttributeGroup.
	Name LocalizedString `json:"name"`
	// Description of the AttributeGroup.
	Description *LocalizedString `json:"description,omitempty"`
	// [Attributes](ctp:api:type:AttributeDefinition) with unique values.
	Attributes []AttributeReference `json:"attributes"`
	// User-defined unique identifier of the AttributeGroup.
	Key *string `json:"key,omitempty"`
}

type AttributeGroupDraft struct {
	// Name of the AttributeGroup.
	Name LocalizedString `json:"name"`
	// Description of the AttributeGroup.
	Description *LocalizedString `json:"description,omitempty"`
	// [Attributes](ctp:api:type:AttributeDefinition) with unique values.
	Attributes []AttributeReference `json:"attributes"`
	// User-defined unique identifier for the AttributeGroup.
	Key *string `json:"key,omitempty"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [AttributeGroup](ctp:api:type:AttributeGroup).
*
 */
type AttributeGroupPagedQueryResponse struct {
	// Number of results requested in the query request.
	Limit int `json:"limit"`
	// Offset supplied by the client or the server default.
	// It is the number of elements skipped, not a page number.
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [AttributeGroups](ctp:api:type:AttributeGroup) matching the query.
	Results []AttributeGroup `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to an [AttributeGroup](ctp:api:type:AttributeGroup).
*
 */
type AttributeGroupReference struct {
	// Platform-generated unique identifier of the referenced [AttributeGroup](ctp:api:type:AttributeGroup).
	ID string `json:"id"`
	// Contains the representation of the expanded AttributeGroup. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for AttributeGroup.
	Obj *AttributeGroup `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeGroupReference) MarshalJSON() ([]byte, error) {
	type Alias AttributeGroupReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "attribute-group", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to an [AttributeGroup](ctp:api:type:AttributeGroup). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type AttributeGroupResourceIdentifier struct {
	// Platform-generated unique identifier of the referenced [AttributeGroup](ctp:api:type:AttributeGroup). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-generated unique identifier of the referenced [AttributeGroup](ctp:api:type:AttributeGroup). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeGroupResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias AttributeGroupResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "attribute-group", Alias: (*Alias)(&obj)})
}

type AttributeGroupUpdate struct {
	// Expected version of the AttributeGroup on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the AttributeGroup.
	Actions []AttributeGroupUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AttributeGroupUpdate) UnmarshalJSON(data []byte) error {
	type Alias AttributeGroupUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorAttributeGroupUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type AttributeGroupUpdateAction interface{}

func mapDiscriminatorAttributeGroupUpdateAction(input interface{}) (AttributeGroupUpdateAction, error) {
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
	case "addAttribute":
		obj := AttributeGroupAddAttributeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := AttributeGroupChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeAttribute":
		obj := AttributeGroupRemoveAttributeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAttributes":
		obj := AttributeGroupSetAttributesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := AttributeGroupSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := AttributeGroupSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type AttributeReference struct {
	// The Attribute's `name` as given in its [AttributeDefinition](ctp:api:type:AttributeDefinition).
	Key string `json:"key"`
}

type AttributeGroupAddAttributeAction struct {
	// Value to add.
	Attribute AttributeReference `json:"attribute"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeGroupAddAttributeAction) MarshalJSON() ([]byte, error) {
	type Alias AttributeGroupAddAttributeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAttribute", Alias: (*Alias)(&obj)})
}

type AttributeGroupChangeNameAction struct {
	// New value to set.
	// Must not be empty.
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeGroupChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias AttributeGroupChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type AttributeGroupRemoveAttributeAction struct {
	// Value to remove.
	Attribute AttributeReference `json:"attribute"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeGroupRemoveAttributeAction) MarshalJSON() ([]byte, error) {
	type Alias AttributeGroupRemoveAttributeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAttribute", Alias: (*Alias)(&obj)})
}

type AttributeGroupSetAttributesAction struct {
	// New unique values to set.
	Attributes []AttributeReference `json:"attributes"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeGroupSetAttributesAction) MarshalJSON() ([]byte, error) {
	type Alias AttributeGroupSetAttributesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAttributes", Alias: (*Alias)(&obj)})
}

type AttributeGroupSetDescriptionAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeGroupSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias AttributeGroupSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type AttributeGroupSetKeyAction struct {
	// If `key` is absent or `null`, the existing key, if any, will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeGroupSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias AttributeGroupSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}
