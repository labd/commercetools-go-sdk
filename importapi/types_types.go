package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

/**
*	Provides a visual representation type for this field. It is only relevant for string-based field types like [CustomFieldStringType](ctp:import:type:CustomFieldStringType) and [CustomFieldLocalizedStringType](ctp:import:type:CustomFieldLocalizedStringType).
*
 */
type TypeTextInputHint string

const (
	TypeTextInputHintSingleLine TypeTextInputHint = "SingleLine"
	TypeTextInputHintMultiLine  TypeTextInputHint = "MultiLine"
)

/**
*	IDs indicating the [customizable resources and data types](/../api/projects/types#list-of-customizable-data-types). Maps to `Type.resourceTypeId`.
*
 */
type ResourceTypeId string

const (
	ResourceTypeIdAddress                     ResourceTypeId = "address"
	ResourceTypeIdAsset                       ResourceTypeId = "asset"
	ResourceTypeIdBusinessUnit                ResourceTypeId = "business-unit"
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
	ResourceTypeIdOrderParcel                 ResourceTypeId = "order-parcel"
	ResourceTypeIdOrderReturnItem             ResourceTypeId = "order-return-item"
	ResourceTypeIdPayment                     ResourceTypeId = "payment"
	ResourceTypeIdPaymentInterfaceInteraction ResourceTypeId = "payment-interface-interaction"
	ResourceTypeIdProductPrice                ResourceTypeId = "product-price"
	ResourceTypeIdProductSelection            ResourceTypeId = "product-selection"
	ResourceTypeIdQuote                       ResourceTypeId = "quote"
	ResourceTypeIdReview                      ResourceTypeId = "review"
	ResourceTypeIdShipping                    ResourceTypeId = "shipping"
	ResourceTypeIdShippingMethod              ResourceTypeId = "shipping-method"
	ResourceTypeIdShoppingList                ResourceTypeId = "shopping-list"
	ResourceTypeIdShoppingListTextLineItem    ResourceTypeId = "shopping-list-text-line-item"
	ResourceTypeIdStandalonePrice             ResourceTypeId = "standalone-price"
	ResourceTypeIdStore                       ResourceTypeId = "store"
	ResourceTypeIdTransaction                 ResourceTypeId = "transaction"
)

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
*	Field type for [DateTime](ctp:import:type:DateTime) values.
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
*	Field type for [Date](ctp:import:type:Date) values.
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
*	Defines an allowed value of a [CustomFieldEnumType](ctp:import:type:CustomFieldEnumType) field.
 */
type CustomFieldEnumValue struct {
	// Key of the value used as a programmatic identifier.
	Key string `json:"key"`
	// Descriptive label of the value.
	Label string `json:"label"`
}

/**
*	Field type for localized enum values.
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
*	Defines an allowed value of a [CustomFieldLocalizedEnumType](ctp:import:type:CustomFieldLocalizedEnumType) field.
 */
type CustomFieldLocalizedEnumValue struct {
	// Key of the value used as a programmatic identifier.
	Key string `json:"key"`
	// Descriptive localized label of the value.
	Label LocalizedString `json:"label"`
}

/**
*	Field type for [LocalizedString](ctp:import:type:LocalizedString) values.
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
*	Field type for [CentPrecisionMoney](ctp:import:type:CentPrecisionMoney) values.
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
*	Field type for [Reference](ctp:import:type:Reference) values.
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
*	Defines which resource type a [CustomFieldReferenceType](ctp:import:type:CustomFieldReferenceType) can reference.
 */
type CustomFieldReferenceValue string

const (
	CustomFieldReferenceValueAssociateRole    CustomFieldReferenceValue = "associate-role"
	CustomFieldReferenceValueBusinessUnit     CustomFieldReferenceValue = "business-unit"
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
*	Values of a SetType Custom Field are sets of values of the specified `elementType` (without duplicate elements).
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
*	Field type for [Time](ctp:import:type:Time) values.
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
*	Defines a [Custom Field](/../api/projects/custom-fields) and its meta-information. Maps to `Type.FieldDefinition`.
*
 */
type FieldDefinition struct {
	// Data type of the Custom Field to define.
	Type FieldType `json:"type"`
	// Name of the Custom Field to define. Must be unique for a given [ResourceTypeId](ctp:import:type:ResourceTypeId). In case there is a FieldDefinition with the same `name` in another Type, both FieldDefinitions must have the same `type`. This value cannot be changed after the Type is imported.
	Name string `json:"name"`
	// A human-readable label for the field.
	Label LocalizedString `json:"label"`
	// Defines whether the field is required to have a value. This value cannot be changed after the Type is imported.
	Required bool `json:"required"`
	// Provides a visual representation type for this field. It is only relevant for string-based field types like [CustomFieldStringType](ctp:import:type:CustomFieldStringType) and [CustomFieldLocalizedStringType](ctp:import:type:CustomFieldLocalizedStringType).
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

/**
*	The data representation for a Type to be imported that is persisted as a [Type](/../api/projects/types#type) in the Project.
*
 */
type TypeImport struct {
	// User-defined unique identifier for the Type. If a [Type](/../api/projects/types#type) with this `key` exists, it will be updated with the imported data.
	Key string `json:"key"`
	// Maps to `Type.name`.
	Name LocalizedString `json:"name"`
	// Maps to `Type.description`.
	Description *LocalizedString `json:"description,omitempty"`
	// Maps to `Type.resourceTypeIds`. This value cannot be changed after the Type is imported.
	ResourceTypeIds []ResourceTypeId `json:"resourceTypeIds"`
	// Maps to `Type.fieldDefinitions`.
	FieldDefinitions []FieldDefinition `json:"fieldDefinitions"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TypeImport) MarshalJSON() ([]byte, error) {
	type Alias TypeImport
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

	if raw["fieldDefinitions"] == nil {
		delete(raw, "fieldDefinitions")
	}

	return json.Marshal(raw)

}
