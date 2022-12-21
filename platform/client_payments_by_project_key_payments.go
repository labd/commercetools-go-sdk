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
func (rb *ByProjectKeyPaymentsRequestBuilder) Get() *ByProjectKeyPaymentsRequestMethodGet {
	return &ByProjectKeyPaymentsRequestMethodGet{
		url:    fmt.Sprintf("/%s/payments", rb.projectKey),
		client: rb.client,
	}
}

/**
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
