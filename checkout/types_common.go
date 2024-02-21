package checkout

// Generated file, please do not change!!!

/**
*	The Region in which the Checkout application is [hosted](/../checkout/installing-checkout#regions-and-hosts).
*
 */
type Region string

const (
	RegionEuropeWest1Gcp         Region = "europe-west1.gcp"
	RegionUsCentral1Gcp          Region = "us-central1.gcp"
	RegionAustraliaSoutheast1Gcp Region = "australia-southeast1.gcp"
)

/**
*	The amount related to a [payment action](ctp:checkout:type:PaymentAction).
*
 */
type Amount struct {
	// Amount in the smallest indivisible unit of a currency, such as:
	//
	// * Cents for EUR and USD, pence for GBP, or centime for CHF (5 CHF is specified as `500`).
	// * The value in the major unit for currencies without minor units, like JPY (5 JPY is specified as `5`).
	CentAmount int `json:"centAmount"`
	// Currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode string `json:"currencyCode"`
}

/**
*	The possible values for a [payment action](ctp:checkout:type:PaymentAction).
*
 */
type PaymentOperation string

const (
	PaymentOperationCapturePayment PaymentOperation = "capturePayment"
	PaymentOperationRefundPayment  PaymentOperation = "refundPayment"
	PaymentOperationCancelPayment  PaymentOperation = "cancelPayment"
)

/**
*	Depending on the action specified, Checkout requests the [Payment Service Provider](/../checkout/configuring-checkout#supported-payment-service-providers) to capture, refund, or cancel the authorization for the given Payment.
*
 */
type PaymentAction struct {
	// Action to execute for the given Payment.
	Action PaymentOperation `json:"action"`
	// Amount to be captured or refunded.
	Amount *Amount `json:"amount,omitempty"`
}
