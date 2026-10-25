package checkout

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyPaymentIntentsByPaymentIdRequestBuilder struct {
	projectKey string
	paymentId  string
	client     *Client
}

/**
*	Requests to the payment [Connector](/checkout/connectors-and-applications#payment-connectors) time out after 30 seconds.
*
*	Specific Error Codes:
*	- [MultipleActionsNotAllowed](ctp:checkout:type:MultipleActionsNotAllowedError)
*	- [RequiredField](ctp:checkout:type:RequiredFieldError)
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [ConnectorTimeout](ctp:checkout:type:ConnectorTimeoutError)
*
 */
func (rb *ByProjectKeyPaymentIntentsByPaymentIdRequestBuilder) Post(body PaymentIntent) *ByProjectKeyPaymentIntentsByPaymentIdRequestMethodPost {
	return &ByProjectKeyPaymentIntentsByPaymentIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/payment-intents/%s", rb.projectKey, rb.paymentId),
		client: rb.client,
	}
}
