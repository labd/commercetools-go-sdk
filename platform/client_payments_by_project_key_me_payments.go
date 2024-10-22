package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMePaymentsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMePaymentsRequestBuilder) WithId(id string) *ByProjectKeyMePaymentsByIDRequestBuilder {
	return &ByProjectKeyMePaymentsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Returns all [Payments](/projects/me-payments#mypayment) that match a given Query Predicate.
 */
func (rb *ByProjectKeyMePaymentsRequestBuilder) Get() *ByProjectKeyMePaymentsRequestMethodGet {
	return &ByProjectKeyMePaymentsRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/payments", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a [Payment](/projects/me-payments#mypayment) exists for a given Query Predicate. Returns a `200 OK` status if any Payments match the Query Predicate or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyMePaymentsRequestBuilder) Head() *ByProjectKeyMePaymentsRequestMethodHead {
	return &ByProjectKeyMePaymentsRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/payments", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a [Payment](/projects/me-payments#mypayment) for the Customer or an anonymous user.
*	Creating a Payment produces the [PaymentCreated](ctp:api:type:PaymentCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyMePaymentsRequestBuilder) Post(body MyPaymentDraft) *ByProjectKeyMePaymentsRequestMethodPost {
	return &ByProjectKeyMePaymentsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/payments", rb.projectKey),
		client: rb.client,
	}
}
