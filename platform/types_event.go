package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	Base representation of an Event containing common fields to all [Event Types](#eventtype).
*
 */
type BaseEvent struct {
	// Unique identifier of the Event.
	ID               string `json:"id"`
	NotificationType string `json:"notificationType"`
	// The type of resource targeted by the Event.
	ResourceType EventSubscriptionResourceTypeId `json:"resourceType"`
	// The type of Event that has occurred.
	Type EventType `json:"type"`
	// An object containing details related to the Event.
	Data interface{} `json:"data"`
	// Date and time (UTC) the Event was generated.
	CreatedAt time.Time `json:"createdAt"`
}

/**
*	Base representation of an Event containing common fields to all [Event Types](#eventtype).
*
 */
type Event interface{}

func mapDiscriminatorEvent(input interface{}) (Event, error) {
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
	case "CheckoutOrderCreationFailed":
		obj := CheckoutOrderCreationFailedEvent{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CheckoutPaymentAuthorizationCancelled":
		obj := CheckoutPaymentAuthorizationCancelledEvent{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CheckoutPaymentAuthorizationFailed":
		obj := CheckoutPaymentAuthorizationFailedEvent{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CheckoutPaymentAuthorized":
		obj := CheckoutPaymentAuthorizedEvent{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CheckoutPaymentCancelAuthorizationFailed":
		obj := CheckoutPaymentCancelAuthorizationFailedEvent{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CheckoutPaymentChargeFailed":
		obj := CheckoutPaymentChargeFailedEvent{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CheckoutPaymentCharged":
		obj := CheckoutPaymentChargedEvent{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CheckoutPaymentRefundFailed":
		obj := CheckoutPaymentRefundFailedEvent{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CheckoutPaymentRefunded":
		obj := CheckoutPaymentRefundedEvent{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ImportContainerCreated":
		obj := ImportContainerCreatedEvent{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ImportContainerDeleted":
		obj := ImportContainerDeletedEvent{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ImportOperationRejected":
		obj := ImportOperationRejectedEvent{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ImportUnresolved":
		obj := ImportUnresolvedEvent{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ImportValidationFailed":
		obj := ImportValidationFailedEvent{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "ImportWaitForMasterVariant":
		obj := ImportWaitForMasterVariantEvent{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Generated when an order creation attempt fails in Checkout. This event includes information about why the order could not be created.
*
 */
type CheckoutOrderCreationFailedEvent struct {
	// Unique identifier of the Event.
	ID               string                          `json:"id"`
	NotificationType string                          `json:"notificationType"`
	ResourceType     EventSubscriptionResourceTypeId `json:"resourceType"`
	// Date and time (UTC) the Event was generated.
	CreatedAt time.Time `json:"createdAt"`
	// An object containing details of the order which could not be created.
	Data CheckoutMessageOrderPayloadBaseData `json:"data"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CheckoutOrderCreationFailedEvent) MarshalJSON() ([]byte, error) {
	type Alias CheckoutOrderCreationFailedEvent
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CheckoutOrderCreationFailed", Alias: (*Alias)(&obj)})
}

/**
*	Generated when a payment authorization is successfully cancelled in Checkout. This event indicates that the payment is cancelled before it is charged.
*
 */
type CheckoutPaymentAuthorizationCancelledEvent struct {
	// Unique identifier of the Event.
	ID               string                          `json:"id"`
	NotificationType string                          `json:"notificationType"`
	ResourceType     EventSubscriptionResourceTypeId `json:"resourceType"`
	// Date and time (UTC) the Event was generated.
	CreatedAt time.Time `json:"createdAt"`
	// An object containing details of the payment authorization that was cancelled.
	Data CheckoutMessagePaymentsPayloadBaseData `json:"data"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CheckoutPaymentAuthorizationCancelledEvent) MarshalJSON() ([]byte, error) {
	type Alias CheckoutPaymentAuthorizationCancelledEvent
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CheckoutPaymentAuthorizationCancelled", Alias: (*Alias)(&obj)})
}

/**
*	Generated when an attempt to authorize a payment fails in Checkout. This failure could result from insufficient funds, incorrect payment details, expired cards, or risk related rejections.
*
 */
type CheckoutPaymentAuthorizationFailedEvent struct {
	// Unique identifier of the Event.
	ID               string                          `json:"id"`
	NotificationType string                          `json:"notificationType"`
	ResourceType     EventSubscriptionResourceTypeId `json:"resourceType"`
	// Date and time (UTC) the Event was generated.
	CreatedAt time.Time `json:"createdAt"`
	// An object containing details of the payment authorization that failed.
	Data CheckoutMessagePaymentsPayloadBaseData `json:"data"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CheckoutPaymentAuthorizationFailedEvent) MarshalJSON() ([]byte, error) {
	type Alias CheckoutPaymentAuthorizationFailedEvent
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CheckoutPaymentAuthorizationFailed", Alias: (*Alias)(&obj)})
}

/**
*	Generated when a payment is successfully authorized in Checkout. This event indicates the payment has been validated and the amount has been reserved but not yet charged.
*
 */
type CheckoutPaymentAuthorizedEvent struct {
	// Unique identifier of the Event.
	ID               string                          `json:"id"`
	NotificationType string                          `json:"notificationType"`
	ResourceType     EventSubscriptionResourceTypeId `json:"resourceType"`
	// Date and time (UTC) the Event was generated.
	CreatedAt time.Time `json:"createdAt"`
	// An object containing details of the successful payment authorization.
	Data CheckoutMessagePaymentsPayloadBaseData `json:"data"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CheckoutPaymentAuthorizedEvent) MarshalJSON() ([]byte, error) {
	type Alias CheckoutPaymentAuthorizedEvent
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CheckoutPaymentAuthorized", Alias: (*Alias)(&obj)})
}

/**
*	Generated when an attempt to cancel a payment authorization fails in Checkout. This could happen if the authorization has already been expired, been captured already or no longer valid.
*
 */
type CheckoutPaymentCancelAuthorizationFailedEvent struct {
	// Unique identifier of the Event.
	ID               string                          `json:"id"`
	NotificationType string                          `json:"notificationType"`
	ResourceType     EventSubscriptionResourceTypeId `json:"resourceType"`
	// Date and time (UTC) the Event was generated.
	CreatedAt time.Time `json:"createdAt"`
	// An object containing details of the payment authorization that could not be cancelled.
	Data CheckoutMessagePaymentsPayloadBaseData `json:"data"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CheckoutPaymentCancelAuthorizationFailedEvent) MarshalJSON() ([]byte, error) {
	type Alias CheckoutPaymentCancelAuthorizationFailedEvent
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CheckoutPaymentCancelAuthorizationFailed", Alias: (*Alias)(&obj)})
}

/**
*	Generated when an attempt to charge a payment fails in Checkout. Even if a payment was previously authorized, charging it may still fail.
*
 */
type CheckoutPaymentChargeFailedEvent struct {
	// Unique identifier of the Event.
	ID               string                          `json:"id"`
	NotificationType string                          `json:"notificationType"`
	ResourceType     EventSubscriptionResourceTypeId `json:"resourceType"`
	// Date and time (UTC) the Event was generated.
	CreatedAt time.Time `json:"createdAt"`
	// An object containing details of the failed payment charge.
	Data CheckoutMessagePaymentsPayloadBaseData `json:"data"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CheckoutPaymentChargeFailedEvent) MarshalJSON() ([]byte, error) {
	type Alias CheckoutPaymentChargeFailedEvent
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CheckoutPaymentChargeFailed", Alias: (*Alias)(&obj)})
}

/**
*	Generated when a payment is successfully charged in Checkout. This event indicates that the authorized amount has been successfully debited from your customer's account.
*
 */
type CheckoutPaymentChargedEvent struct {
	// Unique identifier of the Event.
	ID               string                          `json:"id"`
	NotificationType string                          `json:"notificationType"`
	ResourceType     EventSubscriptionResourceTypeId `json:"resourceType"`
	// Date and time (UTC) the Event was generated.
	CreatedAt time.Time `json:"createdAt"`
	// An object containing details of the successful payment charge.
	Data CheckoutMessagePaymentsPayloadBaseData `json:"data"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CheckoutPaymentChargedEvent) MarshalJSON() ([]byte, error) {
	type Alias CheckoutPaymentChargedEvent
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CheckoutPaymentCharged", Alias: (*Alias)(&obj)})
}

/**
*	Generated when an attempt to refund a payment refund fails in Checkout. This failure indicates that the planned refund amount was not successfully sent to your customer’s account.
*
 */
type CheckoutPaymentRefundFailedEvent struct {
	// Unique identifier of the Event.
	ID               string                          `json:"id"`
	NotificationType string                          `json:"notificationType"`
	ResourceType     EventSubscriptionResourceTypeId `json:"resourceType"`
	// Date and time (UTC) the Event was generated.
	CreatedAt time.Time `json:"createdAt"`
	// An object containing details of the failed payment refund attempt.
	Data CheckoutMessagePaymentsPayloadBaseData `json:"data"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CheckoutPaymentRefundFailedEvent) MarshalJSON() ([]byte, error) {
	type Alias CheckoutPaymentRefundFailedEvent
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CheckoutPaymentRefundFailed", Alias: (*Alias)(&obj)})
}

/**
*	Generated when a payment is successfully refunded in Checkout. This event confirms that the refund has been processed and sent to your customer.
*
 */
type CheckoutPaymentRefundedEvent struct {
	// Unique identifier of the Event.
	ID               string                          `json:"id"`
	NotificationType string                          `json:"notificationType"`
	ResourceType     EventSubscriptionResourceTypeId `json:"resourceType"`
	// Date and time (UTC) the Event was generated.
	CreatedAt time.Time `json:"createdAt"`
	// An object containing details of the successful payment refund.
	Data CheckoutMessagePaymentsPayloadBaseData `json:"data"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CheckoutPaymentRefundedEvent) MarshalJSON() ([]byte, error) {
	type Alias CheckoutPaymentRefundedEvent
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CheckoutPaymentRefunded", Alias: (*Alias)(&obj)})
}

/**
*	Generated when an [Import Container](ctp:import:type:ImportContainer) is created.
 */
type ImportContainerCreatedEvent struct {
	// Unique identifier of the Event.
	ID               string                          `json:"id"`
	NotificationType string                          `json:"notificationType"`
	ResourceType     EventSubscriptionResourceTypeId `json:"resourceType"`
	// Date and time (UTC) the Event was generated.
	CreatedAt time.Time `json:"createdAt"`
	// An object containing details of the created Import Container.
	Data ImportContainerCreatedEventData `json:"data"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ImportContainerCreatedEvent) MarshalJSON() ([]byte, error) {
	type Alias ImportContainerCreatedEvent
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ImportContainerCreated", Alias: (*Alias)(&obj)})
}

/**
*	Generated when an [Import Container](ctp:import:type:ImportContainer) is deleted.
 */
type ImportContainerDeletedEvent struct {
	// Unique identifier of the Event.
	ID               string                          `json:"id"`
	NotificationType string                          `json:"notificationType"`
	ResourceType     EventSubscriptionResourceTypeId `json:"resourceType"`
	// Date and time (UTC) the Event was generated.
	CreatedAt time.Time `json:"createdAt"`
	// An object containing details of the deleted Import Container.
	Data ImportContainerDeletedEventData `json:"data"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ImportContainerDeletedEvent) MarshalJSON() ([]byte, error) {
	type Alias ImportContainerDeletedEvent
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ImportContainerDeleted", Alias: (*Alias)(&obj)})
}

/**
*	Generated when an [Import Operation](ctp:import:type:ImportOperation) has the `rejected` [ProcessingState](ctp:import:type:ProcessingState).
 */
type ImportOperationRejectedEvent struct {
	// Unique identifier of the Event.
	ID               string                          `json:"id"`
	NotificationType string                          `json:"notificationType"`
	ResourceType     EventSubscriptionResourceTypeId `json:"resourceType"`
	// Date and time (UTC) the Event was generated.
	CreatedAt time.Time `json:"createdAt"`
	// An object containing details of the Import Operation with the `rejected` state.
	Data ImportOperationRejectedEventData `json:"data"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ImportOperationRejectedEvent) MarshalJSON() ([]byte, error) {
	type Alias ImportOperationRejectedEvent
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ImportOperationRejected", Alias: (*Alias)(&obj)})
}

/**
*	Generated when an [Import Operation](ctp:import:type:ImportOperation) has the `unresolved` [ProcessingState](ctp:import:type:ProcessingState).
 */
type ImportUnresolvedEvent struct {
	// Unique identifier of the Event.
	ID               string                          `json:"id"`
	NotificationType string                          `json:"notificationType"`
	ResourceType     EventSubscriptionResourceTypeId `json:"resourceType"`
	// Date and time (UTC) the Event was generated.
	CreatedAt time.Time `json:"createdAt"`
	// An object containing details of the Import Operation with the `unresolved` state.
	Data ImportUnresolvedEventData `json:"data"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ImportUnresolvedEvent) MarshalJSON() ([]byte, error) {
	type Alias ImportUnresolvedEvent
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ImportUnresolved", Alias: (*Alias)(&obj)})
}

/**
*	Generated when an [Import Operation](ctp:import:type:ImportOperation) has the `validationFailed` [ProcessingState](ctp:import:type:ProcessingState).
 */
type ImportValidationFailedEvent struct {
	// Unique identifier of the Event.
	ID               string                          `json:"id"`
	NotificationType string                          `json:"notificationType"`
	ResourceType     EventSubscriptionResourceTypeId `json:"resourceType"`
	// Date and time (UTC) the Event was generated.
	CreatedAt time.Time `json:"createdAt"`
	// An object containing details of the Import Operation with the `validationFailed` state.
	Data ImportValidationFailedEventData `json:"data"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ImportValidationFailedEvent) MarshalJSON() ([]byte, error) {
	type Alias ImportValidationFailedEvent
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ImportValidationFailed", Alias: (*Alias)(&obj)})
}

/**
*	Generated when an [Import Operation](ctp:import:type:ImportOperation) has the `waitForMasterVariant` [ProcessingState](ctp:import:type:ProcessingState).
 */
type ImportWaitForMasterVariantEvent struct {
	// Unique identifier of the Event.
	ID               string                          `json:"id"`
	NotificationType string                          `json:"notificationType"`
	ResourceType     EventSubscriptionResourceTypeId `json:"resourceType"`
	// Date and time (UTC) the Event was generated.
	CreatedAt time.Time `json:"createdAt"`
	// An object containing details of the Import Operation with the `waitForMasterVariant` state.
	Data ImportWaitForMasterVariantEventData `json:"data"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ImportWaitForMasterVariantEvent) MarshalJSON() ([]byte, error) {
	type Alias ImportWaitForMasterVariantEvent
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "ImportWaitForMasterVariant", Alias: (*Alias)(&obj)})
}

/**
*	The `data` payload of all related order event messages.
 */
type CheckoutMessageOrderPayloadBaseData struct {
	// `key` of the [Project](ctp:api:type:Project) where the order would belong to.
	ProjectKey string `json:"projectKey"`
	// The [Cart](ctp:api:type:Cart) on which the change or action was performed.
	Cart CartReference `json:"cart"`
	// The [Payments](ctp:api:type:Payment) on which the change or action was performed.
	Payments []PaymentReference `json:"payments"`
	// Errors associated with the order event.
	Errors []ErrorObject `json:"errors"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CheckoutMessageOrderPayloadBaseData) UnmarshalJSON(data []byte) error {
	type Alias CheckoutMessageOrderPayloadBaseData
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

/**
*	The `data` payload of all payment related event messages.
 */
type CheckoutMessagePaymentsPayloadBaseData struct {
	// `key` of the [Project](ctp:api:type:Project) where the payment was made.
	ProjectKey string `json:"projectKey"`
	// The [Payment](ctp:api:type:Payment) on which the change or action was performed.
	Payment PaymentReference `json:"payment"`
	// `id` of the [Transaction](/../api/projects/payments#transaction).
	TransactionId string `json:"transactionId"`
	// The [Cart](ctp:api:type:Cart) on which the change or action was performed.
	Cart *CartReference `json:"cart,omitempty"`
	// The [Order](ctp:api:type:Order) on which the change or action was performed.
	Order *OrderReference `json:"order,omitempty"`
}

/**
*	The `data` of the [Import Container Created Event](ctp:api:type:ImportContainerCreatedEvent).
 */
type ImportContainerCreatedEventData struct {
	// The `key` of the created Import Container.
	Key string `json:"key"`
	// The `version` of the created Import Container.
	Version int `json:"version"`
	// Date and time (UTC) the Import Container was created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Import Container was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
}

/**
*	The `data` of the [Import Container Deleted Event](ctp:api:type:ImportContainerDeletedEvent).
 */
type ImportContainerDeletedEventData struct {
	// The `key` of the deleted Import Container.
	Key string `json:"key"`
	// The `version` of the deleted Import Container.
	Version int `json:"version"`
}

/**
*	The `data` of the [Import Operation Rejected Event](ctp:api:type:ImportOperationRejectedEvent).
 */
type ImportOperationRejectedEventData struct {
	// The `id` of the Import Operation with the `rejected` state.
	ID string `json:"id"`
}

/**
*	The `data` of the [Import Unresolved Event](ctp:api:type:ImportUnresolvedEvent).
 */
type ImportUnresolvedEventData struct {
	// The `id` of the Import Operation with the `unresolved` state.
	ID string `json:"id"`
	// The `version` of the Import Operation with the `unresolved` state.
	Version int `json:"version"`
	// The `key` of the Import Container.
	ImportContainerKey string `json:"importContainerKey"`
}

/**
*	The `data` of the [Import Validation Failed Event](ctp:api:type:ImportValidationFailedEvent).
 */
type ImportValidationFailedEventData struct {
	// The `id` of the Import Operation with the `validationFailed` state.
	ID string `json:"id"`
	// The `version` of the Import Operation with the `validationFailed` state.
	Version int `json:"version"`
	// The `key` of the Import Container.
	ImportContainerKey string `json:"importContainerKey"`
}

/**
*	The `data` of the [Import Wait For Master Variant Event](ctp:api:type:ImportWaitForMasterVariantEvent).
 */
type ImportWaitForMasterVariantEventData struct {
	// The `id` of the Import Operation with the `waitForMasterVariant` state.
	ID string `json:"id"`
	// The `version` of the Import Operation with the `waitForMasterVariant` state.
	Version int `json:"version"`
	// The `key` of the Import Container.
	ImportContainerKey string `json:"importContainerKey"`
}
