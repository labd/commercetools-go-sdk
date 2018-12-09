// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type CustomFieldBooleanType struct{}

func (obj CustomFieldBooleanType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldBooleanType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "Boolean", Alias: (*Alias)(&obj)})
}

type CustomFieldDateTimeType struct{}

func (obj CustomFieldDateTimeType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldDateTimeType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "DateTime", Alias: (*Alias)(&obj)})
}

type CustomFieldDateType struct{}

func (obj CustomFieldDateType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldDateType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "Date", Alias: (*Alias)(&obj)})
}

type CustomFieldEnumType struct {
	Values []CustomFieldEnumValue `json:"values"`
}

func (obj CustomFieldEnumType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldEnumType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "Enum", Alias: (*Alias)(&obj)})
}

type CustomFieldEnumValue struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}

type CustomFieldLocalizedEnumType struct {
	Values []CustomFieldLocalizedEnumValue `json:"values"`
}

func (obj CustomFieldLocalizedEnumType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldLocalizedEnumType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "LocalizedEnum", Alias: (*Alias)(&obj)})
}

type CustomFieldLocalizedEnumValue struct {
	Label *LocalizedString `json:"label"`
	Key   string           `json:"key"`
}

type CustomFieldLocalizedStringType struct{}

func (obj CustomFieldLocalizedStringType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldLocalizedStringType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "LocalizedString", Alias: (*Alias)(&obj)})
}

type CustomFieldMoneyType struct{}

func (obj CustomFieldMoneyType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldMoneyType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "Money", Alias: (*Alias)(&obj)})
}

type CustomFieldNumberType struct{}

func (obj CustomFieldNumberType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldNumberType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "Number", Alias: (*Alias)(&obj)})
}

type CustomFieldReferenceType struct {
	ReferenceTypeID ReferenceTypeID `json:"referenceTypeId"`
}

func (obj CustomFieldReferenceType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldReferenceType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "Reference", Alias: (*Alias)(&obj)})
}

type CustomFieldSetType struct {
	ElementType FieldType `json:"elementType"`
}

func (obj CustomFieldSetType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldSetType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "Set", Alias: (*Alias)(&obj)})
}
func (obj *CustomFieldSetType) UnmarshalJSON(data []byte) error {
	type Alias CustomFieldSetType
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ElementType != nil {
		obj.ElementType = AbstractFieldTypeDiscriminatorMapping(obj.ElementType)
	}

	return nil
}

type CustomFieldStringType struct{}

func (obj CustomFieldStringType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldStringType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "String", Alias: (*Alias)(&obj)})
}

type CustomFieldTimeType struct{}

func (obj CustomFieldTimeType) MarshalJSON() ([]byte, error) {
	type Alias CustomFieldTimeType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "Time", Alias: (*Alias)(&obj)})
}

type CustomFields struct {
	Type   *TypeReference  `json:"type"`
	Fields *FieldContainer `json:"fields"`
}

type CustomFieldsDraft struct {
	Type   *ResourceIdentifier `json:"type"`
	Fields *FieldContainer     `json:"fields,omitempty"`
}
type FieldContainer map[string]string

type FieldDefinition struct {
	Type      FieldType        `json:"type"`
	Required  bool             `json:"required"`
	Name      string           `json:"name"`
	Label     *LocalizedString `json:"label"`
	InputHint string           `json:"inputHint,omitempty"`
}

func (obj *FieldDefinition) UnmarshalJSON(data []byte) error {
	type Alias FieldDefinition
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Type != nil {
		obj.Type = AbstractFieldTypeDiscriminatorMapping(obj.Type)
	}

	return nil
}

type FieldType interface{}
type AbstractFieldType struct{}

func AbstractFieldTypeDiscriminatorMapping(input FieldType) FieldType {
	discriminator := input.(map[string]interface{})["name"]
	switch discriminator {
	case "Boolean":
		new := CustomFieldBooleanType{}
		mapstructure.Decode(input, &new)
		return new
	case "DateTime":
		new := CustomFieldDateTimeType{}
		mapstructure.Decode(input, &new)
		return new
	case "Date":
		new := CustomFieldDateType{}
		mapstructure.Decode(input, &new)
		return new
	case "Enum":
		new := CustomFieldEnumType{}
		mapstructure.Decode(input, &new)
		return new
	case "LocalizedEnum":
		new := CustomFieldLocalizedEnumType{}
		mapstructure.Decode(input, &new)
		return new
	case "LocalizedString":
		new := CustomFieldLocalizedStringType{}
		mapstructure.Decode(input, &new)
		return new
	case "Money":
		new := CustomFieldMoneyType{}
		mapstructure.Decode(input, &new)
		return new
	case "Number":
		new := CustomFieldNumberType{}
		mapstructure.Decode(input, &new)
		return new
	case "Reference":
		new := CustomFieldReferenceType{}
		mapstructure.Decode(input, &new)
		return new
	case "Set":
		new := CustomFieldSetType{}
		mapstructure.Decode(input, &new)
		if new.ElementType != nil {
			new.ElementType = AbstractFieldTypeDiscriminatorMapping(new.ElementType)
		}

		return new
	case "String":
		new := CustomFieldStringType{}
		mapstructure.Decode(input, &new)
		return new
	case "Time":
		new := CustomFieldTimeType{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

type ResourceTypeID string

const (
	ResourceTypeIDAsset                       ResourceTypeID = "asset"
	ResourceTypeIDCategory                    ResourceTypeID = "category"
	ResourceTypeIDChannel                     ResourceTypeID = "channel"
	ResourceTypeIDCustomer                    ResourceTypeID = "customer"
	ResourceTypeIDOrder                       ResourceTypeID = "order"
	ResourceTypeIDOrderEdit                   ResourceTypeID = "order-edit"
	ResourceTypeIDInventoryEntry              ResourceTypeID = "inventory-entry"
	ResourceTypeIDLineItem                    ResourceTypeID = "line-item"
	ResourceTypeIDCustomLineItem              ResourceTypeID = "custom-line-item"
	ResourceTypeIDProductPrice                ResourceTypeID = "product-price"
	ResourceTypeIDPayment                     ResourceTypeID = "payment"
	ResourceTypeIDPaymentInterfaceInteraction ResourceTypeID = "payment-interface-interaction"
	ResourceTypeIDReview                      ResourceTypeID = "review"
	ResourceTypeIDShoppingList                ResourceTypeID = "shopping-list"
	ResourceTypeIDShoppingListTextLineItem    ResourceTypeID = "shopping-list-text-line-item"
	ResourceTypeIDDiscountCode                ResourceTypeID = "discount-code"
	ResourceTypeIDCartDiscount                ResourceTypeID = "cart-discount"
	ResourceTypeIDCustomerGroup               ResourceTypeID = "customer-group"
)

type Type struct {
	Version          int               `json:"version"`
	LastModifiedAt   time.Time         `json:"lastModifiedAt"`
	ID               string            `json:"id"`
	CreatedAt        time.Time         `json:"createdAt"`
	ResourceTypeIds  []ResourceTypeID  `json:"resourceTypeIds"`
	Name             *LocalizedString  `json:"name"`
	Key              string            `json:"key"`
	FieldDefinitions []FieldDefinition `json:"fieldDefinitions"`
	Description      *LocalizedString  `json:"description,omitempty"`
}

type TypeAddEnumValueAction struct {
	Value     *CustomFieldEnumValue `json:"value"`
	FieldName string                `json:"fieldName"`
}

func (obj TypeAddEnumValueAction) MarshalJSON() ([]byte, error) {
	type Alias TypeAddEnumValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addEnumValue", Alias: (*Alias)(&obj)})
}

type TypeAddFieldDefinitionAction struct {
	FieldDefinition *FieldDefinition `json:"fieldDefinition"`
}

func (obj TypeAddFieldDefinitionAction) MarshalJSON() ([]byte, error) {
	type Alias TypeAddFieldDefinitionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addFieldDefinition", Alias: (*Alias)(&obj)})
}

type TypeAddLocalizedEnumValueAction struct {
	Value     *CustomFieldLocalizedEnumValue `json:"value"`
	FieldName string                         `json:"fieldName"`
}

func (obj TypeAddLocalizedEnumValueAction) MarshalJSON() ([]byte, error) {
	type Alias TypeAddLocalizedEnumValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLocalizedEnumValue", Alias: (*Alias)(&obj)})
}

type TypeChangeEnumValueOrderAction struct {
	Keys      []string `json:"keys"`
	FieldName string   `json:"fieldName"`
}

func (obj TypeChangeEnumValueOrderAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeEnumValueOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEnumValueOrder", Alias: (*Alias)(&obj)})
}

type TypeChangeFieldDefinitionLabelAction struct {
	Label     *LocalizedString `json:"label"`
	FieldName string           `json:"fieldName"`
}

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

func (obj TypeChangeFieldDefinitionOrderAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeFieldDefinitionOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeFieldDefinitionOrder", Alias: (*Alias)(&obj)})
}

type TypeChangeKeyAction struct {
	Key string `json:"key"`
}

func (obj TypeChangeKeyAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeKey", Alias: (*Alias)(&obj)})
}

type TypeChangeLabelAction struct {
	Label     *LocalizedString `json:"label"`
	FieldName string           `json:"fieldName"`
}

func (obj TypeChangeLabelAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLabel", Alias: (*Alias)(&obj)})
}

type TypeChangeLocalizedEnumValueOrderAction struct {
	Keys      []string `json:"keys"`
	FieldName string   `json:"fieldName"`
}

func (obj TypeChangeLocalizedEnumValueOrderAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeLocalizedEnumValueOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLocalizedEnumValueOrder", Alias: (*Alias)(&obj)})
}

type TypeChangeNameAction struct {
	Name *LocalizedString `json:"name"`
}

func (obj TypeChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias TypeChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type TypeDraft struct {
	ResourceTypeIds  []ResourceTypeID  `json:"resourceTypeIds"`
	Name             *LocalizedString  `json:"name"`
	Key              string            `json:"key"`
	FieldDefinitions []FieldDefinition `json:"fieldDefinitions,omitempty"`
	Description      *LocalizedString  `json:"description,omitempty"`
}

type TypePagedQueryResponse struct {
	Total   int    `json:"total,omitempty"`
	Offset  int    `json:"offset"`
	Count   int    `json:"count"`
	Results []Type `json:"results"`
}

type TypeReference struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
	Obj *Type  `json:"obj,omitempty"`
}

func (obj TypeReference) MarshalJSON() ([]byte, error) {
	type Alias TypeReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "type", Alias: (*Alias)(&obj)})
}

type TypeRemoveFieldDefinitionAction struct {
	FieldName string `json:"fieldName"`
}

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

func (obj TypeSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias TypeSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type TypeUpdate struct {
	Version int                `json:"version"`
	Actions []TypeUpdateAction `json:"actions"`
}

func (obj *TypeUpdate) UnmarshalJSON(data []byte) error {
	type Alias TypeUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractTypeUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type TypeUpdateAction interface{}
type AbstractTypeUpdateAction struct{}

func AbstractTypeUpdateActionDiscriminatorMapping(input TypeUpdateAction) TypeUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addEnumValue":
		new := TypeAddEnumValueAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addFieldDefinition":
		new := TypeAddFieldDefinitionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addLocalizedEnumValue":
		new := TypeAddLocalizedEnumValueAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeEnumValueOrder":
		new := TypeChangeEnumValueOrderAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeFieldDefinitionLabel":
		new := TypeChangeFieldDefinitionLabelAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeFieldDefinitionOrder":
		new := TypeChangeFieldDefinitionOrderAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeKey":
		new := TypeChangeKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeLabel":
		new := TypeChangeLabelAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeLocalizedEnumValueOrder":
		new := TypeChangeLocalizedEnumValueOrderAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeName":
		new := TypeChangeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeFieldDefinition":
		new := TypeRemoveFieldDefinitionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDescription":
		new := TypeSetDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}
