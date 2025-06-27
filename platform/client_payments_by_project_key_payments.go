package platform

// Generated file, please do not change!!!

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
func (rb *ByProjectKeyPaymentsRequestBuilder) WithId(id string) *ByProjectKeyPaymentsByIDRequestBuilder {
	return &ByProjectKeyPaymentsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves Payments in the Project.
 */
func (rb *ByProjectKeyPaymentsRequestBuilder) Get() *ByProjectKeyPaymentsRequestMethodGet {
	return &ByProjectKeyPaymentsRequestMethodGet{
		url:    fmt.Sprintf("/%s/payments", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Payments exist for the provided query predicate. Returns a `200 OK` status if any Payments match the query predicate, or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyPaymentsRequestBuilder) Head() *ByProjectKeyPaymentsRequestMethodHead {
	return &ByProjectKeyPaymentsRequestMethodHead{
		url:    fmt.Sprintf("/%s/payments", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a Payment in the Project.
*	Creating a Payment produces the [PaymentCreated](ctp:api:type:PaymentCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyPaymentsRequestBuilder) Post(body PaymentDraft) *ByProjectKeyPaymentsRequestMethodPost {
	return &ByProjectKeyPaymentsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/payments", rb.projectKey),
		client: rb.client,
	}
}
