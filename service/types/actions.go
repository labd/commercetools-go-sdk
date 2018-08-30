package types

import (
	"encoding/json"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

// ChangeKey will change the key of the type being updated.
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

// ChangeName will change the name of the type being updated.
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

// SetDescription will set the description of the type being updated. If left
// blank, the description is removed.
type SetDescription struct {
	Description commercetools.LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to add the action value.
func (ua SetDescription) MarshalJSON() ([]byte, error) {
	type Alias SetDescription

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setDescription",
		Alias:  (*Alias)(&ua),
	})
}

// AddFieldDefinition will add a field definition to the type being updated.
type AddFieldDefinition struct {
	FieldDefinition FieldDefinition `json:"fieldDefinition"`
}

// MarshalJSON override to add the action value.
func (ua AddFieldDefinition) MarshalJSON() ([]byte, error) {
	type Alias AddFieldDefinition

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "addFieldDefinition",
		Alias:  (*Alias)(&ua),
	})
}

// RemoveFieldDefinition will remove a field definition to the type being
// updated. The name of the field to remove. The removal of a field definition
// deletes asynchronously all custom fields using this definition as well.
type RemoveFieldDefinition struct {
	FieldName string `json:"fieldName"`
}

// MarshalJSON override to add the action value.
func (ua RemoveFieldDefinition) MarshalJSON() ([]byte, error) {
	type Alias RemoveFieldDefinition

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "removeFieldDefinition",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeLabel will change the label of a field in the type being updated.
type ChangeLabel struct {
	FieldName string                        `json:"fieldName"`
	Label     commercetools.LocalizedString `json:"label"`
}

// MarshalJSON override to add the action value.
func (ua ChangeLabel) MarshalJSON() ([]byte, error) {
	type Alias ChangeLabel

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeLabel",
		Alias:  (*Alias)(&ua),
	})
}

// AddEnumValue will add an enum value to the type being updated. Adds an enum
// to the values of EnumType. It can update an EnumType field definition or a
// Set of EnumType field definition.
type AddEnumValue struct {
	FieldName string                  `json:"fieldName"`
	Value     commercetools.EnumValue `json:"value"`
}

// MarshalJSON override to add the action value.
func (ua AddEnumValue) MarshalJSON() ([]byte, error) {
	type Alias AddEnumValue

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "addEnumValue",
		Alias:  (*Alias)(&ua),
	})
}

// AddLocalizedEnumValue will add a localized enum value to the type being
// updated. Adds an localized enum to the values of LocalizedEnumType. It can
// update a LocalizedEnumType field definition or a Set of LocalizedEnumType
// field definition.
type AddLocalizedEnumValue struct {
	FieldName string                           `json:"fieldName"`
	Value     commercetools.LocalizedEnumValue `json:"value"`
}

// MarshalJSON override to add the action value.
func (ua AddLocalizedEnumValue) MarshalJSON() ([]byte, error) {
	type Alias AddLocalizedEnumValue

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "addLocalizedEnumValue",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeFieldDefinitionsOrder will reorder the fields on the type being
// updated.
type ChangeFieldDefinitionsOrder struct {
	FieldNames []string `json:"fieldNames"`
}

// MarshalJSON override to add the action value.
func (ua ChangeFieldDefinitionsOrder) MarshalJSON() ([]byte, error) {
	type Alias ChangeFieldDefinitionsOrder

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeFieldDefinitionOrder",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeEnumValuesOrder will reorder the enum values on the type being updated.
// This action changes the order of enum values in an EnumType field definition.
// It can update an EnumType field definition or a Set of EnumType field
// definition.
type ChangeEnumValuesOrder struct {
	FieldName string   `json:"fieldName"`
	Keys      []string `json:"keys"`
}

// MarshalJSON override to add the action value.
func (ua ChangeEnumValuesOrder) MarshalJSON() ([]byte, error) {
	type Alias ChangeEnumValuesOrder

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeEnumValueOrder",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeLocalizedEnumValuesOrder will reorder the localized enum values on the
// type being updated. This action changes the order of localized enum values in
// a LocalizedEnumType field definition. It can update a LocalizedEnumType field
// definition or a Set of LocalizedEnumType field definition.
type ChangeLocalizedEnumValuesOrder struct {
	FieldName string   `json:"fieldName"`
	Keys      []string `json:"keys"`
}

// MarshalJSON override to add the action value.
func (ua ChangeLocalizedEnumValuesOrder) MarshalJSON() ([]byte, error) {
	type Alias ChangeLocalizedEnumValuesOrder

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeLocalizedEnumValueOrder",
		Alias:  (*Alias)(&ua),
	})
}
