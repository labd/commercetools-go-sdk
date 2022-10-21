package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	Specifies how an Attribute (or a set of Attributes) should be validated across all variants of a Product:
*
 */
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

/**
*	Describes a Product Attribute and allows you to define meta-information associated with the Attribute (like whether it should be searchable, or its constraints).
*
 */
type AttributeDefinition struct {
	// Describes the Type of the Attribute.
	Type AttributeType `json:"type"`
	// User-defined name of the Attribute that is unique within the [Project](ctp:api:type:Project).
	Name string `json:"name"`
	// Human-readable label for the Attribute.
	Label LocalizedString `json:"label"`
	// If `true`, the Attribute must have a value on a [ProductVariant](ctp:api:type:ProductVariant).
	IsRequired bool `json:"isRequired"`
	// Specifies how Attributes are validated across all variants of a Product.
	AttributeConstraint AttributeConstraintEnum `json:"attributeConstraint"`
	// Provides additional Attribute information to aid content managers configure Product details.
	InputTip *LocalizedString `json:"inputTip,omitempty"`
	// Provides a visual representation directive for values of this Attribute (only relevant for [AttributeTextType](ctp:api:type:AttributeTextType) and [AttributeLocalizableTextType](ctp:api:type:AttributeLocalizableTextType)).
	InputHint TextInputHint `json:"inputHint"`
	// If `true`, the Attribute's values are available for the [Product Projections Search API](/../api/projects/products-search) for use in full-text search queries, filters, and facets.
	//
	// Which exact features are available with this flag depends on the specific [AttributeType](ctp:api:type:AttributeType).
	// The maximum size of a searchable field is **restricted** by the [Field content size limit](/../api/limits#field-content-size).
	// This constraint is enforced at both [Product creation](/../api/projects/products#create-product) and [Product update](/../api/projects/products#update-product).
	// If the length of the input exceeds the maximum size, an [InvalidFieldError](ctp:api:type:InvalidFieldError) is returned.
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

/**
*	Specify the Attribute to be created with the [ProductTypeDraft](ctp:api:type:ProductTypeDraft).
*
 */
type AttributeDefinitionDraft struct {
	// Describes the Type of the Attribute.
	Type AttributeType `json:"type"`
	// User-defined name of the Attribute that is unique with the [Project](ctp:api:type:Project).
	// When using the same `name` for an Attribute in multiple ProductTypes, all fields of the AttributeDefinition of this Attribute must be the same across the ProductTypes. Otherwise an [AttributeDefinitionAlreadyExistsError](ctp:api:type:AttributeDefinitionAlreadyExistsError) will be returned.
	// An exception to this are the values of an `enum` or `lenum` Type and sets thereof.
	Name string `json:"name"`
	// Human-readable label for the Attribute.
	Label LocalizedString `json:"label"`
	// Set to `true` if the Attribute is required to have a value on a [ProductVariant](ctp:api:type:ProductVariant).
	IsRequired bool `json:"isRequired"`
	// Specifies how an Attribute or a combination of Attributes should be validated across all variants of a Product.
	AttributeConstraint *AttributeConstraintEnum `json:"attributeConstraint,omitempty"`
	// Provides additional information about the Attribute that aids content managers when setting Product details.
	InputTip *LocalizedString `json:"inputTip,omitempty"`
	// Provides a visual representation directive for values of this Attribute (only relevant for [AttributeTextType](ctp:api:type:AttributeTextType) and [AttributeLocalizableTextType](ctp:api:type:AttributeLocalizableTextType)).
	InputHint *TextInputHint `json:"inputHint,omitempty"`
	// Set to `true` if the Attribute's values should be available in the [Product Projections Search API](/../api/projects/products-search) and can be used in full-text search queries, filters, and facets.
	// Which exact features are available with this flag depends on the specific [AttributeType](ctp:api:type:AttributeType).
	// The maximum size of a searchable field is **restricted** by the [Field content size limit](/../api/limits#field-content-size).
	// This constraint is enforced at both Product creation and Product update.
	// If the length of the input exceeds the maximum size, an InvalidField error is returned.
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

/**
*	Attribute type for localized enum values. Useful for predefined language-specific values selectable in drop-down menus if only one value can be selected. Use [AttributeSetType](ctp:api:type:AttributeSetType) of AttributeLocalizedEnumValue instead if multiple values can be selected.
*
 */
type AttributeLocalizedEnumValue struct {
	// Key of the value used as a programmatic identifier, for example in facets & filters.
	Key string `json:"key"`
	// Descriptive, localized label of the value.
	Label LocalizedString `json:"label"`
}

/**
*	A plain enum value must be unique within the enum, otherwise a [DuplicateEnumValues](/errors#product-types-400-duplicate-enum-values) error will be returned.
*
 */
type AttributePlainEnumValue struct {
	// Key of the value used as a programmatic identifier, for example in facets & filters.
	Key string `json:"key"`
	// Descriptive label of the value.
	Label string `json:"label"`
}

/**
*	Name of the resource type that the value should reference. Supported resource type identifiers:
*
 */
type AttributeReferenceTypeId string

const (
	AttributeReferenceTypeIdCart             AttributeReferenceTypeId = "cart"
	AttributeReferenceTypeIdCategory         AttributeReferenceTypeId = "category"
	AttributeReferenceTypeIdChannel          AttributeReferenceTypeId = "channel"
	AttributeReferenceTypeIdCustomer         AttributeReferenceTypeId = "customer"
	AttributeReferenceTypeIdKeyValueDocument AttributeReferenceTypeId = "key-value-document"
	AttributeReferenceTypeIdOrder            AttributeReferenceTypeId = "order"
	AttributeReferenceTypeIdProduct          AttributeReferenceTypeId = "product"
	AttributeReferenceTypeIdProductType      AttributeReferenceTypeId = "product-type"
	AttributeReferenceTypeIdReview           AttributeReferenceTypeId = "review"
	AttributeReferenceTypeIdShippingMethod   AttributeReferenceTypeId = "shipping-method"
	AttributeReferenceTypeIdState            AttributeReferenceTypeId = "state"
	AttributeReferenceTypeIdZone             AttributeReferenceTypeId = "zone"
)

/**
*	Umbrella type for specific attribute types discriminated by property `name`.
 */
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

/**
*	Attribute type for Boolean values. Valid values for the Attribute are `true` and `false` (JSON Boolean).
*
 */
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

/**
*	Attribute type for plain enum values. Useful for predefined language-agnostic values selectable in drop downs when only one value should be selected. Use [AttributeSetType](ctp:api:type:AttributeSetType) of AttributeEnumType instead if multiple values can be selected from the list.
*
 */
type AttributeEnumType struct {
	// Available values that can be assigned to Products.
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

/**
*	Attribute type for [LocalizedString](ctp:api:type:LocalizedString) values.
*
 */
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
	// Available values that can be assigned to Products.
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

/**
*	Attribute type for nesting Attributes based on some existing ProductType. It does not support `isSearchable` and is not supported in queries. The only supported AttributeConstraint is `None`.
*
 */
type AttributeNestedType struct {
	// Attributes that can be stored as nested Attributes of the current Attribute.
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
	// Name of the resource type that the value should reference.
	ReferenceTypeId AttributeReferenceTypeId `json:"referenceTypeId"`
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

/**
*	AttributeType that defines a set (without duplicate elements) with values of the given `elementType`. It does not support `isRequired`. Since this type itself is an AttributeType, it is possible to construct an AttributeSetType of an AttributeSetType of any AttributeType, and to continue with this iteration until terminating with any non-AttributeSetType. In case the AttributeSetType iteration terminates with an [AttributeNestedType](ctp:api:type:AttributeNestedType), the iteration can have 5 steps at maximum.
*
 */
type AttributeSetType struct {
	// Attribute type of the elements in the set.
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

/**
*	Attribute type for plain text values.
*
 */
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
	// Unique identifier of the ProductType.
	ID string `json:"id"`
	// Current version of the ProductType.
	Version int `json:"version"`
	// Date and time (UTC) the ProductType was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the ProductType was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the ProductType.
	Key *string `json:"key,omitempty"`
	// Name of the ProductType.
	Name string `json:"name"`
	// Description of the ProductType.
	Description string `json:"description"`
	// Attributes specified for the ProductType.
	Attributes []AttributeDefinition `json:"attributes"`
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["attributes"] == nil {
		delete(raw, "attributes")
	}

	return json.Marshal(raw)

}

type ProductTypeDraft struct {
	// User-defined unique identifier for the ProductType.
	Key *string `json:"key,omitempty"`
	// Name of the ProductType.
	Name string `json:"name"`
	// Description of the ProductType.
	Description string `json:"description"`
	// Attributes to specify for the ProductType. Products of this ProductType have these Attributes available on their [ProductVariants](ctp:api:type:ProductVariant).
	Attributes []AttributeDefinitionDraft `json:"attributes"`
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["attributes"] == nil {
		delete(raw, "attributes")
	}

	return json.Marshal(raw)

}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [ProductType](ctp:api:type:ProductType).
*
 */
type ProductTypePagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [ProductTypes](ctp:api:type:ProductType) matching the query.
	Results []ProductType `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [ProductType](ctp:api:type:ProductType).
*
 */
type ProductTypeReference struct {
	// Unique identifier of the referenced [ProductType](ctp:api:type:ProductType).
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
	// Unique identifier of the referenced [ProductType](ctp:api:type:ProductType). Either `id` or `key` is required.
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
	// Expected version of the ProductType on which the changes should be applied. If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the ProductType.
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

/**
*	A text input hint is a string with one of the following values:
*
 */
type TextInputHint string

const (
	TextInputHintSingleLine TextInputHint = "SingleLine"
	TextInputHintMultiLine  TextInputHint = "MultiLine"
)

type ProductTypeAddAttributeDefinitionAction struct {
	// Value to append to `attributes`.
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

/**
*	Adds a localizable enum to the values of [AttributeLocalizedEnumType](ctp:api:type:AttributeLocalizedEnumType). It can update an AttributeLocalizedEnumType AttributeDefinition or an [AttributeSetType](ctp:api:type:AttributeSetType) of AttributeLocalizedEnumType AttributeDefinition.
*
 */
type ProductTypeAddLocalizedEnumValueAction struct {
	// Name of the AttributeDefinition to update.
	AttributeName string `json:"attributeName"`
	// Value to append to the array.
	Value AttributeLocalizedEnumValue `json:"value"`
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

/**
*	Adds an enum to the values of [AttributeEnumType](ctp:api:type:AttributeEnumType) AttributeDefinition, or [AttributeSetType](ctp:api:type:AttributeSetType) of AttributeEnumType AttributeDefinition.
*
 */
type ProductTypeAddPlainEnumValueAction struct {
	// Name of the AttributeDefinition to update.
	AttributeName string `json:"attributeName"`
	// Value to append to the array.
	Value AttributePlainEnumValue `json:"value"`
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

/**
*	Updates the `attributeConstraint` of an [AttributeDefinition](ctp:api:type:AttributeDefinition). For now only following changes are supported: `SameForAll` to `None` and `Unique` to `None`.
*
 */
type ProductTypeChangeAttributeConstraintAction struct {
	// Name of the AttributeDefinition to update.
	AttributeName string `json:"attributeName"`
	// `None`
	NewValue AttributeConstraintEnumDraft `json:"newValue"`
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

/**
*	Renames an AttributeDefinition and also renames all corresponding Attributes on all [Products](/projects/products) with this ProductType. The renaming of the Attributes is [eventually consistent](/general-concepts#eventual-consistency).
*
 */
type ProductTypeChangeAttributeNameAction struct {
	// Name of the AttributeDefinition to update.
	AttributeName string `json:"attributeName"`
	// New user-defined name of the Attribute that is unique with the [Project](ctp:api:type:Project).
	// When using the same `name` for an Attribute in two or more ProductTypes all fields of the AttributeDefinition of this Attribute need to be the same across the ProductTypes, otherwise an [AttributeDefinitionAlreadyExistsError](ctp:api:type:AttributeDefinitionAlreadyExistsError) will be returned.
	// An exception to this are the values of an `enum` or `lenum` type and sets thereof.
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
	// Names of Attributes to reorder. This array must include all Attributes currently present on a ProductType in a different order.
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
	// New value to set.
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

/**
*	Updates the key of a single enum `value` in an [AttributeEnumType](ctp:api:type:AttributeEnumType) AttributeDefinition, [AttributeLocalizedEnumType](ctp:api:type:AttributeLocalizedEnumType) AttributeDefinition, [AttributeSetType](ctp:api:type:AttributeSetType) of AttributeEnumType AttributeDefinition, or AttributeSetType of AttributeLocalizedEnumType AttributeDefinition.
*
*	All Products will be updated to the new key in an [eventually consistent](/general-concepts#eventual-consistency) way.
*
 */
type ProductTypeChangeEnumKeyAction struct {
	// Name of the AttributeDefinition to update.
	AttributeName string `json:"attributeName"`
	// Existing key to be changed.
	Key string `json:"key"`
	// New key to be set.
	NewKey string `json:"newKey"`
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

/**
*	Updates the `inputHint` of an [AttributeDefinition](ctp:api:type:AttributeDefinition).
*
 */
type ProductTypeChangeInputHintAction struct {
	// Name of the AttributeDefinition to update.
	AttributeName string `json:"attributeName"`
	// `SingleLine` or `MultiLine`
	NewValue TextInputHint `json:"newValue"`
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

/**
*	Following this update the Products are reindexed asynchronously to reflect this change on the search endpoint. When enabling search on an existing Attribute type definition, the constraint regarding the maximum size of a searchable Attribute will not be enforced. Instead, Product AttributeDefinitions exceeding this limit will be treated as not searchable and will not be available for full-text search.
*
 */
type ProductTypeChangeIsSearchableAction struct {
	// Name of the AttributeDefinition to update.
	AttributeName string `json:"attributeName"`
	// Determines whether the Attribute's values can be used in full-text search queries, filters, and facets. See [AttributeDefinition](ctp:api:type:AttributeDefinition) for details.
	IsSearchable bool `json:"isSearchable"`
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
	// Name of the AttributeDefinition to update.
	AttributeName string `json:"attributeName"`
	// New value to set. Must not be empty.
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

/**
*	Updates the label of a single enum `value` in an [AttributeLocalizedEnumType](ctp:api:type:AttributeLocalizedEnumType) AttributeDefinition, or [AttributeSetType](ctp:api:type:AttributeSetType) of AttributeLocalizedEnumType AttributeDefinition.
*
*	All Products will be updated to the new label in an [eventually consistent](/general-concepts#eventual-consistency) way.
*
 */
type ProductTypeChangeLocalizedEnumValueLabelAction struct {
	// Name of the AttributeDefinition to update.
	AttributeName string `json:"attributeName"`
	// New value to set. Must be different from the existing value.
	NewValue AttributeLocalizedEnumValue `json:"newValue"`
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

/**
*	Updates the order of localized enum `values` in an [AttributeLocalizedEnumType](ctp:api:type:AttributeLocalizedEnumType) AttributeDefinition. It can update an AttributeLocalizedEnumType AttributeDefinition or an [AttributeSetType](ctp:api:type:AttributeSetType) of AttributeLocalizedEnumType AttributeDefinition.
*
 */
type ProductTypeChangeLocalizedEnumValueOrderAction struct {
	// Name of the AttributeDefinition to update.
	AttributeName string `json:"attributeName"`
	// Values must be equal to the values of the Attribute enum values (except for the order). If not, an [EnumValuesMustMatch](/errors#product-types-400-enum-values-must-match) error code will be returned.
	Values []AttributeLocalizedEnumValue `json:"values"`
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
	// New value to set.
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

/**
*	Updates the label of a single enum `value` in an [AttributeEnumType](ctp:api:type:AttributeEnumType) AttributeDefinition, or [AttributeSetType](ctp:api:type:AttributeSetType) of AttributeEnumType AttributeDefinition.
*
*	All Products will be updated to the new label in an [eventually consistent](/general-concepts#eventual-consistency) way.
*
 */
type ProductTypeChangePlainEnumValueLabelAction struct {
	// Name of the AttributeDefinition to update.
	AttributeName string `json:"attributeName"`
	// New value to set. Must be different from the existing value.
	NewValue AttributePlainEnumValue `json:"newValue"`
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

/**
*	Updates the order of enum `values` in an [AttributeEnumType](ctp:api:type:AttributeEnumType) AttributeDefinition. It can update an AttributeEnumType AttributeDefinition or an [AttributeSetType](ctp:api:type:AttributeSetType) of AttributeEnumType AttributeDefinition.
*
 */
type ProductTypeChangePlainEnumValueOrderAction struct {
	// Name of the AttributeDefinition to update.
	AttributeName string `json:"attributeName"`
	// Values must be equal to the values of the Attribute enum values (except for the order). If not, an [EnumValuesMustMatch](/errors#product-types-400-enum-values-must-match) error code will be returned.
	Values []AttributePlainEnumValue `json:"values"`
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

/**
*	Removes an AttributeDefinition and also deletes all corresponding Attributes on all [Products](/projects/products) with this ProductType. The removal of the Attributes is [eventually consistent](/general-concepts#eventual-consistency).
*
*	The `CombinationUnique` constraint is not checked when an Attribute is removed, and uniqueness violations may occur when you remove an Attribute with a `CombinationUnique` constraint.
*
 */
type ProductTypeRemoveAttributeDefinitionAction struct {
	// Name of the Attribute to remove.
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

/**
*	Removes enum values from an AttributeDefinition of [AttributeEnumType](ctp:api:type:AttributeEnumType), [AttributeLocalizedEnumType](ctp:api:type:AttributeLocalizedEnumType), [AttributeSetType](ctp:api:type:AttributeSetType) of AttributeEnumType, or AttributeSetType of AttributeLocalizedEnumType.
*
*	If the Attribute is **not** required, the Attributes of all Products using those enum keys will also be removed in an [eventually consistent](/general-concepts#eventual-consistency) way. If the Attribute is required, the operation will fail with the [EnumValueIsUsed](/errors#product-types-400-enum-value-is-used) error code.
*
 */
type ProductTypeRemoveEnumValuesAction struct {
	// Name of the AttributeDefinition to update.
	AttributeName string `json:"attributeName"`
	// Keys of [AttributeEnumType](ctp:api:type:AttributeEnumType) or [AttributeLocalizedEnumType](ctp:api:type:AttributeLocalizedEnumType) to remove.
	Keys []string `json:"keys"`
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
	// Name of the AttributeDefinition to update.
	AttributeName string `json:"attributeName"`
	// Value to set. If empty, any existing value will be removed.
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
	// Value to set. If empty, any existing value will be removed.
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
