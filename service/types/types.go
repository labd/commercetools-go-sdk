package types

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
	// SingleLineTextInputHint allows a single line hint.
	SingleLineTextInputHint TextInputHint = "SingleLine"
	// MultiLineTextInputHint allows a multi line hint.
	MultiLineTextInputHint TextInputHint = "MultiLine"
)

// Type defines custom fields that are used to enhance resources as you need.
type Type struct {
	// The unique ID of the type.
	ID string `json:"id"`
	// The current version of the type.
	Version int `json:"version"`
	// Identifier for the type (max. 256 characters).
	Key         string                        `json:"key"`
	Name        commercetools.LocalizedString `json:"name"`
	Description commercetools.LocalizedString `json:"description,omitempty"`
	// Defines for which resource(s) the type is valid.
	ResourceTypeIds  []string          `json:"resourceTypeIds"`
	FieldDefinitions []FieldDefinition `json:"fieldDefinitions"`
	CreatedAt        time.Time         `json:"createdAt"`
	LastModifiedAt   time.Time         `json:"lastModifiedAt"`
}

// TypeDraft is given as payload for Create Type requests.
type TypeDraft struct {
	// Identifier for the type (max. 256 characters).
	Key         string                        `json:"key"`
	Name        commercetools.LocalizedString `json:"name"`
	Description commercetools.LocalizedString `json:"description,omitempty"`
	// The IDs of the resources that can be customized with this type.
	ResourceTypeIds  []string          `json:"resourceTypeIds"`
	FieldDefinitions []FieldDefinition `json:"fieldDefinitions,omitempty"`
}

// FieldDefinition describe custom fields and allow you to define some meta-information associated with the field.
type FieldDefinition struct {
	// Describes the type of the field.
	Type FieldType `json:"type"`
	// The name of the field.
	// The name must be between two and 36 characters long and can contain the ASCII letters A to Z
	// in lowercase or uppercase, digits, underscores (_) and the hyphen-minus (-).
	// The name must be unique for a given resource type ID. In case there is a field with the same
	// name in another type it has to have the same FieldType also.
	Name string `json:"name"`
	// A human-readable label for the field.
	Label commercetools.LocalizedString `json:"label"`
	//  Whether the field is required to have a value.
	Required bool `json:"required"`
	// Provides a visual representation type for this field. It is only relevant for string-based
	// field types like StringType and LocalizedStringType.
	InputHint TextInputHint `json:"inputHint"`
}

// UnmarshalJSON override to map the field type to the corresponding struct.
func (f *FieldDefinition) UnmarshalJSON(data []byte) error {
	type Alias FieldDefinition
	if err := json.Unmarshal(data, (*Alias)(f)); err != nil {
		return err
	}
	f.Type = fieldTypeMapping(f.Type)
	return nil
}

// BooleanType will allow the storage of a boolan value.
// Valid values for the field are true and false (JSON Boolean).
type BooleanType struct{}

// MarshalJSON override to add the name value.
func (t BooleanType) MarshalJSON() ([]byte, error) {
	type Alias BooleanType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "Boolean",
		Alias: (*Alias)(&t),
	})
}

// StringType will allow the storage of a string value.
type StringType struct{}

// MarshalJSON override to add the name value.
func (t StringType) MarshalJSON() ([]byte, error) {
	type Alias StringType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "String",
		Alias: (*Alias)(&t),
	})
}

// LocalizedStringType will allow the storage of a localized string value.
type LocalizedStringType struct{}

// MarshalJSON override to add the name value.
func (t LocalizedStringType) MarshalJSON() ([]byte, error) {
	type Alias LocalizedStringType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "LocalizedString",
		Alias: (*Alias)(&t),
	})
}

// EnumType will allow the storage of an enum value.
type EnumType struct {
	Values []commercetools.EnumValue `json:"values"`
}

// MarshalJSON override to add the name value.
func (t EnumType) MarshalJSON() ([]byte, error) {
	type Alias EnumType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "Enum",
		Alias: (*Alias)(&t),
	})
}

// LocalizedEnumType will allow the storage of a localized enum value.
type LocalizedEnumType struct {
	Values []commercetools.LocalizedEnumValue `json:"values"`
}

// MarshalJSON override to add the name value.
func (t LocalizedEnumType) MarshalJSON() ([]byte, error) {
	type Alias LocalizedEnumType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "LocalizedEnum",
		Alias: (*Alias)(&t),
	})
}

// NumberType will allow the storage of a number value.
type NumberType struct{}

// MarshalJSON override to add the name value.
func (t NumberType) MarshalJSON() ([]byte, error) {
	type Alias NumberType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "Number",
		Alias: (*Alias)(&t),
	})
}

// MoneyType will allow the storage of a money value.
type MoneyType struct{}

// MarshalJSON override to add the name value.
func (t MoneyType) MarshalJSON() ([]byte, error) {
	type Alias MoneyType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "Money",
		Alias: (*Alias)(&t),
	})
}

// DateType will allow the storage of a date value.
type DateType struct{}

// MarshalJSON override to add the name value.
func (t DateType) MarshalJSON() ([]byte, error) {
	type Alias DateType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "Date",
		Alias: (*Alias)(&t),
	})
}

// TimeType will allow the storage of a time value.
type TimeType struct{}

// MarshalJSON override to add the name value.
func (t TimeType) MarshalJSON() ([]byte, error) {
	type Alias TimeType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "Time",
		Alias: (*Alias)(&t),
	})
}

// DateTimeType will allow the storage of a date-time value.
type DateTimeType struct{}

// MarshalJSON override to add the name value.
func (t DateTimeType) MarshalJSON() ([]byte, error) {
	type Alias DateTimeType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "DateTime",
		Alias: (*Alias)(&t),
	})
}

// ReferenceType will allow the storage of a reference value.
type ReferenceType struct {
	ReferenceTypeID string `json:"referenceTypeId"`
}

// MarshalJSON override to add the name value.
func (t ReferenceType) MarshalJSON() ([]byte, error) {
	type Alias ReferenceType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "Reference",
		Alias: (*Alias)(&t),
	})
}

// SetType will allow the storage of a set of fields.
// The set field type defines a set (array without duplicates) with values of the given elementType.
type SetType struct {
	ElementType FieldType `json:"elementType"`
}

// MarshalJSON override to add the name value.
func (t SetType) MarshalJSON() ([]byte, error) {
	type Alias SetType

	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "Set",
		Alias: (*Alias)(&t),
	})
}

// FieldType interface to allow configuration of various field types.
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
		newType.ElementType = fieldTypeMapping(newType.ElementType)
		return newType
	}
	return nil
}
