package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type StagedQuote struct {
	// The unique ID of the StagedQuote.
	ID string `json:"id"`
	// Current version of the StagedQuote.
	Version int `json:"version"`
	// Date and time (UTC) the StagedQuote was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the StagedQuote was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// User-specific unique identifier of the staged quote.
	Key *string `json:"key,omitempty"`
	// IDs and references that last modified the StagedQuote.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the StagedQuote.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Predefined states tracking the status of the Staged Quote.
	StagedQuoteState StagedQuoteState `json:"stagedQuoteState"`
	// The [Buyer](/../api/quotes-overview#buyer) who requested the Quote.
	Customer *CustomerReference `json:"customer,omitempty"`
	// Quote Request related to the Staged Quote.
	QuoteRequest QuoteRequestReference `json:"quoteRequest"`
	// [Cart](ctp:api:type:Cart) containing the offered items. May contain either [DirectDiscounts](ctp:api:type:DirectDiscount) or [CartDiscounts](ctp:api:type:CartDiscount).
	QuotationCart CartReference `json:"quotationCart"`
	// Expiration date for the Quote.
	ValidTo *time.Time `json:"validTo,omitempty"`
	// Message from the [Seller](/../api/quotes-overview#seller) included in the offer.
	SellerComment *string `json:"sellerComment,omitempty"`
	// Custom Fields of the Staged Quote.
	Custom *CustomFields `json:"custom,omitempty"`
	// [State](ctp:api:type:State) of the Staged Quote.
	// This reference can point to a State in a custom workflow.
	State *StateReference `json:"state,omitempty"`
	// The Purchase Order Number is typically set by the [Buyer](/quotes-overview#buyer) on a [QuoteRequest](ctp:api:type:QuoteRequest) to
	// track the purchase order during the [quote and order flow](/../api/quotes-overview#intended-workflow).
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
	// The [BusinessUnit](ctp:api:type:BusinessUnit) for the Staged Quote.
	BusinessUnit *BusinessUnitKeyReference `json:"businessUnit,omitempty"`
	// The Store to which the [Buyer](/../api/quotes-overview#buyer) belongs.
	Store *StoreKeyReference `json:"store,omitempty"`
}

type StagedQuoteDraft struct {
	// QuoteRequest from which the StagedQuote is created.
	QuoteRequest QuoteRequestResourceIdentifier `json:"quoteRequest"`
	// Current version of the QuoteRequest.
	QuoteRequestVersion int `json:"quoteRequestVersion"`
	// If `true`, the `quoteRequestState` of the referenced [QuoteRequest](ctp:api:type:QuoteRequest) will be set to `Accepted`.
	QuoteRequestStateToAccepted *bool `json:"quoteRequestStateToAccepted,omitempty"`
	// User-defined unique identifier for the StagedQuote.
	Key *string `json:"key,omitempty"`
	// [Custom Fields](/../api/projects/custom-fields) to be added to the StagedQuote.
	//
	// - If specified, the Custom Fields are merged with the Custom Fields on the referenced [QuoteRequest](ctp:api:type:QuoteRequest) and added to the StagedQuote.
	// - If empty, the Custom Fields on the referenced [QuoteRequest](ctp:api:type:QuoteRequest) are added to the StagedQuote automatically.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// [State](ctp:api:type:State) of the Staged Quote.
	// This reference can point to a State in a custom workflow.
	State *StateReference `json:"state,omitempty"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [StagedQuote](ctp:api:type:StagedQuote).
*
 */
type StagedQuotePagedQueryResponse struct {
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
	// Staged Quotes matching the query.
	Results []StagedQuote `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [StagedQuote](ctp:api:type:StagedQuote).
*
 */
type StagedQuoteReference struct {
	// Unique ID of the referenced resource.
	ID string `json:"id"`
	// Contains the representation of the expanded StagedQuote.
	// Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for StagedQuote.
	Obj *StagedQuote `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedQuoteReference) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "staged-quote", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [StagedQuote](ctp:api:type:StagedQuote).
*
 */
type StagedQuoteResourceIdentifier struct {
	// Unique identifier of the referenced resource. Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced resource. Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedQuoteResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "staged-quote", Alias: (*Alias)(&obj)})
}

/**
*	Predefined states tracking the status of the Staged Quote.
*
 */
type StagedQuoteState string

const (
	StagedQuoteStateInProgress StagedQuoteState = "InProgress"
	StagedQuoteStateSent       StagedQuoteState = "Sent"
	StagedQuoteStateClosed     StagedQuoteState = "Closed"
)

type StagedQuoteUpdate struct {
	// Expected version of the [StagedQuote](ctp:api:type:StagedQuote) to which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the [StagedQuote](ctp:api:type:StagedQuote).
	Actions []StagedQuoteUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StagedQuoteUpdate) UnmarshalJSON(data []byte) error {
	type Alias StagedQuoteUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorStagedQuoteUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type StagedQuoteUpdateAction interface{}

func mapDiscriminatorStagedQuoteUpdateAction(input interface{}) (StagedQuoteUpdateAction, error) {
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
	case "changeStagedQuoteState":
		obj := StagedQuoteChangeStagedQuoteStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := StagedQuoteSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := StagedQuoteSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSellerComment":
		obj := StagedQuoteSetSellerCommentAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidTo":
		obj := StagedQuoteSetValidToAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "transitionState":
		obj := StagedQuoteTransitionStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type StagedQuoteChangeStagedQuoteStateAction struct {
	// New state to be set for the Staged Quote.
	StagedQuoteState StagedQuoteState `json:"stagedQuoteState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedQuoteChangeStagedQuoteStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteChangeStagedQuoteStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeStagedQuoteState", Alias: (*Alias)(&obj)})
}

type StagedQuoteSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedQuoteSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type StagedQuoteSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the StagedQuote with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the StagedQuote.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the StagedQuote.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedQuoteSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type StagedQuoteSetSellerCommentAction struct {
	// If `sellerComment` is absent or `null`, this field will be removed if it exists.
	SellerComment *string `json:"sellerComment,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedQuoteSetSellerCommentAction) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteSetSellerCommentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSellerComment", Alias: (*Alias)(&obj)})
}

type StagedQuoteSetValidToAction struct {
	// If `validTo` is absent or `null`, this field will be removed if it exists.
	ValidTo *time.Time `json:"validTo,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedQuoteSetValidToAction) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteSetValidToAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidTo", Alias: (*Alias)(&obj)})
}

/**
*	If the existing [State](ctp:api:type:State) has set `transitions`, there must be a direct transition to the new State. If `transitions` is not set, no validation is performed. This update action produces the [Staged Quote State Transition](ctp:api:type:StagedQuoteStateTransitionMessage) Message.
*
 */
type StagedQuoteTransitionStateAction struct {
	// Value to set.
	// If there is no State yet, the new State must be an initial State.
	State StateResourceIdentifier `json:"state"`
	// Switch validations on or off.
	Force *bool `json:"force,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedQuoteTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedQuoteTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}
