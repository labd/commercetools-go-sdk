package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyPaymentMethodsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves a PaymentMethod with the provided `id`.
*
 */
func (rb *ByProjectKeyPaymentMethodsByIDRequestBuilder) Get() *ByProjectKeyPaymentMethodsByIDRequestMethodGet {
	return &ByProjectKeyPaymentMethodsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/payment-methods/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a PaymentMethod exists with the provided `id`.
*	Returns a `200 OK` status if the PaymentMethod exists; otherwise, returns a [Not Found](/../api/errors#404-not-found).
*
 */
func (rb *ByProjectKeyPaymentMethodsByIDRequestBuilder) Head() *ByProjectKeyPaymentMethodsByIDRequestMethodHead {
	return &ByProjectKeyPaymentMethodsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/payment-methods/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a PaymentMethod in the Project using one or more [update actions](/../api/projects/payment-methods#update-actions).
*
 */
func (rb *ByProjectKeyPaymentMethodsByIDRequestBuilder) Post(body PaymentMethodUpdate) *ByProjectKeyPaymentMethodsByIDRequestMethodPost {
	return &ByProjectKeyPaymentMethodsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/payment-methods/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a PaymentMethod in the Project.
*	This request generates the [PaymentMethodDeleted](ctp:api:type:PaymentMethodDeletedMessage) Message.
*
 */
func (rb *ByProjectKeyPaymentMethodsByIDRequestBuilder) Delete() *ByProjectKeyPaymentMethodsByIDRequestMethodDelete {
	return &ByProjectKeyPaymentMethodsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/payment-methods/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
