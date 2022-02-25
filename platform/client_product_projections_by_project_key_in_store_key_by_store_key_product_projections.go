package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

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

/**
*	You can use the product projections query endpoint to get the current or staged representations of Products.
*	When used with an API client that has the view_published_products:{projectKey} scope,
*	this endpoint only returns published (current) product projections.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/product-projections", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
