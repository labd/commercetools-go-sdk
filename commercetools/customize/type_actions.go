package customize

import (
	"encoding/json"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

type TypeChangeKey struct {
	Key string `json:"key"`
}

func (ua TypeChangeKey) Action() string {
	return "changeKey"
}

// MarshalJSON override to add the Action() value
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

type TypeChangeName struct {
	Name commercetools.LocalizedString `json:"name"`
}

func (ua TypeChangeName) Action() string {
	return "changeName"
}

// MarshalJSON override to add the Action() value
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

type TypeSetDescription struct {
	Description commercetools.LocalizedString `json:"description,omitempty"`
}

func (ua TypeSetDescription) Action() string {
	return "setDescription"
}

// MarshalJSON override to add the Action() value
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

type TypeAddFieldDefinition struct {
	FieldDefinition FieldDefinition `json:"fieldDefinition"`
}

func (ua TypeAddFieldDefinition) Action() string {
	return "addFieldDefinition"
}

// MarshalJSON override to add the Action() value
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

type TypeRemoveFieldDefinition struct {
	FieldName string `json:"fieldName"`
}

func (ua TypeRemoveFieldDefinition) Action() string {
	return "removeFieldDefinition"
}

// MarshalJSON override to add the Action() value
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

type TypeChangeLabel struct {
	FieldName string                        `json:"fieldName"`
	Label     commercetools.LocalizedString `json:"label"`
}

func (ua TypeChangeLabel) Action() string {
	return "changeLabel"
}

// MarshalJSON override to add the Action() value
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

type TypeAddEnumValue struct {
	FieldName string    `json:"fieldName"`
	Value     EnumValue `json:"value"`
}

func (ua TypeAddEnumValue) Action() string {
	return "addEnumValue"
}

// MarshalJSON override to add the Action() value
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

type TypeAddLocalizedEnumValue struct {
	FieldName string             `json:"fieldName"`
	Value     LocalizedEnumValue `json:"value"`
}

func (ua TypeAddLocalizedEnumValue) Action() string {
	return "addLocalizedEnumValue"
}

// MarshalJSON override to add the Action() value
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

type TypeChangeFieldDefinitionsOrder struct {
	FieldNames []string `json:"fieldNames"`
}

func (ua TypeChangeFieldDefinitionsOrder) Action() string {
	return "changeFieldDefinitionOrder"
}

// MarshalJSON override to add the Action() value
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

type TypeChangeEnumValuesOrder struct {
	FieldName string   `json:"fieldName"`
	Keys      []string `json:"keys"`
}

func (ua TypeChangeEnumValuesOrder) Action() string {
	return "changeEnumValueOrder"
}

// MarshalJSON override to add the Action() value
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

type TypeChangeLocalizedEnumValuesOrder struct {
	FieldName string   `json:"fieldName"`
	Keys      []string `json:"keys"`
}

func (ua TypeChangeLocalizedEnumValuesOrder) Action() string {
	return "changeLocalizedEnumValueOrder"
}

// MarshalJSON override to add the Action() value
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
