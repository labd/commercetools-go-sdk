// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// AttributeConstraintEnum is an enum type
type AttributeConstraintEnum string

// Enum values for AttributeConstraintEnum
const (
	AttributeConstraintEnumNone              AttributeConstraintEnum = "None"
	AttributeConstraintEnumUnique            AttributeConstraintEnum = "Unique"
	AttributeConstraintEnumCombinationUnique AttributeConstraintEnum = "CombinationUnique"
	AttributeConstraintEnumSameForAll        AttributeConstraintEnum = "SameForAll"
)

// AttributeConstraintEnumDraft is an enum type
type AttributeConstraintEnumDraft string

// Enum values for AttributeConstraintEnumDraft
const (
	AttributeConstraintEnumDraftNone AttributeConstraintEnumDraft = "None"
)

// TextInputHint is an enum type
type TextInputHint string

// Enum values for TextInputHint
const (
	TextInputHintSingleLine TextInputHint = "SingleLine"
	TextInputHintMultiLine  TextInputHint = "MultiLine"
)

// AttributeType uses name as discriminator attribute
type AttributeType interface{}

func mapDiscriminatorAttributeType(input AttributeType) AttributeType {
	discriminator := input.(map[string]interface{})["name"]
	switch discriminator {
	case "boolean":
		new := AttributeBooleanType{}
		mapstructure.Decode(input, &new)
		return new
	case "datetime":
		new := AttributeDateTimeType{}
		mapstructure.Decode(input, &new)
		return new
	case "date":
		new := AttributeDateType{}
		mapstructure.Decode(input, &new)
		return new
	case "enum":
		new := AttributeEnumType{}
		mapstructure.Decode(input, &new)
		return new
	case "ltext":
		new := AttributeLocalizableTextType{}
		mapstructure.Decode(input, &new)
		return new
	case "lenum":
		new := AttributeLocalizedEnumType{}
		mapstructure.Decode(input, &new)
		return new
	case "money":
		new := AttributeMoneyType{}
		mapstructure.Decode(input, &new)
		return new
	case "nested":
		new := AttributeNestedType{}
		mapstructure.Decode(input, &new)
		return new
	case "number":
		new := AttributeNumberType{}
		mapstructure.Decode(input, &new)
		return new
	case "reference":
		new := AttributeReferenceType{}
		mapstructure.Decode(input, &new)
		return new
	case "set":
		new := AttributeSetType{}
		mapstructure.Decode(input, &new)
		if new.ElementType != nil {
			new.ElementType = mapDiscriminatorAttributeType(new.ElementType)
		}

		return new
	case "text":
		new := AttributeTextType{}
		mapstructure.Decode(input, &new)
		return new
	case "time":
		new := AttributeTimeType{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// ProductTypeUpdateAction uses action as discriminator attribute
type ProductTypeUpdateAction interface{}

func mapDiscriminatorProductTypeUpdateAction(input ProductTypeUpdateAction) ProductTypeUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addAttributeDefinition":
		new := ProductTypeAddAttributeDefinitionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addLocalizedEnumValue":
		new := ProductTypeAddLocalizedEnumValueAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addPlainEnumValue":
		new := ProductTypeAddPlainEnumValueAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeAttributeConstraint":
		new := ProductTypeChangeAttributeConstraintAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeAttributeName":
		new := ProductTypeChangeAttributeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeAttributeOrder":
		new := ProductTypeChangeAttributeOrderAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeDescription":
		new := ProductTypeChangeDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeEnumKey":
		new := ProductTypeChangeEnumKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeInputHint":
		new := ProductTypeChangeInputHintAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeIsSearchable":
		new := ProductTypeChangeIsSearchableAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeLabel":
		new := ProductTypeChangeLabelAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeLocalizedEnumValueLabel":
		new := ProductTypeChangeLocalizedEnumValueLabelAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeLocalizedEnumValueOrder":
		new := ProductTypeChangeLocalizedEnumValueOrderAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeName":
		new := ProductTypeChangeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changePlainEnumValueLabel":
		new := ProductTypeChangePlainEnumValueLabelAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changePlainEnumValueOrder":
		new := ProductTypeChangePlainEnumValueOrderAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeAttributeDefinition":
		new := ProductTypeRemoveAttributeDefinitionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeEnumValues":
		new := ProductTypeRemoveEnumValuesAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setInputTip":
		new := ProductTypeSetInputTipAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := ProductTypeSetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// AttributeBooleanType implements the interface AttributeType
type AttributeBooleanType struct{}

// MarshalJSON override to set the discriminator value
func (obj AttributeBooleanType) MarshalJSON() ([]byte, error) {
	type Alias AttributeBooleanType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "boolean", Alias: (*Alias)(&obj)})
}

// AttributeDateTimeType implements the interface AttributeType
type AttributeDateTimeType struct{}

// MarshalJSON override to set the discriminator value
func (obj AttributeDateTimeType) MarshalJSON() ([]byte, error) {
	type Alias AttributeDateTimeType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "datetime", Alias: (*Alias)(&obj)})
}

// AttributeDateType implements the interface AttributeType
type AttributeDateType struct{}

// MarshalJSON override to set the discriminator value
func (obj AttributeDateType) MarshalJSON() ([]byte, error) {
	type Alias AttributeDateType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "date", Alias: (*Alias)(&obj)})
}

// AttributeDefinition is a standalone struct
type AttributeDefinition struct {
	Type                AttributeType           `json:"type"`
	Name                string                  `json:"name"`
	Label               *LocalizedString        `json:"label"`
	IsSearchable        bool                    `json:"isSearchable"`
	IsRequired          bool                    `json:"isRequired"`
	InputTip            *LocalizedString        `json:"inputTip,omitempty"`
	InputHint           TextInputHint           `json:"inputHint"`
	AttributeConstraint AttributeConstraintEnum `json:"attributeConstraint"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AttributeDefinition) UnmarshalJSON(data []byte) error {
	type Alias AttributeDefinition
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Type != nil {
		obj.Type = mapDiscriminatorAttributeType(obj.Type)
	}

	return nil
}

// AttributeDefinitionDraft is a standalone struct
type AttributeDefinitionDraft struct {
	Type                AttributeType           `json:"type"`
	Name                string                  `json:"name"`
	Label               *LocalizedString        `json:"label"`
	IsSearchable        bool                    `json:"isSearchable,omitempty"`
	IsRequired          bool                    `json:"isRequired"`
	InputTip            *LocalizedString        `json:"inputTip,omitempty"`
	InputHint           TextInputHint           `json:"inputHint,omitempty"`
	AttributeConstraint AttributeConstraintEnum `json:"attributeConstraint,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AttributeDefinitionDraft) UnmarshalJSON(data []byte) error {
	type Alias AttributeDefinitionDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Type != nil {
		obj.Type = mapDiscriminatorAttributeType(obj.Type)
	}

	return nil
}

// AttributeEnumType implements the interface AttributeType
type AttributeEnumType struct {
	Values []AttributePlainEnumValue `json:"values"`
}

// MarshalJSON override to set the discriminator value
func (obj AttributeEnumType) MarshalJSON() ([]byte, error) {
	type Alias AttributeEnumType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "enum", Alias: (*Alias)(&obj)})
}

// AttributeLocalizableTextType implements the interface AttributeType
type AttributeLocalizableTextType struct{}

// MarshalJSON override to set the discriminator value
func (obj AttributeLocalizableTextType) MarshalJSON() ([]byte, error) {
	type Alias AttributeLocalizableTextType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "ltext", Alias: (*Alias)(&obj)})
}

// AttributeLocalizedEnumType implements the interface AttributeType
type AttributeLocalizedEnumType struct {
	Values []AttributeLocalizedEnumValue `json:"values"`
}

// MarshalJSON override to set the discriminator value
func (obj AttributeLocalizedEnumType) MarshalJSON() ([]byte, error) {
	type Alias AttributeLocalizedEnumType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "lenum", Alias: (*Alias)(&obj)})
}

// AttributeLocalizedEnumValue is a standalone struct
type AttributeLocalizedEnumValue struct {
	Label *LocalizedString `json:"label"`
	Key   string           `json:"key"`
}

// AttributeMoneyType implements the interface AttributeType
type AttributeMoneyType struct{}

// MarshalJSON override to set the discriminator value
func (obj AttributeMoneyType) MarshalJSON() ([]byte, error) {
	type Alias AttributeMoneyType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "money", Alias: (*Alias)(&obj)})
}

// AttributeNestedType implements the interface AttributeType
type AttributeNestedType struct {
	TypeReference *ProductTypeReference `json:"typeReference"`
}

// MarshalJSON override to set the discriminator value
func (obj AttributeNestedType) MarshalJSON() ([]byte, error) {
	type Alias AttributeNestedType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "nested", Alias: (*Alias)(&obj)})
}

// AttributeNumberType implements the interface AttributeType
type AttributeNumberType struct{}

// MarshalJSON override to set the discriminator value
func (obj AttributeNumberType) MarshalJSON() ([]byte, error) {
	type Alias AttributeNumberType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "number", Alias: (*Alias)(&obj)})
}

// AttributePlainEnumValue is a standalone struct
type AttributePlainEnumValue struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}

// AttributeReferenceType implements the interface AttributeType
type AttributeReferenceType struct {
	ReferenceTypeID ReferenceTypeID `json:"referenceTypeId"`
}

// MarshalJSON override to set the discriminator value
func (obj AttributeReferenceType) MarshalJSON() ([]byte, error) {
	type Alias AttributeReferenceType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "reference", Alias: (*Alias)(&obj)})
}

// AttributeSetType implements the interface AttributeType
type AttributeSetType struct {
	ElementType AttributeType `json:"elementType"`
}

// MarshalJSON override to set the discriminator value
func (obj AttributeSetType) MarshalJSON() ([]byte, error) {
	type Alias AttributeSetType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "set", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *AttributeSetType) UnmarshalJSON(data []byte) error {
	type Alias AttributeSetType
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ElementType != nil {
		obj.ElementType = mapDiscriminatorAttributeType(obj.ElementType)
	}

	return nil
}

// AttributeTextType implements the interface AttributeType
type AttributeTextType struct{}

// MarshalJSON override to set the discriminator value
func (obj AttributeTextType) MarshalJSON() ([]byte, error) {
	type Alias AttributeTextType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "text", Alias: (*Alias)(&obj)})
}

// AttributeTimeType implements the interface AttributeType
type AttributeTimeType struct{}

// MarshalJSON override to set the discriminator value
func (obj AttributeTimeType) MarshalJSON() ([]byte, error) {
	type Alias AttributeTimeType
	return json.Marshal(struct {
		Name string `json:"name"`
		*Alias
	}{Name: "time", Alias: (*Alias)(&obj)})
}

// ProductType is of type Resource
type ProductType struct {
	Version        int                   `json:"version"`
	LastModifiedAt time.Time             `json:"lastModifiedAt"`
	ID             string                `json:"id"`
	CreatedAt      time.Time             `json:"createdAt"`
	Name           string                `json:"name"`
	Key            string                `json:"key,omitempty"`
	Description    string                `json:"description"`
	Attributes     []AttributeDefinition `json:"attributes,omitempty"`
}

// ProductTypeAddAttributeDefinitionAction implements the interface ProductTypeUpdateAction
type ProductTypeAddAttributeDefinitionAction struct {
	Attribute *AttributeDefinitionDraft `json:"attribute"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeAddAttributeDefinitionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeAddAttributeDefinitionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAttributeDefinition", Alias: (*Alias)(&obj)})
}

// ProductTypeAddLocalizedEnumValueAction implements the interface ProductTypeUpdateAction
type ProductTypeAddLocalizedEnumValueAction struct {
	Value         *AttributeLocalizedEnumValue `json:"value"`
	AttributeName string                       `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeAddLocalizedEnumValueAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeAddLocalizedEnumValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLocalizedEnumValue", Alias: (*Alias)(&obj)})
}

// ProductTypeAddPlainEnumValueAction implements the interface ProductTypeUpdateAction
type ProductTypeAddPlainEnumValueAction struct {
	Value         *AttributePlainEnumValue `json:"value"`
	AttributeName string                   `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeAddPlainEnumValueAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeAddPlainEnumValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPlainEnumValue", Alias: (*Alias)(&obj)})
}

// ProductTypeChangeAttributeConstraintAction implements the interface ProductTypeUpdateAction
type ProductTypeChangeAttributeConstraintAction struct {
	NewValue      AttributeConstraintEnumDraft `json:"newValue"`
	AttributeName string                       `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeChangeAttributeConstraintAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeAttributeConstraintAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAttributeConstraint", Alias: (*Alias)(&obj)})
}

// ProductTypeChangeAttributeNameAction implements the interface ProductTypeUpdateAction
type ProductTypeChangeAttributeNameAction struct {
	NewAttributeName string `json:"newAttributeName"`
	AttributeName    string `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeChangeAttributeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeAttributeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAttributeName", Alias: (*Alias)(&obj)})
}

// ProductTypeChangeAttributeOrderAction implements the interface ProductTypeUpdateAction
type ProductTypeChangeAttributeOrderAction struct {
	Attributes []AttributeDefinition `json:"attributes"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeChangeAttributeOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeAttributeOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAttributeOrder", Alias: (*Alias)(&obj)})
}

// ProductTypeChangeDescriptionAction implements the interface ProductTypeUpdateAction
type ProductTypeChangeDescriptionAction struct {
	Description string `json:"description"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeChangeDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeDescription", Alias: (*Alias)(&obj)})
}

// ProductTypeChangeEnumKeyAction implements the interface ProductTypeUpdateAction
type ProductTypeChangeEnumKeyAction struct {
	NewKey        string `json:"newKey"`
	Key           string `json:"key"`
	AttributeName string `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeChangeEnumKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeEnumKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeEnumKey", Alias: (*Alias)(&obj)})
}

// ProductTypeChangeInputHintAction implements the interface ProductTypeUpdateAction
type ProductTypeChangeInputHintAction struct {
	NewValue      TextInputHint `json:"newValue"`
	AttributeName string        `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeChangeInputHintAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeInputHintAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeInputHint", Alias: (*Alias)(&obj)})
}

// ProductTypeChangeIsSearchableAction implements the interface ProductTypeUpdateAction
type ProductTypeChangeIsSearchableAction struct {
	IsSearchable  bool   `json:"isSearchable"`
	AttributeName string `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeChangeIsSearchableAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeIsSearchableAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsSearchable", Alias: (*Alias)(&obj)})
}

// ProductTypeChangeLabelAction implements the interface ProductTypeUpdateAction
type ProductTypeChangeLabelAction struct {
	Label         *LocalizedString `json:"label"`
	AttributeName string           `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeChangeLabelAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLabel", Alias: (*Alias)(&obj)})
}

// ProductTypeChangeLocalizedEnumValueLabelAction implements the interface ProductTypeUpdateAction
type ProductTypeChangeLocalizedEnumValueLabelAction struct {
	NewValue      *AttributeLocalizedEnumValue `json:"newValue"`
	AttributeName string                       `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeChangeLocalizedEnumValueLabelAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeLocalizedEnumValueLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLocalizedEnumValueLabel", Alias: (*Alias)(&obj)})
}

// ProductTypeChangeLocalizedEnumValueOrderAction implements the interface ProductTypeUpdateAction
type ProductTypeChangeLocalizedEnumValueOrderAction struct {
	Values        []AttributeLocalizedEnumValue `json:"values"`
	AttributeName string                        `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeChangeLocalizedEnumValueOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeLocalizedEnumValueOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLocalizedEnumValueOrder", Alias: (*Alias)(&obj)})
}

// ProductTypeChangeNameAction implements the interface ProductTypeUpdateAction
type ProductTypeChangeNameAction struct {
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// ProductTypeChangePlainEnumValueLabelAction implements the interface ProductTypeUpdateAction
type ProductTypeChangePlainEnumValueLabelAction struct {
	NewValue      *AttributePlainEnumValue `json:"newValue"`
	AttributeName string                   `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeChangePlainEnumValueLabelAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangePlainEnumValueLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePlainEnumValueLabel", Alias: (*Alias)(&obj)})
}

// ProductTypeChangePlainEnumValueOrderAction implements the interface ProductTypeUpdateAction
type ProductTypeChangePlainEnumValueOrderAction struct {
	Values        []AttributePlainEnumValue `json:"values"`
	AttributeName string                    `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeChangePlainEnumValueOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeChangePlainEnumValueOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePlainEnumValueOrder", Alias: (*Alias)(&obj)})
}

// ProductTypeDraft is a standalone struct
type ProductTypeDraft struct {
	Name        string                     `json:"name"`
	Key         string                     `json:"key,omitempty"`
	Description string                     `json:"description"`
	Attributes  []AttributeDefinitionDraft `json:"attributes,omitempty"`
}

// ProductTypePagedQueryResponse is of type PagedQueryResponse
type ProductTypePagedQueryResponse struct {
	Total   int           `json:"total,omitempty"`
	Offset  int           `json:"offset"`
	Count   int           `json:"count"`
	Results []ProductType `json:"results"`
}

// ProductTypeReference implements the interface Reference
type ProductTypeReference struct {
	Key string       `json:"key,omitempty"`
	ID  string       `json:"id,omitempty"`
	Obj *ProductType `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeReference) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "product-type", Alias: (*Alias)(&obj)})
}

// ProductTypeRemoveAttributeDefinitionAction implements the interface ProductTypeUpdateAction
type ProductTypeRemoveAttributeDefinitionAction struct {
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeRemoveAttributeDefinitionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeRemoveAttributeDefinitionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAttributeDefinition", Alias: (*Alias)(&obj)})
}

// ProductTypeRemoveEnumValuesAction implements the interface ProductTypeUpdateAction
type ProductTypeRemoveEnumValuesAction struct {
	Keys          []string `json:"keys"`
	AttributeName string   `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeRemoveEnumValuesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeRemoveEnumValuesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeEnumValues", Alias: (*Alias)(&obj)})
}

// ProductTypeSetInputTipAction implements the interface ProductTypeUpdateAction
type ProductTypeSetInputTipAction struct {
	InputTip      *LocalizedString `json:"inputTip,omitempty"`
	AttributeName string           `json:"attributeName"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeSetInputTipAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeSetInputTipAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setInputTip", Alias: (*Alias)(&obj)})
}

// ProductTypeSetKeyAction implements the interface ProductTypeUpdateAction
type ProductTypeSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProductTypeSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTypeSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// ProductTypeUpdate is of type Update
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
		obj.Actions[i] = mapDiscriminatorProductTypeUpdateAction(obj.Actions[i])
	}

	return nil
}
