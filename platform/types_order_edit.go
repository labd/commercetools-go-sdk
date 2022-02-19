// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type OrderEdit struct {
	// The unique ID of the OrderEdit.
	ID string `json:"id"`
	// The current version of the OrderEdit.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Unique identifier for this edit.
	Key *string `json:"key,omitempty"`
	// The order to be updated with this edit.
	Resource OrderReference `json:"resource"`
	// The actions to apply to the Order.
	// Cannot be updated after the edit has been applied.
	StagedActions []StagedOrderUpdateAction `json:"stagedActions"`
	Custom        *CustomFields             `json:"custom,omitempty"`
	// Contains a preview of the changes in case of unapplied edit.
	// For applied edits, it contains the summary of the changes.
	Result OrderEditResult `json:"result"`
	// This field can be used to add textual information regarding the edit.
	Comment *string `json:"comment,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEdit) UnmarshalJSON(data []byte) error {
	type Alias OrderEdit
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
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

type OrderEditApply struct {
	EditVersion     int `json:"editVersion"`
	ResourceVersion int `json:"resourceVersion"`
}

type OrderEditDraft struct {
	// Unique identifier for this edit.
	Key *string `json:"key,omitempty"`
	// The order to be updated with this edit.
	Resource OrderReference `json:"resource"`
	// The actions to apply to `resource`.
	StagedActions []StagedOrderUpdateAction `json:"stagedActions"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// This field can be used to add additional textual information regarding the edit.
	Comment *string `json:"comment,omitempty"`
	// When set to `true` the edit is applied on the Order without persisting it.
	DryRun *bool `json:"dryRun,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditDraft) UnmarshalJSON(data []byte) error {
	type Alias OrderEditDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["stagedActions"] == nil {
		delete(target, "stagedActions")
	}

	return json.Marshal(target)
}

type OrderEditPagedQueryResponse struct {
	Limit   int         `json:"limit"`
	Count   int         `json:"count"`
	Total   *int        `json:"total,omitempty"`
	Offset  int         `json:"offset"`
	Results []OrderEdit `json:"results"`
}

type OrderEditReference struct {
	// Unique ID of the referenced resource.
	ID  string     `json:"id"`
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

type OrderEditResourceIdentifier struct {
	// Unique ID of the referenced resource. Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced resource. Either `id` or `key` is required.
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
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
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
		return obj, nil
	case "PreviewSuccess":
		obj := OrderEditPreviewSuccess{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type OrderEditApplied struct {
	AppliedAt         time.Time    `json:"appliedAt"`
	ExcerptBeforeEdit OrderExcerpt `json:"excerptBeforeEdit"`
	ExcerptAfterEdit  OrderExcerpt `json:"excerptAfterEdit"`
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

type OrderEditPreviewFailure struct {
	Errors []ErrorObject `json:"errors"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditPreviewFailure) UnmarshalJSON(data []byte) error {
	type Alias OrderEditPreviewFailure
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
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

type OrderEditPreviewSuccess struct {
	Preview         StagedOrder      `json:"preview"`
	MessagePayloads []MessagePayload `json:"messagePayloads"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditPreviewSuccess) UnmarshalJSON(data []byte) error {
	type Alias OrderEditPreviewSuccess
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
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
	Version int                     `json:"version"`
	Actions []OrderEditUpdateAction `json:"actions"`
	DryRun  *bool                   `json:"dryRun,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditUpdate) UnmarshalJSON(data []byte) error {
	type Alias OrderEditUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type OrderEditUpdateAction interface{}

func mapDiscriminatorOrderEditUpdateAction(input interface{}) (OrderEditUpdateAction, error) {

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
		return obj, nil
	}
	return nil, nil
}

type OrderExcerpt struct {
	TotalPrice TypedMoney  `json:"totalPrice"`
	TaxedPrice *TaxedPrice `json:"taxedPrice,omitempty"`
	Version    int         `json:"version"`
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
	// The unique ID of the order.
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
	// The sequence number of the last order message produced by changes to this order.
	// `0` means, that no messages were created yet.
	LastMessageSequenceNumber int `json:"lastMessageSequenceNumber"`
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
func (obj *StagedOrder) UnmarshalJSON(data []byte) error {
	type Alias StagedOrder
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
func (obj StagedOrder) MarshalJSON() ([]byte, error) {
	type Alias StagedOrder
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

type OrderEditAddStagedActionAction struct {
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
	Name string `json:"name"`
	// The value of a Custom Field.
	// The data type of the value depends on the specific [FieldType](/projects/types#fieldtype) given in the `type` field of the [FieldDefinition](/ctp:api:type:FieldDefinition) for a Custom Field.
	// It can be any of the following:
	//
	//  Field type                                                 | Data type                                                                                                                                                                 |
	//  ---------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
	//  [CustomFieldBooleanType](ctp:api:type:CustomFieldBooleanType)                 | Boolean (`true` or `false`)                                                                                                                                                     |
	//  [CustomFieldStringType](ctp:api:type:CustomFieldStringType)                   | String                                                                                                                                                                |
	//  [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType) | [LocalizedString](ctp:api:type:LocalizedString)                                                                                                                             |
	//  [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType)                       | String. Must be a `key` of one of the [EnumValues](ctp:api:type:CustomFieldEnumValue) defined in the [EnumType](ctp:api:type:CustomFiedEnumType)                                     |
	//  [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType)     | String. Must be a `key` of one of the [LocalizedEnumValues](ctp:api:type:CustomFieldLocalizedEnumValue) defined in the [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) |
	//  [CustomFieldNumberType](ctp:api:type:CustomFieldNumberType)                   | Number                                                                                                                                                                |
	//  [CustomFieldMoneyType](ctp:api:type:CustomFieldMoneyType)                     | [CentPrecisionMoney](/../api/types#centprecisionmoney)                                                                                                                                         |
	//  [CustomFieldDateType](ctp:api:type:CustomFieldDateType)                       | [Date](ctp:api:type:Date)                                                                                                                                                   |
	//  [CustomFieldTimeType](ctp:api:type:CustomFieldTimeType)                       | [Time](ctp:api:type:Time)                                                                                                                                                   |
	//  [CustomFieldDateTimeType](ctp:api:type:CustomFieldDateTimeType)               | [DateTime](ctp:api:type:DateTime)                                                                                                                                           |
	//  [CustomFieldReferenceType](ctp:api:type:CustomFieldReferenceType)             | [Reference](/../api/types#reference)                                                                                                                                         |
	//  [CustomFieldSetType](ctp:api:type:CustomFieldSetType)                         | JSON array without duplicates consisting of [CustomFieldValues](ctp:api:type:CustomFieldValue) of a single [FieldType](ctp:api:type:FieldType). For example, a Custom Field of SetType of DateType takes a JSON array of mutually different Dates for its value. The order of items in the array is not fixed. For more examples, see the [example FieldContainer](ctp:api:type:FieldContainer).|
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
	// If set, the custom type is set to this new value.
	// If absent, the custom type and any existing custom fields are removed.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// If set, the custom fields are set to this new value.
	Fields *interface{} `json:"fields,omitempty"`
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
	// If `key` is absent or `null`, this field will be removed if it exists.
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

type OrderEditSetStagedActionsAction struct {
	// The actions to edit the `resource`.
	StagedActions []StagedOrderUpdateAction `json:"stagedActions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *OrderEditSetStagedActionsAction) UnmarshalJSON(data []byte) error {
	type Alias OrderEditSetStagedActionsAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
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

type StagedOrderAddCustomLineItemAction struct {
	// Draft type that stores amounts in cent precision for the specified currency.
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	Money    Money           `json:"money"`
	Name     LocalizedString `json:"name"`
	Quantity *float64        `json:"quantity,omitempty"`
	Slug     string          `json:"slug"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [TaxCategory](ctp:api:type:TaxCategory).
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// The representation used when creating or updating a [customizable data type](/../api/projects/types#list-of-customizable-data-types) with Custom Fields.
	Custom          *CustomFieldsDraft    `json:"custom,omitempty"`
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
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

type StagedOrderAddDeliveryAction struct {
	Items   []DeliveryItem `json:"items"`
	Address *BaseAddress   `json:"address,omitempty"`
	Parcels []ParcelDraft  `json:"parcels"`
	// Custom Fields for the Transaction.
	Custom *CustomFields `json:"custom,omitempty"`
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

type StagedOrderAddDiscountCodeAction struct {
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

type StagedOrderAddItemShippingAddressAction struct {
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

type StagedOrderAddLineItemAction struct {
	// The representation used when creating or updating a [customizable data type](/../api/projects/types#list-of-customizable-data-types) with Custom Fields.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [Channel](ctp:api:type:Channel).
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
	ExternalTaxRate     *ExternalTaxRateDraft      `json:"externalTaxRate,omitempty"`
	ProductId           *string                    `json:"productId,omitempty"`
	VariantId           *int                       `json:"variantId,omitempty"`
	Sku                 *string                    `json:"sku,omitempty"`
	Quantity            *float64                   `json:"quantity,omitempty"`
	AddedAt             *time.Time                 `json:"addedAt,omitempty"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [Channel](ctp:api:type:Channel).
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	// Draft type that stores amounts in cent precision for the specified currency.
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	ExternalPrice      *Money                      `json:"externalPrice,omitempty"`
	ExternalTotalPrice *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ShippingDetails    *ItemShippingDetailsDraft   `json:"shippingDetails,omitempty"`
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

type StagedOrderAddParcelToDeliveryAction struct {
	DeliveryId   string              `json:"deliveryId"`
	Measurements *ParcelMeasurements `json:"measurements,omitempty"`
	TrackingData *TrackingData       `json:"trackingData,omitempty"`
	Items        []DeliveryItem      `json:"items"`
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["items"] == nil {
		delete(target, "items")
	}

	return json.Marshal(target)
}

type StagedOrderAddPaymentAction struct {
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

type StagedOrderAddReturnInfoAction struct {
	ReturnTrackingId *string           `json:"returnTrackingId,omitempty"`
	Items            []ReturnItemDraft `json:"items"`
	ReturnDate       *time.Time        `json:"returnDate,omitempty"`
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

type StagedOrderAddShoppingListAction struct {
	ShoppingList ShoppingListResourceIdentifier `json:"shoppingList"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [Channel](ctp:api:type:Channel).
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [Channel](ctp:api:type:Channel).
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
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
	CustomLineItemId string `json:"customLineItemId"`
	// Draft type that stores amounts in cent precision for the specified currency.
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
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

type StagedOrderChangeCustomLineItemQuantityAction struct {
	CustomLineItemId string  `json:"customLineItemId"`
	Quantity         float64 `json:"quantity"`
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

type StagedOrderChangeLineItemQuantityAction struct {
	LineItemId string  `json:"lineItemId"`
	Quantity   float64 `json:"quantity"`
	// Draft type that stores amounts in cent precision for the specified currency.
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	ExternalPrice      *Money                      `json:"externalPrice,omitempty"`
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

type StagedOrderChangeOrderStateAction struct {
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

type StagedOrderChangePaymentStateAction struct {
	PaymentState *PaymentState `json:"paymentState,omitempty"`
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

type StagedOrderChangeShipmentStateAction struct {
	ShipmentState *ShipmentState `json:"shipmentState,omitempty"`
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

type StagedOrderChangeTaxCalculationModeAction struct {
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

type StagedOrderChangeTaxModeAction struct {
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

type StagedOrderChangeTaxRoundingModeAction struct {
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

type StagedOrderImportCustomLineItemStateAction struct {
	CustomLineItemId string      `json:"customLineItemId"`
	State            []ItemState `json:"state"`
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

type StagedOrderImportLineItemStateAction struct {
	LineItemId string      `json:"lineItemId"`
	State      []ItemState `json:"state"`
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

type StagedOrderRemoveCustomLineItemAction struct {
	CustomLineItemId string `json:"customLineItemId"`
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

type StagedOrderRemoveDeliveryAction struct {
	DeliveryId string `json:"deliveryId"`
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

type StagedOrderRemoveItemShippingAddressAction struct {
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

type StagedOrderRemoveLineItemAction struct {
	LineItemId string   `json:"lineItemId"`
	Quantity   *float64 `json:"quantity,omitempty"`
	// Draft type that stores amounts in cent precision for the specified currency.
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	ExternalPrice           *Money                      `json:"externalPrice,omitempty"`
	ExternalTotalPrice      *ExternalLineItemTotalPrice `json:"externalTotalPrice,omitempty"`
	ShippingDetailsToRemove *ItemShippingDetailsDraft   `json:"shippingDetailsToRemove,omitempty"`
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

type StagedOrderRemoveParcelFromDeliveryAction struct {
	ParcelId string `json:"parcelId"`
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

type StagedOrderSetBillingAddressAction struct {
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
	Name string `json:"name"`
	// The value of a Custom Field.
	// The data type of the value depends on the specific [FieldType](/projects/types#fieldtype) given in the `type` field of the [FieldDefinition](/ctp:api:type:FieldDefinition) for a Custom Field.
	// It can be any of the following:
	//
	//  Field type                                                 | Data type                                                                                                                                                                 |
	//  ---------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
	//  [CustomFieldBooleanType](ctp:api:type:CustomFieldBooleanType)                 | Boolean (`true` or `false`)                                                                                                                                                     |
	//  [CustomFieldStringType](ctp:api:type:CustomFieldStringType)                   | String                                                                                                                                                                |
	//  [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType) | [LocalizedString](ctp:api:type:LocalizedString)                                                                                                                             |
	//  [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType)                       | String. Must be a `key` of one of the [EnumValues](ctp:api:type:CustomFieldEnumValue) defined in the [EnumType](ctp:api:type:CustomFiedEnumType)                                     |
	//  [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType)     | String. Must be a `key` of one of the [LocalizedEnumValues](ctp:api:type:CustomFieldLocalizedEnumValue) defined in the [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) |
	//  [CustomFieldNumberType](ctp:api:type:CustomFieldNumberType)                   | Number                                                                                                                                                                |
	//  [CustomFieldMoneyType](ctp:api:type:CustomFieldMoneyType)                     | [CentPrecisionMoney](/../api/types#centprecisionmoney)                                                                                                                                         |
	//  [CustomFieldDateType](ctp:api:type:CustomFieldDateType)                       | [Date](ctp:api:type:Date)                                                                                                                                                   |
	//  [CustomFieldTimeType](ctp:api:type:CustomFieldTimeType)                       | [Time](ctp:api:type:Time)                                                                                                                                                   |
	//  [CustomFieldDateTimeType](ctp:api:type:CustomFieldDateTimeType)               | [DateTime](ctp:api:type:DateTime)                                                                                                                                           |
	//  [CustomFieldReferenceType](ctp:api:type:CustomFieldReferenceType)             | [Reference](/../api/types#reference)                                                                                                                                         |
	//  [CustomFieldSetType](ctp:api:type:CustomFieldSetType)                         | JSON array without duplicates consisting of [CustomFieldValues](ctp:api:type:CustomFieldValue) of a single [FieldType](ctp:api:type:FieldType). For example, a Custom Field of SetType of DateType takes a JSON array of mutually different Dates for its value. The order of items in the array is not fixed. For more examples, see the [example FieldContainer](ctp:api:type:FieldContainer).|
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
	// [ResourceIdentifier](/../api/types#resourceidentifier) of a [Type](ctp:api:type:Type).
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
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

type StagedOrderSetCountryAction struct {
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
	Name string `json:"name"`
	// The value of a Custom Field.
	// The data type of the value depends on the specific [FieldType](/projects/types#fieldtype) given in the `type` field of the [FieldDefinition](/ctp:api:type:FieldDefinition) for a Custom Field.
	// It can be any of the following:
	//
	//  Field type                                                 | Data type                                                                                                                                                                 |
	//  ---------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
	//  [CustomFieldBooleanType](ctp:api:type:CustomFieldBooleanType)                 | Boolean (`true` or `false`)                                                                                                                                                     |
	//  [CustomFieldStringType](ctp:api:type:CustomFieldStringType)                   | String                                                                                                                                                                |
	//  [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType) | [LocalizedString](ctp:api:type:LocalizedString)                                                                                                                             |
	//  [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType)                       | String. Must be a `key` of one of the [EnumValues](ctp:api:type:CustomFieldEnumValue) defined in the [EnumType](ctp:api:type:CustomFiedEnumType)                                     |
	//  [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType)     | String. Must be a `key` of one of the [LocalizedEnumValues](ctp:api:type:CustomFieldLocalizedEnumValue) defined in the [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) |
	//  [CustomFieldNumberType](ctp:api:type:CustomFieldNumberType)                   | Number                                                                                                                                                                |
	//  [CustomFieldMoneyType](ctp:api:type:CustomFieldMoneyType)                     | [CentPrecisionMoney](/../api/types#centprecisionmoney)                                                                                                                                         |
	//  [CustomFieldDateType](ctp:api:type:CustomFieldDateType)                       | [Date](ctp:api:type:Date)                                                                                                                                                   |
	//  [CustomFieldTimeType](ctp:api:type:CustomFieldTimeType)                       | [Time](ctp:api:type:Time)                                                                                                                                                   |
	//  [CustomFieldDateTimeType](ctp:api:type:CustomFieldDateTimeType)               | [DateTime](ctp:api:type:DateTime)                                                                                                                                           |
	//  [CustomFieldReferenceType](ctp:api:type:CustomFieldReferenceType)             | [Reference](/../api/types#reference)                                                                                                                                         |
	//  [CustomFieldSetType](ctp:api:type:CustomFieldSetType)                         | JSON array without duplicates consisting of [CustomFieldValues](ctp:api:type:CustomFieldValue) of a single [FieldType](ctp:api:type:FieldType). For example, a Custom Field of SetType of DateType takes a JSON array of mutually different Dates for its value. The order of items in the array is not fixed. For more examples, see the [example FieldContainer](ctp:api:type:FieldContainer).|
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
	CustomLineItemId string `json:"customLineItemId"`
	Name             string `json:"name"`
	// The value of a Custom Field.
	// The data type of the value depends on the specific [FieldType](/projects/types#fieldtype) given in the `type` field of the [FieldDefinition](/ctp:api:type:FieldDefinition) for a Custom Field.
	// It can be any of the following:
	//
	//  Field type                                                 | Data type                                                                                                                                                                 |
	//  ---------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
	//  [CustomFieldBooleanType](ctp:api:type:CustomFieldBooleanType)                 | Boolean (`true` or `false`)                                                                                                                                                     |
	//  [CustomFieldStringType](ctp:api:type:CustomFieldStringType)                   | String                                                                                                                                                                |
	//  [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType) | [LocalizedString](ctp:api:type:LocalizedString)                                                                                                                             |
	//  [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType)                       | String. Must be a `key` of one of the [EnumValues](ctp:api:type:CustomFieldEnumValue) defined in the [EnumType](ctp:api:type:CustomFiedEnumType)                                     |
	//  [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType)     | String. Must be a `key` of one of the [LocalizedEnumValues](ctp:api:type:CustomFieldLocalizedEnumValue) defined in the [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) |
	//  [CustomFieldNumberType](ctp:api:type:CustomFieldNumberType)                   | Number                                                                                                                                                                |
	//  [CustomFieldMoneyType](ctp:api:type:CustomFieldMoneyType)                     | [CentPrecisionMoney](/../api/types#centprecisionmoney)                                                                                                                                         |
	//  [CustomFieldDateType](ctp:api:type:CustomFieldDateType)                       | [Date](ctp:api:type:Date)                                                                                                                                                   |
	//  [CustomFieldTimeType](ctp:api:type:CustomFieldTimeType)                       | [Time](ctp:api:type:Time)                                                                                                                                                   |
	//  [CustomFieldDateTimeType](ctp:api:type:CustomFieldDateTimeType)               | [DateTime](ctp:api:type:DateTime)                                                                                                                                           |
	//  [CustomFieldReferenceType](ctp:api:type:CustomFieldReferenceType)             | [Reference](/../api/types#reference)                                                                                                                                         |
	//  [CustomFieldSetType](ctp:api:type:CustomFieldSetType)                         | JSON array without duplicates consisting of [CustomFieldValues](ctp:api:type:CustomFieldValue) of a single [FieldType](ctp:api:type:FieldType). For example, a Custom Field of SetType of DateType takes a JSON array of mutually different Dates for its value. The order of items in the array is not fixed. For more examples, see the [example FieldContainer](ctp:api:type:FieldContainer).|
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
	CustomLineItemId string `json:"customLineItemId"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) of a [Type](ctp:api:type:Type).
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
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
	CustomLineItemId string                    `json:"customLineItemId"`
	ShippingDetails  *ItemShippingDetailsDraft `json:"shippingDetails,omitempty"`
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

type StagedOrderSetCustomLineItemTaxAmountAction struct {
	CustomLineItemId  string                  `json:"customLineItemId"`
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
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

type StagedOrderSetCustomLineItemTaxRateAction struct {
	CustomLineItemId string                `json:"customLineItemId"`
	ExternalTaxRate  *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
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

type StagedOrderSetCustomShippingMethodAction struct {
	ShippingMethodName string            `json:"shippingMethodName"`
	ShippingRate       ShippingRateDraft `json:"shippingRate"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [TaxCategory](ctp:api:type:TaxCategory).
	TaxCategory     *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	ExternalTaxRate *ExternalTaxRateDraft          `json:"externalTaxRate,omitempty"`
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
	// [ResourceIdentifier](/../api/types#resourceidentifier) of a [Type](ctp:api:type:Type).
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
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

type StagedOrderSetCustomerEmailAction struct {
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

type StagedOrderSetCustomerGroupAction struct {
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [CustomerGroup](ctp:api:type:CustomerGroup).
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

type StagedOrderSetCustomerIdAction struct {
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

type StagedOrderSetDeliveryAddressAction struct {
	DeliveryId string       `json:"deliveryId"`
	Address    *BaseAddress `json:"address,omitempty"`
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
	DeliveryId string `json:"deliveryId"`
	Name       string `json:"name"`
	// The value of a Custom Field.
	// The data type of the value depends on the specific [FieldType](/projects/types#fieldtype) given in the `type` field of the [FieldDefinition](/ctp:api:type:FieldDefinition) for a Custom Field.
	// It can be any of the following:
	//
	//  Field type                                                 | Data type                                                                                                                                                                 |
	//  ---------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
	//  [CustomFieldBooleanType](ctp:api:type:CustomFieldBooleanType)                 | Boolean (`true` or `false`)                                                                                                                                                     |
	//  [CustomFieldStringType](ctp:api:type:CustomFieldStringType)                   | String                                                                                                                                                                |
	//  [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType) | [LocalizedString](ctp:api:type:LocalizedString)                                                                                                                             |
	//  [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType)                       | String. Must be a `key` of one of the [EnumValues](ctp:api:type:CustomFieldEnumValue) defined in the [EnumType](ctp:api:type:CustomFiedEnumType)                                     |
	//  [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType)     | String. Must be a `key` of one of the [LocalizedEnumValues](ctp:api:type:CustomFieldLocalizedEnumValue) defined in the [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) |
	//  [CustomFieldNumberType](ctp:api:type:CustomFieldNumberType)                   | Number                                                                                                                                                                |
	//  [CustomFieldMoneyType](ctp:api:type:CustomFieldMoneyType)                     | [CentPrecisionMoney](/../api/types#centprecisionmoney)                                                                                                                                         |
	//  [CustomFieldDateType](ctp:api:type:CustomFieldDateType)                       | [Date](ctp:api:type:Date)                                                                                                                                                   |
	//  [CustomFieldTimeType](ctp:api:type:CustomFieldTimeType)                       | [Time](ctp:api:type:Time)                                                                                                                                                   |
	//  [CustomFieldDateTimeType](ctp:api:type:CustomFieldDateTimeType)               | [DateTime](ctp:api:type:DateTime)                                                                                                                                           |
	//  [CustomFieldReferenceType](ctp:api:type:CustomFieldReferenceType)             | [Reference](/../api/types#reference)                                                                                                                                         |
	//  [CustomFieldSetType](ctp:api:type:CustomFieldSetType)                         | JSON array without duplicates consisting of [CustomFieldValues](ctp:api:type:CustomFieldValue) of a single [FieldType](ctp:api:type:FieldType). For example, a Custom Field of SetType of DateType takes a JSON array of mutually different Dates for its value. The order of items in the array is not fixed. For more examples, see the [example FieldContainer](ctp:api:type:FieldContainer).|
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
	DeliveryId string `json:"deliveryId"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) of a [Type](ctp:api:type:Type).
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
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
	DeliveryId string      `json:"deliveryId"`
	Name       string      `json:"name"`
	Value      interface{} `json:"value,omitempty"`
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
	DeliveryId string `json:"deliveryId"`
	// If set, the custom type is set to this new value.
	// If absent, the custom type and any existing custom fields are removed.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// If set, the custom fields are set to this new value.
	Fields *interface{} `json:"fields,omitempty"`
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

type StagedOrderSetDeliveryItemsAction struct {
	DeliveryId string         `json:"deliveryId"`
	Items      []DeliveryItem `json:"items"`
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

type StagedOrderSetItemShippingAddressCustomFieldAction struct {
	AddressKey string `json:"addressKey"`
	Name       string `json:"name"`
	// The value of a Custom Field.
	// The data type of the value depends on the specific [FieldType](/projects/types#fieldtype) given in the `type` field of the [FieldDefinition](/ctp:api:type:FieldDefinition) for a Custom Field.
	// It can be any of the following:
	//
	//  Field type                                                 | Data type                                                                                                                                                                 |
	//  ---------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
	//  [CustomFieldBooleanType](ctp:api:type:CustomFieldBooleanType)                 | Boolean (`true` or `false`)                                                                                                                                                     |
	//  [CustomFieldStringType](ctp:api:type:CustomFieldStringType)                   | String                                                                                                                                                                |
	//  [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType) | [LocalizedString](ctp:api:type:LocalizedString)                                                                                                                             |
	//  [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType)                       | String. Must be a `key` of one of the [EnumValues](ctp:api:type:CustomFieldEnumValue) defined in the [EnumType](ctp:api:type:CustomFiedEnumType)                                     |
	//  [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType)     | String. Must be a `key` of one of the [LocalizedEnumValues](ctp:api:type:CustomFieldLocalizedEnumValue) defined in the [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) |
	//  [CustomFieldNumberType](ctp:api:type:CustomFieldNumberType)                   | Number                                                                                                                                                                |
	//  [CustomFieldMoneyType](ctp:api:type:CustomFieldMoneyType)                     | [CentPrecisionMoney](/../api/types#centprecisionmoney)                                                                                                                                         |
	//  [CustomFieldDateType](ctp:api:type:CustomFieldDateType)                       | [Date](ctp:api:type:Date)                                                                                                                                                   |
	//  [CustomFieldTimeType](ctp:api:type:CustomFieldTimeType)                       | [Time](ctp:api:type:Time)                                                                                                                                                   |
	//  [CustomFieldDateTimeType](ctp:api:type:CustomFieldDateTimeType)               | [DateTime](ctp:api:type:DateTime)                                                                                                                                           |
	//  [CustomFieldReferenceType](ctp:api:type:CustomFieldReferenceType)             | [Reference](/../api/types#reference)                                                                                                                                         |
	//  [CustomFieldSetType](ctp:api:type:CustomFieldSetType)                         | JSON array without duplicates consisting of [CustomFieldValues](ctp:api:type:CustomFieldValue) of a single [FieldType](ctp:api:type:FieldType). For example, a Custom Field of SetType of DateType takes a JSON array of mutually different Dates for its value. The order of items in the array is not fixed. For more examples, see the [example FieldContainer](ctp:api:type:FieldContainer).|
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
	AddressKey string `json:"addressKey"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) of a [Type](ctp:api:type:Type).
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
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
	LineItemId string `json:"lineItemId"`
	Name       string `json:"name"`
	// The value of a Custom Field.
	// The data type of the value depends on the specific [FieldType](/projects/types#fieldtype) given in the `type` field of the [FieldDefinition](/ctp:api:type:FieldDefinition) for a Custom Field.
	// It can be any of the following:
	//
	//  Field type                                                 | Data type                                                                                                                                                                 |
	//  ---------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
	//  [CustomFieldBooleanType](ctp:api:type:CustomFieldBooleanType)                 | Boolean (`true` or `false`)                                                                                                                                                     |
	//  [CustomFieldStringType](ctp:api:type:CustomFieldStringType)                   | String                                                                                                                                                                |
	//  [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType) | [LocalizedString](ctp:api:type:LocalizedString)                                                                                                                             |
	//  [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType)                       | String. Must be a `key` of one of the [EnumValues](ctp:api:type:CustomFieldEnumValue) defined in the [EnumType](ctp:api:type:CustomFiedEnumType)                                     |
	//  [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType)     | String. Must be a `key` of one of the [LocalizedEnumValues](ctp:api:type:CustomFieldLocalizedEnumValue) defined in the [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) |
	//  [CustomFieldNumberType](ctp:api:type:CustomFieldNumberType)                   | Number                                                                                                                                                                |
	//  [CustomFieldMoneyType](ctp:api:type:CustomFieldMoneyType)                     | [CentPrecisionMoney](/../api/types#centprecisionmoney)                                                                                                                                         |
	//  [CustomFieldDateType](ctp:api:type:CustomFieldDateType)                       | [Date](ctp:api:type:Date)                                                                                                                                                   |
	//  [CustomFieldTimeType](ctp:api:type:CustomFieldTimeType)                       | [Time](ctp:api:type:Time)                                                                                                                                                   |
	//  [CustomFieldDateTimeType](ctp:api:type:CustomFieldDateTimeType)               | [DateTime](ctp:api:type:DateTime)                                                                                                                                           |
	//  [CustomFieldReferenceType](ctp:api:type:CustomFieldReferenceType)             | [Reference](/../api/types#reference)                                                                                                                                         |
	//  [CustomFieldSetType](ctp:api:type:CustomFieldSetType)                         | JSON array without duplicates consisting of [CustomFieldValues](ctp:api:type:CustomFieldValue) of a single [FieldType](ctp:api:type:FieldType). For example, a Custom Field of SetType of DateType takes a JSON array of mutually different Dates for its value. The order of items in the array is not fixed. For more examples, see the [example FieldContainer](ctp:api:type:FieldContainer).|
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
	LineItemId string `json:"lineItemId"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) of a [Type](ctp:api:type:Type).
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
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

type StagedOrderSetLineItemDistributionChannelAction struct {
	LineItemId string `json:"lineItemId"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [Channel](ctp:api:type:Channel).
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

type StagedOrderSetLineItemPriceAction struct {
	LineItemId string `json:"lineItemId"`
	// Draft type that stores amounts in cent precision for the specified currency.
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
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
	LineItemId      string                    `json:"lineItemId"`
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

type StagedOrderSetLineItemTaxAmountAction struct {
	LineItemId        string                  `json:"lineItemId"`
	ExternalTaxAmount *ExternalTaxAmountDraft `json:"externalTaxAmount,omitempty"`
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

type StagedOrderSetLineItemTaxRateAction struct {
	LineItemId      string                `json:"lineItemId"`
	ExternalTaxRate *ExternalTaxRateDraft `json:"externalTaxRate,omitempty"`
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

type StagedOrderSetLineItemTotalPriceAction struct {
	LineItemId         string                      `json:"lineItemId"`
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

type StagedOrderSetOrderTotalTaxAction struct {
	// Draft type that stores amounts in cent precision for the specified currency.
	// For storing money values in fractions of the minor unit in a currency, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft) instead.
	ExternalTotalGross  Money             `json:"externalTotalGross"`
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["externalTaxPortions"] == nil {
		delete(target, "externalTaxPortions")
	}

	return json.Marshal(target)
}

type StagedOrderSetParcelCustomFieldAction struct {
	ParcelId string      `json:"parcelId"`
	Name     string      `json:"name"`
	Value    interface{} `json:"value,omitempty"`
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
	ParcelId string `json:"parcelId"`
	// If set, the custom type is set to this new value.
	// If absent, the custom type and any existing custom fields are removed.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// If set, the custom fields are set to this new value.
	Fields *interface{} `json:"fields,omitempty"`
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

type StagedOrderSetParcelItemsAction struct {
	ParcelId string         `json:"parcelId"`
	Items    []DeliveryItem `json:"items"`
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

type StagedOrderSetParcelMeasurementsAction struct {
	ParcelId     string              `json:"parcelId"`
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

type StagedOrderSetParcelTrackingDataAction struct {
	ParcelId     string        `json:"parcelId"`
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

type StagedOrderSetReturnInfoAction struct {
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

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["items"] == nil {
		delete(target, "items")
	}

	return json.Marshal(target)
}

type StagedOrderSetReturnItemCustomFieldAction struct {
	ReturnItemId string      `json:"returnItemId"`
	Name         string      `json:"name"`
	Value        interface{} `json:"value,omitempty"`
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
	ReturnItemId string `json:"returnItemId"`
	// If set, the custom type is set to this new value.
	// If absent, the custom type and any existing custom fields are removed.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// If set, the custom fields are set to this new value.
	Fields *interface{} `json:"fields,omitempty"`
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

type StagedOrderSetReturnPaymentStateAction struct {
	ReturnItemId string             `json:"returnItemId"`
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

type StagedOrderSetReturnShipmentStateAction struct {
	ReturnItemId  string              `json:"returnItemId"`
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

type StagedOrderSetShippingAddressAction struct {
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

type StagedOrderSetShippingAddressAndCustomShippingMethodAction struct {
	Address            BaseAddress       `json:"address"`
	ShippingMethodName string            `json:"shippingMethodName"`
	ShippingRate       ShippingRateDraft `json:"shippingRate"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [TaxCategory](ctp:api:type:TaxCategory).
	TaxCategory     *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	ExternalTaxRate *ExternalTaxRateDraft          `json:"externalTaxRate,omitempty"`
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

type StagedOrderSetShippingAddressAndShippingMethodAction struct {
	Address         BaseAddress                       `json:"address"`
	ShippingMethod  *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	ExternalTaxRate *ExternalTaxRateDraft             `json:"externalTaxRate,omitempty"`
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
	Name string `json:"name"`
	// The value of a Custom Field.
	// The data type of the value depends on the specific [FieldType](/projects/types#fieldtype) given in the `type` field of the [FieldDefinition](/ctp:api:type:FieldDefinition) for a Custom Field.
	// It can be any of the following:
	//
	//  Field type                                                 | Data type                                                                                                                                                                 |
	//  ---------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
	//  [CustomFieldBooleanType](ctp:api:type:CustomFieldBooleanType)                 | Boolean (`true` or `false`)                                                                                                                                                     |
	//  [CustomFieldStringType](ctp:api:type:CustomFieldStringType)                   | String                                                                                                                                                                |
	//  [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType) | [LocalizedString](ctp:api:type:LocalizedString)                                                                                                                             |
	//  [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType)                       | String. Must be a `key` of one of the [EnumValues](ctp:api:type:CustomFieldEnumValue) defined in the [EnumType](ctp:api:type:CustomFiedEnumType)                                     |
	//  [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType)     | String. Must be a `key` of one of the [LocalizedEnumValues](ctp:api:type:CustomFieldLocalizedEnumValue) defined in the [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) |
	//  [CustomFieldNumberType](ctp:api:type:CustomFieldNumberType)                   | Number                                                                                                                                                                |
	//  [CustomFieldMoneyType](ctp:api:type:CustomFieldMoneyType)                     | [CentPrecisionMoney](/../api/types#centprecisionmoney)                                                                                                                                         |
	//  [CustomFieldDateType](ctp:api:type:CustomFieldDateType)                       | [Date](ctp:api:type:Date)                                                                                                                                                   |
	//  [CustomFieldTimeType](ctp:api:type:CustomFieldTimeType)                       | [Time](ctp:api:type:Time)                                                                                                                                                   |
	//  [CustomFieldDateTimeType](ctp:api:type:CustomFieldDateTimeType)               | [DateTime](ctp:api:type:DateTime)                                                                                                                                           |
	//  [CustomFieldReferenceType](ctp:api:type:CustomFieldReferenceType)             | [Reference](/../api/types#reference)                                                                                                                                         |
	//  [CustomFieldSetType](ctp:api:type:CustomFieldSetType)                         | JSON array without duplicates consisting of [CustomFieldValues](ctp:api:type:CustomFieldValue) of a single [FieldType](ctp:api:type:FieldType). For example, a Custom Field of SetType of DateType takes a JSON array of mutually different Dates for its value. The order of items in the array is not fixed. For more examples, see the [example FieldContainer](ctp:api:type:FieldContainer).|
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
	// [ResourceIdentifier](/../api/types#resourceidentifier) of a [Type](ctp:api:type:Type).
	Type   *TypeResourceIdentifier `json:"type,omitempty"`
	Fields *FieldContainer         `json:"fields,omitempty"`
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

type StagedOrderSetShippingMethodAction struct {
	ShippingMethod  *ShippingMethodResourceIdentifier `json:"shippingMethod,omitempty"`
	ExternalTaxRate *ExternalTaxRateDraft             `json:"externalTaxRate,omitempty"`
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

type StagedOrderSetShippingMethodTaxAmountAction struct {
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

type StagedOrderSetShippingMethodTaxRateAction struct {
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

type StagedOrderSetShippingRateInputAction struct {
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

type StagedOrderTransitionCustomLineItemStateAction struct {
	CustomLineItemId string `json:"customLineItemId"`
	Quantity         int    `json:"quantity"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [State](ctp:api:type:State).
	FromState StateResourceIdentifier `json:"fromState"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [State](ctp:api:type:State).
	ToState              StateResourceIdentifier `json:"toState"`
	ActualTransitionDate *time.Time              `json:"actualTransitionDate,omitempty"`
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

type StagedOrderTransitionLineItemStateAction struct {
	LineItemId string `json:"lineItemId"`
	Quantity   int    `json:"quantity"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [State](ctp:api:type:State).
	FromState StateResourceIdentifier `json:"fromState"`
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [State](ctp:api:type:State).
	ToState              StateResourceIdentifier `json:"toState"`
	ActualTransitionDate *time.Time              `json:"actualTransitionDate,omitempty"`
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

type StagedOrderTransitionStateAction struct {
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [State](ctp:api:type:State).
	State StateResourceIdentifier `json:"state"`
	Force *bool                   `json:"force,omitempty"`
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

type StagedOrderUpdateItemShippingAddressAction struct {
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
	// [ResourceIdentifier](/../api/types#resourceidentifier) to a [Channel](ctp:api:type:Channel).
	Channel    ChannelResourceIdentifier `json:"channel"`
	ExternalId *string                   `json:"externalId,omitempty"`
	SyncedAt   *time.Time                `json:"syncedAt,omitempty"`
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
