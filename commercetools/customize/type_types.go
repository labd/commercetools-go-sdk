package customize

import (
	"encoding/json"
	"time"

	"github.com/mitchellh/mapstructure"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

// TextInputHint provides a visual representation type for a field.
// It is only relevant for string-based field types like StringType and LocalizedStringType.
type TextInputHint string

const (
	SingleLineTextInputHint TextInputHint = "SingleLine"
	MultiLineTextInputHint  TextInputHint = "MultiLine"
)

// Types define custom fields that are used to enhance resources as you need.
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

// TypeDrafts are given as payload for Create Type requests.
type TypeDraft struct {
	Key              string                        `json:"key"`
	Name             commercetools.LocalizedString `json:"name"`
	Description      commercetools.LocalizedString `json:"description,omitempty"`
	ResourceTypeIds  []string                      `json:"resourceTypeIds"`
	FieldDefinitions []FieldDefinition             `json:"fieldDefinitions"`
}

// Field definitions describe custom fields and allow you to define some meta-information associated with the field.
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

// Valid values for the field are true and false (JSON Boolean).
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

// The set field type defines a set (array without duplicates) with values of the given elementType.
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
		newType := BooleanType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "String":
		newType := StringType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "LocalizedString":
		newType := LocalizedStringType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "Enum":
		newType := EnumType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "LocalizedEnum":
		newType := LocalizedEnumType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "Number":
		newType := NumberType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "Money":
		newType := MoneyType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "Date":
		newType := DateType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "Time":
		newType := TimeType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "DateTime":
		newType := DateTimeType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "Reference":
		newType := ReferenceType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "Set":
		newType := SetType{}
		mapstructure.Decode(input, &newType)
		return newType
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
