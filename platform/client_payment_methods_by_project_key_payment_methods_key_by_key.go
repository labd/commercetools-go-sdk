package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyPaymentMethodsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves a PaymentMethod with the provided `key`.
*
 */
func (rb *ByProjectKeyPaymentMethodsKeyByKeyRequestBuilder) Get() *ByProjectKeyPaymentMethodsKeyByKeyRequestMethodGet {
	return &ByProjectKeyPaymentMethodsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/payment-methods/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a PaymentMethod exists with the provided `key`.
*	Returns a `200 OK` status if the PaymentMethod exists; otherwise, returns a [Not Found](/../api/errors#404-not-found).
*
 */
func (rb *ByProjectKeyPaymentMethodsKeyByKeyRequestBuilder) Head() *ByProjectKeyPaymentMethodsKeyByKeyRequestMethodHead {
	return &ByProjectKeyPaymentMethodsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/payment-methods/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a PaymentMethod in the Project using one or more [update actions](/../api/projects/payment-methods#update-actions).
*
 */
func (rb *ByProjectKeyPaymentMethodsKeyByKeyRequestBuilder) Post(body PaymentMethodUpdate) *ByProjectKeyPaymentMethodsKeyByKeyRequestMethodPost {
	return &ByProjectKeyPaymentMethodsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/payment-methods/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a PaymentMethod in the Project.
*	This request generates the [PaymentMethodDeleted](ctp:api:type:PaymentMethodDeletedMessage) Message.
*
 */
func (rb *ByProjectKeyPaymentMethodsKeyByKeyRequestBuilder) Delete() *ByProjectKeyPaymentMethodsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyPaymentMethodsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/payment-methods/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
