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
func (rb *ByProjectKeyMePaymentsRequestBuilder) Get() *ByProjectKeyMePaymentsRequestMethodGet {
	return &ByProjectKeyMePaymentsRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/payments", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a Payment exists for a given Query Predicate. Returns a `200 OK` status if any Payments match the Query Predicate or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyMePaymentsRequestBuilder) Head() *ByProjectKeyMePaymentsRequestMethodHead {
	return &ByProjectKeyMePaymentsRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/payments", rb.projectKey),
		client: rb.client,
	}
}

/**
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
