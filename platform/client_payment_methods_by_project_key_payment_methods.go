package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyPaymentMethodsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyPaymentMethodsRequestBuilder) WithKey(key string) *ByProjectKeyPaymentMethodsKeyByKeyRequestBuilder {
	return &ByProjectKeyPaymentMethodsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyPaymentMethodsRequestBuilder) WithId(id string) *ByProjectKeyPaymentMethodsByIDRequestBuilder {
	return &ByProjectKeyPaymentMethodsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves all PaymentMethods in the Project.
*
 */
func (rb *ByProjectKeyPaymentMethodsRequestBuilder) Get() *ByProjectKeyPaymentMethodsRequestMethodGet {
	return &ByProjectKeyPaymentMethodsRequestMethodGet{
		url:    fmt.Sprintf("/%s/payment-methods", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more PaymentMethods exist for the provided query predicate.
*	Returns a `200 OK` status if any PaymentMethods match the query predicate; otherwise, returns a [Not Found](/../api/errors#404-not-found).
*
 */
func (rb *ByProjectKeyPaymentMethodsRequestBuilder) Head() *ByProjectKeyPaymentMethodsRequestMethodHead {
	return &ByProjectKeyPaymentMethodsRequestMethodHead{
		url:    fmt.Sprintf("/%s/payment-methods", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a PaymentMethod in the Project.
*	This request generates the [PaymentMethodCreated](ctp:api:type:PaymentMethodCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyPaymentMethodsRequestBuilder) Post(body PaymentMethodDraft) *ByProjectKeyPaymentMethodsRequestMethodPost {
	return &ByProjectKeyPaymentMethodsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/payment-methods", rb.projectKey),
		client: rb.client,
	}
}
