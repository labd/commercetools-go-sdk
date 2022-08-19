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
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// The Quote Request related to this Quote.
	QuoteRequest QuoteRequestReference `json:"quoteRequest"`
	// The Staged Quote related to this Quote.
	StagedQuote StagedQuoteReference `json:"stagedQuote"`
	// The [Buyer](/../api/quotes-overview#buyer) who requested this Quote.
	Customer *CustomerReference `json:"customer,omitempty"`
	// Set automatically when `customer` is set and the Customer is a member of a Customer Group.
	// Used for Product Variant price selection.
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// Expiration date for the Quote.
	ValidTo *time.Time `json:"validTo,omitempty"`
	// The text message included in the offer from the [Seller](/../api/quotes-overview#seller).
	SellerComment *string `json:"sellerComment,omitempty"`
	// The Store to which the [Buyer](/../api/quotes-overview#buyer) belongs.
	Store *StoreKeyReference `json:"store,omitempty"`
	// The Line Items for which the Quote is requested.
	LineItems []LineItem `json:"lineItems"`
	// The Custom Line Items for which the Quote is requested.
	CustomLineItems []CustomLineItem `json:"customLineItems"`
	// The sum of all `totalPrice` fields of the `lineItems` and `customLineItems`, as well as the `price` field of `shippingInfo` (if it exists).
	// `totalPrice` may or may not include the taxes: it depends on the taxRate.includedInPrice property of each price.
	TotalPrice TypedMoney `json:"totalPrice"`
	// Not set until the shipping address is set.
	// Will be set automatically in the `Platform` TaxMode.
	// For the `External` tax mode it will be set  as soon as the external tax rates for all line items, custom line items, and shipping in the cart are set.
	TaxedPrice *TaxedPrice `json:"taxedPrice,omitempty"`
	// Used to determine the eligible [ShippingMethods](ctp:api:type:ShippingMethod)
	// and rates as well as the tax rate of the Line Items.
	ShippingAddress *Address `json:"shippingAddress,omitempty"`
	// The address used for invoicing.
	BillingAddress *Address `json:"billingAddress,omitempty"`
	// The inventory mode of the Cart referenced in the [QuoteRequestDraft](ctp:api:type:QuoteRequestDraft).
	InventoryMode *InventoryMode `json:"inventoryMode,omitempty"`
	// The tax mode of the Cart referenced in the [QuoteRequestDraft](ctp:api:type:QuoteRequestDraft).
	TaxMode TaxMode `json:"taxMode"`
	// When calculating taxes for `taxedPrice`, the selected mode is used for rounding.
	TaxRoundingMode RoundingMode `json:"taxRoundingMode"`
	// When calculating taxes for `taxedPrice`, the selected mode is used for calculating the price with `LineItemLevel` (horizontally) or `UnitPriceLevel` (vertically) calculation mode.
	TaxCalculationMode TaxCalculationMode `json:"taxCalculationMode"`
	// Used for Product Variant price selection.
	Country *string `json:"country,omitempty"`
	// Set automatically once the [ShippingMethod](ctp:api:type:ShippingMethod) is set.
	ShippingInfo *ShippingInfo `json:"shippingInfo,omitempty"`
	// Log of payment transactions related to this quote.
	PaymentInfo *PaymentInfo `json:"paymentInfo,omitempty"`
	// Used to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier).
	ShippingRateInput ShippingRateInput `json:"shippingRateInput,omitempty"`
	// Contains addresses for carts with multiple shipping addresses.
	// Line items reference these addresses under their `shippingDetails`.
	// The addresses captured here are not used to determine eligible shipping methods or the applicable tax rate.
	// Only the cart's `shippingAddress` is used for this.
	ItemShippingAddresses []Address `json:"itemShippingAddresses"`
	// Discounts only valid for this Quote, those cannot be associated to any other Cart or Order.
	DirectDiscounts []DirectDiscount `json:"directDiscounts"`
	// Custom Fields of this Quote.
	Custom *CustomFields `json:"custom,omitempty"`
	// [State](ctp:api:type:State) of the Quote.
	// This reference can point to a State in a custom workflow.
	State *StateReference `json:"state,omitempty"`
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
	// The StagedQuote from which this Quote is created.
	StagedQuote StagedQuoteResourceIdentifier `json:"stagedQuote"`
	// Current version of the StagedQuote.
	StagedQuoteVersion int `json:"stagedQuoteVersion"`
	// User-defined unique identifier for the Quote.
	Key *string `json:"key,omitempty"`
	// [Custom Fields](/../api/projects/custom-fields) to be added to the Quote.
	//
	// - If specified, the Custom Fields are merged with the Custom Fields on the referenced [StagedQuote](/../api/projects/staged-quotes#stagedquote) and added to the Quote.
	// - If empty, the Custom Fields on the referenced [StagedQuote](/../api/projects/staged-quotes#stagedquote) are added to the Quote automatically.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// [State](ctp:api:type:State) of the Quote.
	// This reference can point to a State in a custom workflow.
	State *StateReference `json:"state,omitempty"`
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
	QuoteStatePending   QuoteState = "Pending"
	QuoteStateDeclined  QuoteState = "Declined"
	QuoteStateAccepted  QuoteState = "Accepted"
	QuoteStateFailed    QuoteState = "Failed"
	QuoteStateWithdrawn QuoteState = "Withdrawn"
)

type QuoteUpdate struct {
	Version int                 `json:"version"`
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
	case "changeQuoteState":
		obj := QuoteChangeQuoteStateAction{}
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

type QuoteChangeQuoteStateAction struct {
	// The new quote state to be set for the Quote.
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

type QuoteSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
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
