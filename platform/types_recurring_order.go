package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	Information about recurring orders and frequencies.
*
 */
type CustomLineItemRecurrenceInfo struct {
	// [Reference](ctp:api:type:Reference) to a RecurrencePolicy.
	RecurrencePolicy RecurrencePolicyReference `json:"recurrencePolicy"`
}

/**
*	Information about recurring orders and frequencies.
*
 */
type CustomLineItemRecurrenceInfoDraft struct {
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a RecurrencePolicy.
	RecurrencePolicy RecurrencePolicyResourceIdentifier `json:"recurrencePolicy"`
}

/**
*	Information about recurring orders and frequencies.
*
 */
type LineItemRecurrenceInfo struct {
	// [Reference](ctp:api:type:Reference) to a RecurrencePolicy.
	RecurrencePolicy RecurrencePolicyReference `json:"recurrencePolicy"`
	// Indicates how the price of a line item will be selected during order creation.
	PriceSelectionMode PriceSelectionMode `json:"priceSelectionMode"`
}

/**
*	Information about recurring orders and frequencies.
*
 */
type LineItemRecurrenceInfoDraft struct {
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a RecurrencePolicy.
	RecurrencePolicy RecurrencePolicyResourceIdentifier `json:"recurrencePolicy"`
	// Determines how the price of a line item will be selected during order creation.
	PriceSelectionMode PriceSelectionMode `json:"priceSelectionMode"`
}

/**
*	Indicates how the price of a [Line Item](ctp:api:type:LineItem) or [Custom Line Item](ctp:api:type:CustomLineItem) is selected during Order creation.
*
 */
type PriceSelectionMode string

const (
	PriceSelectionModeFixed   PriceSelectionMode = "Fixed"
	PriceSelectionModeDynamic PriceSelectionMode = "Dynamic"
)

type RecurringOrder struct {
	// Unique identifier of the RecurringOrder.
	ID string `json:"id"`
	// Current version of the RecurringOrder.
	Version int `json:"version"`
	// Date and time (UTC) when the RecurringOrder was created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) when the RecurringOrder was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// User-defined unique identifier of the RecurringOrder.
	Key *string `json:"key,omitempty"`
	// [Reference](ctp:api:type:Reference) to the Cart for a RecurringOrder.
	// The referenced Cart will have the `RecurringOrder` [CartOrigin](ctp:api:type:CartOrigin).
	Cart CartReference `json:"cart"`
	// [Reference](ctp:api:type:Reference) to the original [Order](ctp:api:type:Order) that generated this RecurringOrder.
	OriginOrder OrderReference `json:"originOrder"`
	// Date and time (UTC) when the RecurringOrder starts creating new Orders.
	StartsAt time.Time `json:"startsAt"`
	// Date and time (UTC) when the RecurringOrder resumes creating Orders after being unpaused.
	ResumesAt *time.Time `json:"resumesAt,omitempty"`
	// Date and time (UTC) when the RecurringOrder expires.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Date and time (UTC) when the last Order was created from this RecurringOrder.
	LastOrderAt *time.Time `json:"lastOrderAt,omitempty"`
	// Date and time (UTC) when the next Order will be created from this RecurringOrder.
	NextOrderAt *time.Time `json:"nextOrderAt,omitempty"`
	// Information about current and future skips for this RecurringOrder.
	SkipConfiguration SkipConfiguration `json:"skipConfiguration,omitempty"`
	// [Reference](ctp:api:type:Reference) to a Store.
	Store *StoreKeyReference `json:"store,omitempty"`
	// [Reference](ctp:api:type:Reference) to the Business Unit that the RecurringOrder belongs to.
	BusinessUnit *BusinessUnitKeyReference `json:"businessUnit,omitempty"`
	// [State](ctp:api:type:State) of the RecurringOrder in a custom workflow.
	State *StateReference `json:"state,omitempty"`
	// Current state of the RecurringOrder.
	RecurringOrderState RecurringOrderState `json:"recurringOrderState"`
	// Schedule of the RecurringOrder.
	Schedule RecurrencePolicySchedule `json:"schedule"`
	// The [Customer](ctp:api:type:Customer) that the RecurringOrder belongs to.
	Customer *CustomerReference `json:"customer,omitempty"`
	// Email address of the Customer that the RecurringOrder belongs to.
	CustomerEmail *string `json:"customerEmail,omitempty"`
	// Custom Fields of the RecurringOrder.
	Custom *CustomFields `json:"custom,omitempty"`
	// IDs and references that last modified the RecurringOrder.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the RecurringOrder.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *RecurringOrder) UnmarshalJSON(data []byte) error {
	type Alias RecurringOrder
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.SkipConfiguration != nil {
		var err error
		obj.SkipConfiguration, err = mapDiscriminatorSkipConfiguration(obj.SkipConfiguration)
		if err != nil {
			return err
		}
	}
	if obj.Schedule != nil {
		var err error
		obj.Schedule, err = mapDiscriminatorRecurrencePolicySchedule(obj.Schedule)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	Recurring Orders are automatically assigned the Store and/or Business Unit from the associated Cart.
*
 */
type RecurringOrderDraft struct {
	// User-defined unique identifier of the [RecurringOrder](ctp:api:type:RecurringOrder).
	Key *string `json:"key,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to the Cart from which the RecurringOrder is created.
	Cart CartResourceIdentifier `json:"cart"`
	// Current version of the referenced [Cart](ctp:api:type:Cart).
	CartVersion int `json:"cartVersion"`
	// Date and time (UTC) when the RecurringOrder will start.
	StartsAt time.Time `json:"startsAt"`
	// State for the RecurringOrder in a custom workflow.
	State *StateResourceIdentifier `json:"state,omitempty"`
	// Custom Fields for the RecurringOrder.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [RecurringOrder](ctp:api:type:RecurringOrder).
*
 */
type RecurringOrderPagedQueryResponse struct {
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
	// [RecurringOrders](ctp:api:type:RecurringOrder) matching the query.
	Results []RecurringOrder `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [RecurringOrder](ctp:api:type:RecurringOrder).
*
 */
type RecurringOrderReference struct {
	// Unique identifier of the referenced [RecurringOrder](ctp:api:type:RecurringOrder).
	ID string `json:"id"`
	// Contains the representation of the expanded RecurringOrder.
	// Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for RecurringOrders.
	Obj *RecurringOrder `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringOrderReference) MarshalJSON() ([]byte, error) {
	type Alias RecurringOrderReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "recurring-order", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [RecurringOrder](ctp:api:type:RecurringOrder). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type RecurringOrderResourceIdentifier struct {
	// Unique identifier of the referenced [RecurringOrder](ctp:api:type:RecurringOrder). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [RecurringOrder](ctp:api:type:RecurringOrder). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringOrderResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias RecurringOrderResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "recurring-order", Alias: (*Alias)(&obj)})
}

/**
*	Indicates the state of the [RecurringOrder](ctp:api:type:RecurringOrder).
*
 */
type RecurringOrderState string

const (
	RecurringOrderStateActive   RecurringOrderState = "Active"
	RecurringOrderStatePaused   RecurringOrderState = "Paused"
	RecurringOrderStateExpired  RecurringOrderState = "Expired"
	RecurringOrderStateCanceled RecurringOrderState = "Canceled"
	RecurringOrderStateFailed   RecurringOrderState = "Failed"
)

/**
*	Defines the new state for the Recurring Order—for possible values, see [RecurringOrderActive](ctp:api:type:RecurringOrderActive), [RecurringOrderPaused](ctp:api:type:RecurringOrderPaused), [RecurringOrderExpired](ctp:api:type:RecurringOrderExpired), and [RecurringOrderCanceled](ctp:api:type:RecurringOrderCanceled).
*
 */
type RecurringOrderStateDraft interface{}

func mapDiscriminatorRecurringOrderStateDraft(input interface{}) (RecurringOrderStateDraft, error) {
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
	case "active":
		obj := RecurringOrderActive{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "canceled":
		obj := RecurringOrderCanceled{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "expired":
		obj := RecurringOrderExpired{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "paused":
		obj := RecurringOrderPaused{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Changes the Recurring Order state to active.
*
 */
type RecurringOrderActive struct {
	// If set, the Recurring Order will automatically resume at the date and time (UTC) specified.
	ResumesAt *time.Time `json:"resumesAt,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringOrderActive) MarshalJSON() ([]byte, error) {
	type Alias RecurringOrderActive
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "active", Alias: (*Alias)(&obj)})
}

/**
*	Changes the Recurring Order state to canceled.
*
 */
type RecurringOrderCanceled struct {
	// The reason for the cancelation.
	Reason *string `json:"reason,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringOrderCanceled) MarshalJSON() ([]byte, error) {
	type Alias RecurringOrderCanceled
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "canceled", Alias: (*Alias)(&obj)})
}

/**
*	Changes the Recurring Order state to expired.
*
 */
type RecurringOrderExpired struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringOrderExpired) MarshalJSON() ([]byte, error) {
	type Alias RecurringOrderExpired
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "expired", Alias: (*Alias)(&obj)})
}

/**
*	Changes the Recurring Order state to paused.
*
 */
type RecurringOrderPaused struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringOrderPaused) MarshalJSON() ([]byte, error) {
	type Alias RecurringOrderPaused
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "paused", Alias: (*Alias)(&obj)})
}

type RecurringOrderUpdate struct {
	// Expected version of the RecurringOrder on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the RecurringOrder.
	Actions []RecurringOrderUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *RecurringOrderUpdate) UnmarshalJSON(data []byte) error {
	type Alias RecurringOrderUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorRecurringOrderUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type RecurringOrderUpdateAction interface{}

func mapDiscriminatorRecurringOrderUpdateAction(input interface{}) (RecurringOrderUpdateAction, error) {
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
	case "setCustomField":
		obj := RecurringOrderSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := RecurringOrderSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := RecurringOrderSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setOrderSkipConfiguration":
		obj := RecurringOrderSetOrderSkipConfigurationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.SkipConfiguration != nil {
			var err error
			obj.SkipConfiguration, err = mapDiscriminatorSkipConfigurationDraft(obj.SkipConfiguration)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "setSchedule":
		obj := RecurringOrderSetScheduleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setStartsAt":
		obj := RecurringOrderSetStartsAtAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setRecurringOrderState":
		obj := RecurringOrderSetStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.RecurringOrderState != nil {
			var err error
			obj.RecurringOrderState, err = mapDiscriminatorRecurringOrderStateDraft(obj.RecurringOrderState)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "transitionState":
		obj := RecurringOrderTransitionStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Defines how the next orders are going to be skipped.
*
 */
type SkipConfiguration interface{}

func mapDiscriminatorSkipConfiguration(input interface{}) (SkipConfiguration, error) {
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
	case "counter":
		obj := Counter{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Configuration to track skips for a [RecurringOrder](ctp:api:type:RecurringOrder).
*
 */
type Counter struct {
	// Number of Orders that will be skipped.
	TotalToSkip int `json:"totalToSkip"`
	// Number of Orders that were already skipped.
	Skipped int `json:"skipped"`
	// Date and time (UTC) when the last Order creation was skipped.
	LastSkippedAt *time.Time `json:"lastSkippedAt,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Counter) MarshalJSON() ([]byte, error) {
	type Alias Counter
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "counter", Alias: (*Alias)(&obj)})
}

/**
*	Defines how the next orders are going to be skipped.
*
 */
type SkipConfigurationDraft interface{}

func mapDiscriminatorSkipConfigurationDraft(input interface{}) (SkipConfigurationDraft, error) {
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
	case "counter":
		obj := CounterDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Configuration that uses a counter to track the number of Orders that will be skipped.
*
 */
type CounterDraft struct {
	// Number of Orders that will be skipped.
	TotalToSkip int `json:"totalToSkip"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CounterDraft) MarshalJSON() ([]byte, error) {
	type Alias CounterDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "counter", Alias: (*Alias)(&obj)})
}

/**
*	Adding a Custom Field to a Recurring Order generates the [RecurringOrderCustomFieldAdded](ctp:api:type:RecurringOrderCustomFieldAddedMessage) Message, removing one generates the [RecurringOrderCustomFieldRemoved](ctp:api:type:RecurringOrderCustomFieldRemovedMessage) Message, and updating an existing one generates the [RecurringOrderCustomFieldChanged](ctp:api:type:RecurringOrderCustomFieldChangedMessage) Message.
*
 */
type RecurringOrderSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringOrderSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias RecurringOrderSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

/**
*	Adding or updating a Custom Type on a Recurring Order generates the [RecurringOrderCustomTypeSet](ctp:api:type:RecurringOrderCustomTypeSetMessage) Message, removing one generates the [RecurringOrderCustomTypeRemoved](ctp:api:type:RecurringOrderCustomTypeRemovedMessage) Message.
*
 */
type RecurringOrderSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the RecurringOrder with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the RecurringOrder.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the RecurringOrder.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringOrderSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias RecurringOrderSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

/**
*	This update action generates the [RecurringOrderKeySet](ctp:api:type:RecurringOrderKeySetMessage) Message.
*
 */
type RecurringOrderSetKeyAction struct {
	// Value to set.
	// If empty, any existing key will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringOrderSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias RecurringOrderSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type RecurringOrderSetOrderSkipConfigurationAction struct {
	// Configuration for skipping the next orders of the [Recurring Order](ctp:api:type:RecurringOrder).
	SkipConfiguration SkipConfigurationDraft `json:"skipConfiguration,omitempty"`
	// Date and time (UTC) the Recurring Order will resume and start to generate new orders.
	UpdatedExpiresAt *time.Time `json:"updatedExpiresAt,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *RecurringOrderSetOrderSkipConfigurationAction) UnmarshalJSON(data []byte) error {
	type Alias RecurringOrderSetOrderSkipConfigurationAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.SkipConfiguration != nil {
		var err error
		obj.SkipConfiguration, err = mapDiscriminatorSkipConfigurationDraft(obj.SkipConfiguration)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringOrderSetOrderSkipConfigurationAction) MarshalJSON() ([]byte, error) {
	type Alias RecurringOrderSetOrderSkipConfigurationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setOrderSkipConfiguration", Alias: (*Alias)(&obj)})
}

/**
*	To set the [schedule](ctp:api:type:RecurrencePolicySchedule), the [Recurring Order](ctp:api:type:RecurringOrder) must be active, with no active [Skip Configuration](ctp:api:type:SkipConfiguration) and with available prices for all Cart items for the new schedule.
*	Setting the schedule generates the [RecurringOrderScheduleSet](ctp:api:type:RecurringOrderScheduleSetMessage) Message.
*
 */
type RecurringOrderSetScheduleAction struct {
	// Value to set.
	RecurrencePolicy RecurrencePolicyResourceIdentifier `json:"recurrencePolicy"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringOrderSetScheduleAction) MarshalJSON() ([]byte, error) {
	type Alias RecurringOrderSetScheduleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSchedule", Alias: (*Alias)(&obj)})
}

/**
*	To set the start date and time, the [Recurring Order](ctp:api:type:RecurringOrder) must not have been started yet.
*	Setting the start date and time generates the [RecurringOrderStartsAtSet](ctp:api:type:RecurringOrderStartsAtSetMessage) Message.
*
 */
type RecurringOrderSetStartsAtAction struct {
	// Date and time (UTC) the [Recurring Order](ctp:api:type:RecurringOrder) should be started. The date and time must be in the future.
	StartsAt time.Time `json:"startsAt"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringOrderSetStartsAtAction) MarshalJSON() ([]byte, error) {
	type Alias RecurringOrderSetStartsAtAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStartsAt", Alias: (*Alias)(&obj)})
}

/**
*	Setting the [RecurringOrderState](ctp:api:type:RecurringOrderState) generates the [RecurringOrderStateChanged](ctp:api:type:RecurringOrderStateChangedMessage) Message.
*
 */
type RecurringOrderSetStateAction struct {
	// New state of the RecurringOrder.
	RecurringOrderState RecurringOrderStateDraft `json:"recurringOrderState"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *RecurringOrderSetStateAction) UnmarshalJSON(data []byte) error {
	type Alias RecurringOrderSetStateAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.RecurringOrderState != nil {
		var err error
		obj.RecurringOrderState, err = mapDiscriminatorRecurringOrderStateDraft(obj.RecurringOrderState)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringOrderSetStateAction) MarshalJSON() ([]byte, error) {
	type Alias RecurringOrderSetStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRecurringOrderState", Alias: (*Alias)(&obj)})
}

/**
*	If the existing [State](ctp:api:type:State) has set `transitions`, there must be a direct transition to the new State.
*	If `transitions` is not set, no validation is performed.
*
*	This update action produces the [Recurring Order State Transition](ctp:api:type:RecurringOrderStateTransitionMessage) Message.
*
 */
type RecurringOrderTransitionStateAction struct {
	// Value to set.
	// If there is no State yet, the new State must be an initial State.
	State StateResourceIdentifier `json:"state"`
	// Set to `true` to turn off validation.
	Force *bool `json:"force,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurringOrderTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias RecurringOrderTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}
