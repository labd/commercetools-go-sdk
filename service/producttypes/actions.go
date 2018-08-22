package producttypes

import (
	"encoding/json"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

// SetKey will set the key of the product type being updated.
type SetKey struct {
	// If key is absent or null, this field will be removed if it exists.
	// Min. 2 and max. 256 characters
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to add the action value
func (ua SetKey) MarshalJSON() ([]byte, error) {
	type Alias SetKey

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setKey",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeName will change the name of the product type being updated.
type ChangeName struct {
	Name string `json:"name"`
}

// MarshalJSON override to add the action value
func (ua ChangeName) MarshalJSON() ([]byte, error) {
	type Alias ChangeName

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeName",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeDescription will change the description of the product type being updated.
type ChangeDescription struct {
	Description string `json:"description"`
}

// MarshalJSON override to add the action value
func (ua ChangeDescription) MarshalJSON() ([]byte, error) {
	type Alias ChangeDescription

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeDescription",
		Alias:  (*Alias)(&ua),
	})
}

// AddAttributeDefinition will add an attribute definition to the product type being updated.
type AddAttributeDefinition struct {
	Attribute AttributeDefinitionDraft `json:"attribute"`
}

// MarshalJSON override to add the action value
func (ua AddAttributeDefinition) MarshalJSON() ([]byte, error) {
	type Alias AddAttributeDefinition

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "addAttributeDefinition",
		Alias:  (*Alias)(&ua),
	})
}

// RemoveAttributeDefinition will remove an attribute definition to the product type being updated.
// This removal also deletes all corresponding attributes on all Products with this product type.
// The removal of the attributes is eventually consistent.
type RemoveAttributeDefinition struct {
	// The name of the attribute to remove.
	Name string `json:"name"`
}

// MarshalJSON override to add the action value
func (ua RemoveAttributeDefinition) MarshalJSON() ([]byte, error) {
	type Alias RemoveAttributeDefinition

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "removeAttributeDefinition",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeAttributeDefinitionName will change the name of an attribute definition.
// This action also renames all corresponding attributes on all Products with this product type.
// The renaming of the attributes is eventually consistent.
// warning:
// When renaming several attributes at the same time beware of the following effects on updates
// on the affected products:
// - let’s say we have a product type with 2 attribute definitions: aa and bb.
// - we send a first command to update bb to cc.
// - and then we send a second command to update aa to bb.
// - when updating the products the order is not guaranteed as they are 2 separated commands.
//   It’s possible (even if unlikely) that the products are updated in the wrong order:
//   the attribute aa is renamed to bb, and then to cc.
// To avoid that, send multiple "changeAttributeName" updates in one command.
type ChangeAttributeDefinitionName struct {
	// The name of the attribute definition to update.
	AttributeName string `json:"attributeName"`
	// The new name attribute definition name.
	NewAttributeName string `json:"newAttributeName"`
}

// MarshalJSON override to add the action value
func (ua ChangeAttributeDefinitionName) MarshalJSON() ([]byte, error) {
	type Alias ChangeAttributeDefinitionName

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeAttributeName",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeAttributeDefinitionLabel will change the label of an attribute definition.
type ChangeAttributeDefinitionLabel struct {
	// The name of the attribute definition to update.
	AttributeName string                        `json:"attributeName"`
	Label         commercetools.LocalizedString `json:"label"`
}

// MarshalJSON override to add the action value
func (ua ChangeAttributeDefinitionLabel) MarshalJSON() ([]byte, error) {
	type Alias ChangeAttributeDefinitionLabel

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeLabel",
		Alias:  (*Alias)(&ua),
	})
}

// SetAttributeDefinitionInputTip will set the input tip on an attribute definition.
// Allows to set additional information about the specified attribute that aids
// content managers when setting product details.
type SetAttributeDefinitionInputTip struct {
	AttributeName string                        `json:"attributeName"`
	InputTip      commercetools.LocalizedString `json:"inputTip,omitempty"`
}

// MarshalJSON override to add the action value
func (ua SetAttributeDefinitionInputTip) MarshalJSON() ([]byte, error) {
	type Alias SetAttributeDefinitionInputTip

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setInputTip",
		Alias:  (*Alias)(&ua),
	})
}

// AddPlainEnumValue will add a plain enum value to an enum type attribute definition.
// It can update an EnumType attribute definition or a Set of EnumType attribute definition.
type AddPlainEnumValue struct {
	// The name of the attribute definition to update.
	AttributeName string                  `json:"attributeName"`
	Value         commercetools.EnumValue `json:"value"`
}

// MarshalJSON override to add the action value
func (ua AddPlainEnumValue) MarshalJSON() ([]byte, error) {
	type Alias AddPlainEnumValue

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "addPlainEnumValue",
		Alias:  (*Alias)(&ua),
	})
}

// AddLocalizedEnumValue will add a localized enum value to a localizd enum type attribute definition.
// It can update a LocalizableEnum attribute definition or a Set of LocalizableEnumType attribute definition.
type AddLocalizedEnumValue struct {
	AttributeName string                           `json:"attributeName"`
	Value         commercetools.LocalizedEnumValue `json:"value"`
}

// MarshalJSON override to add the action value
func (ua AddLocalizedEnumValue) MarshalJSON() ([]byte, error) {
	type Alias AddLocalizedEnumValue

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "addLocalizedEnumValue",
		Alias:  (*Alias)(&ua),
	})
}

// RemoveEnumValues will remove an enum value of an attribute definition.
// Removes enum values from an attribute definition of EnumType or LocalizableEnumType
// or Set of EnumType or Set of LocalizableEnum.
// All attributes of all products using those enum keys will also be removed in an eventually consistent way.
type RemoveEnumValues struct {
	// The name of the attribute definition to update.
	AttributeName string `json:"attributeName"`
	// The keys of the EnumType or LocalizableEnumType to remove.
	Keys []string `json:"keys"`
}

// MarshalJSON override to add the action value
func (ua RemoveEnumValues) MarshalJSON() ([]byte, error) {
	type Alias RemoveEnumValues

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "removeEnumValues",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeAttributeDefinitionsOrder will reorder the attributes on the product type being updated.
type ChangeAttributeDefinitionsOrder struct {
	// The attributes must be equal to the product type attributes (except for the order).
	Attributes []AttributeDefinition `json:"attributes"`
}

// MarshalJSON override to add the action value
func (ua ChangeAttributeDefinitionsOrder) MarshalJSON() ([]byte, error) {
	type Alias ChangeAttributeDefinitionsOrder

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeAttributeOrder",
		Alias:  (*Alias)(&ua),
	})
}

// ChangePlainEnumValuesOrder will reorder the plain enums on an attribute definition.
// It can update an EnumType attribute definition or a Set of EnumType attribute definition.
type ChangePlainEnumValuesOrder struct {
	AttributeName string `json:"attributeName"`
	// The values must be equal to the values of the attribute enum values (except for the order).
	// If not a EnumValuesMustMatch error code will be returned.
	Values []commercetools.EnumValue `json:"values"`
}

// MarshalJSON override to add the action value
func (ua ChangePlainEnumValuesOrder) MarshalJSON() ([]byte, error) {
	type Alias ChangePlainEnumValuesOrder

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changePlainEnumValueOrder",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeLocalizedEnumValuesOrder will reorder the localized enums on an attribute definition.
// It can update a LocalizableEnumType attribute definition or a Set of LocalizableEnumType attribute definition.
type ChangeLocalizedEnumValuesOrder struct {
	AttributeName string `json:"attributeName"`
	// The values must be equal to the values of the attribute enum values (except for the order).
	// If not a EnumValuesMustMatch error code will be returned.
	Values []commercetools.LocalizedEnumValue `json:"values"`
}

// MarshalJSON override to add the action value
func (ua ChangeLocalizedEnumValuesOrder) MarshalJSON() ([]byte, error) {
	type Alias ChangeLocalizedEnumValuesOrder

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeLocalizedEnumValueOrder",
		Alias:  (*Alias)(&ua),
	})
}

// ChangePlainEnumValueKey will change the key of an enum value.
// This action changes the key of a single enum value in an EnumType or LocalizableEnumType attribute definition.
// It can update an EnumType attribute definition or a LocalizableEnumType attribute definition or a
// Set of EnumType attribute definition or a Set of LocalizableEnumType attribute definition.
// All products will be updated to the new key in an eventually consistent way.
type ChangePlainEnumValueKey struct {
	AttributeName string `json:"attributeName"`
	Key           string `json:"key"`
	NewKey        string `json:"newKey"`
}

// MarshalJSON override to add the action value
func (ua ChangePlainEnumValueKey) MarshalJSON() ([]byte, error) {
	type Alias ChangePlainEnumValueKey

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeEnumKey",
		Alias:  (*Alias)(&ua),
	})
}

// ChangePlainEnumValueLabel will change the label of an enum value.
// It can update an EnumType attribute definition or a Set of EnumType attribute definition.
// All products will be updated to the new label in an eventually consistent way.
type ChangePlainEnumValueLabel struct {
	AttributeName string `json:"attributeName"`
	// The new value must be different from the existing value.
	NewValue commercetools.EnumValue `json:"newValue"`
}

// MarshalJSON override to add the action value
func (ua ChangePlainEnumValueLabel) MarshalJSON() ([]byte, error) {
	type Alias ChangePlainEnumValueLabel

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changePlainEnumValueLabel",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeLocalizedEnumValueLabel will change the label of a localized enum value.
// It can update a LocalizableEnumType attribute definition or a Set of LocalizableEnumType attribute definition.
// All products will be updated to the new label in an eventually consistent way.
type ChangeLocalizedEnumValueLabel struct {
	AttributeName string `json:"attributeName"`
	// The new value must be different from the existing value.
	NewValue commercetools.LocalizedEnumValue `json:"newValue"`
}

// MarshalJSON override to add the action value
func (ua ChangeLocalizedEnumValueLabel) MarshalJSON() ([]byte, error) {
	type Alias ChangeLocalizedEnumValueLabel

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeLocalizedEnumValueLabel",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeIsSearchable changes whether an attribute definition is searchable.
// Following this update the products are reindexed asynchronously to reflect this change on the search endpoint.
// When enabling search on an existing attribute type definition, the constraint regarding the maximum size
// of a searchable attribute will not be enforced. Instead, product attribute definitions exceeding this limit
// will be treated as not searchable and will not be available for full-text search.
type ChangeIsSearchable struct {
	// The name of the attribute definition to update.
	AttributeName string `json:"attributeName"`
	IsSearchable  bool   `json:"isSearchable"`
}

// MarshalJSON override to add the action value
func (ua ChangeIsSearchable) MarshalJSON() ([]byte, error) {
	type Alias ChangeIsSearchable

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeIsSearchable",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeInputHint Changes the input hint on an attribute definition.
type ChangeInputHint struct {
	// The name of the attribute definition to update.
	// Must be of TextType, LocalizableTextType or SetType of TextType/LocalizableTextType.
	AttributeName string                      `json:"attributeName"`
	NewValue      commercetools.TextInputHint `json:"newValue"`
}

// MarshalJSON override to add the action value
func (ua ChangeInputHint) MarshalJSON() ([]byte, error) {
	type Alias ChangeInputHint

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeInputHint",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeAttributeConstraint changes the attribute constraint on an attribute definition.
// For now only following changes are supported: SameForAll to None and Unique to None.
type ChangeAttributeConstraint struct {
	// The name of the attribute definition to update.
	AttributeName string `json:"attributeName"`
	// For now only None is supported.
	NewValue AttributeConstraint `json:"newValue"`
}

// MarshalJSON override to add the action value
func (ua ChangeAttributeConstraint) MarshalJSON() ([]byte, error) {
	type Alias ChangeAttributeConstraint

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeAttributeConstraint",
		Alias:  (*Alias)(&ua),
	})
}
