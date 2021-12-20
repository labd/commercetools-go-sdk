// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type CustomFieldEnumValue struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

type CustomFieldLocalizedEnumValue struct {
	Key   string          `json:"key"`
	Label LocalizedString `json:"label"`
}

type CustomFields struct {
	Type TypeReference `json:"type"`
	// A valid JSON object, based on FieldDefinition.
	Fields FieldContainer `json:"fields"`
}

type CustomFieldsDraft struct {
	// The `id` or the `key` of the type to use.
	Type TypeResourceIdentifier `json:"type"`
	// A valid JSON object, based on the FieldDefinitions of the Type.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// FieldContainer is something
type FieldContainer map[string]interface{}
type FieldDefinition struct {
	// Describes the type of the field.
	Type FieldType `json:"type"`
	// The name of the field.
	// The name must be between two and 36 characters long and can contain the ASCII letters A to Z in lowercase or uppercase, digits, underscores (`_`) and the hyphen-minus (`-`).
	// The name must be unique for a given resource type ID.
	// In case there is a field with the same name in another type it has to have the same FieldType also.
	Name string `json:"name"`
	// A human-readable label for the field.
	Label LocalizedString `json:"label"`
	// Whether the field is required to have a value.
	Required bool `json:"required"`
	// Provides a visual representation type for this field.
	// It is only relevant for string-based field types like StringType and LocalizedStringType.
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
			return nil, errors.New("Error processing discriminator field 'name'")
		}
	} else {
		return nil, errors.New("Invalid data")
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

type CustomFieldBooleanType struct {
}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldBooleanType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldBooleanType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "Boolean", Alias: (*Alias)(&obj)})
}

type CustomFieldDateTimeType struct {
}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldDateTimeType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldDateTimeType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "DateTime", Alias: (*Alias)(&obj)})
}

type CustomFieldDateType struct {
}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldDateType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldDateType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "Date", Alias: (*Alias)(&obj)})
}

type CustomFieldEnumType struct {
	Values []CustomFieldEnumValue `json:"values"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldEnumType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldEnumType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "Enum", Alias: (*Alias)(&obj)})
}

type CustomFieldLocalizedEnumType struct {
	Values []CustomFieldLocalizedEnumValue `json:"values"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldLocalizedEnumType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldLocalizedEnumType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "LocalizedEnum", Alias: (*Alias)(&obj)})
}

type CustomFieldLocalizedStringType struct {
}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldLocalizedStringType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldLocalizedStringType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "LocalizedString", Alias: (*Alias)(&obj)})
}

type CustomFieldMoneyType struct {
}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldMoneyType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldMoneyType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "Money", Alias: (*Alias)(&obj)})
}

type CustomFieldNumberType struct {
}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldNumberType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldNumberType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "Number", Alias: (*Alias)(&obj)})
}

type CustomFieldReferenceType struct {
	ReferenceTypeId ReferenceTypeId `json:"referenceTypeId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldReferenceType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldReferenceType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "Reference", Alias: (*Alias)(&obj)})
}

type CustomFieldSetType struct {
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

// MarshalJSON override to set the discriminator value
func (obj CustomFieldSetType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldSetType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "Set", Alias: (*Alias)(&obj)})
}

type CustomFieldStringType struct {
}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldStringType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldStringType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "String", Alias: (*Alias)(&obj)})
}

type CustomFieldTimeType struct {
}

// MarshalJSON override to set the discriminator value
func (obj CustomFieldTimeType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldTimeType
	return json.Marshal(struct {
		Action string `json:"name"`
		*Alias
	}{Action: "Time", Alias: (*Alias)(&obj)})
}

type ResourceTypeId string

const (
	ResourceTypeIdAddress                     ResourceTypeId = "address"
	ResourceTypeIdAsset                       ResourceTypeId = "asset"
	ResourceTypeIdCategory                    ResourceTypeId = "category"
	ResourceTypeIdChannel                     ResourceTypeId = "channel"
	ResourceTypeIdCustomer                    ResourceTypeId = "customer"
	ResourceTypeIdOrder                       ResourceTypeId = "order"
	ResourceTypeIdOrderEdit                   ResourceTypeId = "order-edit"
	ResourceTypeIdInventoryEntry              ResourceTypeId = "inventory-entry"
	ResourceTypeIdLineItem                    ResourceTypeId = "line-item"
	ResourceTypeIdCustomLineItem              ResourceTypeId = "custom-line-item"
	ResourceTypeIdProductPrice                ResourceTypeId = "product-price"
	ResourceTypeIdPayment                     ResourceTypeId = "payment"
	ResourceTypeIdPaymentInterfaceInteraction ResourceTypeId = "payment-interface-interaction"
	ResourceTypeIdReview                      ResourceTypeId = "review"
	ResourceTypeIdShippingMethod              ResourceTypeId = "shipping-method"
	ResourceTypeIdShoppingList                ResourceTypeId = "shopping-list"
	ResourceTypeIdShoppingListTextLineItem    ResourceTypeId = "shopping-list-text-line-item"
	ResourceTypeIdDiscountCode                ResourceTypeId = "discount-code"
	ResourceTypeIdCartDiscount                ResourceTypeId = "cart-discount"
	ResourceTypeIdCustomerGroup               ResourceTypeId = "customer-group"
	ResourceTypeIdStore                       ResourceTypeId = "store"
)

type Type struct {
	// The unique ID of the type.
	ID string `json:"id"`
	// The current version of the type.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Identifier for the type (max.
	// 256 characters).
	Key         string           `json:"key"`
	Name        LocalizedString  `json:"name"`
	Description *LocalizedString `json:"description,omitempty"`
	// Defines for which resource(s) the type is valid.
	ResourceTypeIds  []ResourceTypeId  `json:"resourceTypeIds"`
	FieldDefinitions []FieldDefinition `json:"fieldDefinitions"`
}

type TypeDraft struct {
	Key         string           `json:"key"`
	Name        LocalizedString  `json:"name"`
	Description *LocalizedString `json:"description,omitempty"`
	// The IDs of the resources that can be customized with this type.
	ResourceTypeIds  []ResourceTypeId  `json:"resourceTypeIds"`
	FieldDefinitions []FieldDefinition `json:"fieldDefinitions,omitempty"`
}

type TypePagedQueryResponse struct {
	Limit   int    `json:"limit"`
	Count   int    `json:"count"`
	Total   *int   `json:"total,omitempty"`
	Offset  int    `json:"offset"`
	Results []Type `json:"results"`
}

type TypeReference struct {
	// Unique ID of the referenced resource.
	ID  string `json:"id"`
	Obj *Type  `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeReference) MarshalJSON() ([]byte, error) {
	type Alias TypeReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "type", Alias: (*Alias)(&obj)})
}

type TypeResourceIdentifier struct {
	// Unique ID of the referenced resource. Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced resource. Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias TypeResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "type", Alias: (*Alias)(&obj)})
}

type TypeTextInputHint string

const (
	TypeTextInputHintSingleLine TypeTextInputHint = "SingleLine"
	TypeTextInputHintMultiLine  TypeTextInputHint = "MultiLine"
)

type TypeUpdate struct {
	Version int                `json:"version"`
	Actions []TypeUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *TypeUpdate) UnmarshalJSON(data []byte) error {
	type Alias TypeUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type TypeUpdateAction interface{}

func mapDiscriminatorTypeUpdateAction(input interface{}) (TypeUpdateAction, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("Invalid data")
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

type TypeAddEnumValueAction struct {
	FieldName string               `json:"fieldName"`
	Value     CustomFieldEnumValue `json:"value"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeAddEnumValueAction) MarshalJSON() ([]byte, error) {
	type Alias TypeAddEnumValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addEnumValue", Alias: (*Alias)(&obj)})
}

type TypeAddFieldDefinitionAction struct {
	FieldDefinition FieldDefinition `json:"fieldDefinition"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeAddFieldDefinitionAction) MarshalJSON() ([]byte, error) {
	type Alias TypeAddFieldDefinitionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addFieldDefinition", Alias: (*Alias)(&obj)})
}

type TypeAddLocalizedEnumValueAction struct {
	FieldName string                        `json:"fieldName"`
	Value     CustomFieldLocalizedEnumValue `json:"value"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeAddLocalizedEnumValueAction) MarshalJSON() ([]byte, error) {
	type Alias TypeAddLocalizedEnumValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLocalizedEnumValue", Alias: (*Alias)(&obj)})
}

type TypeChangeEnumValueLabelAction struct {
	FieldName string               `json:"fieldName"`
	Value     CustomFieldEnumValue `json:"value"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeEnumValueLabelAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeEnumValueLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEnumValueLabel", Alias: (*Alias)(&obj)})
}

type TypeChangeEnumValueOrderAction struct {
	FieldName string   `json:"fieldName"`
	Keys      []string `json:"keys"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeEnumValueOrderAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeEnumValueOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEnumValueOrder", Alias: (*Alias)(&obj)})
}

type TypeChangeFieldDefinitionOrderAction struct {
	FieldNames []string `json:"fieldNames"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeFieldDefinitionOrderAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeFieldDefinitionOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeFieldDefinitionOrder", Alias: (*Alias)(&obj)})
}

type TypeChangeInputHintAction struct {
	FieldName string            `json:"fieldName"`
	InputHint TypeTextInputHint `json:"inputHint"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeInputHintAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeInputHintAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeInputHint", Alias: (*Alias)(&obj)})
}

type TypeChangeKeyAction struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeKeyAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeKey", Alias: (*Alias)(&obj)})
}

type TypeChangeLabelAction struct {
	FieldName string          `json:"fieldName"`
	Label     LocalizedString `json:"label"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeLabelAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLabel", Alias: (*Alias)(&obj)})
}

type TypeChangeLocalizedEnumValueLabelAction struct {
	FieldName string                        `json:"fieldName"`
	Value     CustomFieldLocalizedEnumValue `json:"value"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeLocalizedEnumValueLabelAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeLocalizedEnumValueLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLocalizedEnumValueLabel", Alias: (*Alias)(&obj)})
}

type TypeChangeLocalizedEnumValueOrderAction struct {
	FieldName string   `json:"fieldName"`
	Keys      []string `json:"keys"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeLocalizedEnumValueOrderAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeLocalizedEnumValueOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLocalizedEnumValueOrder", Alias: (*Alias)(&obj)})
}

type TypeChangeNameAction struct {
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type TypeRemoveFieldDefinitionAction struct {
	FieldName string `json:"fieldName"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeRemoveFieldDefinitionAction) MarshalJSON() ([]byte, error) {
	type Alias TypeRemoveFieldDefinitionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeFieldDefinition", Alias: (*Alias)(&obj)})
}

type TypeSetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias TypeSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}
