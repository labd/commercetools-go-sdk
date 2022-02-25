package platform

// Generated file, please do not change!!!

import ()

type ByProjectKeyInStoreKeyByStoreKeyShippingMethodsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Get ShippingMethods for a cart in a store
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsRequestBuilder) MatchingCart() *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
