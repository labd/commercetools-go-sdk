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
*	Retrieves a [Payment](/projects/me-payments#mypayment) with the provided `id` for the authenticated Customer or anonymous user.
 */
func (rb *ByProjectKeyMePaymentsByIDRequestBuilder) Get() *ByProjectKeyMePaymentsByIDRequestMethodGet {
	return &ByProjectKeyMePaymentsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a [Payment](/projects/me-payments#mypayment) exists with the provided `id` for the authenticated Customer or anonymous user. Returns a `200 OK` status if the Payment exists, or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyMePaymentsByIDRequestBuilder) Head() *ByProjectKeyMePaymentsByIDRequestMethodHead {
	return &ByProjectKeyMePaymentsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a [Payment](/projects/me-payments#mypayment) for the authenticated Customer or anonymous user using one or more [update actions](/../api/projects/me-payments#update-actions).
*	You can only update a Payment if it has no [Transactions](ctp:api:type:Transaction).
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
*	Deletes a [Payment](/projects/me-payments#mypayment) with the provided `id` for the authenticated Customer or anonymous user.
*	You can only delete a Payment if it has no [Transactions](ctp:api:type:Transaction).
*
 */
func (rb *ByProjectKeyMePaymentsByIDRequestBuilder) Delete() *ByProjectKeyMePaymentsByIDRequestMethodDelete {
	return &ByProjectKeyMePaymentsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
