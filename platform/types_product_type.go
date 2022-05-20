package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type AttributeConstraintEnum string

const (
	AttributeConstraintEnumNone              AttributeConstraintEnum = "None"
	AttributeConstraintEnumUnique            AttributeConstraintEnum = "Unique"
	AttributeConstraintEnumCombinationUnique AttributeConstraintEnum = "CombinationUnique"
	AttributeConstraintEnumSameForAll        AttributeConstraintEnum = "SameForAll"
)

type AttributeConstraintEnumDraft string

const (
	AttributeConstraintEnumDraftNone AttributeConstraintEnumDraft = "None"
)

type AttributeDefinition struct {
	// Describes the type of the attribute.
	Type AttributeType `json:"type"`
	// The unique name of the attribute used in the API.
	// The name must be between two and 256 characters long and can contain the ASCII letters A to Z in lowercase or uppercase, digits, underscores (`_`) and the hyphen-minus (`-`).
	// When using the same `name` for an attribute in two or more product types all fields of the AttributeDefinition of this attribute need to be the same across the product types, otherwise an AttributeDefinitionAlreadyExists error code will be returned.
	// An exception to this are the values of an `enum` or `lenum` type and sets thereof.
	Name string `json:"name"`
	// A human-readable label for the attribute.
	Label LocalizedString `json:"label"`
	// Whether the attribute is required to have a value.
	IsRequired bool `json:"isRequired"`
	// Describes how an attribute or a set of attributes should be validated across all variants of a product.
	AttributeConstraint AttributeConstraintEnum `json:"attributeConstraint"`
	// Additional information about the attribute that aids content managers when setting product details.
	InputTip *LocalizedString `json:"inputTip,omitempty"`
	// Provides a visual representation type for this attribute.
	// only relevant for text-based attribute types
	// like TextType and LocalizableTextType.
	InputHint TextInputHint `json:"inputHint"`
	// Whether the attribute's values should generally be enabled in product search.
	// This determines whether the value is stored in products for matching terms in the context of full-text search queries  and can be used in facets & filters as part of product search queries.
	// The exact features that are enabled/disabled with this flag depend on the concrete attribute type and are described there.
	// The max size of a searchable field is **restricted to 10922 characters**.
	// This constraint is enforced at both product creation and product update.
	// If the length of the input exceeds the maximum size an InvalidField error is returned.
	IsSearchable bool `json:"isSearchable"`
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

type AttributeDefinitionDraft struct {
	// Describes the type of the attribute.
	Type AttributeType `json:"type"`
	// The unique name of the attribute used in the API.
	// The name must be between two and 256 characters long and can contain the ASCII letters A to Z in lowercase or uppercase, digits, underscores (`_`) and the hyphen-minus (`-`).
	// When using the same `name` for an attribute in two or more product types all fields of the AttributeDefinition of this attribute need to be the same across the product types.
	Name string `json:"name"`
	// A human-readable label for the attribute.
	Label LocalizedString `json:"label"`
	// Whether the attribute is required to have a value.
	IsRequired bool `json:"isRequired"`
	// Describes how an attribute or a set of attributes should be validated across all variants of a product.
	AttributeConstraint *AttributeConstraintEnum `json:"attributeConstraint,omitempty"`
	// Additional information about the attribute that aids content managers when setting product details.
	InputTip *LocalizedString `json:"inputTip,omitempty"`
	// Provides a visual representation type for this attribute.
	// only relevant for text-based attribute types like TextType and LocalizableTextType.
	InputHint *TextInputHint `json:"inputHint,omitempty"`
	// Whether the attribute's values should generally be enabled in product search.
	// This determines whether the value is stored in products for matching terms in the context of full-text search queries and can be used in facets & filters as part of product search queries.
	// The exact features that are enabled/disabled with this flag depend on the concrete attribute type and are described there.
	IsSearchable *bool `json:"isSearchable,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AttributeDefinitionDraft) UnmarshalJSON(data []byte) error {
	type Alias AttributeDefinitionDraft
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

type AttributeLocalizedEnumValue struct {
	Key string `json:"key"`
	// JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values are the corresponding strings used for that language.
	Label LocalizedString `json:"label"`
}

type AttributePlainEnumValue struct {
	Key   string `json:"key"`
	Label string `json:"label"`
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
	// [Reference](ctp:api:type:Reference) to a [ProductType](ctp:api:type:ProductType).
	TypeReference ProductTypeReference `json:"typeReference"`
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
	// supported resource type identifiers:
	ReferenceTypeId ReferenceTypeId `json:"referenceTypeId"`
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

type ProductType struct {
	// Platform-generated unique identifier for the ProductType.
	ID string `json:"id"`
	// The current version of the product type.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the ProductType.
	Key         *string               `json:"key,omitempty"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Attributes  []AttributeDefinition `json:"attributes"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductType) MarshalJSON() ([]byte, error) {
	type Alias ProductType
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["attributes"] == nil {
		delete(target, "attributes")
	}

	return json.Marshal(target)
}

type ProductTypeDraft struct {
	// User-defined unique identifier for the ProductType.
	Key         *string                    `json:"key,omitempty"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Attributes  []AttributeDefinitionDraft `json:"attributes"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeDraft) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeDraft
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["attributes"] == nil {
		delete(target, "attributes")
	}

	return json.Marshal(target)
}

type ProductTypePagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int  `json:"limit"`
	Count int  `json:"count"`
	Total *int `json:"total,omitempty"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset  int           `json:"offset"`
	Results []ProductType `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [ProductType](ctp:api:type:ProductType).
*
 */
type ProductTypeReference struct {
	// Platform-generated unique identifier of the referenced [ProductType](ctp:api:type:ProductType).
	ID string `json:"id"`
	// Contains the representation of the expanded ProductType. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for ProductTypes.
	Obj *ProductType `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeReference) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-type", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [ProductType](ctp:api:type:ProductType).
*
 */
type ProductTypeResourceIdentifier struct {
	// Platform-generated unique identifier of the referenced [ProductType](ctp:api:type:ProductType). Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [ProductType](ctp:api:type:ProductType). Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-type", Alias: (*Alias)(&obj)})
}

type ProductTypeUpdate struct {
	Version int                       `json:"version"`
	Actions []ProductTypeUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductTypeUpdate) UnmarshalJSON(data []byte) error {
	type Alias ProductTypeUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorProductTypeUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}
	return nil
}

type ProductTypeUpdateAction interface{}

func mapDiscriminatorProductTypeUpdateAction(input interface{}) (ProductTypeUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "addAttributeDefinition":
		obj := ProductTypeAddAttributeDefinitionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addLocalizedEnumValue":
		obj := ProductTypeAddLocalizedEnumValueAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addPlainEnumValue":
		obj := ProductTypeAddPlainEnumValueAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAttributeConstraint":
		obj := ProductTypeChangeAttributeConstraintAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAttributeName":
		obj := ProductTypeChangeAttributeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAttributeOrder":
		obj := ProductTypeChangeAttributeOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAttributeOrderByName":
		obj := ProductTypeChangeAttributeOrderByNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeDescription":
		obj := ProductTypeChangeDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeEnumKey":
		obj := ProductTypeChangeEnumKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeInputHint":
		obj := ProductTypeChangeInputHintAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeIsSearchable":
		obj := ProductTypeChangeIsSearchableAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeLabel":
		obj := ProductTypeChangeLabelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeLocalizedEnumValueLabel":
		obj := ProductTypeChangeLocalizedEnumValueLabelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeLocalizedEnumValueOrder":
		obj := ProductTypeChangeLocalizedEnumValueOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := ProductTypeChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changePlainEnumValueLabel":
		obj := ProductTypeChangePlainEnumValueLabelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changePlainEnumValueOrder":
		obj := ProductTypeChangePlainEnumValueOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeAttributeDefinition":
		obj := ProductTypeRemoveAttributeDefinitionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeEnumValues":
		obj := ProductTypeRemoveEnumValuesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setInputTip":
		obj := ProductTypeSetInputTipAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := ProductTypeSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type TextInputHint string

const (
	TextInputHintSingleLine TextInputHint = "SingleLine"
	TextInputHintMultiLine  TextInputHint = "MultiLine"
)

type ProductTypeAddAttributeDefinitionAction struct {
	Attribute AttributeDefinitionDraft `json:"attribute"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeAddAttributeDefinitionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeAddAttributeDefinitionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAttributeDefinition", Alias: (*Alias)(&obj)})
}

type ProductTypeAddLocalizedEnumValueAction struct {
	AttributeName string                      `json:"attributeName"`
	Value         AttributeLocalizedEnumValue `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeAddLocalizedEnumValueAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeAddLocalizedEnumValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLocalizedEnumValue", Alias: (*Alias)(&obj)})
}

type ProductTypeAddPlainEnumValueAction struct {
	AttributeName string                  `json:"attributeName"`
	Value         AttributePlainEnumValue `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeAddPlainEnumValueAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeAddPlainEnumValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPlainEnumValue", Alias: (*Alias)(&obj)})
}

type ProductTypeChangeAttributeConstraintAction struct {
	AttributeName string                       `json:"attributeName"`
	NewValue      AttributeConstraintEnumDraft `json:"newValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeChangeAttributeConstraintAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeAttributeConstraintAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAttributeConstraint", Alias: (*Alias)(&obj)})
}

type ProductTypeChangeAttributeNameAction struct {
	AttributeName    string `json:"attributeName"`
	NewAttributeName string `json:"newAttributeName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeChangeAttributeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeAttributeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAttributeName", Alias: (*Alias)(&obj)})
}

type ProductTypeChangeAttributeOrderAction struct {
	Attributes []AttributeDefinition `json:"attributes"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeChangeAttributeOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeAttributeOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAttributeOrder", Alias: (*Alias)(&obj)})
}

type ProductTypeChangeAttributeOrderByNameAction struct {
	AttributeNames []string `json:"attributeNames"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeChangeAttributeOrderByNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeAttributeOrderByNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAttributeOrderByName", Alias: (*Alias)(&obj)})
}

type ProductTypeChangeDescriptionAction struct {
	Description string `json:"description"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeChangeDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeDescription", Alias: (*Alias)(&obj)})
}

type ProductTypeChangeEnumKeyAction struct {
	AttributeName string `json:"attributeName"`
	Key           string `json:"key"`
	NewKey        string `json:"newKey"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeChangeEnumKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeEnumKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEnumKey", Alias: (*Alias)(&obj)})
}

type ProductTypeChangeInputHintAction struct {
	AttributeName string        `json:"attributeName"`
	NewValue      TextInputHint `json:"newValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeChangeInputHintAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeInputHintAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeInputHint", Alias: (*Alias)(&obj)})
}

type ProductTypeChangeIsSearchableAction struct {
	AttributeName string `json:"attributeName"`
	IsSearchable  bool   `json:"isSearchable"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeChangeIsSearchableAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeIsSearchableAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsSearchable", Alias: (*Alias)(&obj)})
}

type ProductTypeChangeLabelAction struct {
	AttributeName string `json:"attributeName"`
	// JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values are the corresponding strings used for that language.
	Label LocalizedString `json:"label"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeChangeLabelAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLabel", Alias: (*Alias)(&obj)})
}

type ProductTypeChangeLocalizedEnumValueLabelAction struct {
	AttributeName string                      `json:"attributeName"`
	NewValue      AttributeLocalizedEnumValue `json:"newValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeChangeLocalizedEnumValueLabelAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeLocalizedEnumValueLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLocalizedEnumValueLabel", Alias: (*Alias)(&obj)})
}

type ProductTypeChangeLocalizedEnumValueOrderAction struct {
	AttributeName string                        `json:"attributeName"`
	Values        []AttributeLocalizedEnumValue `json:"values"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeChangeLocalizedEnumValueOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeLocalizedEnumValueOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLocalizedEnumValueOrder", Alias: (*Alias)(&obj)})
}

type ProductTypeChangeNameAction struct {
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ProductTypeChangePlainEnumValueLabelAction struct {
	AttributeName string                  `json:"attributeName"`
	NewValue      AttributePlainEnumValue `json:"newValue"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeChangePlainEnumValueLabelAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangePlainEnumValueLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePlainEnumValueLabel", Alias: (*Alias)(&obj)})
}

type ProductTypeChangePlainEnumValueOrderAction struct {
	AttributeName string                    `json:"attributeName"`
	Values        []AttributePlainEnumValue `json:"values"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeChangePlainEnumValueOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangePlainEnumValueOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePlainEnumValueOrder", Alias: (*Alias)(&obj)})
}

type ProductTypeRemoveAttributeDefinitionAction struct {
	// The name of the attribute to remove.
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeRemoveAttributeDefinitionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeRemoveAttributeDefinitionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAttributeDefinition", Alias: (*Alias)(&obj)})
}

type ProductTypeRemoveEnumValuesAction struct {
	AttributeName string   `json:"attributeName"`
	Keys          []string `json:"keys"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeRemoveEnumValuesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeRemoveEnumValuesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeEnumValues", Alias: (*Alias)(&obj)})
}

type ProductTypeSetInputTipAction struct {
	AttributeName string `json:"attributeName"`
	// JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values are the corresponding strings used for that language.
	InputTip *LocalizedString `json:"inputTip,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeSetInputTipAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeSetInputTipAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setInputTip", Alias: (*Alias)(&obj)})
}

type ProductTypeSetKeyAction struct {
	// If `key` is absent or `null`, this field will be removed if it exists.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTypeSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}
