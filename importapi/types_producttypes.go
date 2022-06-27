package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

type AttributeDefinition struct {
	Type AttributeType `json:"type"`
	Name string        `json:"name"`
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	Label               LocalizedString          `json:"label"`
	IsRequired          bool                     `json:"isRequired"`
	AttributeConstraint *AttributeConstraintEnum `json:"attributeConstraint,omitempty"`
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	InputTip     *LocalizedString `json:"inputTip,omitempty"`
	InputHint    *TextInputHint   `json:"inputHint,omitempty"`
	IsSearchable *bool            `json:"isSearchable,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AttributeDefinition) UnmarshalJSON(data []byte) error {
	type Alias AttributeDefinition
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Type != nil {
		var err error
		obj.Type, err = mapDiscriminatorAttributeType(obj.Type)
		if err != nil {
			return err
		}
	}

	return nil
}

type AttributeType interface{}

func mapDiscriminatorAttributeType(input interface{}) (AttributeType, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["name"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'name'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "boolean":
		obj := AttributeBooleanType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "datetime":
		obj := AttributeDateTimeType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "date":
		obj := AttributeDateType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "enum":
		obj := AttributeEnumType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ltext":
		obj := AttributeLocalizableTextType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "lenum":
		obj := AttributeLocalizedEnumType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "money":
		obj := AttributeMoneyType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "nested":
		obj := AttributeNestedType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "number":
		obj := AttributeNumberType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "reference":
		obj := AttributeReferenceType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "set":
		obj := AttributeSetType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.ElementType != nil {
			var err error
			obj.ElementType, err = mapDiscriminatorAttributeType(obj.ElementType)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "text":
		obj := AttributeTextType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "time":
		obj := AttributeTimeType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type AttributeBooleanType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeBooleanType) MarshalJSON() ([]byte, error) {
	type Alias AttributeBooleanType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "boolean", Alias: (*Alias)(&obj)})
}

type AttributeDateTimeType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeDateTimeType) MarshalJSON() ([]byte, error) {
	type Alias AttributeDateTimeType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "datetime", Alias: (*Alias)(&obj)})
}

type AttributeDateType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeDateType) MarshalJSON() ([]byte, error) {
	type Alias AttributeDateType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "date", Alias: (*Alias)(&obj)})
}

type AttributeEnumType struct {
	Values []AttributePlainEnumValue `json:"values"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeEnumType) MarshalJSON() ([]byte, error) {
	type Alias AttributeEnumType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "enum", Alias: (*Alias)(&obj)})
}

type AttributePlainEnumValue struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

type AttributeLocalizableTextType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeLocalizableTextType) MarshalJSON() ([]byte, error) {
	type Alias AttributeLocalizableTextType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "ltext", Alias: (*Alias)(&obj)})
}

type AttributeLocalizedEnumType struct {
	Values []AttributeLocalizedEnumValue `json:"values"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeLocalizedEnumType) MarshalJSON() ([]byte, error) {
	type Alias AttributeLocalizedEnumType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "lenum", Alias: (*Alias)(&obj)})
}

type AttributeLocalizedEnumValue struct {
	Key string `json:"key"`
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	Label LocalizedString `json:"label"`
}

type AttributeMoneyType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeMoneyType) MarshalJSON() ([]byte, error) {
	type Alias AttributeMoneyType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "money", Alias: (*Alias)(&obj)})
}

type AttributeNestedType struct {
	// References a product type by key.
	TypeReference ProductTypeKeyReference `json:"typeReference"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeNestedType) MarshalJSON() ([]byte, error) {
	type Alias AttributeNestedType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "nested", Alias: (*Alias)(&obj)})
}

type AttributeNumberType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeNumberType) MarshalJSON() ([]byte, error) {
	type Alias AttributeNumberType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "number", Alias: (*Alias)(&obj)})
}

type AttributeReferenceType struct {
	// The type of the referenced resource.
	ReferenceTypeId ReferenceType `json:"referenceTypeId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeReferenceType) MarshalJSON() ([]byte, error) {
	type Alias AttributeReferenceType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "reference", Alias: (*Alias)(&obj)})
}

type AttributeSetType struct {
	ElementType AttributeType `json:"elementType"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AttributeSetType) UnmarshalJSON(data []byte) error {
	type Alias AttributeSetType
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ElementType != nil {
		var err error
		obj.ElementType, err = mapDiscriminatorAttributeType(obj.ElementType)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeSetType) MarshalJSON() ([]byte, error) {
	type Alias AttributeSetType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "set", Alias: (*Alias)(&obj)})
}

type AttributeTextType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeTextType) MarshalJSON() ([]byte, error) {
	type Alias AttributeTextType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "text", Alias: (*Alias)(&obj)})
}

type AttributeTimeType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj AttributeTimeType) MarshalJSON() ([]byte, error) {
	type Alias AttributeTimeType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "time", Alias: (*Alias)(&obj)})
}

type AttributeConstraintEnum string

const (
	AttributeConstraintEnumNone              AttributeConstraintEnum = "None"
	AttributeConstraintEnumUnique            AttributeConstraintEnum = "Unique"
	AttributeConstraintEnumCombinationUnique AttributeConstraintEnum = "CombinationUnique"
	AttributeConstraintEnumSameForAll        AttributeConstraintEnum = "SameForAll"
)

type TextInputHint string

const (
	TextInputHintSingleLine TextInputHint = "SingleLine"
	TextInputHintMultiLine  TextInputHint = "MultiLine"
)

/**
*	The data representation for a ProductType to be imported that is persisted as a [ProductType](/../api/projects/productTypes#producttype) in the Project.
*
 */
type ProductTypeImport struct {
	Key string `json:"key"`
	// Maps to `ProductType.name`.
	Name string `json:"name"`
	// Maps to `ProductType.description`.
	Description string `json:"description"`
	// The `attributes` of [ProductType](/../api/projects/productTypes#producttype).
	Attributes []AttributeDefinition `json:"attributes"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeImport) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeImport
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

	return json.Marshal(raw)

}
