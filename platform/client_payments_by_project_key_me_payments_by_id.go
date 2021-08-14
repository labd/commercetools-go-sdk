// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyMePaymentsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get MyPayment by ID
 */
func (rb *ByProjectKeyMePaymentsByIdRequestBuilder) Get() *ByProjectKeyMePaymentsByIdRequestMethodGet {
	return &ByProjectKeyMePaymentsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update MyPayment by ID
 */
func (rb *ByProjectKeyMePaymentsByIdRequestBuilder) Post(body MyPaymentUpdate) *ByProjectKeyMePaymentsByIdRequestMethodPost {
	return &ByProjectKeyMePaymentsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete MyPayment by ID
 */
func (rb *ByProjectKeyMePaymentsByIdRequestBuilder) Delete() *ByProjectKeyMePaymentsByIdRequestMethodDelete {
	return &ByProjectKeyMePaymentsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
