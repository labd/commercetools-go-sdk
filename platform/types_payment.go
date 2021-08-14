// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type Payment struct {
	Id             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for events not tracked.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	// Present on resources created after 1/02/2019 except for events not tracked.
	CreatedBy *CreatedBy `json:"createdBy,omitEmpty"`
	// A reference to the customer this payment belongs to.
	Customer CustomerReference `json:"customer,omitEmpty"`
	// Identifies payments belonging to an anonymous session (the customer has not signed up/in yet).
	AnonymousId string `json:"anonymousId,omitEmpty"`
	ExternalId  string `json:"externalId,omitEmpty"`
	// The identifier that is used by the interface that manages the payment (usually the PSP).
	// Cannot be changed once it has been set.
	// The combination of this ID and the PaymentMethodInfo `paymentInterface` must be unique.
	InterfaceId string `json:"interfaceId,omitEmpty"`
	// How much money this payment intends to receive from the customer.
	// The value usually matches the cart or order gross total.
	AmountPlanned     TypedMoney        `json:"amountPlanned"`
	AmountAuthorized  TypedMoney        `json:"amountAuthorized,omitEmpty"`
	AuthorizedUntil   string            `json:"authorizedUntil,omitEmpty"`
	AmountPaid        TypedMoney        `json:"amountPaid,omitEmpty"`
	AmountRefunded    TypedMoney        `json:"amountRefunded,omitEmpty"`
	PaymentMethodInfo PaymentMethodInfo `json:"paymentMethodInfo"`
	PaymentStatus     PaymentStatus     `json:"paymentStatus"`
	// A list of financial transactions of different TransactionTypes with different TransactionStates.
	Transactions []Transaction `json:"transactions"`
	// Interface interactions can be requests sent to the PSP, responses received from the PSP or notifications received from the PSP.
	// Some interactions may result in a transaction.
	// If so, the `interactionId` in the Transaction should be set to match the ID of the PSP for the interaction.
	// Interactions are managed by the PSP integration and are usually neither written nor read by the user facing frontends or other services.
	InterfaceInteractions []CustomFields `json:"interfaceInteractions"`
	Custom                *CustomFields  `json:"custom,omitEmpty"`
	// User-specific unique identifier for the payment (max.
	// 256 characters).
	Key string `json:"key,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Payment) UnmarshalJSON(data []byte) error {
	type Alias Payment
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.AmountPlanned != nil {
		var err error
		obj.AmountPlanned, err = mapDiscriminatorTypedMoney(obj.AmountPlanned)
		if err != nil {
			return err
		}
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
	// A reference to the customer this payment belongs to.
	Customer CustomerResourceIdentifier `json:"customer,omitEmpty"`
	// Identifies payments belonging to an anonymous session (the customer has not signed up/in yet).
	AnonymousId string `json:"anonymousId,omitEmpty"`
	ExternalId  string `json:"externalId,omitEmpty"`
	// The identifier that is used by the interface that manages the payment (usually the PSP).
	// Cannot be changed once it has been set.
	// The combination of this ID and the PaymentMethodInfo `paymentInterface` must be unique.
	InterfaceId string `json:"interfaceId,omitEmpty"`
	// How much money this payment intends to receive from the customer.
	// The value usually matches the cart or order gross total.
	AmountPlanned     Money               `json:"amountPlanned"`
	AmountAuthorized  *Money              `json:"amountAuthorized,omitEmpty"`
	AuthorizedUntil   string              `json:"authorizedUntil,omitEmpty"`
	AmountPaid        *Money              `json:"amountPaid,omitEmpty"`
	AmountRefunded    *Money              `json:"amountRefunded,omitEmpty"`
	PaymentMethodInfo *PaymentMethodInfo  `json:"paymentMethodInfo,omitEmpty"`
	PaymentStatus     *PaymentStatusDraft `json:"paymentStatus,omitEmpty"`
	// A list of financial transactions of different TransactionTypes with different TransactionStates.
	Transactions []TransactionDraft `json:"transactions,omitEmpty"`
	// Interface interactions can be requests send to the PSP, responses received from the PSP or notifications received from the PSP.
	// Some interactions may result in a transaction.
	// If so, the `interactionId` in the Transaction should be set to match the ID of the PSP for the interaction.
	// Interactions are managed by the PSP integration and are usually neither written nor read by the user facing frontends or other services.
	InterfaceInteractions []CustomFieldsDraft `json:"interfaceInteractions,omitEmpty"`
	Custom                *CustomFieldsDraft  `json:"custom,omitEmpty"`
	// User-specific unique identifier for the payment (max.
	// 256 characters).
	Key string `json:"key,omitEmpty"`
}

type PaymentMethodInfo struct {
	// The interface that handles the payment (usually a PSP).
	// Cannot be changed once it has been set.
	// The combination of Payment`interfaceId` and this field must be unique.
	PaymentInterface string `json:"paymentInterface,omitEmpty"`
	// The payment method that is used, e.g.
	// e.g.
	// a conventional string representing Credit Card, Cash Advance etc.
	Method string `json:"method,omitEmpty"`
	// A human-readable, localized name for the payment method, e.g.
	// 'Credit Card'.
	Name *LocalizedString `json:"name,omitEmpty"`
}

type PaymentPagedQueryResponse struct {
	Limit   int       `json:"limit"`
	Count   int       `json:"count"`
	Total   int       `json:"total,omitEmpty"`
	Offset  int       `json:"offset"`
	Results []Payment `json:"results"`
}

type PaymentReference struct {
	Id  string   `json:"id"`
	Obj *Payment `json:"obj,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentReference) MarshalJSON() ([]byte, error) {
	type Alias PaymentReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "payment", Alias: (*Alias)(&obj)})
}

type PaymentResourceIdentifier struct {
	Id  string `json:"id,omitEmpty"`
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias PaymentResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "payment", Alias: (*Alias)(&obj)})
}

type PaymentStatus struct {
	// A code describing the current status returned by the interface that processes the payment.
	InterfaceCode string `json:"interfaceCode,omitEmpty"`
	// A text describing the current status returned by the interface that processes the payment.
	InterfaceText string         `json:"interfaceText,omitEmpty"`
	State         StateReference `json:"state,omitEmpty"`
}

type PaymentStatusDraft struct {
	InterfaceCode string                  `json:"interfaceCode,omitEmpty"`
	InterfaceText string                  `json:"interfaceText,omitEmpty"`
	State         StateResourceIdentifier `json:"state,omitEmpty"`
}

type PaymentUpdate struct {
	Version int                   `json:"version"`
	Actions []PaymentUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PaymentUpdate) UnmarshalJSON(data []byte) error {
	type Alias PaymentUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type PaymentUpdateAction interface{}

func mapDiscriminatorPaymentUpdateAction(input interface{}) (PaymentUpdateAction, error) {

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
	case "addInterfaceInteraction":
		new := PaymentAddInterfaceInteractionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "addTransaction":
		new := PaymentAddTransactionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeAmountPlanned":
		new := PaymentChangeAmountPlannedAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTransactionInteractionId":
		new := PaymentChangeTransactionInteractionIdAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTransactionState":
		new := PaymentChangeTransactionStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTransactionTimestamp":
		new := PaymentChangeTransactionTimestampAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAmountPaid":
		new := PaymentSetAmountPaidAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAmountRefunded":
		new := PaymentSetAmountRefundedAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAnonymousId":
		new := PaymentSetAnonymousIdAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAuthorization":
		new := PaymentSetAuthorizationAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := PaymentSetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := PaymentSetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomer":
		new := PaymentSetCustomerAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setExternalId":
		new := PaymentSetExternalIdAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setInterfaceId":
		new := PaymentSetInterfaceIdAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := PaymentSetKeyAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMethodInfoInterface":
		new := PaymentSetMethodInfoInterfaceAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMethodInfoMethod":
		new := PaymentSetMethodInfoMethodAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMethodInfoName":
		new := PaymentSetMethodInfoNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setStatusInterfaceCode":
		new := PaymentSetStatusInterfaceCodeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setStatusInterfaceText":
		new := PaymentSetStatusInterfaceTextAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "transitionState":
		new := PaymentTransitionStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type Transaction struct {
	// The unique ID of this object.
	Id string `json:"id"`
	// The time at which the transaction took place.
	Timestamp time.Time `json:"timestamp,omitEmpty"`
	// The type of this transaction.
	Type   TransactionType `json:"type"`
	Amount TypedMoney      `json:"amount"`
	// The identifier that is used by the interface that managed the transaction (usually the PSP).
	// If a matching interaction was logged in the `interfaceInteractions` array, the corresponding interaction should be findable with this ID.
	InteractionId string `json:"interactionId,omitEmpty"`
	// The state of this transaction.
	State TransactionState `json:"state,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Transaction) UnmarshalJSON(data []byte) error {
	type Alias Transaction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Amount != nil {
		var err error
		obj.Amount, err = mapDiscriminatorTypedMoney(obj.Amount)
		if err != nil {
			return err
		}
	}
	return nil
}

type TransactionDraft struct {
	// The time at which the transaction took place.
	Timestamp time.Time `json:"timestamp,omitEmpty"`
	// The type of this transaction.
	Type   TransactionType `json:"type"`
	Amount Money           `json:"amount"`
	// The identifier that is used by the interface that managed the transaction (usually the PSP).
	// If a matching interaction was logged in the `interfaceInteractions` array, the corresponding interaction should be findable with this ID.
	InteractionId string `json:"interactionId,omitEmpty"`
	// The state of this transaction.
	// If not set, defaults to `Initial`.
	State TransactionState `json:"state,omitEmpty"`
}

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

type PaymentAddInterfaceInteractionAction struct {
	Type   TypeResourceIdentifier `json:"type"`
	Fields *FieldContainer        `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentAddInterfaceInteractionAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentAddInterfaceInteractionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addInterfaceInteraction", Alias: (*Alias)(&obj)})
}

type PaymentAddTransactionAction struct {
	Transaction TransactionDraft `json:"transaction"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentAddTransactionAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentAddTransactionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTransaction", Alias: (*Alias)(&obj)})
}

type PaymentChangeAmountPlannedAction struct {
	Amount Money `json:"amount"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentChangeAmountPlannedAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeAmountPlannedAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAmountPlanned", Alias: (*Alias)(&obj)})
}

type PaymentChangeTransactionInteractionIdAction struct {
	TransactionId string `json:"transactionId"`
	InteractionId string `json:"interactionId"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentChangeTransactionInteractionIdAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeTransactionInteractionIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTransactionInteractionId", Alias: (*Alias)(&obj)})
}

type PaymentChangeTransactionStateAction struct {
	TransactionId string           `json:"transactionId"`
	State         TransactionState `json:"state"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentChangeTransactionStateAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeTransactionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTransactionState", Alias: (*Alias)(&obj)})
}

type PaymentChangeTransactionTimestampAction struct {
	TransactionId string    `json:"transactionId"`
	Timestamp     time.Time `json:"timestamp"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentChangeTransactionTimestampAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeTransactionTimestampAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTransactionTimestamp", Alias: (*Alias)(&obj)})
}

type PaymentSetAmountPaidAction struct {
	Amount *Money `json:"amount,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetAmountPaidAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAmountPaidAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAmountPaid", Alias: (*Alias)(&obj)})
}

type PaymentSetAmountRefundedAction struct {
	Amount *Money `json:"amount,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetAmountRefundedAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAmountRefundedAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAmountRefunded", Alias: (*Alias)(&obj)})
}

type PaymentSetAnonymousIdAction struct {
	// Anonymous ID of the anonymous customer that this payment belongs to.
	// If this field is not set any existing `anonymousId` is removed.
	AnonymousId string `json:"anonymousId,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetAnonymousIdAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAnonymousIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAnonymousId", Alias: (*Alias)(&obj)})
}

type PaymentSetAuthorizationAction struct {
	Amount *Money    `json:"amount,omitEmpty"`
	Until  time.Time `json:"until,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetAuthorizationAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAuthorizationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAuthorization", Alias: (*Alias)(&obj)})
}

type PaymentSetCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type PaymentSetCustomTypeAction struct {
	// If set, the custom type is set to this new value.
	// If absent, the custom type and any existing custom fields are removed.
	Type TypeResourceIdentifier `json:"type,omitEmpty"`
	// Sets the custom fields to this value.
	Fields *FieldContainer `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type PaymentSetCustomerAction struct {
	// A reference to the customer this payment belongs to.
	Customer CustomerResourceIdentifier `json:"customer,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomer", Alias: (*Alias)(&obj)})
}

type PaymentSetExternalIdAction struct {
	ExternalId string `json:"externalId,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetExternalIdAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetExternalIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalId", Alias: (*Alias)(&obj)})
}

type PaymentSetInterfaceIdAction struct {
	InterfaceId string `json:"interfaceId"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetInterfaceIdAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetInterfaceIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setInterfaceId", Alias: (*Alias)(&obj)})
}

type PaymentSetKeyAction struct {
	// User-specific unique identifier for the payment (max.
	// 256 characters).
	// If not provided an existing key will be removed.
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type PaymentSetMethodInfoInterfaceAction struct {
	Interface string `json:"interface"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetMethodInfoInterfaceAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoInterfaceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoInterface", Alias: (*Alias)(&obj)})
}

type PaymentSetMethodInfoMethodAction struct {
	// If not provided, the method is unset.
	Method string `json:"method,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetMethodInfoMethodAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoMethod", Alias: (*Alias)(&obj)})
}

type PaymentSetMethodInfoNameAction struct {
	// If not provided, the name is unset.
	Name *LocalizedString `json:"name,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetMethodInfoNameAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoName", Alias: (*Alias)(&obj)})
}

type PaymentSetStatusInterfaceCodeAction struct {
	InterfaceCode string `json:"interfaceCode,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetStatusInterfaceCodeAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetStatusInterfaceCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStatusInterfaceCode", Alias: (*Alias)(&obj)})
}

type PaymentSetStatusInterfaceTextAction struct {
	InterfaceText string `json:"interfaceText"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetStatusInterfaceTextAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetStatusInterfaceTextAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStatusInterfaceText", Alias: (*Alias)(&obj)})
}

type PaymentTransitionStateAction struct {
	State StateResourceIdentifier `json:"state"`
	Force bool                    `json:"force,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}
