package checkout

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	Information for the request to the [Connector](/connectors-and-applications#payment-connectors) to initiate the payment for a specific [Cart](/api/projects/carts).
*
 */
type Transaction struct {
	// Unique identifier of the Transaction.
	ID string `json:"id"`
	// User-defined unique identifier of the Transaction.
	Key *string `json:"key,omitempty"`
	// Current version of the Transaction.
	Version int `json:"version"`
	// [Application](ctp:checkout:type:Application) for which the payment must be executed.
	Application ApplicationReference `json:"application"`
	// Transaction Item associated with the Transaction.
	TransactionItems []TransactionItem `json:"transactionItems"`
	// Reference to the [Cart](ctp:api:type:Cart) for which the payment must be executed.
	Cart CartReference `json:"cart"`
	// Status of the Transaction.
	TransactionStatus TransactionStatus `json:"transactionStatus"`
	// Reference to the [Order](ctp:api:type:Order) created from the Cart when the Transaction is completed.
	Order *OrderReference `json:"order,omitempty"`
	// Date and time (UTC) the Transaction was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Transaction was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Transaction) UnmarshalJSON(data []byte) error {
	type Alias Transaction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.TransactionItems {
		var err error
		obj.TransactionItems[i], err = mapDiscriminatorTransactionItem(obj.TransactionItems[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type TransactionDraft struct {
	// User-defined unique identifier of the Transaction.
	Key *string `json:"key,omitempty"`
	// [Application](ctp:checkout:type:Application) for which the payment is executed.
	Application ApplicationResourceIdentifier `json:"application"`
	// Transaction Item associated with the Transaction.
	TransactionItems []TransactionItemDraft `json:"transactionItems"`
	// [Cart](ctp:api:type:Cart) for which the payment must be executed.
	Cart CartResourceIdentifier `json:"cart"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *TransactionDraft) UnmarshalJSON(data []byte) error {
	type Alias TransactionDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.TransactionItems {
		var err error
		obj.TransactionItems[i], err = mapDiscriminatorTransactionItemDraft(obj.TransactionItems[i])
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	A single error on the [Transaction](/transactions-api#transaction). Multiple errors may be included in the [Transaction Status](/transactions-api#transactionstatus).
*
 */
type TransactionError struct {
	// Error identifier.
	Code string `json:"code"`
	// Plain text description of the cause of the error.
	Message string `json:"message"`
}

func (obj TransactionError) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown TransactionError: failed to parse error response"
}

/**
*	Payment information related to the [Transaction](ctp:checkout:type:Transaction). If `type` is not present, this is a [TransactionItemPaymentIntegration](ctp:checkout:type:TransactionItemPaymentIntegration). Each supported payment flow has its own corresponding Transaction Item type, like [TransactionItemRecurring](ctp:checkout:type:TransactionItemRecurring).
*
 */
type TransactionItem interface{}

func mapDiscriminatorTransactionItem(input interface{}) (TransactionItem, error) {
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
	case "Recurring":
		obj := TransactionItemRecurring{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Base type for creating a Transaction Item. If `type` is omitted, a [TransactionItemPaymentIntegrationDraft](ctp:checkout:type:TransactionItemPaymentIntegrationDraft) is created to process the payment through a Payment Integration. Each supported payment flow has its own corresponding Transaction Item Draft type, like [TransactionItemRecurringDraft](ctp:checkout:type:TransactionItemRecurringDraft).
*
 */
type TransactionItemDraft interface{}

func mapDiscriminatorTransactionItemDraft(input interface{}) (TransactionItemDraft, error) {
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
	case "Recurring":
		obj := TransactionItemRecurringDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	A Transaction Item processed through a [Payment Integration](/checkout/connectors-and-applications#payment-integrations). This is the default Transaction Item type used when `type` is not present.
*
 */
type TransactionItemPaymentIntegration interface{}

func mapDiscriminatorTransactionItemPaymentIntegration(input interface{}) (TransactionItemPaymentIntegration, error) {
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

	}
	return nil, nil
}

/**
*	Creates a Transaction Item that is processed through a [Payment Integration](/checkout/connectors-and-applications#payment-integrations). This is the default Transaction Item type used when `type` is omitted.
*
 */
type TransactionItemPaymentIntegrationDraft interface{}

func mapDiscriminatorTransactionItemPaymentIntegrationDraft(input interface{}) (TransactionItemPaymentIntegrationDraft, error) {
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

	}
	return nil, nil
}

/**
*	A Transaction Item processing a recurring payment for a [PaymentMethod](ctp:api:type:PaymentMethod).
*
 */
type TransactionItemRecurring struct {
	// Money value of the Transaction Item.
	Amount *Amount `json:"amount,omitempty"`
	// Reference to the [Payment](ctp:api:type:Payment) associated with the Transaction Item.
	Payment *PaymentReference `json:"payment,omitempty"`
	// Reference to the [PaymentMethod](ctp:api:type:PaymentMethod) charged for the Transaction Item.
	PaymentMethod PaymentMethodReference `json:"paymentMethod"`
	// Reference to the connector deployment used to execute the payment.
	ConnectorDeployment ConnectorDeploymentReference `json:"connectorDeployment"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TransactionItemRecurring) MarshalJSON() ([]byte, error) {
	type Alias TransactionItemRecurring
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Recurring", Alias: (*Alias)(&obj)})
}

/**
*	Creates a Transaction Item to process a recurring payment for a [PaymentMethod](ctp:api:type:PaymentMethod) using a specific connector deployment. The [Cart](ctp:api:type:Cart) referenced by the Transaction must have the same `customerId` as the referenced PaymentMethod.
*
 */
type TransactionItemRecurringDraft struct {
	// Money value of the Transaction Item. If not present, the Connector resolves the amount from the [Cart](ctp:api:type:Cart).
	Amount *Amount `json:"amount,omitempty"`
	// Reference to the [PaymentMethod](ctp:api:type:PaymentMethod) to charge. The PaymentMethod must belong to the same customer as the Cart referenced by the Transaction.
	PaymentMethod PaymentMethodReference `json:"paymentMethod"`
	// Reference to the connector deployment to use to execute the payment.
	ConnectorDeployment ConnectorDeploymentReference `json:"connectorDeployment"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TransactionItemRecurringDraft) MarshalJSON() ([]byte, error) {
	type Alias TransactionItemRecurringDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "Recurring", Alias: (*Alias)(&obj)})
}

/**
*	The state of the [Transaction](/transactions-api#transaction).
*
 */
type TransactionState string

const (
	TransactionStateInitial   TransactionState = "Initial"
	TransactionStatePending   TransactionState = "Pending"
	TransactionStateCompleted TransactionState = "Completed"
	TransactionStateFailed    TransactionState = "Failed"
)

/**
*	The state of the [Transaction](/transactions-api#transaction) and the related errors in case of a failed Transaction.
*
 */
type TransactionStatus struct {
	// State of the Transaction.
	State TransactionState `json:"state"`
	// Errors returned if the Transaction is in the `Failed` state.
	Errors []TransactionError `json:"errors"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TransactionStatus) MarshalJSON() ([]byte, error) {
	type Alias TransactionStatus
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

	if raw["errors"] == nil {
		delete(raw, "errors")
	}

	return json.Marshal(raw)

}
