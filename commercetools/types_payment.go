// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

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

func (obj *Payment) UnmarshalJSON(data []byte) error {
	type Alias Payment
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.AmountAuthorized != nil {
		obj.AmountAuthorized = AbstractTypedMoneyDiscriminatorMapping(obj.AmountAuthorized)
	}
	if obj.AmountPaid != nil {
		obj.AmountPaid = AbstractTypedMoneyDiscriminatorMapping(obj.AmountPaid)
	}
	if obj.AmountPlanned != nil {
		obj.AmountPlanned = AbstractTypedMoneyDiscriminatorMapping(obj.AmountPlanned)
	}
	if obj.AmountRefunded != nil {
		obj.AmountRefunded = AbstractTypedMoneyDiscriminatorMapping(obj.AmountRefunded)
	}

	return nil
}

type PaymentAddInterfaceInteractionAction struct {
	Type   *TypeReference  `json:"type"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

func (obj PaymentAddInterfaceInteractionAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentAddInterfaceInteractionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addInterfaceInteraction", Alias: (*Alias)(&obj)})
}

type PaymentAddTransactionAction struct {
	Transaction *TransactionDraft `json:"transaction"`
}

func (obj PaymentAddTransactionAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentAddTransactionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTransaction", Alias: (*Alias)(&obj)})
}

type PaymentChangeAmountPlannedAction struct {
	Amount *Money `json:"amount"`
}

func (obj PaymentChangeAmountPlannedAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeAmountPlannedAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAmountPlanned", Alias: (*Alias)(&obj)})
}

type PaymentChangeTransactionInteractionIdAction struct {
	TransactionID string `json:"transactionId"`
	InteractionID string `json:"interactionId"`
}

func (obj PaymentChangeTransactionInteractionIdAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeTransactionInteractionIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTransactionInteractionId", Alias: (*Alias)(&obj)})
}

type PaymentChangeTransactionStateAction struct {
	TransactionID string           `json:"transactionId"`
	State         TransactionState `json:"state"`
}

func (obj PaymentChangeTransactionStateAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeTransactionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTransactionState", Alias: (*Alias)(&obj)})
}

type PaymentChangeTransactionTimestampAction struct {
	TransactionID string    `json:"transactionId"`
	Timestamp     time.Time `json:"timestamp"`
}

func (obj PaymentChangeTransactionTimestampAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentChangeTransactionTimestampAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTransactionTimestamp", Alias: (*Alias)(&obj)})
}

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

type PaymentMethodInfo struct {
	PaymentInterface string           `json:"paymentInterface,omitempty"`
	Name             *LocalizedString `json:"name,omitempty"`
	Method           string           `json:"method,omitempty"`
}

type PaymentPagedQueryResponse struct {
	Total   int       `json:"total,omitempty"`
	Offset  int       `json:"offset"`
	Count   int       `json:"count"`
	Results []Payment `json:"results"`
}

type PaymentReference struct {
	Key string   `json:"key,omitempty"`
	ID  string   `json:"id,omitempty"`
	Obj *Payment `json:"obj,omitempty"`
}

func (obj PaymentReference) MarshalJSON() ([]byte, error) {
	type Alias PaymentReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "payment", Alias: (*Alias)(&obj)})
}

type PaymentSetAmountPaidAction struct {
	Amount *Money `json:"amount,omitempty"`
}

func (obj PaymentSetAmountPaidAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAmountPaidAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAmountPaid", Alias: (*Alias)(&obj)})
}

type PaymentSetAmountRefundedAction struct {
	Amount *Money `json:"amount,omitempty"`
}

func (obj PaymentSetAmountRefundedAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAmountRefundedAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAmountRefunded", Alias: (*Alias)(&obj)})
}

type PaymentSetAnonymousIdAction struct {
	AnonymousID string `json:"anonymousId,omitempty"`
}

func (obj PaymentSetAnonymousIdAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAnonymousIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAnonymousId", Alias: (*Alias)(&obj)})
}

type PaymentSetAuthorizationAction struct {
	Until  time.Time `json:"until,omitempty"`
	Amount *Money    `json:"amount,omitempty"`
}

func (obj PaymentSetAuthorizationAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetAuthorizationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAuthorization", Alias: (*Alias)(&obj)})
}

type PaymentSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

func (obj PaymentSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type PaymentSetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

func (obj PaymentSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type PaymentSetCustomerAction struct {
	Customer *CustomerReference `json:"customer,omitempty"`
}

func (obj PaymentSetCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomer", Alias: (*Alias)(&obj)})
}

type PaymentSetExternalIdAction struct {
	ExternalID string `json:"externalId,omitempty"`
}

func (obj PaymentSetExternalIdAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetExternalIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalId", Alias: (*Alias)(&obj)})
}

type PaymentSetInterfaceIdAction struct {
	InterfaceID string `json:"interfaceId"`
}

func (obj PaymentSetInterfaceIdAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetInterfaceIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setInterfaceId", Alias: (*Alias)(&obj)})
}

type PaymentSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

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

func (obj PaymentSetMethodInfoInterfaceAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoInterfaceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoInterface", Alias: (*Alias)(&obj)})
}

type PaymentSetMethodInfoMethodAction struct {
	Method string `json:"method,omitempty"`
}

func (obj PaymentSetMethodInfoMethodAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoMethodAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoMethod", Alias: (*Alias)(&obj)})
}

type PaymentSetMethodInfoNameAction struct {
	Name *LocalizedString `json:"name,omitempty"`
}

func (obj PaymentSetMethodInfoNameAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetMethodInfoNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMethodInfoName", Alias: (*Alias)(&obj)})
}

type PaymentSetStatusInterfaceCodeAction struct {
	InterfaceCode string `json:"interfaceCode,omitempty"`
}

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

func (obj PaymentSetStatusInterfaceTextAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentSetStatusInterfaceTextAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStatusInterfaceText", Alias: (*Alias)(&obj)})
}

type PaymentStatus struct {
	State         *StateReference `json:"state,omitempty"`
	InterfaceText string          `json:"interfaceText,omitempty"`
	InterfaceCode string          `json:"interfaceCode,omitempty"`
}

type PaymentTransitionStateAction struct {
	State *StateReference `json:"state"`
	Force bool            `json:"force,omitempty"`
}

func (obj PaymentTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}

type PaymentUpdate struct {
	Version int                   `json:"version"`
	Actions []PaymentUpdateAction `json:"actions"`
}

func (obj *PaymentUpdate) UnmarshalJSON(data []byte) error {
	type Alias PaymentUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractPaymentUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type PaymentUpdateAction interface{}
type AbstractPaymentUpdateAction struct{}

func AbstractPaymentUpdateActionDiscriminatorMapping(input PaymentUpdateAction) PaymentUpdateAction {
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
		new := PaymentChangeTransactionInteractionIdAction{}
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
		new := PaymentSetAnonymousIdAction{}
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
		new := PaymentSetExternalIdAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setInterfaceId":
		new := PaymentSetInterfaceIdAction{}
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

type Transaction struct {
	Type          TransactionType  `json:"type"`
	Timestamp     time.Time        `json:"timestamp,omitempty"`
	State         TransactionState `json:"state,omitempty"`
	InteractionID string           `json:"interactionId,omitempty"`
	ID            string           `json:"id"`
	Amount        TypedMoney       `json:"amount"`
}

func (obj *Transaction) UnmarshalJSON(data []byte) error {
	type Alias Transaction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Amount != nil {
		obj.Amount = AbstractTypedMoneyDiscriminatorMapping(obj.Amount)
	}

	return nil
}

type TransactionDraft struct {
	Type          TransactionType  `json:"type"`
	Timestamp     time.Time        `json:"timestamp,omitempty"`
	State         TransactionState `json:"state,omitempty"`
	InteractionID string           `json:"interactionId,omitempty"`
	Amount        *Money           `json:"amount"`
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
