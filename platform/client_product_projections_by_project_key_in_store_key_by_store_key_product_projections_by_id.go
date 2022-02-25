package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Gets the current or staged representation of a product by its ID from a specific Store.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/product-projections/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
