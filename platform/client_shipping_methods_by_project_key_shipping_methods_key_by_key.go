package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyShippingMethodsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves a ShippingMethod with the provided `key`.
 */
func (rb *ByProjectKeyShippingMethodsKeyByKeyRequestBuilder) Get() *ByProjectKeyShippingMethodsKeyByKeyRequestMethodGet {
	return &ByProjectKeyShippingMethodsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/shipping-methods/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a ShippingMethod exists with the provided `key`. Returns a `200 OK` status if the ShippingMethod exists or a [Not Found](/../api/errors#404-not-found) error otherwise.
 */
func (rb *ByProjectKeyShippingMethodsKeyByKeyRequestBuilder) Head() *ByProjectKeyShippingMethodsKeyByKeyRequestMethodHead {
	return &ByProjectKeyShippingMethodsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/shipping-methods/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a ShippingMethod in the Project using one or more [update actions](/../api/projects/shippingMethods#update-actions).
 */
func (rb *ByProjectKeyShippingMethodsKeyByKeyRequestBuilder) Post(body ShippingMethodUpdate) *ByProjectKeyShippingMethodsKeyByKeyRequestMethodPost {
	return &ByProjectKeyShippingMethodsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/shipping-methods/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a ShippingMethod in the Project.
 */
func (rb *ByProjectKeyShippingMethodsKeyByKeyRequestBuilder) Delete() *ByProjectKeyShippingMethodsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyShippingMethodsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/shipping-methods/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
