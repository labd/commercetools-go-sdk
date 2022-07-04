package platform

// Generated file, please do not change!!!

type ByProjectKeyInStoreKeyByStoreKeyProductProjectionsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
