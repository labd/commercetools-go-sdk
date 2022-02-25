package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/shipping-methods/matching-cart", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
