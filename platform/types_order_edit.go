package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type OrderEdit struct {
	// Unique identifier of the Order Edit.
	ID string `json:"id"`
	// Current version of the Order Edit.
	Version int `json:"version"`
	// Date and time (UTC) the Order Edit was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Order Edit was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// User-defined unique identifier of the Order Edit.
	Key *string `json:"key,omitempty"`
	// [Reference](ctp:api:type:Reference) to the Order updated with this edit.
	Resource OrderReference `json:"resource"`
	// Update actions applied to the Order referenced by `resource`.
	StagedActions []StagedOrderUpdateAction `json:"stagedActions"`
	// For applied edits, it's a summary of the changes on the Order.
	// For unapplied edits, it's a preview of the changes.
	Result OrderEditResult `json:"result"`
	// User-defined information regarding the Order Edit.
	Comment *string `json:"comment,omitempty"`
	// Custom Fields of the Order Edit.
	Custom *CustomFields `json:"custom,omitempty"`
	// IDs and references that last modified the OrderEdit.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the OrderEdit.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEdit) UnmarshalJSON(data []byte) error {
	type Alias OrderEdit
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.StagedActions {
		var err error
		obj.StagedActions[i], err = mapDiscriminatorStagedOrderUpdateAction(obj.StagedActions[i])
		if err != nil {
			return err
		}
	}
	if obj.Result != nil {
		var err error
		obj.Result, err = mapDiscriminatorOrderEditResult(obj.Result)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	If the `editVersion` and/or `resourceVersion` do not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
*
 */
type OrderEditApply struct {
	// Current `version` of the OrderEdit to be applied.
	EditVersion int `json:"editVersion"`
	// Current `version` of the [Order](ctp:api:type:Order) to which the OrderEdit is applied.
	ResourceVersion int `json:"resourceVersion"`
}

type OrderEditDraft struct {
	// User-defined unique identifier for the Order Edit.
	Key *string `json:"key,omitempty"`
	// [Reference](ctp:api:type:Reference) to the Order updated with this edit.
	Resource OrderReference `json:"resource"`
	// Update actions to apply to the Order referenced in `resource`.
	// Cannot be updated if the [edit has been applied](ctp:api:endpoint:/{projectKey}/orders/edits/{id}/apply:POST).
	StagedActions []StagedOrderUpdateAction `json:"stagedActions"`
	// Custom Fields for the Order Edit.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// User-defined description regarding the Order Edit.
	Comment *string `json:"comment,omitempty"`
	// Set to `true` if you want to [preview](ctp:api:type:OrderEditPreviewSuccess) the edited Order first without persisting it (dry run).
	// A dry run allows checking for potential [errors](ctp:api:type:OrderEditPreviewFailure) when trying to apply the `stagedActions`.
	//
	// Order [API Extensions](/../api/projects/api-extensions), if any, are also called in dry runs.
	DryRun *bool `json:"dryRun,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditDraft) UnmarshalJSON(data []byte) error {
	type Alias OrderEditDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.StagedActions {
		var err error
		obj.StagedActions[i], err = mapDiscriminatorStagedOrderUpdateAction(obj.StagedActions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderEditDraft) MarshalJSON() ([]byte, error) {
	type Alias OrderEditDraft
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

	if raw["stagedActions"] == nil {
		delete(raw, "stagedActions")
	}

	return json.Marshal(raw)

}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [OrderEdit](ctp:api:type:OrderEdit).
*
 */
type OrderEditPagedQueryResponse struct {
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
	// [OrderEdits](ctp:api:type:OrderEdit) matching the query.
	Results []OrderEdit `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to an [OrderEdit](ctp:api:type:OrderEdit).
*
 */
type OrderEditReference struct {
	// Unique identifier of the referenced [OrderEdit](ctp:api:type:OrderEdit).
	ID string `json:"id"`
	// Contains the representation of the expanded Order Edit. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for Order Edits.
	Obj *OrderEdit `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderEditReference) MarshalJSON() ([]byte, error) {
	type Alias OrderEditReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "order-edit", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to an [OrderEdit](ctp:api:type:OrderEdit). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type OrderEditResourceIdentifier struct {
	// Unique identifier of the referenced [OrderEdit](ctp:api:type:OrderEdit). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [OrderEdit](ctp:api:type:OrderEdit). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderEditResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias OrderEditResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "order-edit", Alias: (*Alias)(&obj)})
}

type OrderEditResult interface{}

func mapDiscriminatorOrderEditResult(input interface{}) (OrderEditResult, error) {
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
	case "Applied":
		obj := OrderEditApplied{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "NotProcessed":
		obj := OrderEditNotProcessed{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "PreviewFailure":
		obj := OrderEditPreviewFailure{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		for i := range obj.Errors {
			var err error
			obj.Errors[i], err = mapDiscriminatorErrorObject(obj.Errors[i])
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "PreviewSuccess":
		obj := OrderEditPreviewSuccess{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		for i := range obj.MessagePayloads {
			var err error
			obj.MessagePayloads[i], err = mapDiscriminatorMessagePayload(obj.MessagePayloads[i])
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Result of a successful application of `stagedActions` to the Order.
*
 */
type OrderEditApplied struct {
	// Date and time (UTC) the Order was edited.
	AppliedAt time.Time `json:"appliedAt"`
	// Prices of the Order before the edit.
	ExcerptBeforeEdit OrderExcerpt `json:"excerptBeforeEdit"`
	// Prices of the Order after the edit.
	ExcerptAfterEdit OrderExcerpt `json:"excerptAfterEdit"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderEditApplied) MarshalJSON() ([]byte, error) {
	type Alias OrderEditApplied
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Applied", Alias: (*Alias)(&obj)})
}

/**
*	Indicates that the edit has not been applied or processed in any way.
*
 */
type OrderEditNotProcessed struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderEditNotProcessed) MarshalJSON() ([]byte, error) {
	type Alias OrderEditNotProcessed
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "NotProcessed", Alias: (*Alias)(&obj)})
}

/**
*	Result of a failed application of `stagedActions` to the Order. The data is calculated on the fly and is not queryable.
*
 */
type OrderEditPreviewFailure struct {
	// Errors returned.
	Errors []ErrorObject `json:"errors"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditPreviewFailure) UnmarshalJSON(data []byte) error {
	type Alias OrderEditPreviewFailure
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Errors {
		var err error
		obj.Errors[i], err = mapDiscriminatorErrorObject(obj.Errors[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderEditPreviewFailure) MarshalJSON() ([]byte, error) {
	type Alias OrderEditPreviewFailure
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PreviewFailure", Alias: (*Alias)(&obj)})
}

/**
*	The data is not persisted but is dynamically pulled by dry-running the update actions from `stagedActions` on the current version of the related [Order](ctp:api:type:Order), not from the Order version at the time the OrderEdit was created. Therefore, it cannot be queried.
*
 */
type OrderEditPreviewSuccess struct {
	// A preview of the edited [Order](ctp:api:type:Order) as it will be after all `stagedActions` (incl. optional Order [API Extensions](/../api/projects/api-extensions)) are applied.
	Preview StagedOrder `json:"preview"`
	// Messages that will be generated if the edit is applied.
	MessagePayloads []MessagePayload `json:"messagePayloads"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditPreviewSuccess) UnmarshalJSON(data []byte) error {
	type Alias OrderEditPreviewSuccess
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.MessagePayloads {
		var err error
		obj.MessagePayloads[i], err = mapDiscriminatorMessagePayload(obj.MessagePayloads[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderEditPreviewSuccess) MarshalJSON() ([]byte, error) {
	type Alias OrderEditPreviewSuccess
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "PreviewSuccess", Alias: (*Alias)(&obj)})
}

type OrderEditUpdate struct {
	// Expected version of the Order Edit on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the Order Edit.
	Actions []OrderEditUpdateAction `json:"actions"`
	// If set to `true`, the Order Edit is applied on the [Order](ctp:api:type:Order) without persisting it.
	DryRun *bool `json:"dryRun,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditUpdate) UnmarshalJSON(data []byte) error {
	type Alias OrderEditUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorOrderEditUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type OrderEditUpdateAction interface{}

func mapDiscriminatorOrderEditUpdateAction(input interface{}) (OrderEditUpdateAction, error) {
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
	case "addStagedAction":
		obj := OrderEditAddStagedActionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.StagedAction != nil {
			var err error
			obj.StagedAction, err = mapDiscriminatorStagedOrderUpdateAction(obj.StagedAction)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "setComment":
		obj := OrderEditSetCommentAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := OrderEditSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := OrderEditSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := OrderEditSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setStagedActions":
		obj := OrderEditSetStagedActionsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		for i := range obj.StagedActions {
			var err error
			obj.StagedActions[i], err = mapDiscriminatorStagedOrderUpdateAction(obj.StagedActions[i])
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Excerpt of the Order extracting the total and the taxed price.
*
 */
type OrderExcerpt struct {
	// Total price of the Order.
	TotalPrice TypedMoney `json:"totalPrice"`
	// Taxed price of the Order.
	TaxedPrice *TaxedPrice `json:"taxedPrice,omitempty"`
	// Current version of the Order.
	Version int `json:"version"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderExcerpt) UnmarshalJSON(data []byte) error {
	type Alias OrderExcerpt
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

type StagedOrder struct {
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
	// Used for [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection).
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// [Anonymous session](ctp:api:type:AnonymousSession) associated with the Order.
	AnonymousId *string `json:"anonymousId,omitempty"`
	// [Reference](ctp:api:type:Reference) to a Business Unit the Order belongs to. Only available for [B2B](/../offering/composable-commerce#composable-commerce-for-b2b)-enabled Projects.
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
	TotalPrice CentPrecisionMoney `json:"totalPrice"`
	// - For `Platform` [TaxMode](ctp:api:type:TaxMode), it is automatically set when a [shipping address is set](ctp:api:type:OrderSetShippingAddressAction).
	// - For `External` [TaxMode](ctp:api:type:TaxMode), it is automatically set when `shippingAddress` and external Tax Rates for all Line Items, Custom Line Items, and Shipping Methods in the Cart are set.
	//
	// If a discount applies on `totalPrice`, this field holds the proportionally discounted value.
	TaxedPrice *TaxedPrice `json:"taxedPrice,omitempty"`
	// Sum of the `taxedPrice` field of [ShippingInfo](ctp:api:type:ShippingInfo) across all Shipping Methods.
	//
	// If a discount applies on `totalPrice`, this field holds the proportionally discounted value.
	TaxedShippingPrice *TaxedPrice `json:"taxedShippingPrice,omitempty"`
	// Discounts that apply on the total price of the Order.
	DiscountOnTotalPrice *DiscountOnTotalPrice `json:"discountOnTotalPrice,omitempty"`
	// Indicates how the total prices on [LineItems](ctp:api:type:LineItem) and [CustomLineItems](ctp:api:type:CustomLineItem) are rounded when calculated.
	PriceRoundingMode *RoundingMode `json:"priceRoundingMode,omitempty"`
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
	// Used for [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection).
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
	// [Reference](ctp:api:type:Reference) to the RecurringOrder that generated this Order.
	RecurringOrder *RecurringOrderReference `json:"recurringOrder,omitempty"`
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
	// Indicates if a combination of discount types can apply on an Order.
	DiscountTypeCombination DiscountTypeCombination `json:"discountTypeCombination,omitempty"`
	// Internal-only field.
	LastMessageSequenceNumber *int `json:"lastMessageSequenceNumber,omitempty"`
	// Custom Fields of the Order.
	Custom *CustomFields `json:"custom,omitempty"`
	// User-defined date and time (UTC) of the Order.
	// Present only on an Order created using [Order Import](ctp:api:endpoint:/{projectKey}/orders/import:POST).
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	// IDs and references that last modified the Order.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the Order.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StagedOrder) UnmarshalJSON(data []byte) error {
	type Alias StagedOrder
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrder) MarshalJSON() ([]byte, error) {
	type Alias StagedOrder
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

/**
*	The `stagedActions` field cannot be updated if the Order Edit `result` is [OrderEdit Applied](/projects/order-edits#orderedit-applied).
*
 */
type OrderEditAddStagedActionAction struct {
	// Order update action to append to the `stagedActions` array.
	StagedAction StagedOrderUpdateAction `json:"stagedAction"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditAddStagedActionAction) UnmarshalJSON(data []byte) error {
	type Alias OrderEditAddStagedActionAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.StagedAction != nil {
		var err error
		obj.StagedAction, err = mapDiscriminatorStagedOrderUpdateAction(obj.StagedAction)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderEditAddStagedActionAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditAddStagedActionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addStagedAction", Alias: (*Alias)(&obj)})
}

type OrderEditSetCommentAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	Comment *string `json:"comment,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderEditSetCommentAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditSetCommentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setComment", Alias: (*Alias)(&obj)})
}

type OrderEditSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderEditSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type OrderEditSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the Order Edit with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Order Edit.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Order Edit.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderEditSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type OrderEditSetKeyAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderEditSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

/**
*	The `stagedActions` field cannot be updated if the Order Edit `result` is [OrderEdit Applied](/projects/order-edits#orderedit-applied).
*
 */
type OrderEditSetStagedActionsAction struct {
	// Value to replace the `stagedActions` of the Order Edit.
	StagedActions []StagedOrderUpdateAction `json:"stagedActions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditSetStagedActionsAction) UnmarshalJSON(data []byte) error {
	type Alias OrderEditSetStagedActionsAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.StagedActions {
		var err error
		obj.StagedActions[i], err = mapDiscriminatorStagedOrderUpdateAction(obj.StagedActions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderEditSetStagedActionsAction) MarshalJSON() ([]byte, error) {
	type Alias OrderEditSetStagedActionsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStagedActions", Alias: (*Alias)(&obj)})
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
type StagedOrderAddCustomLineItemAction struct {
	// Money value of the Custom Line Item. The value can be negative.
	Money Money `json:"money"`
	// Name of the Custom Line Item.
	Name LocalizedString `json:"name"`
	// User-defined unique identifier of the Custom Line Item.
	Key *string `json:"key,omitempty"`
	// Number of Custom Line Items to add to the Cart.
	Quantity *int `json:"quantity,omitempty"`
	// User-defined identifier used in a deep-link URL for the Custom Line Item. It must match the pattern `[a-zA-Z0-9_-]{2,256}`.
	Slug string `json:"slug"`
	// Used to select a Tax Rate when a Cart has the `Platform` [TaxMode](ctp:api:type:TaxMode).
	// If [TaxMode](ctp:api:type:TaxMode) is `Platform`, this field must not be empty.
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// An external Tax Rate can be set if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// Container for Custom Line Item-specific addresses.
	ShippingDetails *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
	// - If `Standard`, Cart Discounts with a matching [CartDiscountCustomLineItemsTarget](ctp:api:type:CartDiscountCustomLineItemsTarget), [MultiBuyCustomLineItemsTarget](ctp:api:type:MultiBuyCustomLineItemsTarget), or [CartDiscountPatternTarget](ctp:api:type:CartDiscountPatternTarget) are applied to the Custom Line Item.
	// - If `External`, Cart Discounts are not considered on the Custom Line Item.
	PriceMode *CustomLineItemPriceMode `json:"priceMode,omitempty"`
	// Custom Fields for the Custom Line Item.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Recurring Order and frequency data.
	RecurrenceInfo *CustomLineItemRecurrenceInfoDraft `json:"recurrenceInfo,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderAddCustomLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddCustomLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addCustomLineItem", Alias: (*Alias)(&obj)})
}

/**
*	A [Delivery](ctp:api:type:Delivery) can only be added to an [Order](ctp:api:type:Order) if
*	its `shippingInfo` (for `shippingMode` = `Single`), or its `shipping` (for `shippingMode` = `Multiple`) exists.
*
*	Produces the [Delivery Added](ctp:api:type:DeliveryAddedMessage) Message.
*
 */
type StagedOrderAddDeliveryAction struct {
	// `key` of an existing [Delivery](ctp:api:type:Delivery).
	DeliveryKey *string `json:"deliveryKey,omitempty"`
	// `key` of the [ShippingMethod](ctp:api:type:ShippingMethod), required for `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
	// Items to be included in the Delivery.
	Items []DeliveryItem `json:"items"`
	// Address the `parcels` should be delivered to.
	Address *BaseAddress `json:"address,omitempty"`
	// Parcels of the Delivery.
	//
	// If provided, this update action also produces the [Parcel Added To Delivery](ctp:api:type:ParcelAddedToDeliveryMessage) Message.
	Parcels []ParcelDraft `json:"parcels"`
	// Custom Fields for the Delivery.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderAddDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddDeliveryAction
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
*	Adds a [DiscountCode](ctp:api:type:DiscountCode) to the Cart to activate the related [Cart Discounts](/../api/projects/cartDiscounts).
*	Adding a Discount Code is only possible if no [DirectDiscount](ctp:api:type:DirectDiscount) has been applied to the Order.
*
*	The maximum number of Discount Codes in a Cart is restricted by a [limit](/../api/limits#carts).
*
*	Specific Error Code: [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError)
*
 */
type StagedOrderAddDiscountCodeAction struct {
	// `code` of a [DiscountCode](ctp:api:type:DiscountCode).
	Code string `json:"code"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderAddDiscountCodeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddDiscountCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addDiscountCode", Alias: (*Alias)(&obj)})
}

/**
*	Adds an address to an Order when shipping to multiple addresses is desired.
*
 */
type StagedOrderAddItemShippingAddressAction struct {
	// Address to append to `itemShippingAddresses`.
	// The new Address must have a `key` that is unique across this Order.
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderAddItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddItemShippingAddressAction
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
 */
type StagedOrderAddLineItemAction struct {
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
	// Sets the external Tax Rate for the Line Item, if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode). If the Cart has `Multiple` [ShippingMode](ctp:api:type:ShippingMode), the Tax Rate is accepted but ignored.
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
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
func (obj StagedOrderAddLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLineItem", Alias: (*Alias)(&obj)})
}

/**
*	To add a Parcel, at least one [Delivery](ctp:api:type:Delivery) must exist.
*
*	Produces the [Parcel Added To Delivery](ctp:api:type:ParcelAddedToDeliveryMessage) Message.
*
 */
type StagedOrderAddParcelToDeliveryAction struct {
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
func (obj StagedOrderAddParcelToDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddParcelToDeliveryAction
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

type StagedOrderAddPaymentAction struct {
	// Payment to add to the [PaymentInfo](ctp:api:type:PaymentInfo).
	// Must not be assigned to another Order or active Cart already.
	Payment PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderAddPaymentAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddPaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPayment", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [Return Info Added](ctp:api:type:ReturnInfoAddedMessage) Message.
*
 */
type StagedOrderAddReturnInfoAction struct {
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
func (obj StagedOrderAddReturnInfoAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddReturnInfoAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addReturnInfo", Alias: (*Alias)(&obj)})
}

/**
*	Adds all [LineItems](ctp:api:type:LineItem) of a [ShoppingList](ctp:api:type:ShoppingList) to the Cart.
*
 */
type StagedOrderAddShoppingListAction struct {
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
func (obj StagedOrderAddShoppingListAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderAddShoppingListAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShoppingList", Alias: (*Alias)(&obj)})
}

type StagedOrderChangeCustomLineItemMoneyAction struct {
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
	// Value to set.
	// Must not be empty.
	// Can be a negative amount.
	Money Money `json:"money"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderChangeCustomLineItemMoneyAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeCustomLineItemMoneyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCustomLineItemMoney", Alias: (*Alias)(&obj)})
}

/**
*	When multiple shipping addresses are set for a Custom Line Item, use the [Add CustomLineItem](ctp:api:type:StagedOrderAddCustomLineItemAction) update action to change the shipping details. Since it is not possible for the API to infer how the overall change in the Custom Line Item quantity should be distributed over the sub-quantities, the `shippingDetails` field is kept in its current state to avoid data loss.
*
*	To change the Custom Line Item quantity and shipping details together, use this update action in combination with the [Set CustomLineItem ShippingDetails](ctp:api:type:StagedOrderSetCustomLineItemShippingDetailsAction) update action in a single Order update command.
*
 */
type StagedOrderChangeCustomLineItemQuantityAction struct {
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
	// New value to set.
	// If `0`, the Custom Line Item is removed from the Order.
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderChangeCustomLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeCustomLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCustomLineItemQuantity", Alias: (*Alias)(&obj)})
}

/**
*	When multiple shipping addresses are set for a Line Item, use the [Remove LineItem](ctp:api:type:StagedOrderRemoveLineItemAction) and [Add LineItem](ctp:api:type:StagedOrderAddLineItemAction) update action to change the shipping details. Since it is not possible for the API to infer how the overall change in the Line Item quantity should be distributed over the sub-quantities, the `shippingDetails` field is kept in its current state to avoid data loss.
*
*	To change the Line Item quantity and shipping details together, use this update action in combination with the [Set LineItem ShippingDetails](ctp:api:type:StagedOrderSetLineItemShippingDetailsAction) update action in a single Order update command.
*
*	The [LineItem](ctp:api:type:LineItem) price is updated as described in [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection).
*
 */
type StagedOrderChangeLineItemQuantityAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// New value to set.
	// If `0`, the LineItem is removed from the Order.
	Quantity int `json:"quantity"`
	// Required when the Line Item uses `ExternalPrice` [LineItemPriceMode](ctp:api:type:LineItemPriceMode).
	// Sets the [LineItem](ctp:api:type:LineItem) `price` to the given value when changing the quantity of a Line Item.
	//
	// The [LineItem](ctp:api:type:LineItem) price is updated as described in [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection).
	ExternalPrice *Money `json:"externalPrice,omitempty"`
	// Sets the [LineItem](ctp:api:type:LineItem) `price` and `totalPrice` to the given value when changing the quantity of a Line Item with the `ExternalTotal` [LineItemPriceMode](ctp:api:type:LineItemPriceMode).
	// If `externalTotalPrice` is not given and the `priceMode` is `ExternalTotal`, the external price is unset and the `priceMode` is set to `Platform`.
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderChangeLineItemQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeLineItemQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLineItemQuantity", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [Order State Changed](ctp:api:type:OrderStateChangedMessage) Message.
*
 */
type StagedOrderChangeOrderStateAction struct {
	// New status of the Order.
	OrderState OrderState `json:"orderState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderChangeOrderStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeOrderStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeOrderState", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [Order Payment State Changed](ctp:api:type:OrderPaymentStateChangedMessage) Message.
*
 */
type StagedOrderChangePaymentStateAction struct {
	// New payment status of the Order.
	PaymentState PaymentState `json:"paymentState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderChangePaymentStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangePaymentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePaymentState", Alias: (*Alias)(&obj)})
}

/**
*	Changing the price rounding mode leads to [recalculation of taxes](/../api/carts-orders-overview#taxes).
*
 */
type StagedOrderChangePriceRoundingModeAction struct {
	// New value to set.
	PriceRoundingMode RoundingMode `json:"priceRoundingMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderChangePriceRoundingModeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangePriceRoundingModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePriceRoundingMode", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [Order Shipment State Changed](ctp:api:type:OrderShipmentStateChangedMessage) Message.
*
 */
type StagedOrderChangeShipmentStateAction struct {
	// New shipment status of the Order.
	ShipmentState ShipmentState `json:"shipmentState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderChangeShipmentStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeShipmentStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeShipmentState", Alias: (*Alias)(&obj)})
}

/**
*	Changing the tax calculation mode leads to [recalculation of taxes](/../api/carts-orders-overview#taxes).
*
 */
type StagedOrderChangeTaxCalculationModeAction struct {
	// New value to set.
	TaxCalculationMode TaxCalculationMode `json:"taxCalculationMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderChangeTaxCalculationModeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeTaxCalculationModeAction
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
type StagedOrderChangeTaxModeAction struct {
	// The new TaxMode.
	TaxMode TaxMode `json:"taxMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderChangeTaxModeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeTaxModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxMode", Alias: (*Alias)(&obj)})
}

/**
*	Changing the tax rounding mode leads to [recalculation of taxes](/../api/carts-orders-overview#taxes).
*
 */
type StagedOrderChangeTaxRoundingModeAction struct {
	// New value to set.
	TaxRoundingMode RoundingMode `json:"taxRoundingMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderChangeTaxRoundingModeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderChangeTaxRoundingModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxRoundingMode", Alias: (*Alias)(&obj)})
}

/**
*	The import of States does not follow any predefined rules and should be only used if no transitions are defined.
*	The `quantity` of the [ItemStates](ctp:api:type:ItemState) must match the sum of all Custom Line Item states' quantities.
*
 */
type StagedOrderImportCustomLineItemStateAction struct {
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
	// New status of the Custom Line Items.
	State []ItemState `json:"state"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderImportCustomLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderImportCustomLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "importCustomLineItemState", Alias: (*Alias)(&obj)})
}

/**
*	The import of States does not follow any predefined rules and should be only used if no transitions are defined.
*	The `quantity` in the [ItemStates](ctp:api:type:ItemState) must match the sum of all Line Item states' quantities.
*
 */
type StagedOrderImportLineItemStateAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// New status of the Line Items.
	State []ItemState `json:"state"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderImportLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderImportLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "importLineItemState", Alias: (*Alias)(&obj)})
}

/**
*	This update action does not support specifying a quantity, unlike the [Remove LineItem](ctp:api:type:StagedOrderRemoveLineItemAction) update action.
*
*	If `shippingDetails` must be partially removed, use the [Change CustomLineItem Quantity](ctp:api:type:StagedOrderChangeCustomLineItemQuantityAction) update action.
*
 */
type StagedOrderRemoveCustomLineItemAction struct {
	// `id` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemId *string `json:"customLineItemId,omitempty"`
	// `key` of the [CustomLineItem](ctp:api:type:CustomLineItem) to update. Either `customLineItemId` or `customLineItemKey` is required.
	CustomLineItemKey *string `json:"customLineItemKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderRemoveCustomLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveCustomLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeCustomLineItem", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [DeliveryRemoved](ctp:api:type:DeliveryRemovedMessage) Message.
*
 */
type StagedOrderRemoveDeliveryAction struct {
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
func (obj StagedOrderRemoveDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeDelivery", Alias: (*Alias)(&obj)})
}

type StagedOrderRemoveDiscountCodeAction struct {
	// Discount Code to remove from the Cart.
	DiscountCode DiscountCodeReference `json:"discountCode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderRemoveDiscountCodeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveDiscountCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeDiscountCode", Alias: (*Alias)(&obj)})
}

/**
*	An address can only be removed if it is not referenced in any [ItemShippingTarget](ctp:api:type:ItemShippingTarget) of the Cart.
*
 */
type StagedOrderRemoveItemShippingAddressAction struct {
	// `key` of the Address to remove from `itemShippingAddresses`.
	AddressKey string `json:"addressKey"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderRemoveItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeItemShippingAddress", Alias: (*Alias)(&obj)})
}

/**
*	The [LineItem](ctp:api:type:LineItem) price is updated as described in [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection).
*
 */
type StagedOrderRemoveLineItemAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// New value to set.
	// If absent or `0`, the Line Item is removed from the Cart.
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
func (obj StagedOrderRemoveLineItemAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveLineItemAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLineItem", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [ParcelRemovedFromDelivery](ctp:api:type:ParcelRemovedFromDeliveryMessage) Message.
*
 */
type StagedOrderRemoveParcelFromDeliveryAction struct {
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
func (obj StagedOrderRemoveParcelFromDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemoveParcelFromDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeParcelFromDelivery", Alias: (*Alias)(&obj)})
}

type StagedOrderRemovePaymentAction struct {
	// Payment to remove from the [PaymentInfo](ctp:api:type:PaymentInfo).
	Payment PaymentResourceIdentifier `json:"payment"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderRemovePaymentAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderRemovePaymentAction
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
type StagedOrderSetBillingAddressAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	Address *BaseAddress `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetBillingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetBillingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddress", Alias: (*Alias)(&obj)})
}

type StagedOrderSetBillingAddressCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetBillingAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetBillingAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddressCustomField", Alias: (*Alias)(&obj)})
}

type StagedOrderSetBillingAddressCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the `billingAddress` with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the `billingAddress`.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `billingAddress`.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetBillingAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetBillingAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBillingAddressCustomType", Alias: (*Alias)(&obj)})
}

/**
*	Updates the Business Unit on the Order. Setting the Order's `businessUnit` does not recalculate prices or discounts on the Order.
*
*	Produces the [OrderBusinessUnitSet](ctp:api:type:OrderBusinessUnitSetMessage) Message.
*
 */
type StagedOrderSetBusinessUnitAction struct {
	// New Business Unit to assign to the Order. If empty, any existing value is removed.
	//
	// If the referenced Business Unit does not exist, a [ReferencedResourceNotFound](ctp:api:type:ReferencedResourceNotFoundError) error is returned.
	BusinessUnit *BusinessUnitResourceIdentifier `json:"businessUnit,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetBusinessUnitAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetBusinessUnitAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setBusinessUnit", Alias: (*Alias)(&obj)})
}

/**
*	Setting the country can lead to changes in the [LineItem](ctp:api:type:LineItem) prices.
*
 */
type StagedOrderSetCountryAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	//
	// If the Cart is bound to a `store`, the provided value must be included in the [Store](ctp:api:type:Store)'s `countries`.
	// Otherwise a [CountryNotConfiguredInStore](ctp:api:type:CountryNotConfiguredInStoreError) error is returned.
	Country *string `json:"country,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetCountryAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCountryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCountry", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCustomLineItemCustomFieldAction struct {
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
func (obj StagedOrderSetCustomLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomField", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCustomLineItemCustomTypeAction struct {
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
func (obj StagedOrderSetCustomLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemCustomType", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCustomLineItemShippingDetailsAction struct {
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
func (obj StagedOrderSetCustomLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

/**
*	Can be used if the Cart has the `ExternalAmount` [TaxMode](ctp:api:type:TaxMode).
*
 */
type StagedOrderSetCustomLineItemTaxAmountAction struct {
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
func (obj StagedOrderSetCustomLineItemTaxAmountAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomLineItemTaxAmountAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemTaxAmount", Alias: (*Alias)(&obj)})
}

/**
*	Can be used if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
*
 */
type StagedOrderSetCustomLineItemTaxRateAction struct {
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
func (obj StagedOrderSetCustomLineItemTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomLineItemTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomLineItemTaxRate", Alias: (*Alias)(&obj)})
}

/**
*	To set the Cart's custom Shipping Method (independent of the [ShippingMethods](ctp:api:type:ShippingMethod) managed through the [Shipping Methods API](/../api/projects/shippingMethods)) the Cart must have the `Single` [ShippingMode](ctp:api:type:ShippingMode) and a `shippingAddress`.
*
*	To unset a custom Shipping Method on a Cart, use the [Set ShippingMethod](ctp:api:type:StagedOrderSetShippingMethodAction) update action without the `shippingMethod` field instead.
*
 */
type StagedOrderSetCustomShippingMethodAction struct {
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
func (obj StagedOrderSetCustomShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomShippingMethod", Alias: (*Alias)(&obj)})
}

type StagedOrderSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the Order Edit with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Order Edit.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Order Edit.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomTypeAction
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
type StagedOrderSetCustomerEmailAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	Email *string `json:"email,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetCustomerEmailAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomerEmailAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerEmail", Alias: (*Alias)(&obj)})
}

/**
*	This update action can only be used if a Customer is not assigned to a Cart.
*	If a Customer is already assigned, the Cart uses the Customer Group of the assigned Customer.
*
*	To reflect the new Customer Group, this update action can result in [updates to the Cart](/../api/carts-orders-overview#update-a-cart). When this occurs, the following errors can be returned: [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError) and [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError).
*
 */
type StagedOrderSetCustomerGroupAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetCustomerGroupAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomerGroupAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerGroup", Alias: (*Alias)(&obj)})
}

/**
*	Setting the Order's `customerId` does not recalculate prices or discounts on the Order.
*	If the Customer belongs to a Customer Group, `customerGroup` on the [Order](ctp:api:type:Order) is updated automatically.
*
*	Produces the [OrderCustomerSet](ctp:api:type:OrderCustomerSetMessage) Message.
*
 */
type StagedOrderSetCustomerIdAction struct {
	// `id` of an existing [Customer](ctp:api:type:Customer).
	// If empty, any existing value is removed.
	CustomerId *string `json:"customerId,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetCustomerIdAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetCustomerIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomerId", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [DeliveryAddressSet](ctp:api:type:DeliveryAddressSetMessage) Message.
*
 */
type StagedOrderSetDeliveryAddressAction struct {
	// `id` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryId *string `json:"deliveryId,omitempty"`
	// `key` of an existing [Delivery](ctp:api:type:Delivery).
	//
	// Either `deliveryId` or `deliveryKey` must be provided.
	DeliveryKey *string `json:"deliveryKey,omitempty"`
	// Value to set.
	// If empty, any existing value will be removed.
	Address *BaseAddress `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetDeliveryAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetDeliveryAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryAddress", Alias: (*Alias)(&obj)})
}

type StagedOrderSetDeliveryAddressCustomFieldAction struct {
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
func (obj StagedOrderSetDeliveryAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetDeliveryAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryAddressCustomField", Alias: (*Alias)(&obj)})
}

type StagedOrderSetDeliveryAddressCustomTypeAction struct {
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
func (obj StagedOrderSetDeliveryAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetDeliveryAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryAddressCustomType", Alias: (*Alias)(&obj)})
}

type StagedOrderSetDeliveryCustomFieldAction struct {
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
func (obj StagedOrderSetDeliveryCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetDeliveryCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryCustomField", Alias: (*Alias)(&obj)})
}

type StagedOrderSetDeliveryCustomTypeAction struct {
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
func (obj StagedOrderSetDeliveryCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetDeliveryCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryCustomType", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [Delivery Items Updated](ctp:api:type:DeliveryItemsUpdatedMessage) Message.
*
 */
type StagedOrderSetDeliveryItemsAction struct {
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
func (obj StagedOrderSetDeliveryItemsAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetDeliveryItemsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDeliveryItems", Alias: (*Alias)(&obj)})
}

/**
*	Adds a [DirectDiscount](ctp:api:type:DirectDiscount), but only if no [DiscountCode](ctp:api:type:DiscountCode) has been added to the Order.
*	Either a Discount Code or a Direct Discount can exist on a Order at the same time.
*
 */
type StagedOrderSetDirectDiscountsAction struct {
	// - If set, all existing Direct Discounts are replaced.
	//   The discounts apply in the order they are added to the list.
	// - If empty, all existing Direct Discounts are removed and all affected prices on the Order are recalculated.
	Discounts []DirectDiscountDraft `json:"discounts"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetDirectDiscountsAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetDirectDiscountsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDirectDiscounts", Alias: (*Alias)(&obj)})
}

type StagedOrderSetItemShippingAddressCustomFieldAction struct {
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
func (obj StagedOrderSetItemShippingAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetItemShippingAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setItemShippingAddressCustomField", Alias: (*Alias)(&obj)})
}

type StagedOrderSetItemShippingAddressCustomTypeAction struct {
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
func (obj StagedOrderSetItemShippingAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetItemShippingAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setItemShippingAddressCustomType", Alias: (*Alias)(&obj)})
}

type StagedOrderSetLineItemCustomFieldAction struct {
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
func (obj StagedOrderSetLineItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomField", Alias: (*Alias)(&obj)})
}

type StagedOrderSetLineItemCustomTypeAction struct {
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
func (obj StagedOrderSetLineItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemCustomType", Alias: (*Alias)(&obj)})
}

/**
*	Setting a distribution channel for a [LineItem](ctp:api:type:LineItem) can lead to an updated `price` as described in [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection).
*
*	Produces the [OrderLineItemDistributionChannelSet](ctp:api:type:OrderLineItemDistributionChannelSetMessage) Message.
*
 */
type StagedOrderSetLineItemDistributionChannelAction struct {
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
func (obj StagedOrderSetLineItemDistributionChannelAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemDistributionChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemDistributionChannel", Alias: (*Alias)(&obj)})
}

/**
*	Sets the [LineItem](ctp:api:type:LineItem) `price` and changes the `priceMode` to `ExternalPrice` [LineItemPriceMode](ctp:api:type:LineItemPriceMode).
*
 */
type StagedOrderSetLineItemPriceAction struct {
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
func (obj StagedOrderSetLineItemPriceAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemPrice", Alias: (*Alias)(&obj)})
}

type StagedOrderSetLineItemShippingDetailsAction struct {
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
func (obj StagedOrderSetLineItemShippingDetailsAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemShippingDetailsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemShippingDetails", Alias: (*Alias)(&obj)})
}

/**
*	Can be used if the Cart has the `ExternalAmount` [TaxMode](ctp:api:type:TaxMode). This update action sets the `taxedPrice` and `taxRate` on a Line Item and must be used after any price-affecting change occurs.
*
 */
type StagedOrderSetLineItemTaxAmountAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Value to set.
	// If empty, any existing value will be removed.
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
	// `key` of the [ShippingMethod](ctp:api:type:ShippingMethod) used for this Line Item.
	// This is required for Carts with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetLineItemTaxAmountAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemTaxAmountAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemTaxAmount", Alias: (*Alias)(&obj)})
}

/**
*	Can be used if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
*
 */
type StagedOrderSetLineItemTaxRateAction struct {
	// `id` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemId *string `json:"lineItemId,omitempty"`
	// `key` of the [LineItem](ctp:api:type:LineItem) to update. Either `lineItemId` or `lineItemKey` is required.
	LineItemKey *string `json:"lineItemKey,omitempty"`
	// Value to set.
	// If empty, any existing value will be removed.
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// `key` of the [ShippingMethod](ctp:api:type:ShippingMethod) used for this Line Item.
	// This is required for Carts with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetLineItemTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemTaxRate", Alias: (*Alias)(&obj)})
}

/**
*	Sets the [LineItem](ctp:api:type:LineItem) `totalPrice` and `price`, and changes the `priceMode` to `ExternalTotal` [LineItemPriceMode](ctp:api:type:LineItemPriceMode).
*
 */
type StagedOrderSetLineItemTotalPriceAction struct {
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
func (obj StagedOrderSetLineItemTotalPriceAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLineItemTotalPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLineItemTotalPrice", Alias: (*Alias)(&obj)})
}

type StagedOrderSetLocaleAction struct {
	// Value to set.
	// Must be one of the [Project](ctp:api:type:Project)'s languages.
	// If empty, any existing value is removed.
	Locale *string `json:"locale,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

type StagedOrderSetOrderNumberAction struct {
	// Value to set. Must be unique across a Project.
	// Once set, the value cannot be changed.
	OrderNumber *string `json:"orderNumber,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetOrderNumberAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetOrderNumberAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setOrderNumber", Alias: (*Alias)(&obj)})
}

/**
*	Updates the total tax amount of the Order if it has the `ExternalAmount` [TaxMode](ctp:api:type:TaxMode).
*
 */
type StagedOrderSetOrderTotalTaxAction struct {
	// Total gross amount of the Order (totalNet + taxes).
	ExternalTotalGross Money `json:"externalTotalGross"`
	// Value to set.
	ExternalTaxPortions []TaxPortionDraft `json:"externalTaxPortions"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetOrderTotalTaxAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetOrderTotalTaxAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setOrderTotalTax", Alias: (*Alias)(&obj)})
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

type StagedOrderSetParcelCustomFieldAction struct {
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
func (obj StagedOrderSetParcelCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetParcelCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelCustomField", Alias: (*Alias)(&obj)})
}

type StagedOrderSetParcelCustomTypeAction struct {
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
func (obj StagedOrderSetParcelCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetParcelCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelCustomType", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [ParcelItemsUpdated](ctp:api:type:ParcelItemsUpdatedMessage) Message.
*
 */
type StagedOrderSetParcelItemsAction struct {
	// `id` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelId *string `json:"parcelId,omitempty"`
	// `key` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelKey *string `json:"parcelKey,omitempty"`
	// Value to set.
	// If empty, any existing value will be removed.
	Items []DeliveryItem `json:"items"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetParcelItemsAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetParcelItemsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelItems", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [ParcelMeasurementsUpdated](ctp:api:type:ParcelMeasurementsUpdatedMessage) Message.
*
 */
type StagedOrderSetParcelMeasurementsAction struct {
	// `id` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelId *string `json:"parcelId,omitempty"`
	// `key` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelKey *string `json:"parcelKey,omitempty"`
	// Value to set.
	// If empty, any existing value will be removed.
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetParcelMeasurementsAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetParcelMeasurementsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelMeasurements", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [ParcelTrackingDataUpdated](ctp:api:type:ParcelTrackingDataUpdatedMessage) Message.
*
 */
type StagedOrderSetParcelTrackingDataAction struct {
	// `id` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelId *string `json:"parcelId,omitempty"`
	// `key` of an existing [Parcel](ctp:api:type:Parcel).
	//
	// Either `parcelId` or `parcelKey` must be provided.
	ParcelKey *string `json:"parcelKey,omitempty"`
	// Value to set.
	// If empty, any existing value will be removed.
	TrackingData *TrackingData `json:"trackingData,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetParcelTrackingDataAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetParcelTrackingDataAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setParcelTrackingData", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [PurchaseOrderNumberSet](ctp:api:type:OrderPurchaseOrderNumberSetMessage) Message.
*
 */
type StagedOrderSetPurchaseOrderNumberAction struct {
	// Value to set.
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetPurchaseOrderNumberAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetPurchaseOrderNumberAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPurchaseOrderNumber", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [Return Info Set](ctp:api:type:ReturnInfoSetMessage) Message.
*
 */
type StagedOrderSetReturnInfoAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	Items []ReturnInfoDraft `json:"items"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetReturnInfoAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetReturnInfoAction
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

type StagedOrderSetReturnItemCustomFieldAction struct {
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
func (obj StagedOrderSetReturnItemCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetReturnItemCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setReturnItemCustomField", Alias: (*Alias)(&obj)})
}

type StagedOrderSetReturnItemCustomTypeAction struct {
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
func (obj StagedOrderSetReturnItemCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetReturnItemCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setReturnItemCustomType", Alias: (*Alias)(&obj)})
}

/**
*	To set a [ReturnPaymentState](ctp:api:type:ReturnPaymentState), the [Order](ctp:api:type:Order) `returnInfo` must have at least one [ReturnItem](ctp:api:type:ReturnItem).
*
 */
type StagedOrderSetReturnPaymentStateAction struct {
	// `id` of the [ReturnItem](ctp:api:type:ReturnItem) to update. Either `returnItemId` or `returnItemKey` is required.
	ReturnItemId *string `json:"returnItemId,omitempty"`
	// `key` of the [ReturnItem](ctp:api:type:ReturnItem) to update. Either `returnItemId` or `returnItemKey` is required.
	ReturnItemKey *string `json:"returnItemKey,omitempty"`
	// New Payment status of the [ReturnItem](ctp:api:type:ReturnItem).
	PaymentState ReturnPaymentState `json:"paymentState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetReturnPaymentStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetReturnPaymentStateAction
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
type StagedOrderSetReturnShipmentStateAction struct {
	// `id` of the [ReturnItem](ctp:api:type:ReturnItem) to update. Either `returnItemId` or `returnItemKey` is required.
	ReturnItemId *string `json:"returnItemId,omitempty"`
	// `key` of the [ReturnItem](ctp:api:type:ReturnItem) to update. Either `returnItemId` or `returnItemKey` is required.
	ReturnItemKey *string `json:"returnItemKey,omitempty"`
	// New shipment state of the [ReturnItem](ctp:api:type:ReturnItem).
	ShipmentState ReturnShipmentState `json:"shipmentState"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetReturnShipmentStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetReturnShipmentStateAction
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
type StagedOrderSetShippingAddressAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	Address *BaseAddress `json:"address,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddress", Alias: (*Alias)(&obj)})
}

/**
*	Sets the shipping address and a custom Shipping Method together to prevent an inconsistent state.
*
 */
type StagedOrderSetShippingAddressAndCustomShippingMethodAction struct {
	// Value to set for `shippingAddress`.
	Address BaseAddress `json:"address"`
	// Value to set.
	ShippingMethodName string `json:"shippingMethodName"`
	// Value to set.
	ShippingRate ShippingRateDraft `json:"shippingRate"`
	// Used to select a Tax Rate when the Order has the `Platform` [TaxMode](ctp:api:type:TaxMode).
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// An external Tax Rate can be set if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
	// Custom Fields for the custom Shipping Method.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetShippingAddressAndCustomShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingAddressAndCustomShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddressAndCustomShippingMethod", Alias: (*Alias)(&obj)})
}

/**
*	Sets the shipping address and Shipping Method together to prevent an inconsistent state.
*
 */
type StagedOrderSetShippingAddressAndShippingMethodAction struct {
	// Value to set for `shippingAddress`.
	Address BaseAddress `json:"address"`
	// Value to set.
	ShippingMethod *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	// An external Tax Rate can be set if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetShippingAddressAndShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingAddressAndShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddressAndShippingMethod", Alias: (*Alias)(&obj)})
}

type StagedOrderSetShippingAddressCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetShippingAddressCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingAddressCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddressCustomField", Alias: (*Alias)(&obj)})
}

type StagedOrderSetShippingAddressCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the `shippingAddress` with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the `shippingAddress`.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `shippingAddress`.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetShippingAddressCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingAddressCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingAddressCustomType", Alias: (*Alias)(&obj)})
}

type StagedOrderSetShippingCustomFieldAction struct {
	// The `shippingKey` of the [Shipping](ctp:api:type:Shipping) to customize. Used to specify which Shipping Method to customize
	// on a Order with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	// Leave this empty to customize the one and only ShippingMethod on a `Single` ShippingMode Order.
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
func (obj StagedOrderSetShippingCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingCustomField", Alias: (*Alias)(&obj)})
}

/**
*	This action sets, overwrites, or removes any existing Custom Type and Custom Fields for the Order's `shippingMethod` or `shipping`.
*
 */
type StagedOrderSetShippingCustomTypeAction struct {
	// The `shippingKey` of the [Shipping](ctp:api:type:Shipping) to customize. Used to specify which Shipping Method to customize
	// on a Order with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	// Leave this empty to customize the one and only ShippingMethod on a `Single` ShippingMode Order.
	ShippingKey *string `json:"shippingKey,omitempty"`
	// Defines the [Type](ctp:api:type:Type) that extends the specified ShippingMethod with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the ShippingMethod.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the `shippingMethod`.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetShippingCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingCustomType", Alias: (*Alias)(&obj)})
}

/**
*	To set the Cart's Shipping Method, the Cart must have the `Single` [ShippingMode](ctp:api:type:ShippingMode) and a `shippingAddress`.
*
 */
type StagedOrderSetShippingMethodAction struct {
	// Value to set. If empty, any existing value will be removed.
	// If the referenced Shipping Method has a predicate that does not match the Cart, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
	ShippingMethod *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	// An external Tax Rate can be set if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetShippingMethodAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingMethod", Alias: (*Alias)(&obj)})
}

/**
*	A Shipping Method tax amount can be set if the Cart has the `ExternalAmount` [TaxMode](ctp:api:type:TaxMode).
*
 */
type StagedOrderSetShippingMethodTaxAmountAction struct {
	// `key` of the [ShippingMethod](ctp:api:type:ShippingMethod) to update. This is required for Orders with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
	// Value to set. If empty, any existing value will be removed.
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetShippingMethodTaxAmountAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingMethodTaxAmountAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingMethodTaxAmount", Alias: (*Alias)(&obj)})
}

/**
*	A Shipping Method Tax Rate can be set if the Cart has the `External` [TaxMode](ctp:api:type:TaxMode).
*
 */
type StagedOrderSetShippingMethodTaxRateAction struct {
	// `key` of the [ShippingMethod](ctp:api:type:ShippingMethod) to update. This is required for Orders with `Multiple` [ShippingMode](ctp:api:type:ShippingMode).
	ShippingKey *string `json:"shippingKey,omitempty"`
	// Value to set.
	// If empty, any existing value is removed.
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetShippingMethodTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingMethodTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingMethodTaxRate", Alias: (*Alias)(&obj)})
}

/**
*	Input used to select a [ShippingRatePriceTier](ctp:api:type:ShippingRatePriceTier). If no matching tier can be found, or the input is not set, the default price for the shipping rate is used.
*
 */
type StagedOrderSetShippingRateInputAction struct {
	// The data type of this field depends on the `shippingRateInputType.type` configured in the [Project](ctp:api:type:Project):
	//
	// - If `CartClassification`, it must be [ClassificationShippingRateInputDraft](ctp:api:type:ClassificationShippingRateInputDraft).
	// - If `CartScore`, it must be [ScoreShippingRateInputDraft](ctp:api:type:ScoreShippingRateInputDraft).
	// - If `CartValue`, it cannot be set.
	ShippingRateInput ShippingRateInputDraft `json:"shippingRateInput,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StagedOrderSetShippingRateInputAction) UnmarshalJSON(data []byte) error {
	type Alias StagedOrderSetShippingRateInputAction
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
func (obj StagedOrderSetShippingRateInputAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetShippingRateInputAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingRateInput", Alias: (*Alias)(&obj)})
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
type StagedOrderSetStoreAction struct {
	// Value to set. If empty, any existing value will be removed.
	//
	// If `store` references the same Store the Order is currently assigned to or if you try to remove the value when no Store is currently assigned, a `400` error is returned.
	Store *StoreResourceIdentifier `json:"store,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderSetStoreAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderSetStoreAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStore", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [Custom Line Item State Transition](ctp:api:type:CustomLineItemStateTransitionMessage) Message.
*
 */
type StagedOrderTransitionCustomLineItemStateAction struct {
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
func (obj StagedOrderTransitionCustomLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderTransitionCustomLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionCustomLineItemState", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [Line Item State Transition](ctp:api:type:LineItemStateTransitionMessage) Message.
*
 */
type StagedOrderTransitionLineItemStateAction struct {
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
func (obj StagedOrderTransitionLineItemStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderTransitionLineItemStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionLineItemState", Alias: (*Alias)(&obj)})
}

/**
*	If the existing [State](ctp:api:type:State) has set `transitions`, there must be a direct transition to the new State. If `transitions` is not set, no validation is performed.
*
*	This update action produces the [Order State Transition](ctp:api:type:OrderStateTransitionMessage) Message.
*
 */
type StagedOrderTransitionStateAction struct {
	// Value to set.
	// If there is no State yet, the new State must be an initial State.
	State StateResourceIdentifier `json:"state"`
	// Set to `true` to turn off validation.
	Force *bool `json:"force,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}

/**
*	Updates an address in `itemShippingAddresses` by keeping the Address `key`.
*
 */
type StagedOrderUpdateItemShippingAddressAction struct {
	// The new Address with the same `key` as the Address it will replace.
	Address BaseAddress `json:"address"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StagedOrderUpdateItemShippingAddressAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderUpdateItemShippingAddressAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateItemShippingAddress", Alias: (*Alias)(&obj)})
}

type StagedOrderUpdateSyncInfoAction struct {
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
func (obj StagedOrderUpdateSyncInfoAction) MarshalJSON() ([]byte, error) {
	type Alias StagedOrderUpdateSyncInfoAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "updateSyncInfo", Alias: (*Alias)(&obj)})
}
