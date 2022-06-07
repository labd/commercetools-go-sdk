package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type StagedOrderUpdateAction interface{}

func mapDiscriminatorStagedOrderUpdateAction(input interface{}) (StagedOrderUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "addCustomLineItem":
		obj := StagedOrderAddCustomLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addDelivery":
		obj := StagedOrderAddDeliveryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addDiscountCode":
		obj := StagedOrderAddDiscountCodeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addItemShippingAddress":
		obj := StagedOrderAddItemShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addLineItem":
		obj := StagedOrderAddLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addParcelToDelivery":
		obj := StagedOrderAddParcelToDeliveryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addPayment":
		obj := StagedOrderAddPaymentAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addReturnInfo":
		obj := StagedOrderAddReturnInfoAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addShoppingList":
		obj := StagedOrderAddShoppingListAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeCustomLineItemMoney":
		obj := StagedOrderChangeCustomLineItemMoneyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeCustomLineItemQuantity":
		obj := StagedOrderChangeCustomLineItemQuantityAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeLineItemQuantity":
		obj := StagedOrderChangeLineItemQuantityAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeOrderState":
		obj := StagedOrderChangeOrderStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changePaymentState":
		obj := StagedOrderChangePaymentStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeShipmentState":
		obj := StagedOrderChangeShipmentStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTaxCalculationMode":
		obj := StagedOrderChangeTaxCalculationModeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTaxMode":
		obj := StagedOrderChangeTaxModeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTaxRoundingMode":
		obj := StagedOrderChangeTaxRoundingModeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "importCustomLineItemState":
		obj := StagedOrderImportCustomLineItemStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "importLineItemState":
		obj := StagedOrderImportLineItemStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeCustomLineItem":
		obj := StagedOrderRemoveCustomLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeDelivery":
		obj := StagedOrderRemoveDeliveryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeDiscountCode":
		obj := StagedOrderRemoveDiscountCodeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeItemShippingAddress":
		obj := StagedOrderRemoveItemShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeLineItem":
		obj := StagedOrderRemoveLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeParcelFromDelivery":
		obj := StagedOrderRemoveParcelFromDeliveryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removePayment":
		obj := StagedOrderRemovePaymentAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setBillingAddress":
		obj := StagedOrderSetBillingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setBillingAddressCustomField":
		obj := StagedOrderSetBillingAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setBillingAddressCustomType":
		obj := StagedOrderSetBillingAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCountry":
		obj := StagedOrderSetCountryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := StagedOrderSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomLineItemCustomField":
		obj := StagedOrderSetCustomLineItemCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomLineItemCustomType":
		obj := StagedOrderSetCustomLineItemCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomLineItemShippingDetails":
		obj := StagedOrderSetCustomLineItemShippingDetailsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomLineItemTaxAmount":
		obj := StagedOrderSetCustomLineItemTaxAmountAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomLineItemTaxRate":
		obj := StagedOrderSetCustomLineItemTaxRateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomShippingMethod":
		obj := StagedOrderSetCustomShippingMethodAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := StagedOrderSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomerEmail":
		obj := StagedOrderSetCustomerEmailAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomerGroup":
		obj := StagedOrderSetCustomerGroupAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomerId":
		obj := StagedOrderSetCustomerIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeliveryAddress":
		obj := StagedOrderSetDeliveryAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeliveryAddressCustomField":
		obj := StagedOrderSetDeliveryAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeliveryAddressCustomType":
		obj := StagedOrderSetDeliveryAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeliveryCustomField":
		obj := StagedOrderSetDeliveryCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeliveryCustomType":
		obj := StagedOrderSetDeliveryCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeliveryItems":
		obj := StagedOrderSetDeliveryItemsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setItemShippingAddressCustomField":
		obj := StagedOrderSetItemShippingAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setItemShippingAddressCustomType":
		obj := StagedOrderSetItemShippingAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemCustomField":
		obj := StagedOrderSetLineItemCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemCustomType":
		obj := StagedOrderSetLineItemCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemDistributionChannel":
		obj := StagedOrderSetLineItemDistributionChannelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemPrice":
		obj := StagedOrderSetLineItemPriceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemShippingDetails":
		obj := StagedOrderSetLineItemShippingDetailsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemTaxAmount":
		obj := StagedOrderSetLineItemTaxAmountAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemTaxRate":
		obj := StagedOrderSetLineItemTaxRateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemTotalPrice":
		obj := StagedOrderSetLineItemTotalPriceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLocale":
		obj := StagedOrderSetLocaleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setOrderNumber":
		obj := StagedOrderSetOrderNumberAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setOrderTotalTax":
		obj := StagedOrderSetOrderTotalTaxAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setParcelCustomField":
		obj := StagedOrderSetParcelCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setParcelCustomType":
		obj := StagedOrderSetParcelCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setParcelItems":
		obj := StagedOrderSetParcelItemsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setParcelMeasurements":
		obj := StagedOrderSetParcelMeasurementsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setParcelTrackingData":
		obj := StagedOrderSetParcelTrackingDataAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setReturnInfo":
		obj := StagedOrderSetReturnInfoAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setReturnItemCustomField":
		obj := StagedOrderSetReturnItemCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setReturnItemCustomType":
		obj := StagedOrderSetReturnItemCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setReturnPaymentState":
		obj := StagedOrderSetReturnPaymentStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setReturnShipmentState":
		obj := StagedOrderSetReturnShipmentStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingAddress":
		obj := StagedOrderSetShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingAddressAndCustomShippingMethod":
		obj := StagedOrderSetShippingAddressAndCustomShippingMethodAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingAddressAndShippingMethod":
		obj := StagedOrderSetShippingAddressAndShippingMethodAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingAddressCustomField":
		obj := StagedOrderSetShippingAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingAddressCustomType":
		obj := StagedOrderSetShippingAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingMethod":
		obj := StagedOrderSetShippingMethodAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingMethodTaxAmount":
		obj := StagedOrderSetShippingMethodTaxAmountAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingMethodTaxRate":
		obj := StagedOrderSetShippingMethodTaxRateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingRateInput":
		obj := StagedOrderSetShippingRateInputAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.ShippingRateInput != nil {
			var err error
			obj.ShippingRateInput, err = mapDiscriminatorShippingRateInputDraft(obj.ShippingRateInput)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "transitionCustomLineItemState":
		obj := StagedOrderTransitionCustomLineItemStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "transitionLineItemState":
		obj := StagedOrderTransitionLineItemStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "transitionState":
		obj := StagedOrderTransitionStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "updateItemShippingAddress":
		obj := StagedOrderUpdateItemShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "updateSyncInfo":
		obj := StagedOrderUpdateSyncInfoAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type Hit struct {
	// Unique identifier of the Order.
	ID string `json:"id"`
	// Current version of the Order.
	Version int `json:"version"`
	// The higher the value is, the more relevant the hit is for the search request.
	Relevance float64 `json:"relevance"`
}

type OrderPagedSearchResponse struct {
	// Total number of results matching the query.
	Total int `json:"total"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset *int `json:"offset,omitempty"`
	// Number of [results requested](/../api/general-concepts#limit).
	Limit *int `json:"limit,omitempty"`
	// Actual results.
	Hits []Hit `json:"hits"`
}

type Delivery struct {
	// Unique identifier of the Delivery.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	// Items which are shipped in this delivery regardless their distribution over several parcels.
	// Can also be specified individually for each Parcel.
	Items   []DeliveryItem `json:"items"`
	Parcels []Parcel       `json:"parcels"`
	Address *Address       `json:"address,omitempty"`
	// Custom Fields for the Transaction.
	Custom *CustomFields `json:"custom,omitempty"`
}

type DeliveryItem struct {
	// Unique identifier of the DeliveryItem.
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}

type DiscountedLineItemPriceDraft struct {
	// Draft type that stores amounts in cent precision for the specified currency.
	//
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	Value             Money                       `json:"value"`
	IncludedDiscounts []DiscountedLineItemPortion `json:"includedDiscounts"`
}

type ItemState struct {
	Quantity int            `json:"quantity"`
	State    StateReference `json:"state"`
}

type LineItemImportDraft struct {
	// ID of the existing product.
	// You also need to specify the ID of the variant if this property is set or alternatively you can just specify SKU of the product variant.
	ProductId *string `json:"productId,omitempty"`
	// The product name.
	Name     LocalizedString           `json:"name"`
	Variant  ProductVariantImportDraft `json:"variant"`
	Price    PriceDraft                `json:"price"`
	Quantity int                       `json:"quantity"`
	State    []ItemState               `json:"state"`
	// Connection to a particular supplier.
	// By providing supply channel information, you can uniquely identify
	// inventory entries that should be reserved.
	// The provided channel should have the
	// InventorySupply role.
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	// The channel is used to select a ProductPrice.
	// The provided channel should have the ProductDistribution role.
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
	TaxRate             *TaxRate                   `json:"taxRate,omitempty"`
	// The custom fields.
	Custom          *CustomFieldsDraft        `json:"custom,omitempty"`
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["state"] == nil {
		delete(target, "state")
	}

	return json.Marshal(target)
}

type Order struct {
	// Unique identifier of the Order.
	ID string `json:"id"`
	// The current version of the order.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// This field will only be present if it was set for Order Import
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	// String that uniquely identifies an order.
	// It can be used to create more human-readable (in contrast to ID) identifier for the order.
	// It should be unique across a project.
	// Once it's set it cannot be changed.
	OrderNumber   *string `json:"orderNumber,omitempty"`
	CustomerId    *string `json:"customerId,omitempty"`
	CustomerEmail *string `json:"customerEmail,omitempty"`
	// Identifies carts and orders belonging to an anonymous session (the customer has not signed up/in yet).
	AnonymousId     *string            `json:"anonymousId,omitempty"`
	Store           *StoreKeyReference `json:"store,omitempty"`
	LineItems       []LineItem         `json:"lineItems"`
	CustomLineItems []CustomLineItem   `json:"customLineItems"`
	TotalPrice      TypedMoney         `json:"totalPrice"`
	// The taxes are calculated based on the shipping address.
	TaxedPrice      *TaxedPrice `json:"taxedPrice,omitempty"`
	ShippingAddress *Address    `json:"shippingAddress,omitempty"`
	BillingAddress  *Address    `json:"billingAddress,omitempty"`
	TaxMode         *TaxMode    `json:"taxMode,omitempty"`
	// When calculating taxes for `taxedPrice`, the selected mode is used for rouding.
	TaxRoundingMode *RoundingMode `json:"taxRoundingMode,omitempty"`
	// Set when the customer is set and the customer is a member of a customer group.
	// Used for product variant price selection.
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	// Used for product variant price selection.
	Country *string `json:"country,omitempty"`
	// One of the four predefined OrderStates.
	OrderState OrderState `json:"orderState"`
	// This reference can point to a state in a custom workflow.
	State         *StateReference `json:"state,omitempty"`
	ShipmentState *ShipmentState  `json:"shipmentState,omitempty"`
	PaymentState  *PaymentState   `json:"paymentState,omitempty"`
	// Set if the ShippingMethod is set.
	ShippingInfo  *ShippingInfo      `json:"shippingInfo,omitempty"`
	SyncInfo      []SyncInfo         `json:"syncInfo"`
	ReturnInfo    []ReturnInfo       `json:"returnInfo"`
	DiscountCodes []DiscountCodeInfo `json:"discountCodes"`
	// Internal-only field.
	LastMessageSequenceNumber *int `json:"lastMessageSequenceNumber,omitempty"`
	// Set when this order was created from a cart.
	// The cart will have the state `Ordered`.
	Cart          *CartReference `json:"cart,omitempty"`
	Custom        *CustomFields  `json:"custom,omitempty"`
	PaymentInfo   *PaymentInfo   `json:"paymentInfo,omitempty"`
	Locale        *string        `json:"locale,omitempty"`
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	Origin        CartOrigin     `json:"origin"`
	// When calculating taxes for `taxedPrice`, the selected mode is used for calculating the price with LineItemLevel (horizontally) or UnitPriceLevel (vertically) calculation mode.
	TaxCalculationMode *TaxCalculationMode `json:"taxCalculationMode,omitempty"`
	// The shippingRateInput is used as an input to select a ShippingRatePriceTier.
	ShippingRateInput ShippingRateInput `json:"shippingRateInput,omitempty"`
	// Contains addresses for orders with multiple shipping addresses.
	ItemShippingAddresses []Address `json:"itemShippingAddresses"`
	// Automatically filled when a line item with LineItemMode `GiftLineItem` is removed from this order.
	RefusedGifts []CartDiscountReference `json:"refusedGifts"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Order) UnmarshalJSON(data []byte) error {
	type Alias Order
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
	if obj.ShippingRateInput != nil {
		var err error
		obj.ShippingRateInput, err = mapDiscriminatorShippingRateInput(obj.ShippingRateInput)
		if err != nil {
			return err
		}
	}
	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Order) MarshalJSON() ([]byte, error) {
	type Alias Order
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["returnInfo"] == nil {
		delete(target, "returnInfo")
	}

	if target["discountCodes"] == nil {
		delete(target, "discountCodes")
	}

	if target["itemShippingAddresses"] == nil {
		delete(target, "itemShippingAddresses")
	}

	return json.Marshal(target)
}

type OrderFromCartDraft struct {
	// Unique identifier of the Cart from which you can create an Order.
	ID *string `json:"id,omitempty"`
	// ResourceIdentifier to the Cart from which this order is created.
	Cart    *CartResourceIdentifier `json:"cart,omitempty"`
	Version int                     `json:"version"`
	// String that uniquely identifies an order.
	// It can be used to create more human-readable (in contrast to ID) identifier for the order.
	// It should be unique across a project.
	// Once it's set it cannot be changed.
	// For easier use on Get, Update and Delete actions we suggest assigning order numbers that match the regular expression `[a-z0-9_\-]{2,36}`.
	OrderNumber   *string        `json:"orderNumber,omitempty"`
	PaymentState  *PaymentState  `json:"paymentState,omitempty"`
	ShipmentState *ShipmentState `json:"shipmentState,omitempty"`
	// Order will be created with `Open` status by default.
	OrderState *OrderState              `json:"orderState,omitempty"`
	State      *StateResourceIdentifier `json:"state,omitempty"`
	// [Custom Fields](/../api/projects/custom-fields) for the Order. The Custom Field type must match the type of the Custom Fields in the referenced [Cart](/../api/projects/carts#cart).
	// If specified, the Custom Fields are merged with the Custom Fields on the referenced [Cart](/../api/projects/carts#cart) and added to the Order.
	// If empty, the Custom Fields on the referenced [Cart](/../api/projects/carts#cart) are added to the Order automatically.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

type OrderImportDraft struct {
	// String that unique identifies an order.
	// It can be used to create more human-readable (in contrast to ID) identifier for the order.
	// It should be unique within a project.
	OrderNumber *string `json:"orderNumber,omitempty"`
	// If given the customer with that ID must exist in the project.
	CustomerId *string `json:"customerId,omitempty"`
	// The customer email can be used when no check against existing Customers is desired during order import.
	CustomerEmail *string `json:"customerEmail,omitempty"`
	// If not given `customLineItems` must not be empty.
	LineItems []LineItemImportDraft `json:"lineItems"`
	// If not given `lineItems` must not be empty.
	CustomLineItems []CustomLineItemImportDraft `json:"customLineItems"`
	TotalPrice      Money                       `json:"totalPrice"`
	// Order Import does not support calculation of taxes.
	// When setting the draft the taxedPrice is to be provided.
	TaxedPrice      *TaxedPriceDraft `json:"taxedPrice,omitempty"`
	ShippingAddress *BaseAddress     `json:"shippingAddress,omitempty"`
	BillingAddress  *BaseAddress     `json:"billingAddress,omitempty"`
	// Set when the customer is set and the customer is a member of a customer group.
	// Used for product variant price selection.
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	// Used for product variant price selection.
	Country *string `json:"country,omitempty"`
	// If not given the `Open` state will be assigned by default.
	OrderState *OrderState `json:"orderState,omitempty"`
	// This reference can point to a state in a custom workflow.
	State         *StateReference `json:"state,omitempty"`
	ShipmentState *ShipmentState  `json:"shipmentState,omitempty"`
	PaymentState  *PaymentState   `json:"paymentState,omitempty"`
	// Set if the ShippingMethod is set.
	ShippingInfo *ShippingInfoImportDraft `json:"shippingInfo,omitempty"`
	PaymentInfo  *PaymentInfo             `json:"paymentInfo,omitempty"`
	CompletedAt  *time.Time               `json:"completedAt,omitempty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// If not given the mode `None` will be assigned by default.
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	// If not given the tax rounding mode `HalfEven` will be assigned by default.
	TaxRoundingMode *RoundingMode `json:"taxRoundingMode,omitempty"`
	// Contains addresses for orders with multiple shipping addresses.
	ItemShippingAddresses []BaseAddress            `json:"itemShippingAddresses"`
	Store                 *StoreResourceIdentifier `json:"store,omitempty"`
	// The default origin is `Customer`.
	Origin *CartOrigin `json:"origin,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderImportDraft) MarshalJSON() ([]byte, error) {
	type Alias OrderImportDraft
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["lineItems"] == nil {
		delete(target, "lineItems")
	}

	if target["customLineItems"] == nil {
		delete(target, "customLineItems")
	}

	if target["itemShippingAddresses"] == nil {
		delete(target, "itemShippingAddresses")
	}

	return json.Marshal(target)
}

type OrderPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int  `json:"limit"`
	Count int  `json:"count"`
	Total *int `json:"total,omitempty"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset  int     `json:"offset"`
	Results []Order `json:"results"`
}

type OrderReference struct {
	// Unique ID of the referenced resource.
	ID  string `json:"id"`
	Obj *Order `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderReference) MarshalJSON() ([]byte, error) {
	type Alias OrderReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "order", Alias: (*Alias)(&obj)})
}

type OrderResourceIdentifier struct {
	// Unique identifier of the referenced resource. Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced resource. Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias OrderResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "order", Alias: (*Alias)(&obj)})
}

type OrderSearchQuery map[string]interface{}
type OrderSearchRequest struct {
	// The Order search query.
	Query OrderSearchQuery `json:"query"`
	// Controls how results to your query are sorted. If not provided, the results are sorted by relevance in descending order.
	Sort *string `json:"sort,omitempty"`
	// The maximum number of search results to be returned.
	Limit *int `json:"limit,omitempty"`
	// The number of search results to be skipped in the response for pagination.
	Offset *int `json:"offset,omitempty"`
}

type OrderState string

const (
	OrderStateOpen      OrderState = "Open"
	OrderStateConfirmed OrderState = "Confirmed"
	OrderStateComplete  OrderState = "Complete"
	OrderStateCancelled OrderState = "Cancelled"
)

type OrderUpdate struct {
	Version int                 `json:"version"`
	Actions []OrderUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderUpdate) UnmarshalJSON(data []byte) error {
	type Alias OrderUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorOrderUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}
	return nil
}

type OrderUpdateAction interface{}

func mapDiscriminatorOrderUpdateAction(input interface{}) (OrderUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "addDelivery":
		obj := OrderAddDeliveryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addItemShippingAddress":
		obj := OrderAddItemShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addParcelToDelivery":
		obj := OrderAddParcelToDeliveryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addPayment":
		obj := OrderAddPaymentAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addReturnInfo":
		obj := OrderAddReturnInfoAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeOrderState":
		obj := OrderChangeOrderStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changePaymentState":
		obj := OrderChangePaymentStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeShipmentState":
		obj := OrderChangeShipmentStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "importCustomLineItemState":
		obj := OrderImportCustomLineItemStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "importLineItemState":
		obj := OrderImportLineItemStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeDelivery":
		obj := OrderRemoveDeliveryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeItemShippingAddress":
		obj := OrderRemoveItemShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeParcelFromDelivery":
		obj := OrderRemoveParcelFromDeliveryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removePayment":
		obj := OrderRemovePaymentAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setBillingAddress":
		obj := OrderSetBillingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setBillingAddressCustomField":
		obj := OrderSetBillingAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setBillingAddressCustomType":
		obj := OrderSetBillingAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := OrderSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomLineItemCustomField":
		obj := OrderSetCustomLineItemCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomLineItemCustomType":
		obj := OrderSetCustomLineItemCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomLineItemShippingDetails":
		obj := OrderSetCustomLineItemShippingDetailsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := OrderSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomerEmail":
		obj := OrderSetCustomerEmailAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomerId":
		obj := OrderSetCustomerIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeliveryAddress":
		obj := OrderSetDeliveryAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeliveryAddressCustomField":
		obj := OrderSetDeliveryAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeliveryAddressCustomType":
		obj := OrderSetDeliveryAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeliveryCustomField":
		obj := OrderSetDeliveryCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeliveryCustomType":
		obj := OrderSetDeliveryCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeliveryItems":
		obj := OrderSetDeliveryItemsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setItemShippingAddressCustomField":
		obj := OrderSetItemShippingAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setItemShippingAddressCustomType":
		obj := OrderSetItemShippingAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemCustomField":
		obj := OrderSetLineItemCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemCustomType":
		obj := OrderSetLineItemCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemShippingDetails":
		obj := OrderSetLineItemShippingDetailsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLocale":
		obj := OrderSetLocaleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setOrderNumber":
		obj := OrderSetOrderNumberAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setParcelCustomField":
		obj := OrderSetParcelCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setParcelCustomType":
		obj := OrderSetParcelCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setParcelItems":
		obj := OrderSetParcelItemsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setParcelMeasurements":
		obj := OrderSetParcelMeasurementsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setParcelTrackingData":
		obj := OrderSetParcelTrackingDataAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setReturnInfo":
		obj := OrderSetReturnInfoAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setReturnItemCustomField":
		obj := OrderSetReturnItemCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setReturnItemCustomType":
		obj := OrderSetReturnItemCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setReturnPaymentState":
		obj := OrderSetReturnPaymentStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setReturnShipmentState":
		obj := OrderSetReturnShipmentStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingAddress":
		obj := OrderSetShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingAddressCustomField":
		obj := OrderSetShippingAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingAddressCustomType":
		obj := OrderSetShippingAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setStore":
		obj := OrderSetStoreAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "transitionCustomLineItemState":
		obj := OrderTransitionCustomLineItemStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "transitionLineItemState":
		obj := OrderTransitionLineItemStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "transitionState":
		obj := OrderTransitionStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "updateItemShippingAddress":
		obj := OrderUpdateItemShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "updateSyncInfo":
		obj := OrderUpdateSyncInfoAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type Parcel struct {
	// Unique identifier of the Parcel.
	ID           string              `json:"id"`
	CreatedAt    time.Time           `json:"createdAt"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	TrackingData *TrackingData       `json:"trackingData,omitempty"`
	// The delivery items contained in this parcel.
	Items []DeliveryItem `json:"items"`
	// Custom Fields of this parcel.
	Custom *CustomFields `json:"custom,omitempty"`
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["items"] == nil {
		delete(target, "items")
	}

	return json.Marshal(target)
}

type ParcelDraft struct {
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	TrackingData *TrackingData       `json:"trackingData,omitempty"`
	// The delivery items contained in this parcel.
	Items []DeliveryItem `json:"items"`
	// Custom Fields of this parcel.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ParcelDraft) MarshalJSON() ([]byte, error) {
	type Alias ParcelDraft
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["items"] == nil {
		delete(target, "items")
	}

	return json.Marshal(target)
}

type ParcelMeasurements struct {
	HeightInMillimeter *int `json:"heightInMillimeter,omitempty"`
	LengthInMillimeter *int `json:"lengthInMillimeter,omitempty"`
	WidthInMillimeter  *int `json:"widthInMillimeter,omitempty"`
	WeightInGram       *int `json:"weightInGram,omitempty"`
}

type PaymentInfo struct {
	Payments []PaymentReference `json:"payments"`
}

type PaymentState string

const (
	PaymentStateBalanceDue PaymentState = "BalanceDue"
	PaymentStateFailed     PaymentState = "Failed"
	PaymentStatePending    PaymentState = "Pending"
	PaymentStateCreditOwed PaymentState = "CreditOwed"
	PaymentStatePaid       PaymentState = "Paid"
)

type ProductVariantImportDraft struct {
	// The sequential ID of the variant within the product.
	// The variant with provided ID should exist in some existing product, so you also need to specify the productId if this property is set,
	// or alternatively you can just specify SKU of the product variant.
	ID *int `json:"id,omitempty"`
	// The SKU of the existing variant.
	Sku *string `json:"sku,omitempty"`
	// The [EmbeddedPrices](ctp:api:type:EmbeddedPrice) of the variant.
	// The prices should not contain two prices for the same price scope (same currency, country and customer group).
	// If this property is defined, then it will override the `prices` property from the original product variant, otherwise `prices` property from the original product variant would be copied in the resulting order.
	Prices []PriceDraft `json:"prices"`
	// If this property is defined, then it will override the `attributes` property from the original
	// product variant, otherwise `attributes` property from the original product variant would be copied in the resulting order.
	Attributes []Attribute `json:"attributes"`
	// If this property is defined, then it will override the `images` property from the original
	// product variant, otherwise `images` property from the original product variant would be copied in the resulting order.
	Images []Image `json:"images"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariantImportDraft) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantImportDraft
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["prices"] == nil {
		delete(target, "prices")
	}

	if target["attributes"] == nil {
		delete(target, "attributes")
	}

	if target["images"] == nil {
		delete(target, "images")
	}

	return json.Marshal(target)
}

type ReturnInfo struct {
	Items []ReturnItem `json:"items"`
	// Identifies, which return tracking ID is connected to this particular return.
	ReturnTrackingId *string    `json:"returnTrackingId,omitempty"`
	ReturnDate       *time.Time `json:"returnDate,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReturnInfo) UnmarshalJSON(data []byte) error {
	type Alias ReturnInfo
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Items {
		var err error
		obj.Items[i], err = mapDiscriminatorReturnItem(obj.Items[i])
		if err != nil {
			return err
		}
	}
	return nil
}

type ReturnInfoDraft struct {
	Items []ReturnItemDraft `json:"items"`
	// Identifies, which return tracking ID is connected to this particular return.
	ReturnTrackingId *string    `json:"returnTrackingId,omitempty"`
	ReturnDate       *time.Time `json:"returnDate,omitempty"`
}

type ReturnItem interface{}

func mapDiscriminatorReturnItem(input interface{}) (ReturnItem, error) {
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
	case "CustomLineItemReturnItem":
		obj := CustomLineItemReturnItem{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "LineItemReturnItem":
		obj := LineItemReturnItem{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type CustomLineItemReturnItem struct {
	// Unique identifier of the ReturnItem.
	ID            string              `json:"id"`
	Quantity      int                 `json:"quantity"`
	Comment       *string             `json:"comment,omitempty"`
	ShipmentState ReturnShipmentState `json:"shipmentState"`
	PaymentState  ReturnPaymentState  `json:"paymentState"`
	// Custom Fields of this return item.
	Custom           *CustomFields `json:"custom,omitempty"`
	LastModifiedAt   time.Time     `json:"lastModifiedAt"`
	CreatedAt        time.Time     `json:"createdAt"`
	CustomLineItemId string        `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomLineItemReturnItem) MarshalJSON() ([]byte, error) {
	type Alias CustomLineItemReturnItem
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomLineItemReturnItem", Alias: (*Alias)(&obj)})
}

type LineItemReturnItem struct {
	// Unique identifier of the ReturnItem.
	ID            string              `json:"id"`
	Quantity      int                 `json:"quantity"`
	Comment       *string             `json:"comment,omitempty"`
	ShipmentState ReturnShipmentState `json:"shipmentState"`
	PaymentState  ReturnPaymentState  `json:"paymentState"`
	// Custom Fields of this return item.
	Custom         *CustomFields `json:"custom,omitempty"`
	LastModifiedAt time.Time     `json:"lastModifiedAt"`
	CreatedAt      time.Time     `json:"createdAt"`
	LineItemId     string        `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LineItemReturnItem) MarshalJSON() ([]byte, error) {
	type Alias LineItemReturnItem
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "LineItemReturnItem", Alias: (*Alias)(&obj)})
}

type ReturnItemDraft struct {
	Quantity         int                 `json:"quantity"`
	LineItemId       *string             `json:"lineItemId,omitempty"`
	CustomLineItemId *string             `json:"customLineItemId,omitempty"`
	Comment          *string             `json:"comment,omitempty"`
	ShipmentState    ReturnShipmentState `json:"shipmentState"`
	// Custom Fields of this return item.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

type ReturnPaymentState string

const (
	ReturnPaymentStateNonRefundable ReturnPaymentState = "NonRefundable"
	ReturnPaymentStateInitial       ReturnPaymentState = "Initial"
	ReturnPaymentStateRefunded      ReturnPaymentState = "Refunded"
	ReturnPaymentStateNotRefunded   ReturnPaymentState = "NotRefunded"
)

type ReturnShipmentState string

const (
	ReturnShipmentStateAdvised     ReturnShipmentState = "Advised"
	ReturnShipmentStateReturned    ReturnShipmentState = "Returned"
	ReturnShipmentStateBackInStock ReturnShipmentState = "BackInStock"
	ReturnShipmentStateUnusable    ReturnShipmentState = "Unusable"
)

type ShipmentState string

const (
	ShipmentStateShipped   ShipmentState = "Shipped"
	ShipmentStateReady     ShipmentState = "Ready"
	ShipmentStatePending   ShipmentState = "Pending"
	ShipmentStateDelayed   ShipmentState = "Delayed"
	ShipmentStatePartial   ShipmentState = "Partial"
	ShipmentStateBackorder ShipmentState = "Backorder"
)

type ShippingInfoImportDraft struct {
	ShippingMethodName string `json:"shippingMethodName"`
	Price              Money  `json:"price"`
	// The shipping rate used to determine the price.
	ShippingRate ShippingRateDraft              `json:"shippingRate"`
	TaxRate      *TaxRate                       `json:"taxRate,omitempty"`
	TaxCategory  *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// Not set if custom shipping method is used.
	ShippingMethod *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	// Deliveries are compilations of information on how the articles are being delivered to the customers.
	Deliveries      []Delivery                    `json:"deliveries"`
	DiscountedPrice *DiscountedLineItemPriceDraft `json:"discountedPrice,omitempty"`
	// Indicates whether the ShippingMethod referenced is allowed for the cart or not.
	ShippingMethodState *ShippingMethodState `json:"shippingMethodState,omitempty"`
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["deliveries"] == nil {
		delete(target, "deliveries")
	}

	return json.Marshal(target)
}

type SyncInfo struct {
	// Connection to a particular synchronization destination.
	Channel ChannelReference `json:"channel"`
	// Can be used to reference an external order instance, file etc.
	ExternalId *string   `json:"externalId,omitempty"`
	SyncedAt   time.Time `json:"syncedAt"`
}

type TaxedItemPriceDraft struct {
	// Draft type that stores amounts in cent precision for the specified currency.
	//
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	TotalNet Money `json:"totalNet"`
	// Draft type that stores amounts in cent precision for the specified currency.
	//
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	TotalGross Money `json:"totalGross"`
}

type TrackingData struct {
	// The ID to track one parcel.
	TrackingId *string `json:"trackingId,omitempty"`
	// The carrier that delivers the parcel.
	Carrier             *string `json:"carrier,omitempty"`
	Provider            *string `json:"provider,omitempty"`
	ProviderTransaction *string `json:"providerTransaction,omitempty"`
	// Flag to distinguish if the parcel is on the way to the customer (false) or on the way back (true).
	IsReturn *bool `json:"isReturn,omitempty"`
}

type OrderAddDeliveryAction struct {
	Items   []DeliveryItem `json:"items"`
	Address *BaseAddress   `json:"address,omitempty"`
	Parcels []ParcelDraft  `json:"parcels"`
	// Custom Fields for the Transaction.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderAddDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddDeliveryAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDelivery", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["items"] == nil {
		delete(target, "items")
	}

	if target["parcels"] == nil {
		delete(target, "parcels")
	}

	return json.Marshal(target)
}

type OrderAddItemShippingAddressAction struct {
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderAddItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addItemShippingAddress", Alias: (*Alias)(&obj)})
}

type OrderAddParcelToDeliveryAction struct {
	DeliveryId   string              `json:"deliveryId"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	TrackingData *TrackingData       `json:"trackingData,omitempty"`
	Items        []DeliveryItem      `json:"items"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderAddParcelToDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddParcelToDeliveryAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addParcelToDelivery", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["items"] == nil {
		delete(target, "items")
	}

	return json.Marshal(target)
}

type OrderAddPaymentAction struct {
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Payment](ctp:api:type:Payment).
	Payment PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderAddPaymentAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddPaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPayment", Alias: (*Alias)(&obj)})
}

type OrderAddReturnInfoAction struct {
	ReturnTrackingId *string           `json:"returnTrackingId,omitempty"`
	Items            []ReturnItemDraft `json:"items"`
	ReturnDate       *time.Time        `json:"returnDate,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderAddReturnInfoAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddReturnInfoAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addReturnInfo", Alias: (*Alias)(&obj)})
}

type OrderChangeOrderStateAction struct {
	OrderState OrderState `json:"orderState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderChangeOrderStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderChangeOrderStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeOrderState", Alias: (*Alias)(&obj)})
}

type OrderChangePaymentStateAction struct {
	PaymentState *PaymentState `json:"paymentState,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderChangePaymentStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderChangePaymentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePaymentState", Alias: (*Alias)(&obj)})
}

type OrderChangeShipmentStateAction struct {
	ShipmentState *ShipmentState `json:"shipmentState,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderChangeShipmentStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderChangeShipmentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeShipmentState", Alias: (*Alias)(&obj)})
}

type OrderImportCustomLineItemStateAction struct {
	CustomLineItemId string      `json:"customLineItemId"`
	State            []ItemState `json:"state"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderImportCustomLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderImportCustomLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "importCustomLineItemState", Alias: (*Alias)(&obj)})
}

type OrderImportLineItemStateAction struct {
	LineItemId string      `json:"lineItemId"`
	State      []ItemState `json:"state"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderImportLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderImportLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "importLineItemState", Alias: (*Alias)(&obj)})
}

type OrderRemoveDeliveryAction struct {
	DeliveryId string `json:"deliveryId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderRemoveDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias OrderRemoveDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeDelivery", Alias: (*Alias)(&obj)})
}

type OrderRemoveItemShippingAddressAction struct {
	AddressKey string `json:"addressKey"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderRemoveItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderRemoveItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeItemShippingAddress", Alias: (*Alias)(&obj)})
}

type OrderRemoveParcelFromDeliveryAction struct {
	ParcelId string `json:"parcelId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderRemoveParcelFromDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias OrderRemoveParcelFromDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeParcelFromDelivery", Alias: (*Alias)(&obj)})
}

type OrderRemovePaymentAction struct {
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Payment](ctp:api:type:Payment).
	Payment PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderRemovePaymentAction) MarshalJSON() ([]byte, error) {
	type Alias OrderRemovePaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePayment", Alias: (*Alias)(&obj)})
}

type OrderSetBillingAddressAction struct {
	Address *BaseAddress `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddress", Alias: (*Alias)(&obj)})
}

type OrderSetBillingAddressCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetBillingAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetBillingAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddressCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetBillingAddressCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the `billingAddress` with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the `billingAddress`.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `billingAddress`.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetBillingAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetBillingAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddressCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetCustomLineItemCustomFieldAction struct {
	CustomLineItemId string `json:"customLineItemId"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetCustomLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetCustomLineItemCustomTypeAction struct {
	CustomLineItemId string `json:"customLineItemId"`
	// Defines the [Type](ctp:api:type:Type) that extends the CustomLineItem with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the CustomLineItem.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the CustomLineItem.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetCustomLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetCustomLineItemShippingDetailsAction struct {
	CustomLineItemId string                    `json:"customLineItemId"`
	ShippingDetails  *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetCustomLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

type OrderSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the Order with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Order.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Order.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetCustomerEmailAction struct {
	Email *string `json:"email,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetCustomerEmailAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomerEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerEmail", Alias: (*Alias)(&obj)})
}

type OrderSetCustomerIdAction struct {
	CustomerId *string `json:"customerId,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetCustomerIdAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomerIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerId", Alias: (*Alias)(&obj)})
}

type OrderSetDeliveryAddressAction struct {
	DeliveryId string       `json:"deliveryId"`
	Address    *BaseAddress `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetDeliveryAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetDeliveryAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryAddress", Alias: (*Alias)(&obj)})
}

type OrderSetDeliveryAddressCustomFieldAction struct {
	DeliveryId string `json:"deliveryId"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetDeliveryAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetDeliveryAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryAddressCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetDeliveryAddressCustomTypeAction struct {
	DeliveryId string `json:"deliveryId"`
	// Defines the [Type](ctp:api:type:Type) that extends the `address` in a Delivery with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the `address` in a Delivery.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `address` in a Delivery.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetDeliveryAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetDeliveryAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryAddressCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetDeliveryCustomFieldAction struct {
	DeliveryId string `json:"deliveryId"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetDeliveryCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetDeliveryCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetDeliveryCustomTypeAction struct {
	DeliveryId string `json:"deliveryId"`
	// Defines the [Type](ctp:api:type:Type) that extends the Delivery with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Delivery.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Delivery.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetDeliveryCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetDeliveryCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetDeliveryItemsAction struct {
	DeliveryId string         `json:"deliveryId"`
	Items      []DeliveryItem `json:"items"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetDeliveryItemsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetDeliveryItemsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryItems", Alias: (*Alias)(&obj)})
}

type OrderSetItemShippingAddressCustomFieldAction struct {
	AddressKey string `json:"addressKey"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetItemShippingAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetItemShippingAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setItemShippingAddressCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetItemShippingAddressCustomTypeAction struct {
	AddressKey string `json:"addressKey"`
	// Defines the [Type](ctp:api:type:Type) that extends the `itemShippingAddress` with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the `itemShippingAddress`.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `itemShippingAddress`.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetItemShippingAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetItemShippingAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setItemShippingAddressCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetLineItemCustomFieldAction struct {
	LineItemId string `json:"lineItemId"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetLineItemCustomTypeAction struct {
	LineItemId string `json:"lineItemId"`
	// Defines the [Type](ctp:api:type:Type) that extends the LineItem with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the LineItem.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the LineItem.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetLineItemShippingDetailsAction struct {
	LineItemId      string                    `json:"lineItemId"`
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

type OrderSetLocaleAction struct {
	Locale *string `json:"locale,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

type OrderSetOrderNumberAction struct {
	OrderNumber *string `json:"orderNumber,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetOrderNumberAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetOrderNumberAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setOrderNumber", Alias: (*Alias)(&obj)})
}

type OrderSetParcelCustomFieldAction struct {
	ParcelId string `json:"parcelId"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetParcelCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetParcelCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetParcelCustomTypeAction struct {
	ParcelId string `json:"parcelId"`
	// Defines the [Type](ctp:api:type:Type) that extends the Parcel with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Parcel.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Parcel.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetParcelCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetParcelCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetParcelItemsAction struct {
	ParcelId string         `json:"parcelId"`
	Items    []DeliveryItem `json:"items"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetParcelItemsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetParcelItemsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelItems", Alias: (*Alias)(&obj)})
}

type OrderSetParcelMeasurementsAction struct {
	ParcelId     string              `json:"parcelId"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetParcelMeasurementsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetParcelMeasurementsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelMeasurements", Alias: (*Alias)(&obj)})
}

type OrderSetParcelTrackingDataAction struct {
	ParcelId     string        `json:"parcelId"`
	TrackingData *TrackingData `json:"trackingData,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetParcelTrackingDataAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetParcelTrackingDataAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelTrackingData", Alias: (*Alias)(&obj)})
}

type OrderSetReturnInfoAction struct {
	Items []ReturnInfoDraft `json:"items"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetReturnInfoAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetReturnInfoAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setReturnInfo", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["items"] == nil {
		delete(target, "items")
	}

	return json.Marshal(target)
}

type OrderSetReturnItemCustomFieldAction struct {
	ReturnItemId string `json:"returnItemId"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetReturnItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetReturnItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setReturnItemCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetReturnItemCustomTypeAction struct {
	ReturnItemId string `json:"returnItemId"`
	// Defines the [Type](ctp:api:type:Type) that extends the ReturnItem with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the ReturnItem.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the ReturnItem.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetReturnItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetReturnItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setReturnItemCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetReturnPaymentStateAction struct {
	ReturnItemId string             `json:"returnItemId"`
	PaymentState ReturnPaymentState `json:"paymentState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetReturnPaymentStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetReturnPaymentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setReturnPaymentState", Alias: (*Alias)(&obj)})
}

type OrderSetReturnShipmentStateAction struct {
	ReturnItemId  string              `json:"returnItemId"`
	ShipmentState ReturnShipmentState `json:"shipmentState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetReturnShipmentStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetReturnShipmentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setReturnShipmentState", Alias: (*Alias)(&obj)})
}

type OrderSetShippingAddressAction struct {
	Address *BaseAddress `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddress", Alias: (*Alias)(&obj)})
}

type OrderSetShippingAddressCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetShippingAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetShippingAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddressCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetShippingAddressCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the `shippingAddress` with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the `shippingAddress`.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `shippingAddress`.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetShippingAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetShippingAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddressCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetStoreAction struct {
	Store *StoreResourceIdentifier `json:"store,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetStoreAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetStoreAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStore", Alias: (*Alias)(&obj)})
}

type OrderTransitionCustomLineItemStateAction struct {
	CustomLineItemId string `json:"customLineItemId"`
	Quantity         int    `json:"quantity"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [State](ctp:api:type:State).
	FromState StateResourceIdentifier `json:"fromState"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [State](ctp:api:type:State).
	ToState              StateResourceIdentifier `json:"toState"`
	ActualTransitionDate *time.Time              `json:"actualTransitionDate,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderTransitionCustomLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderTransitionCustomLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionCustomLineItemState", Alias: (*Alias)(&obj)})
}

type OrderTransitionLineItemStateAction struct {
	LineItemId string `json:"lineItemId"`
	Quantity   int    `json:"quantity"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [State](ctp:api:type:State).
	FromState StateResourceIdentifier `json:"fromState"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [State](ctp:api:type:State).
	ToState              StateResourceIdentifier `json:"toState"`
	ActualTransitionDate *time.Time              `json:"actualTransitionDate,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderTransitionLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderTransitionLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionLineItemState", Alias: (*Alias)(&obj)})
}

type OrderTransitionStateAction struct {
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [State](ctp:api:type:State).
	State StateResourceIdentifier `json:"state"`
	Force *bool                   `json:"force,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}

type OrderUpdateItemShippingAddressAction struct {
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderUpdateItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderUpdateItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateItemShippingAddress", Alias: (*Alias)(&obj)})
}

type OrderUpdateSyncInfoAction struct {
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Channel](ctp:api:type:Channel).
	Channel    ChannelResourceIdentifier `json:"channel"`
	ExternalId *string                   `json:"externalId,omitempty"`
	SyncedAt   *time.Time                `json:"syncedAt,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderUpdateSyncInfoAction) MarshalJSON() ([]byte, error) {
	type Alias OrderUpdateSyncInfoAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateSyncInfo", Alias: (*Alias)(&obj)})
}
