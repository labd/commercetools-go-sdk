package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type Quote struct {
	// Unique identifier of the Quote.
	ID string `json:"id"`
	// Current version of the Quote.
	Version int `json:"version"`
	// Date and time (UTC) the Quote was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Quote was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// User-defined unique identifier of the Quote.
	Key *string `json:"key,omitempty"`
	// IDs and references that last modified the Quote.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the Quote.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Quote Request related to the Quote.
	QuoteRequest QuoteRequestReference `json:"quoteRequest"`
	// Staged Quote related to the Quote.
	StagedQuote StagedQuoteReference `json:"stagedQuote"`
	// The [Buyer](/../api/quotes-overview#buyer) who owns the Quote.
	Customer *CustomerReference `json:"customer,omitempty"`
	// Set automatically when `customer` is set and the Customer is a member of a Customer Group.
	// Not updated if Customer is changed after Quote creation.
	// Used for Product Variant price selection.
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// Expiration date for the Quote.
	ValidTo *time.Time `json:"validTo,omitempty"`
	// Message from the [Seller](/../api/quotes-overview#seller) included in the offer.
	SellerComment *string `json:"sellerComment,omitempty"`
	// Message from the [Buyer](/../api/quotes-overview#buyer) included in the [renegotiation request](ctp:api:type:QuoteRequestQuoteRenegotiationAction).
	BuyerComment *string `json:"buyerComment,omitempty"`
	// The Store to which the [Buyer](/../api/quotes-overview#buyer) belongs.
	Store *StoreKeyReference `json:"store,omitempty"`
	// The Line Items for which the Quote is requested.
	LineItems []LineItem `json:"lineItems"`
	// The Custom Line Items for which the Quote is requested.
	CustomLineItems []CustomLineItem `json:"customLineItems"`
	// Sum of all `totalPrice` fields of the `lineItems` and `customLineItems`, as well as the `price` field of `shippingInfo` (if it exists).
	// `totalPrice` may or may not include the taxes: it depends on the taxRate.includedInPrice property of each price.
	TotalPrice TypedMoney `json:"totalPrice"`
	// Not set until the shipping address is set.
	// Will be set automatically in the `Platform` TaxMode.
	// For the `External` tax mode it will be set  as soon as the external tax rates for all line items, custom line items, and shipping in the cart are set.
	TaxedPrice *TaxedPrice `json:"taxedPrice,omitempty"`
	// Used to determine the eligible [ShippingMethods](ctp:api:type:ShippingMethod)
	// and rates as well as the tax rate of the Line Items.
	ShippingAddress *Address `json:"shippingAddress,omitempty"`
	// Address used for invoicing.
	BillingAddress *Address `json:"billingAddress,omitempty"`
	// Inventory mode of the Cart referenced in the [QuoteRequestDraft](ctp:api:type:QuoteRequestDraft).
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	// Tax mode of the Cart referenced in the [QuoteRequestDraft](ctp:api:type:QuoteRequestDraft).
	TaxMode TaxMode `json:"taxMode"`
	// When calculating taxes for `taxedPrice`, the selected mode is used for rounding.
	TaxRoundingMode RoundingMode `json:"taxRoundingMode"`
	// When calculating taxes for `taxedPrice`, the selected mode is used for calculating the price with `LineItemLevel` (horizontally) or `UnitPriceLevel` (vertically) calculation mode.
	TaxCalculationMode TaxCalculationMode `json:"taxCalculationMode"`
	// Used for Product Variant price selection.
	Country *string `json:"country,omitempty"`
	// Set automatically once the [ShippingMethod](ctp:api:type:ShippingMethod) is set.
	ShippingInfo *ShippingInfo `json:"shippingInfo,omitempty"`
	// Log of payment transactions related to the Quote.
	PaymentInfo *PaymentInfo `json:"paymentInfo,omitempty"`
	// Used to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	ShippingRateInput ShippingRateInput `json:"shippingRateInput,omitempty"`
	// Contains addresses for carts with multiple shipping addresses.
	// Line items reference these addresses under their `shippingDetails`.
	// The addresses captured here are not used to determine eligible shipping methods or the applicable tax rate.
	// Only the cart's `shippingAddress` is used for this.
	ItemShippingAddresses []Address `json:"itemShippingAddresses"`
	// Discounts that are only valid for the Quote and cannot be associated to any other Cart or Order.
	DirectDiscounts []DirectDiscount `json:"directDiscounts"`
	// Custom Fields on the Quote.
	Custom *CustomFields `json:"custom,omitempty"`
	// Predefined states tracking the status of the Quote.
	QuoteState QuoteState `json:"quoteState"`
	// [State](ctp:api:type:State) of the Quote.
	// This reference can point to a State in a custom workflow.
	State *StateReference `json:"state,omitempty"`
	// The Purchase Order Number is typically set by the [Buyer](/quotes-overview#buyer) on a [QuoteRequest](ctp:api:type:QuoteRequest) to
	// track the purchase order during the [quote and order flow](/../api/quotes-overview#intended-workflow).
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
	// The [BusinessUnit](ctp:api:type:BusinessUnit) for the Quote.
	BusinessUnit *BusinessUnitKeyReference `json:"businessUnit,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Quote) UnmarshalJSON(data []byte) error {
	type Alias Quote
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
func (obj Quote) MarshalJSON() ([]byte, error) {
	type Alias Quote
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

	if raw["directDiscounts"] == nil {
		delete(raw, "directDiscounts")
	}

	return json.Marshal(raw)

}

type QuoteDraft struct {
	// User-defined unique identifier for the Quote.
	Key *string `json:"key,omitempty"`
	// StagedQuote from which the Quote is created.
	StagedQuote StagedQuoteResourceIdentifier `json:"stagedQuote"`
	// Current version of the StagedQuote.
	StagedQuoteVersion int `json:"stagedQuoteVersion"`
	// If `true`, the `stagedQuoteState` of the referenced [StagedQuote](/../api/projects/staged-quotes#stagedquote) will be set to `Sent`.
	StagedQuoteStateToSent *bool `json:"stagedQuoteStateToSent,omitempty"`
	// [State](ctp:api:type:State) of the Quote.
	// This reference can point to a State in a custom workflow.
	State *StateReference `json:"state,omitempty"`
	// [Custom Fields](/../api/projects/custom-fields) to be added to the Quote.
	//
	// - If specified, the Custom Fields are merged with the Custom Fields on the referenced [StagedQuote](/../api/projects/staged-quotes#stagedquote) and added to the Quote.
	// - If empty, the Custom Fields on the referenced [StagedQuote](/../api/projects/staged-quotes#stagedquote) are added to the Quote automatically.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [Quote](ctp:api:type:Quote).
*
 */
type QuotePagedQueryResponse struct {
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
	// Quotes matching the query.
	Results []Quote `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [Quote](ctp:api:type:Quote).
*
 */
type QuoteReference struct {
	// Unique ID of the referenced resource.
	ID string `json:"id"`
	// Contains the representation of the expanded Quote.
	// Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for Quote.
	Obj *Quote `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteReference) MarshalJSON() ([]byte, error) {
	type Alias QuoteReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "quote", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Quote](ctp:api:type:Quote).
*
 */
type QuoteResourceIdentifier struct {
	// Unique identifier of the referenced resource. Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced resource. Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias QuoteResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "quote", Alias: (*Alias)(&obj)})
}

/**
*	Predefined states tracking the status of the Quote.
*
 */
type QuoteState string

const (
	QuoteStatePending                  QuoteState = "Pending"
	QuoteStateDeclined                 QuoteState = "Declined"
	QuoteStateDeclinedForRenegotiation QuoteState = "DeclinedForRenegotiation"
	QuoteStateRenegotiationAddressed   QuoteState = "RenegotiationAddressed"
	QuoteStateAccepted                 QuoteState = "Accepted"
	QuoteStateWithdrawn                QuoteState = "Withdrawn"
)

type QuoteUpdate struct {
	// Expected version of the [Quote](ctp:api:type:Quote) to which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the [Quote](ctp:api:type:Quote).
	Actions []QuoteUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *QuoteUpdate) UnmarshalJSON(data []byte) error {
	type Alias QuoteUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorQuoteUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type QuoteUpdateAction interface{}

func mapDiscriminatorQuoteUpdateAction(input interface{}) (QuoteUpdateAction, error) {
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
	case "changeCustomer":
		obj := QuoteChangeCustomerAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeQuoteState":
		obj := QuoteChangeQuoteStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "requestQuoteRenegotiation":
		obj := QuoteRequestQuoteRenegotiationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := QuoteSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := QuoteSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "transitionState":
		obj := QuoteTransitionStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Changes the owner of a Quote to a different Customer.
*	Customer Group is not updated.
*	This update action produces the [Quote Customer Changed](ctp:api:type:QuoteCustomerChangedMessage) Message.
*
 */
type QuoteChangeCustomerAction struct {
	// New Customer to own the Quote.
	Customer CustomerResourceIdentifier `json:"customer"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteChangeCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias QuoteChangeCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCustomer", Alias: (*Alias)(&obj)})
}

type QuoteChangeQuoteStateAction struct {
	// New state to be set for the Quote.
	QuoteState QuoteState `json:"quoteState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteChangeQuoteStateAction) MarshalJSON() ([]byte, error) {
	type Alias QuoteChangeQuoteStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeQuoteState", Alias: (*Alias)(&obj)})
}

/**
*	Represents the Buyer requesting renegotiation for a Quote. Valid for Quotes in a `Pending` [state](ctp:api:type:QuoteState).
*
 */
type QuoteRequestQuoteRenegotiationAction struct {
	// Message from the [Buyer](/api/quotes-overview#buyer) regarding the Quote renegotiation request.
	BuyerComment *string `json:"buyerComment,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteRequestQuoteRenegotiationAction) MarshalJSON() ([]byte, error) {
	type Alias QuoteRequestQuoteRenegotiationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "requestQuoteRenegotiation", Alias: (*Alias)(&obj)})
}

type QuoteSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias QuoteSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type QuoteSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the Quote with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Quote.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Quote.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias QuoteSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

/**
*	If the existing [State](ctp:api:type:State) has set `transitions`, there must be a direct transition to the new State. If `transitions` is not set, no validation is performed. This update action produces the [Quote State Transition](ctp:api:type:QuoteStateTransitionMessage) Message.
*
 */
type QuoteTransitionStateAction struct {
	// Value to set.
	// If there is no State yet, this must be an initial State.
	State StateResourceIdentifier `json:"state"`
	// Switch validations on or off.
	Force *bool `json:"force,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj QuoteTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias QuoteTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}
