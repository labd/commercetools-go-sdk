package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	Represents the value of an Attribute of a Product Variant.
*
 */
type Attribute interface{}

func mapDiscriminatorAttribute(input interface{}) (Attribute, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "boolean":
		obj := BooleanAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "boolean-set":
		obj := BooleanSetAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "date":
		obj := DateAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "date-set":
		obj := DateSetAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "datetime":
		obj := DateTimeAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "datetime-set":
		obj := DateTimeSetAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "enum":
		obj := EnumAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "enum-set":
		obj := EnumSetAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "lenum":
		obj := LocalizableEnumAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "lenum-set":
		obj := LocalizableEnumSetAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ltext":
		obj := LocalizableTextAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ltext-set":
		obj := LocalizableTextSetAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "money":
		obj := MoneyAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Value != nil {
			var err error
			obj.Value, err = mapDiscriminatorTypedMoney(obj.Value)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "money-set":
		obj := MoneySetAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		for i := range obj.Value {
			var err error
			obj.Value[i], err = mapDiscriminatorTypedMoney(obj.Value[i])
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "number":
		obj := NumberAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "number-set":
		obj := NumberSetAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "reference":
		obj := ReferenceAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Value != nil {
			var err error
			obj.Value, err = mapDiscriminatorKeyReference(obj.Value)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "reference-set":
		obj := ReferenceSetAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		for i := range obj.Value {
			var err error
			obj.Value[i], err = mapDiscriminatorKeyReference(obj.Value[i])
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "text":
		obj := TextAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "text-set":
		obj := TextSetAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "time":
		obj := TimeAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "time-set":
		obj := TimeSetAttribute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	This type represents an attribute whose value is either "true" or "false".
*
 */
type BooleanAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// `true` or `false`
	Value bool `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BooleanAttribute) MarshalJSON() ([]byte, error) {
	type Alias BooleanAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "boolean", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is set of boolean values.
*
 */
type BooleanSetAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A set of boolean values.
	Value []bool `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BooleanSetAttribute) MarshalJSON() ([]byte, error) {
	type Alias BooleanSetAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "boolean-set", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a date.
*
 */
type DateAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A date in the format `YYYY-MM-DD`.
	Value Date `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DateAttribute) MarshalJSON() ([]byte, error) {
	type Alias DateAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "date", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a set of dates.
*
 */
type DateSetAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A set of dates in the format `YYYY-MM-DD`.
	Value []Date `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DateSetAttribute) MarshalJSON() ([]byte, error) {
	type Alias DateSetAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "date-set", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a date with time.
*
 */
type DateTimeAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A date with time in the format `YYYY-MM-DDTHH:mm:ss.SSSZ`.
	// The time zone is optional and defaults to UTC if not specified.
	// If the time zone is specified, it must be in the format `±HH:mm` or `Z` for UTC.
	Value time.Time `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DateTimeAttribute) MarshalJSON() ([]byte, error) {
	type Alias DateTimeAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "datetime", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a set of dates with time.
*
 */
type DateTimeSetAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A set of dates with time in the format `YYYY-MM-DDTHH:mm:ss.SSSZ`.
	// The time zone is optional and defaults to UTC if not specified.
	// If the time zone is specified, it must be in the format `±HH:mm` or `Z` for UTC.
	Value []time.Time `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DateTimeSetAttribute) MarshalJSON() ([]byte, error) {
	type Alias DateTimeSetAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "datetime-set", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is an enum.
*	The attribute value refers to the key of the enum value.
*
 */
type EnumAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// The key of the enum value.
	// Must match the key of an [AttributePlainEnumValue](ctp:api:type:AttributePlainEnumValue) in the Product Type.
	Value string `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj EnumAttribute) MarshalJSON() ([]byte, error) {
	type Alias EnumAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "enum", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is an enum.
*	The attribute value refers to the key of the enum value.
*
 */
type EnumSetAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A set of enum values, each represented by its key.
	// Each key must match the key of an [AttributePlainEnumValue](ctp:api:type:AttributePlainEnumValue) in the Product Type.
	Value []string `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj EnumSetAttribute) MarshalJSON() ([]byte, error) {
	type Alias EnumSetAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "enum-set", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a localized enum.
*	The attribute value refers to the key of the enum value.
*
 */
type LocalizableEnumAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// The key of the localized enum value.
	// Must match the key of an [AttributeLocalizedEnumValue](ctp:api:type:AttributeLocalizedEnumValue) in the Product Type.
	Value string `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LocalizableEnumAttribute) MarshalJSON() ([]byte, error) {
	type Alias LocalizableEnumAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "lenum", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a localized enum.
*	The attribute value refers to the key of the enum value.
*
 */
type LocalizableEnumSetAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A set of localized enum values, each represented by its key.
	// Each key must match the key of an [AttributeLocalizedEnumValue](ctp:api:type:AttributeLocalizedEnumValue) in the Product Type.
	Value []string `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LocalizableEnumSetAttribute) MarshalJSON() ([]byte, error) {
	type Alias LocalizableEnumSetAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "lenum-set", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a localized text.
*
 */
type LocalizableTextAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A localized string.
	Value LocalizedString `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LocalizableTextAttribute) MarshalJSON() ([]byte, error) {
	type Alias LocalizableTextAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ltext", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a localized text.
*
 */
type LocalizableTextSetAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A set of localized strings.
	Value []LocalizedString `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LocalizableTextSetAttribute) MarshalJSON() ([]byte, error) {
	type Alias LocalizableTextSetAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ltext-set", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a money object.
*
 */
type MoneyAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A money value in cent precision format.
	Value TypedMoney `json:"value"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MoneyAttribute) UnmarshalJSON(data []byte) error {
	type Alias MoneyAttribute
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorTypedMoney(obj.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MoneyAttribute) MarshalJSON() ([]byte, error) {
	type Alias MoneyAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "money", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a set of money objects.
*
 */
type MoneySetAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A set of money values in cent precision format.
	Value []TypedMoney `json:"value"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MoneySetAttribute) UnmarshalJSON(data []byte) error {
	type Alias MoneySetAttribute
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Value {
		var err error
		obj.Value[i], err = mapDiscriminatorTypedMoney(obj.Value[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MoneySetAttribute) MarshalJSON() ([]byte, error) {
	type Alias MoneySetAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "money-set", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a number.
*
 */
type NumberAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A number value.
	// Can be an integer or a floating-point number.
	Value float64 `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj NumberAttribute) MarshalJSON() ([]byte, error) {
	type Alias NumberAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "number", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a set of numbers.
*
 */
type NumberSetAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A set of number values.
	// Each value can be an integer or a floating-point number.
	Value []float64 `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj NumberSetAttribute) MarshalJSON() ([]byte, error) {
	type Alias NumberSetAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "number-set", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a key reference.
*
 */
type ReferenceAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// References a resource by key.
	Value KeyReference `json:"value"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReferenceAttribute) UnmarshalJSON(data []byte) error {
	type Alias ReferenceAttribute
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorKeyReference(obj.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReferenceAttribute) MarshalJSON() ([]byte, error) {
	type Alias ReferenceAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "reference", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a set of references.
*
 */
type ReferenceSetAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A set of references, each referencing a resource by key.
	// Each reference must match the key of an existing resource in the project.
	Value []KeyReference `json:"value"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReferenceSetAttribute) UnmarshalJSON(data []byte) error {
	type Alias ReferenceSetAttribute
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Value {
		var err error
		obj.Value[i], err = mapDiscriminatorKeyReference(obj.Value[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReferenceSetAttribute) MarshalJSON() ([]byte, error) {
	type Alias ReferenceSetAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "reference-set", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a string.
*
 */
type TextAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A text value.
	Value string `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TextAttribute) MarshalJSON() ([]byte, error) {
	type Alias TextAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "text", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a set of strings.
*
 */
type TextSetAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A set of text values.
	Value []string `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TextSetAttribute) MarshalJSON() ([]byte, error) {
	type Alias TextSetAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "text-set", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a time.
*
 */
type TimeAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A time value in the format `HH:mm:ss.SSS`.
	// The time zone is optional and defaults to UTC if not specified.
	// If the time zone is specified, it must be in the format `±HH:mm` or `Z` for UTC.
	Value time.Time `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TimeAttribute) MarshalJSON() ([]byte, error) {
	type Alias TimeAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "time", Alias: (*Alias)(&obj)})
}

/**
*	This type represents an attribute whose value is a set of times.
*
 */
type TimeSetAttribute struct {
	// Required if used for [ProductVariantImport](ctp:import:type:ProductVariantImport).
	// Must not be set if used for [ProductVariantPatch](ctp:import:type:ProductVariantPatch).
	//
	// Must match `name` of an [AttributeDefinition](ctp:api:type:AttributeDefinition) of the Product Type.
	Name *string `json:"name,omitempty"`
	// A set of time values in the format `HH:mm:ss.SSS`.
	// The time zone is optional and defaults to UTC if not specified.
	// If the time zone is specified, it must be in the format `±HH:mm` or `Z` for UTC.
	Value []time.Time `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TimeSetAttribute) MarshalJSON() ([]byte, error) {
	type Alias TimeSetAttribute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "time-set", Alias: (*Alias)(&obj)})
}

/**
*	Represents the data used to import a ProductVariant. Once imported, this data is persisted as a [ProductVariant](ctp:api:type:ProductVariant) in the Project.
*
 */
type ProductVariantImport struct {
	// User-defined unique identifier. If a [ProductVariant](ctp:api:type:ProductVariant) with this `key` exists on the specified `product`, it is updated with the imported data.
	Key string `json:"key"`
	// Maps to `ProductVariant.sku`.
	Sku *string `json:"sku,omitempty"`
	// - When creating a new ProductVariant, set to `false`; otherwise, the import operation will fail with a [NewMasterVariantAdditionNotAllowed](ctp:import:type:NewMasterVariantAdditionNotAllowedError) error.
	// - Set to `true` if the ProductVariant exists and you want to set this ProductVariant as the Master Variant.
	IsMasterVariant bool `json:"isMasterVariant"`
	// Maps to `ProductVariant.attributes`.
	// The referenced attribute must be defined in an already existing ProductType in the project, or the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be `unresolved`.
	Attributes []Attribute `json:"attributes"`
	// Maps to `ProductVariant.images`.
	Images []Image `json:"images"`
	// Maps to `ProductVariant.assets`.
	Assets []Asset `json:"assets"`
	// If `publish` is set to either `true` or `false`, both staged and current projections are set to the same value provided by the import data.
	// If `publish` is not set, the staged projection is set to the provided import data, but the current projection stays unchanged.
	// However, if the import data contains no update, that is, if it matches the staged projection of the existing Product, the import induces no change in the existing Product whether `publish` is set or not.
	Publish *bool `json:"publish,omitempty"`
	// - Set to `false` to update both the [current and staged projections](/../api/projects/productProjections#current--staged) of the [Product](ctp:api:type:Product) with the new Product Variant data.
	// - Leave empty or set to `true` to only update the staged projection.
	Staged *bool `json:"staged,omitempty"`
	// The [Product](ctp:api:type:ProductVariant) containing this ProductVariant. If the referenced Product does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Product is created.
	Product ProductKeyReference `json:"product"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductVariantImport) UnmarshalJSON(data []byte) error {
	type Alias ProductVariantImport
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Attributes {
		var err error
		obj.Attributes[i], err = mapDiscriminatorAttribute(obj.Attributes[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariantImport) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantImport
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["attributes"] == nil {
		delete(raw, "attributes")
	}

	if raw["images"] == nil {
		delete(raw, "images")
	}

	if raw["assets"] == nil {
		delete(raw, "assets")
	}

	return json.Marshal(raw)

}

/**
*	Represents the data used to update a [ProductVariant](ctp:api:type:ProductVariant).
*
 */
type ProductVariantPatch struct {
	// Reference to the [ProductVariant](ctp:api:type:ProductVariant) to update.
	ProductVariant ProductVariantKeyReference `json:"productVariant"`
	// Maps to `ProductVariant.attributes`.
	// - The referenced Attribute must be defined in an existing [ProductType](ctp:api:type:ProductType), or the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be `validationFailed`.
	// - Setting the value of a non-required Attribute to `null` will remove the Attribute.
	// - Attempting to set a `null` value to a required Attribute will make the import operation fail with an [InvalidOperation](ctp:import:type:InvalidOperation) error.
	// - Importing [LocalizableTextAttributes](ctp:import:type:LocalizableTextAttribute) or [LocalizableTextSetAttributes](ctp:import:type:LocalizableTextSetAttribute) follows an override pattern, meaning that omitted localized fields will be deleted, new fields will be created, and existing fields will be updated. You can delete localized fields by setting their value to `null`.
	Attributes *Attributes `json:"attributes,omitempty"`
	// If `false`, the attribute changes are applied to both [current and staged projected representations](/projects/productProjections#current--staged) of the [Product](ctp:api:type:Product).
	Staged *bool `json:"staged,omitempty"`
	// Reference to the [Product](ctp:api:type:Product) that contains the ProductVariant.
	//
	// We recommend to set this value to minimize concurrency errors.
	// If set, this field is required for every ProductVariantPatch in the [ProductVariantPatchRequest](ctp:import:type:ProductVariantPatchRequest).
	//
	// If the referenced Product does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Product is created.
	Product *ProductKeyReference `json:"product,omitempty"`
}

type Attributes map[string]interface{}
