package platform

// Generated file, please do not change!!!

type ByProjectKeyInStoreKeyByStoreKeyProductsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsRequestBuilder) WithProductId(productId string) *ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDRequestBuilder{
		productId:  productId,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsRequestBuilder) WithProductKey(productKey string) *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyRequestBuilder{
		productKey: productKey,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
