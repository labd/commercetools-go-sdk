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
*	Gets the current or staged representation of a [Product](ctp:api:type:Product) by its ID from the specified [Store](ctp:api:type:Store).
*	If the Store has defined some languages, countries, distribution or supply Channels,
*	they are used for projections based on [locale](ctp:api:type:ProductProjectionLocales), [price](ctp:api:type:ProductProjectionPrices)
*	and [inventory](ctp:api:type:ProductProjectionInventoryEntries).
*
*	When used with an API Client that has the `view_published_products:{projectKey}` scope, this endpoint only returns published (current) Product Projections.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/product-projections/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
