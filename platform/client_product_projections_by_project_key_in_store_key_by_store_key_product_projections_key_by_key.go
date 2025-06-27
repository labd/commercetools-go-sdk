package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestBuilder struct {
	projectKey string
	storeKey   string
	key        string
	client     *Client
}

/**
*	Gets the current or staged representation of a [Product](ctp:api:type:Product) by its key in the specified [Store](ctp:api:type:Store).
*	If the Store has defined some languages, countries, distribution, supply Channels, and/or Product Selection,
*	they are used for projections based on [locale](ctp:api:type:ProductProjectionLocales), [price](ctp:api:type:ProductProjectionPrices),
*	and [inventory](ctp:api:type:ProductProjectionInventoryEntries).
*
*	If [ProductSelection](ctp:api:type:ProductSelection) is used, it affects the [availability of the Product](/projects/stores#products-available-in-store) in the specified Store.
*
*	If a [ProductTailoring](ctp:api:type:ProductTailoring) exists for the Product with the given `key` and the given Store, this endpoint returns the ProductProjection with tailored data.
*
*	When used with an API Client that has the `view_published_products:{projectKey}` scope, this endpoint only returns published (current) Product Projections.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/product-projections/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if the current or staged representations of a Product exists with the provided `key` in the specified [Store](ctp:api:type:Store). Returns a `200 OK` status if the ProductProjection exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/product-projections/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
