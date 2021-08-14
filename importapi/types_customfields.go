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
	Fields *FieldContainer `json:"fields,omitEmpty"`
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
		new := BooleanField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "String":
		new := StringField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "LocalizedString":
		new := LocalizedStringField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Enum":
		new := EnumField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "LocalizedEnum":
		new := LocalizedEnumField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Number":
		new := NumberField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Money":
		new := MoneyField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Date":
		new := DateField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Time":
		new := TimeField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DateTime":
		new := DateTimeField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Reference":
		new := ReferenceField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "BooleanSet":
		new := BooleanSetField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "StringSet":
		new := StringSetField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "LocalizedStringSet":
		new := LocalizedStringSetField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "EnumSet":
		new := EnumSetField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "LocalizedEnumSet":
		new := LocalizedEnumSetField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "NumberSet":
		new := NumberSetField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "MoneySet":
		new := MoneySetField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DateSet":
		new := DateSetField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "TimeSet":
		new := TimeSetField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DateTimeSet":
		new := DateTimeSetField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "ReferenceSet":
		new := ReferenceSetField{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

/**
*	A field with a boolean value.
 */
type BooleanField struct {
	Value bool `json:"value"`
}

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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
	// References a resource by its key
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
func (obj ReferenceSetField) MarshalJSON() ([]byte, error) {
	type Alias ReferenceSetField
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ReferenceSet", Alias: (*Alias)(&obj)})
}
