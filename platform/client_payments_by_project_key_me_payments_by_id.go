package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMePaymentsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Returns a [Payment](/projects/me-payments#mypayment) for a given `id`.
 */
func (rb *ByProjectKeyMePaymentsByIDRequestBuilder) Get() *ByProjectKeyMePaymentsByIDRequestMethodGet {
	return &ByProjectKeyMePaymentsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a [Payment](/projects/me-payments#mypayment) exists for a given `id`. Returns a `200 OK` status if the Payment exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyMePaymentsByIDRequestBuilder) Head() *ByProjectKeyMePaymentsByIDRequestMethodHead {
	return &ByProjectKeyMePaymentsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a [Payment](/projects/me-payments#mypayment) for a given `id`.
*	You can only update a [Payment](/projects/me-payments#mypayment) if it has no [Transactions](ctp:api:type:Transaction).
*
 */
func (rb *ByProjectKeyMePaymentsByIDRequestBuilder) Post(body MyPaymentUpdate) *ByProjectKeyMePaymentsByIDRequestMethodPost {
	return &ByProjectKeyMePaymentsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes the [Payment](/projects/me-payments#mypayment) for a given `id`.
*	You can only delete a [Payment](/projects/me-payments#mypayment) if it has no [Transactions](ctp:api:type:Transaction).
*
 */
func (rb *ByProjectKeyMePaymentsByIDRequestBuilder) Delete() *ByProjectKeyMePaymentsByIDRequestMethodDelete {
	return &ByProjectKeyMePaymentsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
