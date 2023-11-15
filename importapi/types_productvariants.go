package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	This type represents the value of an attribute of a product variant.
*	The name and type property must match the name and type property of an attribute definition of the product type.
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string `json:"name,omitempty"`
	Value bool    `json:"value"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string `json:"name,omitempty"`
	Value []bool  `json:"value"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string `json:"name,omitempty"`
	Value Date    `json:"value"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string `json:"name,omitempty"`
	Value []Date  `json:"value"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string   `json:"name,omitempty"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string     `json:"name,omitempty"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string `json:"name,omitempty"`
	Value string  `json:"value"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string  `json:"name,omitempty"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string `json:"name,omitempty"`
	Value string  `json:"value"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string  `json:"name,omitempty"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name *string `json:"name,omitempty"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string           `json:"name,omitempty"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string    `json:"name,omitempty"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string      `json:"name,omitempty"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string `json:"name,omitempty"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string   `json:"name,omitempty"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string        `json:"name,omitempty"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string `json:"name,omitempty"`
	Value string  `json:"value"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string  `json:"name,omitempty"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string   `json:"name,omitempty"`
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
	// The name of this attribute must match a name of the product types attribute definitions.
	// The name is required if this type is used in a product variant and must not be set when
	// used in a product variant patch.
	Name  *string     `json:"name,omitempty"`
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
*	The data representation for a ProductVariant to be imported that is persisted as a [ProductVariant](/../api/projects/products#productvariant) in the Project.
*
 */
type ProductVariantImport struct {
	// User-defined unique identifier. If a [ProductVariant](/../api/projects/products#productvariant) with this `key` exists on the specified `product`, it will be updated with the imported data.
	Key string `json:"key"`
	// Maps to `ProductVariant.sku`.
	Sku *string `json:"sku,omitempty"`
	// Maps to `ProductVariant.isMasterVariant`.
	IsMasterVariant bool `json:"isMasterVariant"`
	// Maps to `ProductVariant.attributes`.
	// The referenced attribute must be defined in an already existing ProductType in the project, or the `state` of the [ImportOperation](/import-operation#importoperation) will be `unresolved`.
	Attributes []Attribute `json:"attributes"`
	// Maps to `ProductVariant.images`.
	Images []Image `json:"images"`
	// Maps to `ProductVariant.assets`.
	Assets []Asset `json:"assets"`
	// If `publish` is set to either `true` or `false`, both staged and current projections are set to the same value provided by the import data.
	// If `publish` is not set, the staged projection is set to the provided import data, but the current projection stays unchanged.
	// However, if the import data contains no update, that is, if it matches the staged projection of the existing Product, the import induces no change in the existing Product whether `publish` is set or not.
	Publish *bool `json:"publish,omitempty"`
	// The [Product](/../api/projects/products#productvariant) to which this Product Variant belongs. Maps to `ProductVariant.product`.
	// The Reference to the [Product](/../api/projects/products#product) with which the ProductVariant is associated.
	// If referenced Product does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary Product is created.
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
*	Representation for an update of a [ProductVariant](/../api/projects/products#productvariant). Use this type to import updates for existing
*	[ProductVariants](/../api/projects/products#productvariant) in a Project.
*
 */
type ProductVariantPatch struct {
	// Reference to the [ProductVariant](/../api/projects/products#productvariant) to update.
	// If the referenced ProductVariant does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary ProductVariant is created.
	ProductVariant ProductVariantKeyReference `json:"productVariant"`
	// Maps to `ProductVariant.attributes`.
	// - The referenced Attribute must be defined in an existing [ProductType](/../api/projects/productTypes#producttype), or the `state` of the [ImportOperation](/import-operation#importoperation) will be `validationFailed`.
	// - Setting the value of a non-required Attribute to `null` will remove the Attribute.
	// - Attempting to set a `null` value to a required Attribute will make the import operation fail with an [InvalidOperation](/error#invalidoperation) error.
	// - Importing [LocalizableTextAttributes](/product-variant#localizabletextattribute) or [LocalizableTextSetAttributes](/product-variant#localizabletextsetattribute) follows an override pattern, meaning that omitted localized fields will be deleted, new fields will be created, and existing fields will be updated. You can also delete localized fields by setting their value to `null`.
	Attributes *Attributes `json:"attributes,omitempty"`
	// If `false`, the attribute changes are applied to both [current and staged projected representations](/../api/projects/productProjections#current--staged) of the [Product](/../api/projects/products#product).
	Staged *bool `json:"staged,omitempty"`
	// Reference to the [Product](/../api/projects/products#product) which contains the ProductVariant. Setting a value will batch process the import operations to minimize concurrency errors. If set, this field is required for every ProductVariantPatch in the [ProductVariantPatchRequest](ctp:import:type:ProductVariantPatchRequest).
	Product *ProductKeyReference `json:"product,omitempty"`
}

type Attributes map[string]interface{}
