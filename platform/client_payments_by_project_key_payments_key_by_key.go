package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyPaymentsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves a Payment with the provided `key`.
 */
func (rb *ByProjectKeyPaymentsKeyByKeyRequestBuilder) Get() *ByProjectKeyPaymentsKeyByKeyRequestMethodGet {
	return &ByProjectKeyPaymentsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/payments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Payment exists with the provided `key`. Returns a `200 OK` status if the Payment exists, or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyPaymentsKeyByKeyRequestBuilder) Head() *ByProjectKeyPaymentsKeyByKeyRequestMethodHead {
	return &ByProjectKeyPaymentsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/payments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a Payment in the Project using one or more [update actions](/../api/projects/payments#update-actions).
 */
func (rb *ByProjectKeyPaymentsKeyByKeyRequestBuilder) Post(body PaymentUpdate) *ByProjectKeyPaymentsKeyByKeyRequestMethodPost {
	return &ByProjectKeyPaymentsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/payments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a Payment in the Project.
 */
func (rb *ByProjectKeyPaymentsKeyByKeyRequestBuilder) Delete() *ByProjectKeyPaymentsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyPaymentsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/payments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
