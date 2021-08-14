// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyPaymentsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyPaymentsRequestBuilder) WithKey(key string) *ByProjectKeyPaymentsKeyByKeyRequestBuilder {
	return &ByProjectKeyPaymentsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyPaymentsRequestBuilder) WithId(id string) *ByProjectKeyPaymentsByIdRequestBuilder {
	return &ByProjectKeyPaymentsByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Query payments
 */
func (rb *ByProjectKeyPaymentsRequestBuilder) Get() *ByProjectKeyPaymentsRequestMethodGet {
	return &ByProjectKeyPaymentsRequestMethodGet{
		url:    fmt.Sprintf("/%s/payments", rb.projectKey),
		client: rb.client,
	}
}

/**
*	To create a payment object a payment draft object has to be given with the request.
 */
func (rb *ByProjectKeyPaymentsRequestBuilder) Post(body PaymentDraft) *ByProjectKeyPaymentsRequestMethodPost {
	return &ByProjectKeyPaymentsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/payments", rb.projectKey),
		client: rb.client,
	}
}
