package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type Payment struct {
	// Unique identifier of the Payment.
	ID string `json:"id"`
	// Current version of the Payment.
	Version int `json:"version"`
	// Date and time (UTC) the Payment was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Payment was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/general-concepts#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/general-concepts#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Reference to a [Customer](ctp:api:type:Customer) associated with the Payment.
	Customer *CustomerReference `json:"customer,omitempty"`
	// [Anonymous session](ctp:api:type:AnonymousSession) associated with the Payment.
	AnonymousId *string `json:"anonymousId,omitempty"`
	// Additional identifier for external systems like Customer Relationship Management (CRM) or Enterprise Resource Planning (ERP).
	ExternalId *string `json:"externalId,omitempty"`
	// Identifier used by the payment service that processes the Payment (for example, a PSP).
	// The combination of `interfaceId` and the `paymentInterface` field on [PaymentMethodInfo](ctp:api:type:PaymentMethodInfo) must be unique.
	InterfaceId *string `json:"interfaceId,omitempty"`
	// Money value the Payment intends to receive from the customer.
	// The value typically matches the [Cart](ctp:api:type:Cart) or [Order](ctp:api:type:Order) gross total.
	AmountPlanned CentPrecisionMoney `json:"amountPlanned"`
	// Deprecated because its value can be calculated from the total amounts saved in the [Transactions](ctp:api:type:Transaction).
	AmountAuthorized TypedMoney `json:"amountAuthorized,omitempty"`
	// Deprecated because this field is of little practical value, as it is either not reliably known, or the authorization time is fixed for a PSP.
	AuthorizedUntil *string `json:"authorizedUntil,omitempty"`
	// Deprecated because its value can be calculated from the total amounts saved in the [Transactions](ctp:api:type:Transaction).
	AmountPaid TypedMoney `json:"amountPaid,omitempty"`
	// Deprecated because its value can be calculated from the total amounts saved in the [Transactions](ctp:api:type:Transaction).
	AmountRefunded TypedMoney `json:"amountRefunded,omitempty"`
	// Information regarding the payment interface (for example, a PSP), and the specific payment method used.
	PaymentMethodInfo PaymentMethodInfo `json:"paymentMethodInfo"`
	// Current status of the Payment.
	PaymentStatus PaymentStatus `json:"paymentStatus"`
	// Financial transactions of the Payment. Each Transaction has a [TransactionType](ctp:api:type:TransactionType) and a [TransactionState](ctp:api:type:TransactionState).
	Transactions []Transaction `json:"transactions"`
	// Represents information exchange with the payment service, for example, a PSP. An interaction may be a request sent, or a response or notification received from the payment service.
	InterfaceInteractions []CustomFields `json:"interfaceInteractions"`
	// Custom Fields for the Payment.
	Custom *CustomFields `json:"custom,omitempty"`
	// User-defined unique identifier of the Payment.
	Key *string `json:"key,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Payment) UnmarshalJSON(data []byte) error {
	type Alias Payment
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.AmountAuthorized != nil {
		var err error
		obj.AmountAuthorized, err = mapDiscriminatorTypedMoney(obj.AmountAuthorized)
		if err != nil {
			return err
		}
	}
	if obj.AmountPaid != nil {
		var err error
		obj.AmountPaid, err = mapDiscriminatorTypedMoney(obj.AmountPaid)
		if err != nil {
			return err
		}
	}
	if obj.AmountRefunded != nil {
		var err error
		obj.AmountRefunded, err = mapDiscriminatorTypedMoney(obj.AmountRefunded)
		if err != nil {
			return err
		}
	}

	return nil
}

type PaymentDraft struct {
	// Reference to a [Customer](ctp:api:type:Customer) associated with the Payment.
	Customer *CustomerResourceIdentifier `json:"customer,omitempty"`
	// [Anonymous session](ctp:api:type:AnonymousSession) associated with the Payment.
	AnonymousId *string `json:"anonymousId,omitempty"`
	// Additional identifier for external systems like Customer Relationship Management (CRM) or Enterprise Resource Planning (ERP).
	ExternalId *string `json:"externalId,omitempty"`
	// Identifier used by the payment service that processes the Payment (for example, a PSP).
	// The combination of `interfaceId` and the `paymentInterface` field on [PaymentMethodInfo](ctp:api:type:PaymentMethodInfo) must be unique.
	// Once set, it cannot be changed.
	InterfaceId *string `json:"interfaceId,omitempty"`
	// Money value the Payment intends to receive from the customer.
	// The value typically matches the [Cart](ctp:api:type:Cart) or [Order](ctp:api:type:Order) gross total.
	AmountPlanned Money `json:"amountPlanned"`
	// Deprecated because the value can be calculated from the total amounts saved in the [Transactions](ctp:api:type:Transaction).
	AmountAuthorized *Money `json:"amountAuthorized,omitempty"`
	// Deprecated because this field is of little practical value, as it is either not reliably known, or the authorization time is fixed for a PSP.
	AuthorizedUntil *string `json:"authorizedUntil,omitempty"`
	// Deprecated because the value can be calculated from the total amounts saved in the [Transactions](ctp:api:type:Transaction).
	AmountPaid *Money `json:"amountPaid,omitempty"`
	// Deprecated because the value can be calculated from the total amounts saved in the [Transactions](ctp:api:type:Transaction).
	AmountRefunded *Money `json:"amountRefunded,omitempty"`
	// Information regarding the payment interface (for example, a PSP), and the specific payment method used.
	PaymentMethodInfo *PaymentMethodInfo `json:"paymentMethodInfo,omitempty"`
	// Current status of the Payment.
	PaymentStatus *PaymentStatusDraft `json:"paymentStatus,omitempty"`
	// Financial transactions of the Payment. Each Transaction has a [TransactionType](ctp:api:type:TransactionType) and a [TransactionState](ctp:api:type:TransactionState).
	Transactions []TransactionDraft `json:"transactions"`
	// Represents information exchange with the payment service, for example, a PSP. An interaction may be a request sent, or a response or notification received from the payment service.
	InterfaceInteractions []CustomFieldsDraft `json:"interfaceInteractions"`
	// Custom Fields for the Payment.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// User-defined unique identifier for the Payment.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentDraft) MarshalJSON() ([]byte, error) {
	type Alias PaymentDraft
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

	if raw["transactions"] == nil {
		delete(raw, "transactions")
	}

	if raw["interfaceInteractions"] == nil {
		delete(raw, "interfaceInteractions")
	}

	return json.Marshal(raw)

}

type PaymentMethodInfo struct {
	// Payment service that processes the Payment (for example, a PSP).
	// Once set, it cannot be changed.
	// The combination of `paymentInterface` and the `interfaceId` of a [Payment](ctp:api:type:Payment) must be unique.
	PaymentInterface *string `json:"paymentInterface,omitempty"`
	// Payment method used, for example, credit card, or cash advance.
	Method *string `json:"method,omitempty"`
	// Localizable name of the payment method.
	Name *LocalizedString `json:"name,omitempty"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [Payment](ctp:api:type:Payment).
*
 */
type PaymentPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// [Payments](ctp:api:type:Payment) matching the query.
	Results []Payment `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [Payment](ctp:api:type:Payment).
*
 */
type PaymentReference struct {
	// Unique identifier of the referenced [Payment](ctp:api:type:Payment).
	ID string `json:"id"`
	// Contains the representation of the expanded Payment. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for Payments.
	Obj *Payment `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentReference) MarshalJSON() ([]byte, error) {
	type Alias PaymentReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "payment", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) of a [Payment](ctp:api:type:Payment). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type PaymentResourceIdentifier struct {
	// Unique identifier of the referenced [Payment](ctp:api:type:Payment). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [Payment](ctp:api:type:Payment). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias PaymentResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "payment", Alias: (*Alias)(&obj)})
}

type PaymentStatus struct {
	// External reference that identifies the current status of the Payment.
	InterfaceCode *string `json:"interfaceCode,omitempty"`
	// Text describing the current status of the Payment.
	InterfaceText *string `json:"interfaceText,omitempty"`
	// [Reference](ctp:api:type:Reference) to a [State](ctp:api:type:State).
	State *StateReference `json:"state,omitempty"`
}

type PaymentStatusDraft struct {
	// External reference that identifies the current status of the Payment.
	InterfaceCode *string `json:"interfaceCode,omitempty"`
	// Text describing the current status of the Payment.
	InterfaceText *string `json:"interfaceText,omitempty"`
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [State](ctp:api:type:State).
	State *StateResourceIdentifier `json:"state,omitempty"`
}

type PaymentUpdate struct {
	// Expected version of the Payment on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the Payment.
	Actions []PaymentUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PaymentUpdate) UnmarshalJSON(data []byte) error {
	type Alias PaymentUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorPaymentUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type PaymentUpdateAction interface{}

func mapDiscriminatorPaymentUpdateAction(input interface{}) (PaymentUpdateAction, error) {
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
	case "addInterfaceInteraction":
		obj := PaymentAddInterfaceInteractionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addTransaction":
		obj := PaymentAddTransactionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAmountPlanned":
		obj := PaymentChangeAmountPlannedAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTransactionInteractionId":
		obj := PaymentChangeTransactionInteractionIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTransactionState":
		obj := PaymentChangeTransactionStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTransactionTimestamp":
		obj := PaymentChangeTransactionTimestampAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAmountPaid":
		obj := PaymentSetAmountPaidAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAmountRefunded":
		obj := PaymentSetAmountRefundedAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAnonymousId":
		obj := PaymentSetAnonymousIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAuthorization":
		obj := PaymentSetAuthorizationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := PaymentSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := PaymentSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomer":
		obj := PaymentSetCustomerAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setExternalId":
		obj := PaymentSetExternalIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setInterfaceId":
		obj := PaymentSetInterfaceIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := PaymentSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMethodInfoInterface":
		obj := PaymentSetMethodInfoInterfaceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMethodInfoMethod":
		obj := PaymentSetMethodInfoMethodAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMethodInfoName":
		obj := PaymentSetMethodInfoNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setStatusInterfaceCode":
		obj := PaymentSetStatusInterfaceCodeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setStatusInterfaceText":
		obj := PaymentSetStatusInterfaceTextAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTransactionCustomField":
		obj := PaymentSetTransactionCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTransactionCustomType":
		obj := PaymentSetTransactionCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "transitionState":
		obj := PaymentTransitionStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Represents a financial transaction typically created as a result of a notification from the payment service.
*
 */
type Transaction struct {
	// Unique identifier of the Transaction.
	ID string `json:"id"`
	// Date and time (UTC) the Transaction took place.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Type of the Transaction. For example, `Authorization`.
	Type TransactionType `json:"type"`
	// Money value of the Transaction.
	Amount CentPrecisionMoney `json:"amount"`
	// Identifier used by the interface that manages the Transaction (usually the PSP).
	// If a matching interaction was logged in the `interfaceInteractions` array, the corresponding interaction can be found with this ID.
	InteractionId *string `json:"interactionId,omitempty"`
	// State of the Transaction.
	State TransactionState `json:"state"`
	// Custom Fields defined for the Transaction.
	Custom *CustomFields `json:"custom,omitempty"`
}

type TransactionDraft struct {
	// Date and time (UTC) the Transaction took place.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Type of the Transaction.
	Type TransactionType `json:"type"`
	// Money value for the Transaction.
	Amount Money `json:"amount"`
	// Identifier used by the payment service that manages the Transaction.
	// Can be used to correlate the Transaction to an interface interaction.
	InteractionId *string `json:"interactionId,omitempty"`
	// State of the Transaction.
	State *TransactionState `json:"state,omitempty"`
	// Custom Fields of the Transaction.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

/**
*	Transactions can be in one of the following States:
*
 */
type TransactionState string

const (
	TransactionStateInitial TransactionState = "Initial"
	TransactionStatePending TransactionState = "Pending"
	TransactionStateSuccess TransactionState = "Success"
	TransactionStateFailure TransactionState = "Failure"
)

type TransactionType string

const (
	TransactionTypeAuthorization       TransactionType = "Authorization"
	TransactionTypeCancelAuthorization TransactionType = "CancelAuthorization"
	TransactionTypeCharge              TransactionType = "Charge"
	TransactionTypeRefund              TransactionType = "Refund"
	TransactionTypeChargeback          TransactionType = "Chargeback"
)

/**
*	Adding a Payment interaction generates the [PaymentInteractionAdded](ctp:api:type:PaymentInteractionAddedMessage) Message.
*
 */
type PaymentAddInterfaceInteractionAction struct {
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) of a [Type](ctp:api:type:Type).
	Type TypeResourceIdentifier `json:"type"`
	// [Custom Fields](/../api/projects/custom-fields) as per [FieldDefinitions](ctp:api:type:FieldDefinition) of the [Type](ctp:api:type:Type).
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentAddInterfaceInteractionAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentAddInterfaceInteractionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addInterfaceInteraction", Alias: (*Alias)(&obj)})
}

/**
*	Adding a Transaction to a Payment generates the [PaymentTransactionAdded](ctp:api:type:PaymentTransactionAddedMessage) Message.
*
 */
type PaymentAddTransactionAction struct {
	// Value to append to the `transactions` array.
	Transaction TransactionDraft `json:"transaction"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentAddTransactionAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentAddTransactionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTransaction", Alias: (*Alias)(&obj)})
}

/**
*	Can be used to update the Payment if a customer changes the [Cart](ctp:api:type:Cart), or adds or removes a [CartDiscount](ctp:api:type:CartDiscount) during checkout.
*
 */
type PaymentChangeAmountPlannedAction struct {
	// New value to set.
	Amount Money `json:"amount"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentChangeAmountPlannedAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeAmountPlannedAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAmountPlanned", Alias: (*Alias)(&obj)})
}

type PaymentChangeTransactionInteractionIdAction struct {
	// Unique identifier of the [Transaction](ctp:api:type:Transaction).
	TransactionId string `json:"transactionId"`
	// New value to set.
	InteractionId string `json:"interactionId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentChangeTransactionInteractionIdAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeTransactionInteractionIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTransactionInteractionId", Alias: (*Alias)(&obj)})
}

/**
*	Changing the [TransactionState](ctp:api:type:TransactionState) generates the [PaymentTransactionStateChanged](ctp:api:type:PaymentTransactionStateChangedMessage) Message.
*
 */
type PaymentChangeTransactionStateAction struct {
	// Unique identifier of the [Transaction](ctp:api:type:Transaction).
	TransactionId string `json:"transactionId"`
	// New TransactionState.
	State TransactionState `json:"state"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentChangeTransactionStateAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeTransactionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTransactionState", Alias: (*Alias)(&obj)})
}

type PaymentChangeTransactionTimestampAction struct {
	// Unique identifier of the [Transaction](ctp:api:type:Transaction).
	TransactionId string `json:"transactionId"`
	// Timestamp of the Transaction as reported by the payment service.
	Timestamp time.Time `json:"timestamp"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentChangeTransactionTimestampAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeTransactionTimestampAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTransactionTimestamp", Alias: (*Alias)(&obj)})
}

type PaymentSetAmountPaidAction struct {
	// Draft type that stores amounts only in cent precision for the specified currency.
	Amount *Money `json:"amount,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetAmountPaidAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAmountPaidAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAmountPaid", Alias: (*Alias)(&obj)})
}

type PaymentSetAmountRefundedAction struct {
	// Draft type that stores amounts only in cent precision for the specified currency.
	Amount *Money `json:"amount,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetAmountRefundedAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAmountRefundedAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAmountRefunded", Alias: (*Alias)(&obj)})
}

type PaymentSetAnonymousIdAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	AnonymousId *string `json:"anonymousId,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetAnonymousIdAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAnonymousIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAnonymousId", Alias: (*Alias)(&obj)})
}

type PaymentSetAuthorizationAction struct {
	// Draft type that stores amounts only in cent precision for the specified currency.
	Amount *Money     `json:"amount,omitempty"`
	Until  *time.Time `json:"until,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetAuthorizationAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAuthorizationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAuthorization", Alias: (*Alias)(&obj)})
}

type PaymentSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type PaymentSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the Payment with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Payment.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Payment.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type PaymentSetCustomerAction struct {
	// Value to set.
	// If empty, any existing reference is removed.
	Customer *CustomerResourceIdentifier `json:"customer,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomer", Alias: (*Alias)(&obj)})
}

type PaymentSetExternalIdAction struct {
	ExternalId *string `json:"externalId,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetExternalIdAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetExternalIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalId", Alias: (*Alias)(&obj)})
}

type PaymentSetInterfaceIdAction struct {
	// Value to set.
	// Once set, the `interfaceId` cannot be changed.
	InterfaceId string `json:"interfaceId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetInterfaceIdAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetInterfaceIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setInterfaceId", Alias: (*Alias)(&obj)})
}

type PaymentSetKeyAction struct {
	// Value to set.
	// If `key` is absent or `null`, the existing key, if any, will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type PaymentSetMethodInfoInterfaceAction struct {
	// Value to set.
	// Once set, the `paymentInterface` of the `paymentMethodInfo` cannot be changed.
	Interface string `json:"interface"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetMethodInfoInterfaceAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoInterfaceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoInterface", Alias: (*Alias)(&obj)})
}

type PaymentSetMethodInfoMethodAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	Method *string `json:"method,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetMethodInfoMethodAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoMethod", Alias: (*Alias)(&obj)})
}

type PaymentSetMethodInfoNameAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetMethodInfoNameAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoName", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [PaymentStatusInterfaceCodeSet](ctp:api:type:PaymentStatusInterfaceCodeSetMessage) Message.
 */
type PaymentSetStatusInterfaceCodeAction struct {
	// Value to set. If empty, any existing value will be removed.
	InterfaceCode *string `json:"interfaceCode,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetStatusInterfaceCodeAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetStatusInterfaceCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStatusInterfaceCode", Alias: (*Alias)(&obj)})
}

type PaymentSetStatusInterfaceTextAction struct {
	// Value to set. If empty, any existing value will be removed.
	InterfaceText string `json:"interfaceText"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetStatusInterfaceTextAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetStatusInterfaceTextAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStatusInterfaceText", Alias: (*Alias)(&obj)})
}

type PaymentSetTransactionCustomFieldAction struct {
	// Unique identifier of the [Transaction](ctp:api:type:Transaction).
	TransactionId string `json:"transactionId"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetTransactionCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetTransactionCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTransactionCustomField", Alias: (*Alias)(&obj)})
}

type PaymentSetTransactionCustomTypeAction struct {
	// Unique identifier of the [Transaction](ctp:api:type:Transaction). If the specified `transactionId` does not exist, the request will fail with an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	TransactionId string `json:"transactionId"`
	// Defines the [Type](ctp:api:type:Type) that extends the Transaction with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Transaction.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Transaction.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentSetTransactionCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetTransactionCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTransactionCustomType", Alias: (*Alias)(&obj)})
}

/**
*	If the Payment has no current [State](ctp:api:type:State), `initial` must be `true` for the new State.
*	If the existing State has transitions set, the new State must be a valid transition.
*	If the existing State has no transitions set, no validations are performed when transitioning to the new State.
*
*	Transitioning the State of a Payment produces the [PaymentStatusStateTransition](ctp:api:type:PaymentStatusStateTransitionMessage) Message.
*
 */
type PaymentTransitionStateAction struct {
	// [ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [State](ctp:api:type:State).
	State StateResourceIdentifier `json:"state"`
	// Set to `true` to skip validations when transitioning to the new State.
	Force *bool `json:"force,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}
