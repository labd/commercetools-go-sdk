package customize

import (
	"encoding/json"
	"time"

	"github.com/mitchellh/mapstructure"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

type Type struct {
	ID               string                        `json:"id"`
	Version          int                           `json:"version"`
	Key              string                        `json:"key"`
	Name             commercetools.LocalizedString `json:"name"`
	Description      commercetools.LocalizedString `json:"description,omitempty"`
	ResourceTypeIds  []string                      `json:"resourceTypeIds"`
	FieldDefinitions []FieldDefinition             `json:"fieldDefinitions"`
	CreatedAt        time.Time                     `json:"createdAt"`
	LastModifiedAt   time.Time                     `json:"lastModifiedAt"`
}

type TypeDraft struct {
	Key              string                        `json:"key"`
	Name             commercetools.LocalizedString `json:"name"`
	Description      commercetools.LocalizedString `json:"description,omitempty"`
	ResourceTypeIds  []string                      `json:"resourceTypeIds"`
	FieldDefinitions []FieldDefinition             `json:"fieldDefinitions"`
}

type FieldDefinition struct {
	Type      FieldType                     `json:"type"`
	Name      string                        `json:"name"`
	Label     commercetools.LocalizedString `json:"label"`
	Required  bool                          `json:"required"`
	InputHint TextInputHint                 `json:"inputHint"`
}

func (f *FieldDefinition) UnmarshalJSON(data []byte) error {
	type Alias FieldDefinition
	if err := json.Unmarshal(data, (*Alias)(f)); err != nil {
		return err
	}
	f.Type = fieldTypeMapping(f.Type)
	return nil
}

type BooleanType struct{}

func (t BooleanType) Name() string {
	return "Boolean"
}

// MarshalJSON override to add the Name() value
func (t BooleanType) MarshalJSON() ([]byte, error) {
	type Alias BooleanType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  t.Name(),
		Alias: (*Alias)(&t),
	})
}

type StringType struct{}

func (t StringType) Name() string {
	return "String"
}

// MarshalJSON override to add the Name() value
func (t StringType) MarshalJSON() ([]byte, error) {
	type Alias StringType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  t.Name(),
		Alias: (*Alias)(&t),
	})
}

type LocalizedStringType struct{}

func (t LocalizedStringType) Name() string {
	return "LocalizedString"
}

// MarshalJSON override to add the Name() value
func (t LocalizedStringType) MarshalJSON() ([]byte, error) {
	type Alias LocalizedStringType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  t.Name(),
		Alias: (*Alias)(&t),
	})
}

type EnumType struct {
	Values []EnumValue `json:"values"`
}

func (t EnumType) Name() string {
	return "Enum"
}

// MarshalJSON override to add the Name() value
func (t EnumType) MarshalJSON() ([]byte, error) {
	type Alias EnumType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  t.Name(),
		Alias: (*Alias)(&t),
	})
}

type LocalizedEnumType struct {
	Values []LocalizedEnumValue `json:"values"`
}

func (t LocalizedEnumType) Name() string {
	return "LocalizedEnum"
}

// MarshalJSON override to add the Name() value
func (t LocalizedEnumType) MarshalJSON() ([]byte, error) {
	type Alias LocalizedEnumType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  t.Name(),
		Alias: (*Alias)(&t),
	})
}

type NumberType struct{}

func (t NumberType) Name() string {
	return "Number"
}

// MarshalJSON override to add the Name() value
func (t NumberType) MarshalJSON() ([]byte, error) {
	type Alias NumberType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  t.Name(),
		Alias: (*Alias)(&t),
	})
}

type MoneyType struct{}

func (t MoneyType) Name() string {
	return "Money"
}

// MarshalJSON override to add the Name() value
func (t MoneyType) MarshalJSON() ([]byte, error) {
	type Alias MoneyType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  t.Name(),
		Alias: (*Alias)(&t),
	})
}

type DateType struct{}

func (t DateType) Name() string {
	return "Date"
}

// MarshalJSON override to add the Name() value
func (t DateType) MarshalJSON() ([]byte, error) {
	type Alias DateType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  t.Name(),
		Alias: (*Alias)(&t),
	})
}

type TimeType struct{}

func (t TimeType) Name() string {
	return "Time"
}

// MarshalJSON override to add the Name() value
func (t TimeType) MarshalJSON() ([]byte, error) {
	type Alias TimeType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  t.Name(),
		Alias: (*Alias)(&t),
	})
}

type DateTimeType struct{}

func (t DateTimeType) Name() string {
	return "DateTime"
}

// MarshalJSON override to add the Name() value
func (t DateTimeType) MarshalJSON() ([]byte, error) {
	type Alias DateTimeType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  t.Name(),
		Alias: (*Alias)(&t),
	})
}

type ReferenceType struct {
	ReferenceTypeId string `json:"referenceTypeId"`
}

func (t ReferenceType) Name() string {
	return "Reference"
}

// MarshalJSON override to add the Name() value
func (t ReferenceType) MarshalJSON() ([]byte, error) {
	type Alias ReferenceType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  t.Name(),
		Alias: (*Alias)(&t),
	})
}

type SetType struct {
	ElementType FieldType `json:"elementType"`
}

func (t SetType) Name() string {
	return "Set"
}

// MarshalJSON override to add the Name() value
func (t SetType) MarshalJSON() ([]byte, error) {
	type Alias SetType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  t.Name(),
		Alias: (*Alias)(&t),
	})
}

type FieldType interface{}

func fieldTypeMapping(input FieldType) FieldType {
	FieldType := input.(map[string]interface{})["name"]
	switch FieldType {
	case "Boolean":
		new := BooleanType{}
		mapstructure.Decode(input, &new)
		return new
	case "String":
		new := StringType{}
		mapstructure.Decode(input, &new)
		return new
	case "LocalizedString":
		new := LocalizedStringType{}
		mapstructure.Decode(input, &new)
		return new
	case "Enum":
		new := EnumType{}
		mapstructure.Decode(input, &new)
		return new
	case "LocalizedEnum":
		new := LocalizedEnumType{}
		mapstructure.Decode(input, &new)
		return new
	case "Number":
		new := NumberType{}
		mapstructure.Decode(input, &new)
		return new
	case "Money":
		new := MoneyType{}
		mapstructure.Decode(input, &new)
		return new
	case "Date":
		new := DateType{}
		mapstructure.Decode(input, &new)
		return new
	case "Time":
		new := TimeType{}
		mapstructure.Decode(input, &new)
		return new
	case "DateTime":
		new := DateTimeType{}
		mapstructure.Decode(input, &new)
		return new
	case "Reference":
		new := ReferenceType{}
		mapstructure.Decode(input, &new)
		return new
	case "Set":
		new := SetType{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

type EnumValue struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

type LocalizedEnumValue struct {
	Key   string                        `json:"key"`
	Label commercetools.LocalizedString `json:"label"`
}

type TextInputHint string

const (
	SingleLineTextInputHint TextInputHint = "SingleLine"
	MultiLineTextInputHint  TextInputHint = "MultiLine"
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
