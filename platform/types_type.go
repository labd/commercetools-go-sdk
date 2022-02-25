package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	Defines an allowed value of a [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType) field.
*
 */
type CustomFieldEnumValue struct {
	// Key of the value used as a programmatic identifier.
	Key string `json:"key"`
	// Descriptive label of the value.
	Label string `json:"label"`
}

/**
*	Defines an allowed value of a [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) field.
*
 */
type CustomFieldLocalizedEnumValue struct {
	// Key of the value used as a programmatic identifier.
	Key string `json:"key"`
	// Descriptive localized label of the value.
	Label LocalizedString `json:"label"`
}

/**
*	Defines which resource type a [CustomFieldReferenceType](ctp:api:type:CustomFieldReferenceType) can reference.
*
 */
type CustomFieldReferenceValue string

const (
	CustomFieldReferenceValueCart             CustomFieldReferenceValue = "cart"
	CustomFieldReferenceValueCategory         CustomFieldReferenceValue = "category"
	CustomFieldReferenceValueChannel          CustomFieldReferenceValue = "channel"
	CustomFieldReferenceValueCustomer         CustomFieldReferenceValue = "customer"
	CustomFieldReferenceValueKeyValueDocument CustomFieldReferenceValue = "key-value-document"
	CustomFieldReferenceValueOrder            CustomFieldReferenceValue = "order"
	CustomFieldReferenceValueProduct          CustomFieldReferenceValue = "product"
	CustomFieldReferenceValueProductType      CustomFieldReferenceValue = "product-type"
	CustomFieldReferenceValueReview           CustomFieldReferenceValue = "review"
	CustomFieldReferenceValueState            CustomFieldReferenceValue = "state"
	CustomFieldReferenceValueShippingMethod   CustomFieldReferenceValue = "shipping-method"
	CustomFieldReferenceValueZone             CustomFieldReferenceValue = "zone"
)

/**
*	Serves as value of the `custom` field on a resource or data type customized with a [Type](ctp:api:type:Type).
*
 */
type CustomFields struct {
	// Reference to the [Type](ctp:api:type:Type) that holds the [FieldDefinitions](ctp:api:type:FieldDefinition) for the Custom Fields.
	Type TypeReference `json:"type"`
	// Object containing the Custom Fields for the [customized resource or data type](/../api/projects/types#list-of-customizable-data-types).
	Fields FieldContainer `json:"fields"`
}

/**
*	The representation used when creating or updating a [customizable data type](/../api/projects/types#list-of-customizable-data-types) with Custom Fields.
*
 */
type CustomFieldsDraft struct {
	// `id` or `key` of the [Type](ctp:api:type:Type).
	Type TypeResourceIdentifier `json:"type"`
	// Object containing the Custom Fields for the [customized resource or data type](/../api/projects/types#list-of-customizable-data-types).
	Fields *FieldContainer `json:"fields,omitempty"`
}

type FieldContainer map[string]interface{}

/**
*	Defines a [Custom Field](/../api/projects/custom-fields) and its meta-information.
*	This FieldDefinition is similar to an [AttributeDefinition](ctp:api:type:AttributeDefinition) of [Product Types](/../api/projects/productTypes).
*
 */
type FieldDefinition struct {
	// Data type of the Custom Field to define.
	Type FieldType `json:"type"`
	// Name of the Custom Field to define.
	// Must be unique for a given [ResourceTypeId](ctp:api:type:ResourceTypeId).
	// In case there is a FieldDefinition with the same `name` in another [Type](ctp:api:type:Type), both FieldDefinitions must have the same `type`.
	Name string `json:"name"`
	// A human-readable label for the field.
	Label LocalizedString `json:"label"`
	// Defines whether the field is required to have a value.
	Required bool `json:"required"`
	// Must be either `SingleLine` or `MultiLine`.
	// Defines the visual representation of the field in user interfaces like the Merchant Center.
	// It is only relevant for string-based [FieldTypes](ctp:api:type:FieldType) like [CustomFieldStringType](ctp:api:type:CustomFieldStringType) and [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType).
	InputHint *TypeTextInputHint `json:"inputHint,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *FieldDefinition) UnmarshalJSON(data []byte) error {
	type Alias FieldDefinition
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Type != nil {
		var err error
		obj.Type, err = mapDiscriminatorFieldType(obj.Type)
		if err != nil {
			return err
		}
	}
	return nil
}

type FieldType interface{}

func mapDiscriminatorFieldType(input interface{}) (FieldType, error) {
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
	case "Boolean":
		obj := CustomFieldBooleanType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "DateTime":
		obj := CustomFieldDateTimeType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Date":
		obj := CustomFieldDateType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Enum":
		obj := CustomFieldEnumType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "LocalizedEnum":
		obj := CustomFieldLocalizedEnumType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "LocalizedString":
		obj := CustomFieldLocalizedStringType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Money":
		obj := CustomFieldMoneyType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Number":
		obj := CustomFieldNumberType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Reference":
		obj := CustomFieldReferenceType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Set":
		obj := CustomFieldSetType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.ElementType != nil {
			var err error
			obj.ElementType, err = mapDiscriminatorFieldType(obj.ElementType)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "String":
		obj := CustomFieldStringType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Time":
		obj := CustomFieldTimeType{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Field type for Boolean values.
*
 */
type CustomFieldBooleanType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomFieldBooleanType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldBooleanType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "Boolean", Alias: (*Alias)(&obj)})
}

/**
*	Field type for [DateTime](ctp:api:type:DateTime) values.
*
 */
type CustomFieldDateTimeType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomFieldDateTimeType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldDateTimeType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "DateTime", Alias: (*Alias)(&obj)})
}

/**
*	Field type for [Date](ctp:api:type:Date) values.
*
 */
type CustomFieldDateType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomFieldDateType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldDateType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "Date", Alias: (*Alias)(&obj)})
}

/**
*	Field type for enum values.
*
 */
type CustomFieldEnumType struct {
	// Allowed values.
	Values []CustomFieldEnumValue `json:"values"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomFieldEnumType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldEnumType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "Enum", Alias: (*Alias)(&obj)})
}

/**
*	Field type for localized enum values.
*
 */
type CustomFieldLocalizedEnumType struct {
	// Allowed values.
	Values []CustomFieldLocalizedEnumValue `json:"values"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomFieldLocalizedEnumType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldLocalizedEnumType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "LocalizedEnum", Alias: (*Alias)(&obj)})
}

/**
*	Field type for [LocalizedString](ctp:api:type:LocalizedString) values.
*
 */
type CustomFieldLocalizedStringType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomFieldLocalizedStringType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldLocalizedStringType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "LocalizedString", Alias: (*Alias)(&obj)})
}

/**
*	Field type for [CentPrecisionMoney](ctp:api:type:CentPrecisionMoney) values.
*
 */
type CustomFieldMoneyType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomFieldMoneyType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldMoneyType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "Money", Alias: (*Alias)(&obj)})
}

/**
*	Field type for number values.
*
 */
type CustomFieldNumberType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomFieldNumberType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldNumberType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "Number", Alias: (*Alias)(&obj)})
}

/**
*	Field type for [Reference](ctp:api:type:Reference) values.
*
 */
type CustomFieldReferenceType struct {
	// Resource type the Custom Field can reference.
	ReferenceTypeId CustomFieldReferenceValue `json:"referenceTypeId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomFieldReferenceType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldReferenceType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "Reference", Alias: (*Alias)(&obj)})
}

/**
*	Values of a SetType Custom Field are sets of values of the specified `elementType`.
*
 */
type CustomFieldSetType struct {
	// Field type of the elements in the set.
	ElementType FieldType `json:"elementType"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomFieldSetType) UnmarshalJSON(data []byte) error {
	type Alias CustomFieldSetType
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ElementType != nil {
		var err error
		obj.ElementType, err = mapDiscriminatorFieldType(obj.ElementType)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomFieldSetType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldSetType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "Set", Alias: (*Alias)(&obj)})
}

/**
*	Field type for string values.
*
 */
type CustomFieldStringType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomFieldStringType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldStringType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "String", Alias: (*Alias)(&obj)})
}

/**
*	Field type for [Time](ctp:api:type:Time) values.
*
 */
type CustomFieldTimeType struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomFieldTimeType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldTimeType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "Time", Alias: (*Alias)(&obj)})
}

/**
*	IDs indicating the [customizable resources and data types](/../api/projects/types#list-of-customizable-data-types).
*
 */
type ResourceTypeId string

const (
	ResourceTypeIdAddress                     ResourceTypeId = "address"
	ResourceTypeIdAsset                       ResourceTypeId = "asset"
	ResourceTypeIdCartDiscount                ResourceTypeId = "cart-discount"
	ResourceTypeIdCategory                    ResourceTypeId = "category"
	ResourceTypeIdChannel                     ResourceTypeId = "channel"
	ResourceTypeIdCustomer                    ResourceTypeId = "customer"
	ResourceTypeIdCustomerGroup               ResourceTypeId = "customer-group"
	ResourceTypeIdCustomLineItem              ResourceTypeId = "custom-line-item"
	ResourceTypeIdDiscountCode                ResourceTypeId = "discount-code"
	ResourceTypeIdInventoryEntry              ResourceTypeId = "inventory-entry"
	ResourceTypeIdLineItem                    ResourceTypeId = "line-item"
	ResourceTypeIdOrder                       ResourceTypeId = "order"
	ResourceTypeIdOrderEdit                   ResourceTypeId = "order-edit"
	ResourceTypeIdOrderDelivery               ResourceTypeId = "order-delivery"
	ResourceTypeIdPayment                     ResourceTypeId = "payment"
	ResourceTypeIdPaymentInterfaceInteraction ResourceTypeId = "payment-interface-interaction"
	ResourceTypeIdProductPrice                ResourceTypeId = "product-price"
	ResourceTypeIdProductSelection            ResourceTypeId = "product-selection"
	ResourceTypeIdReview                      ResourceTypeId = "review"
	ResourceTypeIdShippingMethod              ResourceTypeId = "shipping-method"
	ResourceTypeIdShoppingList                ResourceTypeId = "shopping-list"
	ResourceTypeIdShoppingListTextLineItem    ResourceTypeId = "shopping-list-text-line-item"
	ResourceTypeIdStore                       ResourceTypeId = "store"
	ResourceTypeIdTransaction                 ResourceTypeId = "transaction"
)

type Type struct {
	// Unique ID of the Type.
	ID string `json:"id"`
	// Current version of the Type.
	Version int `json:"version"`
	// Date and time (UTC) the Type was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Type was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier for the Type.
	Key string `json:"key"`
	// Name of the Type.
	Name LocalizedString `json:"name"`
	// Description of the Type.
	Description *LocalizedString `json:"description,omitempty"`
	// Resources and/or data types for which the Type is defined.
	ResourceTypeIds []ResourceTypeId `json:"resourceTypeIds"`
	// Defines Custom Fields.
	FieldDefinitions []FieldDefinition `json:"fieldDefinitions"`
}

type TypeDraft struct {
	// User-defined unique identifier for the Type.
	Key string `json:"key"`
	// Name of the Type.
	Name LocalizedString `json:"name"`
	// Description of the Type.
	Description *LocalizedString `json:"description,omitempty"`
	// Resources and/or data types for which the Type is defined.
	ResourceTypeIds []ResourceTypeId `json:"resourceTypeIds"`
	// Defines Custom Fields.
	FieldDefinitions []FieldDefinition `json:"fieldDefinitions"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeDraft) MarshalJSON() ([]byte, error) {
	type Alias TypeDraft
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

	if target["fieldDefinitions"] == nil {
		delete(target, "fieldDefinitions")
	}

	return json.Marshal(target)
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [Types](ctp:api:type:Type).
*
 */
type TypePagedQueryResponse struct {
	// Number of results requested in the query request.
	Limit int `json:"limit"`
	// Offset supplied by the client or server default. It is the number of elements skipped, not a page number.
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](ctp:api:type:QueryPredicate), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [Types](ctp:api:type:Type) matching the query.
	Results []Type `json:"results"`
}

/**
*	[Reference](/../api/types#reference) to a [Type](ctp:api:type:Type).
*
 */
type TypeReference struct {
	// Unique ID of the referenced [Type](ctp:api:type:Type).
	ID string `json:"id"`
	// Contains the representation of the expanded Type.
	// Only present in responses to requests with [Reference Expansion](ctp:api:type:Expansion) for Types.
	Obj *Type `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeReference) MarshalJSON() ([]byte, error) {
	type Alias TypeReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "type", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](/../api/types#resourceidentifier) of a [Type](ctp:api:type:Type).
*
 */
type TypeResourceIdentifier struct {
	// Unique ID of the referenced [Type](ctp:api:type:Type). Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced [Type](ctp:api:type:Type). Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias TypeResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "type", Alias: (*Alias)(&obj)})
}

/**
*	Provides a visual representation type for this field. It is only relevant for string-based field types like [CustomFieldStringType](ctp:api:type:CustomFieldStringType) and [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType). Following values are supported:
*
 */
type TypeTextInputHint string

const (
	TypeTextInputHintSingleLine TypeTextInputHint = "SingleLine"
	TypeTextInputHintMultiLine  TypeTextInputHint = "MultiLine"
)

type TypeUpdate struct {
	// Expected version of the type on which the changes should be applied.
	// If the expected version does not match the actual version, a 409 Conflict will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the Type.
	Actions []TypeUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *TypeUpdate) UnmarshalJSON(data []byte) error {
	type Alias TypeUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorTypeUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}
	return nil
}

type TypeUpdateAction interface{}

func mapDiscriminatorTypeUpdateAction(input interface{}) (TypeUpdateAction, error) {
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
	case "addEnumValue":
		obj := TypeAddEnumValueAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addFieldDefinition":
		obj := TypeAddFieldDefinitionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addLocalizedEnumValue":
		obj := TypeAddLocalizedEnumValueAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeEnumValueLabel":
		obj := TypeChangeEnumValueLabelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeEnumValueOrder":
		obj := TypeChangeEnumValueOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeFieldDefinitionLabel":
		obj := TypeChangeFieldDefinitionLabelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeFieldDefinitionOrder":
		obj := TypeChangeFieldDefinitionOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeInputHint":
		obj := TypeChangeInputHintAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeKey":
		obj := TypeChangeKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeLabel":
		obj := TypeChangeLabelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeLocalizedEnumValueLabel":
		obj := TypeChangeLocalizedEnumValueLabelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeLocalizedEnumValueOrder":
		obj := TypeChangeLocalizedEnumValueOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := TypeChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeFieldDefinition":
		obj := TypeRemoveFieldDefinitionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := TypeSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Adds a value to an [EnumType](ctp:api:type:CustomFieldEnumType).
*	This update action can be used to update an [EnumType](ctp:api:type:CustomFieldEnumType) FieldDefinition and a [SetType](ctp:api:type:CustomFieldSetType) FieldDefinition of [EnumType](ctp:api:type:CustomFieldEnumType).
*
 */
type TypeAddEnumValueAction struct {
	// `name` of the [Field Definition](ctp:api:type:FieldDefinition) to update.
	FieldName string `json:"fieldName"`
	// Value to append to the array.
	Value CustomFieldEnumValue `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeAddEnumValueAction) MarshalJSON() ([]byte, error) {
	type Alias TypeAddEnumValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addEnumValue", Alias: (*Alias)(&obj)})
}

type TypeAddFieldDefinitionAction struct {
	// Value to append to the array.
	FieldDefinition FieldDefinition `json:"fieldDefinition"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeAddFieldDefinitionAction) MarshalJSON() ([]byte, error) {
	type Alias TypeAddFieldDefinitionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addFieldDefinition", Alias: (*Alias)(&obj)})
}

/**
*	Adds a value to a [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType).
*	This update action can be used to update a [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) FieldDefinition and a [SetType](ctp:api:type:CustomFieldSetType) FieldDefinition of [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType).
*
 */
type TypeAddLocalizedEnumValueAction struct {
	// `name` of the [FieldDefinition](ctp:api:type:FieldDefinition) to update.
	FieldName string `json:"fieldName"`
	// Value to append to the array.
	Value CustomFieldLocalizedEnumValue `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeAddLocalizedEnumValueAction) MarshalJSON() ([]byte, error) {
	type Alias TypeAddLocalizedEnumValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLocalizedEnumValue", Alias: (*Alias)(&obj)})
}

/**
*	Changes the `label` of an [EnumValue](ctp:api:type:CustomFieldEnumValue) of an [EnumType](ctp:api:type:CustomFieldEnumType) FieldDefinition.
*
 */
type TypeChangeEnumValueLabelAction struct {
	// `name` of the [FieldDefinition](ctp:api:type:FieldDefinition) to update.
	FieldName string `json:"fieldName"`
	// New value to set.
	// Must not be empty.
	Value CustomFieldEnumValue `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeChangeEnumValueLabelAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeEnumValueLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEnumValueLabel", Alias: (*Alias)(&obj)})
}

/**
*	Changes the order of [EnumValues](ctp:api:type:CustomFieldEnumValue) in an [EnumType](ctp:api:type:CustomFieldEnumType) FieldDefinition.
*	This update action can be used to update an [EnumType](ctp:api:type:CustomFieldEnumType) FieldDefinition and a [SetType](ctp:api:type:CustomFieldSetType) FieldDefinition of [EnumType](ctp:api:type:CustomFieldEnumType).
*
 */
type TypeChangeEnumValueOrderAction struct {
	// `name` of the [FieldDefinition](ctp:api:type:FieldDefinition) to update.
	FieldName string `json:"fieldName"`
	// Must match the set of `key`s of the EnumValues in the FieldDefinition (apart from their order).
	Keys []string `json:"keys"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeChangeEnumValueOrderAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeEnumValueOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEnumValueOrder", Alias: (*Alias)(&obj)})
}

type TypeChangeFieldDefinitionLabelAction struct {
	// `name` of the [FieldDefinition](ctp:api:type:FieldDefinition) to update.
	FieldName string `json:"fieldName"`
	// New value to set.
	// Must not be empty.
	Label LocalizedString `json:"label"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeChangeFieldDefinitionLabelAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeFieldDefinitionLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeFieldDefinitionLabel", Alias: (*Alias)(&obj)})
}

type TypeChangeFieldDefinitionOrderAction struct {
	// Must match the set of `name`s of FieldDefinitions (up to order).
	FieldNames []string `json:"fieldNames"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeChangeFieldDefinitionOrderAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeFieldDefinitionOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeFieldDefinitionOrder", Alias: (*Alias)(&obj)})
}

/**
*	Changes the `inputHint` of [CustomFieldStringType](ctp:api:type:CustomFieldStringType) [FieldDefinition](ctp:api:type:FieldDefinition), a [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType) [FieldDefinition](ctp:api:type:FieldDefinition), and [CustomFieldSetType](ctp:api:type:CustomFieldSetType) [FieldDefinition](ctp:api:type:FieldDefinition) of these string-based FieldTypes.
*
 */
type TypeChangeInputHintAction struct {
	// `name` of the [Field Definition](ctp:api:type:FieldDefinition) to update.
	FieldName string `json:"fieldName"`
	// New value to set.
	// Must not be empty.
	InputHint TypeTextInputHint `json:"inputHint"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeChangeInputHintAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeInputHintAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeInputHint", Alias: (*Alias)(&obj)})
}

type TypeChangeKeyAction struct {
	// New value to set.
	// Must not be empty.
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeChangeKeyAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeKey", Alias: (*Alias)(&obj)})
}

type TypeChangeLabelAction struct {
	// Name of the [Field Definition](ctp:api:type:FieldDefinition) to update.
	FieldName string          `json:"fieldName"`
	Label     LocalizedString `json:"label"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeChangeLabelAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLabel", Alias: (*Alias)(&obj)})
}

/**
*	Changes the `label` of a [LocalizedEnumValue](ctp:api:type:CustomFieldLocalizedEnumValue) of an [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) FieldDefinition.
*
 */
type TypeChangeLocalizedEnumValueLabelAction struct {
	// `name` of the [FieldDefinition](ctp:api:type:FieldDefinition) to update.
	FieldName string `json:"fieldName"`
	// New value to set.
	// Must not be empty.
	Value CustomFieldLocalizedEnumValue `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeChangeLocalizedEnumValueLabelAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeLocalizedEnumValueLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLocalizedEnumValueLabel", Alias: (*Alias)(&obj)})
}

/**
*	Changes the order of [LocalizedEnumValues](ctp:api:type:CustomFieldLocalizedEnumValue) in a [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) FieldDefinition.
*	This update action can be used to update a [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) FieldDefinition and a [SetType](ctp:api:type:CustomFieldSetType) of [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) FieldDefinitions.
*
 */
type TypeChangeLocalizedEnumValueOrderAction struct {
	// `name` of the [Field Definition](ctp:api:type:FieldDefinition) to update.
	FieldName string `json:"fieldName"`
	// Must match the set of `key`s of the LocalizedEnumValues in the FieldDefinition (up to order).
	Keys []string `json:"keys"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeChangeLocalizedEnumValueOrderAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeLocalizedEnumValueOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLocalizedEnumValueOrder", Alias: (*Alias)(&obj)})
}

type TypeChangeNameAction struct {
	// New value to set.
	// Must not be empty.
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type TypeRemoveFieldDefinitionAction struct {
	// `name` of the [FieldDefinition](ctp:api:type:FieldDefinition) to remove.
	// The removal of a FieldDefinition deletes [asynchronously](/../api/general-concepts#eventual-consistency) all Custom Fields using the FieldDefinition as well.
	FieldName string `json:"fieldName"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeRemoveFieldDefinitionAction) MarshalJSON() ([]byte, error) {
	type Alias TypeRemoveFieldDefinitionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeFieldDefinition", Alias: (*Alias)(&obj)})
}

type TypeSetDescriptionAction struct {
	// Value to set. If empty, any existing value will be removed.
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias TypeSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}
