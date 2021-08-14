// Generated file, please do not change!!!
package history

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
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
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
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
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
	ID   string          `json:"id"`
	Name LocalizedString `json:"name"`
}

/**
*	Shape of the value for cart discounts line item and custom line items target.
 */
type ChangeTargetCustomLineItemsChangeValue struct {
	Predicate string `json:"predicate"`
}

// MarshalJSON override to set the discriminator value
func (obj ChangeTargetCustomLineItemsChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeTargetCustomLineItemsChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "customLineItems", Alias: (*Alias)(&obj)})
}

/**
*	Shape of the value for cart discounts line item target.
 */
type ChangeTargetLineItemsChangeValue struct {
	Predicate string `json:"predicate"`
}

// MarshalJSON override to set the discriminator value
func (obj ChangeTargetLineItemsChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeTargetLineItemsChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "lineItems", Alias: (*Alias)(&obj)})
}

/**
*	Shape of the value for cart discounts multiBuyCustomLineItems target.
 */
type ChangeTargetMultiBuyCustomLineItemsChangeValue struct {
	Predicate string `json:"predicate"`
	// Quantity of line items that need to be present in order to trigger an application of this discount.
	TriggerQuantity int `json:"triggerQuantity"`
	// Quantity of line items that are discounted per application of this discount.
	DiscountedQuantity int `json:"discountedQuantity"`
	// Maximum number of applications of this discount.
	MaxOccurrence int           `json:"maxOccurrence"`
	SelectionMode SelectionMode `json:"selectionMode"`
}

// MarshalJSON override to set the discriminator value
func (obj ChangeTargetMultiBuyCustomLineItemsChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeTargetMultiBuyCustomLineItemsChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "multiBuyCustomLineItems", Alias: (*Alias)(&obj)})
}

/**
*	Shape of the value for cart discounts multiBuyLineItems target.
 */
type ChangeTargetMultiBuyLineItemsChangeValue struct {
	Predicate string `json:"predicate"`
	// Quantity of line items that need to be present in order to trigger an application of this discount.
	TriggerQuantity int `json:"triggerQuantity"`
	// Quantity of line items that are discounted per application of this discount.
	DiscountedQuantity int `json:"discountedQuantity"`
	// Maximum number of applications of this discount.
	MaxOccurrence int           `json:"maxOccurrence"`
	SelectionMode SelectionMode `json:"selectionMode"`
}

// MarshalJSON override to set the discriminator value
func (obj ChangeTargetMultiBuyLineItemsChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeTargetMultiBuyLineItemsChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "multiBuyLineItems", Alias: (*Alias)(&obj)})
}

/**
*	Shape of the value for cart discounts shipping target.
 */
type ChangeTargetShippingChangeValue struct {
}

// MarshalJSON override to set the discriminator value
func (obj ChangeTargetShippingChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeTargetShippingChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "shipping", Alias: (*Alias)(&obj)})
}

/**
*	Shape of the value for cart discounts absolute value.
 */
type ChangeValueAbsoluteChangeValue struct {
	Money []Money `json:"money"`
}

// MarshalJSON override to set the discriminator value
func (obj ChangeValueAbsoluteChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeValueAbsoluteChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "absolute", Alias: (*Alias)(&obj)})
}

/**
*	Shape of the value for product discounts external value.
 */
type ChangeValueExternalChangeValue struct {
}

// MarshalJSON override to set the discriminator value
func (obj ChangeValueExternalChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeValueExternalChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "external", Alias: (*Alias)(&obj)})
}

/**
*	Shape of the value for cart discounts gift line item value.
 */
type ChangeValueGiftLineItemChangeValue struct {
	Product             Reference  `json:"product"`
	VariantId           int        `json:"variantId"`
	SupplyChannel       *Reference `json:"supplyChannel,omitempty"`
	DistributionChannel Reference  `json:"distributionChannel"`
}

// MarshalJSON override to set the discriminator value
func (obj ChangeValueGiftLineItemChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeValueGiftLineItemChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "giftLineItem", Alias: (*Alias)(&obj)})
}

/**
*	Shape of the value for cart discounts relative value.
 */
type ChangeValueRelativeChangeValue struct {
	Permyriad int `json:"permyriad"`
}

// MarshalJSON override to set the discriminator value
func (obj ChangeValueRelativeChangeValue) MarshalJSON() ([]byte, error) {
	type Alias ChangeValueRelativeChangeValue
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "relative", Alias: (*Alias)(&obj)})
}

/**
*	Only available if `expand` is set to true
 */
type CustomFieldExpandedValue struct {
	// Name of a custom field.
	Name  string          `json:"name"`
	Value interface{}     `json:"value"`
	Label LocalizedString `json:"label"`
}

type CustomShippingMethodChangeValue struct {
	Name string `json:"name"`
}

type DeliveryChangeValue struct {
	Items   []DeliveryItem `json:"items"`
	Address Address        `json:"address"`
	Parcels []Parcel       `json:"parcels"`
}

type EnumValue struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

/**
*	Shape of the value for action `changeFieldDefinitionOrder`
 */
type FieldDefinitionOrderValue struct {
	Name  string          `json:"name"`
	Label LocalizedString `json:"label"`
}

type InventoryQuantityValue struct {
	QuantityOnStock   int `json:"quantityOnStock"`
	AvailableQuantity int `json:"availableQuantity"`
}

type LocalizedEnumValue struct {
	Key   string          `json:"key"`
	Label LocalizedString `json:"label"`
}

type ParcelChangeValue struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

/**
*	Shape of the cart classification shipping input rate value.
 */
type SetCartClassificationShippingRateInputValue struct {
	Type  string          `json:"type"`
	Key   string          `json:"key"`
	Label LocalizedString `json:"label"`
}

/**
*	Shape of the cart score shipping input rate value.
 */
type SetCartScoreShippingRateInputValue struct {
	Type  string `json:"type"`
	Score int    `json:"score"`
}

type ShippingMethodChangeValue struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ShoppingListLineItemValue struct {
	ID        string          `json:"id"`
	Name      LocalizedString `json:"name"`
	VariantId int             `json:"variantId"`
}

type TextLineItemValue struct {
	ID   string          `json:"id"`
	Name LocalizedString `json:"name"`
}

type TransactionChangeValue struct {
	ID            string `json:"id"`
	InteractionId string `json:"interactionId"`
	Timestamp     string `json:"timestamp"`
}

/**
*	Shape of the value for `setValidFromAndUntil` action
 */
type ValidFromAndUntilValue struct {
	ValidFrom  string `json:"validFrom"`
	ValidUntil string `json:"validUntil"`
}
