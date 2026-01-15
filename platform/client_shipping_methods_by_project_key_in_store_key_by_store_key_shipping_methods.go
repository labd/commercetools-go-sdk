package platform

// Generated file, please do not change!!!

type ByProjectKeyInStoreKeyByStoreKeyShippingMethodsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Get matching ShippingMethods for a Cart in a Store
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsRequestBuilder) MatchingCart() *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
