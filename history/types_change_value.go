package history

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

type ChangeTargetChangeValue interface{}

func mapDiscriminatorChangeTargetChangeValue(input interface{}) (ChangeTargetChangeValue, error) {
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
	case "customLineItems":
		obj := ChangeTargetCustomLineItemsChangeValue{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "lineItems":
		obj := ChangeTargetLineItemsChangeValue{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "multiBuyCustomLineItems":
		obj := ChangeTargetMultiBuyCustomLineItemsChangeValue{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "multiBuyLineItems":
		obj := ChangeTargetMultiBuyLineItemsChangeValue{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "shipping":
		obj := ChangeTargetShippingChangeValue{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ChangeValueChangeValue interface{}

func mapDiscriminatorChangeValueChangeValue(input interface{}) (ChangeValueChangeValue, error) {
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
	case "absolute":
		obj := ChangeValueAbsoluteChangeValue{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "external":
		obj := ChangeValueExternalChangeValue{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "giftLineItem":
		obj := ChangeValueGiftLineItemChangeValue{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "relative":
		obj := ChangeValueRelativeChangeValue{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type AssetChangeValue struct {
	// `id` of the [Asset](ctp:api:type:Asset).
	ID string `json:"id"`
	// Name of the Asset.
	Name LocalizedString `json:"name"`
}

type AttributeValue struct {
	// Name of the Attribute set.
	Name string `json:"name"`
	// Value set for the Attribute determined by the [AttributeType](ctp:api:type:AttributeType):
	//
	// - For [Enum Type](ctp:api:type:AttributeEnumType) and [Localized Enum Type](ctp:api:type:AttributeLocalizedEnumType), `value` is the `key` of the [Plain Enum Value](ctp:api:type:AttributePlainEnumValue) or [Localized Enum Value](ctp:api:type:AttributeLocalizedEnumValue) objects,
	//   or the complete objects.
	// - For [Localizable Text Type](ctp:api:type:AttributeLocalizableTextType), `value` is the [LocalizedString](ctp:api:type:LocalizedString) object.
	// - For [Money Type](ctp:api:type:AttributeMoneyType) Attributes, `value` is the [Money](ctp:api:type:Money) object.
	// - For [Set Type](ctp:api:type:AttributeSetType) Attributes, `value` is the entire `set` object.
	// - For [Nested Type](ctp:api:type:AttributeNestedType) Attributes, `value` is the list of values of all Attributes of the nested Product.
	// - For [Reference Type](ctp:api:type:AttributeReferenceType) Attributes, `value` is the [Reference](ctp:api:type:Reference) object.
	Value interface{} `json:"value"`
}

type ChangeTargetCustomLineItemsChangeValue struct {
	// Valid [CustomLineItem target predicate](/../api/projects/predicates#customlineitem-field-identifiers).
	Predicate string `json:"predicate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeTargetCustomLineItemsChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeTargetCustomLineItemsChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "customLineItems", Alias: (*Alias)(&obj)})
}

type ChangeTargetLineItemsChangeValue struct {
	// Valid [LineItem target predicate](/../api/projects/predicates#lineitem-field-identifiers).
	Predicate string `json:"predicate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeTargetLineItemsChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeTargetLineItemsChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "lineItems", Alias: (*Alias)(&obj)})
}

type ChangeTargetMultiBuyCustomLineItemsChangeValue struct {
	// Valid [CustomLineItem target predicate](/../api/projects/predicates#customlineitem-field-identifiers).
	Predicate string `json:"predicate"`
	// Quantity of Custom Line Items that triggered the application of the discount.
	TriggerQuantity int `json:"triggerQuantity"`
	// Quantity of Custom Line Items discounted per application of this discount.
	DiscountedQuantity int `json:"discountedQuantity"`
	// Maximum number of times the discount is applicable.
	MaxOccurrence int `json:"maxOccurrence"`
	// SelectionMode based on which particular Custom Line Items were discounted.
	SelectionMode SelectionMode `json:"selectionMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeTargetMultiBuyCustomLineItemsChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeTargetMultiBuyCustomLineItemsChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "multiBuyCustomLineItems", Alias: (*Alias)(&obj)})
}

type ChangeTargetMultiBuyLineItemsChangeValue struct {
	// Valid [LineItem target predicate](/../api/projects/predicates#lineitem-field-identifiers).
	Predicate string `json:"predicate"`
	// Quantity of Line Items that triggered the application of the discount.
	TriggerQuantity int `json:"triggerQuantity"`
	// Quantity of Line Items discounted per application of this discount.
	DiscountedQuantity int `json:"discountedQuantity"`
	// Maximum number of times the discount is applicable.
	MaxOccurrence int `json:"maxOccurrence"`
	// SelectionMode based on which particular Line Items were discounted.
	SelectionMode SelectionMode `json:"selectionMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeTargetMultiBuyLineItemsChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeTargetMultiBuyLineItemsChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "multiBuyLineItems", Alias: (*Alias)(&obj)})
}

type ChangeTargetShippingChangeValue struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeTargetShippingChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeTargetShippingChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "shipping", Alias: (*Alias)(&obj)})
}

type ChangeValueAbsoluteChangeValue struct {
	// Money values in different currencies.
	Money []Money `json:"money"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeValueAbsoluteChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeValueAbsoluteChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "absolute", Alias: (*Alias)(&obj)})
}

type ChangeValueExternalChangeValue struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeValueExternalChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeValueExternalChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "external", Alias: (*Alias)(&obj)})
}

type ChangeValueGiftLineItemChangeValue struct {
	// Reference to a [Product](ctp:api:type:Product).
	Product Reference `json:"product"`
	// `id` of the [ProductVariant](ctp:api:type:ProductVariant).
	VariantId int `json:"variantId"`
	// Channel with [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum) `InventorySupply`.
	SupplyChannel *Reference `json:"supplyChannel,omitempty"`
	// Channel with [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum) `ProductDistribution`.
	DistributionChannel Reference `json:"distributionChannel"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeValueGiftLineItemChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeValueGiftLineItemChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "giftLineItem", Alias: (*Alias)(&obj)})
}

type ChangeValueRelativeChangeValue struct {
	// Fraction (per ten thousand) the price is reduced by. For example, 1000 results in a 10% price reduction.
	Permyriad int `json:"permyriad"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ChangeValueRelativeChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeValueRelativeChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "relative", Alias: (*Alias)(&obj)})
}

/**
*	Only present if `expand` is set to `true`.
 */
type CustomFieldExpandedValue struct {
	// Name of the Custom Field.
	Name string `json:"name"`
	// [CustomFieldValue](ctp:api:type:CustomFieldValue) based on the [FieldType](ctp:api:type:FieldType).
	Value interface{} `json:"value"`
	// User-defined label of the Custom Field.
	Label LocalizedString `json:"label"`
}

type CustomShippingMethodChangeValue struct {
	// Name of the Custom ShippingMethod.
	Name string `json:"name"`
}

type DeliveryChangeValue struct {
	// Line Items or Custom Line Items shipped in the [Delivery](ctp:api:type:Delivery).
	Items []DeliveryItem `json:"items"`
	// Address to which the parcels are delivered.
	Address Address `json:"address"`
	// Parcels included in the [Delivery](ctp:api:type:Delivery).
	Parcels []Parcel `json:"parcels"`
}

type EnumValue struct {
	// Key of the value used as a programmatic identifier.
	Key string `json:"key"`
	// Descriptive label of the value.
	Label string `json:"label"`
}

type FieldDefinitionOrderValue struct {
	// Name of the [FieldDefinition](ctp:api:type:FieldDefinition).
	Name string `json:"name"`
	// Descriptive label of the field.
	Label LocalizedString `json:"label"`
}

type InventoryQuantityValue struct {
	// Overall amount of stock (`availableQuantity` + reserved).
	QuantityOnStock int `json:"quantityOnStock"`
	// Available amount of stock (`quantityOnStock` - reserved).
	AvailableQuantity int `json:"availableQuantity"`
}

type LocalizedEnumValue struct {
	// Key of the value used as a programmatic identifier.
	Key string `json:"key"`
	// Descriptive localized label of the value.
	Label LocalizedString `json:"label"`
}

type ParcelChangeValue struct {
	// `id` of the [Parcel](ctp:api:type:Parcel).
	ID string `json:"id"`
	// Date and time (UTC) the Parcel was created.
	CreatedAt string `json:"createdAt"`
}

type SetCartClassificationShippingRateInputValue struct {
	Type string `json:"type"`
	// Key of the value used as a programmatic identifier.
	Key string `json:"key"`
	// Descriptive localized label of the value.
	Label LocalizedString `json:"label"`
}

type SetCartScoreShippingRateInputValue struct {
	Type string `json:"type"`
	// Abstract value for categorizing a Cart.
	Score int `json:"score"`
}

type ShippingMethodChangeValue struct {
	// `id` of the [ShippingMethod](ctp:api:type:ShippingMethod).
	ID string `json:"id"`
	// Name of the ShippingMethod.
	Name string `json:"name"`
}

type ShippingMethodTaxAmountChangeValue struct {
	// Taxed price for the Shipping Method based on `taxRate`.
	TaxedPrice TaxedPrice `json:"taxedPrice"`
	// Tax rate set externally for the Shipping Method.
	TaxRate TaxRate `json:"taxRate"`
}

type ShoppingListLineItemValue struct {
	// `id` of the [ShoppingListLineItem](ctp:api:type:ShoppingListLineItem).
	ID string `json:"id"`
	// Name of the corresponding Product the Product Variant belongs to.
	Name LocalizedString `json:"name"`
	// `id` of the [ProductVariant](ctp:api:type:ProductVariant) the ShoppingListLineItem refers to.
	VariantId int `json:"variantId"`
}

type TextLineItemValue struct {
	// `id` of the [TextLineItem](ctp:api:type:TextLineItem).
	ID string `json:"id"`
	// Name of the TextLineItem.
	Name LocalizedString `json:"name"`
}

type TransactionChangeValue struct {
	// `id` of the [Transaction](ctp:api:type:Transaction).
	ID string `json:"id"`
	// Identifier used by the interface that manages the Transaction (usually the PSP).
	InteractionId string `json:"interactionId"`
	// Date and time (UTC) the Transaction took place.
	Timestamp string `json:"timestamp"`
}

type ValidFromAndUntilValue struct {
	// Date and time (UTC) from when the Discount is effective.
	ValidFrom string `json:"validFrom"`
	// Date and time (UTC) until when the Discount is effective.
	ValidUntil string `json:"validUntil"`
}
