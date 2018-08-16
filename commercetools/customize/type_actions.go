package customize

import (
	"encoding/json"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

// TypeChangeKey will change the key of the type being updated.
type TypeChangeKey struct {
	Key string `json:"key"`
}

// Action override to provide the expected name for this update action.
func (ua TypeChangeKey) Action() string {
	return "changeKey"
}

// MarshalJSON override to add the Action() value.
func (ua TypeChangeKey) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeKey

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: ua.Action(),
		Alias:  (*Alias)(&ua),
	})
}

// TypeChangeName will change the name of the type being updated.
type TypeChangeName struct {
	Name commercetools.LocalizedString `json:"name"`
}

// Action override to provide the expected name for this update action.
func (ua TypeChangeName) Action() string {
	return "changeName"
}

// MarshalJSON override to add the Action() value.
func (ua TypeChangeName) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeName

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: ua.Action(),
		Alias:  (*Alias)(&ua),
	})
}

// TypeSetDescription will set the description of the type being updated.
// If left blank, the description is removed.
type TypeSetDescription struct {
	Description commercetools.LocalizedString `json:"description,omitempty"`
}

// Action override to provide the expected name for this update action.
func (ua TypeSetDescription) Action() string {
	return "setDescription"
}

// MarshalJSON override to add the Action() value.
func (ua TypeSetDescription) MarshalJSON() ([]byte, error) {
	type Alias TypeSetDescription

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: ua.Action(),
		Alias:  (*Alias)(&ua),
	})
}

// TypeAddFieldDefinition will add a field definition to the type being updated.
type TypeAddFieldDefinition struct {
	FieldDefinition FieldDefinition `json:"fieldDefinition"`
}

// Action override to provide the expected name for this update action.
func (ua TypeAddFieldDefinition) Action() string {
	return "addFieldDefinition"
}

// MarshalJSON override to add the Action() value.
func (ua TypeAddFieldDefinition) MarshalJSON() ([]byte, error) {
	type Alias TypeAddFieldDefinition

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: ua.Action(),
		Alias:  (*Alias)(&ua),
	})
}

// TypeRemoveFieldDefinition will remove a field definition to the type being updated.
// The name of the field to remove.
// The removal of a field definition deletes asynchronously all custom fields using this definition as well.
type TypeRemoveFieldDefinition struct {
	FieldName string `json:"fieldName"`
}

// Action override to provide the expected name for this update action.
func (ua TypeRemoveFieldDefinition) Action() string {
	return "removeFieldDefinition"
}

// MarshalJSON override to add the Action() value.
func (ua TypeRemoveFieldDefinition) MarshalJSON() ([]byte, error) {
	type Alias TypeRemoveFieldDefinition

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: ua.Action(),
		Alias:  (*Alias)(&ua),
	})
}

// TypeChangeLabel will change the label of a field in the type being updated.
type TypeChangeLabel struct {
	FieldName string                        `json:"fieldName"`
	Label     commercetools.LocalizedString `json:"label"`
}

// Action override to provide the expected name for this update action.
func (ua TypeChangeLabel) Action() string {
	return "changeLabel"
}

// MarshalJSON override to add the Action() value.
func (ua TypeChangeLabel) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeLabel

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: ua.Action(),
		Alias:  (*Alias)(&ua),
	})
}

// TypeAddEnumValue will add an enum value to the type being updated.
// Adds an enum to the values of EnumType.
// It can update an EnumType field definition or a Set of EnumType field definition.
type TypeAddEnumValue struct {
	FieldName string    `json:"fieldName"`
	Value     EnumValue `json:"value"`
}

// Action override to provide the expected name for this update action.
func (ua TypeAddEnumValue) Action() string {
	return "addEnumValue"
}

// MarshalJSON override to add the Action() value.
func (ua TypeAddEnumValue) MarshalJSON() ([]byte, error) {
	type Alias TypeAddEnumValue

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: ua.Action(),
		Alias:  (*Alias)(&ua),
	})
}

// TypeAddLocalizedEnumValue will add a localized enum value to the type being updated.
// Adds an localized enum to the values of LocalizedEnumType.
// It can update a LocalizedEnumType field definition or a Set of LocalizedEnumType field definition.
type TypeAddLocalizedEnumValue struct {
	FieldName string             `json:"fieldName"`
	Value     LocalizedEnumValue `json:"value"`
}

// Action override to provide the expected name for this update action.
func (ua TypeAddLocalizedEnumValue) Action() string {
	return "addLocalizedEnumValue"
}

// MarshalJSON override to add the Action() value.
func (ua TypeAddLocalizedEnumValue) MarshalJSON() ([]byte, error) {
	type Alias TypeAddLocalizedEnumValue

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: ua.Action(),
		Alias:  (*Alias)(&ua),
	})
}

// TypeChangeFieldDefinitionsOrder will reorder the fields on the type being updated.
type TypeChangeFieldDefinitionsOrder struct {
	FieldNames []string `json:"fieldNames"`
}

// Action override to provide the expected name for this update action.
func (ua TypeChangeFieldDefinitionsOrder) Action() string {
	return "changeFieldDefinitionOrder"
}

// MarshalJSON override to add the Action() value.
func (ua TypeChangeFieldDefinitionsOrder) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeFieldDefinitionsOrder

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: ua.Action(),
		Alias:  (*Alias)(&ua),
	})
}

// TypeChangeEnumValuesOrder will reorder the enum values on the type being updated.
// This action changes the order of enum values in an EnumType field definition.
// It can update an EnumType field definition or a Set of EnumType field definition.
type TypeChangeEnumValuesOrder struct {
	FieldName string   `json:"fieldName"`
	Keys      []string `json:"keys"`
}

// Action override to provide the expected name for this update action.
func (ua TypeChangeEnumValuesOrder) Action() string {
	return "changeEnumValueOrder"
}

// MarshalJSON override to add the Action() value.
func (ua TypeChangeEnumValuesOrder) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeEnumValuesOrder

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: ua.Action(),
		Alias:  (*Alias)(&ua),
	})
}

// TypeChangeLocalizedEnumValuesOrder will reorder the localized enum values on the type being updated.
// This action changes the order of localized enum values in a LocalizedEnumType field definition.
// It can update a LocalizedEnumType field definition or a Set of LocalizedEnumType field definition.
type TypeChangeLocalizedEnumValuesOrder struct {
	FieldName string   `json:"fieldName"`
	Keys      []string `json:"keys"`
}

// Action override to provide the expected name for this update action.
func (ua TypeChangeLocalizedEnumValuesOrder) Action() string {
	return "changeLocalizedEnumValueOrder"
}

// MarshalJSON override to add the Action() value.
func (ua TypeChangeLocalizedEnumValuesOrder) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeLocalizedEnumValuesOrder

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: ua.Action(),
		Alias:  (*Alias)(&ua),
	})
}
