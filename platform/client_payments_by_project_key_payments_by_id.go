package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyPaymentsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves a Payment with the provided `id`.
 */
func (rb *ByProjectKeyPaymentsByIDRequestBuilder) Get() *ByProjectKeyPaymentsByIDRequestMethodGet {
	return &ByProjectKeyPaymentsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Payment exists with the provided `id`. Returns a `200 OK` status if the Payment exists, or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyPaymentsByIDRequestBuilder) Head() *ByProjectKeyPaymentsByIDRequestMethodHead {
	return &ByProjectKeyPaymentsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a Payment in the Project using one or more [update actions](/../api/projects/payments#update-actions).
 */
func (rb *ByProjectKeyPaymentsByIDRequestBuilder) Post(body PaymentUpdate) *ByProjectKeyPaymentsByIDRequestMethodPost {
	return &ByProjectKeyPaymentsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a Payment in the Project.
 */
func (rb *ByProjectKeyPaymentsByIDRequestBuilder) Delete() *ByProjectKeyPaymentsByIDRequestMethodDelete {
	return &ByProjectKeyPaymentsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
