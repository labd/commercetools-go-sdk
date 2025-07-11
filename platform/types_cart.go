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
	// Current version of the Cart.
	Version int `json:"version"`
	// Date and time (UTC) the Cart was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Cart was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// User-defined unique identifier of the Cart.
	Key *string `json:"key,omitempty"`
	// `id` of the [Customer](ctp:api:type:Customer) that the Cart belongs to.
	CustomerId *string `json:"customerId,omitempty"`
	// Email address of the Customer that the Cart belongs to.
	CustomerEmail *string `json:"customerEmail,omitempty"`
	// [Reference](ctp:api:type:Reference) to the Customer Group of the Customer that the Cart belongs to. Used for [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection).
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// [Anonymous session](ctp:api:type:AnonymousSession) associated with the Cart.
	AnonymousId *string `json:"anonymousId,omitempty"`
	// [Reference](ctp:api:type:Reference) to a Business Unit the Cart belongs to. Only available for [B2B](/../offering/composable-commerce#composable-commerce-for-b2b)-enabled Projects.
	BusinessUnit *BusinessUnitKeyReference `json:"businessUnit,omitempty"`
	// [Reference](ctp:api:type:Reference) to a Store the Cart belongs to.
	Store *StoreKeyReference `json:"store,omitempty"`
	// [Line Items](ctp:api:type:LineItems) added to the Cart.
	LineItems []LineItem `json:"lineItems"`
	// [Custom Line Items](ctp:api:type:CustomLineItems) added to the Cart.
	CustomLineItems []CustomLineItem `json:"customLineItems"`
	// Sum of all [LineItem](ctp:api:type:LineItem) quantities, excluding [CustomLineItems](ctp:api:type:CustomLineItem). Only present when the Cart has at least one LineItem.
	TotalLineItemQuantity *int `json:"totalLineItemQuantity,omitempty"`
	// Sum of the `totalPrice` field of all [LineItems](ctp:api:type:LineItem) and [CustomLineItems](ctp:api:type:CustomLineItem), and if available, the `price` field of [ShippingInfo](ctp:api:type:ShippingInfo).
	// If a discount applies on `totalPrice`, this field holds the discounted value.
	//
	// Taxes are included if [TaxRate](ctp:api:type:TaxRate) `includedInPrice` is `true` for each price.
	TotalPrice CentPrecisionMoney `json:"totalPrice"`
	// - For a Cart with `Platform` [TaxMode](ctp:api:type:TaxMode), it is automatically set when a [shipping address is set](ctp:api:type:CartSetShippingAddressAction). For Carts with `Multiple` [ShippingMode](ctp:api:type:ShippingMode), all Line Items and Custom Line Items must be fully distributed between the Shipping Methods (via `shippingDetails`), otherwise `taxedPrice` is not automatically set.
	// - For a Cart with `External` [TaxMode](ctp:api:type:TaxMode), it is automatically set when `shippingAddress` and external Tax Rates for all Line Items, Custom Line Items, and Shipping Methods in the Cart are set. For Carts with `Multiple` [ShippingMode](ctp:api:type:ShippingMode), all allocations must have their respective tax rates present in `perMethodTaxRate`, otherwise `taxedPrice` is not automatically set.
	//
	// If a discount applies on `totalPrice`, this field holds the proportionally discounted value.
	TaxedPrice *TaxedPrice `json:"taxedPrice,omitempty"`
	// Sum of the `taxedPrice` field of [ShippingInfo](ctp:api:type:ShippingInfo) across all Shipping Methods.
	//
	// If a discount applies on `totalPrice`, this field holds the proportionally discounted value.
	TaxedShippingPrice *TaxedPrice `json:"taxedShippingPrice,omitempty"`
	// Discounts that apply on the Cart `totalPrice`.
	DiscountOnTotalPrice *DiscountOnTotalPrice `json:"discountOnTotalPrice,omitempty"`
	// Indicates how Tax Rates are set.
	TaxMode TaxMode `json:"taxMode"`
	// Indicates how the total prices on [LineItems](ctp:api:type:LineItem) and [CustomLineItems](ctp:api:type:CustomLineItem) are rounded when calculated. Configured in [Project settings](ctp:api:type:CartsConfiguration).
	PriceRoundingMode RoundingMode `json:"priceRoundingMode"`
	// Indicates how monetary values are rounded when calculating taxes for `taxedPrice`. Configured in [Project settings](ctp:api:type:CartsConfiguration).
	TaxRoundingMode RoundingMode `json:"taxRoundingMode"`
	// Indicates how taxes are calculated when calculating taxes for `taxedPrice`.
	TaxCalculationMode TaxCalculationMode `json:"taxCalculationMode"`
	// Indicates how stock quantities are tracked for Line Items in the Cart.
	InventoryMode InventoryMode `json:"inventoryMode"`
	// Current status of the Cart.
	CartState CartState `json:"cartState"`
	// Billing address associated with the Cart.
	BillingAddress *Address `json:"billingAddress,omitempty"`
	// Shipping address for a Cart with `Single` [ShippingMode](ctp:api:type:ShippingMode). Determines eligible [ShippingMethod](ctp:api:type:ShippingMethod) rates and Tax Rates of Line Items.
	ShippingAddress *Address `json:"shippingAddress,omitempty"`
	// Indicates whether the Cart has one or multiple Shipping Methods.
	ShippingMode ShippingMode `json:"shippingMode"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Single` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
	// Shipping-related information of a Cart with `Single` [ShippingMode](ctp:api:type:ShippingMode). Automatically set when a [Shipping Method is set](ctp:api:type:CartSetShippingMethodAction).
	ShippingInfo *ShippingInfo `json:"shippingInfo,omitempty"`
	// Input used to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	// The data type of this field depends on the `shippingRateInputType.type` configured in the [Project](ctp:api:type:Project):
	//
	// - If `CartClassification`, it is [ClassificationShippingRateInput](ctp:api:type:ClassificationShippingRateInput).
	// - If `CartScore`, it is [ScoreShippingRateInput](ctp:api:type:ScoreShippingRateInput).
	// - If `CartValue`, it cannot be used.
	ShippingRateInput ShippingRateInput `json:"shippingRateInput,omitempty"`
	// Custom Fields of the Shipping Method in a Cart with `Single` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingCustomFields *CustomFields `json:"shippingCustomFields,omitempty"`
	// Shipping-related information of a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode). Updated automatically each time a new [Shipping Method is added](ctp:api:type:CartAddShippingMethodAction).
	Shipping []Shipping `json:"shipping"`
	// Additional shipping addresses of the Cart as specified by [LineItems](ctp:api:type:LineItem) using the `shippingDetails` field.
	//
	// For Carts with `Single` [ShippingMode](ctp:api:type:ShippingMode): eligible Shipping Methods or applicable Tax Rates are determined by the address in `shippingAddress`, and not `itemShippingAddresses`.
	ItemShippingAddresses []Address `json:"itemShippingAddresses"`
	// Discount Codes applied to the Cart. A Cart that has `directDiscounts` cannot have `discountCodes`.
	DiscountCodes []DiscountCodeInfo `json:"discountCodes"`
	// Direct Discounts added to the Cart. A Cart that has `discountCodes` cannot have `directDiscounts`.
	DirectDiscounts []DirectDiscount `json:"directDiscounts"`
	// Automatically set when a Line Item with `GiftLineItem` [LineItemMode](ctp:api:type:LineItemMode) is [removed](ctp:api:type:CartRemoveLineItemAction) from the Cart.
	RefusedGifts []CartDiscountReference `json:"refusedGifts"`
	// Payment information related to the Cart.
	PaymentInfo *PaymentInfo `json:"paymentInfo,omitempty"`
	// Used for [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection).
	Country *string `json:"country,omitempty"`
	// Languages of the Cart. Can only contain languages supported by the [Project](ctp:api:type:Project).
	Locale *string `json:"locale,omitempty"`
	// Indicates how the Cart was created.
	Origin CartOrigin `json:"origin"`
	// Custom Fields of the Cart.
	Custom *CustomFields `json:"custom,omitempty"`
	// Indicates if a combination of discount types can apply on a Cart.
	DiscountTypeCombination DiscountTypeCombination `json:"discountTypeCombination,omitempty"`
	// Number of days after the last modification before a Cart is deleted. Configured in [Project settings](ctp:api:type:CartsConfiguration).
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
	// IDs and references that last modified the Cart.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the Cart.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Cart) UnmarshalJSON(data []byte) error {
	type Alias Cart
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
	if obj.DiscountTypeCombination != nil {
		var err error
		obj.DiscountTypeCombination, err = mapDiscriminatorDiscountTypeCombination(obj.DiscountTypeCombination)
		if err != nil {
			return err
		}
	}

	return nil
}

type CartDraft struct {
	// Currency the Cart uses.
	Currency string `json:"currency"`
	// User-defined unique identifier for the Cart.
	Key *string `json:"key,omitempty"`
	// `id` of the [Customer](ctp:api:type:Customer) that the Cart belongs to.
	CustomerId *string `json:"customerId,omitempty"`
	// Email address of the Customer that the Cart belongs to.
	CustomerEmail *string `json:"customerEmail,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to the Customer Group of the Customer that the Cart belongs to. Used for [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection).
	//
	// You can set either a `customerId` or a `customerGroup`.
	// If the Customer referenced in `customerId` belongs to a Customer Group then `customerGroup` is set automatically.
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	// [Anonymous session](ctp:api:type:AnonymousSession) associated with the Cart.
	AnonymousId *string `json:"anonymousId,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to the Business Unit the Cart should belong to. When the `customerId` of the Cart is also set, the [Customer](ctp:api:type:Customer) must be an [Associate](ctp:api:type:Associate) of the Business Unit. Only available for [B2B](/../offering/composable-commerce#composable-commerce-for-b2b)-enabled Projects.
	BusinessUnit *BusinessUnitResourceIdentifier `json:"businessUnit,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to the Store the Cart should belong to. Once set, it cannot be updated.
	Store *StoreResourceIdentifier `json:"store,omitempty"`
	// [Line Items](ctp:api:type:LineItems) to add to the Cart.
	LineItems []LineItemDraft `json:"lineItems"`
	// [Custom Line Items](ctp:api:type:CustomLineItems) to add to the Cart.
	CustomLineItems []CustomLineItemDraft `json:"customLineItems"`
	// Determines how Tax Rates are set.
	TaxMode *TaxMode `json:"taxMode,omitempty"`
	// External Tax Rate for the `shippingMethod` if the Cart has `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRateForShippingMethod *ExternalTaxRateDraft `json:"externalTaxRateForShippingMethod,omitempty"`
	// Determines how the total prices on [LineItems](ctp:api:type:LineItem) and [CustomLineItems](ctp:api:type:CustomLineItem) are rounded when calculated. If not set, the [default value](ctp:api:type:CartsConfiguration) configured in the [Project](ctp:api:type:Project) is used.
	PriceRoundingMode *RoundingMode `json:"priceRoundingMode,omitempty"`
	// Determines how monetary values are rounded when calculating taxes for `taxedPrice`. If not set, the [default value](ctp:api:type:CartsConfiguration) configured in the [Project](ctp:api:type:Project) is used.
	TaxRoundingMode *RoundingMode `json:"taxRoundingMode,omitempty"`
	// Determines how taxes are calculated for `taxedPrice`.
	TaxCalculationMode *TaxCalculationMode `json:"taxCalculationMode,omitempty"`
	// Determines how stock quantities are tracked for Line Items in the Cart.
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	// Billing address associated with the Cart.
	BillingAddress *BaseAddress `json:"billingAddress,omitempty"`
	// Shipping address for a Cart with `Single` [ShippingMode](ctp:api:type:ShippingMode). Determines eligible [ShippingMethod](ctp:api:type:ShippingMethod) rates and Tax Rates of Line Items.
	// Must be one of the `itemShippingAddresses` when that field is also provided.
	ShippingAddress *BaseAddress `json:"shippingAddress,omitempty"`
	// Shipping Method for a Cart with `Single` [ShippingMode](ctp:api:type:ShippingMode). If the referenced [ShippingMethod](ctp:api:type:ShippingMethod) has a `predicate` that does not match the Cart, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned when [creating a Cart](ctp:api:endpoint:/{projectKey}/carts:POST).
	ShippingMethod *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	// Used as an input to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	// The data type of this field depends on the `shippingRateInputType.type` configured in the [Project](ctp:api:type:Project):
	//
	// - If `CartClassification`, it must be [ClassificationShippingRateInputDraft](ctp:api:type:ClassificationShippingRateInputDraft).
	// - If `CartScore`, it must be [ScoreShippingRateInputDraft](ctp:api:type:ScoreShippingRateInputDraft).
	// - If `CartValue`, it cannot be set.
	ShippingRateInput ShippingRateInputDraft `json:"shippingRateInput,omitempty"`
	// - If set to `Single`, only a single Shipping Method can be added to the Cart.
	// - If set to `Multiple`, multiple Shipping Methods can be added to the Cart.
	ShippingMode *ShippingMode `json:"shippingMode,omitempty"`
	// Custom Shipping Methods for a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	CustomShipping []CustomShippingDraft `json:"customShipping"`
	// Shipping Methods for a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	Shipping []ShippingDraft `json:"shipping"`
	// Multiple shipping addresses of the Cart. Each address must contain a `key` that is unique in this Cart.
	// The keys are used by [LineItems](ctp:api:type:LineItem) to reference these addresses under their `shippingDetails`.
	//
	// For Carts with `Single` [ShippingMode](ctp:api:type:ShippingMode): eligible Shipping Methods or applicable Tax Rates are determined by the address `shippingAddress`, and not `itemShippingAddresses`.
	ItemShippingAddresses []BaseAddress `json:"itemShippingAddresses"`
	// `code` of the existing [DiscountCodes](ctp:api:type:DiscountCode) to add to the Cart.
	DiscountCodes []string `json:"discountCodes"`
	// Used for [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection).
	// If used for [Create Cart in Store](ctp:api:endpoint:/{projectKey}/in-store/carts:POST), the provided country must be one of the [Store's](ctp:api:type:Store) `countries`.
	Country *string `json:"country,omitempty"`
	// Languages of the Cart. Can only contain languages supported by the [Project](ctp:api:type:Project).
	Locale *string `json:"locale,omitempty"`
	// Indicates how the Cart was created.
	Origin *CartOrigin `json:"origin,omitempty"`
	// Number of days after the last modification before a Cart is deleted.
	// If not provided, the default value for this field configured in [Project settings](ctp:api:type:CartsConfiguration) is assigned.
	//
	// Create a [ChangeSubscription](ctp:api:type:ChangeSubscription) for Carts to receive a [ResourceDeletedDeliveryPayload](ctp:api:type:ResourceDeletedDeliveryPayload) upon deletion of the Cart.
	DeleteDaysAfterLastModification *int `json:"deleteDaysAfterLastModification,omitempty"`
	// Custom Fields for the Cart.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
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

/**
*	Indicates who created the Cart.
*
 */
type CartOrigin string

const (
	CartOriginCustomer       CartOrigin = "Customer"
	CartOriginMerchant       CartOrigin = "Merchant"
	CartOriginQuote          CartOrigin = "Quote"
	CartOriginRecurringOrder CartOrigin = "RecurringOrder"
)

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [Cart](ctp:api:type:Cart).
*
 */
type CartPagedQueryResponse struct {
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
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [Carts](ctp:api:type:Cart) matching the query.
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
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Cart](ctp:api:type:Cart). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type CartResourceIdentifier struct {
	// Unique identifier of the referenced [Cart](ctp:api:type:Cart). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [Cart](ctp:api:type:Cart). Required if `id` is absent.
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

/**
*	Indicates the current status of a Cart.
*
 */
type CartState string

const (
	CartStateActive  CartState = "Active"
	CartStateMerged  CartState = "Merged"
	CartStateOrdered CartState = "Ordered"
	CartStateFrozen  CartState = "Frozen"
)

type CartUpdate struct {
	// Expected version of the Cart on which the changes apply.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the Cart.
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
	case "changeLineItemsOrder":
		obj := CartChangeLineItemsOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changePriceRoundingMode":
		obj := CartChangePriceRoundingModeAction{}
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
	case "freezeCart":
		obj := CartFreezeCartAction{}
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
	case "setBusinessUnit":
		obj := CartSetBusinessUnitAction{}
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
	case "setCustomLineItemRecurrenceInfo":
		obj := CartSetCustomLineItemRecurrenceInfoAction{}
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
	case "setLineItemInventoryMode":
		obj := CartSetLineItemInventoryModeAction{}
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
	case "setLineItemRecurrenceInfo":
		obj := CartSetLineItemRecurrenceInfoAction{}
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
	case "unfreezeCart":
		obj := CartUnfreezeCartAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
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

/**
*	A generic item that can be added to the Cart but is not bound to a Product that can be used for discounts (negative money), vouchers, complex cart rules, additional services, or fees.
*	You control the lifecycle of this item.
*
 */
type CustomLineItem struct {
	// Unique identifier of the Custom Line Item.
	ID string `json:"id"`
	// User-defined unique identifier of the Custom Line Item.
	Key *string `json:"key,omitempty"`
	// Name of the Custom Line Item.
	Name LocalizedString `json:"name"`
	// Money value of the Custom Line Item.
	Money TypedMoney `json:"money"`
	// Automatically set after the `taxRate` is set.
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
	// Total taxed prices based on the quantity of the Custom Line Item assigned to each [Shipping Method](ctp:api:type:ShippingMethod). Only applicable for Carts with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	// Automatically set after `perMethodTaxRate` is set.
	TaxedPricePortions []MethodTaxedPrice `json:"taxedPricePortions"`
	// Total price of the Custom Line Item (`money` multiplied by `quantity`).
	// If the Custom Line Item is discounted, the total price is `discountedPricePerQuantity` multiplied by `quantity`.
	//
	// Includes taxes if the [TaxRate](ctp:api:type:TaxRate) `includedInPrice` is `true`.
	TotalPrice CentPrecisionMoney `json:"totalPrice"`
	// User-defined identifier used in a deep-link URL for the Custom Line Item.
	// It matches the pattern `[a-zA-Z0-9_-]{2,256}`.
	Slug string `json:"slug"`
	// Number of Custom Line Items in the [Cart](ctp:api:type:Cart) or [Order](ctp:api:type:Order).
	Quantity int `json:"quantity"`
	// Tracks specific quantities of the Custom Line Item within a given State. When a Custom Line Item is added to a Cart, its full quantity is set to the built-in "Initial" state. State transitions for Custom Line Items are managed on the [Order](ctp:api:type:Order).
	State []ItemState `json:"state"`
	// Used to select a Tax Rate when a Cart has the `Platform` [TaxMode](ctp:api:type:TaxMode).
	TaxCategory *TaxCategoryReference `json:"taxCategory,omitempty"`
	// - For a Cart with `Platform` [TaxMode](ctp:api:type:TaxMode), the `taxRate` of Custom Line Items is set automatically once a shipping address is set. The rate is based on the [TaxCategory](ctp:api:type:TaxCategory) that applies for the shipping address.
	// - For a Cart with `External` TaxMode, the `taxRate` of Custom Line Items can be set using [ExternalTaxRateDraft](ctp:api:type:ExternalTaxRateDraft).
	TaxRate *TaxRate `json:"taxRate,omitempty"`
	// Tax Rate per Shipping Method for a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode). For a Cart with `Platform` [TaxMode](ctp:api:type:TaxMode) it is automatically set after the [Shipping Method is added](ctp:api:type:CartAddShippingMethodAction).
	// For a Cart with `External` [TaxMode](ctp:api:type:TaxMode), the Tax Rate must be set with [ExternalTaxRateDraft](ctp:api:type:ExternalTaxRateDraft).
	PerMethodTaxRate []MethodTaxRate `json:"perMethodTaxRate"`
	// Discounted price of a single quantity of the Custom Line Item.
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	// Custom Fields of the Custom Line Item.
	Custom *CustomFields `json:"custom,omitempty"`
	// Container for Custom Line Item-specific addresses.
	ShippingDetails *ItemShippingDetails `json:"shippingDetails,omitempty"`
	// Indicates whether Cart Discounts with a matching [CartDiscountCustomLineItemsTarget](ctp:api:type:CartDiscountCustomLineItemsTarget), [MultiBuyCustomLineItemsTarget](ctp:api:type:MultiBuyCustomLineItemsTarget), or [CartDiscountPatternTarget](ctp:api:type:CartDiscountPatternTarget) are applied to the Custom Line Item.
	PriceMode CustomLineItemPriceMode `json:"priceMode"`
	// Recurring Order and frequency data.
	RecurrenceInfo *CustomLineItemRecurrenceInfo `json:"recurrenceInfo,omitempty"`
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

	return nil
}

type CustomLineItemDraft struct {
	// Name of the Custom Line Item.
	Name LocalizedString `json:"name"`
	// User-defined unique identifier of the Custom Line Item.
	Key *string `json:"key,omitempty"`
	// Number of Custom Line Items to add to the Cart.
	Quantity *int `json:"quantity,omitempty"`
	// Money value of the Custom Line Item.
	// The value can be negative.
	Money Money `json:"money"`
	// User-defined identifier used in a deep-link URL for the Custom Line Item.
	// It must match the pattern `[a-zA-Z0-9_-]{2,256}`.
	Slug string `json:"slug"`
	// Used to select a Tax Rate when a Cart has the `Platform` [TaxMode](ctp:api:type:TaxMode).
	// This field is required for `Platform` [TaxMode](ctp:api:type:TaxMode).
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// External Tax Rate for the Custom Line Item if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// Custom Fields for the Custom Line Item.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Container for Custom Line Item-specific addresses.
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	// - If `Standard`, Cart Discounts with a matching [CartDiscountCustomLineItemsTarget](ctp:api:type:CartDiscountCustomLineItemsTarget), [MultiBuyCustomLineItemsTarget](ctp:api:type:MultiBuyCustomLineItemsTarget), or [CartDiscountPatternTarget](ctp:api:type:CartDiscountPatternTarget) are applied to the Custom Line Item.
	// - If `External`, Cart Discounts are not considered on the Custom Line Item.
	PriceMode *CustomLineItemPriceMode `json:"priceMode,omitempty"`
	// Recurring Order and frequency data.
	RecurrenceInfo *CustomLineItemRecurrenceInfoDraft `json:"recurrenceInfo,omitempty"`
}

/**
*	Determines if Cart Discounts can be applied to a Custom Line Item in the Cart.
*
 */
type CustomLineItemPriceMode string

const (
	CustomLineItemPriceModeStandard CustomLineItemPriceMode = "Standard"
	CustomLineItemPriceModeExternal CustomLineItemPriceMode = "External"
)

type CustomShippingDraft struct {
	// User-defined unique identifier of the custom Shipping Method in the Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	Key string `json:"key"`
	// Name of the custom Shipping Method.
	ShippingMethodName string `json:"shippingMethodName"`
	// Determines the shipping rate and Tax Rate of the associated Line Items.
	ShippingAddress *BaseAddress `json:"shippingAddress,omitempty"`
	// Determines the shipping price.
	ShippingRate ShippingRateDraft `json:"shippingRate"`
	// Input used to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	// The data type of this field depends on the `shippingRateInputType.type` configured in the [Project](ctp:api:type:Project):
	//
	// - If `CartClassification`, it must be [ClassificationShippingRateInputDraft](ctp:api:type:ClassificationShippingRateInputDraft).
	// - If `CartScore`, it must be [ScoreShippingRateInputDraft](ctp:api:type:ScoreShippingRateInputDraft).
	// - If `CartValue`, it cannot be set.
	ShippingRateInput ShippingRateInputDraft `json:"shippingRateInput,omitempty"`
	// Tax Category used to determine a shipping Tax Rate if the Cart has the `Platform` [TaxMode](ctp:api:type:TaxMode).
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// Tax Rate used to tax a shipping expense if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// Deliveries to be shipped with the custom Shipping Method.
	Deliveries []DeliveryDraft `json:"deliveries"`
	// Custom Fields for the custom Shipping Method.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomShippingDraft) MarshalJSON() ([]byte, error) {
	type Alias CustomShippingDraft
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
*	Represents a [CartDiscount](ctp:api:type:CartDiscount) that is only associated with a single Cart or Order.
*
 */
type DirectDiscount struct {
	// Unique identifier of the Direct Discount.
	ID string `json:"id"`
	// Effect of the Discount on the Cart.
	Value CartDiscountValue `json:"value"`
	// Segment of the Cart that is discounted.
	//
	// Empty when the `value` is set to `giftLineItem`.
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

/**
*	Represents a [CartDiscount](ctp:api:type:CartDiscount) that can only be associated with a single Cart or Order.
*
*	For an introduction to Direct Discounts and to understand how they work in Composable Commerce, see the [Direct Discounts overview](/pricing-and-discounts-overview#direct-discounts).
*
 */
type DirectDiscountDraft struct {
	// Defines the effect the Discount will have.
	Value CartDiscountValueDraft `json:"value"`
	// Defines what segment of the Cart will be discounted.
	//
	// If `value` is set to `giftLineItem`, this must not be set.
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
		obj.Value, err = mapDiscriminatorCartDiscountValueDraft(obj.Value)
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

/**
*	[Reference](ctp:api:type:Reference) to a [DirectDiscount](ctp:api:type:DirectDiscount).
*
 */
type DirectDiscountReference struct {
	// Unique identifier of the referenced [DirectDiscount](ctp:api:type:DirectDiscount).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DirectDiscountReference) MarshalJSON() ([]byte, error) {
	type Alias DirectDiscountReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "direct-discount", Alias: (*Alias)(&obj)})
}

type DiscountCodeInfo struct {
	// Discount Code associated with the Cart or Order.
	DiscountCode DiscountCodeReference `json:"discountCode"`
	// Indicates the state of the Discount Code applied to the Cart or Order.
	State DiscountCodeState `json:"state"`
}

/**
*	Indicates the state of a Discount Code in a Cart.
*
*	If an Order is created from a Cart with a state other than `MatchesCart` or `ApplicationStoppedByGroupBestDeal`, a [DiscountCodeNonApplicable](ctp:api:type:DiscountCodeNonApplicableError) error is returned.
*
*	For Orders created from a Cart with a `ApplicationStoppedByGroupBestDeal` state, the discount code is not applied.
*
 */
type DiscountCodeState string

const (
	DiscountCodeStateNotActive                            DiscountCodeState = "NotActive"
	DiscountCodeStateNotValid                             DiscountCodeState = "NotValid"
	DiscountCodeStateDoesNotMatchCart                     DiscountCodeState = "DoesNotMatchCart"
	DiscountCodeStateMatchesCart                          DiscountCodeState = "MatchesCart"
	DiscountCodeStateMaxApplicationReached                DiscountCodeState = "MaxApplicationReached"
	DiscountCodeStateApplicationStoppedByPreviousDiscount DiscountCodeState = "ApplicationStoppedByPreviousDiscount"
	DiscountCodeStateApplicationStoppedByGroupBestDeal    DiscountCodeState = "ApplicationStoppedByGroupBestDeal"
)

type DiscountOnTotalPrice struct {
	// Money value of the discount on the total price of the Cart or Order.
	DiscountedAmount TypedMoney `json:"discountedAmount"`
	// Discounts that impact the total price of the Cart or Order.
	IncludedDiscounts []DiscountedTotalPricePortion `json:"includedDiscounts"`
	// Money value of the discount on the total net price of the Cart or Order.
	//
	// The same percentage of discount applies as on the `discountedAmount`.
	// Present only when `taxedPrice` of the Cart or Order exists.
	DiscountedNetAmount TypedMoney `json:"discountedNetAmount,omitempty"`
	// Money value of the discount on the total gross price of the Cart or Order.
	//
	// The same percentage of discount applies as on the `discountedAmount`.
	// Present only when `taxedPrice` of the Cart or Order exists.
	DiscountedGrossAmount TypedMoney `json:"discountedGrossAmount,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountOnTotalPrice) UnmarshalJSON(data []byte) error {
	type Alias DiscountOnTotalPrice
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
	if obj.DiscountedNetAmount != nil {
		var err error
		obj.DiscountedNetAmount, err = mapDiscriminatorTypedMoney(obj.DiscountedNetAmount)
		if err != nil {
			return err
		}
	}
	if obj.DiscountedGrossAmount != nil {
		var err error
		obj.DiscountedGrossAmount, err = mapDiscriminatorTypedMoney(obj.DiscountedGrossAmount)
		if err != nil {
			return err
		}
	}

	return nil
}

type DiscountTypeCombination interface{}

func mapDiscriminatorDiscountTypeCombination(input interface{}) (DiscountTypeCombination, error) {
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
	case "BestDeal":
		obj := BestDeal{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "Stacking":
		obj := Stacking{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Indicates the best deal logic applies to a Cart or Order and indicates the discount type that offers the best deal.
*
 */
type BestDeal struct {
	// Discount type that offers the best deal; the value can be `product-discount` or `cart-discount`.
	ChosenDiscountType string `json:"chosenDiscountType"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj BestDeal) MarshalJSON() ([]byte, error) {
	type Alias BestDeal
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "BestDeal", Alias: (*Alias)(&obj)})
}

type DiscountedLineItemPortion struct {
	// A [CartDiscountReference](ctp:api:type:CartDiscountReference) or [DirectDiscountReference](ctp:api:type:DirectDiscountReference) of the applicable discount on the Line Item.
	Discount Reference `json:"discount"`
	// Money value of the applicable discount.
	DiscountedAmount TypedMoney `json:"discountedAmount"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountedLineItemPortion) UnmarshalJSON(data []byte) error {
	type Alias DiscountedLineItemPortion
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Discount != nil {
		var err error
		obj.Discount, err = mapDiscriminatorReference(obj.Discount)
		if err != nil {
			return err
		}
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

type DiscountedLineItemPortionDraft struct {
	// A [CartDiscountReference](ctp:api:type:CartDiscountReference) or [DirectDiscountReference](ctp:api:type:DirectDiscountReference) for the discount applicable on the Line Item.
	Discount Reference `json:"discount"`
	// Money value for the discount applicable.
	DiscountedAmount TypedMoneyDraft `json:"discountedAmount"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountedLineItemPortionDraft) UnmarshalJSON(data []byte) error {
	type Alias DiscountedLineItemPortionDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Discount != nil {
		var err error
		obj.Discount, err = mapDiscriminatorReference(obj.Discount)
		if err != nil {
			return err
		}
	}
	if obj.DiscountedAmount != nil {
		var err error
		obj.DiscountedAmount, err = mapDiscriminatorTypedMoneyDraft(obj.DiscountedAmount)
		if err != nil {
			return err
		}
	}

	return nil
}

type DiscountedLineItemPrice struct {
	// Money value of the discounted Line Item or Custom Line Item.
	Value TypedMoney `json:"value"`
	// Discount applicable on the Line Item or Custom Line Item.
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
	// Number of Line Items or Custom Line Items in the Cart.
	Quantity int `json:"quantity"`
	// Discounted price of the Line Item or Custom Line Item.
	DiscountedPrice DiscountedLineItemPrice `json:"discountedPrice"`
}

type DiscountedTotalPricePortion struct {
	// A [CartDiscountReference](ctp:api:type:CartDiscountReference) or [DirectDiscountReference](ctp:api:type:DirectDiscountReference) to the discount applied on the Cart `totalPrice`.
	Discount Reference `json:"discount"`
	// Money value of the discount.
	DiscountedAmount TypedMoney `json:"discountedAmount"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *DiscountedTotalPricePortion) UnmarshalJSON(data []byte) error {
	type Alias DiscountedTotalPricePortion
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Discount != nil {
		var err error
		obj.Discount, err = mapDiscriminatorReference(obj.Discount)
		if err != nil {
			return err
		}
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

type ExternalLineItemTotalPrice struct {
	// Price of the Line Item.
	Price Money `json:"price"`
	// Total price of the Line Item.
	TotalPrice Money `json:"totalPrice"`
}

/**
*	Cannot be used in [LineItemDraft](ctp:api:type:LineItemDraft) or [CustomLineItemDraft](ctp:api:type:CustomLineItemDraft).
*
*	Can only be set by these update actions:
*
*	- [Set LineItem TaxAmount](ctp:api:type:CartSetLineItemTaxAmountAction), [Set CustomLineItem TaxAmount](ctp:api:type:CartSetCustomLineItemTaxAmountAction), or [Set ShippingMethod TaxAmount](ctp:api:type:CartSetShippingMethodTaxAmountAction) on Carts
*	- [Set LineItem TaxAmount](ctp:api:type:StagedOrderSetLineItemTaxAmountAction), [Set CustomLineItem TaxAmount](ctp:api:type:StagedOrderSetCustomLineItemTaxAmountAction), or [Set ShippingMethod TaxAmount](ctp:api:type:StagedOrderSetShippingMethodTaxAmountAction) on Order Edits
*
 */
type ExternalTaxAmountDraft struct {
	// Total gross amount (`totalNet` + `taxPortions`) of the Line Item or Custom Line Item.
	TotalGross Money `json:"totalGross"`
	// Tax Rates and subrates of states and countries.
	TaxRate ExternalTaxRateDraft `json:"taxRate"`
}

/**
*	Controls calculation of taxed prices for Line Items, Custom Line Items, and Shipping Methods as explained in [Cart tax calculation](ctp:api:type:CartTaxCalculation).
*
 */
type ExternalTaxRateDraft struct {
	// Name of the Tax Rate.
	Name string `json:"name"`
	// Percentage in the range of 0-1.
	//
	// - If no `subRates` are specified, a value must be defined.
	// - If `subRates` are specified, this can be omitted or its value must be the sum of all `subRates` amounts.
	Amount *float64 `json:"amount,omitempty"`
	// - If set to `false`, the related price is considered the net price and the provided `amount` is applied to calculate the gross price.
	// - If set to `true`, the related price is considered the gross price, and the provided `amount` is applied to calculate the net price.
	IncludedInPrice *bool `json:"includedInPrice,omitempty"`
	// Country for which the tax applies.
	Country string `json:"country"`
	// State within the specified country.
	State *string `json:"state,omitempty"`
	// Used when the total tax is a combination of multiple taxes (for example, local, state/provincial, and/or federal taxes). The total of all subrates must equal the TaxRate `amount`.
	// These subrates are used to calculate the `taxPortions` field of a [Cart](ctp:api:type:Cart) or [Order](ctp:api:type:Order) and the `taxedPrice` field of [LineItems](ctp:api:type:LineItem), [CustomLineItems](ctp:api:type:CustomLineItem), and [ShippingInfos](ctp:api:type:ShippingInfo).
	SubRates []SubRate `json:"subRates"`
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

/**
*	Indicates how Line Items in a Cart are tracked.
*
 */
type InventoryMode string

const (
	InventoryModeNone           InventoryMode = "None"
	InventoryModeTrackOnly      InventoryMode = "TrackOnly"
	InventoryModeReserveOnOrder InventoryMode = "ReserveOnOrder"
)

type ItemShippingDetails struct {
	// Holds information on the quantity of Line Items or Custom Line Items and the address it is shipped.
	Targets []ItemShippingTarget `json:"targets"`
	// - `true` if the quantity of Line Items or Custom Line Items is equal to the sum of sub-quantities defined in `targets`.
	// - `false` if the quantity of Line Items or Custom Line Items is not equal to the sum of sub-quantities defined in `targets`.
	//   Ordering a Cart when the value is `false` returns an [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError) error.
	Valid bool `json:"valid"`
}

/**
*	For order creation and updates, the sum of the `targets` must match the quantity of the Line Items or Custom Line Items.
*
 */
type ItemShippingDetailsDraft struct {
	// Holds information on the quantity of Line Items or Custom Line Items and the address it is shipped.
	//
	// If multiple shipping addresses are present for a Line Item or Custom Line Item, sub-quantities must be specified.
	Targets []ItemShippingTarget `json:"targets"`
}

/**
*	Determines the address (as a reference to an address in `itemShippingAddresses`) and the quantity shipped to the address.
*
*	If multiple shipping addresses are present for a Line Item or Custom Line Item, sub-quantities must be specified.
*	An array of addresses and sub-quantities is stored per Line Item or Custom Line Item.
*
 */
type ItemShippingTarget struct {
	// Key of the address in the [Cart](ctp:api:type:Cart) `itemShippingAddresses`.
	// Duplicate address keys are not allowed.
	AddressKey string `json:"addressKey"`
	// Quantity of Line Items or Custom Line Items shipped to the address with the specified `addressKey`.
	//
	// If a quantity is updated to `0` when defining [ItemShippingDetailsDraft](ctp:api:type:ItemShippingDetailsDraft), the `targets` are removed from a Line Item or Custom Line Item in the resulting [ItemShippingDetails](ctp:api:type:ItemShippingDetails).
	Quantity int `json:"quantity"`
	// User-defined unique identifier of the Shipping Method in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	//
	// It connects Line Item or Custom Line Item quantities with individual Shipping Methods.
	ShippingMethodKey *string `json:"shippingMethodKey,omitempty"`
}

/**
*	The representation of a [Line Item](/../api/carts-orders-overview#line-items) in a [Cart](ctp:api:type:Cart) or in an [Order](ctp:api:type:Order).
*
 */
type LineItem struct {
	// Unique identifier of the LineItem.
	ID string `json:"id"`
	// User-defined unique identifier of the LineItem.
	Key *string `json:"key,omitempty"`
	// `id` of the [Product](ctp:api:type:Product) the Line Item is based on.
	ProductId string `json:"productId"`
	// `key` of the [Product](ctp:api:type:Product).
	//
	// This field is only present on:
	//
	// - Line Items in a [Cart](ctp:api:type:Cart) when the `key` is available on that specific Product at the time the LineItem was created or updated on the Cart.
	// - Line Items in an [Order](ctp:api:type:Order) when the `key` is available on the specific Product at the time the Order was created from the Cart.
	//
	// Present on resources created or updated after 3 December 2021.
	ProductKey *string `json:"productKey,omitempty"`
	// Name of the Product.
	Name LocalizedString `json:"name"`
	// `slug` of the current version of the Product. Updated automatically if the `slug` changes. Empty if the Product has been deleted.
	// The `productSlug` field of LineItem is not expanded when using [Reference Expansion](/../api/general-concepts#reference-expansion).
	ProductSlug *LocalizedString `json:"productSlug,omitempty"`
	// Product Type of the Product.
	ProductType ProductTypeReference `json:"productType"`
	// Holds the data of the Product Variant added to the Cart.
	//
	// The data is saved at the time the Product Variant is added to the Cart and is not updated automatically when Product Variant data changes.
	// Must be updated using the [Recalculate](ctp:api:type:CartRecalculateAction) update action.
	Variant ProductVariant `json:"variant"`
	// Price of a Line Item selected from the Product Variant according to the [Product](ctp:api:type:Product) `priceMode`. If the `priceMode` is `Embedded` [ProductPriceMode](ctp:api:type:ProductPriceModeEnum) and the `variant` field hasn't been updated, the price may not correspond to a price in `variant.prices`.
	Price Price `json:"price"`
	// Number of Line Items of the given Product Variant present in the [Cart](ctp:api:type:Cart) or [Order](ctp:api:type:Order).
	Quantity int `json:"quantity"`
	// Total price of this Line Item equalling `price` multiplied by `quantity`. If the Line Item is discounted, the total price is the `discountedPricePerQuantity` multiplied by `quantity`.
	// Includes taxes if the [TaxRate](ctp:api:type:TaxRate) `includedInPrice` is `true`.
	//
	// If `ExternalPrice` [LineItemPriceMode](#ctp:api:type:LineItemPriceMode) is used with high-precision money, then the total price is rounded by using the `HalfEven` rounding mode.
	TotalPrice CentPrecisionMoney `json:"totalPrice"`
	// Discounted price of a single quantity of the Line Item.
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	// Automatically set after `taxRate` is set.
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
	// Total taxed prices based on the quantity of Line Item assigned to each [Shipping Method](ctp:api:type:ShippingMethod). Only applicable for Carts with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	// Automatically set after `perMethodTaxRate` is set.
	TaxedPricePortions []MethodTaxedPrice `json:"taxedPricePortions"`
	// Tracks specific quantities of the Line Item within a given State. When a Line Item is added to a Cart, its full quantity is set to the built-in "Initial" state. State transitions for Line Items are managed on the [Order](ctp:api:type:Order).
	State []ItemState `json:"state"`
	// - For a Cart with `Platform` [TaxMode](ctp:api:type:TaxMode), the `taxRate` of Line Items is set automatically once a shipping address is set. The rate is based on the [TaxCategory](ctp:api:type:TaxCategory) that applies for the shipping address.
	// - For a Cart with `External` TaxMode, the `taxRate` of Line Items can be set using [ExternalTaxRateDraft](ctp:api:type:ExternalTaxRateDraft).
	TaxRate *TaxRate `json:"taxRate,omitempty"`
	// Tax Rate per Shipping Method for a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode). For a Cart with `Platform` [TaxMode](ctp:api:type:TaxMode) it is automatically set after the [Shipping Method is added](ctp:api:type:CartAddShippingMethodAction).
	// For a Cart with `External` [TaxMode](ctp:api:type:TaxMode), the Tax Rate must be set with [ExternalTaxRateDraft](ctp:api:type:ExternalTaxRateDraft).
	PerMethodTaxRate []MethodTaxRate `json:"perMethodTaxRate"`
	// Identifies [Inventory entries](/../api/projects/inventory) that are reserved. The referenced Channel has the `InventorySupply` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum).
	SupplyChannel *ChannelReference `json:"supplyChannel,omitempty"`
	// Used to [select](/../api/pricing-and-discounts-overview#line-item-price-selection) a Product Price. The referenced Channel has the `ProductDistribution` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum).
	DistributionChannel *ChannelReference `json:"distributionChannel,omitempty"`
	// Indicates how the Price for the Line Item is set.
	PriceMode LineItemPriceMode `json:"priceMode"`
	// Indicates how the Line Item is added to the Cart.
	LineItemMode LineItemMode `json:"lineItemMode"`
	// Inventory mode specific to this Line Item only, and valid for the entire `quantity` of the Line Item.
	// Only present if the inventory mode is different from the `inventoryMode` specified on the [Cart](ctp:api:type:Cart).
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	// Container for Line Item-specific addresses.
	ShippingDetails *ItemShippingDetails `json:"shippingDetails,omitempty"`
	// Custom Fields of the Line Item.
	Custom *CustomFields `json:"custom,omitempty"`
	// Date and time (UTC) the Line Item was added to the Cart.
	AddedAt *time.Time `json:"addedAt,omitempty"`
	// Date and time (UTC) the Line Item was last updated.
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`
	// Recurring Order and frequency data.
	RecurrenceInfo *LineItemRecurrenceInfo `json:"recurrenceInfo,omitempty"`
}

/**
*	For Product Variant identification, either the `productId` and `variantId`, or `sku` must be provided.
*	Product Attributes are merged with Variant Attributes to ensure the full Attribute context of the Product Variant.
*
 */
type LineItemDraft struct {
	// User-defined unique identifier of the LineItem.
	Key *string `json:"key,omitempty"`
	// `id` of a published [Product](ctp:api:type:Product).
	ProductId *string `json:"productId,omitempty"`
	// `id` of the [ProductVariant](ctp:api:type:ProductVariant) in the Product.
	// If not provided, the Master Variant is used.
	VariantId *int `json:"variantId,omitempty"`
	// `sku` of the [ProductVariant](ctp:api:type:ProductVariant).
	Sku *string `json:"sku,omitempty"`
	// Quantity of the Product Variant to add to the Cart.
	Quantity *int `json:"quantity,omitempty"`
	// Date and time (UTC) the Product Variant is added to the Cart.
	// If not set, it defaults to the current date and time.
	//
	// Optional for backwards compatibility reasons.
	AddedAt *time.Time `json:"addedAt,omitempty"`
	// Used to [select](/../api/pricing-and-discounts-overview#line-item-price-selection) a Product Price.
	// The referenced Channel must have the `ProductDistribution` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum).
	//
	// If the Cart is bound to a [Store](ctp:api:type:Store) with `distributionChannels` set,
	// the Channel must match one of the Store's distribution channels.
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
	// Used to identify [Inventory entries](/../api/projects/inventory) that must be reserved.
	// The referenced Channel must have the `InventorySupply` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum).
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	// Sets the [LineItem](ctp:api:type:LineItem) `price` value, and the `priceMode` to `ExternalPrice` [LineItemPriceMode](ctp:api:type:LineItemPriceMode).
	ExternalPrice *Money `json:"externalPrice,omitempty"`
	// Sets the [LineItem](ctp:api:type:LineItem) `price` and `totalPrice` values, and the `priceMode` to `ExternalTotal` [LineItemPriceMode](ctp:api:type:LineItemPriceMode).
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	// Sets the external Tax Rate for the Line Item, if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode) and `Single` [ShippingMode](ctp:api:type:ShippingMode). If the Cart has `Multiple` [ShippingMode](ctp:api:type:ShippingMode), the Tax Rate is accepted but ignored.
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// Sets the external Tax Rates for individual Shipping Methods, if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode) and `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	PerMethodExternalTaxRate []MethodExternalTaxRateDraft `json:"perMethodExternalTaxRate"`
	// Inventory mode specific to the Line Item only, and valid for the entire `quantity` of the Line Item.
	// Set only if the inventory mode should be different from the `inventoryMode` specified on the [Cart](ctp:api:type:Cart).
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	// Container for Line Item-specific addresses.
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	// Custom Fields for the Line Item.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Recurring Order and frequency data.
	RecurrenceInfo *LineItemRecurrenceInfoDraft `json:"recurrenceInfo,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj LineItemDraft) MarshalJSON() ([]byte, error) {
	type Alias LineItemDraft
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

	if raw["perMethodExternalTaxRate"] == nil {
		delete(raw, "perMethodExternalTaxRate")
	}

	return json.Marshal(raw)

}

/**
*	Indicates how a Line Item was added to a Cart.
*
 */
type LineItemMode string

const (
	LineItemModeStandard     LineItemMode = "Standard"
	LineItemModeGiftLineItem LineItemMode = "GiftLineItem"
)

/**
*	This mode indicates how the price is set for the Line Item.
*
 */
type LineItemPriceMode string

const (
	LineItemPriceModePlatform      LineItemPriceMode = "Platform"
	LineItemPriceModeExternalPrice LineItemPriceMode = "ExternalPrice"
	LineItemPriceModeExternalTotal LineItemPriceMode = "ExternalTotal"
)

type MethodExternalTaxRateDraft struct {
	// User-defined unique identifier of the Shipping Method in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingMethodKey string `json:"shippingMethodKey"`
	// External Tax Rate for the Shipping Method, if the Cart has `External` [TaxMode](ctp:api:type:TaxMode).
	TaxRate *ExternalTaxRateDraft `json:"taxRate,omitempty"`
}

type MethodTaxRate struct {
	// User-defined unique identifier of the Shipping Method in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingMethodKey string `json:"shippingMethodKey"`
	// Tax Rate for the Shipping Method.
	TaxRate *TaxRate `json:"taxRate,omitempty"`
}

type MethodTaxedPrice struct {
	// User-defined unique identifier of the [Shipping Method](ctp:api:type:ShippingMethod) in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingMethodKey string `json:"shippingMethodKey"`
	// Total taxed price based on the quantity of the Line Item or Custom Line Item assigned to the Shipping Method identified by `shippingMethodKey`.
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
}

/**
*	Used for [replicating an existing Cart](ctp:api:endpoint:/{projectKey}/carts/replicate:POST) or Order.
*
 */
type ReplicaCartDraft struct {
	// [Reference](ctp:api:type:Reference) to a [Cart](ctp:api:type:Cart) or [Order](ctp:api:type:Order) that is replicated.
	Reference interface{} `json:"reference"`
	// User-defined unique identifier for the Cart.
	Key *string `json:"key,omitempty"`
}

/**
*	Determines how monetary values are rounded.
*
 */
type RoundingMode string

const (
	RoundingModeHalfEven RoundingMode = "HalfEven"
	RoundingModeHalfUp   RoundingMode = "HalfUp"
	RoundingModeHalfDown RoundingMode = "HalfDown"
)

type Shipping struct {
	// User-defined unique identifier of the Shipping in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey string `json:"shippingKey"`
	// Automatically set when the [Shipping Method is added](ctp:api:type:CartAddShippingMethodAction).
	ShippingInfo ShippingInfo `json:"shippingInfo"`
	// Determines the shipping rates and Tax Rates of associated Line Items.
	ShippingAddress Address `json:"shippingAddress"`
	// Used as an input to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	// The data type of this field depends on the `shippingRateInputType.type` configured in the [Project](ctp:api:type:Project):
	//
	// - If `CartClassification`, it is [ClassificationShippingRateInput](ctp:api:type:ClassificationShippingRateInput).
	// - If `CartScore`, it is [ScoreShippingRateInput](ctp:api:type:ScoreShippingRateInput).
	// - If `CartValue`, it cannot be used.
	ShippingRateInput ShippingRateInput `json:"shippingRateInput,omitempty"`
	// Custom Fields of Shipping with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
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
	// User-defined unique identifier for the Shipping in a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	Key string `json:"key"`
	// Shipping Methods added to the Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingMethod *ShippingMethodReference `json:"shippingMethod,omitempty"`
	// Determines the shipping rate and Tax Rate of the associated Line Items.
	ShippingAddress BaseAddress `json:"shippingAddress"`
	// Input used to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	// The data type of this field depends on the `shippingRateInputType.type` configured in the [Project](ctp:api:type:Project):
	//
	// - If `CartClassification`, it must be [ClassificationShippingRateInputDraft](ctp:api:type:ClassificationShippingRateInputDraft).
	// - If `CartScore`, it must be [ScoreShippingRateInputDraft](ctp:api:type:ScoreShippingRateInputDraft).
	// - If `CartValue`, it cannot be set.
	ShippingRateInput ShippingRateInputDraft `json:"shippingRateInput,omitempty"`
	// Tax Rate used for taxing a shipping expense if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// Deliveries to be shipped with the Shipping Method.
	Deliveries []DeliveryDraft `json:"deliveries"`
	// Custom Fields for Shipping.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingDraft) MarshalJSON() ([]byte, error) {
	type Alias ShippingDraft
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

type ShippingInfo struct {
	// Name of the Shipping Method.
	ShippingMethodName string `json:"shippingMethodName"`
	// Determined based on the [ShippingRate](ctp:api:type:ShippingRate) and its tiered prices, and either the sum of [LineItem](ctp:api:type:LineItem) prices or the `shippingRateInput` field.
	Price CentPrecisionMoney `json:"price"`
	// Used to determine the price.
	ShippingRate ShippingRate `json:"shippingRate"`
	// Automatically set after the `taxRate` is set.
	TaxedPrice *TaxedItemPrice `json:"taxedPrice,omitempty"`
	// Automatically set in the `Platform` [TaxMode](ctp:api:type:TaxMode) after the [shipping address is set](ctp:api:type:CartSetShippingAddressAction).
	//
	// For the `External` [TaxMode](ctp:api:type:TaxMode) the Tax Rate must be set explicitly with the [ExternalTaxRateDraft](ctp:api:type:ExternalTaxRateDraft).
	TaxRate *TaxRate `json:"taxRate,omitempty"`
	// Used to select a Tax Rate when a Cart has the `Platform` [TaxMode](ctp:api:type:TaxMode).
	TaxCategory *TaxCategoryReference `json:"taxCategory,omitempty"`
	// Not set if a custom Shipping Method is used.
	ShippingMethod *ShippingMethodReference `json:"shippingMethod,omitempty"`
	// Information on how items are delivered to customers.
	Deliveries []Delivery `json:"deliveries"`
	// Discounted price of the Shipping Method.
	DiscountedPrice *DiscountedLineItemPrice `json:"discountedPrice,omitempty"`
	// Indicates whether the [ShippingMethod](ctp:api:type:ShippingMethod) referenced in this ShippingInfo is allowed for the Cart.
	ShippingMethodState ShippingMethodState `json:"shippingMethodState"`
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

/**
*	Determines whether the selected [ShippingMethod](ctp:api:type:ShippingMethod) is allowed for the Cart. For more information, see [Predicates](/shipping-delivery-overview#predicates).
*
 */
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
	// Key of the value used as a programmatic identifier.
	Key string `json:"key"`
	// Descriptive localized label of the value.
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
	// Abstract value for categorizing a Cart.
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
*	Generic type holding specific ShippingRateInputDraft types.
 */
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
	// Key of the value used as a programmatic identifier.
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
	// Abstract value for categorizing a Cart.
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

/**
*	Indicates both Product Discounts and Cart Discounts apply to a Cart and Order.
*
 */
type Stacking struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Stacking) MarshalJSON() ([]byte, error) {
	type Alias Stacking
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Stacking", Alias: (*Alias)(&obj)})
}

/**
*	Determines in which [Tax calculation mode](/carts-orders-overview#tax-calculation-mode) taxed prices are calculated.
*
 */
type TaxCalculationMode string

const (
	TaxCalculationModeLineItemLevel  TaxCalculationMode = "LineItemLevel"
	TaxCalculationModeUnitPriceLevel TaxCalculationMode = "UnitPriceLevel"
)

/**
*	Indicates how taxes are set on the Cart.
*
 */
type TaxMode string

const (
	TaxModePlatform       TaxMode = "Platform"
	TaxModeExternal       TaxMode = "External"
	TaxModeExternalAmount TaxMode = "ExternalAmount"
	TaxModeDisabled       TaxMode = "Disabled"
)

/**
*	The tax portions are calculated from the [TaxRates](ctp:api:type:TaxRate).
*	If a Tax Rate has [SubRates](ctp:api:type:SubRate), they are used and can be identified by name.
*	Tax portions from Line Items with the same `rate` and `name` are accumulated to the same tax portion.
*
 */
type TaxPortion struct {
	// Name of the tax portion.
	Name *string `json:"name,omitempty"`
	// A number in the range 0-1.
	Rate float64 `json:"rate"`
	// Money value of the tax portion.
	Amount CentPrecisionMoney `json:"amount"`
}

/**
*	Represents the portions that sum up to the `totalGross` field of a [TaxedPrice](ctp:api:type:TaxedPrice).
*
*	The portions are calculated from the [TaxRates](ctp:api:type:TaxRate).
*	If a Tax Rate has [SubRates](ctp:api:type:SubRate), they are used and can be identified by name.
*	Tax portions from Line Items with the same `rate` and `name` will be accumulated to the same tax portion.
*
 */
type TaxPortionDraft struct {
	// Name of the tax portion.
	Name *string `json:"name,omitempty"`
	// A number in the range 0-1.
	Rate float64 `json:"rate"`
	// Money value for the tax portion.
	Amount Money `json:"amount"`
}

type TaxedItemPrice struct {
	// Total net amount of the Line Item or Custom Line Item.
	TotalNet CentPrecisionMoney `json:"totalNet"`
	// Total gross amount of the Line Item or Custom Line Item.
	TotalGross CentPrecisionMoney `json:"totalGross"`
	// Taxable portions added to the total net price.
	//
	// Calculated from the [TaxRates](ctp:api:type:TaxRate).
	TaxPortions []TaxPortion `json:"taxPortions"`
	// Total tax applicable for the Line Item or Custom Line Item.
	// Automatically calculated as the difference between the `totalGross` and `totalNet` values.
	TotalTax *CentPrecisionMoney `json:"totalTax,omitempty"`
}

type TaxedPrice struct {
	// Total net price of the Cart or Order.
	TotalNet CentPrecisionMoney `json:"totalNet"`
	// Total gross price of the Cart or Order.
	TotalGross CentPrecisionMoney `json:"totalGross"`
	// Taxable portions added to the total net price.
	//
	// Calculated from the [TaxRates](ctp:api:type:TaxRate).
	TaxPortions []TaxPortion `json:"taxPortions"`
	// Total tax applicable for the Cart or Order.
	//
	// Automatically calculated as the difference between the `totalGross` and `totalNet` values.
	TotalTax *CentPrecisionMoney `json:"totalTax,omitempty"`
}

type TaxedPriceDraft struct {
	// Total net price of the Cart or Order.
	TotalNet Money `json:"totalNet"`
	// Total gross price of the Cart or Order.
	TotalGross Money `json:"totalGross"`
	// Taxable portions added to the `totalGross`.
	//
	// Calculated from the [TaxRates](ctp:api:type:TaxRate).
	TaxPortions []TaxPortionDraft `json:"taxPortions"`
	// Total tax applicable for the Cart or Order.
	TotalTax TypedMoneyDraft `json:"totalTax,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *TaxedPriceDraft) UnmarshalJSON(data []byte) error {
	type Alias TaxedPriceDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.TotalTax != nil {
		var err error
		obj.TotalTax, err = mapDiscriminatorTypedMoneyDraft(obj.TotalTax)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	If the Cart already contains a [CustomLineItem](ctp:api:type:CustomLineItem) with the same `slug`, `name`, `money`, `taxCategory`, `state`,
*	and Custom Fields, then only the quantity of the existing Custom Line Item is increased.
*	If [CustomLineItem](ctp:api:type:CustomLineItem) `shippingDetails` are set, they are merged with the `targets` that already exist on the
*	[ItemShippingDetails](ctp:api:type:ItemShippingDetails) of the Custom Line Item.
*	In case of overlapping address keys the [ItemShippingTarget](ctp:api:type:ItemShippingTarget) `quantity` is summed up.
*
*	If the Cart already contains a Custom Line Item with the same slug that is otherwise not identical, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
*
*	If the Tax Rate is not set, a [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError) error is returned.
*
 */
type CartAddCustomLineItemAction struct {
	// Money value of the Custom Line Item.
	// The value can be negative.
	Money Money `json:"money"`
	// Name of the Custom Line Item.
	Name LocalizedString `json:"name"`
	// User-defined unique identifier of the Custom Line Item.
	Key *string `json:"key,omitempty"`
	// Number of Custom Line Items to add to the Cart.
	Quantity *int `json:"quantity,omitempty"`
	// User-defined identifier used in a deep-link URL for the Custom Line Item.
	// It must match the pattern `[a-zA-Z0-9_-]{2,256}`.
	Slug string `json:"slug"`
	// Used to select a Tax Rate when a Cart has the `Platform` [TaxMode](ctp:api:type:TaxMode).
	//
	// If [TaxMode](ctp:api:type:TaxMode) is `Platform`, this field must not be empty.
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// An external Tax Rate can be set if the Cart has `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// Container for Custom Line Item-specific addresses.
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	// Custom Fields for the Custom Line Item.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// - If `Standard`, Cart Discounts with a matching [CartDiscountCustomLineItemsTarget](ctp:api:type:CartDiscountCustomLineItemsTarget), [MultiBuyCustomLineItemsTarget](ctp:api:type:MultiBuyCustomLineItemsTarget), or [CartDiscountPatternTarget](ctp:api:type:CartDiscountPatternTarget) are applied to the Custom Line Item.
	// - If `External`, Cart Discounts are not considered on the Custom Line Item.
	PriceMode *CustomLineItemPriceMode `json:"priceMode,omitempty"`
	// Recurring Order and frequency data.
	RecurrenceInfo *CustomLineItemRecurrenceInfoDraft `json:"recurrenceInfo,omitempty"`
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

/**
*	To add a custom Shipping Method (independent of the [ShippingMethods](ctp:api:type:ShippingMethod) managed through
*	the [Shipping Methods API](/projects/shippingMethods)) to the Cart, it **must have** the `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
*
 */
type CartAddCustomShippingMethodAction struct {
	// User-defined identifier for the custom Shipping Method that must be unique across the Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey string `json:"shippingKey"`
	// Name of the custom Shipping Method.
	ShippingMethodName string `json:"shippingMethodName"`
	// Determines the shipping rate and Tax Rate of the associated Line Items.
	ShippingAddress BaseAddress `json:"shippingAddress"`
	// Determines the shipping price.
	ShippingRate ShippingRateDraft `json:"shippingRate"`
	// Input used to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	// The data type of this field depends on the `shippingRateInputType.type` configured in the [Project](ctp:api:type:Project):
	//
	// - If `CartClassification`, it must be [ClassificationShippingRateInputDraft](ctp:api:type:ClassificationShippingRateInputDraft).
	// - If `CartScore`, it must be [ScoreShippingRateInputDraft](ctp:api:type:ScoreShippingRateInputDraft).
	// - If `CartValue`, it cannot be set.
	ShippingRateInput ShippingRateInputDraft `json:"shippingRateInput,omitempty"`
	// Tax Category used to determine a shipping Tax Rate if the Cart has the `Platform` [TaxMode](ctp:api:type:TaxMode).
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// Tax Rate used to tax a shipping expense if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// Deliveries to be shipped with the custom Shipping Method.
	Deliveries []DeliveryDraft `json:"deliveries"`
	// Custom Fields for the custom Shipping Method.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
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
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addCustomShippingMethod", Alias: (*Alias)(&obj)})
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
*	Adds a [DiscountCode](ctp:api:type:DiscountCode) to the Cart to activate the related [Cart Discounts](/../api/projects/cartDiscounts).
*	Adding a Discount Code is only possible if no [DirectDiscount](ctp:api:type:DirectDiscount) has been applied to the Cart.
*	Discount Codes can be added to [frozen Carts](ctp:api:type:FrozenCarts), but their [DiscountCodeState](ctp:api:type:DiscountCodeState) is then `DoesNotMatchCart`.
*
*	The maximum number of Discount Codes in a Cart is restricted by a [limit](/../api/limits#carts).
*
*	Specific Error Code: [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError)
*
 */
type CartAddDiscountCodeAction struct {
	// `code` of a [DiscountCode](ctp:api:type:DiscountCode).
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

/**
*	Adds an address to a Cart when shipping to multiple addresses is desired.
*
 */
type CartAddItemShippingAddressAction struct {
	// Address to append to `itemShippingAddresses`.
	//
	// The new Address must have a `key` that is unique across this Cart.
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

/**
*	If the Cart contains a [LineItem](ctp:api:type:LineItem) for a Product Variant with the same [LineItemMode](ctp:api:type:LineItemMode), [Custom Fields](/../api/projects/custom-fields), supply and distribution channel, then only the quantity of the existing Line Item is increased.
*	If [LineItem](ctp:api:type:LineItem) `shippingDetails` is set, it is merged. All addresses will be present afterwards and, for address keys present in both shipping details, the quantity will be summed up.
*	A new Line Item is added when the `externalPrice` or `externalTotalPrice` is set in this update action.
*	The [LineItem](ctp:api:type:LineItem) price is set as described in [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection).
*
*	If the Tax Rate is not set, a [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError) error is returned.
*
*	If the Line Items do not have a Price according to the [Product](ctp:api:type:Product) `priceMode` value for a selected currency and/or country, Customer Group, or Channel, a [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError) error is returned.
*
*	If the Line Items are added to a Cart bound to a Store with active Product Selections, the selected Product Variant must be [available in that Store](/../api/projects/stores#products-available-in-store), otherwise an [InvalidInput](ctp:api:type:InvalidInputError) error is returned.
*
 */
type CartAddLineItemAction struct {
	// User-defined unique identifier of the LineItem.
	Key *string `json:"key,omitempty"`
	// `id` of the published [Product](ctp:api:type:Product).
	//
	// Either the `productId` and `variantId`, or `sku` must be provided.
	ProductId *string `json:"productId,omitempty"`
	// `id` of the [ProductVariant](ctp:api:type:ProductVariant) in the Product.
	// If not provided, the Master Variant is used.
	//
	// Either the `productId` and `variantId`, or `sku` must be provided.
	VariantId *int `json:"variantId,omitempty"`
	// SKU of the [ProductVariant](ctp:api:type:ProductVariant).
	//
	// Either the `productId` and `variantId`, or `sku` must be provided.
	Sku *string `json:"sku,omitempty"`
	// Quantity of the Product Variant to add to the Cart.
	Quantity *int `json:"quantity,omitempty"`
	// Date and time (UTC) the Product Variant is added to the Cart.
	// If not set, it defaults to the current date and time.
	//
	// Optional for backwards compatibility reasons.
	AddedAt *time.Time `json:"addedAt,omitempty"`
	// Used to [select](/../api/pricing-and-discounts-overview#line-item-price-selection) a Product Price.
	// The Channel must have the `ProductDistribution` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum).
	// If the Cart is bound to a [Store](ctp:api:type:Store) with `distributionChannels` set, the Channel must match one of the Store's distribution channels.
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
	// Used to identify [Inventory entries](/../api/projects/inventory) that must be reserved.
	// The Channel must have the `InventorySupply` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum).
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	// Sets the [LineItem](ctp:api:type:LineItem) `price` value, and the `priceMode` to `ExternalPrice` [LineItemPriceMode](ctp:api:type:LineItemPriceMode).
	ExternalPrice *Money `json:"externalPrice,omitempty"`
	// Sets the [LineItem](ctp:api:type:LineItem) `price` and `totalPrice` values, and the `priceMode` to `ExternalTotal` [LineItemPriceMode](ctp:api:type:LineItemPriceMode).
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	// Sets the external Tax Rate for the Line Item, if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode) and `Single` [ShippingMode](ctp:api:type:ShippingMode). If the Cart has `Multiple` [ShippingMode](ctp:api:type:ShippingMode), the Tax Rate is accepted but ignored.
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// Sets the external Tax Rates for individual Shipping Methods, if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode) and `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	PerMethodExternalTaxRate []MethodExternalTaxRateDraft `json:"perMethodExternalTaxRate"`
	// Inventory mode specific to the Line Item only, and valid for the entire `quantity` of the Line Item.
	// Set only if the inventory mode should be different from the `inventoryMode` specified on the [Cart](ctp:api:type:Cart).
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	// Container for Line Item-specific addresses.
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	// Recurring Order and frequency data.
	RecurrenceInfo *LineItemRecurrenceInfoDraft `json:"recurrenceInfo,omitempty"`
	// Custom Fields for the Line Item.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias CartAddLineItemAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["perMethodExternalTaxRate"] == nil {
		delete(raw, "perMethodExternalTaxRate")
	}

	return json.Marshal(raw)

}

type CartAddPaymentAction struct {
	// Payment to add to the Cart.
	// Must not be assigned to another Order or active Cart already.
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
*	Adds a Shipping Method for a specified shipping address to a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
*
 */
type CartAddShippingMethodAction struct {
	// User-defined identifier for the [Shipping](ctp:api:type:Shipping) that must be unique across the Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey string `json:"shippingKey"`
	// ResourceIdentifier to a [ShippingMethod](ctp:api:type:ShippingMethod) to add to the Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	// If the referenced Shipping Method has a predicate that does not match the Cart, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
	ShippingMethod ShippingMethodResourceIdentifier `json:"shippingMethod"`
	// Determines the shipping rate and Tax Rate of the Line Items.
	ShippingAddress BaseAddress `json:"shippingAddress"`
	// Input used to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	// The data type of this field depends on the `shippingRateInputType.type` configured in the [Project](ctp:api:type:Project):
	//
	// - If `CartClassification`, it must be [ClassificationShippingRateInputDraft](ctp:api:type:ClassificationShippingRateInputDraft).
	// - If `CartScore`, it must be [ScoreShippingRateInputDraft](ctp:api:type:ScoreShippingRateInputDraft).
	// - If `CartValue`, it cannot be set.
	ShippingRateInput ShippingRateInputDraft `json:"shippingRateInput,omitempty"`
	// Tax Rate used to tax a shipping expense if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// Deliveries to be shipped with the referenced Shipping Method.
	Deliveries []DeliveryDraft `json:"deliveries"`
	// Custom Fields for the Shipping Method.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
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
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingMethod", Alias: (*Alias)(&obj)})
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
*	Adds all [LineItems](ctp:api:type:LineItem) of a [ShoppingList](ctp:api:type:ShoppingList) to the Cart.
*
 */
type CartAddShoppingListAction struct {
	// Shopping List that contains the Line Items to be added.
	ShoppingList ShoppingListResourceIdentifier `json:"shoppingList"`
	// `distributionChannel` to set for all [LineItems](ctp:api:type:LineItem) that are added to the Cart.
	// The Channel must have the `ProductDistribution` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum).
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
	// `supplyChannel` to set for all [LineItems](ctp:api:type:LineItem) that are added to the Cart.
	// The Channel must have the `InventorySupply` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum).
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
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
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
	// Using positive or negative quantities increases or decreases the number of items shipped to an address.
	TargetsDelta []ItemShippingTarget `json:"targetsDelta"`
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

/**
*	To override the shipping details, see [Set LineItem ShippingDetails](ctp:api:type:CartSetLineItemShippingDetailsAction).
*
 */
type CartApplyDeltaToLineItemShippingDetailsTargetsAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Using positive or negative quantities increases or decreases the number of items shipped to an address.
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
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
	// Value to set. Must not be empty. Can be a negative amount.
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
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
	// New value to set. Must not be empty.
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

/**
*	When multiple shipping addresses are set for a Custom Line Item,
*	use the [Add CustomLineItem](ctp:api:type:CartAddCustomLineItemAction) update action to change the shipping details.
*	Since it is not possible for the API to infer how the overall change in the Custom Line Item quantity should be distributed over the sub-quantities,
*	the `shippingDetails` field is kept in its current state to avoid data loss.
*
*	To change the Custom Line Item quantity and shipping details together,
*	use this update action in combination with the [Set CustomLineItem ShippingDetails](ctp:api:type:CartSetCustomLineItemShippingDetailsAction) update action
*	in a single Cart update command.
*
 */
type CartChangeCustomLineItemQuantityAction struct {
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
	// New value to set.
	//
	// If `0`, the Custom Line Item is removed from the Cart.
	Quantity int `json:"quantity"`
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

/**
*	When multiple shipping addresses are set for a Line Item,
*	use the [Remove LineItem](ctp:api:type:CartRemoveLineItemAction) and [Add LineItem](ctp:api:type:CartAddLineItemAction) update action
*	to change the shipping details.
*	Since it is not possible for the API to infer how the overall change in the Line Item quantity should be distributed over the sub-quantities,
*	the `shippingDetails` field is kept in its current state to avoid data loss.
*
*	To change the Line Item quantity and shipping details together,
*	use this update action in combination with the [Set LineItem ShippingDetails](ctp:api:type:CartSetLineItemShippingDetailsAction) update action
*	in a single Cart update command.
*
*	The [LineItem](ctp:api:type:LineItem) price is set as described in [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection).
*
 */
type CartChangeLineItemQuantityAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// New value to set.
	//
	// If `0`, the Line Item is removed from the Cart.
	Quantity int `json:"quantity"`
	// Required when the Line Item uses `ExternalPrice` [LineItemPriceMode](ctp:api:type:LineItemPriceMode).
	// Sets the [LineItem](ctp:api:type:LineItem) `price` to the given value when changing the quantity of a Line Item.
	//
	// The LineItem price is updated as described in Line Item price selection.
	ExternalPrice *Money `json:"externalPrice,omitempty"`
	// Sets the [LineItem](ctp:api:type:LineItem) `price` and `totalPrice` to the given value when changing the quantity of a Line Item with the `ExternalTotal` [LineItemPriceMode](ctp:api:type:LineItemPriceMode).
	// If `externalTotalPrice` is not given and the `priceMode` is `ExternalTotal`, the external price is unset and the `priceMode` is set to `Platform`.
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

type CartChangeLineItemsOrderAction struct {
	// All existing [LineItem](ctp:api:type:LineItem) `id`s in the desired new order.
	LineItemOrder []string `json:"lineItemOrder"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartChangeLineItemsOrderAction) MarshalJSON() ([]byte, error) {
	type Alias CartChangeLineItemsOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemsOrder", Alias: (*Alias)(&obj)})
}

/**
*	Changing the price rounding mode leads to [recalculation of taxes](/../api/carts-orders-overview#taxes).
*
 */
type CartChangePriceRoundingModeAction struct {
	// New value to set.
	PriceRoundingMode RoundingMode `json:"priceRoundingMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartChangePriceRoundingModeAction) MarshalJSON() ([]byte, error) {
	type Alias CartChangePriceRoundingModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePriceRoundingMode", Alias: (*Alias)(&obj)})
}

/**
*	Changing the tax calculation mode leads to [recalculation of taxes](/../api/carts-orders-overview#taxes).
*
 */
type CartChangeTaxCalculationModeAction struct {
	// New value to set.
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

/**
*	- When `External` [TaxMode](ctp:api:type:TaxMode) is changed to `Platform` or `Disabled`, all previously set external Tax Rates are removed.
*	- When set to `Platform`, Line Items, Custom Line Items, and Shipping Method require a Tax Category with a Tax Rate for the Cart's `shippingAddress`.
*
 */
type CartChangeTaxModeAction struct {
	// The new TaxMode.
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

/**
*	Changing the tax rounding mode leads to [recalculation of taxes](/../api/carts-orders-overview#taxes).
*
 */
type CartChangeTaxRoundingModeAction struct {
	// New value to set.
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

/**
*	Changes the [CartState](ctp:api:type:CartState) from `Active` to `Frozen`. Results in a [Frozen Cart](ctp:api:type:FrozenCarts).
*	Fails with [InvalidOperation](ctp:api:type:InvalidOperationError) error when the Cart is empty.
*
 */
type CartFreezeCartAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartFreezeCartAction) MarshalJSON() ([]byte, error) {
	type Alias CartFreezeCartAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "freezeCart", Alias: (*Alias)(&obj)})
}

/**
*	This update action does not set any Cart field in particular, but it triggers several [Cart updates](/../api/carts-orders-overview#update-a-cart)
*	to bring prices and discounts to the latest state. Those can become stale over time when no Cart updates have been performed for a while and
*	prices on related Products have changed in the meanwhile.
*
*	If the `priceMode` of the [Product](ctp:api:type:Product) related to a Line Item is of `Embedded` [ProductPriceMode](ctp:api:type:ProductPriceModeEnum),
*	the updated `price` of that [LineItem](ctp:api:type:LineItem) may not correspond to a Price in the `variant.prices` anymore.
*
 */
type CartRecalculateAction struct {
	// - Leave empty or set as `false` to update only the Prices and TaxRates of the Line Items.
	// - Set as `true` to update the Product data (such as `name`, `variant`, `productType`, and Product Attributes) of the Line Items.
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

/**
*	This update action does not support specifying quantities, unlike the [Remove LineItem](ctp:api:type:CartRemoveLineItemAction) update action.
*
*	If `shippingDetails` must be partially removed, use the [Change CustomLineItem Quantity](ctp:api:type:CartChangeCustomLineItemQuantityAction) update action.
*
 */
type CartRemoveCustomLineItemAction struct {
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
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
	// Discount Code to remove from the Cart.
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

/**
*	An address can only be removed if it is not referenced in any [ItemShippingTarget](ctp:api:type:ItemShippingTarget) of the Cart.
*
 */
type CartRemoveItemShippingAddressAction struct {
	// `key` of the Address to remove from `itemShippingAddresses`.
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

/**
*	The [LineItem](ctp:api:type:LineItem) price is updated as described in [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection).
*
 */
type CartRemoveLineItemAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Amount to subtract from the LineItem's `quantity`.
	// If absent, the LineItem is removed from the Cart.
	Quantity *int `json:"quantity,omitempty"`
	// Sets the [LineItem](ctp:api:type:LineItem) `price` to the given value when decreasing the quantity of a Line Item with the `ExternalPrice` [LineItemPriceMode](ctp:api:type:LineItemPriceMode).
	ExternalPrice *Money `json:"externalPrice,omitempty"`
	// Sets the [LineItem](ctp:api:type:LineItem) `price` and `totalPrice` to the given value when decreasing the quantity of a Line Item with the `ExternalTotal` [LineItemPriceMode](ctp:api:type:LineItemPriceMode).
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	// Container for Line Item-specific addresses to remove.
	ShippingDetailsToRemove *ItemShippingDetailsDraft `json:"shippingDetailsToRemove,omitempty"`
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
	// Payment to remove from the Cart.
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

/**
*	Removes a Shipping Method from a Cart that has the `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
*
 */
type CartRemoveShippingMethodAction struct {
	// User-defined unique identifier of the Shipping Method to remove from the Cart.
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

/**
*	If the Cart is already associated with a Customer, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
*
 */
type CartSetAnonymousIdAction struct {
	// Value to set.
	// If empty, any existing value is removed.
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
	// Value to set.
	// If empty, any existing value is removed.
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
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
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

/**
*	Updates the Business Unit on the Cart. The Cart must have an existing Business Unit assigned already.
*
 */
type CartSetBusinessUnitAction struct {
	// New Business Unit to assign to the Cart, which must have access to the [Store](/../api/projects/stores) that is set on the Cart.
	BusinessUnit BusinessUnitResourceIdentifier `json:"businessUnit"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetBusinessUnitAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetBusinessUnitAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBusinessUnit", Alias: (*Alias)(&obj)})
}

/**
*	Can be used if the Cart has the `ExternalAmount` [TaxMode](ctp:api:type:TaxMode). This update action adds the `taxedPrice` field to the Cart. It sets the `totalGross` amount, and Composable Commerce calculates the `totalNet` and `totalTax` values based on the provided `externalTotalGross`.
*	You must use this update action after any price-affecting change occurs within the Cart.
*
 */
type CartSetCartTotalTaxAction struct {
	// The total gross amount of the Cart, including tax. This value is used to calculate the `totalNet` and `totalTax` fields of the Cart's `taxedPrice`.
	ExternalTotalGross Money `json:"externalTotalGross"`
	// Set if the `externalTotalGross` price is a sum of portions with different tax rates.
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

/**
*	Setting the country can lead to changes in the [LineItem](ctp:api:type:LineItem) prices.
*
 */
type CartSetCountryAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	//
	// If the Cart is bound to a `store`, the provided value must be included in the [Store's](ctp:api:type:Store) `countries`.
	// Otherwise a [CountryNotConfiguredInStore](ctp:api:type:CountryNotConfiguredInStoreError) error is returned.
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
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
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
func (obj CartSetCustomLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomField", Alias: (*Alias)(&obj)})
}

type CartSetCustomLineItemCustomTypeAction struct {
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
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

/**
*	Sets the recurrence information on the [CustomLineItem](ctp:api:type:CustomLineItem).
*	If the Cart is already associated with a Recurring Order, this action will fail.
*
 */
type CartSetCustomLineItemRecurrenceInfoAction struct {
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
	// Value to set.
	// If empty, any existing value will be removed.
	RecurrenceInfo *CustomLineItemRecurrenceInfoDraft `json:"recurrenceInfo,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetCustomLineItemRecurrenceInfoAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomLineItemRecurrenceInfoAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemRecurrenceInfo", Alias: (*Alias)(&obj)})
}

type CartSetCustomLineItemShippingDetailsAction struct {
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
func (obj CartSetCustomLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

/**
*	Can be used if the Cart has the `ExternalAmount` [TaxMode](ctp:api:type:TaxMode).
*
 */
type CartSetCustomLineItemTaxAmountAction struct {
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
	// Value to set.
	// If empty, any existing value is removed.
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
	// `key` of the [ShippingMethod](ctp:api:type:ShippingMethod) used for this Custom Line Item.
	// This is required for Carts with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	Can be used if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
*
 */
type CartSetCustomLineItemTaxRateAction struct {
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
	// Value to set.
	// If empty, an existing value is removed.
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// `key` of the [ShippingMethod](ctp:api:type:ShippingMethod) used for this Custom Line Item.
	// This is required for Carts with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
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

/**
*	To set the Cart's custom Shipping Method (independent of the [ShippingMethods](ctp:api:type:ShippingMethod) managed through
*	the [Shipping Methods API](/projects/shippingMethods)) the Cart must have
*	the `Single` [ShippingMode](ctp:api:type:ShippingMode) and a `shippingAddress`.
*
*	To unset a custom Shipping Method on a Cart, use the [Set ShippingMethod](ctp:api:type:CartSetShippingMethodAction) update action
*	without the `shippingMethod` field instead.
*
 */
type CartSetCustomShippingMethodAction struct {
	// Name of the custom Shipping Method.
	ShippingMethodName string `json:"shippingMethodName"`
	// Determines the shipping price.
	ShippingRate ShippingRateDraft `json:"shippingRate"`
	// Tax Category used to determine the Tax Rate when the Cart has the `Platform` [TaxMode](ctp:api:type:TaxMode).
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// External Tax Rate for the `shippingRate` to be set if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// Custom Fields for the custom Shipping Method.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
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
	// Value to set.
	// If empty, any existing value is removed.
	Email *string `json:"email,omitempty"`
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

/**
*	This update action can only be used if a Customer is not assigned to the Cart.
*	If a Customer is already assigned, the Cart uses the Customer Group of the assigned Customer.
*
*	To reflect the new Customer Group, this update action can result in [updates to the Cart](/api/carts-orders-overview#update-a-cart). When this occurs, the following errors can be returned: [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError) and [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError).
*
 */
type CartSetCustomerGroupAction struct {
	// Value to set.
	// If empty, any existing value is removed.
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

/**
*	Setting the Cart's `customerId` can lead to updates on all its [LineItem](ctp:api:type:LineItem) `prices`.
*
*	If the Customer with the specified `id` cannot be found, this update action returns a
*	[ReferencedResourceNotFound](ctp:api:type:ReferencedResourceNotFoundError) error.
*
 */
type CartSetCustomerIdAction struct {
	// `id` of an existing [Customer](ctp:api:type:Customer).
	// If the Customer is assigned to a [CustomerGroup](ctp:api:type:CustomerGroup), this update action also sets the value for the `customerGroup` field.
	// If empty, the update action removes the value for both `customerId` and `customerGroup`.
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

/**
*	Number of days after the last modification before a Cart is deleted.
*
*	If a [ChangeSubscription](ctp:api:type:ChangeSubscription) exists for Carts, a [ResourceDeletedDeliveryPayload](ctp:api:type:ResourceDeletedDeliveryPayload) is sent.
*
 */
type CartSetDeleteDaysAfterLastModificationAction struct {
	// Value to set.
	// If not provided, the default value for this field configured in [Project settings](ctp:api:type:CartsConfiguration) is assigned.
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

/**
*	Adds a [DirectDiscount](ctp:api:type:DirectDiscount), but only if no [DiscountCode](ctp:api:type:DiscountCode) has been added to the Cart.
*	Either a Discount Code or a Direct Discount can exist on a Cart at the same time.
*
 */
type CartSetDirectDiscountsAction struct {
	// - If set, all existing Direct Discounts are replaced.
	//   The discounts apply in the order they are added to the list.
	// - If empty, all existing Direct Discounts are removed and all affected prices on the Cart or Order are recalculated.
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
	// `key` of the [Address](ctp:api:type:Address) in `itemShippingAddress`.
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
func (obj CartSetItemShippingAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetItemShippingAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setItemShippingAddressCustomField", Alias: (*Alias)(&obj)})
}

type CartSetItemShippingAddressCustomTypeAction struct {
	// `key` of the [Address](ctp:api:type:Address) in `itemShippingAddress`.
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
	// Value to set.
	// If empty, any existing key will be removed.
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
func (obj CartSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

type CartSetLineItemCustomTypeAction struct {
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
func (obj CartSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

/**
*	Setting a distribution channel for a [LineItem](ctp:api:type:LineItem) can lead to an updated `price` as described in [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection).
*
 */
type CartSetLineItemDistributionChannelAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// - If present, a [Reference](ctp:api:type:Reference) to the Channel is set for the [LineItem](ctp:api:type:LineItem) specified by `lineItemId`.
	// - If not present, the current [Reference](ctp:api:type:Reference) to a distribution channel is removed from the [LineItem](ctp:api:type:LineItem) specified by `lineItemId`.
	//   The Channel must have the `ProductDistribution` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum).
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

type CartSetLineItemInventoryModeAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Inventory mode specific to the Line Item only, and valid for the entire `quantity` of the Line Item.
	// Set only if the inventory mode should be different from the `inventoryMode` specified on the [Cart](ctp:api:type:Cart).
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetLineItemInventoryModeAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemInventoryModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemInventoryMode", Alias: (*Alias)(&obj)})
}

/**
*	Sets the [LineItem](ctp:api:type:LineItem) `price` and changes the `priceMode` to `ExternalPrice` [LineItemPriceMode](ctp:api:type:LineItemPriceMode).
*
 */
type CartSetLineItemPriceAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Value to set.
	// If `externalPrice` is not given and the `priceMode` is `ExternalPrice`, the external price is unset and the `priceMode` is set to `Platform`.
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

/**
*	Sets the recurrence information on the [LineItem](ctp:api:type:LineItem).
*	If the Cart is already associated with a Recurring Order, this action will fail.
*
 */
type CartSetLineItemRecurrenceInfoAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Value to set.
	// If empty, any existing value will be removed.
	RecurrenceInfo *LineItemRecurrenceInfoDraft `json:"recurrenceInfo,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartSetLineItemRecurrenceInfoAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemRecurrenceInfoAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemRecurrenceInfo", Alias: (*Alias)(&obj)})
}

type CartSetLineItemShippingDetailsAction struct {
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
func (obj CartSetLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

/**
*	Performing this action does not reserve stock. Stock is only reserved at Order creation if the [InventoryMode](ctp:api:type:InventoryMode) of the Cart is `TrackOnly` or `ReserveOnOrder`.
*
 */
type CartSetLineItemSupplyChannelAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// - If present, a [Reference](ctp:api:type:Reference) to the Channel is set for the [LineItem](ctp:api:type:LineItem) specified by `lineItemId`.
	// - If not present, the current [Reference](ctp:api:type:Reference) to a supply channel will be removed from the [LineItem](ctp:api:type:LineItem) specified by `lineItemId`.
	//   The Channel must have the `InventorySupply` [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum).
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

/**
*	Can be used if the Cart has the `ExternalAmount` [TaxMode](ctp:api:type:TaxMode). This update action sets the `taxedPrice` and `taxRate` on a Line Item and must be used after any price-affecting change occurs.
*
 */
type CartSetLineItemTaxAmountAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Value to set.
	// If empty, any existing value is removed.
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
	// `key` of the [ShippingMethod](ctp:api:type:ShippingMethod) used for this Line Item.
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

/**
*	Can be used if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
*
 */
type CartSetLineItemTaxRateAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Value to set.
	// If empty, any existing value is removed.
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// `key` of the [ShippingMethod](ctp:api:type:ShippingMethod) used for this Line Item.
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

/**
*	Sets the [LineItem](ctp:api:type:LineItem) `totalPrice` and `price`, and changes the `priceMode` to `ExternalTotal` [LineItemPriceMode](ctp:api:type:LineItemPriceMode).
*
 */
type CartSetLineItemTotalPriceAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Value to set.
	// If `externalTotalPrice` is not given and the `priceMode` is `ExternalTotal`, the external price is unset and the `priceMode` is set to `Platform`.
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
	// Value to set.
	// Must be one of the [Project](ctp:api:type:Project)'s `languages`.
	// If empty, any existing value will be removed.
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

/**
*	Setting the shipping address also sets the [TaxRate](ctp:api:type:TaxRate) of Line Items and calculates the [TaxedPrice](ctp:api:type:TaxedPrice).
*
*	If a matching price cannot be found for the given shipping address during [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection),
*	a [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError) error is returned.
*
*	If you want to allow shipping to states inside a country that are not explicitly covered by a TaxRate,
*	set the `countryTaxRateFallbackEnabled` field to `true` in the [CartsConfiguration](ctp:api:type:CartsConfiguration) by using
*	the [Change CountryTaxRateFallbackEnabled](ctp:api:type:ProjectChangeCountryTaxRateFallbackEnabledAction) update action.
*
 */
type CartSetShippingAddressAction struct {
	// Value to set.
	// If not set, the shipping address is unset, and the `taxedPrice` and `taxRate` are unset in all Line Items of the Cart.
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
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
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
	// The `shippingKey` of the [Shipping](ctp:api:type:Shipping) to customize. Used to specify which Shipping Method to customize
	// on a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	// Leave this empty to customize the one and only ShippingMethod on a `Single` ShippingMode Cart.
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

/**
*	This action sets, overwrites, or removes any existing Custom Type and Custom Fields for the Cart's `shippingMethod` or `shipping`.
*
 */
type CartSetShippingCustomTypeAction struct {
	// The `shippingKey` of the [Shipping](ctp:api:type:Shipping) to customize. Used to specify which Shipping Method to customize
	// on a Cart with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	// Leave this empty to customize the one and only ShippingMethod on a `Single` ShippingMode Cart.
	ShippingKey *string `json:"shippingKey,omitempty"`
	// Defines the [Type](ctp:api:type:Type) that extends the specified ShippingMethod with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the ShippingMethod.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `shippingMethod`.
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

/**
*	To set the Cart's Shipping Method the Cart must have the `Single` [ShippingMode](ctp:api:type:ShippingMode) and a `shippingAddress`.
*
 */
type CartSetShippingMethodAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	//
	// If the referenced Shipping Method is inactive, or has a predicate that does not match the Cart, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
	ShippingMethod *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	// An external Tax Rate can be set if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
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

/**
*	A Shipping Method tax amount can be set if the Cart has the `ExternalAmount` [TaxMode](ctp:api:type:TaxMode).
*
 */
type CartSetShippingMethodTaxAmountAction struct {
	// `key` of the [ShippingMethod](ctp:api:type:ShippingMethod) to update. This is required for Carts with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
	// Value to set.
	// If empty, any existing value is removed.
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

/**
*	A Shipping Method Tax Rate can be set if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
*
 */
type CartSetShippingMethodTaxRateAction struct {
	// `key` of the [ShippingMethod](ctp:api:type:ShippingMethod) to update. This is required for Carts with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
	// Value to set.
	// If empty, any existing value is removed.
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

/**
*	Input used to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
*	If no matching tier can be found, or the input is not set, the default price for the shipping rate is used.
*
 */
type CartSetShippingRateInputAction struct {
	// The data type of this field depends on the `shippingRateInputType.type` configured in the [Project](ctp:api:type:Project):
	//
	// - If `CartClassification`, it must be [ClassificationShippingRateInputDraft](ctp:api:type:ClassificationShippingRateInputDraft).
	// - If `CartScore`, it must be [ScoreShippingRateInputDraft](ctp:api:type:ScoreShippingRateInputDraft).
	// - If `CartValue`, it cannot be set.
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

/**
*	Changes the [CartState](ctp:api:type:CartState) from `Frozen` to `Active`. Reactivates a [Frozen Cart](ctp:api:type:FrozenCarts).
*	This action updates all prices in the Cart according to latest Prices on related Product Variants and Shipping Methods and by applying all discounts currently being active and applicable for the Cart.
*
 */
type CartUnfreezeCartAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartUnfreezeCartAction) MarshalJSON() ([]byte, error) {
	type Alias CartUnfreezeCartAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "unfreezeCart", Alias: (*Alias)(&obj)})
}

/**
*	Updates an address in `itemShippingAddresses` by keeping the Address `key`.
*
 */
type CartUpdateItemShippingAddressAction struct {
	// The new Address with the same `key` as the Address it will replace.
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

type ProductTailoringUpdate struct {
	// Expected version of the ProductTailoring on which the changes apply.
	// If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) is returned.
	Version int `json:"version"`
	// Update actions to be performed on the ProductTailoring.
	Actions []ProductTailoringUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductTailoringUpdate) UnmarshalJSON(data []byte) error {
	type Alias ProductTailoringUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorProductTailoringUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	The scope controls which part of the product information is published.
 */
type ProductPublishScope string

const (
	ProductPublishScopeAll    ProductPublishScope = "All"
	ProductPublishScopePrices ProductPublishScope = "Prices"
)
