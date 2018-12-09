// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type Channel struct {
	Version                int                     `json:"version"`
	LastModifiedAt         time.Time               `json:"lastModifiedAt"`
	ID                     string                  `json:"id"`
	CreatedAt              time.Time               `json:"createdAt"`
	Roles                  []ChannelRoleEnum       `json:"roles"`
	ReviewRatingStatistics *ReviewRatingStatistics `json:"reviewRatingStatistics,omitempty"`
	Name                   *LocalizedString        `json:"name,omitempty"`
	Key                    string                  `json:"key"`
	GeoLocation            *GeoJsonPoint           `json:"geoLocation,omitempty"`
	Description            *LocalizedString        `json:"description,omitempty"`
	Custom                 *CustomFields           `json:"custom,omitempty"`
	Address                *Address                `json:"address,omitempty"`
}

type ChannelAddRolesAction struct {
	Roles []ChannelRoleEnum `json:"roles"`
}

func (obj ChannelAddRolesAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelAddRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addRoles", Alias: (*Alias)(&obj)})
}

type ChannelChangeDescriptionAction struct {
	Description *LocalizedString `json:"description"`
}

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

func (obj ChannelChangeKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelChangeKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeKey", Alias: (*Alias)(&obj)})
}

type ChannelChangeNameAction struct {
	Name *LocalizedString `json:"name"`
}

func (obj ChannelChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ChannelDraft struct {
	Roles       []ChannelRoleEnum  `json:"roles,omitempty"`
	Name        *LocalizedString   `json:"name,omitempty"`
	Key         string             `json:"key"`
	GeoLocation *GeoJsonPoint      `json:"geoLocation,omitempty"`
	Description *LocalizedString   `json:"description,omitempty"`
	Custom      *CustomFieldsDraft `json:"custom,omitempty"`
	Address     *Address           `json:"address,omitempty"`
}

type ChannelPagedQueryResponse struct {
	Total   int       `json:"total,omitempty"`
	Offset  int       `json:"offset"`
	Count   int       `json:"count"`
	Results []Channel `json:"results"`
}

type ChannelReference struct {
	Key string   `json:"key,omitempty"`
	ID  string   `json:"id,omitempty"`
	Obj *Channel `json:"obj,omitempty"`
}

func (obj ChannelReference) MarshalJSON() ([]byte, error) {
	type Alias ChannelReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "channel", Alias: (*Alias)(&obj)})
}

type ChannelRemoveRolesAction struct {
	Roles []ChannelRoleEnum `json:"roles"`
}

func (obj ChannelRemoveRolesAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelRemoveRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeRoles", Alias: (*Alias)(&obj)})
}

type ChannelRoleEnum string

const (
	ChannelRoleEnumInventorySupply     ChannelRoleEnum = "InventorySupply"
	ChannelRoleEnumProductDistribution ChannelRoleEnum = "ProductDistribution"
	ChannelRoleEnumOrderExport         ChannelRoleEnum = "OrderExport"
	ChannelRoleEnumOrderImport         ChannelRoleEnum = "OrderImport"
	ChannelRoleEnumPrimary             ChannelRoleEnum = "Primary"
)

type ChannelSetAddressAction struct {
	Address *Address `json:"address,omitempty"`
}

func (obj ChannelSetAddressAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAddress", Alias: (*Alias)(&obj)})
}

type ChannelSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

func (obj ChannelSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type ChannelSetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

func (obj ChannelSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type ChannelSetGeoLocationAction struct {
	GeoLocation *GeoJsonPoint `json:"geoLocation,omitempty"`
}

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

func (obj ChannelSetRolesAction) MarshalJSON() ([]byte, error) {
	type Alias ChannelSetRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRoles", Alias: (*Alias)(&obj)})
}

type ChannelUpdate struct {
	Version int                   `json:"version"`
	Actions []ChannelUpdateAction `json:"actions"`
}

func (obj *ChannelUpdate) UnmarshalJSON(data []byte) error {
	type Alias ChannelUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractChannelUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type ChannelUpdateAction interface{}
type AbstractChannelUpdateAction struct{}

func AbstractChannelUpdateActionDiscriminatorMapping(input ChannelUpdateAction) ChannelUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addRoles":
		new := ChannelAddRolesAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeDescription":
		new := ChannelChangeDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeKey":
		new := ChannelChangeKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeName":
		new := ChannelChangeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeRoles":
		new := ChannelRemoveRolesAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAddress":
		new := ChannelSetAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomField":
		new := ChannelSetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomType":
		new := ChannelSetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setGeoLocation":
		new := ChannelSetGeoLocationAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setRoles":
		new := ChannelSetRolesAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}
