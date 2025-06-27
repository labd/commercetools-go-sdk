package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type ItemState struct {
	// Number of Line Items or Custom Line Items in this State.
	Quantity int `json:"quantity"`
	// State of the Line Items or Custom Line Items in a custom workflow. If the referenced [State](ctp:api:type:State) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced State is created.
	State StateKeyReference `json:"state"`
}

/**
*	Determines the address (as a reference to an address in `itemShippingAddresses`) and the quantity shipped to the address.
*
 */
type ItemShippingTarget struct {
	// Key of the address in the [Cart](ctp:api:type:Cart) `itemShippingAddresses`. Duplicate address keys are not allowed.
	AddressKey string `json:"addressKey"`
	// Quantity of Line Items or Custom Line Items shipped to the address with the specified `addressKey`.
	Quantity int `json:"quantity"`
}

/**
*	The sum of the `targets` must match the quantity of the Line Items or Custom Line Items
 */
type ItemShippingDetailsDraft struct {
	// Holds information on the quantity of Line Items or Custom Line Items and the address it is shipped.
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
	// Maps to `Price.customerGroup`. References a customer group by key. If the referenced [CustomerGroup](ctp:api:type:CustomerGroup) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced CustomerGroup is created.
	CustomerGroup *CustomerGroupKeyReference `json:"customerGroup,omitempty"`
	// Maps to `Price.channel`. References a channel by key. If the referenced [Channel](ctp:api:type:Channel) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Channel is created.
	Channel *ChannelKeyReference `json:"channel,omitempty"`
	// Sets a discounted price from an external service.
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	// Maps to `Price.tiers`.
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
	// Maps to `ProductVariant.product`. If the referenced [ProductVariant](ctp:api:type:ProductVariant) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced ProductVariant is created.
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
*	Product Attributes are merged with Variant Attributes to ensure the full Attribute context of the Product Variant.
*
 */
type LineItemImportDraft struct {
	// Maps to `LineItem.productId`. If the referenced [Product](ctp:api:type:Product) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Product is created.
	Product *ProductKeyReference `json:"product,omitempty"`
	// Maps to `LineItem.name`.
	Name LocalizedString `json:"name"`
	// Maps to `ProductVariantImportDraft`.
	Variant LineItemProductVariantImportDraft `json:"variant"`
	// Maps to `LineItem.price`.
	Price LineItemPrice `json:"price"`
	// Maps to `LineItem.quantity`.
	Quantity int `json:"quantity"`
	// Maps to `LineItem.state`.
	State []ItemState `json:"state"`
	// Maps to `LineItem.supplyChannel`. If the referenced [Channel](ctp:api:type:Channel) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Channel is created.
	SupplyChannel *ChannelKeyReference `json:"supplyChannel,omitempty"`
	// Maps to `LineItem.distributionChannel`. If the referenced [Channel](ctp:api:type:Channel) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Channel is created.
	DistributionChannel *ChannelKeyReference `json:"distributionChannel,omitempty"`
	// Maps to `LineItem.taxRate`.
	TaxRate *TaxRate `json:"taxRate,omitempty"`
	// Maps to `LineItem.shippingDetails`.
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	// Maps to `LineItem.custom`.
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
	// Currency amount of the ShippingRate.
	Price Money `json:"price"`
	// Free shipping is applied if the sum of the (Custom) Line Item Prices reaches the specified value.
	FreeAbove *Money `json:"freeAbove,omitempty"`
	// Price tiers for the ShippingRate.
	Tiers []ShippingRatePriceTier `json:"tiers"`
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
	// Height of the Parcel.
	HeightInMillimeter *int `json:"heightInMillimeter,omitempty"`
	// Length of the Parcel.
	LengthInMillimeter *int `json:"lengthInMillimeter,omitempty"`
	// Width of the Parcel.
	WidthInMillimeter *int `json:"widthInMillimeter,omitempty"`
	// Weight of the Parcel.
	WeightInGram *int `json:"weightInGram,omitempty"`
}

type TrackingData struct {
	// Identifier to track the Parcel.
	TrackingId *string `json:"trackingId,omitempty"`
	// Name of the carrier that delivers the Parcel.
	Carrier *string `json:"carrier,omitempty"`
	// Name of the provider that serves as facade to several carriers.
	Provider *string `json:"provider,omitempty"`
	// Transaction identifier with the `provider`.
	ProviderTransaction *string `json:"providerTransaction,omitempty"`
	// - If `true`, the Parcel is being returned.
	// - If `false`, the Parcel is being delivered to the customer.
	IsReturn *bool `json:"isReturn,omitempty"`
}

type DeliveryItem struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) or [CustomLineItem](ctp:api:type:CustomLineItem) delivered.
	ID string `json:"id"`
	// Number of Line Items or Custom Line Items delivered.
	Quantity int `json:"quantity"`
}

type Parcel struct {
	// Unique identifier of the Parcel.
	ID string `json:"id"`
	// Date and time (UTC) the Parcel was created.
	CreatedAt time.Time `json:"createdAt"`
	// Information about the dimensions of the Parcel.
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	// Shipment tracking information of the Parcel.
	TrackingData *TrackingData `json:"trackingData,omitempty"`
	// Line Items or Custom Line Items delivered in this Parcel.
	Items []DeliveryItem `json:"items"`
	// Custom Fields of the Parcel.
	Custom *Custom `json:"custom,omitempty"`
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
	// Unique identifier of the Delivery.
	ID string `json:"id"`
	// Date and time (UTC) the Delivery was created.
	CreatedAt time.Time `json:"createdAt"`
	// Line Items or Custom Line Items that are delivered.
	Items []DeliveryItem `json:"items"`
	// Information regarding the appearance, content, and shipment of a Parcel.
	Parcels []Parcel `json:"parcels"`
	// Address to which Parcels are delivered.
	Address *Address `json:"address,omitempty"`
}

type DiscountedLineItemPortion struct {
	// References a cart discount by key. If the referenced [CartDiscount](ctp:api:type:CartDiscount) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced CartDiscount is created.
	Discount CartDiscountKeyReference `json:"discount"`
	// Money value for the discount applicable.
	DiscountedAmount Money `json:"discountedAmount"`
}

type DiscountedLineItemPriceDraft struct {
	// Discounted money value.
	Value Money `json:"value"`
	// Discounts to be applied.
	IncludedDiscounts []DiscountedLineItemPortion `json:"includedDiscounts"`
}

type ShippingMethodState string

const (
	ShippingMethodStateDoesNotMatchCart ShippingMethodState = "DoesNotMatchCart"
	ShippingMethodStateMatchesCart      ShippingMethodState = "MatchesCart"
)

/**
*	Maps to an Order's `shippingInfo` property. This field is usually populated by the Cart associated with the Order, but when importing Orders you must provide a draft representation as a part of the OrderImport.
*
 */
type ShippingInfoImportDraft struct {
	// Maps to `shippingInfo.shippingMethodName`.
	ShippingMethodName string `json:"shippingMethodName"`
	// Maps to `shippingInfo.price`.
	Price TypedMoney `json:"price"`
	// Used to determine the price.
	ShippingRate ShippingRateDraft `json:"shippingRate"`
	// Maps to `shippingInfo.taxRate`.
	TaxRate *TaxRate `json:"taxRate,omitempty"`
	// Maps to `shippingInfo.taxCategory`. If the referenced [TaxCategory](ctp:api:type:TaxCategory) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced TaxCategory is created.
	TaxCategory *TaxCategoryKeyReference `json:"taxCategory,omitempty"`
	// Maps to `shippingInfo.shippingMethod`. If the referenced [ShippingMethod](ctp:api:type:ShippingMethod) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced ShippingMethod is created.
	ShippingMethod *ShippingMethodKeyReference `json:"shippingMethod,omitempty"`
	// Maps to `shippingInfo.deliveries`. You cannot add a `DeliveryItem` on import, as `LineItems` and `CustomLineItems` are not yet referenceable by an `id`.
	Deliveries []Delivery `json:"deliveries"`
	// Maps to `shippingInfo.discountedPrice`.
	DiscountedPrice *DiscountedLineItemPriceDraft `json:"discountedPrice,omitempty"`
	// Maps to `shippingInfo.shippingMethodState`.
	ShippingMethodState *ShippingMethodState `json:"shippingMethodState,omitempty"`
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
	// Name of the Tax Rate.
	Name string `json:"name"`
	// Percentage in the range of 0-1.
	//
	// - If no `subRates` are specified, a value must be defined.
	// - If `subRates` are specified, this can be omitted or its value must be the sum of all `subRates` amounts.
	Amount *float64 `json:"amount,omitempty"`
	// Country for which the tax applies.
	Country string `json:"country"`
	// State within the specified country.
	State *string `json:"state,omitempty"`
	// Used when the total tax is a combination of multiple taxes (for example, local, state/provincial, and/or federal taxes). The total of all subrates must equal the TaxRate `amount`.
	// These subrates are used to calculate the `taxPortions` field of a [Cart](ctp:api:type:Cart) or [Order](ctp:api:type:Order) and the `taxedPrice` field of [LineItems](ctp:api:type:LineItem), [CustomLineItems](ctp:api:type:CustomLineItem), and [ShippingInfos](ctp:api:type:ShippingInfo).
	SubRates []SubRate `json:"subRates"`
	// - If set to `false`, the related price is considered the net price and the provided `amount` is applied to calculate the gross price.
	// - If set to `true`, the related price is considered the gross price, and the provided `amount` is applied to calculate the net price.
	IncludedInPrice *bool `json:"includedInPrice,omitempty"`
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
	// Maps to `CustomLineItem.name`.
	Name LocalizedString `json:"name"`
	// Maps to `CustomLineItem.money`.
	Money TypedMoney `json:"money"`
	// Maps to `CustomLineItem.taxedPrice`.
	TaxedPrice *CustomLineItemTaxedPrice `json:"taxedPrice,omitempty"`
	// Maps to `CustomLineItem.totalPrice`.
	TotalPrice TypedMoney `json:"totalPrice"`
	// Maps to `CustomLineItem.slug`.
	Slug string `json:"slug"`
	// Maps to `CustomLineItem.quantity`.
	Quantity int `json:"quantity"`
	// Maps to `CustomLineItem.state`.
	State []ItemState `json:"state"`
	// Maps to `CustomLineItem.taxCategory`. References a tax category by key. If the referenced [TaxCategory](ctp:api:type:TaxCategory) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced TaxCategory is created.
	TaxCategory *TaxCategoryKeyReference `json:"taxCategory,omitempty"`
	// Maps to `CustomLineItem.taxRate`.
	TaxRate *TaxRate `json:"taxRate,omitempty"`
	// External Tax Rate for the Custom Line Item if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// Maps to `CustomLineItem.discountedPricePerQuantity`.
	DiscountedPricePerQuantity []DiscountedLineItemPriceDraft `json:"discountedPricePerQuantity"`
	// Maps to `CustomLineItem.shippingDetails`.
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
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
	// Name of the tax portion.
	Name *string `json:"name,omitempty"`
	// A number in the range 0-1.
	Rate float64 `json:"rate"`
	// Money value of the tax portion.
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
*	The rounding mode representation used in `Order.priceRoundingMode` and `Order.taxRoundingMode`.
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
	// Maps to `SyncInfo.channel`. If the referenced [Channel](ctp:api:type:Channel) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Channel is created.
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
	// References a DiscountCode by key. If the referenced [DiscountCode](ctp:api:type:DiscountCode) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced DiscountCode is created.
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
	Score int `json:"score"`
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
*	Represents the data used to import an Order. Once imported, this data is persisted as an [Order](ctp:api:type:Order) in the Project.
*
*	An OrderImport is a snapshot of an order at the time it was imported.
*
 */
type OrderImport struct {
	// Maps to `Order.orderNumber`, String that uniquely identifies an order. It should be unique across a project. Once it's set it cannot be changed.
	OrderNumber string `json:"orderNumber"`
	// `key` of the [Customer](ctp:api:type:Customer) that the Order belongs to. If the referenced Customer does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Customer is created.
	Customer *CustomerKeyReference `json:"customer,omitempty"`
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
	// Maps to `Order.customerGroup`. If the referenced [CustomerGroup](ctp:api:type:CustomerGroup) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced CustomerGroup is created.
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
	// Maps to `Order.store`. If the referenced [Store](ctp:api:type:Store) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Store is created.
	Store *StoreKeyReference `json:"store,omitempty"`
	// Maps to `Order.state`. If the referenced [State](ctp:api:type:State) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced State is created.
	State *StateKeyReference `json:"state,omitempty"`
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
