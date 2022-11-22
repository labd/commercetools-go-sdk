package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type Channel struct {
	// Unique identifier of the Channel.
	ID string `json:"id"`
	// Current version of the Channel.
	Version int `json:"version"`
	// Date and time (UTC) the Channel was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Channel was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the Channel.
	Key string `json:"key"`
	// Roles of the Channel.
	Roles []ChannelRoleEnum `json:"roles"`
	// Name of the Channel.
	Name *LocalizedString `json:"name,omitempty"`
	// Description of the Channel.
	Description *LocalizedString `json:"description,omitempty"`
	// Address where the Channel is located (for example, if the Channel is a physical store).
	Address *Address `json:"address,omitempty"`
	// Statistics about the review ratings taken into account for the Channel.
	ReviewRatingStatistics *ReviewRatingStatistics `json:"reviewRatingStatistics,omitempty"`
	// Custom Fields defined for the Channel.
	Custom *CustomFields `json:"custom,omitempty"`
	// GeoJSON geometry object encoding the geo location of the Channel.
	GeoLocation GeoJson `json:"geoLocation,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Channel) UnmarshalJSON(data []byte) error {
	type Alias Channel
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.GeoLocation != nil {
		var err error
		obj.GeoLocation, err = mapDiscriminatorGeoJson(obj.GeoLocation)
		if err != nil {
			return err
		}
	}

	return nil
}

type ChannelDraft struct {
	// User-defined unique identifier for the Channel.
	Key string `json:"key"`
	// Roles of the Channel.
	// Each channel must have at least one role.
	// If not specified, then `InventorySupply` is assigned by default.
	Roles []ChannelRoleEnum `json:"roles"`
	// Name of the Channel.
	Name *LocalizedString `json:"name,omitempty"`
	// Description of the Channel.
	Description *LocalizedString `json:"description,omitempty"`
	// Address where the Channel is located.
	Address *BaseAddress `json:"address,omitempty"`
	// Custom fields defined for the Channel.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// GeoJSON geometry object encoding the geo location of the Channel.
	// Currently, only the [Point](ctp:api:type:GeoJsonPoint) type is supported.
	GeoLocation GeoJson `json:"geoLocation,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ChannelDraft) UnmarshalJSON(data []byte) error {
	type Alias ChannelDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.GeoLocation != nil {
		var err error
		obj.GeoLocation, err = mapDiscriminatorGeoJson(obj.GeoLocation)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelDraft) MarshalJSON() ([]byte, error) {
	type Alias ChannelDraft
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

	if raw["roles"] == nil {
		delete(raw, "roles")
	}

	return json.Marshal(raw)

}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [Channel](ctp:api:type:Channel).
*
 */
type ChannelPagedQueryResponse struct {
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
	// [Channels](ctp:api:type:Channel) matching the query.
	Results []Channel `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [Channel](ctp:api:type:Channel).
*
 */
type ChannelReference struct {
	// Unique identifier of the referenced [Channel](ctp:api:type:Channel).
	ID string `json:"id"`
	// Contains the representation of the expanded Channel.
	// Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for Channels.
	Obj *Channel `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelReference) MarshalJSON() ([]byte, error) {
	type Alias ChannelReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "channel", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Channel](ctp:api:type:Channel).
*
 */
type ChannelResourceIdentifier struct {
	// Unique identifier of the referenced [Channel](ctp:api:type:Channel). Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [Channel](ctp:api:type:Channel). Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ChannelResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "channel", Alias: (*Alias)(&obj)})
}

/**
*	Describes the purpose and type of the Channel. A Channel can have one or more roles.
*
 */
type ChannelRoleEnum string

const (
	ChannelRoleEnumInventorySupply     ChannelRoleEnum = "InventorySupply"
	ChannelRoleEnumProductDistribution ChannelRoleEnum = "ProductDistribution"
	ChannelRoleEnumOrderExport         ChannelRoleEnum = "OrderExport"
	ChannelRoleEnumOrderImport         ChannelRoleEnum = "OrderImport"
	ChannelRoleEnumPrimary             ChannelRoleEnum = "Primary"
)

type ChannelUpdate struct {
	// Expected version of the Channel on which the changes should be applied. If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error is returned.
	Version int `json:"version"`
	// Update actions to be performed on the Channel.
	Actions []ChannelUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ChannelUpdate) UnmarshalJSON(data []byte) error {
	type Alias ChannelUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorChannelUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type ChannelUpdateAction interface{}

func mapDiscriminatorChannelUpdateAction(input interface{}) (ChannelUpdateAction, error) {
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
	case "addRoles":
		obj := ChannelAddRolesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeDescription":
		obj := ChannelChangeDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeKey":
		obj := ChannelChangeKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := ChannelChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeRoles":
		obj := ChannelRemoveRolesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAddress":
		obj := ChannelSetAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAddressCustomField":
		obj := ChannelSetAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAddressCustomType":
		obj := ChannelSetAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := ChannelSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := ChannelSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setGeoLocation":
		obj := ChannelSetGeoLocationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.GeoLocation != nil {
			var err error
			obj.GeoLocation, err = mapDiscriminatorGeoJson(obj.GeoLocation)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "setRoles":
		obj := ChannelSetRolesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ChannelAddRolesAction struct {
	// Value to append to the array.
	Roles []ChannelRoleEnum `json:"roles"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelAddRolesAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelAddRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addRoles", Alias: (*Alias)(&obj)})
}

type ChannelChangeDescriptionAction struct {
	// New value to set. Must not be empty.
	Description LocalizedString `json:"description"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelChangeDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelChangeDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeDescription", Alias: (*Alias)(&obj)})
}

type ChannelChangeKeyAction struct {
	// New value to set. Must not be empty.
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelChangeKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelChangeKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeKey", Alias: (*Alias)(&obj)})
}

type ChannelChangeNameAction struct {
	// New value to set. Must not be empty.
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ChannelRemoveRolesAction struct {
	// Value to remove from the array.
	Roles []ChannelRoleEnum `json:"roles"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelRemoveRolesAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelRemoveRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeRoles", Alias: (*Alias)(&obj)})
}

type ChannelSetAddressAction struct {
	// Value to set. If empty, any existing value will be removed.
	Address *BaseAddress `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelSetAddressAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAddress", Alias: (*Alias)(&obj)})
}

type ChannelSetAddressCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// Specifies the format of the value of the Custom Field defined by `name`.
	// If `value` is absent or `null`, this field will be removed, if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelSetAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAddressCustomField", Alias: (*Alias)(&obj)})
}

type ChannelSetAddressCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the `address` with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the `address`.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `address`.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelSetAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAddressCustomType", Alias: (*Alias)(&obj)})
}

type ChannelSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type ChannelSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the Channel with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Channel.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Channel.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type ChannelSetGeoLocationAction struct {
	// Value to set.
	GeoLocation GeoJson `json:"geoLocation,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ChannelSetGeoLocationAction) UnmarshalJSON(data []byte) error {
	type Alias ChannelSetGeoLocationAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.GeoLocation != nil {
		var err error
		obj.GeoLocation, err = mapDiscriminatorGeoJson(obj.GeoLocation)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelSetGeoLocationAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetGeoLocationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setGeoLocation", Alias: (*Alias)(&obj)})
}

type ChannelSetRolesAction struct {
	// Value to set. If not specified, then `InventorySupply` is assigned by default.
	Roles []ChannelRoleEnum `json:"roles"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChannelSetRolesAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRoles", Alias: (*Alias)(&obj)})
}
