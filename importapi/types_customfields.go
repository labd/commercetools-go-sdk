// Generated file, please do not change!!!
package importapi

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	The representation to be sent to the server when creating a resource with custom fields.
 */
type Custom struct {
	// The type that provides the field definitions for this object.
	Type TypeKeyReference `json:"type"`
	// The custom fields of this object.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// FieldContainer is something
type FieldContainer map[string]CustomField

/**
*	Provides the value for a custom field of a specific type.
 */
type CustomField interface{}

func mapDiscriminatorCustomField(input interface{}) (CustomField, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "Boolean":
		obj := BooleanField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "String":
		obj := StringField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "LocalizedString":
		obj := LocalizedStringField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Enum":
		obj := EnumField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "LocalizedEnum":
		obj := LocalizedEnumField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Number":
		obj := NumberField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Money":
		obj := MoneyField{}
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
	case "Date":
		obj := DateField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Time":
		obj := TimeField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DateTime":
		obj := DateTimeField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Reference":
		obj := ReferenceField{}
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
	case "BooleanSet":
		obj := BooleanSetField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "StringSet":
		obj := StringSetField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "LocalizedStringSet":
		obj := LocalizedStringSetField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "EnumSet":
		obj := EnumSetField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "LocalizedEnumSet":
		obj := LocalizedEnumSetField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "NumberSet":
		obj := NumberSetField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "MoneySet":
		obj := MoneySetField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DateSet":
		obj := DateSetField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "TimeSet":
		obj := TimeSetField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DateTimeSet":
		obj := DateTimeSetField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ReferenceSet":
		obj := ReferenceSetField{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	A field with a boolean value.
 */
type BooleanField struct {
	Value bool `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BooleanField) MarshalJSON() ([]byte, error) {
	type Alias BooleanField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Boolean", Alias: (*Alias)(&obj)})
}

/**
*	A field with a string value.
 */
type StringField struct {
	Value string `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StringField) MarshalJSON() ([]byte, error) {
	type Alias StringField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "String", Alias: (*Alias)(&obj)})
}

/**
*	A field with a localized string value.
 */
type LocalizedStringField struct {
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	Value LocalizedString `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LocalizedStringField) MarshalJSON() ([]byte, error) {
	type Alias LocalizedStringField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "LocalizedString", Alias: (*Alias)(&obj)})
}

/**
*	A field with a enum value.
 */
type EnumField struct {
	Value string `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj EnumField) MarshalJSON() ([]byte, error) {
	type Alias EnumField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Enum", Alias: (*Alias)(&obj)})
}

/**
*	A field with a localized enum value.
 */
type LocalizedEnumField struct {
	Value string `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LocalizedEnumField) MarshalJSON() ([]byte, error) {
	type Alias LocalizedEnumField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "LocalizedEnum", Alias: (*Alias)(&obj)})
}

/**
*	A field with a number value.
 */
type NumberField struct {
	Value float64 `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj NumberField) MarshalJSON() ([]byte, error) {
	type Alias NumberField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Number", Alias: (*Alias)(&obj)})
}

/**
*	A field with a money value.
 */
type MoneyField struct {
	Value TypedMoney `json:"value"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *MoneyField) UnmarshalJSON(data []byte) error {
	type Alias MoneyField
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
func (obj MoneyField) MarshalJSON() ([]byte, error) {
	type Alias MoneyField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Money", Alias: (*Alias)(&obj)})
}

/**
*	A field with a date value.
 */
type DateField struct {
	Value Date `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DateField) MarshalJSON() ([]byte, error) {
	type Alias DateField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Date", Alias: (*Alias)(&obj)})
}

/**
*	A field with a time value.
 */
type TimeField struct {
	Value time.Time `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TimeField) MarshalJSON() ([]byte, error) {
	type Alias TimeField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Time", Alias: (*Alias)(&obj)})
}

/**
*	A field with a date time value.
 */
type DateTimeField struct {
	Value time.Time `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DateTimeField) MarshalJSON() ([]byte, error) {
	type Alias DateTimeField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DateTime", Alias: (*Alias)(&obj)})
}

/**
*	A field with a reference value.
 */
type ReferenceField struct {
	// References a resource by key
	Value KeyReference `json:"value"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReferenceField) UnmarshalJSON(data []byte) error {
	type Alias ReferenceField
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
func (obj ReferenceField) MarshalJSON() ([]byte, error) {
	type Alias ReferenceField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Reference", Alias: (*Alias)(&obj)})
}

/**
*	A field with a boolean set value.
 */
type BooleanSetField struct {
	Value []bool `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BooleanSetField) MarshalJSON() ([]byte, error) {
	type Alias BooleanSetField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BooleanSet", Alias: (*Alias)(&obj)})
}

/**
*	A field with a string set value.
 */
type StringSetField struct {
	Value []string `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StringSetField) MarshalJSON() ([]byte, error) {
	type Alias StringSetField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "StringSet", Alias: (*Alias)(&obj)})
}

/**
*	A field with a localized string set value.
 */
type LocalizedStringSetField struct {
	Value []LocalizedString `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LocalizedStringSetField) MarshalJSON() ([]byte, error) {
	type Alias LocalizedStringSetField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "LocalizedStringSet", Alias: (*Alias)(&obj)})
}

/**
*	A field with a enum set value.
 */
type EnumSetField struct {
	Value []string `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj EnumSetField) MarshalJSON() ([]byte, error) {
	type Alias EnumSetField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "EnumSet", Alias: (*Alias)(&obj)})
}

/**
*	A field with a localized enum set value.
 */
type LocalizedEnumSetField struct {
	Value []string `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LocalizedEnumSetField) MarshalJSON() ([]byte, error) {
	type Alias LocalizedEnumSetField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "LocalizedEnumSet", Alias: (*Alias)(&obj)})
}

/**
*	A field with a number value.
 */
type NumberSetField struct {
	Value []float64 `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj NumberSetField) MarshalJSON() ([]byte, error) {
	type Alias NumberSetField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "NumberSet", Alias: (*Alias)(&obj)})
}

/**
*	A field with a money set value.
 */
type MoneySetField struct {
	Value []Money `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MoneySetField) MarshalJSON() ([]byte, error) {
	type Alias MoneySetField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "MoneySet", Alias: (*Alias)(&obj)})
}

/**
*	A field with a date set value.
 */
type DateSetField struct {
	Value []Date `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DateSetField) MarshalJSON() ([]byte, error) {
	type Alias DateSetField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DateSet", Alias: (*Alias)(&obj)})
}

/**
*	A field with a time set value.
 */
type TimeSetField struct {
	Value []time.Time `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TimeSetField) MarshalJSON() ([]byte, error) {
	type Alias TimeSetField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "TimeSet", Alias: (*Alias)(&obj)})
}

/**
*	A field with a date time set value.
 */
type DateTimeSetField struct {
	Value []time.Time `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DateTimeSetField) MarshalJSON() ([]byte, error) {
	type Alias DateTimeSetField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "DateTimeSet", Alias: (*Alias)(&obj)})
}

/**
*	A field with a reference set value.
 */
type ReferenceSetField struct {
	Value []KeyReference `json:"value"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReferenceSetField) UnmarshalJSON(data []byte) error {
	type Alias ReferenceSetField
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReferenceSetField) MarshalJSON() ([]byte, error) {
	type Alias ReferenceSetField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReferenceSet", Alias: (*Alias)(&obj)})
}
