// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type Cart struct {
	Version                         int                     `json:"version"`
	LastModifiedAt                  time.Time               `json:"lastModifiedAt"`
	ID                              string                  `json:"id"`
	CreatedAt                       time.Time               `json:"createdAt"`
	TotalPrice                      TypedMoney              `json:"totalPrice"`
	TaxedPrice                      *TaxedPrice             `json:"taxedPrice,omitempty"`
	TaxRoundingMode                 RoundingMode            `json:"taxRoundingMode"`
	TaxMode                         TaxMode                 `json:"taxMode"`
	TaxCalculationMode              TaxCalculationMode      `json:"taxCalculationMode"`
	ShippingRateInput               ShippingRateInput       `json:"shippingRateInput,omitempty"`
	ShippingInfo                    *ShippingInfo           `json:"shippingInfo,omitempty"`
	ShippingAddress                 *Address                `json:"shippingAddress,omitempty"`
	RefusedGifts                    []CartDiscountReference `json:"refusedGifts"`
	PaymentInfo                     *PaymentInfo            `json:"paymentInfo,omitempty"`
	Origin                          CartOrigin              `json:"origin"`
	Locale                          string                  `json:"locale,omitempty"`
	LineItems                       []LineItem              `json:"lineItems"`
	ItemShippingAddresses           []Address               `json:"itemShippingAddresses,omitempty"`
	InventoryMode                   InventoryMode           `json:"inventoryMode,omitempty"`
	DiscountCodes                   []DiscountCodeInfo      `json:"discountCodes,omitempty"`
	DeleteDaysAfterLastModification int                     `json:"deleteDaysAfterLastModification,omitempty"`
	CustomerID                      string                  `json:"customerId,omitempty"`
	CustomerGroup                   *CustomerGroupReference `json:"customerGroup,omitempty"`
	CustomerEmail                   string                  `json:"customerEmail,omitempty"`
	CustomLineItems                 []CustomLineItem        `json:"customLineItems"`
	Custom                          *CustomFields           `json:"custom,omitempty"`
	Country                         CountryCode             `json:"country,omitempty"`
	CartState                       CartState               `json:"cartState"`
	BillingAddress                  *Address                `json:"billingAddress,omitempty"`
	AnonymousID                     string                  `json:"anonymousId,omitempty"`
}

func (obj *Cart) UnmarshalJSON(data []byte) error {
	type Alias Cart
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInput != nil {
		obj.ShippingRateInput = AbstractShippingRateInputDiscriminatorMapping(obj.ShippingRateInput)
	}
	if obj.TotalPrice != nil {
		obj.TotalPrice = AbstractTypedMoneyDiscriminatorMapping(obj.TotalPrice)
	}

	return nil
}

type CartAddCustomLineItemAction struct {
	TaxCategory     *TaxCategoryReference `json:"taxCategory,omitempty"`
	Slug            string                `json:"slug"`
	Quantity        float64               `json:"quantity"`
	Name            *LocalizedString      `json:"name"`
	Money           *Money                `json:"money"`
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	Custom          *CustomFieldsDraft    `json:"custom,omitempty"`
}

func (obj CartAddCustomLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias CartAddCustomLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addCustomLineItem", Alias: (*Alias)(&obj)})
}

type CartAddDiscountCodeAction struct {
	Code string `json:"code"`
}

func (obj CartAddDiscountCodeAction) MarshalJSON() ([]byte, error) {
	type Alias CartAddDiscountCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDiscountCode", Alias: (*Alias)(&obj)})
}

type CartAddItemShippingAddressAction struct {
	Address *Address `json:"address"`
}

func (obj CartAddItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CartAddItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addItemShippingAddress", Alias: (*Alias)(&obj)})
}

type CartAddLineItemAction struct {
	VariantID           int                         `json:"variantId,omitempty"`
	SupplyChannel       *ChannelReference           `json:"supplyChannel,omitempty"`
	Sku                 string                      `json:"sku,omitempty"`
	ShippingDetails     *ItemShippingDetailsDraft   `json:"shippingDetails,omitempty"`
	Quantity            float64                     `json:"quantity,omitempty"`
	ProductID           string                      `json:"productId,omitempty"`
	ExternalTotalPrice  *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ExternalTaxRate     *ExternalTaxRateDraft       `json:"externalTaxRate,omitempty"`
	ExternalPrice       *Money                      `json:"externalPrice,omitempty"`
	DistributionChannel *ChannelReference           `json:"distributionChannel,omitempty"`
	Custom              *CustomFieldsDraft          `json:"custom,omitempty"`
}

func (obj CartAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias CartAddLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
}

type CartAddPaymentAction struct {
	Payment *PaymentReference `json:"payment"`
}

func (obj CartAddPaymentAction) MarshalJSON() ([]byte, error) {
	type Alias CartAddPaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPayment", Alias: (*Alias)(&obj)})
}

type CartAddShoppingListAction struct {
	SupplyChannel       *ChannelReference      `json:"supplyChannel,omitempty"`
	ShoppingList        *ShoppingListReference `json:"shoppingList"`
	DistributionChannel *ChannelReference      `json:"distributionChannel,omitempty"`
}

func (obj CartAddShoppingListAction) MarshalJSON() ([]byte, error) {
	type Alias CartAddShoppingListAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShoppingList", Alias: (*Alias)(&obj)})
}

type CartApplyDeltaToCustomLineItemShippingDetailsTargetsAction struct {
	TargetsDelta     []ItemShippingTarget `json:"targetsDelta"`
	CustomLineItemID string               `json:"customLineItemId"`
}

func (obj CartApplyDeltaToCustomLineItemShippingDetailsTargetsAction) MarshalJSON() ([]byte, error) {
	type Alias CartApplyDeltaToCustomLineItemShippingDetailsTargetsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "applyDeltaToCustomLineItemShippingDetailsTargets", Alias: (*Alias)(&obj)})
}

type CartApplyDeltaToLineItemShippingDetailsTargetsAction struct {
	TargetsDelta []ItemShippingTarget `json:"targetsDelta"`
	LineItemID   string               `json:"lineItemId"`
}

func (obj CartApplyDeltaToLineItemShippingDetailsTargetsAction) MarshalJSON() ([]byte, error) {
	type Alias CartApplyDeltaToLineItemShippingDetailsTargetsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "applyDeltaToLineItemShippingDetailsTargets", Alias: (*Alias)(&obj)})
}

type CartChangeCustomLineItemMoneyAction struct {
	Money            *Money `json:"money"`
	CustomLineItemID string `json:"customLineItemId"`
}

func (obj CartChangeCustomLineItemMoneyAction) MarshalJSON() ([]byte, error) {
	type Alias CartChangeCustomLineItemMoneyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCustomLineItemMoney", Alias: (*Alias)(&obj)})
}

type CartChangeCustomLineItemQuantityAction struct {
	Quantity         float64 `json:"quantity"`
	CustomLineItemID string  `json:"customLineItemId"`
}

func (obj CartChangeCustomLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias CartChangeCustomLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCustomLineItemQuantity", Alias: (*Alias)(&obj)})
}

type CartChangeLineItemQuantityAction struct {
	Quantity           float64                     `json:"quantity"`
	LineItemID         string                      `json:"lineItemId"`
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ExternalPrice      *Money                      `json:"externalPrice,omitempty"`
}

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

func (obj CartChangeTaxRoundingModeAction) MarshalJSON() ([]byte, error) {
	type Alias CartChangeTaxRoundingModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxRoundingMode", Alias: (*Alias)(&obj)})
}

type CartDraft struct {
	TaxRoundingMode                  RoundingMode             `json:"taxRoundingMode,omitempty"`
	TaxMode                          TaxMode                  `json:"taxMode,omitempty"`
	TaxCalculationMode               TaxCalculationMode       `json:"taxCalculationMode,omitempty"`
	ShippingRateInput                ShippingRateInputDraft   `json:"shippingRateInput,omitempty"`
	ShippingMethod                   *ShippingMethodReference `json:"shippingMethod,omitempty"`
	ShippingAddress                  *Address                 `json:"shippingAddress,omitempty"`
	Origin                           CartOrigin               `json:"origin,omitempty"`
	Locale                           string                   `json:"locale,omitempty"`
	LineItems                        []LineItemDraft          `json:"lineItems,omitempty"`
	ItemShippingAddresses            []Address                `json:"itemShippingAddresses,omitempty"`
	InventoryMode                    InventoryMode            `json:"inventoryMode,omitempty"`
	ExternalTaxRateForShippingMethod *ExternalTaxRateDraft    `json:"externalTaxRateForShippingMethod,omitempty"`
	DeleteDaysAfterLastModification  int                      `json:"deleteDaysAfterLastModification,omitempty"`
	CustomerID                       string                   `json:"customerId,omitempty"`
	CustomerGroup                    *CustomerGroupReference  `json:"customerGroup,omitempty"`
	CustomerEmail                    string                   `json:"customerEmail,omitempty"`
	CustomLineItems                  []CustomLineItemDraft    `json:"customLineItems,omitempty"`
	Custom                           *CustomFieldsDraft       `json:"custom,omitempty"`
	Currency                         CurrencyCode             `json:"currency"`
	Country                          string                   `json:"country,omitempty"`
	BillingAddress                   *Address                 `json:"billingAddress,omitempty"`
	AnonymousID                      string                   `json:"anonymousId,omitempty"`
}

func (obj *CartDraft) UnmarshalJSON(data []byte) error {
	type Alias CartDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInput != nil {
		obj.ShippingRateInput = AbstractShippingRateInputDraftDiscriminatorMapping(obj.ShippingRateInput)
	}

	return nil
}

type CartOrigin string

const (
	CartOriginCustomer CartOrigin = "Customer"
	CartOriginMerchant CartOrigin = "Merchant"
)

type CartPagedQueryResponse struct {
	Total   int    `json:"total,omitempty"`
	Offset  int    `json:"offset"`
	Count   int    `json:"count"`
	Results []Cart `json:"results"`
}

type CartRecalculateAction struct {
	UpdateProductData bool `json:"updateProductData,omitempty"`
}

func (obj CartRecalculateAction) MarshalJSON() ([]byte, error) {
	type Alias CartRecalculateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "recalculate", Alias: (*Alias)(&obj)})
}

type CartReference struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
	Obj *Cart  `json:"obj,omitempty"`
}

func (obj CartReference) MarshalJSON() ([]byte, error) {
	type Alias CartReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "cart", Alias: (*Alias)(&obj)})
}

type CartRemoveCustomLineItemAction struct {
	CustomLineItemID string `json:"customLineItemId"`
}

func (obj CartRemoveCustomLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias CartRemoveCustomLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeCustomLineItem", Alias: (*Alias)(&obj)})
}

type CartRemoveDiscountCodeAction struct {
	DiscountCode *DiscountCodeReference `json:"discountCode"`
}

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

func (obj CartRemoveItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CartRemoveItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeItemShippingAddress", Alias: (*Alias)(&obj)})
}

type CartRemoveLineItemAction struct {
	ShippingDetailsToRemove *ItemShippingDetailsDraft   `json:"shippingDetailsToRemove,omitempty"`
	Quantity                float64                     `json:"quantity,omitempty"`
	LineItemID              string                      `json:"lineItemId"`
	ExternalTotalPrice      *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ExternalPrice           *Money                      `json:"externalPrice,omitempty"`
}

func (obj CartRemoveLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias CartRemoveLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLineItem", Alias: (*Alias)(&obj)})
}

type CartRemovePaymentAction struct {
	Payment *PaymentReference `json:"payment"`
}

func (obj CartRemovePaymentAction) MarshalJSON() ([]byte, error) {
	type Alias CartRemovePaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePayment", Alias: (*Alias)(&obj)})
}

type CartSetAnonymousIdAction struct {
	AnonymousID string `json:"anonymousId,omitempty"`
}

func (obj CartSetAnonymousIdAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetAnonymousIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAnonymousId", Alias: (*Alias)(&obj)})
}

type CartSetBillingAddressAction struct {
	Address *Address `json:"address,omitempty"`
}

func (obj CartSetBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddress", Alias: (*Alias)(&obj)})
}

type CartSetCartTotalTaxAction struct {
	ExternalTotalGross  *Money       `json:"externalTotalGross"`
	ExternalTaxPortions []TaxPortion `json:"externalTaxPortions,omitempty"`
}

func (obj CartSetCartTotalTaxAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCartTotalTaxAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCartTotalTax", Alias: (*Alias)(&obj)})
}

type CartSetCountryAction struct {
	Country CountryCode `json:"country,omitempty"`
}

func (obj CartSetCountryAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCountryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCountry", Alias: (*Alias)(&obj)})
}

type CartSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

func (obj CartSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type CartSetCustomLineItemCustomFieldAction struct {
	Value            interface{} `json:"value,omitempty"`
	Name             string      `json:"name"`
	CustomLineItemID string      `json:"customLineItemId"`
}

func (obj CartSetCustomLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomField", Alias: (*Alias)(&obj)})
}

type CartSetCustomLineItemCustomTypeAction struct {
	Type             *TypeReference  `json:"type,omitempty"`
	Fields           *FieldContainer `json:"fields,omitempty"`
	CustomLineItemID string          `json:"customLineItemId"`
}

func (obj CartSetCustomLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomType", Alias: (*Alias)(&obj)})
}

type CartSetCustomLineItemShippingDetailsAction struct {
	ShippingDetails  *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	CustomLineItemID string                    `json:"customLineItemId"`
}

func (obj CartSetCustomLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

type CartSetCustomLineItemTaxAmountAction struct {
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
	CustomLineItemID  string                  `json:"customLineItemId"`
}

func (obj CartSetCustomLineItemTaxAmountAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomLineItemTaxAmountAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemTaxAmount", Alias: (*Alias)(&obj)})
}

type CartSetCustomLineItemTaxRateAction struct {
	ExternalTaxRate  *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	CustomLineItemID string                `json:"customLineItemId"`
}

func (obj CartSetCustomLineItemTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomLineItemTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemTaxRate", Alias: (*Alias)(&obj)})
}

type CartSetCustomShippingMethodAction struct {
	TaxCategory        *TaxCategoryReference `json:"taxCategory,omitempty"`
	ShippingRate       *ShippingRateDraft    `json:"shippingRate"`
	ShippingMethodName string                `json:"shippingMethodName"`
	ExternalTaxRate    *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
}

func (obj CartSetCustomShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomShippingMethod", Alias: (*Alias)(&obj)})
}

type CartSetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

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

func (obj CartSetCustomerEmailAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomerEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerEmail", Alias: (*Alias)(&obj)})
}

type CartSetCustomerGroupAction struct {
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
}

func (obj CartSetCustomerGroupAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomerGroupAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerGroup", Alias: (*Alias)(&obj)})
}

type CartSetCustomerIdAction struct {
	CustomerID string `json:"customerId,omitempty"`
}

func (obj CartSetCustomerIdAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetCustomerIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerId", Alias: (*Alias)(&obj)})
}

type CartSetDeleteDaysAfterLastModificationAction struct {
	DeleteDaysAfterLastModification int `json:"deleteDaysAfterLastModification,omitempty"`
}

func (obj CartSetDeleteDaysAfterLastModificationAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetDeleteDaysAfterLastModificationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeleteDaysAfterLastModification", Alias: (*Alias)(&obj)})
}

type CartSetLineItemCustomFieldAction struct {
	Value      interface{} `json:"value,omitempty"`
	Name       string      `json:"name"`
	LineItemID string      `json:"lineItemId"`
}

func (obj CartSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

type CartSetLineItemCustomTypeAction struct {
	Type       *TypeReference  `json:"type,omitempty"`
	LineItemID string          `json:"lineItemId"`
	Fields     *FieldContainer `json:"fields,omitempty"`
}

func (obj CartSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

type CartSetLineItemPriceAction struct {
	LineItemID    string `json:"lineItemId"`
	ExternalPrice *Money `json:"externalPrice,omitempty"`
}

func (obj CartSetLineItemPriceAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemPrice", Alias: (*Alias)(&obj)})
}

type CartSetLineItemShippingDetailsAction struct {
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	LineItemID      string                    `json:"lineItemId"`
}

func (obj CartSetLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

type CartSetLineItemTaxAmountAction struct {
	LineItemID        string                  `json:"lineItemId"`
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
}

func (obj CartSetLineItemTaxAmountAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemTaxAmountAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemTaxAmount", Alias: (*Alias)(&obj)})
}

type CartSetLineItemTaxRateAction struct {
	LineItemID      string                `json:"lineItemId"`
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
}

func (obj CartSetLineItemTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemTaxRate", Alias: (*Alias)(&obj)})
}

type CartSetLineItemTotalPriceAction struct {
	LineItemID         string                      `json:"lineItemId"`
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
}

func (obj CartSetLineItemTotalPriceAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLineItemTotalPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemTotalPrice", Alias: (*Alias)(&obj)})
}

type CartSetLocaleAction struct {
	Locale string `json:"locale,omitempty"`
}

func (obj CartSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

type CartSetShippingAddressAction struct {
	Address *Address `json:"address,omitempty"`
}

func (obj CartSetShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddress", Alias: (*Alias)(&obj)})
}

type CartSetShippingMethodAction struct {
	ShippingMethod  *TypeReference        `json:"shippingMethod,omitempty"`
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
}

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

func (obj CartSetShippingMethodTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetShippingMethodTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingMethodTaxRate", Alias: (*Alias)(&obj)})
}

type CartSetShippingRateInputAction struct {
	ShippingRateInput ShippingRateInputDraft `json:"shippingRateInput,omitempty"`
}

func (obj CartSetShippingRateInputAction) MarshalJSON() ([]byte, error) {
	type Alias CartSetShippingRateInputAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingRateInput", Alias: (*Alias)(&obj)})
}
func (obj *CartSetShippingRateInputAction) UnmarshalJSON(data []byte) error {
	type Alias CartSetShippingRateInputAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInput != nil {
		obj.ShippingRateInput = AbstractShippingRateInputDraftDiscriminatorMapping(obj.ShippingRateInput)
	}

	return nil
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

func (obj *CartUpdate) UnmarshalJSON(data []byte) error {
	type Alias CartUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractCartUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type CartUpdateAction interface{}
type AbstractCartUpdateAction struct{}

func AbstractCartUpdateActionDiscriminatorMapping(input CartUpdateAction) CartUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addCustomLineItem":
		new := CartAddCustomLineItemAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addDiscountCode":
		new := CartAddDiscountCodeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addItemShippingAddress":
		new := CartAddItemShippingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addLineItem":
		new := CartAddLineItemAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addPayment":
		new := CartAddPaymentAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addShoppingList":
		new := CartAddShoppingListAction{}
		mapstructure.Decode(input, &new)
		return new
	case "applyDeltaToCustomLineItemShippingDetailsTargets":
		new := CartApplyDeltaToCustomLineItemShippingDetailsTargetsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "applyDeltaToLineItemShippingDetailsTargets":
		new := CartApplyDeltaToLineItemShippingDetailsTargetsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeCustomLineItemMoney":
		new := CartChangeCustomLineItemMoneyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeCustomLineItemQuantity":
		new := CartChangeCustomLineItemQuantityAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeLineItemQuantity":
		new := CartChangeLineItemQuantityAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeTaxCalculationMode":
		new := CartChangeTaxCalculationModeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeTaxMode":
		new := CartChangeTaxModeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeTaxRoundingMode":
		new := CartChangeTaxRoundingModeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "recalculate":
		new := CartRecalculateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeCustomLineItem":
		new := CartRemoveCustomLineItemAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeDiscountCode":
		new := CartRemoveDiscountCodeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeItemShippingAddress":
		new := CartRemoveItemShippingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeLineItem":
		new := CartRemoveLineItemAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removePayment":
		new := CartRemovePaymentAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAnonymousId":
		new := CartSetAnonymousIdAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setBillingAddress":
		new := CartSetBillingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCartTotalTax":
		new := CartSetCartTotalTaxAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCountry":
		new := CartSetCountryAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomField":
		new := CartSetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomLineItemCustomField":
		new := CartSetCustomLineItemCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomLineItemCustomType":
		new := CartSetCustomLineItemCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomLineItemShippingDetails":
		new := CartSetCustomLineItemShippingDetailsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomLineItemTaxAmount":
		new := CartSetCustomLineItemTaxAmountAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomLineItemTaxRate":
		new := CartSetCustomLineItemTaxRateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomShippingMethod":
		new := CartSetCustomShippingMethodAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomType":
		new := CartSetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomerEmail":
		new := CartSetCustomerEmailAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomerGroup":
		new := CartSetCustomerGroupAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomerId":
		new := CartSetCustomerIdAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDeleteDaysAfterLastModification":
		new := CartSetDeleteDaysAfterLastModificationAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemCustomField":
		new := CartSetLineItemCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemCustomType":
		new := CartSetLineItemCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemPrice":
		new := CartSetLineItemPriceAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemShippingDetails":
		new := CartSetLineItemShippingDetailsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemTaxAmount":
		new := CartSetLineItemTaxAmountAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemTaxRate":
		new := CartSetLineItemTaxRateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLineItemTotalPrice":
		new := CartSetLineItemTotalPriceAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLocale":
		new := CartSetLocaleAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setShippingAddress":
		new := CartSetShippingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setShippingMethod":
		new := CartSetShippingMethodAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setShippingMethodTaxAmount":
		new := CartSetShippingMethodTaxAmountAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setShippingMethodTaxRate":
		new := CartSetShippingMethodTaxRateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setShippingRateInput":
		new := CartSetShippingRateInputAction{}
		mapstructure.Decode(input, &new)
		if new.ShippingRateInput != nil {
			new.ShippingRateInput = AbstractShippingRateInputDraftDiscriminatorMapping(new.ShippingRateInput)
		}

		return new
	case "updateItemShippingAddress":
		new := CartUpdateItemShippingAddressAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

type CartUpdateItemShippingAddressAction struct {
	Address *Address `json:"address"`
}

func (obj CartUpdateItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias CartUpdateItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateItemShippingAddress", Alias: (*Alias)(&obj)})
}

type ClassificationShippingRateInput struct {
	Label *LocalizedString `json:"label"`
	Key   string           `json:"key"`
}

func (obj ClassificationShippingRateInput) MarshalJSON() ([]byte, error) {
	type Alias ClassificationShippingRateInput
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "Classification", Alias: (*Alias)(&obj)})
}

type ClassificationShippingRateInputDraft struct {
	Key string `json:"key"`
}

func (obj ClassificationShippingRateInputDraft) MarshalJSON() ([]byte, error) {
	type Alias ClassificationShippingRateInputDraft
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "Classification", Alias: (*Alias)(&obj)})
}

type CustomLineItem struct {
	TotalPrice                 TypedMoney                           `json:"totalPrice"`
	TaxedPrice                 *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
	TaxRate                    *TaxRate                             `json:"taxRate,omitempty"`
	TaxCategory                *TaxCategoryReference                `json:"taxCategory,omitempty"`
	State                      []ItemState                          `json:"state"`
	Slug                       string                               `json:"slug"`
	ShippingDetails            *ItemShippingDetails                 `json:"shippingDetails,omitempty"`
	Quantity                   float64                              `json:"quantity"`
	Name                       *LocalizedString                     `json:"name"`
	Money                      TypedMoney                           `json:"money"`
	ID                         string                               `json:"id"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	Custom                     *CustomFields                        `json:"custom,omitempty"`
}

func (obj *CustomLineItem) UnmarshalJSON(data []byte) error {
	type Alias CustomLineItem
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Money != nil {
		obj.Money = AbstractTypedMoneyDiscriminatorMapping(obj.Money)
	}
	if obj.TotalPrice != nil {
		obj.TotalPrice = AbstractTypedMoneyDiscriminatorMapping(obj.TotalPrice)
	}

	return nil
}

type CustomLineItemDraft struct {
	TaxCategory     *TaxCategoryReference     `json:"taxCategory,omitempty"`
	Slug            string                    `json:"slug"`
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	Quantity        float64                   `json:"quantity"`
	Name            *LocalizedString          `json:"name"`
	Money           *Money                    `json:"money"`
	ExternalTaxRate *ExternalTaxRateDraft     `json:"externalTaxRate,omitempty"`
	Custom          *CustomFields             `json:"custom,omitempty"`
}

type DiscountCodeInfo struct {
	State        DiscountCodeState      `json:"state"`
	DiscountCode *DiscountCodeReference `json:"discountCode"`
}
type DiscountCodeState string

const (
	DiscountCodeStateNotActive             DiscountCodeState = "NotActive"
	DiscountCodeStateDoesNotMatchCart      DiscountCodeState = "DoesNotMatchCart"
	DiscountCodeStateMatchesCart           DiscountCodeState = "MatchesCart"
	DiscountCodeStateMaxApplicationReached DiscountCodeState = "MaxApplicationReached"
)

type DiscountedLineItemPortion struct {
	DiscountedAmount *Money                 `json:"discountedAmount"`
	Discount         *CartDiscountReference `json:"discount"`
}

type DiscountedLineItemPrice struct {
	Value             TypedMoney                  `json:"value"`
	IncludedDiscounts []DiscountedLineItemPortion `json:"includedDiscounts"`
}

func (obj *DiscountedLineItemPrice) UnmarshalJSON(data []byte) error {
	type Alias DiscountedLineItemPrice
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		obj.Value = AbstractTypedMoneyDiscriminatorMapping(obj.Value)
	}

	return nil
}

type DiscountedLineItemPriceForQuantity struct {
	Quantity        float64                  `json:"quantity"`
	DiscountedPrice *DiscountedLineItemPrice `json:"discountedPrice"`
}

type ExternalLineItemTotalPrice struct {
	TotalPrice *Money `json:"totalPrice"`
	Price      *Money `json:"price"`
}

type ExternalTaxAmountDraft struct {
	TotalGross *Money                `json:"totalGross"`
	TaxRate    *ExternalTaxRateDraft `json:"taxRate"`
}

type ExternalTaxRateDraft struct {
	SubRates []SubRate `json:"subRates,omitempty"`
	State    string    `json:"state,omitempty"`
	Name     string    `json:"name"`
	Country  string    `json:"country"`
	Amount   float64   `json:"amount,omitempty"`
}
type InventoryMode string

const (
	InventoryModeTrackOnly      InventoryMode = "TrackOnly"
	InventoryModeReserveOnOrder InventoryMode = "ReserveOnOrder"
	InventoryModeNone           InventoryMode = "None"
)

type ItemShippingDetails struct {
	Valid   bool                 `json:"valid"`
	Targets []ItemShippingTarget `json:"targets"`
}

type ItemShippingDetailsDraft struct {
	Targets []ItemShippingTarget `json:"targets"`
}

type ItemShippingTarget struct {
	Quantity   float64 `json:"quantity"`
	AddressKey string  `json:"addressKey"`
}

type LineItem struct {
	Variant                    *ProductVariant                      `json:"variant"`
	TotalPrice                 *Money                               `json:"totalPrice"`
	TaxedPrice                 *TaxedItemPrice                      `json:"taxedPrice,omitempty"`
	TaxRate                    *TaxRate                             `json:"taxRate,omitempty"`
	SupplyChannel              *ChannelReference                    `json:"supplyChannel,omitempty"`
	State                      []ItemState                          `json:"state"`
	ShippingDetails            *ItemShippingDetails                 `json:"shippingDetails,omitempty"`
	Quantity                   int                                  `json:"quantity"`
	ProductType                *ProductTypeReference                `json:"productType"`
	ProductSlug                *LocalizedString                     `json:"productSlug,omitempty"`
	ProductID                  string                               `json:"productId"`
	PriceMode                  LineItemPriceMode                    `json:"priceMode"`
	Price                      *Price                               `json:"price"`
	Name                       *LocalizedString                     `json:"name"`
	LineItemMode               LineItemMode                         `json:"lineItemMode"`
	ID                         string                               `json:"id"`
	DistributionChannel        *ChannelReference                    `json:"distributionChannel,omitempty"`
	DiscountedPricePerQuantity []DiscountedLineItemPriceForQuantity `json:"discountedPricePerQuantity"`
	Custom                     *CustomFields                        `json:"custom,omitempty"`
}

type LineItemDraft struct {
	VariantID           int                         `json:"variantId,omitempty"`
	SupplyChannel       *ChannelReference           `json:"supplyChannel,omitempty"`
	Sku                 string                      `json:"sku,omitempty"`
	ShippingDetails     *ItemShippingDetailsDraft   `json:"shippingDetails,omitempty"`
	Quantity            int                         `json:"quantity,omitempty"`
	ProductID           string                      `json:"productId,omitempty"`
	ExternalTotalPrice  *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ExternalTaxRate     *ExternalTaxRateDraft       `json:"externalTaxRate,omitempty"`
	ExternalPrice       *Money                      `json:"externalPrice,omitempty"`
	DistributionChannel *ChannelReference           `json:"distributionChannel,omitempty"`
	Custom              *CustomFieldsDraft          `json:"custom,omitempty"`
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

type ProductPublishScope string

const (
	ProductPublishScopeAll    ProductPublishScope = "All"
	ProductPublishScopePrices ProductPublishScope = "Prices"
)

type ReplicaCartDraft struct {
	Reference Reference `json:"reference"`
}

func (obj *ReplicaCartDraft) UnmarshalJSON(data []byte) error {
	type Alias ReplicaCartDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Reference != nil {
		obj.Reference = AbstractReferenceDiscriminatorMapping(obj.Reference)
	}

	return nil
}

type RoundingMode string

const (
	RoundingModeHalfEven RoundingMode = "HalfEven"
	RoundingModeHalfUp   RoundingMode = "HalfUp"
	RoundingModeHalfDown RoundingMode = "HalfDown"
)

type ScoreShippingRateInput struct {
	Score float64 `json:"score"`
}

func (obj ScoreShippingRateInput) MarshalJSON() ([]byte, error) {
	type Alias ScoreShippingRateInput
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "Score", Alias: (*Alias)(&obj)})
}

type ScoreShippingRateInputDraft struct {
	Score float64 `json:"score"`
}

func (obj ScoreShippingRateInputDraft) MarshalJSON() ([]byte, error) {
	type Alias ScoreShippingRateInputDraft
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "Score", Alias: (*Alias)(&obj)})
}

type ShippingInfo struct {
	TaxedPrice          *TaxedItemPrice          `json:"taxedPrice,omitempty"`
	TaxRate             *TaxRate                 `json:"taxRate,omitempty"`
	TaxCategory         *TaxCategoryReference    `json:"taxCategory,omitempty"`
	ShippingRate        *ShippingRate            `json:"shippingRate"`
	ShippingMethodState ShippingMethodState      `json:"shippingMethodState"`
	ShippingMethodName  string                   `json:"shippingMethodName"`
	ShippingMethod      *ShippingMethodReference `json:"shippingMethod,omitempty"`
	Price               TypedMoney               `json:"price"`
	DiscountedPrice     *DiscountedLineItemPrice `json:"discountedPrice,omitempty"`
	Deliveries          []Delivery               `json:"deliveries,omitempty"`
}

func (obj *ShippingInfo) UnmarshalJSON(data []byte) error {
	type Alias ShippingInfo
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Price != nil {
		obj.Price = AbstractTypedMoneyDiscriminatorMapping(obj.Price)
	}

	return nil
}

type ShippingMethodState string

const (
	ShippingMethodStateDoesNotMatchCart ShippingMethodState = "DoesNotMatchCart"
	ShippingMethodStateMatchesCart      ShippingMethodState = "MatchesCart"
)

type ShippingRateInput interface{}
type AbstractShippingRateInput struct{}

func AbstractShippingRateInputDiscriminatorMapping(input ShippingRateInput) ShippingRateInput {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "Classification":
		new := ClassificationShippingRateInput{}
		mapstructure.Decode(input, &new)
		return new
	case "Score":
		new := ScoreShippingRateInput{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

type ShippingRateInputDraft interface{}
type AbstractShippingRateInputDraft struct{}

func AbstractShippingRateInputDraftDiscriminatorMapping(input ShippingRateInputDraft) ShippingRateInputDraft {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "Classification":
		new := ClassificationShippingRateInputDraft{}
		mapstructure.Decode(input, &new)
		return new
	case "Score":
		new := ScoreShippingRateInputDraft{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
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
	Rate   float64 `json:"rate"`
	Name   string  `json:"name,omitempty"`
	Amount *Money  `json:"amount"`
}

type TaxedItemPrice struct {
	TotalNet   TypedMoney `json:"totalNet"`
	TotalGross TypedMoney `json:"totalGross"`
}

func (obj *TaxedItemPrice) UnmarshalJSON(data []byte) error {
	type Alias TaxedItemPrice
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.TotalGross != nil {
		obj.TotalGross = AbstractTypedMoneyDiscriminatorMapping(obj.TotalGross)
	}
	if obj.TotalNet != nil {
		obj.TotalNet = AbstractTypedMoneyDiscriminatorMapping(obj.TotalNet)
	}

	return nil
}

type TaxedPrice struct {
	TotalNet    *Money       `json:"totalNet"`
	TotalGross  *Money       `json:"totalGross"`
	TaxPortions []TaxPortion `json:"taxPortions"`
}
