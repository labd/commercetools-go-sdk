// Generated file, please do not change!!!
package platform

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
			return nil, errors.New("Error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "addCustomLineItem":
		new := StagedOrderAddCustomLineItemAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addDelivery":
		new := StagedOrderAddDeliveryAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addDiscountCode":
		new := StagedOrderAddDiscountCodeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addItemShippingAddress":
		new := StagedOrderAddItemShippingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addLineItem":
		new := StagedOrderAddLineItemAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addParcelToDelivery":
		new := StagedOrderAddParcelToDeliveryAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addPayment":
		new := StagedOrderAddPaymentAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addReturnInfo":
		new := StagedOrderAddReturnInfoAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addShoppingList":
		new := StagedOrderAddShoppingListAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeCustomLineItemMoney":
		new := StagedOrderChangeCustomLineItemMoneyAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeCustomLineItemQuantity":
		new := StagedOrderChangeCustomLineItemQuantityAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeLineItemQuantity":
		new := StagedOrderChangeLineItemQuantityAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeOrderState":
		new := StagedOrderChangeOrderStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changePaymentState":
		new := StagedOrderChangePaymentStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeShipmentState":
		new := StagedOrderChangeShipmentStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTaxCalculationMode":
		new := StagedOrderChangeTaxCalculationModeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTaxMode":
		new := StagedOrderChangeTaxModeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTaxRoundingMode":
		new := StagedOrderChangeTaxRoundingModeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "importCustomLineItemState":
		new := StagedOrderImportCustomLineItemStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "importLineItemState":
		new := StagedOrderImportLineItemStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeCustomLineItem":
		new := StagedOrderRemoveCustomLineItemAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeDelivery":
		new := StagedOrderRemoveDeliveryAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeDiscountCode":
		new := StagedOrderRemoveDiscountCodeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeItemShippingAddress":
		new := StagedOrderRemoveItemShippingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeLineItem":
		new := StagedOrderRemoveLineItemAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeParcelFromDelivery":
		new := StagedOrderRemoveParcelFromDeliveryAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removePayment":
		new := StagedOrderRemovePaymentAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setBillingAddress":
		new := StagedOrderSetBillingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setBillingAddressCustomField":
		new := StagedOrderSetBillingAddressCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setBillingAddressCustomType":
		new := StagedOrderSetBillingAddressCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCountry":
		new := StagedOrderSetCountryAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := StagedOrderSetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomLineItemCustomField":
		new := StagedOrderSetCustomLineItemCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomLineItemCustomType":
		new := StagedOrderSetCustomLineItemCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomLineItemShippingDetails":
		new := StagedOrderSetCustomLineItemShippingDetailsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomLineItemTaxAmount":
		new := StagedOrderSetCustomLineItemTaxAmountAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomLineItemTaxRate":
		new := StagedOrderSetCustomLineItemTaxRateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomShippingMethod":
		new := StagedOrderSetCustomShippingMethodAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := StagedOrderSetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomerEmail":
		new := StagedOrderSetCustomerEmailAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomerGroup":
		new := StagedOrderSetCustomerGroupAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomerId":
		new := StagedOrderSetCustomerIdAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeliveryAddress":
		new := StagedOrderSetDeliveryAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeliveryAddressCustomField":
		new := StagedOrderSetDeliveryAddressCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeliveryAddressCustomType":
		new := StagedOrderSetDeliveryAddressCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeliveryItems":
		new := StagedOrderSetDeliveryItemsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setItemShippingAddressCustomField":
		new := StagedOrderSetItemShippingAddressCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setItemShippingAddressCustomType":
		new := StagedOrderSetItemShippingAddressCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomField":
		new := StagedOrderSetLineItemCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomType":
		new := StagedOrderSetLineItemCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemDistributionChannel":
		new := StagedOrderSetLineItemDistributionChannelAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemPrice":
		new := StagedOrderSetLineItemPriceAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemShippingDetails":
		new := StagedOrderSetLineItemShippingDetailsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemTaxAmount":
		new := StagedOrderSetLineItemTaxAmountAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemTaxRate":
		new := StagedOrderSetLineItemTaxRateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemTotalPrice":
		new := StagedOrderSetLineItemTotalPriceAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLocale":
		new := StagedOrderSetLocaleAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setOrderNumber":
		new := StagedOrderSetOrderNumberAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setOrderTotalTax":
		new := StagedOrderSetOrderTotalTaxAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setParcelItems":
		new := StagedOrderSetParcelItemsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setParcelMeasurements":
		new := StagedOrderSetParcelMeasurementsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setParcelTrackingData":
		new := StagedOrderSetParcelTrackingDataAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setReturnPaymentState":
		new := StagedOrderSetReturnPaymentStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setReturnShipmentState":
		new := StagedOrderSetReturnShipmentStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingAddress":
		new := StagedOrderSetShippingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingAddressAndCustomShippingMethod":
		new := StagedOrderSetShippingAddressAndCustomShippingMethodAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingAddressAndShippingMethod":
		new := StagedOrderSetShippingAddressAndShippingMethodAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingAddressCustomField":
		new := StagedOrderSetShippingAddressCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingAddressCustomType":
		new := StagedOrderSetShippingAddressCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingMethod":
		new := StagedOrderSetShippingMethodAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingMethodTaxAmount":
		new := StagedOrderSetShippingMethodTaxAmountAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingMethodTaxRate":
		new := StagedOrderSetShippingMethodTaxRateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingRateInput":
		new := StagedOrderSetShippingRateInputAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "transitionCustomLineItemState":
		new := StagedOrderTransitionCustomLineItemStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "transitionLineItemState":
		new := StagedOrderTransitionLineItemStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "transitionState":
		new := StagedOrderTransitionStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "updateItemShippingAddress":
		new := StagedOrderUpdateItemShippingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "updateSyncInfo":
		new := StagedOrderUpdateSyncInfoAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type Delivery struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	// Items which are shipped in this delivery regardless their distribution over several parcels.
	// Can also be specified individually for each Parcel.
	Items   []DeliveryItem `json:"items"`
	Parcels []Parcel       `json:"parcels"`
	Address *Address       `json:"address,omitEmpty"`
}

type DeliveryItem struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"`
}

type DiscountedLineItemPriceDraft struct {
	Value             Money                       `json:"value"`
	IncludedDiscounts []DiscountedLineItemPortion `json:"includedDiscounts"`
}

type ItemState struct {
	Quantity float64        `json:"quantity"`
	State    StateReference `json:"state"`
}

type LineItemImportDraft struct {
	// ID of the existing product.
	// You also need to specify the ID of the variant if this property is set or alternatively you can just specify SKU of the product variant.
	ProductId string `json:"productId,omitEmpty"`
	// The product name.
	Name     LocalizedString           `json:"name"`
	Variant  ProductVariantImportDraft `json:"variant"`
	Price    PriceDraft                `json:"price"`
	Quantity float64                   `json:"quantity"`
	State    []ItemState               `json:"state,omitEmpty"`
	// Optional connection to a particular supplier.
	// By providing supply channel information, you can uniquely identify
	// inventory entries that should be reserved.
	// The provided channel should have the
	// InventorySupply role.
	SupplyChannel ChannelResourceIdentifier `json:"supplyChannel,omitEmpty"`
	// The channel is used to select a ProductPrice.
	// The provided channel should have the ProductDistribution role.
	DistributionChannel ChannelResourceIdentifier `json:"distributionChannel,omitEmpty"`
	TaxRate             *TaxRate                  `json:"taxRate,omitEmpty"`
	// The custom fields.
	Custom          *CustomFieldsDraft        `json:"custom,omitEmpty"`
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitEmpty"`
}

type Order struct {
	// The unique ID of the order.
	Id string `json:"id"`
	// The current version of the order.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for events not tracked.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	// Present on resources created after 1/02/2019 except for events not tracked.
	CreatedBy *CreatedBy `json:"createdBy,omitEmpty"`
	// This field will only be present if it was set for Order Import
	CompletedAt time.Time `json:"completedAt,omitEmpty"`
	// String that uniquely identifies an order.
	// It can be used to create more human-readable (in contrast to ID) identifier for the order.
	// It should be unique across a project.
	// Once it's set it cannot be changed.
	OrderNumber   string `json:"orderNumber,omitEmpty"`
	CustomerId    string `json:"customerId,omitEmpty"`
	CustomerEmail string `json:"customerEmail,omitEmpty"`
	// Identifies carts and orders belonging to an anonymous session (the customer has not signed up/in yet).
	AnonymousId     string            `json:"anonymousId,omitEmpty"`
	Store           StoreKeyReference `json:"store,omitEmpty"`
	LineItems       []LineItem        `json:"lineItems"`
	CustomLineItems []CustomLineItem  `json:"customLineItems"`
	TotalPrice      TypedMoney        `json:"totalPrice"`
	// The taxes are calculated based on the shipping address.
	TaxedPrice      *TaxedPrice `json:"taxedPrice,omitEmpty"`
	ShippingAddress *Address    `json:"shippingAddress,omitEmpty"`
	BillingAddress  *Address    `json:"billingAddress,omitEmpty"`
	TaxMode         TaxMode     `json:"taxMode,omitEmpty"`
	// When calculating taxes for `taxedPrice`, the selected mode is used for rouding.
	TaxRoundingMode RoundingMode `json:"taxRoundingMode,omitEmpty"`
	// Set when the customer is set and the customer is a member of a customer group.
	// Used for product variant price selection.
	CustomerGroup CustomerGroupReference `json:"customerGroup,omitEmpty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	// Used for product variant price selection.
	Country string `json:"country,omitEmpty"`
	// One of the four predefined OrderStates.
	OrderState OrderState `json:"orderState"`
	// This reference can point to a state in a custom workflow.
	State         StateReference `json:"state,omitEmpty"`
	ShipmentState ShipmentState  `json:"shipmentState,omitEmpty"`
	PaymentState  PaymentState   `json:"paymentState,omitEmpty"`
	// Set if the ShippingMethod is set.
	ShippingInfo  *ShippingInfo      `json:"shippingInfo,omitEmpty"`
	SyncInfo      []SyncInfo         `json:"syncInfo"`
	ReturnInfo    []ReturnInfo       `json:"returnInfo,omitEmpty"`
	DiscountCodes []DiscountCodeInfo `json:"discountCodes,omitEmpty"`
	// The sequence number of the last order message produced by changes to this order.
	// `0` means, that no messages were created yet.
	LastMessageSequenceNumber int `json:"lastMessageSequenceNumber"`
	// Set when this order was created from a cart.
	// The cart will have the state `Ordered`.
	Cart          CartReference `json:"cart,omitEmpty"`
	Custom        *CustomFields `json:"custom,omitEmpty"`
	PaymentInfo   *PaymentInfo  `json:"paymentInfo,omitEmpty"`
	Locale        string        `json:"locale,omitEmpty"`
	InventoryMode InventoryMode `json:"inventoryMode,omitEmpty"`
	Origin        CartOrigin    `json:"origin"`
	// When calculating taxes for `taxedPrice`, the selected mode is used for calculating the price with LineItemLevel (horizontally) or UnitPriceLevel (vertically) calculation mode.
	TaxCalculationMode TaxCalculationMode `json:"taxCalculationMode,omitEmpty"`
	// The shippingRateInput is used as an input to select a ShippingRatePriceTier.
	ShippingRateInput ShippingRateInput `json:"shippingRateInput,omitEmpty"`
	// Contains addresses for orders with multiple shipping addresses.
	ItemShippingAddresses []Address `json:"itemShippingAddresses,omitEmpty"`
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

type OrderFromCartDraft struct {
	// The unique id of the cart from which an order is created.
	Id string `json:"id,omitEmpty"`
	// ResourceIdentifier to the Cart from which this order is created.
	Cart    CartResourceIdentifier `json:"cart,omitEmpty"`
	Version int                    `json:"version"`
	// String that uniquely identifies an order.
	// It can be used to create more human-readable (in contrast to ID) identifier for the order.
	// It should be unique across a project.
	// Once it's set it cannot be changed.
	// For easier use on Get, Update and Delete actions we suggest assigning order numbers that match the regular expression `[a-z0-9_\-]{2,36}`.
	OrderNumber   string        `json:"orderNumber,omitEmpty"`
	PaymentState  PaymentState  `json:"paymentState,omitEmpty"`
	ShipmentState ShipmentState `json:"shipmentState,omitEmpty"`
	// Order will be created with `Open` status by default.
	OrderState OrderState              `json:"orderState,omitEmpty"`
	State      StateResourceIdentifier `json:"state,omitEmpty"`
}

type OrderImportDraft struct {
	// String that unique identifies an order.
	// It can be used to create more human-readable (in contrast to ID) identifier for the order.
	// It should be unique within a project.
	OrderNumber string `json:"orderNumber,omitEmpty"`
	// If given the customer with that ID must exist in the project.
	CustomerId string `json:"customerId,omitEmpty"`
	// The customer email can be used when no check against existing Customers is desired during order import.
	CustomerEmail string `json:"customerEmail,omitEmpty"`
	// If not given `customLineItems` must not be empty.
	LineItems []LineItemImportDraft `json:"lineItems,omitEmpty"`
	// If not given `lineItems` must not be empty.
	CustomLineItems []CustomLineItemDraft `json:"customLineItems,omitEmpty"`
	TotalPrice      Money                 `json:"totalPrice"`
	// Order Import does not support calculation of taxes.
	// When setting the draft the taxedPrice is to be provided.
	TaxedPrice      *TaxedPriceDraft `json:"taxedPrice,omitEmpty"`
	ShippingAddress *BaseAddress     `json:"shippingAddress,omitEmpty"`
	BillingAddress  *BaseAddress     `json:"billingAddress,omitEmpty"`
	// Set when the customer is set and the customer is a member of a customer group.
	// Used for product variant price selection.
	CustomerGroup CustomerGroupResourceIdentifier `json:"customerGroup,omitEmpty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	// Used for product variant price selection.
	Country string `json:"country,omitEmpty"`
	// If not given the `Open` state will be assigned by default.
	OrderState    OrderState    `json:"orderState,omitEmpty"`
	ShipmentState ShipmentState `json:"shipmentState,omitEmpty"`
	PaymentState  PaymentState  `json:"paymentState,omitEmpty"`
	// Set if the ShippingMethod is set.
	ShippingInfo *ShippingInfoImportDraft `json:"shippingInfo,omitEmpty"`
	CompletedAt  time.Time                `json:"completedAt,omitEmpty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitEmpty"`
	// If not given the mode `None` will be assigned by default.
	InventoryMode InventoryMode `json:"inventoryMode,omitEmpty"`
	// If not given the tax rounding mode `HalfEven` will be assigned by default.
	TaxRoundingMode RoundingMode `json:"taxRoundingMode,omitEmpty"`
	// Contains addresses for orders with multiple shipping addresses.
	ItemShippingAddresses []BaseAddress           `json:"itemShippingAddresses,omitEmpty"`
	Store                 StoreResourceIdentifier `json:"store,omitEmpty"`
	// The default origin is `Customer`.
	Origin CartOrigin `json:"origin,omitEmpty"`
}

type OrderPagedQueryResponse struct {
	Limit   int     `json:"limit"`
	Count   int     `json:"count"`
	Total   int     `json:"total,omitEmpty"`
	Offset  int     `json:"offset"`
	Results []Order `json:"results"`
}

type OrderReference struct {
	Id  string `json:"id"`
	Obj *Order `json:"obj,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderReference) MarshalJSON() ([]byte, error) {
	type Alias OrderReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "order", Alias: (*Alias)(&obj)})
}

type OrderResourceIdentifier struct {
	Id  string `json:"id,omitEmpty"`
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias OrderResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "order", Alias: (*Alias)(&obj)})
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

	return nil
}

type OrderUpdateAction interface{}

func mapDiscriminatorOrderUpdateAction(input interface{}) (OrderUpdateAction, error) {

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
	case "addDelivery":
		new := OrderAddDeliveryAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addItemShippingAddress":
		new := OrderAddItemShippingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addParcelToDelivery":
		new := OrderAddParcelToDeliveryAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addPayment":
		new := OrderAddPaymentAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addReturnInfo":
		new := OrderAddReturnInfoAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeOrderState":
		new := OrderChangeOrderStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changePaymentState":
		new := OrderChangePaymentStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeShipmentState":
		new := OrderChangeShipmentStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "importCustomLineItemState":
		new := OrderImportCustomLineItemStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "importLineItemState":
		new := OrderImportLineItemStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeDelivery":
		new := OrderRemoveDeliveryAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeItemShippingAddress":
		new := OrderRemoveItemShippingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeParcelFromDelivery":
		new := OrderRemoveParcelFromDeliveryAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removePayment":
		new := OrderRemovePaymentAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setBillingAddress":
		new := OrderSetBillingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setBillingAddressCustomField":
		new := OrderSetBillingAddressCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setBillingAddressCustomType":
		new := OrderSetBillingAddressCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := OrderSetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomLineItemCustomField":
		new := OrderSetCustomLineItemCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomLineItemCustomType":
		new := OrderSetCustomLineItemCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomLineItemShippingDetails":
		new := OrderSetCustomLineItemShippingDetailsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := OrderSetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomerEmail":
		new := OrderSetCustomerEmailAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomerId":
		new := OrderSetCustomerIdAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeliveryAddress":
		new := OrderSetDeliveryAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeliveryAddressCustomField":
		new := OrderSetDeliveryAddressCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeliveryAddressCustomType":
		new := OrderSetDeliveryAddressCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDeliveryItems":
		new := OrderSetDeliveryItemsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setItemShippingAddressCustomField":
		new := OrderSetItemShippingAddressCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setItemShippingAddressCustomType":
		new := OrderSetItemShippingAddressCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomField":
		new := OrderSetLineItemCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemCustomType":
		new := OrderSetLineItemCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLineItemShippingDetails":
		new := OrderSetLineItemShippingDetailsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLocale":
		new := OrderSetLocaleAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setOrderNumber":
		new := OrderSetOrderNumberAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setParcelItems":
		new := OrderSetParcelItemsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setParcelMeasurements":
		new := OrderSetParcelMeasurementsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setParcelTrackingData":
		new := OrderSetParcelTrackingDataAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setReturnPaymentState":
		new := OrderSetReturnPaymentStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setReturnShipmentState":
		new := OrderSetReturnShipmentStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingAddress":
		new := OrderSetShippingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingAddressCustomField":
		new := OrderSetShippingAddressCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setShippingAddressCustomType":
		new := OrderSetShippingAddressCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setStore":
		new := OrderSetStoreAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "transitionCustomLineItemState":
		new := OrderTransitionCustomLineItemStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "transitionLineItemState":
		new := OrderTransitionLineItemStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "transitionState":
		new := OrderTransitionStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "updateItemShippingAddress":
		new := OrderUpdateItemShippingAddressAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "updateSyncInfo":
		new := OrderUpdateSyncInfoAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type Parcel struct {
	Id           string              `json:"id"`
	CreatedAt    time.Time           `json:"createdAt"`
	Measurements *ParcelMeasurements `json:"measurements,omitEmpty"`
	TrackingData *TrackingData       `json:"trackingData,omitEmpty"`
	// The delivery items contained in this parcel.
	Items []DeliveryItem `json:"items,omitEmpty"`
}

type ParcelDraft struct {
	Measurements *ParcelMeasurements `json:"measurements,omitEmpty"`
	TrackingData *TrackingData       `json:"trackingData,omitEmpty"`
	// The delivery items contained in this parcel.
	Items []DeliveryItem `json:"items,omitEmpty"`
}

type ParcelMeasurements struct {
	HeightInMillimeter float64 `json:"heightInMillimeter,omitEmpty"`
	LengthInMillimeter float64 `json:"lengthInMillimeter,omitEmpty"`
	WidthInMillimeter  float64 `json:"widthInMillimeter,omitEmpty"`
	WeightInGram       float64 `json:"weightInGram,omitEmpty"`
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
	Id int `json:"id,omitEmpty"`
	// The SKU of the existing variant.
	Sku string `json:"sku,omitEmpty"`
	// The prices of the variant.
	// The prices should not contain two prices for the same price scope (same currency, country and customer group).
	// If this property is defined, then it will override the `prices` property from the original product variant, otherwise `prices` property from the original product variant would be copied in the resulting order.
	Prices []PriceDraft `json:"prices,omitEmpty"`
	// If this property is defined, then it will override the `attributes` property from the original
	// product variant, otherwise `attributes` property from the original product variant would be copied in the resulting order.
	Attributes []Attribute `json:"attributes,omitEmpty"`
	// If this property is defined, then it will override the `images` property from the original
	// product variant, otherwise `images` property from the original product variant would be copied in the resulting order.
	Images []Image `json:"images,omitEmpty"`
}

type ReturnInfo struct {
	Items []ReturnItem `json:"items"`
	// Identifies, which return tracking ID is connected to this particular return.
	ReturnTrackingId string    `json:"returnTrackingId,omitEmpty"`
	ReturnDate       time.Time `json:"returnDate,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReturnInfo) UnmarshalJSON(data []byte) error {
	type Alias ReturnInfo
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type ReturnItem interface{}

func mapDiscriminatorReturnItem(input interface{}) (ReturnItem, error) {

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
	case "CustomLineItemReturnItem":
		new := CustomLineItemReturnItem{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "LineItemReturnItem":
		new := LineItemReturnItem{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type CustomLineItemReturnItem struct {
	Id               string              `json:"id"`
	Quantity         int                 `json:"quantity"`
	Comment          string              `json:"comment,omitEmpty"`
	ShipmentState    ReturnShipmentState `json:"shipmentState"`
	PaymentState     ReturnPaymentState  `json:"paymentState"`
	LastModifiedAt   time.Time           `json:"lastModifiedAt"`
	CreatedAt        time.Time           `json:"createdAt"`
	CustomLineItemId string              `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomLineItemReturnItem) MarshalJSON() ([]byte, error) {
	type Alias CustomLineItemReturnItem
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CustomLineItemReturnItem", Alias: (*Alias)(&obj)})
}

type LineItemReturnItem struct {
	Id             string              `json:"id"`
	Quantity       int                 `json:"quantity"`
	Comment        string              `json:"comment,omitEmpty"`
	ShipmentState  ReturnShipmentState `json:"shipmentState"`
	PaymentState   ReturnPaymentState  `json:"paymentState"`
	LastModifiedAt time.Time           `json:"lastModifiedAt"`
	CreatedAt      time.Time           `json:"createdAt"`
	LineItemId     string              `json:"lineItemId"`
}

// MarshalJSON override to set the discriminator value
func (obj LineItemReturnItem) MarshalJSON() ([]byte, error) {
	type Alias LineItemReturnItem
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "LineItemReturnItem", Alias: (*Alias)(&obj)})
}

type ReturnItemDraft struct {
	Quantity         int                 `json:"quantity"`
	LineItemId       string              `json:"lineItemId,omitEmpty"`
	CustomLineItemId string              `json:"customLineItemId,omitEmpty"`
	Comment          string              `json:"comment,omitEmpty"`
	ShipmentState    ReturnShipmentState `json:"shipmentState"`
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
	ShippingRate ShippingRateDraft             `json:"shippingRate"`
	TaxRate      *TaxRate                      `json:"taxRate,omitEmpty"`
	TaxCategory  TaxCategoryResourceIdentifier `json:"taxCategory,omitEmpty"`
	// Not set if custom shipping method is used.
	ShippingMethod ShippingMethodResourceIdentifier `json:"shippingMethod,omitEmpty"`
	// Deliveries are compilations of information on how the articles are being delivered to the customers.
	Deliveries      []Delivery                    `json:"deliveries,omitEmpty"`
	DiscountedPrice *DiscountedLineItemPriceDraft `json:"discountedPrice,omitEmpty"`
	// Indicates whether the ShippingMethod referenced is allowed for the cart or not.
	ShippingMethodState ShippingMethodState `json:"shippingMethodState,omitEmpty"`
}

type SyncInfo struct {
	// Connection to a particular synchronization destination.
	Channel ChannelReference `json:"channel"`
	// Can be used to reference an external order instance, file etc.
	ExternalId string    `json:"externalId,omitEmpty"`
	SyncedAt   time.Time `json:"syncedAt"`
}

type TaxedItemPriceDraft struct {
	TotalNet   Money `json:"totalNet"`
	TotalGross Money `json:"totalGross"`
}

type TrackingData struct {
	// The ID to track one parcel.
	TrackingId string `json:"trackingId,omitEmpty"`
	// The carrier that delivers the parcel.
	Carrier             string `json:"carrier,omitEmpty"`
	Provider            string `json:"provider,omitEmpty"`
	ProviderTransaction string `json:"providerTransaction,omitEmpty"`
	// Flag to distinguish if the parcel is on the way to the customer (false) or on the way back (true).
	IsReturn bool `json:"isReturn,omitEmpty"`
}

type OrderAddDeliveryAction struct {
	Items   []DeliveryItem `json:"items,omitEmpty"`
	Address *BaseAddress   `json:"address,omitEmpty"`
	Parcels []ParcelDraft  `json:"parcels,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderAddDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDelivery", Alias: (*Alias)(&obj)})
}

type OrderAddItemShippingAddressAction struct {
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderAddItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addItemShippingAddress", Alias: (*Alias)(&obj)})
}

type OrderAddParcelToDeliveryAction struct {
	DeliveryId   string              `json:"deliveryId"`
	Measurements *ParcelMeasurements `json:"measurements,omitEmpty"`
	TrackingData *TrackingData       `json:"trackingData,omitEmpty"`
	Items        []DeliveryItem      `json:"items,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderAddParcelToDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddParcelToDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addParcelToDelivery", Alias: (*Alias)(&obj)})
}

type OrderAddPaymentAction struct {
	Payment PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderAddPaymentAction) MarshalJSON() ([]byte, error) {
	type Alias OrderAddPaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPayment", Alias: (*Alias)(&obj)})
}

type OrderAddReturnInfoAction struct {
	ReturnTrackingId string            `json:"returnTrackingId,omitEmpty"`
	Items            []ReturnItemDraft `json:"items"`
	ReturnDate       time.Time         `json:"returnDate,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
func (obj OrderChangeOrderStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderChangeOrderStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeOrderState", Alias: (*Alias)(&obj)})
}

type OrderChangePaymentStateAction struct {
	PaymentState PaymentState `json:"paymentState,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderChangePaymentStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderChangePaymentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePaymentState", Alias: (*Alias)(&obj)})
}

type OrderChangeShipmentStateAction struct {
	ShipmentState ShipmentState `json:"shipmentState,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
func (obj OrderRemoveParcelFromDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias OrderRemoveParcelFromDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeParcelFromDelivery", Alias: (*Alias)(&obj)})
}

type OrderRemovePaymentAction struct {
	Payment PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderRemovePaymentAction) MarshalJSON() ([]byte, error) {
	type Alias OrderRemovePaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePayment", Alias: (*Alias)(&obj)})
}

type OrderSetBillingAddressAction struct {
	Address *BaseAddress `json:"address,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddress", Alias: (*Alias)(&obj)})
}

type OrderSetBillingAddressCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetBillingAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetBillingAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddressCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetBillingAddressCustomTypeAction struct {
	Type   TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields *FieldContainer        `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetBillingAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetBillingAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddressCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetCustomLineItemCustomFieldAction struct {
	CustomLineItemId string      `json:"customLineItemId"`
	Name             string      `json:"name"`
	Value            interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetCustomLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetCustomLineItemCustomTypeAction struct {
	CustomLineItemId string                 `json:"customLineItemId"`
	Type             TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields           *FieldContainer        `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetCustomLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetCustomLineItemShippingDetailsAction struct {
	CustomLineItemId string                    `json:"customLineItemId"`
	ShippingDetails  *ItemShippingDetailsDraft `json:"shippingDetails,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetCustomLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

type OrderSetCustomTypeAction struct {
	Type   TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields *FieldContainer        `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetCustomerEmailAction struct {
	Email string `json:"email,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetCustomerEmailAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomerEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerEmail", Alias: (*Alias)(&obj)})
}

type OrderSetCustomerIdAction struct {
	CustomerId string `json:"customerId,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetCustomerIdAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetCustomerIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerId", Alias: (*Alias)(&obj)})
}

type OrderSetDeliveryAddressAction struct {
	DeliveryId string       `json:"deliveryId"`
	Address    *BaseAddress `json:"address,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetDeliveryAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetDeliveryAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryAddress", Alias: (*Alias)(&obj)})
}

type OrderSetDeliveryAddressCustomFieldAction struct {
	DeliveryId string      `json:"deliveryId"`
	Name       string      `json:"name"`
	Value      interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetDeliveryAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetDeliveryAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryAddressCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetDeliveryAddressCustomTypeAction struct {
	DeliveryId string                 `json:"deliveryId"`
	Type       TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields     *FieldContainer        `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetDeliveryAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetDeliveryAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryAddressCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetDeliveryItemsAction struct {
	DeliveryId string         `json:"deliveryId"`
	Items      []DeliveryItem `json:"items"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetDeliveryItemsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetDeliveryItemsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryItems", Alias: (*Alias)(&obj)})
}

type OrderSetItemShippingAddressCustomFieldAction struct {
	AddressKey string      `json:"addressKey"`
	Name       string      `json:"name"`
	Value      interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetItemShippingAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetItemShippingAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setItemShippingAddressCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetItemShippingAddressCustomTypeAction struct {
	AddressKey string                 `json:"addressKey"`
	Type       TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields     *FieldContainer        `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetItemShippingAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetItemShippingAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setItemShippingAddressCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetLineItemCustomFieldAction struct {
	LineItemId string      `json:"lineItemId"`
	Name       string      `json:"name"`
	Value      interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetLineItemCustomTypeAction struct {
	LineItemId string                 `json:"lineItemId"`
	Type       TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields     *FieldContainer        `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetLineItemShippingDetailsAction struct {
	LineItemId      string                    `json:"lineItemId"`
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

type OrderSetLocaleAction struct {
	Locale string `json:"locale,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

type OrderSetOrderNumberAction struct {
	OrderNumber string `json:"orderNumber,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetOrderNumberAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetOrderNumberAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setOrderNumber", Alias: (*Alias)(&obj)})
}

type OrderSetParcelItemsAction struct {
	ParcelId string         `json:"parcelId"`
	Items    []DeliveryItem `json:"items"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetParcelItemsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetParcelItemsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelItems", Alias: (*Alias)(&obj)})
}

type OrderSetParcelMeasurementsAction struct {
	ParcelId     string              `json:"parcelId"`
	Measurements *ParcelMeasurements `json:"measurements,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetParcelMeasurementsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetParcelMeasurementsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelMeasurements", Alias: (*Alias)(&obj)})
}

type OrderSetParcelTrackingDataAction struct {
	ParcelId     string        `json:"parcelId"`
	TrackingData *TrackingData `json:"trackingData,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetParcelTrackingDataAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetParcelTrackingDataAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelTrackingData", Alias: (*Alias)(&obj)})
}

type OrderSetReturnPaymentStateAction struct {
	ReturnItemId string             `json:"returnItemId"`
	PaymentState ReturnPaymentState `json:"paymentState"`
}

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
func (obj OrderSetReturnShipmentStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetReturnShipmentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setReturnShipmentState", Alias: (*Alias)(&obj)})
}

type OrderSetShippingAddressAction struct {
	Address *BaseAddress `json:"address,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddress", Alias: (*Alias)(&obj)})
}

type OrderSetShippingAddressCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetShippingAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetShippingAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddressCustomField", Alias: (*Alias)(&obj)})
}

type OrderSetShippingAddressCustomTypeAction struct {
	Type   TypeResourceIdentifier `json:"type,omitEmpty"`
	Fields *FieldContainer        `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetShippingAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetShippingAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddressCustomType", Alias: (*Alias)(&obj)})
}

type OrderSetStoreAction struct {
	Store StoreResourceIdentifier `json:"store,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderSetStoreAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetStoreAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStore", Alias: (*Alias)(&obj)})
}

type OrderTransitionCustomLineItemStateAction struct {
	CustomLineItemId     string                  `json:"customLineItemId"`
	Quantity             int                     `json:"quantity"`
	FromState            StateResourceIdentifier `json:"fromState"`
	ToState              StateResourceIdentifier `json:"toState"`
	ActualTransitionDate time.Time               `json:"actualTransitionDate,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderTransitionCustomLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderTransitionCustomLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionCustomLineItemState", Alias: (*Alias)(&obj)})
}

type OrderTransitionLineItemStateAction struct {
	LineItemId           string                  `json:"lineItemId"`
	Quantity             int                     `json:"quantity"`
	FromState            StateResourceIdentifier `json:"fromState"`
	ToState              StateResourceIdentifier `json:"toState"`
	ActualTransitionDate time.Time               `json:"actualTransitionDate,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderTransitionLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias OrderTransitionLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionLineItemState", Alias: (*Alias)(&obj)})
}

type OrderTransitionStateAction struct {
	State StateResourceIdentifier `json:"state"`
	Force bool                    `json:"force,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
func (obj OrderUpdateItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias OrderUpdateItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateItemShippingAddress", Alias: (*Alias)(&obj)})
}

type OrderUpdateSyncInfoAction struct {
	Channel    ChannelResourceIdentifier `json:"channel"`
	ExternalId string                    `json:"externalId,omitEmpty"`
	SyncedAt   time.Time                 `json:"syncedAt,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj OrderUpdateSyncInfoAction) MarshalJSON() ([]byte, error) {
	type Alias OrderUpdateSyncInfoAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateSyncInfo", Alias: (*Alias)(&obj)})
}
