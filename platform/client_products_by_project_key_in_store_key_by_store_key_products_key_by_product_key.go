package platform

// Generated file, please do not change!!!

type ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyRequestBuilder struct {
	projectKey string
	storeKey   string
	productKey string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyRequestBuilder) ProductTailoring() *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		productKey: rb.productKey,
		client:     rb.client,
	}
}
