package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyShippingMethodsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves a ShippingMethod with the provided `id`.
 */
func (rb *ByProjectKeyShippingMethodsByIDRequestBuilder) Get() *ByProjectKeyShippingMethodsByIDRequestMethodGet {
	return &ByProjectKeyShippingMethodsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/shipping-methods/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a ShippingMethod exists with the provided `id`. Returns a `200 OK` status if the ShippingMethod exists or a [Not Found](/../api/errors#404-not-found) error otherwise.
 */
func (rb *ByProjectKeyShippingMethodsByIDRequestBuilder) Head() *ByProjectKeyShippingMethodsByIDRequestMethodHead {
	return &ByProjectKeyShippingMethodsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/shipping-methods/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a ShippingMethod in the Project using one or more [update actions](/../api/projects/shippingMethods#update-actions).
 */
func (rb *ByProjectKeyShippingMethodsByIDRequestBuilder) Post(body ShippingMethodUpdate) *ByProjectKeyShippingMethodsByIDRequestMethodPost {
	return &ByProjectKeyShippingMethodsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/shipping-methods/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a ShippingMethod in the Project.
 */
func (rb *ByProjectKeyShippingMethodsByIDRequestBuilder) Delete() *ByProjectKeyShippingMethodsByIDRequestMethodDelete {
	return &ByProjectKeyShippingMethodsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/shipping-methods/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
