package platform

// Generated file, please do not change!!!

type ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDRequestBuilder struct {
	projectKey string
	storeKey   string
	productId  string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDRequestBuilder) ProductTailoring() *ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDProductTailoringRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDProductTailoringRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		productId:  rb.productId,
		client:     rb.client,
	}
}
