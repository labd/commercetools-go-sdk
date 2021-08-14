// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type Channel struct {
	// The unique ID of the channel.
	Id             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for events not tracked.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	// Present on resources created after 1/02/2019 except for events not tracked.
	CreatedBy *CreatedBy `json:"createdBy,omitEmpty"`
	// Any arbitrary string key that uniquely identifies this channel within the project.
	Key string `json:"key"`
	// The roles of this channel.
	// Each channel must have at least one role.
	Roles []ChannelRoleEnum `json:"roles"`
	// A human-readable name of the channel.
	Name *LocalizedString `json:"name,omitEmpty"`
	// A human-readable description of the channel.
	Description *LocalizedString `json:"description,omitEmpty"`
	// The address where this channel is located (e.g.
	// if the channel is a physical store).
	Address *Address `json:"address,omitEmpty"`
	// Statistics about the review ratings taken into account for this channel.
	ReviewRatingStatistics *ReviewRatingStatistics `json:"reviewRatingStatistics,omitEmpty"`
	Custom                 *CustomFields           `json:"custom,omitEmpty"`
	// A GeoJSON geometry object encoding the geo location of the channel.
	GeoLocation GeoJson `json:"geoLocation,omitEmpty"`
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
	Roles       []ChannelRoleEnum `json:"roles,omitEmpty"`
	Name        *LocalizedString  `json:"name,omitEmpty"`
	Description *LocalizedString  `json:"description,omitEmpty"`
	Address     *BaseAddress      `json:"address,omitEmpty"`
	// The custom fields.
	Custom      *CustomFieldsDraft `json:"custom,omitEmpty"`
	GeoLocation GeoJson            `json:"geoLocation,omitEmpty"`
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

type ChannelPagedQueryResponse struct {
	Limit   int       `json:"limit"`
	Count   int       `json:"count"`
	Total   int       `json:"total,omitEmpty"`
	Offset  int       `json:"offset"`
	Results []Channel `json:"results"`
}

type ChannelReference struct {
	Id  string   `json:"id"`
	Obj *Channel `json:"obj,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelReference) MarshalJSON() ([]byte, error) {
	type Alias ChannelReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "channel", Alias: (*Alias)(&obj)})
}

type ChannelResourceIdentifier struct {
	Id  string `json:"id,omitEmpty"`
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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
		new := ChannelAddRolesAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeDescription":
		new := ChannelChangeDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeKey":
		new := ChannelChangeKeyAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := ChannelChangeNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeRoles":
		new := ChannelRemoveRolesAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAddress":
		new := ChannelSetAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAddressCustomField":
		new := ChannelSetAddressCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAddressCustomType":
		new := ChannelSetAddressCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := ChannelSetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := ChannelSetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setGeoLocation":
		new := ChannelSetGeoLocationAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setRoles":
		new := ChannelSetRolesAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type ChannelAddRolesAction struct {
	Roles []ChannelRoleEnum `json:"roles"`
}

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
func (obj ChannelRemoveRolesAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelRemoveRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeRoles", Alias: (*Alias)(&obj)})
}

type ChannelSetAddressAction struct {
	Address *BaseAddress `json:"address,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelSetAddressAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAddress", Alias: (*Alias)(&obj)})
}

type ChannelSetAddressCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelSetAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAddressCustomField", Alias: (*Alias)(&obj)})
}

type ChannelSetAddressCustomTypeAction struct {
	Type   TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields *FieldContainer        `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelSetAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAddressCustomType", Alias: (*Alias)(&obj)})
}

type ChannelSetCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type ChannelSetCustomTypeAction struct {
	Type   TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields *FieldContainer        `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ChannelSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type ChannelSetGeoLocationAction struct {
	GeoLocation GeoJson `json:"geoLocation,omitEmpty"`
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
func (obj ChannelSetRolesAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRoles", Alias: (*Alias)(&obj)})
}
