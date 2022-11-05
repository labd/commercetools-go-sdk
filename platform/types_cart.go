package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type Cart struct {
	// Unique identifier of the Cart.
	ID string `json:"id"`
	// The current version of the cart.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// User-defined unique identifier of the Cart.
	Key *string `json:"key,omitempty"`
	// Present on resources updated after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy     *CreatedBy `json:"createdBy,omitempty"`
	CustomerId    *string    `json:"customerId,omitempty"`
	CustomerEmail *string    `json:"customerEmail,omitempty"`
	// Identifies carts and orders belonging to an anonymous session (the customer has not signed up/in yet).
	AnonymousId *string `json:"anonymousId,omitempty"`
	// The Business Unit the Cart belongs to.
	BusinessUnit    *BusinessUnitKeyReference `json:"businessUnit,omitempty"`
	Store           *StoreKeyReference        `json:"store,omitempty"`
	LineItems       []LineItem                `json:"lineItems"`
	CustomLineItems []CustomLineItem          `json:"customLineItems"`
	// The sum of all `totalPrice` fields of the `lineItems` and `customLineItems`, as well as the `price` field of `shippingInfo` (if it exists).
	// `totalPrice` may or may not include the taxes: it depends on the taxRate.includedInPrice property of each price.
	TotalPrice TypedMoney `json:"totalPrice"`
	// Not set until the shipping address is set.
	// Will be set automatically in the `Platform` TaxMode.
	// For the `External` tax mode it will be set  as soon as the external tax rates for all line items, custom line items, and shipping in the cart are set.
	TaxedPrice *TaxedPrice `json:"taxedPrice,omitempty"`
	// Sum of `taxedPrice` of [ShippingInfo](ctp:api:type:ShippingInfo) across all Shipping Methods.
	// For `Platform` [TaxMode](ctp:api:type:TaxMode), it is set automatically only if [shipping address is set](ctp:api:type:CartSetShippingAddressAction) or [Shipping Method is added](ctp:api:type:CartAddShippingMethodAction) to the Cart.
	TaxedShippingPrice *TaxedPrice `json:"taxedShippingPrice,omitempty"`
	CartState          CartState   `json:"cartState"`
	// The shipping address is used to determine the eligible shipping methods and rates as well as the tax rate of the line items.
	ShippingAddress *Address `json:"shippingAddress,omitempty"`
	BillingAddress  *Address `json:"billingAddress,omitempty"`
	// Indicates whether one or multiple Shipping Methods are added to the Cart.
	ShippingMode ShippingMode `json:"shippingMode"`
	// Holds all shipping-related information per Shipping Method of a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	//
	// It is automatically updated after the [Shipping Method is added](ctp:api:type:CartAddShippingMethodAction).
	Shipping      []Shipping     `json:"shipping"`
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	TaxMode       TaxMode        `json:"taxMode"`
	// When calculating taxes for `taxedPrice`, the selected mode is used for rounding.
	TaxRoundingMode RoundingMode `json:"taxRoundingMode"`
	// When calculating taxes for `taxedPrice`, the selected mode is used for calculating the price with `LineItemLevel` (horizontally) or `UnitPriceLevel` (vertically) calculation mode.
	TaxCalculationMode TaxCalculationMode `json:"taxCalculationMode"`
	// Set automatically when the customer is set and the customer is a member of a customer group.
	// Used for product variant
	// price selection.
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	// Used for product variant price selection.
	Country *string `json:"country,omitempty"`
	// Shipping-related information of a Cart with `Single` [ShippingMode](ctp:api:type:ShippingMode).
	// Set automatically once the ShippingMethod is set.
	ShippingInfo    *ShippingInfo      `json:"shippingInfo,omitempty"`
	DiscountCodes   []DiscountCodeInfo `json:"discountCodes"`
	DirectDiscounts []DirectDiscount   `json:"directDiscounts"`
	Custom          *CustomFields      `json:"custom,omitempty"`
	PaymentInfo     *PaymentInfo       `json:"paymentInfo,omitempty"`
	Locale          *string            `json:"locale,omitempty"`
	// The cart will be deleted automatically if it hasn't been modified for the specified amount of days and it is in the `Active` CartState.
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
	// Automatically filled when a line item with LineItemMode `GiftLineItem` is removed from the cart.
	RefusedGifts []CartDiscountReference `json:"refusedGifts"`
	// The origin field indicates how this cart was created.
	// The value `Customer` indicates, that the cart was created by the customer.
	Origin CartOrigin `json:"origin"`
	// The shippingRateInput is used as an input to select a ShippingRatePriceTier.
	ShippingRateInput ShippingRateInput `json:"shippingRateInput,omitempty"`
	// Contains addresses for carts with multiple shipping addresses.
	// Line items reference these addresses under their `shippingDetails`.
	// The addresses captured here are not used to determine eligible shipping methods or the applicable tax rate.
	// Only the cart's `shippingAddress` is used for this.
	ItemShippingAddresses []Address `json:"itemShippingAddresses"`
	// The sum off all the [Line Items](ctp:api:type:LineItem) quantities. Does not take [Custom Line Items](ctp:api:type:CustomLineItem) into consideration.
	TotalLineItemQuantity *int `json:"totalLineItemQuantity,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Cart) UnmarshalJSON(data []byte) error {
	type Alias Cart
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
func (obj Cart) MarshalJSON() ([]byte, error) {
	type Alias Cart
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

	if raw["discountCodes"] == nil {
		delete(raw, "discountCodes")
	}

	if raw["directDiscounts"] == nil {
		delete(raw, "directDiscounts")
	}

	if raw["itemShippingAddresses"] == nil {
		delete(raw, "itemShippingAddresses")
	}

	return json.Marshal(raw)

}

type CartDraft struct {
	// A three-digit currency code as per [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Currency string `json:"currency"`
	// User-defined unique identifier for the Cart.
	Key *string `json:"key,omitempty"`
	// Id of an existing Customer.
	CustomerId    *string `json:"customerId,omitempty"`
	CustomerEmail *string `json:"customerEmail,omitempty"`
	// Will be set automatically when the `customerId` is set and the customer is a member of a customer group.
	// Can be set explicitly when no `customerId` is present.
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	// Assigns the new cart to an anonymous session (the customer has not signed up/in yet).
	AnonymousId *string `json:"anonymousId,omitempty"`
	// The Business Unit the Cart belongs to.
	BusinessUnit *BusinessUnitResourceIdentifier `json:"businessUnit,omitempty"`
	// Assigns the new cart to the store.
	// The store assignment can not be modified.
	Store *StoreResourceIdentifier `json:"store,omitempty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country *string `json:"country,omitempty"`
	// Default inventory mode is `None`.
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	// The default tax mode is `Platform`.
	TaxMode *TaxMode `json:"taxMode,omitempty"`
	// The default tax rounding mode is `HalfEven`.
	TaxRoundingMode *RoundingMode `json:"taxRoundingMode,omitempty"`
	// The default tax calculation mode is `LineItemLevel`.
	TaxCalculationMode *TaxCalculationMode   `json:"taxCalculationMode,omitempty"`
	LineItems          []LineItemDraft       `json:"lineItems"`
	CustomLineItems    []CustomLineItemDraft `json:"customLineItems"`
	// The shipping address is used to determine the eligible shipping methods and rates as well as the tax rate of the line items.
	ShippingAddress *BaseAddress                      `json:"shippingAddress,omitempty"`
	BillingAddress  *BaseAddress                      `json:"billingAddress,omitempty"`
	ShippingMethod  *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	// An external tax rate can be set for the `shippingMethod` if the cart has the `External` TaxMode.
	ExternalTaxRateForShippingMethod *ExternalTaxRateDraft `json:"externalTaxRateForShippingMethod,omitempty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Must be one of the languages supported for this project
	Locale *string `json:"locale,omitempty"`
	// The cart will be deleted automatically if it hasn't been modified for the specified amount of days and it is in the `Active` CartState.
	// If a ChangeSubscription for carts exists, a `ResourceDeleted` notification will be sent.
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
	// The default origin is `Customer`.
	Origin *CartOrigin `json:"origin,omitempty"`
	// - If `Single`, only a single Shipping Method can be added to the Cart.
	// - If `Multiple`, multiple Shipping Methods can be added to the Cart.
	ShippingMode *ShippingMode `json:"shippingMode,omitempty"`
	// Custom Shipping Methods for a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	CustomShipping []CustomShippingDraft `json:"customShipping"`
	// Shipping Methods for a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	Shipping []ShippingDraft `json:"shipping"`
	// The shippingRateInput is used as an input to select a ShippingRatePriceTier.
	// Based on the definition of ShippingRateInputType.
	// If CartClassification is defined, it must be ClassificationShippingRateInput.
	// If CartScore is defined, it must be ScoreShippingRateInput.
	// Otherwise it can not bet set.
	ShippingRateInput ShippingRateInputDraft `json:"shippingRateInput,omitempty"`
	// Contains addresses for carts with multiple shipping addresses.
	// Each address must contain a key which is unique in this cart.
	// Line items will use these keys to reference the addresses under their `shippingDetails`.
	// The addresses captured here are not used to determine eligible shipping methods or the applicable tax rate.
	// Only the cart's `shippingAddress` is used for this.
	ItemShippingAddresses []BaseAddress `json:"itemShippingAddresses"`
	// The code of existing DiscountCodes.
	DiscountCodes []string `json:"discountCodes"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartDraft) UnmarshalJSON(data []byte) error {
	type Alias CartDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInput != nil {
		var err error
		obj.ShippingRateInput, err = mapDiscriminatorShippingRateInputDraft(obj.ShippingRateInput)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDraft) MarshalJSON() ([]byte, error) {
	type Alias CartDraft
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

	if raw["customShipping"] == nil {
		delete(raw, "customShipping")
	}

	if raw["shipping"] == nil {
		delete(raw, "shipping")
	}

	if raw["itemShippingAddresses"] == nil {
		delete(raw, "itemShippingAddresses")
	}

	if raw["discountCodes"] == nil {
		delete(raw, "discountCodes")
	}

	return json.Marshal(raw)

}

type CartOrigin string

const (
	CartOriginCustomer CartOrigin = "Customer"
	CartOriginMerchant CartOrigin = "Merchant"
	CartOriginQuote    CartOrigin = "Quote"
)

type CartPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int  `json:"limit"`
	Count int  `json:"count"`
	Total *int `json:"total,omitempty"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset  int    `json:"offset"`
	Results []Cart `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [Cart](ctp:api:type:Cart).
*
 */
type CartReference struct {
	// Unique identifier of the referenced [Cart](ctp:api:type:Cart).
	ID string `json:"id"`
	// Contains the representation of the expanded Cart. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for Carts.
	Obj *Cart `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartReference) MarshalJSON() ([]byte, error) {
	type Alias CartReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "cart", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Cart](ctp:api:type:Cart).
*
 */
type CartResourceIdentifier struct {
	// Unique identifier of the referenced [Cart](ctp:api:type:Cart). Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [Cart](ctp:api:type:Cart). Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias CartResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "cart", Alias: (*Alias)(&obj)})
}

type CartState string

const (
	CartStateActive  CartState = "Active"
	CartStateMerged  CartState = "Merged"
	CartStateOrdered CartState = "Ordered"
)

type CartUpdate struct {
	Version int                `json:"version"`
	Actions []CartUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartUpdate) UnmarshalJSON(data []byte) error {
	type Alias CartUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorCartUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type CartUpdateAction interface{}

func mapDiscriminatorCartUpdateAction(input interface{}) (CartUpdateAction, error) {
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
		obj := CartAddCustomLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addCustomShippingMethod":
		obj := CartAddCustomShippingMethodAction{}
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
	case "addDiscountCode":
		obj := CartAddDiscountCodeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addItemShippingAddress":
		obj := CartAddItemShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addLineItem":
		obj := CartAddLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addPayment":
		obj := CartAddPaymentAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addShippingMethod":
		obj := CartAddShippingMethodAction{}
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
	case "addShoppingList":
		obj := CartAddShoppingListAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "applyDeltaToCustomLineItemShippingDetailsTargets":
		obj := CartApplyDeltaToCustomLineItemShippingDetailsTargetsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "applyDeltaToLineItemShippingDetailsTargets":
		obj := CartApplyDeltaToLineItemShippingDetailsTargetsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeCustomLineItemMoney":
		obj := CartChangeCustomLineItemMoneyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeCustomLineItemPriceMode":
		obj := CartChangeCustomLineItemPriceModeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeCustomLineItemQuantity":
		obj := CartChangeCustomLineItemQuantityAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeLineItemQuantity":
		obj := CartChangeLineItemQuantityAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTaxCalculationMode":
		obj := CartChangeTaxCalculationModeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTaxMode":
		obj := CartChangeTaxModeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTaxRoundingMode":
		obj := CartChangeTaxRoundingModeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "recalculate":
		obj := CartRecalculateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeCustomLineItem":
		obj := CartRemoveCustomLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeDiscountCode":
		obj := CartRemoveDiscountCodeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeItemShippingAddress":
		obj := CartRemoveItemShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeLineItem":
		obj := CartRemoveLineItemAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removePayment":
		obj := CartRemovePaymentAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeShippingMethod":
		obj := CartRemoveShippingMethodAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAnonymousId":
		obj := CartSetAnonymousIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setBillingAddress":
		obj := CartSetBillingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setBillingAddressCustomField":
		obj := CartSetBillingAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setBillingAddressCustomType":
		obj := CartSetBillingAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCartTotalTax":
		obj := CartSetCartTotalTaxAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCountry":
		obj := CartSetCountryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := CartSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomLineItemCustomField":
		obj := CartSetCustomLineItemCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomLineItemCustomType":
		obj := CartSetCustomLineItemCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomLineItemShippingDetails":
		obj := CartSetCustomLineItemShippingDetailsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomLineItemTaxAmount":
		obj := CartSetCustomLineItemTaxAmountAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomLineItemTaxRate":
		obj := CartSetCustomLineItemTaxRateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomShippingMethod":
		obj := CartSetCustomShippingMethodAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := CartSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomerEmail":
		obj := CartSetCustomerEmailAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomerGroup":
		obj := CartSetCustomerGroupAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomerId":
		obj := CartSetCustomerIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeleteDaysAfterLastModification":
		obj := CartSetDeleteDaysAfterLastModificationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeliveryAddressCustomField":
		obj := CartSetDeliveryAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDeliveryAddressCustomType":
		obj := CartSetDeliveryAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDirectDiscounts":
		obj := CartSetDirectDiscountsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setItemShippingAddressCustomField":
		obj := CartSetItemShippingAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setItemShippingAddressCustomType":
		obj := CartSetItemShippingAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := CartSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemCustomField":
		obj := CartSetLineItemCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemCustomType":
		obj := CartSetLineItemCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemDistributionChannel":
		obj := CartSetLineItemDistributionChannelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemPrice":
		obj := CartSetLineItemPriceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemShippingDetails":
		obj := CartSetLineItemShippingDetailsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemSupplyChannel":
		obj := CartSetLineItemSupplyChannelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemTaxAmount":
		obj := CartSetLineItemTaxAmountAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemTaxRate":
		obj := CartSetLineItemTaxRateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLineItemTotalPrice":
		obj := CartSetLineItemTotalPriceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLocale":
		obj := CartSetLocaleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingAddress":
		obj := CartSetShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingAddressCustomField":
		obj := CartSetShippingAddressCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingAddressCustomType":
		obj := CartSetShippingAddressCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingCustomField":
		obj := CartSetShippingCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingCustomType":
		obj := CartSetShippingCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingMethod":
		obj := CartSetShippingMethodAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingMethodTaxAmount":
		obj := CartSetShippingMethodTaxAmountAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingMethodTaxRate":
		obj := CartSetShippingMethodTaxRateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setShippingRateInput":
		obj := CartSetShippingRateInputAction{}
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
	case "updateItemShippingAddress":
		obj := CartUpdateItemShippingAddressAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type CustomLineItem struct {
	// Unique identifier of the CustomLineItem.
	ID string `json:"id"`
	// The name of this CustomLineItem.
	Name LocalizedString `json:"name"`
	// The cost to add to the cart.
	// The amount can be negative.
	Money TypedMoney `json:"money"`
	// Set once the `taxRate` is set.
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
	// The total price of this custom line item.
	// If custom line item is discounted, then the `totalPrice` would be the discounted custom line item price multiplied by `quantity`.
	// Otherwise a total price is just a `money` multiplied by the `quantity`.
	// `totalPrice` may or may not include the taxes: it depends on the taxRate.includedInPrice property.
	TotalPrice TypedMoney `json:"totalPrice"`
	// A unique String in the cart to identify this CustomLineItem.
	Slug string `json:"slug"`
	// The amount of a CustomLineItem in the cart.
	// Must be a positive integer.
	Quantity    int                   `json:"quantity"`
	State       []ItemState           `json:"state"`
	TaxCategory *TaxCategoryReference `json:"taxCategory,omitempty"`
	// Will be set automatically in the `Platform` TaxMode once the shipping address is set is set.
	// For the `External` tax mode the tax rate has to be set explicitly with the ExternalTaxRateDraft.
	TaxRate                    *TaxRate                             `json:"taxRate,omitempty"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	Custom                     *CustomFields                        `json:"custom,omitempty"`
	// Container for custom line item specific address(es).
	// CustomLineItem fields that can be used in query predicates: `slug`, `name`, `quantity`,
	// `money`, `state`, `discountedPricePerQuantity`.
	ShippingDetails *ItemShippingDetails `json:"shippingDetails,omitempty"`
	// Specifies whether Cart Discounts with a matching [CartDiscountCustomLineItemsTarget](ctp:api:type:CartDiscountCustomLineItemsTarget)
	// are applied to the Custom Line Item.
	PriceMode CustomLineItemPriceMode `json:"priceMode"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomLineItem) UnmarshalJSON(data []byte) error {
	type Alias CustomLineItem
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

type CustomLineItemDraft struct {
	Name LocalizedString `json:"name"`
	// The amount of a CustomLineItemin the cart.
	// Must be a positive integer.
	Quantity int    `json:"quantity"`
	Money    Money  `json:"money"`
	Slug     string `json:"slug"`
	// The given tax category will be used to select a tax rate when a cart has the TaxMode `Platform`.
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// An external tax rate can be set if the cart has the `External` TaxMode.
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Container for custom line item specific address(es).
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	// - If `Standard`, Cart Discounts with a matching [CartDiscountCustomLineItemsTarget](ctp:api:type:CartDiscountCustomLineItemsTarget)
	// are applied to the Custom Line Item.
	// - If `External`, Cart Discounts are not considered on the Custom Line Item.
	PriceMode CustomLineItemPriceMode `json:"priceMode"`
}

type CustomLineItemPriceMode string

const (
	CustomLineItemPriceModeStandard CustomLineItemPriceMode = "Standard"
	CustomLineItemPriceModeExternal CustomLineItemPriceMode = "External"
)

type CustomShippingDraft struct {
	// User-defined unique identifier of the custom Shipping Method in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	Key string `json:"key"`
	// Name of the custom Shipping Method.
	ShippingMethodName string `json:"shippingMethodName"`
	// Determines the shipping rate and Tax Rate of the associated Line Items.
	ShippingAddress *BaseAddress `json:"shippingAddress,omitempty"`
	// Determines the shipping price.
	ShippingRate ShippingRateDraft `json:"shippingRate"`
	// Used as an input to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	//
	// - Must be [ClassificationShippingRateInput](ctp:api:type:ClassificationShippingRateInput) if [ShippingRateInputType](ctp:api:type:ShippingRateInputType) is [CartClassificationType](ctp:api:type:CartClassificationType).
	// - Must be [ScoreShippingRateInput](ctp:api:type:ScoreShippingRateInput) if [ShippingRateInputType](ctp:api:type:ShippingRateInputType) is [CartScoreType](ctp:api:type:CartScoreType).
	//
	// The `shippingRateInput` cannot be set on the Cart if [CartValueType](ctp:api:type:CartValueType) is defined.
	ShippingRateInput ShippingRateInputDraft `json:"shippingRateInput,omitempty"`
	// Tax Category used to determine a shipping Tax Rate if a Cart has the `Platform` [TaxMode](ctp:api:type:TaxMode).
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// Tax Rate used to tax a shipping expense if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *string `json:"externalTaxRate,omitempty"`
	// Deliveries tied to a Shipping Method in a multi-shipping method Cart.
	// It holds information on how items are delivered to customers.
	Deliveries []Delivery `json:"deliveries"`
	// Custom Fields for the custom Shipping Method.
	Custom *string `json:"custom,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomShippingDraft) UnmarshalJSON(data []byte) error {
	type Alias CustomShippingDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInput != nil {
		var err error
		obj.ShippingRateInput, err = mapDiscriminatorShippingRateInputDraft(obj.ShippingRateInput)
		if err != nil {
			return err
		}
	}

	return nil
}

type DirectDiscount struct {
	// The unique ID of the cart discount.
	ID    string            `json:"id"`
	Value CartDiscountValue `json:"value"`
	// Empty when the `value` has type `giftLineItem`, otherwise a CartDiscountTarget is set.
	Target CartDiscountTarget `json:"target,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DirectDiscount) UnmarshalJSON(data []byte) error {
	type Alias DirectDiscount
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorCartDiscountValue(obj.Value)
		if err != nil {
			return err
		}
	}
	if obj.Target != nil {
		var err error
		obj.Target, err = mapDiscriminatorCartDiscountTarget(obj.Target)
		if err != nil {
			return err
		}
	}

	return nil
}

type DirectDiscountDraft struct {
	Value CartDiscountValue `json:"value"`
	// Empty when the `value` has type `giftLineItem`, otherwise a CartDiscountTarget is set.
	Target CartDiscountTarget `json:"target,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DirectDiscountDraft) UnmarshalJSON(data []byte) error {
	type Alias DirectDiscountDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorCartDiscountValue(obj.Value)
		if err != nil {
			return err
		}
	}
	if obj.Target != nil {
		var err error
		obj.Target, err = mapDiscriminatorCartDiscountTarget(obj.Target)
		if err != nil {
			return err
		}
	}

	return nil
}

type DiscountCodeInfo struct {
	DiscountCode DiscountCodeReference `json:"discountCode"`
	State        DiscountCodeState     `json:"state"`
}

type DiscountCodeState string

const (
	DiscountCodeStateNotActive                            DiscountCodeState = "NotActive"
	DiscountCodeStateDoesNotMatchCart                     DiscountCodeState = "DoesNotMatchCart"
	DiscountCodeStateMatchesCart                          DiscountCodeState = "MatchesCart"
	DiscountCodeStateMaxApplicationReached                DiscountCodeState = "MaxApplicationReached"
	DiscountCodeStateApplicationStoppedByPreviousDiscount DiscountCodeState = "ApplicationStoppedByPreviousDiscount"
	DiscountCodeStateNotValid                             DiscountCodeState = "NotValid"
)

type DiscountedLineItemPortion struct {
	Discount         CartDiscountReference `json:"discount"`
	DiscountedAmount TypedMoney            `json:"discountedAmount"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountedLineItemPortion) UnmarshalJSON(data []byte) error {
	type Alias DiscountedLineItemPortion
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.DiscountedAmount != nil {
		var err error
		obj.DiscountedAmount, err = mapDiscriminatorTypedMoney(obj.DiscountedAmount)
		if err != nil {
			return err
		}
	}

	return nil
}

type DiscountedLineItemPrice struct {
	Value             TypedMoney                  `json:"value"`
	IncludedDiscounts []DiscountedLineItemPortion `json:"includedDiscounts"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountedLineItemPrice) UnmarshalJSON(data []byte) error {
	type Alias DiscountedLineItemPrice
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

type DiscountedLineItemPriceForQuantity struct {
	Quantity        int                     `json:"quantity"`
	DiscountedPrice DiscountedLineItemPrice `json:"discountedPrice"`
}

type ExternalLineItemTotalPrice struct {
	Price      Money `json:"price"`
	TotalPrice Money `json:"totalPrice"`
}

type ExternalTaxAmountDraft struct {
	// The total gross amount of the item (totalNet + taxes).
	TotalGross Money                `json:"totalGross"`
	TaxRate    ExternalTaxRateDraft `json:"taxRate"`
}

type ExternalTaxRateDraft struct {
	Name string `json:"name"`
	// Percentage in the range of [0..1].
	// Must be supplied if no `subRates` are specified.
	// If `subRates` are specified
	// then the `amount` can be omitted or it must be the sum of the amounts of all `subRates`.
	Amount *float64 `json:"amount,omitempty"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	// The state in the country
	State *string `json:"state,omitempty"`
	// For countries (e.g.
	// the US) where the total tax is a combination of multiple taxes (e.g.
	// state and local taxes).
	SubRates []SubRate `json:"subRates"`
	// The default value for `includedInPrice` is FALSE.
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

type InventoryMode string

const (
	InventoryModeTrackOnly      InventoryMode = "TrackOnly"
	InventoryModeReserveOnOrder InventoryMode = "ReserveOnOrder"
	InventoryModeNone           InventoryMode = "None"
)

type ItemShippingDetails struct {
	// Used to map what sub-quantity should be shipped to which address.
	// Duplicate address keys are not allowed.
	Targets []ItemShippingTarget `json:"targets"`
	// `true` if the quantity of the (custom) line item is equal to the sum of the sub-quantities in `targets`, `false` otherwise.
	// A cart cannot be ordered when the value is `false`.
	// The error InvalidItemShippingDetails will be triggered.
	Valid bool `json:"valid"`
}

type ItemShippingDetailsDraft struct {
	// Used to capture one or more (custom) line item specific shipping addresses.
	// By specifying sub-quantities, it is possible to set multiple shipping addresses for one line item.
	// A cart can have `shippingDetails` where the `targets` sum does not match the quantity of the line item or custom line item.
	// For the order creation and order updates the `targets` sum must match the quantity.
	Targets []ItemShippingTarget `json:"targets"`
}

type ItemShippingTarget struct {
	// The key of the address in the cart's `itemShippingAddresses`
	AddressKey string `json:"addressKey"`
	// The quantity of items that should go to the address with the specified `addressKey`.
	// Only positive values are allowed.
	// Using `0` as quantity is also possible in a draft object, but the element will not be present in the resulting ItemShippingDetails.
	Quantity int `json:"quantity"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	//
	// It connects Line Item quantities with individual shipping addresses.
	ShippingMethodKey *string `json:"shippingMethodKey,omitempty"`
}

type LineItem struct {
	// Unique identifier of the LineItem.
	ID        string `json:"id"`
	ProductId string `json:"productId"`
	// User-defined unique identifier of the [Product](ctp:api:type:Product).
	// Only present on Line Items in a [Cart](ctp:api:type:Cart) when the `key` is available on that specific Product at the time the Line Item is created or updated on the Cart. On [Order](/ctp:api:type:Order) resources this field is only present when the `key` is available on the specific Product at the time the Order is created from the Cart. This field is in general not present on Carts that had no updates until 3 December 2021 and on Orders created before this date.
	ProductKey *string `json:"productKey,omitempty"`
	// The product name.
	Name LocalizedString `json:"name"`
	// The slug of a product is inserted on the fly.
	// It is always up-to-date and can therefore be used to link to the product detail page of the product.
	// It is empty if the product has been deleted.
	// The slug is also empty if the cart or order is retrieved via Reference Expansion or is a snapshot in a Message.
	ProductSlug *LocalizedString     `json:"productSlug,omitempty"`
	ProductType ProductTypeReference `json:"productType"`
	// The variant data is saved when the variant is added to the cart, and not updated automatically.
	// It can manually be updated with the Recalculate update action.
	Variant ProductVariant `json:"variant"`
	// The price of a line item is selected from the product variant according to the Product's [priceMode](ctp:api:type:Product) value.
	// If the `priceMode` is `Embedded` [ProductPriceMode](ctp:api:type:ProductPriceModeEnum) and the `variant` field hasn't been updated, the price may not correspond to a price in `variant.prices`.
	Price Price `json:"price"`
	// Set once the `taxRate` is set.
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
	// Taxed price of the Shipping Method that is set automatically after `perMethodTaxRate` is set.
	TaxedPricePortions []MethodTaxedPrice `json:"taxedPricePortions"`
	// The total price of this line item.
	// If the line item is discounted, then the `totalPrice` is the DiscountedLineItemPriceForQuantity multiplied by `quantity`.
	// Otherwise the total price is the product price multiplied by the `quantity`.
	// `totalPrice` may or may not include the taxes: it depends on the taxRate.includedInPrice property.
	TotalPrice TypedMoney `json:"totalPrice"`
	// The amount of a LineItem in the cart.
	// Must be a positive integer.
	Quantity int `json:"quantity"`
	// When the line item was added to the cart. Optional for backwards
	// compatibility reasons only.
	AddedAt *time.Time  `json:"addedAt,omitempty"`
	State   []ItemState `json:"state"`
	// Will be set automatically in the `Platform` TaxMode once the shipping address is set is set.
	// For the `External` tax mode the tax rate has to be set explicitly with the ExternalTaxRateDraft.
	TaxRate *TaxRate `json:"taxRate,omitempty"`
	// Tax Rate per Shipping Method that is automatically set after the [Shipping Method is added](ctp:api:type:CartAddShippingMethodAction) to a Cart with the `Platform` [TaxMode](ctp:api:type:TaxMode) and `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	//
	// For the `External` [TaxMode](ctp:api:type:TaxMode), the Tax Rate must be set with [ExternalTaxRateDraft](ctp:api:type:ExternalTaxRateDraft).
	PerMethodTaxRate []MethodTaxRate `json:"perMethodTaxRate"`
	// The supply channel identifies the inventory entries that should be reserved.
	// The channel has
	// the role InventorySupply.
	SupplyChannel *ChannelReference `json:"supplyChannel,omitempty"`
	// The distribution channel is used to select a ProductPrice.
	// The channel has the role ProductDistribution.
	DistributionChannel        *ChannelReference                    `json:"distributionChannel,omitempty"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	PriceMode                  LineItemPriceMode                    `json:"priceMode"`
	LineItemMode               LineItemMode                         `json:"lineItemMode"`
	Custom                     *CustomFields                        `json:"custom,omitempty"`
	// Inventory mode specific to the line item only, valid for the entire `quantity` of the line item.
	// Only present if inventory mode is different from the `inventoryMode` specified on the [Cart](ctp:api:type:Cart).
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	// Container for line item specific address(es).
	ShippingDetails *ItemShippingDetails `json:"shippingDetails,omitempty"`
	// The date when the LineItem was last modified by one of the following actions
	// setLineItemShippingDetails, addLineItem, removeLineItem, or changeLineItemQuantity.
	// Optional only for backwards compatible reasons. When the LineItem is created lastModifiedAt is set to addedAt.
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *LineItem) UnmarshalJSON(data []byte) error {
	type Alias LineItem
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

type LineItemDraft struct {
	ProductId *string `json:"productId,omitempty"`
	VariantId *int    `json:"variantId,omitempty"`
	Sku       *string `json:"sku,omitempty"`
	// The amount of a `LineItem`in the cart.
	// Must be a positive integer.
	Quantity *int `json:"quantity,omitempty"`
	// When the line item was added to the cart. Optional for backwards
	// compatibility reasons only.
	AddedAt *time.Time `json:"addedAt,omitempty"`
	// By providing supply channel information, you can unique identify
	// inventory entries that should be reserved.
	// The provided channel should have
	// the InventorySupply role.
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	// The channel is used to select a ProductPrice.
	// The provided channel should have the ProductDistribution role.
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
	// An external tax rate can be set if the cart has the `External` TaxMode.
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Sets the line item `price` to the given value and sets the line item `priceMode` to `ExternalPrice` LineItemPriceMode.
	ExternalPrice *Money `json:"externalPrice,omitempty"`
	// Sets the line item `price` and `totalPrice` to the given values and sets the line item `priceMode` to `ExternalTotal` LineItemPriceMode.
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	// Inventory mode specific to the line item only, valid for the entire `quantity` of the line item.
	// Set only if inventory mode should be different from the `inventoryMode` specified on the [Cart](ctp:api:type:Cart).
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	// Container for line item specific address(es).
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
}

type LineItemMode string

const (
	LineItemModeStandard     LineItemMode = "Standard"
	LineItemModeGiftLineItem LineItemMode = "GiftLineItem"
)

type LineItemPriceMode string

const (
	LineItemPriceModePlatform      LineItemPriceMode = "Platform"
	LineItemPriceModeExternalTotal LineItemPriceMode = "ExternalTotal"
	LineItemPriceModeExternalPrice LineItemPriceMode = "ExternalPrice"
)

type MethodTaxRate struct {
	// User-defined unique identifier of the Shipping Method in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingMethodKey string `json:"shippingMethodKey"`
	// Tax Rate for the Shipping Method.
	TaxRate *TaxRate `json:"taxRate,omitempty"`
}

type MethodTaxedPrice struct {
	// User-defined unique identifier of the Shipping Method in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingMethodKey string `json:"shippingMethodKey"`
	// Taxed price for the Shipping Method.
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
}

type ReplicaCartDraft struct {
	Reference interface{} `json:"reference"`
	// User-specific unique identifier of the cart.
	Key *string `json:"key,omitempty"`
}

type RoundingMode string

const (
	RoundingModeHalfEven RoundingMode = "HalfEven"
	RoundingModeHalfUp   RoundingMode = "HalfUp"
	RoundingModeHalfDown RoundingMode = "HalfDown"
)

type Shipping struct {
	// User-defined unique identifier of the Shipping Method in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey string `json:"shippingKey"`
	// Automatically set when the [Shipping Method is added](ctp:api:type:CartAddShippingMethodAction).
	ShippingInfo ShippingInfo `json:"shippingInfo"`
	// Determines the shipping rates and Tax Rates of the associated Line Item quantities.
	ShippingAddress Address `json:"shippingAddress"`
	// Used as an input to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	//
	// - Must be [ClassificationShippingRateInput](ctp:api:type:ClassificationShippingRateInput) if [ShippingRateInputType](ctp:api:type:ShippingRateInputType) is [CartClassificationType](ctp:api:type:CartClassificationType).
	// - Must be [ScoreShippingRateInput](ctp:api:type:ScoreShippingRateInput) if [ShippingRateInputType](ctp:api:type:ShippingRateInputType) is [CartScoreType](ctp:api:type:CartScoreType).
	ShippingRateInput ShippingRateInput `json:"shippingRateInput,omitempty"`
	// Custom Fields of Shipping.
	ShippingCustomFields *CustomFields `json:"shippingCustomFields,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Shipping) UnmarshalJSON(data []byte) error {
	type Alias Shipping
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
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

/**
*	Wraps all shipping-related information (such as address, rate, deliveries) per Shipping Method for Carts with multiple Shipping Methods.
*
 */
type ShippingDraft struct {
	// User-defined unique identifier of the Shipping Method in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	Key string `json:"key"`
	// Shipping Methods added to the Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingMethod *ShippingMethodReference `json:"shippingMethod,omitempty"`
	// Determines the shipping rate and Tax Rate of the associated Line Items.
	ShippingAddress *BaseAddress `json:"shippingAddress,omitempty"`
	// Used as an input to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	//
	// - Must be [ClassificationShippingRateInput](ctp:api:type:ClassificationShippingRateInput) if [ShippingRateInputType](ctp:api:type:ShippingRateInputType) is [CartClassificationType](ctp:api:type:CartClassificationType).
	// - Must be [ScoreShippingRateInput](ctp:api:type:ScoreShippingRateInput) if [ShippingRateInputType](ctp:api:type:ShippingRateInputType) is [CartScoreType](ctp:api:type:CartScoreType).
	//
	// The `shippingRateInput` cannot be set on the Cart if [CartValueType](ctp:api:type:CartValueType) is defined.
	ShippingRateInput ShippingRateInputDraft `json:"shippingRateInput,omitempty"`
	// Tax Rate used for taxing a shipping expense if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *string `json:"externalTaxRate,omitempty"`
	// Holds information on how items are delivered to customers.
	Deliveries []Delivery `json:"deliveries"`
	// Custom Fields for Shipping.
	Custom *string `json:"custom,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingDraft) UnmarshalJSON(data []byte) error {
	type Alias ShippingDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInput != nil {
		var err error
		obj.ShippingRateInput, err = mapDiscriminatorShippingRateInputDraft(obj.ShippingRateInput)
		if err != nil {
			return err
		}
	}

	return nil
}

type ShippingInfo struct {
	ShippingMethodName string `json:"shippingMethodName"`
	// Determined based on the ShippingRate and its tiered prices, and either the sum of LineItem prices or the `shippingRateInput` field.
	Price TypedMoney `json:"price"`
	// The shipping rate used to determine the price.
	ShippingRate ShippingRate `json:"shippingRate"`
	// Set once the `taxRate` is set.
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
	// Will be set automatically in the `Platform` TaxMode once the shipping address is set is set.
	// For the `External` tax mode the tax rate has to be set explicitly with the ExternalTaxRateDraft.
	TaxRate     *TaxRate              `json:"taxRate,omitempty"`
	TaxCategory *TaxCategoryReference `json:"taxCategory,omitempty"`
	// Not set if custom shipping method is used.
	ShippingMethod *ShippingMethodReference `json:"shippingMethod,omitempty"`
	// Deliveries are compilations of information on how the articles are being delivered to the customers.
	Deliveries      []Delivery               `json:"deliveries"`
	DiscountedPrice *DiscountedLineItemPrice `json:"discountedPrice,omitempty"`
	// Indicates whether the ShippingMethod referenced in this ShippingInfo is allowed for the cart or not.
	ShippingMethodState ShippingMethodState `json:"shippingMethodState"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingInfo) UnmarshalJSON(data []byte) error {
	type Alias ShippingInfo
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
func (obj ShippingInfo) MarshalJSON() ([]byte, error) {
	type Alias ShippingInfo
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

type ShippingMethodState string

const (
	ShippingMethodStateDoesNotMatchCart ShippingMethodState = "DoesNotMatchCart"
	ShippingMethodStateMatchesCart      ShippingMethodState = "MatchesCart"
)

type ShippingMode string

const (
	ShippingModeSingle   ShippingMode = "Single"
	ShippingModeMultiple ShippingMode = "Multiple"
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
	// JSON object where the keys are of type [Locale](ctp:api:type:Locale), and the values are the strings used for the corresponding language.
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

type ShippingRateInputDraft interface{}

func mapDiscriminatorShippingRateInputDraft(input interface{}) (ShippingRateInputDraft, error) {
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
		obj := ClassificationShippingRateInputDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Score":
		obj := ScoreShippingRateInputDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ClassificationShippingRateInputDraft struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ClassificationShippingRateInputDraft) MarshalJSON() ([]byte, error) {
	type Alias ClassificationShippingRateInputDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Classification", Alias: (*Alias)(&obj)})
}

type ScoreShippingRateInputDraft struct {
	Score int `json:"score"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ScoreShippingRateInputDraft) MarshalJSON() ([]byte, error) {
	type Alias ScoreShippingRateInputDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Score", Alias: (*Alias)(&obj)})
}

type TaxCalculationMode string

const (
	TaxCalculationModeLineItemLevel  TaxCalculationMode = "LineItemLevel"
	TaxCalculationModeUnitPriceLevel TaxCalculationMode = "UnitPriceLevel"
)

type TaxMode string

const (
	TaxModePlatform       TaxMode = "Platform"
	TaxModeExternal       TaxMode = "External"
	TaxModeExternalAmount TaxMode = "ExternalAmount"
	TaxModeDisabled       TaxMode = "Disabled"
)

type TaxPortion struct {
	Name *string `json:"name,omitempty"`
	// A number in the range [0..1]
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

type TaxPortionDraft struct {
	Name *string `json:"name,omitempty"`
	Rate float64 `json:"rate"`
	// Draft type that stores amounts in cent precision for the specified currency.
	//
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	Amount Money `json:"amount"`
}

type TaxedItemPrice struct {
	TotalNet TypedMoney `json:"totalNet"`
	// TaxedItemPrice fields can not be used in query predicates.
	TotalGross TypedMoney `json:"totalGross"`
	// Calculated automatically as the subtraction of `totalGross` - `totalNet`.
	TotalTax TypedMoney `json:"totalTax,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *TaxedItemPrice) UnmarshalJSON(data []byte) error {
	type Alias TaxedItemPrice
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
	if obj.TotalTax != nil {
		var err error
		obj.TotalTax, err = mapDiscriminatorTypedMoney(obj.TotalTax)
		if err != nil {
			return err
		}
	}

	return nil
}

type TaxedPrice struct {
	TotalNet   TypedMoney `json:"totalNet"`
	TotalGross TypedMoney `json:"totalGross"`
	// TaxedPrice fields that can be used in query predicates: `totalNet`, `totalGross`.
	TaxPortions []TaxPortion `json:"taxPortions"`
	// Calculated automatically as the subtraction of `totalGross` - `totalNet`.
	TotalTax TypedMoney `json:"totalTax,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *TaxedPrice) UnmarshalJSON(data []byte) error {
	type Alias TaxedPrice
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
	if obj.TotalTax != nil {
		var err error
		obj.TotalTax, err = mapDiscriminatorTypedMoney(obj.TotalTax)
		if err != nil {
			return err
		}
	}

	return nil
}

type TaxedPriceDraft struct {
	// Draft type that stores amounts in cent precision for the specified currency.
	//
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	TotalNet Money `json:"totalNet"`
	// Draft type that stores amounts in cent precision for the specified currency.
	//
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	TotalGross  Money             `json:"totalGross"`
	TaxPortions []TaxPortionDraft `json:"taxPortions"`
}

type CartAddCustomLineItemAction struct {
	// Draft type that stores amounts in cent precision for the specified currency.
	//
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	Money Money `json:"money"`
	// JSON object where the keys are of type [Locale](ctp:api:type:Locale), and the values are the strings used for the corresponding language.
	Name     LocalizedString `json:"name"`
	Quantity int             `json:"quantity"`
	Slug     string          `json:"slug"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [TaxCategory](ctp:api:type:TaxCategory).
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// The representation used when creating or updating a [customizable data type](/../api/projects/types#list-of-customizable-data-types) with Custom Fields.
	Custom          *CustomFieldsDraft    `json:"custom,omitempty"`
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// - If `Standard`, Cart Discounts with a matching [CartDiscountCustomLineItemsTarget](ctp:api:type:CartDiscountCustomLineItemsTarget)
	// are applied to the Custom Line Item.
	// - If `External`, Cart Discounts are not considered on the Custom Line Item.
	PriceMode *CustomLineItemPriceMode `json:"priceMode,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartAddCustomLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias CartAddCustomLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addCustomLineItem", Alias: (*Alias)(&obj)})
}

type CartAddCustomShippingMethodAction struct {
	// User-defined unique identifier of the custom Shipping Method in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey string `json:"shippingKey"`
	// Name of the custom Shipping Method.
	ShippingMethodName string `json:"shippingMethodName"`
	// Determines the shipping rate and Tax Rate of the associated Line Items.
	ShippingAddress BaseAddress `json:"shippingAddress"`
	// Determines the shipping price.
	ShippingRate ShippingRateDraft `json:"shippingRate"`
	// Used as an input to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	//
	// - Must be [ClassificationShippingRateInput](ctp:api:type:ClassificationShippingRateInput) if [ShippingRateInputType](ctp:api:type:ShippingRateInputType) is [CartClassificationType](ctp:api:type:CartClassificationType).
	// - Must be [ScoreShippingRateInput](ctp:api:type:ScoreShippingRateInput) if [ShippingRateInputType](ctp:api:type:ShippingRateInputType) is [CartScoreType](ctp:api:type:CartScoreType).
	//
	// The `shippingRateInput` cannot be set on the Cart if [CartValueType](ctp:api:type:CartValueType) is defined.
	ShippingRateInput ShippingRateInputDraft `json:"shippingRateInput,omitempty"`
	// Tax Category used to determine a shipping Tax Rate if a Cart has the `Platform` [TaxMode](ctp:api:type:TaxMode).
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// Tax Rate used to tax a shipping expense if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *string `json:"externalTaxRate,omitempty"`
	// Deliveries tied to a Shipping Method in a multi-shipping method Cart.
	// It holds information on how items are delivered to customers.
	Deliveries []Delivery `json:"deliveries"`
	// Custom Fields for the custom Shipping Method.
	Custom *string `json:"custom,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartAddCustomShippingMethodAction) UnmarshalJSON(data []byte) error {
	type Alias CartAddCustomShippingMethodAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInput != nil {
		var err error
		obj.ShippingRateInput, err = mapDiscriminatorShippingRateInputDraft(obj.ShippingRateInput)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartAddCustomShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias CartAddCustomShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addCustomShippingMethod", Alias: (*Alias)(&obj)})
}

type CartAddDiscountCodeAction struct {
	Code string `json:"code"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartAddDiscountCodeAction) MarshalJSON() ([]byte, error) {
	type Alias CartAddDiscountCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDiscountCode", Alias: (*Alias)(&obj)})
}

type CartAddItemShippingAddressAction struct {
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartAddItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CartAddItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addItemShippingAddress", Alias: (*Alias)(&obj)})
}

type CartAddLineItemAction struct {
	// The representation used when creating or updating a [customizable data type](/../api/projects/types#list-of-customizable-data-types) with Custom Fields.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Channel](ctp:api:type:Channel).
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
	ExternalTaxRate     *ExternalTaxRateDraft      `json:"externalTaxRate,omitempty"`
	ProductId           *string                    `json:"productId,omitempty"`
	VariantId           *int                       `json:"variantId,omitempty"`
	Sku                 *string                    `json:"sku,omitempty"`
	Quantity            *int                       `json:"quantity,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Channel](ctp:api:type:Channel).
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	// Draft type that stores amounts in cent precision for the specified currency.
	//
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	ExternalPrice      *Money                      `json:"externalPrice,omitempty"`
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ShippingDetails    *ItemShippingDetailsDraft   `json:"shippingDetails,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias CartAddLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
}

type CartAddPaymentAction struct {
	Payment PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartAddPaymentAction) MarshalJSON() ([]byte, error) {
	type Alias CartAddPaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPayment", Alias: (*Alias)(&obj)})
}

/**
*	This update action fails with an [InvalidOperation](ctp:api:type:InvalidOperationError) error if the referenced shipping method has a predicate that does not match the Cart.
*
 */
type CartAddShippingMethodAction struct {
	// User-defined unique identifier of the Shipping Method in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey string `json:"shippingKey"`
	// Value to set.
	// If empty, any existing value is removed.
	ShippingMethod ShippingMethodReference `json:"shippingMethod"`
	// Determines the shipping rate and Tax Rate of the Line Items.
	ShippingAddress BaseAddress `json:"shippingAddress"`
	// Used as an input to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	//
	// - Must be [ClassificationShippingRateInput](ctp:api:type:ClassificationShippingRateInput) if [ShippingRateInputType](ctp:api:type:ShippingRateInputType) is [CartClassificationType](ctp:api:type:CartClassificationType).
	// - Must be [ScoreShippingRateInput](ctp:api:type:ScoreShippingRateInput) if [ShippingRateInputType](ctp:api:type:ShippingRateInputType) is [CartScoreType](ctp:api:type:CartScoreType).
	//
	// The `shippingRateInput` cannot be set on the Cart if [CartValueType](ctp:api:type:CartValueType) is defined.
	ShippingRateInput ShippingRateInputDraft `json:"shippingRateInput,omitempty"`
	// Tax Rate used to tax a shipping expense if a Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *string `json:"externalTaxRate,omitempty"`
	// Deliveries tied to a Shipping Method in a multi-shipping method Cart.
	// It holds information on how items are delivered to customers.
	Deliveries []Delivery `json:"deliveries"`
	// Custom Fields for the Shipping Method.
	Custom *string `json:"custom,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartAddShippingMethodAction) UnmarshalJSON(data []byte) error {
	type Alias CartAddShippingMethodAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInput != nil {
		var err error
		obj.ShippingRateInput, err = mapDiscriminatorShippingRateInputDraft(obj.ShippingRateInput)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartAddShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias CartAddShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingMethod", Alias: (*Alias)(&obj)})
}

type CartAddShoppingListAction struct {
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [ShoppingList](ctp:api:type:ShoppingList).
	ShoppingList ShoppingListResourceIdentifier `json:"shoppingList"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Channel](ctp:api:type:Channel).
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Channel](ctp:api:type:Channel).
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartAddShoppingListAction) MarshalJSON() ([]byte, error) {
	type Alias CartAddShoppingListAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShoppingList", Alias: (*Alias)(&obj)})
}

type CartApplyDeltaToCustomLineItemShippingDetailsTargetsAction struct {
	CustomLineItemId string               `json:"customLineItemId"`
	TargetsDelta     []ItemShippingTarget `json:"targetsDelta"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartApplyDeltaToCustomLineItemShippingDetailsTargetsAction) MarshalJSON() ([]byte, error) {
	type Alias CartApplyDeltaToCustomLineItemShippingDetailsTargetsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "applyDeltaToCustomLineItemShippingDetailsTargets", Alias: (*Alias)(&obj)})
}

type CartApplyDeltaToLineItemShippingDetailsTargetsAction struct {
	LineItemId   string               `json:"lineItemId"`
	TargetsDelta []ItemShippingTarget `json:"targetsDelta"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartApplyDeltaToLineItemShippingDetailsTargetsAction) MarshalJSON() ([]byte, error) {
	type Alias CartApplyDeltaToLineItemShippingDetailsTargetsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "applyDeltaToLineItemShippingDetailsTargets", Alias: (*Alias)(&obj)})
}

type CartChangeCustomLineItemMoneyAction struct {
	CustomLineItemId string `json:"customLineItemId"`
	// Draft type that stores amounts in cent precision for the specified currency.
	//
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	Money Money `json:"money"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartChangeCustomLineItemMoneyAction) MarshalJSON() ([]byte, error) {
	type Alias CartChangeCustomLineItemMoneyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCustomLineItemMoney", Alias: (*Alias)(&obj)})
}

type CartChangeCustomLineItemPriceModeAction struct {
	// ID of the Custom Line Item to be updated.
	CustomLineItemId string `json:"customLineItemId"`
	// New value to set.
	Mode CustomLineItemPriceMode `json:"mode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartChangeCustomLineItemPriceModeAction) MarshalJSON() ([]byte, error) {
	type Alias CartChangeCustomLineItemPriceModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCustomLineItemPriceMode", Alias: (*Alias)(&obj)})
}

type CartChangeCustomLineItemQuantityAction struct {
	CustomLineItemId string `json:"customLineItemId"`
	Quantity         int    `json:"quantity"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartChangeCustomLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias CartChangeCustomLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCustomLineItemQuantity", Alias: (*Alias)(&obj)})
}

type CartChangeLineItemQuantityAction struct {
	LineItemId string `json:"lineItemId"`
	Quantity   int    `json:"quantity"`
	// Draft type that stores amounts in cent precision for the specified currency.
	//
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	ExternalPrice      *Money                      `json:"externalPrice,omitempty"`
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartChangeLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias CartChangeLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemQuantity", Alias: (*Alias)(&obj)})
}

type CartChangeTaxCalculationModeAction struct {
	TaxCalculationMode TaxCalculationMode `json:"taxCalculationMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartChangeTaxCalculationModeAction) MarshalJSON() ([]byte, error) {
	type Alias CartChangeTaxCalculationModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxCalculationMode", Alias: (*Alias)(&obj)})
}

type CartChangeTaxModeAction struct {
	TaxMode TaxMode `json:"taxMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartChangeTaxModeAction) MarshalJSON() ([]byte, error) {
	type Alias CartChangeTaxModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxMode", Alias: (*Alias)(&obj)})
}

type CartChangeTaxRoundingModeAction struct {
	TaxRoundingMode RoundingMode `json:"taxRoundingMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartChangeTaxRoundingModeAction) MarshalJSON() ([]byte, error) {
	type Alias CartChangeTaxRoundingModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxRoundingMode", Alias: (*Alias)(&obj)})
}

type CartRecalculateAction struct {
	// If set to `true`, the line item product data (`name`, `variant` and `productType`) will also be updated.
	// If set to `false`, only the prices and tax rates of the line item will be updated.
	// Notice that if the Product's [priceMode](ctp:api:type:Product) value is `Embedded` [ProductPriceMode](ctp:api:type:ProductPriceModeEnum), the updated price of a line item may not correspond to a price in `variant.prices` anymore.
	UpdateProductData *bool `json:"updateProductData,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartRecalculateAction) MarshalJSON() ([]byte, error) {
	type Alias CartRecalculateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "recalculate", Alias: (*Alias)(&obj)})
}

type CartRemoveCustomLineItemAction struct {
	CustomLineItemId string `json:"customLineItemId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartRemoveCustomLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias CartRemoveCustomLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeCustomLineItem", Alias: (*Alias)(&obj)})
}

type CartRemoveDiscountCodeAction struct {
	// [Reference](ctp:api:type:Reference) to a [DiscountCode](ctp:api:type:DiscountCode).
	DiscountCode DiscountCodeReference `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartRemoveDiscountCodeAction) MarshalJSON() ([]byte, error) {
	type Alias CartRemoveDiscountCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeDiscountCode", Alias: (*Alias)(&obj)})
}

type CartRemoveItemShippingAddressAction struct {
	AddressKey string `json:"addressKey"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartRemoveItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CartRemoveItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeItemShippingAddress", Alias: (*Alias)(&obj)})
}

type CartRemoveLineItemAction struct {
	LineItemId string `json:"lineItemId"`
	Quantity   *int   `json:"quantity,omitempty"`
	// Draft type that stores amounts in cent precision for the specified currency.
	//
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	ExternalPrice           *Money                      `json:"externalPrice,omitempty"`
	ExternalTotalPrice      *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ShippingDetailsToRemove *ItemShippingDetailsDraft   `json:"shippingDetailsToRemove,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartRemoveLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias CartRemoveLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLineItem", Alias: (*Alias)(&obj)})
}

type CartRemovePaymentAction struct {
	Payment PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartRemovePaymentAction) MarshalJSON() ([]byte, error) {
	type Alias CartRemovePaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePayment", Alias: (*Alias)(&obj)})
}

type CartRemoveShippingMethodAction struct {
	// User-defined unique identifier of the Shipping Method to remove in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey string `json:"shippingKey"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartRemoveShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias CartRemoveShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeShippingMethod", Alias: (*Alias)(&obj)})
}

type CartSetAnonymousIdAction struct {
	// If not set, any existing anonymous ID will be removed.
	AnonymousId *string `json:"anonymousId,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetAnonymousIdAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetAnonymousIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAnonymousId", Alias: (*Alias)(&obj)})
}

type CartSetBillingAddressAction struct {
	Address *BaseAddress `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddress", Alias: (*Alias)(&obj)})
}

type CartSetBillingAddressCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetBillingAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetBillingAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddressCustomField", Alias: (*Alias)(&obj)})
}

type CartSetBillingAddressCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the `billingAddress` with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the `billingAddress`.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `billingAddress`.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetBillingAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetBillingAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddressCustomType", Alias: (*Alias)(&obj)})
}

type CartSetCartTotalTaxAction struct {
	// The total gross amount of the cart (totalNet + taxes).
	ExternalTotalGross  Money             `json:"externalTotalGross"`
	ExternalTaxPortions []TaxPortionDraft `json:"externalTaxPortions"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetCartTotalTaxAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCartTotalTaxAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCartTotalTax", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["externalTaxPortions"] == nil {
		delete(raw, "externalTaxPortions")
	}

	return json.Marshal(raw)

}

type CartSetCountryAction struct {
	// Two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country *string `json:"country,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetCountryAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCountryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCountry", Alias: (*Alias)(&obj)})
}

type CartSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type CartSetCustomLineItemCustomFieldAction struct {
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
func (obj CartSetCustomLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomField", Alias: (*Alias)(&obj)})
}

type CartSetCustomLineItemCustomTypeAction struct {
	CustomLineItemId string `json:"customLineItemId"`
	// Defines the [Type](ctp:api:type:Type) that extends the CustomLineItem with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the CustomLineItem.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the CustomLineItem.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetCustomLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomType", Alias: (*Alias)(&obj)})
}

type CartSetCustomLineItemShippingDetailsAction struct {
	CustomLineItemId string                    `json:"customLineItemId"`
	ShippingDetails  *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetCustomLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

type CartSetCustomLineItemTaxAmountAction struct {
	CustomLineItemId  string                  `json:"customLineItemId"`
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetCustomLineItemTaxAmountAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomLineItemTaxAmountAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemTaxAmount", Alias: (*Alias)(&obj)})
}

type CartSetCustomLineItemTaxRateAction struct {
	CustomLineItemId string                `json:"customLineItemId"`
	ExternalTaxRate  *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetCustomLineItemTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomLineItemTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemTaxRate", Alias: (*Alias)(&obj)})
}

type CartSetCustomShippingMethodAction struct {
	ShippingMethodName string            `json:"shippingMethodName"`
	ShippingRate       ShippingRateDraft `json:"shippingRate"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [TaxCategory](ctp:api:type:TaxCategory).
	TaxCategory     *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	ExternalTaxRate *ExternalTaxRateDraft          `json:"externalTaxRate,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetCustomShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomShippingMethod", Alias: (*Alias)(&obj)})
}

type CartSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the Cart with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Cart.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Cart.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type CartSetCustomerEmailAction struct {
	Email string `json:"email"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetCustomerEmailAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomerEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerEmail", Alias: (*Alias)(&obj)})
}

type CartSetCustomerGroupAction struct {
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [CustomerGroup](ctp:api:type:CustomerGroup).
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetCustomerGroupAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomerGroupAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerGroup", Alias: (*Alias)(&obj)})
}

type CartSetCustomerIdAction struct {
	// If set, a customer with the given ID must exist in the project.
	CustomerId *string `json:"customerId,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetCustomerIdAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomerIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerId", Alias: (*Alias)(&obj)})
}

type CartSetDeleteDaysAfterLastModificationAction struct {
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetDeleteDaysAfterLastModificationAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetDeleteDaysAfterLastModificationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeleteDaysAfterLastModification", Alias: (*Alias)(&obj)})
}

type CartSetDeliveryAddressCustomFieldAction struct {
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
func (obj CartSetDeliveryAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetDeliveryAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryAddressCustomField", Alias: (*Alias)(&obj)})
}

type CartSetDeliveryAddressCustomTypeAction struct {
	DeliveryId string `json:"deliveryId"`
	// Defines the [Type](ctp:api:type:Type) that extends the `address` in a Delivery with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the `address` in a Delivery.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `address` in a Delivery.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetDeliveryAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetDeliveryAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryAddressCustomType", Alias: (*Alias)(&obj)})
}

type CartSetDirectDiscountsAction struct {
	Discounts []DirectDiscountDraft `json:"discounts"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetDirectDiscountsAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetDirectDiscountsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDirectDiscounts", Alias: (*Alias)(&obj)})
}

type CartSetItemShippingAddressCustomFieldAction struct {
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
func (obj CartSetItemShippingAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetItemShippingAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setItemShippingAddressCustomField", Alias: (*Alias)(&obj)})
}

type CartSetItemShippingAddressCustomTypeAction struct {
	AddressKey string `json:"addressKey"`
	// Defines the [Type](ctp:api:type:Type) that extends the `itemShippingAddress` with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the `itemShippingAddress`.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `itemShippingAddress`.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetItemShippingAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetItemShippingAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setItemShippingAddressCustomType", Alias: (*Alias)(&obj)})
}

type CartSetKeyAction struct {
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type CartSetLineItemCustomFieldAction struct {
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
func (obj CartSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

type CartSetLineItemCustomTypeAction struct {
	LineItemId string `json:"lineItemId"`
	// Defines the [Type](ctp:api:type:Type) that extends the LineItem with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the LineItem.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the LineItem.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

type CartSetLineItemDistributionChannelAction struct {
	LineItemId string `json:"lineItemId"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Channel](ctp:api:type:Channel).
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetLineItemDistributionChannelAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemDistributionChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemDistributionChannel", Alias: (*Alias)(&obj)})
}

type CartSetLineItemPriceAction struct {
	LineItemId string `json:"lineItemId"`
	// Draft type that stores amounts in cent precision for the specified currency.
	//
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	ExternalPrice *Money `json:"externalPrice,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetLineItemPriceAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemPrice", Alias: (*Alias)(&obj)})
}

type CartSetLineItemShippingDetailsAction struct {
	LineItemId      string                    `json:"lineItemId"`
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

type CartSetLineItemSupplyChannelAction struct {
	LineItemId string `json:"lineItemId"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Channel](ctp:api:type:Channel).
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetLineItemSupplyChannelAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemSupplyChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemSupplyChannel", Alias: (*Alias)(&obj)})
}

type CartSetLineItemTaxAmountAction struct {
	LineItemId        string                  `json:"lineItemId"`
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
	// `key` of the [ShippingMethod](ctp:api:type:ShippingMethod) used for this Line Item.```
	// This is required for Carts with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetLineItemTaxAmountAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemTaxAmountAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemTaxAmount", Alias: (*Alias)(&obj)})
}

type CartSetLineItemTaxRateAction struct {
	LineItemId      string                `json:"lineItemId"`
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// `key` of the [ShippingMethod](ctp:api:type:ShippingMethod) used for this Line Item.```
	// This is required for Carts with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetLineItemTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemTaxRate", Alias: (*Alias)(&obj)})
}

type CartSetLineItemTotalPriceAction struct {
	LineItemId         string                      `json:"lineItemId"`
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetLineItemTotalPriceAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemTotalPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemTotalPrice", Alias: (*Alias)(&obj)})
}

type CartSetLocaleAction struct {
	Locale *string `json:"locale,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

type CartSetShippingAddressAction struct {
	Address *BaseAddress `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddress", Alias: (*Alias)(&obj)})
}

type CartSetShippingAddressCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetShippingAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetShippingAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddressCustomField", Alias: (*Alias)(&obj)})
}

type CartSetShippingAddressCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the `shippingAddress` with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the `shippingAddress`.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `shippingAddress`.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetShippingAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetShippingAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddressCustomType", Alias: (*Alias)(&obj)})
}

type CartSetShippingCustomFieldAction struct {
	// User-defined unique identifier of the Shipping Method in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetShippingCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetShippingCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingCustomField", Alias: (*Alias)(&obj)})
}

type CartSetShippingCustomTypeAction struct {
	// User-defined unique identifier of the Shipping Method in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
	// Defines the [Type](ctp:api:type:Type) that extends the `shippingAddress` with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the `shippingAddress`.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `shippingAddress`.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetShippingCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetShippingCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingCustomType", Alias: (*Alias)(&obj)})
}

type CartSetShippingMethodAction struct {
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [ShippingMethod](ctp:api:type:ShippingMethod).
	ShippingMethod  *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	ExternalTaxRate *ExternalTaxRateDraft             `json:"externalTaxRate,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingMethod", Alias: (*Alias)(&obj)})
}

type CartSetShippingMethodTaxAmountAction struct {
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetShippingMethodTaxAmountAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetShippingMethodTaxAmountAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingMethodTaxAmount", Alias: (*Alias)(&obj)})
}

type CartSetShippingMethodTaxRateAction struct {
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetShippingMethodTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetShippingMethodTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingMethodTaxRate", Alias: (*Alias)(&obj)})
}

type CartSetShippingRateInputAction struct {
	// Based on the definition of ShippingRateInputType.
	// If CartClassification is defined, it must be ClassificationShippingRateInput.
	// If CartScore is defined, it must be ScoreShippingRateInput.
	// Otherwise it can not bet set.
	ShippingRateInput ShippingRateInputDraft `json:"shippingRateInput,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartSetShippingRateInputAction) UnmarshalJSON(data []byte) error {
	type Alias CartSetShippingRateInputAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInput != nil {
		var err error
		obj.ShippingRateInput, err = mapDiscriminatorShippingRateInputDraft(obj.ShippingRateInput)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetShippingRateInputAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetShippingRateInputAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingRateInput", Alias: (*Alias)(&obj)})
}

type CartUpdateItemShippingAddressAction struct {
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartUpdateItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CartUpdateItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateItemShippingAddress", Alias: (*Alias)(&obj)})
}

type CustomLineItemImportDraft struct {
	Name LocalizedString `json:"name"`
	// The amount of a CustomLineItem in the cart.
	// Must be a positive integer.
	Quantity int `json:"quantity"`
	// The cost to add to the cart. The amount can be negative.
	Money       Money                          `json:"money"`
	Slug        string                         `json:"slug"`
	State       []ItemState                    `json:"state"`
	TaxRate     *TaxRate                       `json:"taxRate,omitempty"`
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// The custom fields.
	Custom          *CustomFieldsDraft        `json:"custom,omitempty"`
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	// - If `Standard`, Cart Discounts with a matching [CartDiscountCustomLineItemsTarget](ctp:api:type:CartDiscountCustomLineItemsTarget)
	// are applied to the Custom Line Item.
	// - If `External`, Cart Discounts are not considered on the Custom Line Item.
	PriceMode CustomLineItemPriceMode `json:"priceMode"`
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
*	The scope controls which part of the product information is published.
 */
type ProductPublishScope string

const (
	ProductPublishScopeAll    ProductPublishScope = "All"
	ProductPublishScopePrices ProductPublishScope = "Prices"
)
