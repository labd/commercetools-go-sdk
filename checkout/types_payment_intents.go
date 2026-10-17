package checkout

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
)

type PaymentIntent struct {
	// Action to execute for the given Payment.
	Actions []PaymentIntentAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PaymentIntent) UnmarshalJSON(data []byte) error {
	type Alias PaymentIntent
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorPaymentIntentAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	Depending on the action specified, Checkout requests the [payment service provider](/connectors-and-applications#supported-psps-payment-integration-types-and-payment-methods) (PSP) or gift card management system to capture, refund, or cancel the authorization for the given Payment.
*
 */
type PaymentIntentAction interface{}

func mapDiscriminatorPaymentIntentAction(input interface{}) (PaymentIntentAction, error) {
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
	case "cancelPayment":
		obj := PaymentIntentCancelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "capturePayment":
		obj := PaymentIntentCaptureAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "refundPayment":
		obj := PaymentIntentRefundAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "reversePayment":
		obj := PaymentIntentReverseAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Requests to [cancel the authorization](/payments-lifecycle#authorization-cancellation) for a Payment. Checkout will cancel the [Payment](ctp:api:type:Payment) and will request the PSP or gift card management system to proceed with the financial process to cancel the authorization.
*
*	You cannot request to cancel the authorization for a Payment that has already been [captured](/payments-lifecycle#payment-capture).
*
 */
type PaymentIntentCancelAction struct {
	// A merchant-defined identifier associated with the [Payment](ctp:api:type:Payment) to track and reconcile the [Payment Intent Action](ctp:checkout:type:PaymentIntentAction) on the merchant's side. For example, an invoice number.
	MerchantReference *string `json:"merchantReference,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentIntentCancelAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentIntentCancelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "cancelPayment", Alias: (*Alias)(&obj)})
}

/**
*	Requests to [capture](/payments-lifecycle#payment-capture) the given amount from the customer. Checkout will request the PSP or gift card management system to proceed with the financial process to capture the amount.
*
 */
type PaymentIntentCaptureAction struct {
	// Amount to be captured. It must be less than or equal to the [authorized](/payments-lifecycle#authorization) amount.
	Amount Amount `json:"amount"`
	// A merchant-defined identifier associated with the [Payment](ctp:api:type:Payment) to track and reconcile the [Payment Intent Action](ctp:checkout:type:PaymentIntentAction) on the merchant's side. For example, an invoice number.
	MerchantReference *string `json:"merchantReference,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentIntentCaptureAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentIntentCaptureAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "capturePayment", Alias: (*Alias)(&obj)})
}

/**
*	The possible values for a [Payment Intent Action](ctp:checkout:type:PaymentIntentAction).
*
 */
type PaymentIntentOperation string

const (
	PaymentIntentOperationCapturePayment PaymentIntentOperation = "capturePayment"
	PaymentIntentOperationRefundPayment  PaymentIntentOperation = "refundPayment"
	PaymentIntentOperationCancelPayment  PaymentIntentOperation = "cancelPayment"
	PaymentIntentOperationReversePayment PaymentIntentOperation = "reversePayment"
)

/**
*	The outcome returned by the payment Connector after processing a Payment Intent.
*
 */
type PaymentIntentOutcome string

const (
	PaymentIntentOutcomeApproved PaymentIntentOutcome = "approved"
	PaymentIntentOutcomeRejected PaymentIntentOutcome = "rejected"
	PaymentIntentOutcomeReceived PaymentIntentOutcome = "received"
)

/**
*	Requests to [refund](/payments-lifecycle#refund) the given amount to the customer. Checkout will request the PSP or gift card management system to proceed with the financial process to refund the amount.
*
 */
type PaymentIntentRefundAction struct {
	// Amount to be refunded. It must be less than or equal to the [captured](/payments-lifecycle#payment-capture) amount.
	Amount Amount `json:"amount"`
	// The identifier of the capture transaction that is associated with the refund action.
	TransactionId *string `json:"transactionId,omitempty"`
	// A merchant-defined identifier associated with the [Payment](ctp:api:type:Payment) to track and reconcile the [Payment Intent Action](ctp:checkout:type:PaymentIntentAction) on the merchant's side. For example, an invoice number.
	MerchantReference *string `json:"merchantReference,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentIntentRefundAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentIntentRefundAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "refundPayment", Alias: (*Alias)(&obj)})
}

/**
*	Returned by Checkout after forwarding a Payment Intent request to the payment Connector.
*	If the Connector response does not contain a valid `outcome` value, Checkout returns a `500 Internal Server Error`.
*
 */
type PaymentIntentResponse struct {
	// The outcome of the Payment Intent as reported by the payment [Connector](/connectors-and-applications#payment-connectors).
	Outcome PaymentIntentOutcome `json:"outcome"`
}

/**
*	Requests to [reverse](/payment-integration-predicates) a [Payment](ctp:api:type:Payment). Checkout reverses the Payment, and then requests the PSP or gift card management system to proceed with the relevant process to reverse the Payment.
*
 */
type PaymentIntentReverseAction struct {
	// A merchant-defined identifier associated with the [Payment](ctp:api:type:Payment) to track and reconcile the [Payment Intent Action](ctp:checkout:type:PaymentIntentAction) on the merchant's side. For example, an invoice number.
	MerchantReference *string `json:"merchantReference,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentIntentReverseAction) MarshalJSON() ([]byte, error) {
	type Alias PaymentIntentReverseAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "reversePayment", Alias: (*Alias)(&obj)})
}
