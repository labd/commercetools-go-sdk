package channels

import (
	"encoding/json"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

// ChangeKey will change the key of the channel being updated.
type ChangeKey struct {
	Key string `json:"key"`
}

// MarshalJSON override to add the action value.
func (ua ChangeKey) MarshalJSON() ([]byte, error) {
	type Alias ChangeKey

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeKey",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeName will change the name of the channel being updated.
type ChangeName struct {
	Name commercetools.LocalizedString `json:"name"`
}

// MarshalJSON override to add the action value.
func (ua ChangeName) MarshalJSON() ([]byte, error) {
	type Alias ChangeName

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeName",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeDescription will change the description of the channel being updated.
type ChangeDescription struct {
	Description commercetools.LocalizedString `json:"description"`
}

// MarshalJSON override to add the action value.
func (ua ChangeDescription) MarshalJSON() ([]byte, error) {
	type Alias ChangeDescription

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeDescription",
		Alias:  (*Alias)(&ua),
	})
}

// SetRoles will set the roles of the channel being updated.
type SetRoles struct {
	Roles []ChannelRole `json:"roles"`
}

// MarshalJSON override to add the action value.
func (ua SetRoles) MarshalJSON() ([]byte, error) {
	type Alias SetRoles

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setRoles",
		Alias:  (*Alias)(&ua),
	})
}

// AddRoles will add roles to the channel being updated.
type AddRoles struct {
	Roles []ChannelRole `json:"roles"`
}

// MarshalJSON override to add the action value.
func (ua AddRoles) MarshalJSON() ([]byte, error) {
	type Alias AddRoles

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "addRoles",
		Alias:  (*Alias)(&ua),
	})
}

// RemoveRoles will remove roles from the channel being updated.
type RemoveRoles struct {
	Roles []ChannelRole `json:"roles"`
}

// MarshalJSON override to add the action value.
func (ua RemoveRoles) MarshalJSON() ([]byte, error) {
	type Alias RemoveRoles

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "removeRoles",
		Alias:  (*Alias)(&ua),
	})
}

// SetAddress will set the address of the channel being updated.
type SetAddress struct {
	Address commercetools.Address `json:"address,omitempty"`
}

// MarshalJSON override to add the action value.
func (ua SetAddress) MarshalJSON() ([]byte, error) {
	type Alias SetAddress

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setAddress",
		Alias:  (*Alias)(&ua),
	})
}

// SetCustomType will set the custom types of the channel being updated.
type SetCustomType struct {
	// If absent, the custom type and any existing custom fields are removed.
	Type   commercetools.ResourceIdentifier `json:"type,omitempty"`
	Fields map[string]string                `json:"fields,omitempty"`
}

// MarshalJSON override to add the action value.
func (ua SetCustomType) MarshalJSON() ([]byte, error) {
	type Alias SetCustomType

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setCustomType",
		Alias:  (*Alias)(&ua),
	})
}

// SetCustomField will set the custom fields of the channel being updated.
type SetCustomField struct {
	Name string `json:"name"`

	// If value is absent or null, this field will be removed if it exists. If
	// value is provided, set the  value of the field defined by the name. The
	// FieldDefinition determines the format for the value to be provided. In
	// particular, for the fields definitions, value has to be provided as:
	// - for EnumType and LocalizedEnumType fields: the key of the EnumValue or
	//   the LocalizedEnumValue
	// - for LocalizedStringType fields: the LocalizedString object
	// - for MoneyType fields: the Money object
	// - for SetType fields: the entire Set object
	// - for ReferenceType fields: the Reference object
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to add the action value.
func (ua SetCustomField) MarshalJSON() ([]byte, error) {
	type Alias SetCustomField

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setCustomField",
		Alias:  (*Alias)(&ua),
	})
}

// SetGeoLocation will set the geo location of the channel being updated.
type SetGeoLocation struct {
	GeoLocation commercetools.GeoJSONGeometry `json:"geoLocation,omitempty"`
}

// MarshalJSON override to add the action value.
func (ua SetGeoLocation) MarshalJSON() ([]byte, error) {
	type Alias SetGeoLocation

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setGeoLocation",
		Alias:  (*Alias)(&ua),
	})
}
