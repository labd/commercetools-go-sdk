package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	The item's state.
*
 */
type ItemState struct {
	Quantity float64 `json:"quantity"`
	// Maps to `ItemState.state`.
	State StateKeyReference `json:"state"`
}

/**
*	The item's shipping target.
*
 */
type ItemShippingTarget struct {
	// Maps to `ItemShippingTarget.addressKey`.
	AddressKey string `json:"addressKey"`
	// Maps to `ItemShippingTarget.quantity`.
	Quantity float64 `json:"quantity"`
}

type ItemShippingDetailsDraft struct {
	// Maps to `ItemShippingDetailsDraft.targets`.
	Targets []ItemShippingTarget `json:"targets"`
}

type LineItemPrice struct {
	// Maps to `Price.value`.
	Value TypedMoney `json:"value"`
	// Maps to `Price.county`.
	Country *string `json:"country,omitempty"`
	// Maps to `Price.validFrom`.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Maps to `Price.validUntil`.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// References a customer group by key.
	CustomerGroup *CustomerGroupKeyReference `json:"customerGroup,omitempty"`
	// References a channel by key.
	Channel *ChannelKeyReference `json:"channel,omitempty"`
	// Sets a discounted price from an external service.
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	// The tiered prices for this price.
	Tiers []PriceTier `json:"tiers"`
	// Maps to `Price.custom`.
	Custom *Custom `json:"custom,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *LineItemPrice) UnmarshalJSON(data []byte) error {
	type Alias LineItemPrice
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorTypedMoney(obj.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LineItemPrice) MarshalJSON() ([]byte, error) {
	type Alias LineItemPrice
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

	if raw["tiers"] == nil {
		delete(raw, "tiers")
	}

	return json.Marshal(raw)

}

type LineItemProductVariantImportDraft struct {
	// Maps to `ProductVariant.product`.
	ProductVariant *ProductVariantKeyReference `json:"productVariant,omitempty"`
	// Maps to `ProductVariantImportDraft.sku`.
	Sku *string `json:"sku,omitempty"`
	// Maps to `ProductVariantImportDraft.prices`
	Prices []LineItemPrice `json:"prices"`
	// Maps to `ProductVariantImportDraft.attributes`
	Attributes []Attribute `json:"attributes"`
	// Maps to `ProductVariantImportDraft.images`.
	Images []Image `json:"images"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *LineItemProductVariantImportDraft) UnmarshalJSON(data []byte) error {
	type Alias LineItemProductVariantImportDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Attributes {
		var err error
		obj.Attributes[i], err = mapDiscriminatorAttribute(obj.Attributes[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LineItemProductVariantImportDraft) MarshalJSON() ([]byte, error) {
	type Alias LineItemProductVariantImportDraft
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

	if raw["prices"] == nil {
		delete(raw, "prices")
	}

	if raw["attributes"] == nil {
		delete(raw, "attributes")
	}

	if raw["images"] == nil {
		delete(raw, "images")
	}

	return json.Marshal(raw)

}

/**
*	Represents an individual Line Item in an Order. A line item is a snapshot of a product at the time it was added to the order.
*
*	You cannot create an Order that includes line item operations that do not exist in the Project or have been deleted.
*	Products and Product Variants referenced by a line item must already exist in the Project.
*
 */
type LineItemImportDraft struct {
	// Maps to `LineItem.productId`.
	Product *ProductKeyReference `json:"product,omitempty"`
	// Maps to `LineItem.name`.
	Name LocalizedString `json:"name"`
	// Maps to `ProductVariantImportDraft`.
	Variant LineItemProductVariantImportDraft `json:"variant"`
	// Maps to `LineItem.price`.
	Price LineItemPrice `json:"price"`
	// Maps to `LineItem.quantity`.
	Quantity float64     `json:"quantity"`
	State    []ItemState `json:"state"`
	// Maps to `LineItem.supplyChannel`.
	// The Reference to the Supply [Channel](/../api/projects/channels#channel) with which the LineItem is associated.
	// If referenced Supply Channel does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary Supply Channel is created.
	SupplyChannel *ChannelKeyReference `json:"supplyChannel,omitempty"`
	// Maps to `LineItem.distributionChannel`.
	// The Reference to the Distribution [Channel](/../api/projects/channels#channel) with which the LineItem is associated.
	// If referenced CustomerGroup does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary Distribution Channel is created.
	DistributionChannel *ChannelKeyReference `json:"distributionChannel,omitempty"`
	// Maps to `LineItem.taxRate`.
	TaxRate *TaxRate `json:"taxRate,omitempty"`
	// Maps to LineItem.shippingDetails.
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	// Custom Fields for this Line Item.
	Custom *Custom `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LineItemImportDraft) MarshalJSON() ([]byte, error) {
	type Alias LineItemImportDraft
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

	if raw["state"] == nil {
		delete(raw, "state")
	}

	return json.Marshal(raw)

}

type ShippingRateTierType string

const (
	ShippingRateTierTypeCartValue          ShippingRateTierType = "CartValue"
	ShippingRateTierTypeCartClassification ShippingRateTierType = "CartClassification"
	ShippingRateTierTypeCartScore          ShippingRateTierType = "CartScore"
)

type ShippingRatePriceTier interface{}

func mapDiscriminatorShippingRatePriceTier(input interface{}) (ShippingRatePriceTier, error) {
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
	case "CartClassification":
		obj := CartClassificationTier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		for i := range obj.Tiers {
			var err error
			obj.Tiers[i], err = mapDiscriminatorShippingRatePriceTier(obj.Tiers[i])
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	}
	return nil, nil
}

type CartClassificationTier struct {
	Value      string                  `json:"value"`
	Price      Money                   `json:"price"`
	Tiers      []ShippingRatePriceTier `json:"tiers"`
	IsMatching *bool                   `json:"isMatching,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartClassificationTier) UnmarshalJSON(data []byte) error {
	type Alias CartClassificationTier
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Tiers {
		var err error
		obj.Tiers[i], err = mapDiscriminatorShippingRatePriceTier(obj.Tiers[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartClassificationTier) MarshalJSON() ([]byte, error) {
	type Alias CartClassificationTier
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CartClassification", Alias: (*Alias)(&obj)})
}

type ShippingRateDraft struct {
	Price     Money                   `json:"price"`
	FreeAbove *Money                  `json:"freeAbove,omitempty"`
	Tiers     []ShippingRatePriceTier `json:"tiers"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingRateDraft) UnmarshalJSON(data []byte) error {
	type Alias ShippingRateDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Tiers {
		var err error
		obj.Tiers[i], err = mapDiscriminatorShippingRatePriceTier(obj.Tiers[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingRateDraft) MarshalJSON() ([]byte, error) {
	type Alias ShippingRateDraft
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

	if raw["tiers"] == nil {
		delete(raw, "tiers")
	}

	return json.Marshal(raw)

}

type ParcelMeasurements struct {
	HeightInMillimeter *float64 `json:"heightInMillimeter,omitempty"`
	LengthInMillimeter *float64 `json:"lengthInMillimeter,omitempty"`
	WidthInMillimeter  *float64 `json:"widthInMillimeter,omitempty"`
	WeightInGram       *float64 `json:"weightInGram,omitempty"`
}

type TrackingData struct {
	TrackingId          *string `json:"trackingId,omitempty"`
	Carrier             *string `json:"carrier,omitempty"`
	Provider            *string `json:"provider,omitempty"`
	ProviderTransaction *string `json:"providerTransaction,omitempty"`
	IsReturn            *bool   `json:"isReturn,omitempty"`
}

type DeliveryItem struct {
	ID       string  `json:"id"`
	Quantity float64 `json:"quantity"`
}

type Parcel struct {
	ID           string              `json:"id"`
	CreatedAt    time.Time           `json:"createdAt"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	TrackingData *TrackingData       `json:"trackingData,omitempty"`
	Items        []DeliveryItem      `json:"items"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Parcel) MarshalJSON() ([]byte, error) {
	type Alias Parcel
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

	if raw["items"] == nil {
		delete(raw, "items")
	}

	return json.Marshal(raw)

}

type Delivery struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	Items     []DeliveryItem `json:"items"`
	Parcels   []Parcel       `json:"parcels"`
	Address   *Address       `json:"address,omitempty"`
}

type DiscountedLineItemPortion struct {
	// References a cart discount by key.
	Discount         CartDiscountKeyReference `json:"discount"`
	DiscountedAmount Money                    `json:"discountedAmount"`
}

type DiscountedLineItemPriceDraft struct {
	Value             Money                       `json:"value"`
	IncludedDiscounts []DiscountedLineItemPortion `json:"includedDiscounts"`
}

type ShippingMethodState string

const (
	ShippingMethodStateDoesNotMatchCart ShippingMethodState = "DoesNotMatchCart"
	ShippingMethodStateMatchesCart      ShippingMethodState = "MatchesCart"
)

/**
*	Maps to an order's `shippingInfo` property. This field is usually populated by the cart assosciated with
*	the order, but when importing orders you must provide a draft representation as a part of the OrderImport.
*
 */
type ShippingInfoImportDraft struct {
	ShippingMethodName string            `json:"shippingMethodName"`
	Price              TypedMoney        `json:"price"`
	ShippingRate       ShippingRateDraft `json:"shippingRate"`
	TaxRate            *TaxRate          `json:"taxRate,omitempty"`
	// References a tax category by key.
	TaxCategory *TaxCategoryKeyReference `json:"taxCategory,omitempty"`
	// References a shipping method by key.
	ShippingMethod *ShippingMethodKeyReference `json:"shippingMethod,omitempty"`
	// Note that you can not add a `DeliveryItem` on import, as `LineItems` and `CustomLineItems` are not yet referencable by an `id`.
	Deliveries          []Delivery                    `json:"deliveries"`
	DiscountedPrice     *DiscountedLineItemPriceDraft `json:"discountedPrice,omitempty"`
	ShippingMethodState *ShippingMethodState          `json:"shippingMethodState,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingInfoImportDraft) UnmarshalJSON(data []byte) error {
	type Alias ShippingInfoImportDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Price != nil {
		var err error
		obj.Price, err = mapDiscriminatorTypedMoney(obj.Price)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingInfoImportDraft) MarshalJSON() ([]byte, error) {
	type Alias ShippingInfoImportDraft
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

	if raw["deliveries"] == nil {
		delete(raw, "deliveries")
	}

	return json.Marshal(raw)

}

type ExternalTaxRateDraft struct {
	Name            string    `json:"name"`
	Amount          *float64  `json:"amount,omitempty"`
	Country         string    `json:"country"`
	State           *string   `json:"state,omitempty"`
	SubRates        []SubRate `json:"subRates"`
	IncludedInPrice *bool     `json:"includedInPrice,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ExternalTaxRateDraft) MarshalJSON() ([]byte, error) {
	type Alias ExternalTaxRateDraft
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

	if raw["subRates"] == nil {
		delete(raw, "subRates")
	}

	return json.Marshal(raw)

}

type CustomLineItemTaxedPrice struct {
	TotalNet   TypedMoney `json:"totalNet"`
	TotalGross TypedMoney `json:"totalGross"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomLineItemTaxedPrice) UnmarshalJSON(data []byte) error {
	type Alias CustomLineItemTaxedPrice
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.TotalNet != nil {
		var err error
		obj.TotalNet, err = mapDiscriminatorTypedMoney(obj.TotalNet)
		if err != nil {
			return err
		}
	}
	if obj.TotalGross != nil {
		var err error
		obj.TotalGross, err = mapDiscriminatorTypedMoney(obj.TotalGross)
		if err != nil {
			return err
		}
	}

	return nil
}

type CustomLineItemDraft struct {
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	Name       LocalizedString           `json:"name"`
	Money      TypedMoney                `json:"money"`
	TaxedPrice *CustomLineItemTaxedPrice `json:"taxedPrice,omitempty"`
	TotalPrice TypedMoney                `json:"totalPrice"`
	Slug       string                    `json:"slug"`
	Quantity   float64                   `json:"quantity"`
	State      []ItemState               `json:"state"`
	// References a tax category by key.
	TaxCategory                *TaxCategoryKeyReference       `json:"taxCategory,omitempty"`
	TaxRate                    *TaxRate                       `json:"taxRate,omitempty"`
	ExternalTaxRate            *ExternalTaxRateDraft          `json:"externalTaxRate,omitempty"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceDraft `json:"discountedPricePerQuantity"`
	ShippingDetails            *ItemShippingDetailsDraft      `json:"shippingDetails,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomLineItemDraft) UnmarshalJSON(data []byte) error {
	type Alias CustomLineItemDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Money != nil {
		var err error
		obj.Money, err = mapDiscriminatorTypedMoney(obj.Money)
		if err != nil {
			return err
		}
	}
	if obj.TotalPrice != nil {
		var err error
		obj.TotalPrice, err = mapDiscriminatorTypedMoney(obj.TotalPrice)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomLineItemDraft) MarshalJSON() ([]byte, error) {
	type Alias CustomLineItemDraft
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

	if raw["state"] == nil {
		delete(raw, "state")
	}

	if raw["discountedPricePerQuantity"] == nil {
		delete(raw, "discountedPricePerQuantity")
	}

	return json.Marshal(raw)

}

type TaxPortion struct {
	Name   *string    `json:"name,omitempty"`
	Rate   float64    `json:"rate"`
	Amount TypedMoney `json:"amount"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *TaxPortion) UnmarshalJSON(data []byte) error {
	type Alias TaxPortion
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Amount != nil {
		var err error
		obj.Amount, err = mapDiscriminatorTypedMoney(obj.Amount)
		if err != nil {
			return err
		}
	}

	return nil
}

type TaxedPrice struct {
	// Maps to `TaxedPrice.totalNet`.
	TotalNet Money `json:"totalNet"`
	// Maps to `TaxedPrice.totalGross`.
	TotalGross Money `json:"totalGross"`
	// Maps to `TaxedPrice.taxPortions`.
	TaxPortions []TaxPortion `json:"taxPortions"`
}

/**
*	Maps to `Order.taxMode`
 */
type TaxMode string

const (
	TaxModePlatform       TaxMode = "Platform"
	TaxModeExternal       TaxMode = "External"
	TaxModeExternalAmount TaxMode = "ExternalAmount"
	TaxModeDisabled       TaxMode = "Disabled"
)

/**
*	Maps to `Order.orderState`.
 */
type OrderState string

const (
	OrderStateOpen      OrderState = "Open"
	OrderStateConfirmed OrderState = "Confirmed"
	OrderStateComplete  OrderState = "Complete"
	OrderStateCancelled OrderState = "Cancelled"
)

/**
*	Maps to `Order.shipmentState`.
 */
type ShipmentState string

const (
	ShipmentStateShipped   ShipmentState = "Shipped"
	ShipmentStateReady     ShipmentState = "Ready"
	ShipmentStatePending   ShipmentState = "Pending"
	ShipmentStateDelayed   ShipmentState = "Delayed"
	ShipmentStatePartial   ShipmentState = "Partial"
	ShipmentStateBackorder ShipmentState = "Backorder"
)

/**
*	Maps to `Order.paymentState`.
 */
type PaymentState string

const (
	PaymentStateBalanceDue PaymentState = "BalanceDue"
	PaymentStateFailed     PaymentState = "Failed"
	PaymentStatePending    PaymentState = "Pending"
	PaymentStateCreditOwed PaymentState = "CreditOwed"
	PaymentStatePaid       PaymentState = "Paid"
)

/**
*	Maps to `Order.inventoryMode`.
 */
type InventoryMode string

const (
	InventoryModeTrackOnly      InventoryMode = "TrackOnly"
	InventoryModeReserveOnOrder InventoryMode = "ReserveOnOrder"
)

/**
*	Maps to `Order.taxRoundingMode`.
 */
type RoundingMode string

const (
	RoundingModeHalfEven RoundingMode = "HalfEven"
	RoundingModeHalfUp   RoundingMode = "HalfUp"
	RoundingModeHalfDown RoundingMode = "HalfDown"
)

/**
*	Maps to `Order.taxCalculationMode`.
 */
type TaxCalculationMode string

const (
	TaxCalculationModeLineItemLevel  TaxCalculationMode = "LineItemLevel"
	TaxCalculationModeUnitPriceLevel TaxCalculationMode = "UnitPriceLevel"
)

/**
*	Maps to `Order.origin`.
 */
type CartOrigin string

const (
	CartOriginCustomer CartOrigin = "Customer"
	CartOriginMerchant CartOrigin = "Merchant"
)

type SyncInfo struct {
	// Maps to `SyncInfo.channel`
	Channel ChannelKeyReference `json:"channel"`
	// Maps to `SyncInfo.externalId`
	ExternalId *string `json:"externalId,omitempty"`
	// Maps to `SyncInfo.syncedAt`
	SyncedAt time.Time `json:"syncedAt"`
}

/**
*	Maps to `DiscountCodeInfo.state`
 */
type DiscountCodeState string

const (
	DiscountCodeStateNotActive                            DiscountCodeState = "NotActive"
	DiscountCodeStateNotValid                             DiscountCodeState = "NotValid"
	DiscountCodeStateDoesNotMatchCart                     DiscountCodeState = "DoesNotMatchCart"
	DiscountCodeStateMatchesCart                          DiscountCodeState = "MatchesCart"
	DiscountCodeStateMaxApplicationReached                DiscountCodeState = "MaxApplicationReached"
	DiscountCodeStateApplicationStoppedByPreviousDiscount DiscountCodeState = "ApplicationStoppedByPreviousDiscount"
)

type DiscountCodeInfo struct {
	// References a discount code by key.
	DiscountCode DiscountCodeKeyReference `json:"discountCode"`
	// Maps to `DiscountCodeInfo.state`
	State *DiscountCodeState `json:"state,omitempty"`
}

type ShippingRateInputType string

const (
	ShippingRateInputTypeClassification ShippingRateInputType = "Classification"
	ShippingRateInputTypeScore          ShippingRateInputType = "Score"
)

type ShippingRateInput interface{}

func mapDiscriminatorShippingRateInput(input interface{}) (ShippingRateInput, error) {
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
	case "Classification":
		obj := ClassificationShippingRateInput{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Score":
		obj := ScoreShippingRateInput{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ClassificationShippingRateInput struct {
	Key string `json:"key"`
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	Label LocalizedString `json:"label"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ClassificationShippingRateInput) MarshalJSON() ([]byte, error) {
	type Alias ClassificationShippingRateInput
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Classification", Alias: (*Alias)(&obj)})
}

type ScoreShippingRateInput struct {
	Score float64 `json:"score"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ScoreShippingRateInput) MarshalJSON() ([]byte, error) {
	type Alias ScoreShippingRateInput
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Score", Alias: (*Alias)(&obj)})
}

/**
*	The data representation for an Order to be imported that is persisted as an [Order](/../api/projects/orders#top) in the Project.
*
*	In commercetools, you can import an Order using the
*	[Create Order by Import](/../api/projects/orders-import#create-an-order-by-import)
*	endpoint method instead of creating it from a Cart.
*
*	An OrderImport is a snapshot of an order at the time it was imported.
*
 */
type OrderImport struct {
	// Maps to `Order.orderNumber`, String that uniquely identifies an order. It should be unique across a project. Once it's set it cannot be changed.
	OrderNumber string                `json:"orderNumber"`
	Customer    *CustomerKeyReference `json:"customer,omitempty"`
	// Maps to `Order.customerEmail`.
	CustomerEmail *string `json:"customerEmail,omitempty"`
	// Maps to `Order.lineItems`.
	LineItems []LineItemImportDraft `json:"lineItems"`
	// Maps to `Order.customLineItems`
	CustomLineItems []CustomLineItemDraft `json:"customLineItems"`
	// Maps to `Order.totalPrice`.
	TotalPrice TypedMoney `json:"totalPrice"`
	// Maps to `Order.taxedPrice`.
	TaxedPrice *TaxedPrice `json:"taxedPrice,omitempty"`
	// Maps to `Order.shippingAddress`.
	ShippingAddress *Address `json:"shippingAddress,omitempty"`
	// Maps to `Order.billingAddress`.
	BillingAddress *Address `json:"billingAddress,omitempty"`
	// Maps to `Order.customerGroup`.
	CustomerGroup *CustomerGroupKeyReference `json:"customerGroup,omitempty"`
	// Maps to `Order.country`.
	Country *string `json:"country,omitempty"`
	// Maps to `Order.orderState`.
	OrderState *OrderState `json:"orderState,omitempty"`
	// Maps to `Order.shipmentState`.
	ShipmentState *ShipmentState `json:"shipmentState,omitempty"`
	// Maps to `Order.paymentState`.
	PaymentState *PaymentState `json:"paymentState,omitempty"`
	// Maps to `Order.shippingInfo`.
	ShippingInfo *ShippingInfoImportDraft `json:"shippingInfo,omitempty"`
	// Maps to `Order.completedAt`.
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	// Maps to `Order.custom`.
	Custom *Custom `json:"custom,omitempty"`
	// Maps to `Order.inventoryMode`.
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	// Maps to `Order.taxRoundingMode`.
	TaxRoundingMode *RoundingMode `json:"taxRoundingMode,omitempty"`
	// Maps to `Order.taxCalculationMode`.
	TaxCalculationMode *TaxCalculationMode `json:"taxCalculationMode,omitempty"`
	// Maps to `Order.origin`.
	Origin *CartOrigin `json:"origin,omitempty"`
	// Maps to `Order.itemShippingAddresses`.
	ItemShippingAddresses []Address `json:"itemShippingAddresses"`
	// Reference to the Store in which the Order is associated. If referenced Store does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary Store exists.
	Store *StoreKeyReference `json:"store,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderImport) UnmarshalJSON(data []byte) error {
	type Alias OrderImport
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.TotalPrice != nil {
		var err error
		obj.TotalPrice, err = mapDiscriminatorTypedMoney(obj.TotalPrice)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderImport) MarshalJSON() ([]byte, error) {
	type Alias OrderImport
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

	if raw["lineItems"] == nil {
		delete(raw, "lineItems")
	}

	if raw["customLineItems"] == nil {
		delete(raw, "customLineItems")
	}

	if raw["itemShippingAddresses"] == nil {
		delete(raw, "itemShippingAddresses")
	}

	return json.Marshal(raw)

}
