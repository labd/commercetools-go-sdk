// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type Channel struct {
	// The unique ID of the channel.
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Any arbitrary string key that uniquely identifies this channel within the project.
	Key string `json:"key"`
	// The roles of this channel.
	// Each channel must have at least one role.
	Roles []ChannelRoleEnum `json:"roles"`
	// A human-readable name of the channel.
	Name *LocalizedString `json:"name,omitempty"`
	// A human-readable description of the channel.
	Description *LocalizedString `json:"description,omitempty"`
	// The address where this channel is located (e.g.
	// if the channel is a physical store).
	Address *Address `json:"address,omitempty"`
	// Statistics about the review ratings taken into account for this channel.
	ReviewRatingStatistics *ReviewRatingStatistics `json:"reviewRatingStatistics,omitempty"`
	Custom                 *CustomFields           `json:"custom,omitempty"`
	// A GeoJSON geometry object encoding the geo location of the channel.
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
	Key string `json:"key"`
	// If not specified, then channel will get InventorySupply role by default
	Roles       []ChannelRoleEnum `json:"roles"`
	Name        *LocalizedString  `json:"name,omitempty"`
	Description *LocalizedString  `json:"description,omitempty"`
	Address     *BaseAddress      `json:"address,omitempty"`
	// The custom fields.
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
	GeoLocation GeoJson            `json:"geoLocation,omitempty"`
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["roles"] == nil {
		delete(target, "roles")
	}

	return json.Marshal(target)
}

type ChannelPagedQueryResponse struct {
	Limit   int       `json:"limit"`
	Count   int       `json:"count"`
	Total   *int      `json:"total,omitempty"`
	Offset  int       `json:"offset"`
	Results []Channel `json:"results"`
}

type ChannelReference struct {
	// Unique ID of the referenced resource.
	ID  string   `json:"id"`
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

type ChannelResourceIdentifier struct {
	// Unique ID of the referenced resource. Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced resource. Either `id` or `key` is required.
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

type ChannelRoleEnum string

const (
	ChannelRoleEnumInventorySupply     ChannelRoleEnum = "InventorySupply"
	ChannelRoleEnumProductDistribution ChannelRoleEnum = "ProductDistribution"
	ChannelRoleEnumOrderExport         ChannelRoleEnum = "OrderExport"
	ChannelRoleEnumOrderImport         ChannelRoleEnum = "OrderImport"
	ChannelRoleEnumPrimary             ChannelRoleEnum = "Primary"
)

type ChannelUpdate struct {
	Version int                   `json:"version"`
	Actions []ChannelUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ChannelUpdate) UnmarshalJSON(data []byte) error {
	type Alias ChannelUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type ChannelUpdateAction interface{}

func mapDiscriminatorChannelUpdateAction(input interface{}) (ChannelUpdateAction, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("Invalid data")
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
	Name  string      `json:"name"`
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
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
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
	Name  string      `json:"name"`
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
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
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
