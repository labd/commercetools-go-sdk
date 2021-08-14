// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyShippingMethodsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get ShippingMethod by ID
 */
func (rb *ByProjectKeyShippingMethodsByIdRequestBuilder) Get() *ByProjectKeyShippingMethodsByIdRequestMethodGet {
	return &ByProjectKeyShippingMethodsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/shipping-methods/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update ShippingMethod by ID
 */
func (rb *ByProjectKeyShippingMethodsByIdRequestBuilder) Post(body ShippingMethodUpdate) *ByProjectKeyShippingMethodsByIdRequestMethodPost {
	return &ByProjectKeyShippingMethodsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/shipping-methods/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete ShippingMethod by ID
 */
func (rb *ByProjectKeyShippingMethodsByIdRequestBuilder) Delete() *ByProjectKeyShippingMethodsByIdRequestMethodDelete {
	return &ByProjectKeyShippingMethodsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/shipping-methods/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
