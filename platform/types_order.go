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
	case "setDirectDiscounts":
		obj := StagedOrderSetDirectDiscountsAction{}
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
	case "setPurchaseOrderNumber":
		obj := StagedOrderSetPurchaseOrderNumberAction{}
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
	case "setStore":
		obj := StagedOrderSetStoreAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
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
	Relevance *float64 `json:"relevance,omitempty"`
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

type OrderSearchMatchType string

const (
	OrderSearchMatchTypeAny OrderSearchMatchType = "any"
	OrderSearchMatchTypeAll OrderSearchMatchType = "all"
)

type OrderSearchQueryExpressionValue struct {
	Field      string  `json:"field"`
	Boost      *int    `json:"boost,omitempty"`
	CustomType *string `json:"customType,omitempty"`
}

type OrderSearchAnyValue struct {
	Field           string      `json:"field"`
	Boost           *int        `json:"boost,omitempty"`
	CustomType      *string     `json:"customType,omitempty"`
	Value           interface{} `json:"value"`
	Language        *string     `json:"language,omitempty"`
	CaseInsensitive *bool       `json:"caseInsensitive,omitempty"`
}

type OrderSearchDateRangeValue struct {
	Field      string     `json:"field"`
	Boost      *int       `json:"boost,omitempty"`
	CustomType *string    `json:"customType,omitempty"`
	Gte        *time.Time `json:"gte,omitempty"`
	Lte        *time.Time `json:"lte,omitempty"`
}

type OrderSearchFullTextValue struct {
	Field      string                `json:"field"`
	Boost      *int                  `json:"boost,omitempty"`
	CustomType *string               `json:"customType,omitempty"`
	Value      string                `json:"value"`
	Language   *string               `json:"language,omitempty"`
	MustMatch  *OrderSearchMatchType `json:"mustMatch,omitempty"`
}

type OrderSearchLongRangeValue struct {
	Field      string  `json:"field"`
	Boost      *int    `json:"boost,omitempty"`
	CustomType *string `json:"customType,omitempty"`
	Gte        *int    `json:"gte,omitempty"`
	Lte        *int    `json:"lte,omitempty"`
}

type OrderSearchNumberRangeValue struct {
	Field      string   `json:"field"`
	Boost      *int     `json:"boost,omitempty"`
	CustomType *string  `json:"customType,omitempty"`
	Gte        *float64 `json:"gte,omitempty"`
	Lte        *float64 `json:"lte,omitempty"`
}

type OrderSearchSortMode string

const (
	OrderSearchSortModeMin OrderSearchSortMode = "min"
	OrderSearchSortModeMax OrderSearchSortMode = "max"
	OrderSearchSortModeAvg OrderSearchSortMode = "avg"
	OrderSearchSortModeSum OrderSearchSortMode = "sum"
)

type OrderSearchSortOrder string

const (
	OrderSearchSortOrderAsc  OrderSearchSortOrder = "asc"
	OrderSearchSortOrderDesc OrderSearchSortOrder = "desc"
)

type OrderSearchStringValue struct {
	Field           string  `json:"field"`
	Boost           *int    `json:"boost,omitempty"`
	CustomType      *string `json:"customType,omitempty"`
	Value           string  `json:"value"`
	Language        *string `json:"language,omitempty"`
	CaseInsensitive *bool   `json:"caseInsensitive,omitempty"`
}

/**
*	Custom Line Items contain generic user-defined items that are not linked to Products.
*
 */
type CustomLineItemImportDraft struct {
	// Name of the Custom Line Item.
	Name LocalizedString `json:"name"`
	// User-defined unique identifier of the Custom Line Item.
	Key *string `json:"key,omitempty"`
	// User-defined identifier used in a deep-link URL for the Custom Line Item. This value should match the pattern `[a-zA-Z0-9_-]{2,256}`.
	Slug string `json:"slug"`
	// The number of items in the Custom Line Item. Can be a negative value.
	Quantity int `json:"quantity"`
	// The cost of individual items in the Custom Line Item. The amount can be negative.
	Money Money `json:"money"`
	// The tax rate used to calculate the `taxedPrice` of the Order.
	TaxRate *TaxRate `json:"taxRate,omitempty"`
	// Include a value to associate a Tax Category with the Custom Line Item.
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// - If `Standard`, Cart Discounts with a matching [CartDiscountCustomLineItemsTarget](ctp:api:type:CartDiscountCustomLineItemsTarget)
	// are applied to the Custom Line Item.
	// - If `External`, Cart Discounts are not considered on the Custom Line Item.
	PriceMode *CustomLineItemPriceMode `json:"priceMode,omitempty"`
	// Container for Custom Line Item-specific addresses.
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	// State of the Custom Line Items.
	State []ItemState `json:"state"`
	// Custom Fields of the CustomLineItem.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomLineItemImportDraft) MarshalJSON() ([]byte, error) {
	type Alias CustomLineItemImportDraft
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

/**
*	Contains information on how items are shipped to Customers, for example, a delivery note.
*
 */
type Delivery struct {
	// Unique identifier of the Delivery.
	ID string `json:"id"`
	// User-defined unique identifier of the Delivery.
	Key *string `json:"key,omitempty"`
	// Date and time (UTC) the Delivery was created.
	CreatedAt time.Time `json:"createdAt"`
	// Line Items or Custom Line Items that are delivered.
	Items []DeliveryItem `json:"items"`
	// Information regarding the appearance, content, and shipment of a Parcel.
	Parcels []Parcel `json:"parcels"`
	// Address to which Parcels are delivered.
	Address *Address `json:"address,omitempty"`
	// Custom Fields of the Delivery.
	Custom *CustomFields `json:"custom,omitempty"`
}

type DeliveryDraft struct {
	// User-defined unique identifier of the Delivery.
	Key *string `json:"key,omitempty"`
	// Line Items or Custom Line Items to deliver.
	// It can also be specified individually for each [Parcel](ctp:api:type:Parcel).
	Items []DeliveryItem `json:"items"`
	// Information regarding the appearance, content, and shipment of a parcel.
	Parcels []ParcelDraft `json:"parcels"`
	// Address to which the Parcels are delivered.
	Address *AddressDraft `json:"address,omitempty"`
	// Custom Fields for the Delivery.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DeliveryDraft) MarshalJSON() ([]byte, error) {
	type Alias DeliveryDraft
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

	if raw["parcels"] == nil {
		delete(raw, "parcels")
	}

	return json.Marshal(raw)

}

type DeliveryItem struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) or [CustomLineItem](ctp:api:type:CustomLineItem) delivered.
	ID string `json:"id"`
	// Number of Line Items or Custom Line Items delivered.
	Quantity int `json:"quantity"`
}

type DiscountedLineItemPriceDraft struct {
	// Discounted money value.
	Value Money `json:"value"`
	// Discounts to be applied.
	IncludedDiscounts []DiscountedLineItemPortionDraft `json:"includedDiscounts"`
}

type ItemState struct {
	// Number of Line Items or Custom Line Items in this State.
	Quantity int `json:"quantity"`
	// State of the Line Items or Custom Line Items in a custom workflow.
	State StateReference `json:"state"`
}

/**
*	Represents a snapshot of a Product Variant at the time it was imported with the Order. The Product Variant can be specified by providing a `productId` and `variant.id`, or by providing a `variant.sku`.
*
 */
type LineItemImportDraft struct {
	// Name of the Line Item.
	Name LocalizedString `json:"name"`
	// User-defined unique identifier of the Line Item.
	Key *string `json:"key,omitempty"`
	// The Product Variant to use as a [Line Item](ctp:api:type:LineItem).
	Variant ProductVariantImportDraft `json:"variant"`
	// `id` of the [Product](ctp:api:type:Product) the Product Variant belongs to.
	//
	// If provided, you must also set `variant.id`.
	ProductId *string `json:"productId,omitempty"`
	// The number of Product Variants in the LineItem. Can be a negative value.
	Quantity int `json:"quantity"`
	// The Line Item price for `quantity` = `1`. The amount can be negative.
	Price PriceDraft `json:"price"`
	// The tax rate used to calculate the `taxedPrice` of the Order.
	TaxRate *TaxRate `json:"taxRate,omitempty"`
	// The Channel used to [select a Price](ctp:api:type:LineItemPriceSelection).
	// This Channel must have the `ProductDistribution` role.
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
	// The Channel used to supply Line Items.
	// By providing supply Channel information, you can uniquely identify [Inventory entries](ctp:api:type:InventoryEntry) that should be reserved.
	// This Channel must have the `InventorySupply` role.
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	// Inventory mode specific to the LineItem, valid for the entire `quantity` of the LineItem.
	// Set only if Inventory mode should be different from the `inventoryMode` specified on the [OrderImportDraft](ctp:api:type:OrderImportDraft).
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	// Container for Line Item-specific addresses.
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	// States of the Line Item.
	State []ItemState `json:"state"`
	// Custom Fields of the LineItem.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
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

type Order struct {
	// Unique identifier of the Order.
	ID string `json:"id"`
	// Current version of the Order.
	Version int `json:"version"`
	// Date and time (UTC) the Order was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Order was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// User-defined identifier of the Order that is unique across a Project.
	OrderNumber *string `json:"orderNumber,omitempty"`
	// User-defined identifier of a purchase Order.
	//
	// It is typically set by the [Buyer](ctp:api:type:Buyer) and can be used with [Quotes](/quotes-overview) to track the purchase Order during the [quote and order flow](/../api/quotes-overview#intended-workflow).
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
	// `id` of the [Customer](ctp:api:type:Customer) that the Order belongs to.
	CustomerId *string `json:"customerId,omitempty"`
	// Email address of the Customer that the Order belongs to.
	CustomerEmail *string `json:"customerEmail,omitempty"`
	// [Reference](ctp:api:type:Reference) to the Customer Group of the Customer that the Order belongs to.
	// Used for [LineItem Price selection](ctp:api:type:LineItemPriceSelection).
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// [Anonymous session](ctp:api:type:AnonymousSession) associated with the Order.
	AnonymousId *string `json:"anonymousId,omitempty"`
	// [Reference](ctp:api:type:Reference) to a Business Unit the Order belongs to.
	BusinessUnit *BusinessUnitKeyReference `json:"businessUnit,omitempty"`
	// [Reference](ctp:api:type:Reference) to a Store the Order belongs to.
	Store *StoreKeyReference `json:"store,omitempty"`
	// [Line Items](ctp:api:type:LineItems) that are part of the Order.
	LineItems []LineItem `json:"lineItems"`
	// [Custom Line Items](ctp:api:type:CustomLineItems) that are part of the Order.
	CustomLineItems []CustomLineItem `json:"customLineItems"`
	// Sum of the `totalPrice` field of all [LineItems](ctp:api:type:LineItem) and [CustomLineItems](ctp:api:type:CustomLineItem), and if available, the `price` field of [ShippingInfo](ctp:api:type:ShippingInfo).
	// If a discount applies on `totalPrice`, this field holds the discounted value.
	//
	// Taxes are included if [TaxRate](ctp:api:type:TaxRate) `includedInPrice` is `true` for each price.
	TotalPrice TypedMoney `json:"totalPrice"`
	// - For `Platform` [TaxMode](ctp:api:type:TaxMode), it is automatically set when a [shipping address is set](ctp:api:type:OrderSetShippingAddressAction).
	// - For `External` [TaxMode](ctp:api:type:TaxMode), it is automatically set when `shippingAddress` and external Tax Rates for all Line Items, Custom Line Items, and Shipping Methods in the Cart are set.
	//
	// If a discount applies on `totalPrice`, this field holds the discounted values.
	TaxedPrice *TaxedPrice `json:"taxedPrice,omitempty"`
	// Sum of the `taxedPrice` field of [ShippingInfo](ctp:api:type:ShippingInfo) across all Shipping Methods.
	TaxedShippingPrice *TaxedPrice `json:"taxedShippingPrice,omitempty"`
	// Discounts that apply on the total price of the Order.
	DiscountOnTotalPrice *DiscountOnTotalPrice `json:"discountOnTotalPrice,omitempty"`
	// Indicates how Tax Rates are set.
	TaxMode *TaxMode `json:"taxMode,omitempty"`
	// Indicates how monetary values are rounded when calculating taxes for `taxedPrice`.
	TaxRoundingMode *RoundingMode `json:"taxRoundingMode,omitempty"`
	// Indicates how taxes are calculated when calculating taxes for `taxedPrice`.
	TaxCalculationMode *TaxCalculationMode `json:"taxCalculationMode,omitempty"`
	// Indicates how stock quantities are tracked for Line Items in the Order.
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	// Billing address associated with the Order.
	BillingAddress *Address `json:"billingAddress,omitempty"`
	// Shipping address associated with the Order.
	// Determines eligible [ShippingMethod](ctp:api:type:ShippingMethod) rates and Tax Rates of Line Items.
	ShippingAddress *Address `json:"shippingAddress,omitempty"`
	// Indicates whether there can be one or multiple Shipping Methods.
	ShippingMode ShippingMode `json:"shippingMode"`
	// `key` of the [ShippingMethod](ctp:api:type:ShippingMethod) for `Single` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
	// Shipping-related information for `Single` [ShippingMode](ctp:api:type:ShippingMode).
	// Automatically set when a [Shipping Method is set](ctp:api:type:StagedOrderSetShippingMethodAction).
	ShippingInfo *ShippingInfo `json:"shippingInfo,omitempty"`
	// Input used to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	// The data type of this field depends on the `shippingRateInputType.type` configured in the [Project](ctp:api:type:Project):
	//
	// - If `CartClassification`, it is [ClassificationShippingRateInput](ctp:api:type:ClassificationShippingRateInput).
	// - If `CartScore`, it is [ScoreShippingRateInput](ctp:api:type:ScoreShippingRateInput).
	// - If `CartValue`, it cannot be used.
	ShippingRateInput ShippingRateInput `json:"shippingRateInput,omitempty"`
	// Custom Fields of the Shipping Method for `Single` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingCustomFields *CustomFields `json:"shippingCustomFields,omitempty"`
	// Shipping-related information for `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	// Updated automatically each time a new [Shipping Method is added](ctp:api:type:CartAddShippingMethodAction).
	Shipping []Shipping `json:"shipping"`
	// Additional shipping addresses of the Order as specified by [LineItems](ctp:api:type:LineItem) using the `shippingDetails` field.
	// Eligible Shipping Methods or applicable Tax Rates are determined by the address in `shippingAddress`, and not `itemShippingAddresses`.
	ItemShippingAddresses []Address `json:"itemShippingAddresses"`
	// Discount Codes added to the Order.
	// An Order that has `directDiscounts` cannot have `discountCodes`.
	DiscountCodes []DiscountCodeInfo `json:"discountCodes"`
	// Direct Discounts added to the Order.
	// An Order that has `discountCodes` cannot have `directDiscounts`.
	DirectDiscounts []DirectDiscount `json:"directDiscounts"`
	// Automatically set when a Line Item with `GiftLineItem` [LineItemMode](ctp:api:type:LineItemMode) is [removed](ctp:api:type:StagedOrderRemoveLineItemAction) from the Order.
	RefusedGifts []CartDiscountReference `json:"refusedGifts"`
	// Payment information related to the Order.
	PaymentInfo *PaymentInfo `json:"paymentInfo,omitempty"`
	// Used for [LineItem Price selection](ctp:api:type:LineItemPriceSelection).
	Country *string `json:"country,omitempty"`
	// Languages of the Order.
	// Can only contain languages supported by the [Project](ctp:api:type:Project).
	Locale *string `json:"locale,omitempty"`
	// Indicates the origin of the Cart from which the Order was created.
	Origin CartOrigin `json:"origin"`
	// [Reference](ctp:api:type:Reference) to the Cart for an [Order created from Cart](ctp:api:endpoint:/{projectKey}/orders:POST).
	// The referenced Cart will have the `Ordered` [CartState](ctp:api:type:CartState).
	Cart *CartReference `json:"cart,omitempty"`
	// [Reference](ctp:api:type:Reference) to the Quote for an [Order created from Quote](ctp:api:endpoint:/{projectKey}/orders/quotes:POST).
	Quote *QuoteReference `json:"quote,omitempty"`
	// Current status of the Order.
	OrderState OrderState `json:"orderState"`
	// Shipment status of the Order.
	ShipmentState *ShipmentState `json:"shipmentState,omitempty"`
	// Payment status of the Order.
	PaymentState *PaymentState `json:"paymentState,omitempty"`
	// [State](ctp:api:type:State) of the Order.
	// This reference can point to a State in a custom workflow.
	State *StateReference `json:"state,omitempty"`
	// Contains synchronization activity information of the Order (like export or import).
	// Can only be set with [Update SyncInfo](ctp:api:type:OrderUpdateSyncInfoAction) update action.
	SyncInfo []SyncInfo `json:"syncInfo"`
	// Contains information regarding the returns associated with the Order.
	ReturnInfo []ReturnInfo `json:"returnInfo"`
	// Internal-only field.
	LastMessageSequenceNumber *int `json:"lastMessageSequenceNumber,omitempty"`
	// Custom Fields of the Order.
	Custom *CustomFields `json:"custom,omitempty"`
	// User-defined date and time (UTC) of the Order.
	// Present only on an Order created using [Order Import](ctp:api:endpoint:/{projectKey}/orders/import:POST).
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/general-concepts#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/general-concepts#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["itemShippingAddresses"] == nil {
		delete(raw, "itemShippingAddresses")
	}

	if raw["discountCodes"] == nil {
		delete(raw, "discountCodes")
	}

	if raw["directDiscounts"] == nil {
		delete(raw, "directDiscounts")
	}

	if raw["returnInfo"] == nil {
		delete(raw, "returnInfo")
	}

	return json.Marshal(raw)

}

type OrderFromCartDraft struct {
	// `id` of the [Cart](ctp:api:type:Cart) used to create the Order.
	ID *string `json:"id,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to the Cart from which the Order is created.
	//
	// This field is required, but is marked as optional for backwards compatibility reasons.
	Cart *CartResourceIdentifier `json:"cart,omitempty"`
	// `version` of the [Cart](ctp:api:type:Cart) from which the Order is created.
	Version int `json:"version"`
	// User-defined identifier for the Order that is unique across a Project.
	// Once set, the value cannot be changed.
	OrderNumber *string `json:"orderNumber,omitempty"`
	// User-defined identifier for a purchase Order.
	//
	// It is typically set by the [Buyer](ctp:api:type:Buyer) and can be used with [Quotes](/quotes-overview) to track the purchase Order during the [quote and order flow](/../api/quotes-overview#intended-workflow).
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
	// Payment status for the Order.
	PaymentState *PaymentState `json:"paymentState,omitempty"`
	// Shipment status for the Order.
	ShipmentState *ShipmentState `json:"shipmentState,omitempty"`
	// Current status for the Order.
	OrderState *OrderState `json:"orderState,omitempty"`
	// State for the Order in a custom workflow.
	State *StateResourceIdentifier `json:"state,omitempty"`
	// Custom Fields for the Order.
	// The Custom Fields' type must match the Custom Fields' type in the referenced [Cart](ctp:api:type:Cart).
	//
	// - If empty, the Custom Fields on the referenced [Cart](ctp:api:type:Cart) are added to the Order automatically.
	// - If specified, the Custom Fields are merged with the Custom Fields on the referenced [Cart](ctp:api:type:Cart) and added to the Order.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

type OrderFromQuoteDraft struct {
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to the Quote from which the Order is created.
	// If the referenced [Quote](ctp:api:type:Quote) has expired (`validTo` check) or its `quoteState` is `Accepted`, `Declined`, or `Withdrawn`, the Order creation will fail.
	Quote QuoteResourceIdentifier `json:"quote"`
	// `version` of the [Quote](ctp:api:type:Quote) from which the Order is created.
	Version int `json:"version"`
	// If `true`, the `quoteState` of the referenced [Quote](ctp:api:type:Quote) will be set to `Accepted`.
	QuoteStateToAccepted *bool `json:"quoteStateToAccepted,omitempty"`
	// User-defined identifier for the Order that is unique across a Project.
	// Once set, the value cannot be changed.
	OrderNumber *string `json:"orderNumber,omitempty"`
	// Payment status for the Order.
	PaymentState *PaymentState `json:"paymentState,omitempty"`
	// Shipment status for the Order.
	ShipmentState *ShipmentState `json:"shipmentState,omitempty"`
	// Current status for the Order.
	OrderState *OrderState `json:"orderState,omitempty"`
	// State of the Order in a custom workflow.
	State *StateResourceIdentifier `json:"state,omitempty"`
}

/**
*	A snapshot of an Order at the time it was imported.
*
 */
type OrderImportDraft struct {
	// User-defined identifier of the Order. Must be unique across a Project.
	// Once set, the value cannot be changed.
	OrderNumber *string `json:"orderNumber,omitempty"`
	// User-defined identifier for a purchase Order.
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
	// The `id` of the [Customer](ctp:api:type:Customer) the Order belongs to.
	CustomerId *string `json:"customerId,omitempty"`
	// The Email address of the Customer the Order belongs to. Can be used instead of `customerId` when no check against existing [Customers](ctp:api:type:Customer) is required.
	CustomerEmail *string `json:"customerEmail,omitempty"`
	// The Customer Group of the Customer the Order belongs to.
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to the Business Unit the Order should belong to.
	// When the `customerId` of the Order is also set, the [Customer](ctp:api:type:Customer) must be an [Associate](ctp:api:type:Associate) of the Business Unit.
	BusinessUnit *BusinessUnitResourceIdentifier `json:"businessUnit,omitempty"`
	// The Store the Order belongs to.
	// Used for [filtering](#filtering).
	//
	// If a [LineItemImportDraft](ctp:api:type:LineItemImportDraft) or a [CustomLineItemImportDraft](ctp:api:type:CustomLineItemImportDraft) specifies a `distributionChannel` or a `supplyChannel` that is not defined for the referenced Store, the Order Import gets rejected.
	// The same applies when the provided `country` is not defined for the referenced Store.
	Store *StoreResourceIdentifier `json:"store,omitempty"`
	// [Line Items](ctp:api:type:LineItems) to add to the Order.
	//
	// If not specified, `customLineItems` must not be empty.
	LineItems []LineItemImportDraft `json:"lineItems"`
	// [Custom Line Items](ctp:api:type:CustomLineItems) to add to the Cart.
	//
	// If not specified, `lineItems` must not be empty.
	CustomLineItems []CustomLineItemImportDraft `json:"customLineItems"`
	// The total Price of the Order. The amount can be negative.
	TotalPrice Money `json:"totalPrice"`
	// Include TaxedPrice information for the Order. If not included, and if you have Tax Rates set for Line Items and Custom Line Items, the Order total will not be recalculated.
	TaxedPrice *TaxedPriceDraft `json:"taxedPrice,omitempty"`
	// Determines how monetary values are rounded when calculating taxes for `taxedPrice`.
	TaxRoundingMode *RoundingMode `json:"taxRoundingMode,omitempty"`
	// Determines how taxes are calculated for `taxedPrice`.
	TaxCalculationMode *TaxCalculationMode `json:"taxCalculationMode,omitempty"`
	// Determines how stock quantities are tracked for Line Items in the Cart.
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	// Billing address associated with the Order.
	BillingAddress *BaseAddress `json:"billingAddress,omitempty"`
	// Shipping address associated with the Order.
	ShippingAddress *BaseAddress `json:"shippingAddress,omitempty"`
	// Addresses for Orders with multiple shipping addresses. Addresses must include a value for `key`.
	ItemShippingAddresses []BaseAddress `json:"itemShippingAddresses"`
	// Shipping-related information of the Order.
	ShippingInfo *ShippingInfoImportDraft `json:"shippingInfo,omitempty"`
	// Payment information associated with the Order.
	PaymentInfo *PaymentInfo `json:"paymentInfo,omitempty"`
	// Payment status of the Order.
	PaymentState *PaymentState `json:"paymentState,omitempty"`
	// Shipment status of the Order.
	ShipmentState *ShipmentState `json:"shipmentState,omitempty"`
	// Current status of the Order.
	OrderState *OrderState `json:"orderState,omitempty"`
	// State of the Order in a custom workflow.
	State *StateReference `json:"state,omitempty"`
	// Include a value to associate a country with the Order.
	Country *string `json:"country,omitempty"`
	// Indicates the origin of the Order.
	Origin *CartOrigin `json:"origin,omitempty"`
	// User-defined date and time for the Order. This value does not influence the `createdAt` or `lastModifiedAt` values of the Order created by the Order Import.
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	// Custom Fields for the Order.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
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

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [Order](ctp:api:type:Order).
*
 */
type OrderPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](ctp:api:type:QueryPredicate), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [Orders](ctp:api:type:Order) matching the query.
	Results []Order `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to an [Order](ctp:api:type:Order).
*
 */
type OrderReference struct {
	// Unique identifier of the referenced [Order](ctp:api:type:Order).
	ID string `json:"id"`
	// Contains the representation of the expanded Order. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for Orders.
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

type OrderSearchQuery map[string]interface{}
type OrderSearchCompoundExpression map[string]interface{}
type OrderSearchAndExpression struct {
	And []OrderSearchQuery `json:"and"`
}

type OrderSearchFilterExpression struct {
	Filter []OrderSearchQueryExpression `json:"filter"`
}

type OrderSearchNotExpression struct {
	Not []OrderSearchQuery `json:"not"`
}

type OrderSearchOrExpression struct {
	Or []OrderSearchQuery `json:"or"`
}

type OrderSearchQueryExpression map[string]interface{}
type OrderSearchDateRangeExpression struct {
	Range OrderSearchDateRangeValue `json:"range"`
}

type OrderSearchExactExpression struct {
	Exact OrderSearchAnyValue `json:"exact"`
}

type OrderSearchExistsExpression struct {
	Exists OrderSearchQueryExpressionValue `json:"exists"`
}

type OrderSearchFullTextExpression struct {
	FullText OrderSearchFullTextValue `json:"fullText"`
}

type OrderSearchLongRangeExpression struct {
	Range OrderSearchLongRangeValue `json:"range"`
}

type OrderSearchNumberRangeExpression struct {
	Range OrderSearchNumberRangeValue `json:"range"`
}

type OrderSearchPrefixExpression struct {
	Prefix OrderSearchStringValue `json:"prefix"`
}

type OrderSearchWildCardExpression struct {
	Wildcard OrderSearchStringValue `json:"wildcard"`
}

type OrderSearchRequest struct {
	// The Order search query.
	Query OrderSearchQuery `json:"query"`
	// Controls how results to your query are sorted. If not provided, the results are sorted by relevance in descending order.
	Sort []OrderSearchSorting `json:"sort"`
	// The maximum number of search results to be returned.
	Limit *int `json:"limit,omitempty"`
	// The number of search results to be skipped in the response for pagination.
	Offset *int `json:"offset,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSearchRequest) MarshalJSON() ([]byte, error) {
	type Alias OrderSearchRequest
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

	if raw["sort"] == nil {
		delete(raw, "sort")
	}

	return json.Marshal(raw)

}

type OrderSearchSorting struct {
	Field    string                      `json:"field"`
	Language *string                     `json:"language,omitempty"`
	Order    *OrderSearchSortOrder       `json:"order,omitempty"`
	Mode     *OrderSearchSortMode        `json:"mode,omitempty"`
	Filter   *OrderSearchQueryExpression `json:"filter,omitempty"`
}

/**
*	Indicates the state of the Order.
*
 */
type OrderState string

const (
	OrderStateOpen      OrderState = "Open"
	OrderStateConfirmed OrderState = "Confirmed"
	OrderStateComplete  OrderState = "Complete"
	OrderStateCancelled OrderState = "Cancelled"
)

type OrderUpdate struct {
	// Expected version of the Order on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the Order.
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
	case "setPurchaseOrderNumber":
		obj := OrderSetPurchaseOrderNumberAction{}
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

/**
*	Information regarding the appearance, content, and shipment of a Parcel.
*
 */
type Parcel struct {
	// Unique identifier of the Parcel.
	ID string `json:"id"`
	// User-defined unique identifier of the Parcel.
	Key *string `json:"key,omitempty"`
	// Date and time (UTC) the Parcel was created.
	CreatedAt time.Time `json:"createdAt"`
	// Information about the dimensions of the Parcel.
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	// Shipment tracking information of the Parcel.
	TrackingData *TrackingData `json:"trackingData,omitempty"`
	// Line Items or Custom Line Items delivered in this Parcel.
	Items []DeliveryItem `json:"items"`
	// Custom Fields of the Parcel.
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["items"] == nil {
		delete(raw, "items")
	}

	return json.Marshal(raw)

}

type ParcelDraft struct {
	// User-defined unique identifier of the Parcel.
	Key *string `json:"key,omitempty"`
	// Information about the dimensions for the Parcel.
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	// Shipment tracking information for the Parcel.
	TrackingData *TrackingData `json:"trackingData,omitempty"`
	// Line Items or Custom Line Items delivered in this Parcel.
	Items []DeliveryItem `json:"items"`
	// Custom Fields for the Parcel.
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["items"] == nil {
		delete(raw, "items")
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

type PaymentInfo struct {
	// [References](ctp:api:type:Reference) to the Payments associated with the Order.
	Payments []PaymentReference `json:"payments"`
}

/**
*	Indicates the payment status for the Order.
*
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
*	Contains the Product Variant to be used in the [LineItemImportDraft](ctp:api:type:LineItemImportDraft).
*
 */
type ProductVariantImportDraft struct {
	// The `id` of the [ProductVariant](ctp:api:type:ProductVariant). Required if you do not set a value for `sku`.
	// If set, you must specify a `productId` in the [LineItemImportDraft](ctp:api:type:LineItemImportDraft) also.
	ID *int `json:"id,omitempty"`
	// The `sku` of the [ProductVariant](ctp:api:type:ProductVariant). Required if you do not set a value for `id`.
	Sku *string `json:"sku,omitempty"`
	// The [Prices](ctp:api:type:Price) of the Product Variant if you want to override the `prices` property in the referenced [ProductVariant](ctp:api:type:ProductVariant).
	// If not set, the `prices` from the referenced [ProductVariant](ctp:api:type:ProductVariant) are used in the resulting Order.
	// If set, each Price must have its unique price scope (same `value.currencyCode`, `country`, `customerGroup`, `channel`, `validFrom` and `validUntil`).
	Prices []PriceDraft `json:"prices"`
	// The [Attributes](ctp:api:type:Attribute) of the Product Variant if you want to override the `attributes` property in the referenced [ProductVariant](ctp:api:type:ProductVariant).
	// If not set, the `attributes` from the referenced [ProductVariant](ctp:api:type:ProductVariant) are copied to the resulting Order.
	Attributes []Attribute `json:"attributes"`
	// The [Images](ctp:api:type:Image) of the Product Variant if you want to override the `images` property in the referenced [ProductVariant](ctp:api:type:ProductVariant).
	// If not set, the `images` from the referenced [ProductVariant](ctp:api:type:ProductVariant) are copied to the resulting Order.
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
*	Stores information about returns connected to an Order.
*
 */
type ReturnInfo struct {
	// Information on the Line Items or Custom Line Items returned.
	Items []ReturnItem `json:"items"`
	// User-defined identifier to track the return.
	ReturnTrackingId *string `json:"returnTrackingId,omitempty"`
	// Date and time (UTC) the return is initiated.
	ReturnDate *time.Time `json:"returnDate,omitempty"`
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
	// Information on the Line Items or Custom Line Items returned.
	Items []ReturnItemDraft `json:"items"`
	// User-defined identifier for tracking the return.
	ReturnTrackingId *string `json:"returnTrackingId,omitempty"`
	// Date and time (UTC) the return is initiated.
	ReturnDate *time.Time `json:"returnDate,omitempty"`
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
	// Unique identifier of the Return Item.
	ID string `json:"id"`
	// User-defined unique identifier of the CustomLineItemReturnItem.
	Key *string `json:"key,omitempty"`
	// Number of Custom Line Items returned.
	Quantity int `json:"quantity"`
	// User-defined description for the return.
	Comment *string `json:"comment,omitempty"`
	// Shipment status of the Return Item.
	ShipmentState ReturnShipmentState `json:"shipmentState"`
	// Payment status of the Return Item:
	//
	// - `NonRefundable`, for items in the `Advised` [ReturnShipmentState](ctp:api:type:ReturnShipmentState)
	// - `Initial`, for items in the `Returned` [ReturnShipmentState](ctp:api:type:ReturnShipmentState)
	PaymentState ReturnPaymentState `json:"paymentState"`
	// Custom Fields of the Return Item.
	Custom *CustomFields `json:"custom,omitempty"`
	// Date and time (UTC) the Return Item was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Date and time (UTC) the Return Item was intitially created.
	CreatedAt time.Time `json:"createdAt"`
	// `id` of the returned [CustomLineItem](ctp:api:type:CustomLineItem).
	CustomLineItemId string `json:"customLineItemId"`
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
	// Unique identifier of the Return Item.
	ID string `json:"id"`
	// User-defined unique identifier of the LineItemReturnItem.
	Key *string `json:"key,omitempty"`
	// Number of Line Items returned.
	Quantity int `json:"quantity"`
	// User-defined description for the return.
	Comment *string `json:"comment,omitempty"`
	// Shipment status of the Return Item.
	ShipmentState ReturnShipmentState `json:"shipmentState"`
	// Payment status of the Return Item:
	//
	// - `NonRefundable`, for items in the `Advised` [ReturnShipmentState](ctp:api:type:ReturnShipmentState)
	// - `Initial`, for items in the `Returned` [ReturnShipmentState](ctp:api:type:ReturnShipmentState)
	PaymentState ReturnPaymentState `json:"paymentState"`
	// Custom Fields of the Return Item.
	Custom *CustomFields `json:"custom,omitempty"`
	// Date and time (UTC) the Return Item was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Date and time (UTC) the Return Item was intitially created.
	CreatedAt time.Time `json:"createdAt"`
	// `id` of the returned [LineItem](ctp:api:type:LineItem).
	LineItemId string `json:"lineItemId"`
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
	// User-defined unique identifier of the Return Item.
	Key *string `json:"key,omitempty"`
	// Number of Line Items or Custom Line Items to return.
	Quantity int `json:"quantity"`
	// `id` of the [LineItem](ctp:api:type:LineItem) to return.
	//
	//  Required if Line Items are returned, to create a [LineItemReturnItem](ctp:api:type:LineItemReturnItem).
	LineItemId *string `json:"lineItemId,omitempty"`
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to return.
	//
	//  Required if Custom Line Items are returned, to create a [CustomLineItemReturnItem](ctp:api:type:CustomLineItemReturnItem).
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// User-defined description for the return.
	Comment *string `json:"comment,omitempty"`
	// Shipment status of the item to be returned.
	// Can either be `Advised` or `Returned` only.
	ShipmentState ReturnShipmentState `json:"shipmentState"`
	// Custom Fields for the Return Item.
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

/**
*	Indicates the shipment status of the Order.
*
 */
type ShipmentState string

const (
	ShipmentStateShipped   ShipmentState = "Shipped"
	ShipmentStateDelivered ShipmentState = "Delivered"
	ShipmentStateReady     ShipmentState = "Ready"
	ShipmentStatePending   ShipmentState = "Pending"
	ShipmentStateDelayed   ShipmentState = "Delayed"
	ShipmentStatePartial   ShipmentState = "Partial"
	ShipmentStateBackorder ShipmentState = "Backorder"
)

/**
*	Becomes the `shippingInfo` of the imported Order.
*
 */
type ShippingInfoImportDraft struct {
	// Name of the Shipping Method.
	ShippingMethodName string `json:"shippingMethodName"`
	// The base price for the Shipping Method.
	Price Money `json:"price"`
	// Shipping rate information for the Order.
	ShippingRate ShippingRateDraft `json:"shippingRate"`
	// Include a Tax Rate for the Shipping Method.
	TaxRate *TaxRate `json:"taxRate,omitempty"`
	// Include a value to associate a Tax Category with the shipping information.
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// Include a value to associate a Shipping Method with the Order.
	ShippingMethod *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	// Information on how items are to be delivered to customers.
	Deliveries []DeliveryDraft `json:"deliveries"`
	// Discounted Price of the Shipping Method.
	DiscountedPrice *DiscountedLineItemPriceDraft `json:"discountedPrice,omitempty"`
	// Indicates if the [ShippingMethod](ctp:api:type:ShippingMethod) referenced is allowed for the Order or not.
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["deliveries"] == nil {
		delete(raw, "deliveries")
	}

	return json.Marshal(raw)

}

/**
*	Contains synchronization activity information of the Order (like export or import).
*
 */
type SyncInfo struct {
	// Connection to a synchronization destination.
	Channel ChannelReference `json:"channel"`
	// Identifier of an external order instance, file, or other resource.
	ExternalId *string `json:"externalId,omitempty"`
	// Date and time (UTC) the information was synced.
	SyncedAt time.Time `json:"syncedAt"`
}

type TaxedItemPriceDraft struct {
	// Draft type that stores amounts only in cent precision for the specified currency.
	TotalNet Money `json:"totalNet"`
	// Draft type that stores amounts only in cent precision for the specified currency.
	TotalGross Money `json:"totalGross"`
}

/**
*	Information that helps track a Parcel.
*
 */
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

/**
*	A [Delivery](ctp:api:type:Delivery) can only be added to an [Order](ctp:api:type:Order) if
*	its `shippingInfo` (for `shippingMode` = `Single`), or its `shipping` (for `shippingMode` = `Multiple`) exists.
*
*	Produces the [Delivery Added](ctp:api:type:DeliveryAddedMessage) Message.
*
 */
type OrderAddDeliveryAction struct {
	// `key` of an existing [Delivery](ctp:api:type:Delivery).
	DeliveryKey *string `json:"deliveryKey,omitempty"`
	// `key` of the [ShippingMethod](ctp:api:type:ShippingMethod), required for `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
	// Line Items or Custom Line Items to be included in the Delivery.
	Items []DeliveryItem `json:"items"`
	// Address the `parcels` should be delivered to.
	Address *BaseAddress `json:"address,omitempty"`
	// Parcels of the Delivery.
	//
	// If provided, this update action produces the [Parcel Added To Delivery](ctp:api:type:ParcelAddedToDeliveryMessage) Message.
	Parcels []ParcelDraft `json:"parcels"`
	// Custom Fields for the Delivery.
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["items"] == nil {
		delete(raw, "items")
	}

	if raw["parcels"] == nil {
		delete(raw, "parcels")
	}

	return json.Marshal(raw)

}

/**
*	Adds an address to an Order when shipping to multiple addresses is desired.
*
 */
type OrderAddItemShippingAddressAction struct {
	// Address to append to `itemShippingAddresses`.
	// The new Address must have a `key` that is unique across this Order.
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

/**
*	To add a Parcel, at least one [Delivery](ctp:api:type:Delivery) must exist.
*
*	Produces the [Parcel Added To Delivery](ctp:api:type:ParcelAddedToDeliveryMessage) Message.
*
 */
type OrderAddParcelToDeliveryAction struct {
	// `id` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryId *string `json:"deliveryId,omitempty"`
	// `key` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryKey *string `json:"deliveryKey,omitempty"`
	// `key` of an existing [Parcel](ctp:api:type:Parcel).
	ParcelKey *string `json:"parcelKey,omitempty"`
	// Value to set.
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	// Value to set.
	TrackingData *TrackingData `json:"trackingData,omitempty"`
	// Value to set.
	Items []DeliveryItem `json:"items"`
	// Custom Fields for the Parcel.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["items"] == nil {
		delete(raw, "items")
	}

	return json.Marshal(raw)

}

type OrderAddPaymentAction struct {
	// Payment to add to the [PaymentInfo](ctp:api:type:PaymentInfo).
	// Must not be assigned to another Order or active Cart already.
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

/**
*	Produces the [Return Info Added](ctp:api:type:ReturnInfoAddedMessage) Message.
*
 */
type OrderAddReturnInfoAction struct {
	// Value to set.
	ReturnTrackingId *string `json:"returnTrackingId,omitempty"`
	// Items to be returned.
	// Must not be empty.
	Items []ReturnItemDraft `json:"items"`
	// Value to set.
	// If not set, it defaults to the current date and time.
	ReturnDate *time.Time `json:"returnDate,omitempty"`
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

/**
*	Produces the [Order State Changed](ctp:api:type:OrderStateChangedMessage) Message.
*
 */
type OrderChangeOrderStateAction struct {
	// New status of the Order.
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

/**
*	Produces the [Order Payment State Changed](ctp:api:type:OrderPaymentStateChangedMessage) Message.
*
 */
type OrderChangePaymentStateAction struct {
	// New payment status of the Order.
	PaymentState PaymentState `json:"paymentState"`
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

/**
*	Produces the [Order Shipment State Changed](ctp:api:type:OrderShipmentStateChangedMessage) Message.
*
 */
type OrderChangeShipmentStateAction struct {
	// New shipment status of the Order.
	ShipmentState ShipmentState `json:"shipmentState"`
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

/**
*	The import of States does not follow any predefined rules and should be only used if no transitions are defined.
*	The `quantity` in the [ItemStates](ctp:api:type:ItemState) must match the sum of all Custom Line Item states' quantities.
*
 */
type OrderImportCustomLineItemStateAction struct {
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
	// New status of the Custom Line Items.
	State []ItemState `json:"state"`
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

/**
*	The import of States does not follow any predefined rules and should be only used if no transitions are defined.
*	The `quantity` in the [ItemStates](ctp:api:type:ItemState) must match the sum of all Line Items states' quantities.
*
 */
type OrderImportLineItemStateAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// New status of the Line Items.
	State []ItemState `json:"state"`
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

/**
*	Produces the [DeliveryRemoved](ctp:api:type:DeliveryRemovedMessage) Message.
*
 */
type OrderRemoveDeliveryAction struct {
	// `id` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryId *string `json:"deliveryId,omitempty"`
	// `key` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryKey *string `json:"deliveryKey,omitempty"`
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

/**
*	An address can only be removed if it is not referenced in any [ItemShippingTarget](ctp:api:type:ItemShippingTarget) of the Cart.
*	In such case, change the Line Item shipping address to a different `addressKey` first using the [Set LineItemShippingDetails](ctp:api:type:OrderSetLineItemShippingDetailsAction) update action, before you remove the obsolete address.
*
 */
type OrderRemoveItemShippingAddressAction struct {
	// `key` of the Address to remove from `itemShippingAddresses`.
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

/**
*	Produces the [ParcelRemovedFromDelivery](ctp:api:type:ParcelRemovedFromDeliveryMessage) Message.
*
 */
type OrderRemoveParcelFromDeliveryAction struct {
	// `id` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelId *string `json:"parcelId,omitempty"`
	// `key` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelKey *string `json:"parcelKey,omitempty"`
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
	// Payment to remove from the [PaymentInfo](ctp:api:type:PaymentInfo).
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

/**
*	This action updates the `billingAddress` on the Order, but it does not change the billing address on the referenced [Cart](ctp:api:type:Cart) from which the Order is created.
*
*	Produces the [Order Billing Address Set](ctp:api:type:OrderBillingAddressSetMessage) Message.
*
 */
type OrderSetBillingAddressAction struct {
	// Value to set.
	// If empty, any existing value is removed.
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
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
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
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
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
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
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
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
	// Defines the [Type](ctp:api:type:Type) that extends the Custom Line Item with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Custom Line Item.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Custom Line Item.
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
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
	// Value to set.
	// If empty, any existing value is removed.
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
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

/**
*	This action updates the `customerEmail` on the Order, but it does not change the Customer email on the [Cart](ctp:api:type:Cart) the Order has been created from.
*
*	Produces the [Order Customer Email Set](ctp:api:type:OrderCustomerEmailSetMessage) Message.
*
 */
type OrderSetCustomerEmailAction struct {
	// Value to set.
	// If empty, any existing value is removed.
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

/**
*	Setting the Order's `customerId` does not recalculate prices or discounts on the Order.
*	If the Customer belongs to a Customer Group, `customerGroup` on the [Order](ctp:api:type:Order) is updated automatically.
*
*	Produces the [OrderCustomerSet](ctp:api:type:OrderCustomerSetMessage) Message.
*
 */
type OrderSetCustomerIdAction struct {
	// `id` of an existing [Customer](ctp:api:type:Customer).
	// If empty, any existing value is removed.
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

/**
*	Produces the [DeliveryAddressSet](ctp:api:type:DeliveryAddressSetMessage) Message.
*
 */
type OrderSetDeliveryAddressAction struct {
	// `id` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryId *string `json:"deliveryId,omitempty"`
	// `key` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryKey *string `json:"deliveryKey,omitempty"`
	// Value to set.
	// If empty, any existing value is removed.
	Address *BaseAddress `json:"address,omitempty"`
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
	// `id` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryId *string `json:"deliveryId,omitempty"`
	// `key` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryKey *string `json:"deliveryKey,omitempty"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
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
	// `id` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryId *string `json:"deliveryId,omitempty"`
	// `key` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryKey *string `json:"deliveryKey,omitempty"`
	// Defines the [Type](ctp:api:type:Type) that extends the [Delivery](ctp:api:type:Delivery) `address` with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the [Delivery](ctp:api:type:Delivery) `address`.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the [Delivery](ctp:api:type:Delivery) `address`.
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
	// `id` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryId *string `json:"deliveryId,omitempty"`
	// `key` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryKey *string `json:"deliveryKey,omitempty"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
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
	// `id` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryId *string `json:"deliveryId,omitempty"`
	// `key` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryKey *string `json:"deliveryKey,omitempty"`
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

/**
*	Produces the [Delivery Items Updated](ctp:api:type:DeliveryItemsUpdatedMessage) Message.
*
 */
type OrderSetDeliveryItemsAction struct {
	// `id` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryId *string `json:"deliveryId,omitempty"`
	// `key` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryKey *string `json:"deliveryKey,omitempty"`
	// Value to set.
	// If empty, any existing value is removed.
	Items []DeliveryItem `json:"items"`
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
	// `key` of the [Address](ctp:api:type:Address) in `itemShippingAddresses`.
	AddressKey string `json:"addressKey"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
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
	// `key` of the [Address](ctp:api:type:Address) in `itemShippingAddresses`.
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
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
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
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Defines the [Type](ctp:api:type:Type) that extends the Line Item with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Line Item.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Line Item.
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
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Value to set.
	// If empty, the existing value is removed.
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
	// Value to set.
	// Must be one of the [Project](ctp:api:type:Project)'s languages.
	// If empty, any existing value is removed.
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
	// Value to set.
	// Must be unique across a Project.
	// Once set, the value cannot be changed.
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
	// `id` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelId *string `json:"parcelId,omitempty"`
	// `key` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelKey *string `json:"parcelKey,omitempty"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
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
	// `id` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelId *string `json:"parcelId,omitempty"`
	// `key` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelKey *string `json:"parcelKey,omitempty"`
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

/**
*	Produces the [ParcelItemsUpdated](ctp:api:type:ParcelItemsUpdatedMessage) Message.
*
 */
type OrderSetParcelItemsAction struct {
	// `id` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelId *string `json:"parcelId,omitempty"`
	// `key` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelKey *string `json:"parcelKey,omitempty"`
	// Value to set.
	// If empty, any existing value is removed.
	Items []DeliveryItem `json:"items"`
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

/**
*	Produces the [ParcelMeasurementsUpdated](ctp:api:type:ParcelMeasurementsUpdatedMessage) Message.
*
 */
type OrderSetParcelMeasurementsAction struct {
	// `id` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelId *string `json:"parcelId,omitempty"`
	// `key` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelKey *string `json:"parcelKey,omitempty"`
	// Value to set.
	// If empty, any existing value is removed.
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

/**
*	Produces the [ParcelTrackingDataUpdated](ctp:api:type:ParcelTrackingDataUpdatedMessage) Message.
*
 */
type OrderSetParcelTrackingDataAction struct {
	// `id` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelId *string `json:"parcelId,omitempty"`
	// `key` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelKey *string `json:"parcelKey,omitempty"`
	// Value to set.
	// If empty, any existing value is removed.
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

/**
*	Produces the [PurchaseOrderNumberSet](ctp:api:type:OrderPurchaseOrderNumberSetMessage) Message.
*
 */
type OrderSetPurchaseOrderNumberAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderSetPurchaseOrderNumberAction) MarshalJSON() ([]byte, error) {
	type Alias OrderSetPurchaseOrderNumberAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPurchaseOrderNumber", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [Return Info Set](ctp:api:type:ReturnInfoSetMessage) Message.
*
 */
type OrderSetReturnInfoAction struct {
	// Value to set.
	// If empty, any existing value is removed.
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["items"] == nil {
		delete(raw, "items")
	}

	return json.Marshal(raw)

}

type OrderSetReturnItemCustomFieldAction struct {
	// `id` of the [ReturnItem](ctp:api:type:ReturnItem) to update. Either `returnItemId` or `returnItemKey` is required.
	ReturnItemId *string `json:"returnItemId,omitempty"`
	// `key` of the [ReturnItem](ctp:api:type:ReturnItem) to update. Either `returnItemId` or `returnItemKey` is required.
	ReturnItemKey *string `json:"returnItemKey,omitempty"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
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
	// `id` of the [ReturnItem](ctp:api:type:ReturnItem) to update. Either `returnItemId` or `returnItemKey` is required.
	ReturnItemId *string `json:"returnItemId,omitempty"`
	// `key` of the [ReturnItem](ctp:api:type:ReturnItem) to update. Either `returnItemId` or `returnItemKey` is required.
	ReturnItemKey *string `json:"returnItemKey,omitempty"`
	// Defines the [Type](ctp:api:type:Type) that extends the Return Item with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Return Item.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Return Item.
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

/**
*	To set a [ReturnPaymentState](ctp:api:type:ReturnPaymentState), the [Order](ctp:api:type:Order) `returnInfo` must have at least one [ReturnItem](ctp:api:type:ReturnItem).
*
 */
type OrderSetReturnPaymentStateAction struct {
	// `id` of the [ReturnItem](ctp:api:type:ReturnItem) to update. Either `returnItemId` or `returnItemKey` is required.
	ReturnItemId *string `json:"returnItemId,omitempty"`
	// `key` of the [ReturnItem](ctp:api:type:ReturnItem) to update. Either `returnItemId` or `returnItemKey` is required.
	ReturnItemKey *string `json:"returnItemKey,omitempty"`
	// New Payment status of the [ReturnItem](ctp:api:type:ReturnItem).
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

/**
*	To set a `ReturnShipmentState`, the [Order](ctp:api:type:Order) `returnInfo` must have at least one [ReturnItem](ctp:api:type:ReturnItem).
*
*	Produces the [Order Return Shipment State Changed](ctp:api:type:OrderReturnShipmentStateChangedMessage) Message.
*
 */
type OrderSetReturnShipmentStateAction struct {
	// `id` of the [ReturnItem](ctp:api:type:ReturnItem) to update. Either `returnItemId` or `returnItemKey` is required.
	ReturnItemId *string `json:"returnItemId,omitempty"`
	// `key` of the [ReturnItem](ctp:api:type:ReturnItem) to update. Either `returnItemId` or `returnItemKey` is required.
	ReturnItemKey *string `json:"returnItemKey,omitempty"`
	// New shipment state of the [ReturnItem](ctp:api:type:ReturnItem).
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

/**
*	This action updates the `shippingAddress` on the Order, but it does not change the shipping address on the referenced [Cart](ctp:api:type:Cart) from which the Order is created.
*	Also, it does not recalculate the Cart as taxes may not fit to the shipping address anymore.
*
*	Produces the [Order Shipping Address Set](ctp:api:type:OrderShippingAddressSetMessage) Message.
*
 */
type OrderSetShippingAddressAction struct {
	// Value to set.
	// If empty, any existing value is removed.
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
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
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

/**
*	Sets the [Store](ctp:api:type:Store) the Order is assigned to.
*	It should be used to migrate Orders to a new Store.
*	No validations are performed (such as that the Customer is allowed to create Orders in the Store).
*
*	Produces the [Order Store Set](ctp:api:type:OrderStoreSetMessage) Message.
*	Returns a `400` error if `store` references the same Store the Order is currently assigned to, including if you try to remove the value when no Store is currently assigned.
*
 */
type OrderSetStoreAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	//
	// If `store` references the same Store the Order is currently assigned to or if you try to remove the value when no Store is currently assigned, a `400` error is returned.
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

/**
*	Produces the [Custom Line Item State Transition](ctp:api:type:CustomLineItemStateTransitionMessage) Message.
*
 */
type OrderTransitionCustomLineItemStateAction struct {
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
	// Number of Custom Line Items that should transition [State](ctp:api:type:State).
	Quantity int `json:"quantity"`
	// [State](ctp:api:type:State) the Custom Line Item should transition from.
	FromState StateResourceIdentifier `json:"fromState"`
	// [State](ctp:api:type:State) the Custom Line Item should transition to.
	ToState StateResourceIdentifier `json:"toState"`
	// Date and time (UTC) to perform the [State](ctp:api:type:State) transition.
	ActualTransitionDate *time.Time `json:"actualTransitionDate,omitempty"`
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

/**
*	Produces the [Line Item State Transition](ctp:api:type:LineItemStateTransitionMessage) Message.
*
 */
type OrderTransitionLineItemStateAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Number of Line Items that should transition [State](ctp:api:type:State).
	Quantity int `json:"quantity"`
	// [State](ctp:api:type:State) the Line Item should transition from.
	FromState StateResourceIdentifier `json:"fromState"`
	// [State](ctp:api:type:State) the Line Item should transition to.
	ToState StateResourceIdentifier `json:"toState"`
	// Date and time (UTC) to perform the [State](ctp:api:type:State) transition.
	ActualTransitionDate *time.Time `json:"actualTransitionDate,omitempty"`
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

/**
*	If the existing [State](ctp:api:type:State) has set `transitions`, there must be a direct transition to the new State.
*	If `transitions` is not set, no validation is performed.
*
*	This update action produces the [Order State Transition](ctp:api:type:OrderStateTransitionMessage) Message.
*
 */
type OrderTransitionStateAction struct {
	// Value to set.
	// If there is no State yet, the new State must be an initial State.
	State StateResourceIdentifier `json:"state"`
	// Set to `true` to turn off validation.
	Force *bool `json:"force,omitempty"`
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

/**
*	Updates an address in `itemShippingAddresses` by keeping the Address `key`.
*
 */
type OrderUpdateItemShippingAddressAction struct {
	// The new Address with the same `key` as the Address it will replace.
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
	// Set this to identify an external order instance, file, or other resource.
	ExternalId *string `json:"externalId,omitempty"`
	// The synchronization destination to set. Must not be empty.
	// The referenced Channel must have the [Channel Role](ctp:api:type:ChannelRoleEnum) `OrderExport` or `OrderImport`.
	// Otherwise this update action returns an [InvalidInput](ctp:api:type:InvalidInputError) error.
	Channel ChannelResourceIdentifier `json:"channel"`
	// If not set, it defaults to the current date and time.
	SyncedAt *time.Time `json:"syncedAt,omitempty"`
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
