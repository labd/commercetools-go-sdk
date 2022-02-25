package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyShippingMethodsMatchingCartRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyShippingMethodsMatchingCartRequestBuilder) Get() *ByProjectKeyShippingMethodsMatchingCartRequestMethodGet {
	return &ByProjectKeyShippingMethodsMatchingCartRequestMethodGet{
		url:    fmt.Sprintf("/%s/shipping-methods/matching-cart", rb.projectKey),
		client: rb.client,
	}
}
