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
*	Specific Error Codes:
*	- [MultipleActionsNotAllowed](ctp:checkout:type:MultipleActionsNotAllowedError)
*	- [RequiredField](ctp:checkout:type:RequiredFieldError)
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*
 */
func (rb *ByProjectKeyPaymentIntentsByPaymentIdRequestBuilder) Post(body Payment) *ByProjectKeyPaymentIntentsByPaymentIdRequestMethodPost {
	return &ByProjectKeyPaymentIntentsByPaymentIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/payment-intents/%s", rb.projectKey, rb.paymentId),
		client: rb.client,
	}
}
