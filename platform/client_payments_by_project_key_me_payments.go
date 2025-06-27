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
*	Retrieves [Payments](/projects/me-payments#mypayment) for the authenticated Customer or anonymous user.
 */
func (rb *ByProjectKeyMePaymentsRequestBuilder) Get() *ByProjectKeyMePaymentsRequestMethodGet {
	return &ByProjectKeyMePaymentsRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/payments", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more [Payments](/projects/me-payments#mypayment) exist for the provided query predicate for the authenticated Customer or anonymous user. Returns a `200 OK` status if any Payments match the query predicate, or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyMePaymentsRequestBuilder) Head() *ByProjectKeyMePaymentsRequestMethodHead {
	return &ByProjectKeyMePaymentsRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/payments", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a [Payment](/projects/me-payments#mypayment) for the authenticated Customer or anonymous user.
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
