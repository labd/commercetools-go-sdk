// Generated file, please do not change!!!
package importapi

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
	// Maps to `Price.value`. TypedMoney is what is called BaseMoney in the HTTP API.
	Value TypedMoney `json:"value"`
	// Maps to `Price.county`.
	Country string `json:"country,omitEmpty"`
	// Maps to `Price.validFrom`.
	ValidFrom time.Time `json:"validFrom,omitEmpty"`
	// Maps to `Price.validUntil`.
	ValidUntil time.Time `json:"validUntil,omitEmpty"`
	// References a customer group by its key.
	CustomerGroup CustomerGroupKeyReference `json:"customerGroup,omitEmpty"`
	// References a channel by its key.
	Channel ChannelKeyReference `json:"channel,omitEmpty"`
	// Sets a discounted price from an external service.
	Discounted *DiscountedPrice `json:"discounted,omitEmpty"`
	// The tiered prices for this price.
	Tiers []PriceTier `json:"tiers,omitEmpty"`
	// Maps to `Price.custom`.
	Custom *Custom `json:"custom,omitEmpty"`
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

type LineItemProductVariantImportDraft struct {
	// Maps to `ProductVariant.product`.
	ProductVariant ProductVariantKeyReference `json:"productVariant,omitEmpty"`
	// Maps to `ProductVariantImportDraft.sku`.
	Sku string `json:"sku,omitEmpty"`
	// Maps to `ProductVariantImportDraft.prices`
	Prices []LineItemPrice `json:"prices,omitEmpty"`
	// Maps to `ProductVariantImportDraft.attributes`
	Attributes []Attribute `json:"attributes,omitEmpty"`
	// Maps to `ProductVariantImportDraft.images`.
	Images []Image `json:"images,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *LineItemProductVariantImportDraft) UnmarshalJSON(data []byte) error {
	type Alias LineItemProductVariantImportDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

/**
*	Represents an individual Line Item in an Order. A line item is a snapshot of a product at the time it was added to the order.
*
*	You cannot create an Order that includes line item operations that do not exist in the Project or have been deleted.
*	Products and Product Variants referenced by a line item must already exist in the commercetools Project.
*
 */
type LineItemImportDraft struct {
	// Maps to `LineItem.productId`.
	Product ProductKeyReference `json:"product,omitEmpty"`
	// Maps to `LineItem.name`.
	Name LocalizedString `json:"name"`
	// Maps to `ProductVariantImportDraft`.
	Variant LineItemProductVariantImportDraft `json:"variant"`
	// Maps to `LineItem.price`.
	Price LineItemPrice `json:"price"`
	// Maps to `LineItem.quantity`.
	Quantity float64     `json:"quantity"`
	State    []ItemState `json:"state,omitEmpty"`
	// Maps to `LineItem.supplyChannel`.
	// The Reference to the Supply [Channel](/../api/projects/channels#channel) with which the LineItem is associated.
	// If referenced Supply Channel does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `Unresolved` until the necessary Supply Channel is created.
	SupplyChannel ChannelKeyReference `json:"supplyChannel,omitEmpty"`
	// Maps to `LineItem.distributionChannel`.
	// The Reference to the Distribution [Channel](/../api/projects/channels#channel) with which the LineItem is associated.
	// If referenced CustomerGroup does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `Unresolved` until the necessary Distribution Channel is created.
	DistributionChannel ChannelKeyReference `json:"distributionChannel,omitEmpty"`
	// Maps to `LineItem.taxRate`.
	TaxRate *TaxRate `json:"taxRate,omitEmpty"`
	// Maps to LineItem.shippingDetails.
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitEmpty"`
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
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "CartClassification":
		new := CartClassificationTier{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type CartClassificationTier struct {
	Value      string                  `json:"value"`
	Price      Money                   `json:"price"`
	Tiers      []ShippingRatePriceTier `json:"tiers"`
	IsMatching bool                    `json:"isMatching,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartClassificationTier) UnmarshalJSON(data []byte) error {
	type Alias CartClassificationTier
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

// MarshalJSON override to set the discriminator value
func (obj CartClassificationTier) MarshalJSON() ([]byte, error) {
	type Alias CartClassificationTier
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CartClassification", Alias: (*Alias)(&obj)})
}

type ShippingRateDraft struct {
	Price     Money                   `json:"price"`
	FreeAbove Money                   `json:"freeAbove,omitEmpty"`
	Tiers     []ShippingRatePriceTier `json:"tiers,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingRateDraft) UnmarshalJSON(data []byte) error {
	type Alias ShippingRateDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type ParcelMeasurements struct {
	HeightInMillimeter float64 `json:"heightInMillimeter,omitEmpty"`
	LengthInMillimeter float64 `json:"lengthInMillimeter,omitEmpty"`
	WidthInMillimeter  float64 `json:"widthInMillimeter,omitEmpty"`
	WeightInGram       float64 `json:"weightInGram,omitEmpty"`
}

type TrackingData struct {
	TrackingId          string `json:"trackingId,omitEmpty"`
	Carrier             string `json:"carrier,omitEmpty"`
	Provider            string `json:"provider,omitEmpty"`
	ProviderTransaction string `json:"providerTransaction,omitEmpty"`
	IsReturn            bool   `json:"isReturn,omitEmpty"`
}

type DeliveryItem struct {
	Id       string  `json:"id"`
	Quantity float64 `json:"quantity"`
}

type Parcel struct {
	Id           string              `json:"id"`
	CreatedAt    time.Time           `json:"createdAt"`
	Measurements *ParcelMeasurements `json:"measurements,omitEmpty"`
	TrackingData *TrackingData       `json:"trackingData,omitEmpty"`
	Items        []DeliveryItem      `json:"items,omitEmpty"`
}

type Delivery struct {
	Id        string         `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	Items     []DeliveryItem `json:"items"`
	Parcels   []Parcel       `json:"parcels"`
	Address   *Address       `json:"address,omitEmpty"`
}

type DiscountedLineItemPortion struct {
	// References a cart discount by its key.
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
	ShippingMethodName string `json:"shippingMethodName"`
	// TypedMoney is what is called BaseMoney in the HTTP API.
	Price        TypedMoney        `json:"price"`
	ShippingRate ShippingRateDraft `json:"shippingRate"`
	TaxRate      *TaxRate          `json:"taxRate,omitEmpty"`
	// References a tax category by its key.
	TaxCategory TaxCategoryKeyReference `json:"taxCategory,omitEmpty"`
	// References a shipping method by its key.
	ShippingMethod ShippingMethodKeyReference `json:"shippingMethod,omitEmpty"`
	// Note that you can not add a `DeliveryItem` on import, as `LineItems` and `CustomLineItems` are not yet referencable by an `id`.
	Deliveries          []Delivery                    `json:"deliveries,omitEmpty"`
	DiscountedPrice     *DiscountedLineItemPriceDraft `json:"discountedPrice,omitEmpty"`
	ShippingMethodState ShippingMethodState           `json:"shippingMethodState,omitEmpty"`
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

type ExternalTaxRateDraft struct {
	Name            string    `json:"name"`
	Amount          float64   `json:"amount,omitEmpty"`
	Country         string    `json:"country"`
	State           string    `json:"state,omitEmpty"`
	SubRates        []SubRate `json:"subRates,omitEmpty"`
	IncludedInPrice bool      `json:"includedInPrice,omitEmpty"`
}

type CustomLineItemTaxedPrice struct {
	// TypedMoney is what is called BaseMoney in the HTTP API.
	TotalNet TypedMoney `json:"totalNet"`
	// TypedMoney is what is called BaseMoney in the HTTP API.
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
	Name LocalizedString `json:"name"`
	// TypedMoney is what is called BaseMoney in the HTTP API.
	Money      TypedMoney                `json:"money"`
	TaxedPrice *CustomLineItemTaxedPrice `json:"taxedPrice,omitEmpty"`
	// TypedMoney is what is called BaseMoney in the HTTP API.
	TotalPrice TypedMoney  `json:"totalPrice"`
	Slug       string      `json:"slug"`
	Quantity   float64     `json:"quantity"`
	State      []ItemState `json:"state,omitEmpty"`
	// References a tax category by its key.
	TaxCategory                TaxCategoryKeyReference        `json:"taxCategory,omitEmpty"`
	TaxRate                    *TaxRate                       `json:"taxRate,omitEmpty"`
	ExternalTaxRate            *ExternalTaxRateDraft          `json:"externalTaxRate,omitEmpty"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceDraft `json:"discountedPricePerQuantity,omitEmpty"`
	ShippingDetails            *ItemShippingDetailsDraft      `json:"shippingDetails,omitEmpty"`
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

type TaxPortion struct {
	Name string  `json:"name,omitEmpty"`
	Rate float64 `json:"rate"`
	// TypedMoney is what is called BaseMoney in the HTTP API.
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
	ExternalId string `json:"externalId,omitEmpty"`
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
	// References a discount code by its key.
	DiscountCode DiscountCodeKeyReference `json:"discountCode"`
	// Maps to `DiscountCodeInfo.state`
	State DiscountCodeState `json:"state,omitEmpty"`
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
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "Classification":
		new := ClassificationShippingRateInput{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "Score":
		new := ScoreShippingRateInput{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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
*	[Create Order by Import](https://docs.commercetools.com/http-api-projects-orders-import.html#create-an-order-by-import)
*	endpoint method instead of creating it from a Cart.
*
*	An OrderImport is a snapshot of an order at the time it was imported.
*
 */
type OrderImport struct {
	// Maps to `Order.orderNumber`, String that uniquely identifies an order. It should be unique across a project. Once it's set it cannot be changed.
	OrderNumber string               `json:"orderNumber"`
	Customer    CustomerKeyReference `json:"customer,omitEmpty"`
	// Maps to `Order.customerEmail`.
	CustomerEmail string `json:"customerEmail,omitEmpty"`
	// Maps to `Order.lineItems`.
	LineItems []LineItemImportDraft `json:"lineItems,omitEmpty"`
	// Maps to `Order.customLineItems`
	CustomLineItems []CustomLineItemDraft `json:"customLineItems,omitEmpty"`
	// Maps to `Order.totalPrice`. TypedMoney is what is called BaseMoney in the HTTP API.
	TotalPrice TypedMoney `json:"totalPrice"`
	// Maps to `Order.taxedPrice`.
	TaxedPrice *TaxedPrice `json:"taxedPrice,omitEmpty"`
	// Maps to `Order.shippingAddress`.
	ShippingAddress *Address `json:"shippingAddress,omitEmpty"`
	// Maps to `Order.billingAddress`.
	BillingAddress *Address `json:"billingAddress,omitEmpty"`
	// Maps to `Order.customerGroup`.
	CustomerGroup CustomerGroupKeyReference `json:"customerGroup,omitEmpty"`
	// Maps to `Order.country`.
	Country string `json:"country,omitEmpty"`
	// Maps to `Order.orderState`.
	OrderState OrderState `json:"orderState,omitEmpty"`
	// Maps to `Order.shipmentState`.
	ShipmentState ShipmentState `json:"shipmentState,omitEmpty"`
	// Maps to `Order.paymentState`.
	PaymentState PaymentState `json:"paymentState,omitEmpty"`
	// Maps to `Order.shippingInfo`.
	ShippingInfo *ShippingInfoImportDraft `json:"shippingInfo,omitEmpty"`
	// Maps to `Order.completedAt`.
	CompletedAt time.Time `json:"completedAt,omitEmpty"`
	// Maps to `Order.custom`.
	Custom *Custom `json:"custom,omitEmpty"`
	// Maps to `Order.inventoryMode`.
	InventoryMode InventoryMode `json:"inventoryMode,omitEmpty"`
	// Maps to `Order.taxRoundingMode`.
	TaxRoundingMode RoundingMode `json:"taxRoundingMode,omitEmpty"`
	// Maps to `Order.taxCalculationMode`.
	TaxCalculationMode TaxCalculationMode `json:"taxCalculationMode,omitEmpty"`
	// Maps to `Order.origin`.
	Origin CartOrigin `json:"origin,omitEmpty"`
	// Maps to `Order.itemShippingAddresses`.
	ItemShippingAddresses []Address `json:"itemShippingAddresses,omitEmpty"`
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
