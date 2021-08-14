// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyPaymentsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get Payment by ID
 */
func (rb *ByProjectKeyPaymentsByIdRequestBuilder) Get() *ByProjectKeyPaymentsByIdRequestMethodGet {
	return &ByProjectKeyPaymentsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update Payment by ID
 */
func (rb *ByProjectKeyPaymentsByIdRequestBuilder) Post(body PaymentUpdate) *ByProjectKeyPaymentsByIdRequestMethodPost {
	return &ByProjectKeyPaymentsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete Payment by ID
 */
func (rb *ByProjectKeyPaymentsByIdRequestBuilder) Delete() *ByProjectKeyPaymentsByIdRequestMethodDelete {
	return &ByProjectKeyPaymentsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
