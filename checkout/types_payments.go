package checkout

// Generated file, please do not change!!!

import (
	"encoding/json"
)

/**
*	Requests to [capture](/payments-lifecycle#payment-capture) the given amount from the customer. Checkout will request the PSP to proceed with the financial process to capture the amount.
*
 */
type CapturePaymentAction struct {
	// Amount to be captured. It must be less than or equal to the [authorized](/payments-lifecycle#authorization) amount.
	Amount Amount `json:"amount"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CapturePaymentAction) MarshalJSON() ([]byte, error) {
	type Alias CapturePaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "capturePayment", Alias: (*Alias)(&obj)})
}

/**
*	Requests to [refund](/payments-lifecycle#refund) the given amount to the customer. Checkout will request the PSP to proceed with the financial process to refund the amount.
*
 */
type RefundPaymentAction struct {
	// Amount to be refunded. It must be less than or equal to the [captured](/payments-lifecycle#payment-capture) amount.
	Amount *Amount `json:"amount,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RefundPaymentAction) MarshalJSON() ([]byte, error) {
	type Alias RefundPaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "refundPayment", Alias: (*Alias)(&obj)})
}

/**
*	Requests to [cancel the authorization](/payments-lifecycle#authorization-cancellation) for a Payment. Checkout will cancel the [Payment](/../api/projects/payments#payment) and will request the PSP to proceed with the financial process to cancel the authorization.
*
*	You cannot request to cancel the authorization for a Payment that has already been [captured](/payments-lifecycle#payment-capture).
*
 */
type CancelPaymentAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CancelPaymentAction) MarshalJSON() ([]byte, error) {
	type Alias CancelPaymentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "cancelPayment", Alias: (*Alias)(&obj)})
}

type Payment struct {
	// Action to execute for the given Payment.
	Actions []PaymentAction `json:"actions"`
}
