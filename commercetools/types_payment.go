// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// TransactionState is an enum type
type TransactionState string

// Enum values for TransactionState
const (
	TransactionStateInitial TransactionState = "Initial"
	TransactionStatePending TransactionState = "Pending"
	TransactionStateSuccess TransactionState = "Success"
	TransactionStateFailure TransactionState = "Failure"
)

// TransactionType is an enum type
type TransactionType string

// Enum values for TransactionType
const (
	TransactionTypeAuthorization       TransactionType = "Authorization"
	TransactionTypeCancelAuthorization TransactionType = "CancelAuthorization"
	TransactionTypeCharge              TransactionType = "Charge"
	TransactionTypeRefund              TransactionType = "Refund"
	TransactionTypeChargeback          TransactionType = "Chargeback"
)

// PaymentUpdateAction uses action as discriminator attribute
type PaymentUpdateAction interface{}

func mapDiscriminatorPaymentUpdateAction(input PaymentUpdateAction) PaymentUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addInterfaceInteraction":
		new := PaymentAddInterfaceInteractionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addTransaction":
		new := PaymentAddTransactionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeAmountPlanned":
		new := PaymentChangeAmountPlannedAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeTransactionInteractionId":
		new := PaymentChangeTransactionInteractionIDAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeTransactionState":
		new := PaymentChangeTransactionStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeTransactionTimestamp":
		new := PaymentChangeTransactionTimestampAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAmountPaid":
		new := PaymentSetAmountPaidAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAmountRefunded":
		new := PaymentSetAmountRefundedAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAnonymousId":
		new := PaymentSetAnonymousIDAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAuthorization":
		new := PaymentSetAuthorizationAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomField":
		new := PaymentSetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomType":
		new := PaymentSetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomer":
		new := PaymentSetCustomerAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setExternalId":
		new := PaymentSetExternalIDAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setInterfaceId":
		new := PaymentSetInterfaceIDAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := PaymentSetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setMethodInfoInterface":
		new := PaymentSetMethodInfoInterfaceAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setMethodInfoMethod":
		new := PaymentSetMethodInfoMethodAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setMethodInfoName":
		new := PaymentSetMethodInfoNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setStatusInterfaceCode":
		new := PaymentSetStatusInterfaceCodeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setStatusInterfaceText":
		new := PaymentSetStatusInterfaceTextAction{}
		mapstructure.Decode(input, &new)
		return new
	case "transitionState":
		new := PaymentTransitionStateAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// Payment is of type Resource
type Payment struct {
	Version               int                `json:"version"`
	LastModifiedAt        time.Time          `json:"lastModifiedAt"`
	ID                    string             `json:"id"`
	CreatedAt             time.Time          `json:"createdAt"`
	Transactions          []Transaction      `json:"transactions"`
	PaymentStatus         *PaymentStatus     `json:"paymentStatus"`
	PaymentMethodInfo     *PaymentMethodInfo `json:"paymentMethodInfo"`
	Key                   string             `json:"key,omitempty"`
	InterfaceInteractions []CustomFields     `json:"interfaceInteractions"`
	InterfaceID           string             `json:"interfaceId,omitempty"`
	ExternalID            string             `json:"externalId,omitempty"`
	Customer              *CustomerReference `json:"customer,omitempty"`
	Custom                *CustomFields      `json:"custom,omitempty"`
	AuthorizedUntil       string             `json:"authorizedUntil,omitempty"`
	AnonymousID           string             `json:"anonymousId,omitempty"`
	AmountRefunded        TypedMoney         `json:"amountRefunded,omitempty"`
	AmountPlanned         TypedMoney         `json:"amountPlanned"`
	AmountPaid            TypedMoney         `json:"amountPaid,omitempty"`
	AmountAuthorized      TypedMoney         `json:"amountAuthorized,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Payment) UnmarshalJSON(data []byte) error {
	type Alias Payment
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.AmountAuthorized != nil {
		obj.AmountAuthorized = mapDiscriminatorTypedMoney(obj.AmountAuthorized)
	}
	if obj.AmountPaid != nil {
		obj.AmountPaid = mapDiscriminatorTypedMoney(obj.AmountPaid)
	}
	if obj.AmountPlanned != nil {
		obj.AmountPlanned = mapDiscriminatorTypedMoney(obj.AmountPlanned)
	}
	if obj.AmountRefunded != nil {
		obj.AmountRefunded = mapDiscriminatorTypedMoney(obj.AmountRefunded)
	}

	return nil
}

// PaymentAddInterfaceInteractionAction implements the interface PaymentUpdateAction
type PaymentAddInterfaceInteractionAction struct {
	Type   *TypeReference  `json:"type"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentAddInterfaceInteractionAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentAddInterfaceInteractionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addInterfaceInteraction", Alias: (*Alias)(&obj)})
}

// PaymentAddTransactionAction implements the interface PaymentUpdateAction
type PaymentAddTransactionAction struct {
	Transaction *TransactionDraft `json:"transaction"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentAddTransactionAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentAddTransactionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTransaction", Alias: (*Alias)(&obj)})
}

// PaymentChangeAmountPlannedAction implements the interface PaymentUpdateAction
type PaymentChangeAmountPlannedAction struct {
	Amount *Money `json:"amount"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentChangeAmountPlannedAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeAmountPlannedAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAmountPlanned", Alias: (*Alias)(&obj)})
}

// PaymentChangeTransactionInteractionIDAction implements the interface PaymentUpdateAction
type PaymentChangeTransactionInteractionIDAction struct {
	TransactionID string `json:"transactionId"`
	InteractionID string `json:"interactionId"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentChangeTransactionInteractionIDAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeTransactionInteractionIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTransactionInteractionId", Alias: (*Alias)(&obj)})
}

// PaymentChangeTransactionStateAction implements the interface PaymentUpdateAction
type PaymentChangeTransactionStateAction struct {
	TransactionID string           `json:"transactionId"`
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

// PaymentChangeTransactionTimestampAction implements the interface PaymentUpdateAction
type PaymentChangeTransactionTimestampAction struct {
	TransactionID string    `json:"transactionId"`
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

// PaymentDraft is a standalone struct
type PaymentDraft struct {
	Transactions          []TransactionDraft  `json:"transactions,omitempty"`
	PaymentStatus         *PaymentStatus      `json:"paymentStatus,omitempty"`
	PaymentMethodInfo     *PaymentMethodInfo  `json:"paymentMethodInfo,omitempty"`
	Key                   string              `json:"key,omitempty"`
	InterfaceInteractions []CustomFieldsDraft `json:"interfaceInteractions,omitempty"`
	InterfaceID           string              `json:"interfaceId,omitempty"`
	ExternalID            string              `json:"externalId,omitempty"`
	Customer              *CustomerReference  `json:"customer,omitempty"`
	Custom                *CustomFieldsDraft  `json:"custom,omitempty"`
	AuthorizedUntil       string              `json:"authorizedUntil,omitempty"`
	AnonymousID           string              `json:"anonymousId,omitempty"`
	AmountRefunded        *Money              `json:"amountRefunded,omitempty"`
	AmountPlanned         *Money              `json:"amountPlanned"`
	AmountPaid            *Money              `json:"amountPaid,omitempty"`
	AmountAuthorized      *Money              `json:"amountAuthorized,omitempty"`
}

// PaymentMethodInfo is a standalone struct
type PaymentMethodInfo struct {
	PaymentInterface string           `json:"paymentInterface,omitempty"`
	Name             *LocalizedString `json:"name,omitempty"`
	Method           string           `json:"method,omitempty"`
}

// PaymentPagedQueryResponse is of type PagedQueryResponse
type PaymentPagedQueryResponse struct {
	Total   int       `json:"total,omitempty"`
	Offset  int       `json:"offset"`
	Count   int       `json:"count"`
	Results []Payment `json:"results"`
}

// PaymentReference implements the interface Reference
type PaymentReference struct {
	Key string   `json:"key,omitempty"`
	ID  string   `json:"id,omitempty"`
	Obj *Payment `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentReference) MarshalJSON() ([]byte, error) {
	type Alias PaymentReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "payment", Alias: (*Alias)(&obj)})
}

// PaymentSetAmountPaidAction implements the interface PaymentUpdateAction
type PaymentSetAmountPaidAction struct {
	Amount *Money `json:"amount,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetAmountPaidAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAmountPaidAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAmountPaid", Alias: (*Alias)(&obj)})
}

// PaymentSetAmountRefundedAction implements the interface PaymentUpdateAction
type PaymentSetAmountRefundedAction struct {
	Amount *Money `json:"amount,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetAmountRefundedAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAmountRefundedAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAmountRefunded", Alias: (*Alias)(&obj)})
}

// PaymentSetAnonymousIDAction implements the interface PaymentUpdateAction
type PaymentSetAnonymousIDAction struct {
	AnonymousID string `json:"anonymousId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetAnonymousIDAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAnonymousIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAnonymousId", Alias: (*Alias)(&obj)})
}

// PaymentSetAuthorizationAction implements the interface PaymentUpdateAction
type PaymentSetAuthorizationAction struct {
	Until  time.Time `json:"until,omitempty"`
	Amount *Money    `json:"amount,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetAuthorizationAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAuthorizationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAuthorization", Alias: (*Alias)(&obj)})
}

// PaymentSetCustomFieldAction implements the interface PaymentUpdateAction
type PaymentSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// PaymentSetCustomTypeAction implements the interface PaymentUpdateAction
type PaymentSetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// PaymentSetCustomerAction implements the interface PaymentUpdateAction
type PaymentSetCustomerAction struct {
	Customer *CustomerReference `json:"customer,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomer", Alias: (*Alias)(&obj)})
}

// PaymentSetExternalIDAction implements the interface PaymentUpdateAction
type PaymentSetExternalIDAction struct {
	ExternalID string `json:"externalId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetExternalIDAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetExternalIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalId", Alias: (*Alias)(&obj)})
}

// PaymentSetInterfaceIDAction implements the interface PaymentUpdateAction
type PaymentSetInterfaceIDAction struct {
	InterfaceID string `json:"interfaceId"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetInterfaceIDAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetInterfaceIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setInterfaceId", Alias: (*Alias)(&obj)})
}

// PaymentSetKeyAction implements the interface PaymentUpdateAction
type PaymentSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// PaymentSetMethodInfoInterfaceAction implements the interface PaymentUpdateAction
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

// PaymentSetMethodInfoMethodAction implements the interface PaymentUpdateAction
type PaymentSetMethodInfoMethodAction struct {
	Method string `json:"method,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetMethodInfoMethodAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoMethod", Alias: (*Alias)(&obj)})
}

// PaymentSetMethodInfoNameAction implements the interface PaymentUpdateAction
type PaymentSetMethodInfoNameAction struct {
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetMethodInfoNameAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoName", Alias: (*Alias)(&obj)})
}

// PaymentSetStatusInterfaceCodeAction implements the interface PaymentUpdateAction
type PaymentSetStatusInterfaceCodeAction struct {
	InterfaceCode string `json:"interfaceCode,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentSetStatusInterfaceCodeAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetStatusInterfaceCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStatusInterfaceCode", Alias: (*Alias)(&obj)})
}

// PaymentSetStatusInterfaceTextAction implements the interface PaymentUpdateAction
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

// PaymentStatus is a standalone struct
type PaymentStatus struct {
	State         *StateReference `json:"state,omitempty"`
	InterfaceText string          `json:"interfaceText,omitempty"`
	InterfaceCode string          `json:"interfaceCode,omitempty"`
}

// PaymentTransitionStateAction implements the interface PaymentUpdateAction
type PaymentTransitionStateAction struct {
	State *StateReference `json:"state"`
	Force bool            `json:"force,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj PaymentTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}

// PaymentUpdate is of type Update
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
	for i := range obj.Actions {
		obj.Actions[i] = mapDiscriminatorPaymentUpdateAction(obj.Actions[i])
	}

	return nil
}

// Transaction is a standalone struct
type Transaction struct {
	Type          TransactionType  `json:"type"`
	Timestamp     time.Time        `json:"timestamp,omitempty"`
	State         TransactionState `json:"state,omitempty"`
	InteractionID string           `json:"interactionId,omitempty"`
	ID            string           `json:"id"`
	Amount        TypedMoney       `json:"amount"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Transaction) UnmarshalJSON(data []byte) error {
	type Alias Transaction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Amount != nil {
		obj.Amount = mapDiscriminatorTypedMoney(obj.Amount)
	}

	return nil
}

// TransactionDraft is a standalone struct
type TransactionDraft struct {
	Type          TransactionType  `json:"type"`
	Timestamp     time.Time        `json:"timestamp,omitempty"`
	State         TransactionState `json:"state,omitempty"`
	InteractionID string           `json:"interactionId,omitempty"`
	Amount        *Money           `json:"amount"`
}
