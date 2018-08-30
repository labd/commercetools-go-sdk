package producttypes

import (
	"encoding/json"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/mitchellh/mapstructure"
)

// AttributeConstraint enum tells how an attribute or a set of attributes should
// be validated across all variants of a product.
type AttributeConstraint string

const (
	// NoneAttributeConstraint indicates no constraints are applied to the
	// attribute.
	NoneAttributeConstraint AttributeConstraint = "None"

	// UniqueAttributeConstraint indicates attribute values should be different
	// in each variant.
	UniqueAttributeConstraint AttributeConstraint = "Unique"

	// CombinationUniqueAttributeConstraint indicates that a set of attributes,
	// that have this constraint, should have different combinations in each
	// variant.
	CombinationUniqueAttributeConstraint AttributeConstraint = "CombinationUnique"

	// SameForAllAttributeConstraint indicates attribute values should be the
	// same in all variants.
	SameForAllAttributeConstraint AttributeConstraint = "SameForAll"
)

// ProductType is used to describe common characteristics, most importantly
// common custom attributes, of many concrete products.
type ProductType struct {
	// The unique ID of the product type.
	ID string `json:"id"`

	// The current version of the product type.
	Version int `json:"version"`

	// User-specific unique identifier for the product type (max. 256
	// characters).
	Key string `json:"key,omitempty"`

	Name           string                `json:"name"`
	Description    string                `json:"description"`
	Attributes     []AttributeDefinition `json:"attributes"`
	CreatedAt      time.Time             `json:"createdAt"`
	LastModifiedAt time.Time             `json:"lastModifiedAt"`
}

// ProductTypeDraft is given as payload for Create Product Type requests.
type ProductTypeDraft struct {
	Name string `json:"name"`

	// User-specific unique identifier for the product type (min. 2 and max. 256
	// characters).
	Key string `json:"key,omitempty"`

	Description string                `json:"description"`
	Attributes  []AttributeDefinition `json:"attributes,omitempty"`
}

// AttributeDefinition describes a product attribute and allow you to define
// some meta-information associated with the attribute (like whether it should
// be searchable or its constraints).
type AttributeDefinition struct {
	// Describes the type of the attribute.
	Type AttributeType `json:"type"`

	// The unique name of the attribute used in the API. The name must be
	// between two and 256 characters long and can contain the ASCII letters A
	// to Z in lowercase or uppercase, digits, underscores (_) and the
	// hyphen-minus (-). When using the same name for an attribute in two or
	// more product types all fields of the AttributeDefinition of this
	// attribute need to be the same across the product types, otherwise an
	// AttributeDefinitionAlreadyExists error code will be returned. An
	// exception to this are the values of an enum or lenum type and sets
	// thereof.
	Name string `json:"name"`

	// A human-readable label for the attribute.
	Label commercetools.LocalizedString `json:"label"`

	// Whether the attribute is required to have a value.
	IsRequired bool `json:"isRequired"`

	// Describes how an attribute or a set of attributes should be validated
	// across all variants of a product.
	AttributeConstraint AttributeConstraint `json:"attributeConstraint"`

	// Additional information about the attribute that aids content managers
	// when setting product details.
	InputTip commercetools.LocalizedString `json:"inputTip,omitempty"`

	// Provides a visual representation type for this attribute. Only relevant
	// for text-based attribute types like TextType and LocalizableTextType.
	InputHint commercetools.TextInputHint `json:"inputHint"`

	// Whether the attribute’s values should generally be enabled in product
	// search. This determines whether the value is stored in products for
	// matching terms in the context of full-text search queries and can be used
	// in facets & filters as part of product search queries. The exact features
	// that are enabled/disabled with this flag depend on the concrete attribute
	// type and are described there. The max size of a searchable field is
	// restricted to 10922 characters. This constraint is enforced at both
	// product creation and product update. If the length of the input exceeds
	// the maximum size an InvalidField error is returned.
	IsSearchable bool `json:"isSearchable"`
}

// UnmarshalJSON override to map the attribute definition to the corresponding
// struct.
func (a *AttributeDefinition) UnmarshalJSON(data []byte) error {
	type Alias AttributeDefinition
	if err := json.Unmarshal(data, (*Alias)(a)); err != nil {
		return err
	}
	a.Type = attributeTypeMapping(a.Type)
	return nil
}

// AttributeDefinitionDraft is given as payload for Create Attribute Definition
// requests.
type AttributeDefinitionDraft struct {
	// Describes the type of the attribute.
	Type AttributeType `json:"type"`

	// The unique name of the attribute used in the API. The name must be
	// between two and 256 characters long and can contain the ASCII letters A
	// to Z in lowercase or uppercase, digits, underscores (_) and the
	// hyphen-minus (-). When using the same name for an attribute in two or
	// more product types all fields of the AttributeDefinition of this
	// attribute need to be the same across the product types.
	Name string `json:"name"`

	// A human-readable label for the attribute.
	Label commercetools.LocalizedString `json:"label"`

	// Whether the attribute is required to have a value.
	IsRequired bool `json:"isRequired"`

	// Describes how an attribute or a set of attributes should be validated
	// across all variants of a product.
	AttributeConstraint AttributeConstraint `json:"attributeConstraint"`

	// Additional information about the attribute that aids content managers
	// when setting product details.
	InputTip commercetools.LocalizedString `json:"inputTip,omitempty"`

	// Provides a visual representation type for this attribute. Only relevant
	// for text-based attribute types like TextType and LocalizableTextType.
	InputHint commercetools.TextInputHint `json:"inputHint,omitempty"`

	// Whether the attribute’s values should generally be enabled in product
	// search. This determines whether the value is stored in products for
	// matching terms in the context of full-text search queries and can be used
	// in facets & filters as part of product search queries. The exact features
	// that are enabled/disabled with this flag depend on the concrete attribute
	// type and are described there.
	IsSearchable bool `json:"isSearchable"`
}

// UnmarshalJSON override to map the attribute definition draft to the
// corresponding struct.
func (a *AttributeDefinitionDraft) UnmarshalJSON(data []byte) error {
	type Alias AttributeDefinitionDraft
	if err := json.Unmarshal(data, (*Alias)(a)); err != nil {
		return err
	}
	a.Type = attributeTypeMapping(a.Type)
	return nil
}

// BooleanType indicates the attribute can store a boolean. Valid values for the
// attribute are true and false (JSON Boolean).
type BooleanType struct{}

// MarshalJSON override to add the name value
func (t BooleanType) MarshalJSON() ([]byte, error) {
	type Alias BooleanType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "boolean",
		Alias: (*Alias)(&t),
	})
}

// TextType indicates the attribute can store text.
type TextType struct{}

// MarshalJSON override to add the name value
func (t TextType) MarshalJSON() ([]byte, error) {
	type Alias TextType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "text",
		Alias: (*Alias)(&t),
	})
}

// LocalizedTextType indicates the attribute can store localized text.
type LocalizedTextType struct{}

// MarshalJSON override to add the name value
func (t LocalizedTextType) MarshalJSON() ([]byte, error) {
	type Alias LocalizedTextType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "ltext",
		Alias: (*Alias)(&t),
	})
}

// EnumType indicates the attribute can store an enum.
type EnumType struct {
	// The available values that can be assigned to products.
	Values []commercetools.EnumValue `json:"values"`
}

// MarshalJSON override to add the name value
func (t EnumType) MarshalJSON() ([]byte, error) {
	type Alias EnumType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "enum",
		Alias: (*Alias)(&t),
	})
}

// LocalizedEnumType indicates the attribute can store a localized enum.
type LocalizedEnumType struct {
	// The available values that can be assigned to products.
	Values []commercetools.LocalizedEnumValue `json:"values"`
}

// MarshalJSON override to add the name value
func (t LocalizedEnumType) MarshalJSON() ([]byte, error) {
	type Alias LocalizedEnumType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "lenum",
		Alias: (*Alias)(&t),
	})
}

// NumberType indicates the attribute can store a number.
type NumberType struct{}

// MarshalJSON override to add the name value
func (t NumberType) MarshalJSON() ([]byte, error) {
	type Alias NumberType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "number",
		Alias: (*Alias)(&t),
	})
}

// MoneyType indicates the attribute can store a money value.
type MoneyType struct{}

// MarshalJSON override to add the name value
func (t MoneyType) MarshalJSON() ([]byte, error) {
	type Alias MoneyType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "money",
		Alias: (*Alias)(&t),
	})
}

// DateType indicates the attribute can store a date.
type DateType struct{}

// MarshalJSON override to add the name value
func (t DateType) MarshalJSON() ([]byte, error) {
	type Alias DateType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "date",
		Alias: (*Alias)(&t),
	})
}

// TimeType indicates the attribute can store a time.
type TimeType struct{}

// MarshalJSON override to add the name value
func (t TimeType) MarshalJSON() ([]byte, error) {
	type Alias TimeType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "time",
		Alias: (*Alias)(&t),
	})
}

// DateTimeType indicates the attribute can store a date-time value.
type DateTimeType struct{}

// MarshalJSON override to add the name value
func (t DateTimeType) MarshalJSON() ([]byte, error) {
	type Alias DateTimeType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "datetime",
		Alias: (*Alias)(&t),
	})
}

// ReferenceType indicates the attribute can store a reference.
type ReferenceType struct {
	// The name of the resource type that the value should reference.
	ReferenceTypeID string `json:"referenceTypeId"`
}

// MarshalJSON override to add the name value
func (t ReferenceType) MarshalJSON() ([]byte, error) {
	type Alias ReferenceType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "reference",
		Alias: (*Alias)(&t),
	})
}

// SetType defines a set (array without duplicates) with values of the given
// elementType. It does not support isRequired.
type SetType struct {
	ElementType AttributeType `json:"elementType"`
}

// MarshalJSON override to add the name value
func (t SetType) MarshalJSON() ([]byte, error) {
	type Alias SetType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "set",
		Alias: (*Alias)(&t),
	})
}

// NestedType allows you to nest attributes based on some existing product type.
// It does not support isSearchable nor it is supported in queries at the
// moment. The only supported AttributeConstraint is None. The value of the
// nested attribute is an array of values. It reflects the structure of the
// attributes property of product variant, where every element of array is a
// JSON object with properties name and value.
type NestedType struct {
	// ProductType defines, which attributes are allowed to be stored as a
	// nested attributes of the current attribute.
	TypeReference string `json:"typeReference"`
}

// MarshalJSON override to add the name value TODO: Test with nested response.
func (t NestedType) MarshalJSON() ([]byte, error) {
	type Alias NestedType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{
		Name:  "nested",
		Alias: (*Alias)(&t),
	})
}

// AttributeType all have a name. Some have additional fields such as values in
// enums or elementType in sets.
type AttributeType interface{}

func attributeTypeMapping(input AttributeType) AttributeType {
	AttributeType := input.(map[string]interface{})["name"]
	switch AttributeType {
	case "boolean":
		newType := BooleanType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "text":
		newType := TextType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "ltext":
		newType := LocalizedTextType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "enum":
		newType := EnumType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "lenum":
		newType := LocalizedEnumType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "number":
		newType := NumberType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "money":
		newType := MoneyType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "date":
		newType := DateType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "time":
		newType := TimeType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "datetime":
		newType := DateTimeType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "reference":
		newType := ReferenceType{}
		mapstructure.Decode(input, &newType)
		return newType
	case "set":
		newType := SetType{}
		mapstructure.Decode(input, &newType)
		newType.ElementType = attributeTypeMapping(newType.ElementType)
		return newType
	case "nested":
		newType := NestedType{}
		mapstructure.Decode(input, &newType)
		return newType
	}
	return nil
}
