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
	Fields *FieldContainer `json:"fields,omitEmpty"`
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
	InputHint TypeTextInputHint `json:"inputHint,omitEmpty"`
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
		new := CustomFieldBooleanType{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "DateTime":
		new := CustomFieldDateTimeType{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Date":
		new := CustomFieldDateType{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Enum":
		new := CustomFieldEnumType{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "LocalizedEnum":
		new := CustomFieldLocalizedEnumType{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "LocalizedString":
		new := CustomFieldLocalizedStringType{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Money":
		new := CustomFieldMoneyType{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Number":
		new := CustomFieldNumberType{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Reference":
		new := CustomFieldReferenceType{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Set":
		new := CustomFieldSetType{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "String":
		new := CustomFieldStringType{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Time":
		new := CustomFieldTimeType{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
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
)

type Type struct {
	// The unique ID of the type.
	Id string `json:"id"`
	// The current version of the type.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for events not tracked.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	// Present on resources created after 1/02/2019 except for events not tracked.
	CreatedBy *CreatedBy `json:"createdBy,omitEmpty"`
	// Identifier for the type (max.
	// 256 characters).
	Key         string           `json:"key"`
	Name        LocalizedString  `json:"name"`
	Description *LocalizedString `json:"description,omitEmpty"`
	// Defines for which resource(s) the type is valid.
	ResourceTypeIds  []ResourceTypeId  `json:"resourceTypeIds"`
	FieldDefinitions []FieldDefinition `json:"fieldDefinitions"`
}

type TypeDraft struct {
	Key         string           `json:"key"`
	Name        LocalizedString  `json:"name"`
	Description *LocalizedString `json:"description,omitEmpty"`
	// The IDs of the resources that can be customized with this type.
	ResourceTypeIds  []ResourceTypeId  `json:"resourceTypeIds"`
	FieldDefinitions []FieldDefinition `json:"fieldDefinitions,omitEmpty"`
}

type TypePagedQueryResponse struct {
	Limit   int    `json:"limit"`
	Count   int    `json:"count"`
	Total   int    `json:"total,omitEmpty"`
	Offset  int    `json:"offset"`
	Results []Type `json:"results"`
}

type TypeReference struct {
	Id  string `json:"id"`
	Obj *Type  `json:"obj,omitEmpty"`
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
	Id  string `json:"id,omitEmpty"`
	Key string `json:"key,omitEmpty"`
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
		new := TypeAddEnumValueAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addFieldDefinition":
		new := TypeAddFieldDefinitionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addLocalizedEnumValue":
		new := TypeAddLocalizedEnumValueAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeEnumValueLabel":
		new := TypeChangeEnumValueLabelAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeEnumValueOrder":
		new := TypeChangeEnumValueOrderAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeFieldDefinitionLabel":
		new := TypeChangeFieldDefinitionLabelAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeFieldDefinitionOrder":
		new := TypeChangeFieldDefinitionOrderAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeInputHint":
		new := TypeChangeInputHintAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeKey":
		new := TypeChangeKeyAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLabel":
		new := TypeChangeLabelAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLocalizedEnumValueLabel":
		new := TypeChangeLocalizedEnumValueLabelAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLocalizedEnumValueOrder":
		new := TypeChangeLocalizedEnumValueOrderAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := TypeChangeNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeFieldDefinition":
		new := TypeRemoveFieldDefinitionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := TypeSetDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
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

type TypeChangeFieldDefinitionLabelAction struct {
	FieldName string          `json:"fieldName"`
	Label     LocalizedString `json:"label"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeChangeFieldDefinitionLabelAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeFieldDefinitionLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeFieldDefinitionLabel", Alias: (*Alias)(&obj)})
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
	Description *LocalizedString `json:"description,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj TypeSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias TypeSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}
