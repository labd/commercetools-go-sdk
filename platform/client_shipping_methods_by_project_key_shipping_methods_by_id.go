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

func (rb *ByProjectKeyShippingMethodsByIDRequestBuilder) Get() *ByProjectKeyShippingMethodsByIDRequestMethodGet {
	return &ByProjectKeyShippingMethodsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/shipping-methods/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a ShippingMethod exists with the provided `id`. Returns a `200 OK` status if the ShippingMethod exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyShippingMethodsByIDRequestBuilder) Head() *ByProjectKeyShippingMethodsByIDRequestMethodHead {
	return &ByProjectKeyShippingMethodsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/shipping-methods/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyShippingMethodsByIDRequestBuilder) Post(body ShippingMethodUpdate) *ByProjectKeyShippingMethodsByIDRequestMethodPost {
	return &ByProjectKeyShippingMethodsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/shipping-methods/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyShippingMethodsByIDRequestBuilder) Delete() *ByProjectKeyShippingMethodsByIDRequestMethodDelete {
	return &ByProjectKeyShippingMethodsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/shipping-methods/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
