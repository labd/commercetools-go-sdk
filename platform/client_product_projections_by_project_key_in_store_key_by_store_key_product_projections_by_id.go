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
*	Retrieves the [projected](/../api/projects/productProjections#projection-dimensions) representation of a [Product](ctp:api:type:Product) by its ID in the specified [Store](ctp:api:type:Store).
*
*	If the Store has defined some languages, countries, distribution, supply Channels, and/or Product Selection,
*	they are used for projections based on [locale](ctp:api:type:ProductProjectionLocales), [price](ctp:api:type:ProductProjectionPrices),
*	and [inventory](ctp:api:type:ProductProjectionInventoryEntries).
*	If [ProductSelection](ctp:api:type:ProductSelection) is used, it affects the [availability of the Product](/api/project-configuration-overview#products-available-in-store) in the specified Store.
*	If a [ProductTailoring](ctp:api:type:ProductTailoring) exists for the Product with the given `key` and the given Store, this endpoint returns the ProductProjection with tailored data.
*
*	By default, this endpoint returns the `current` representation of Products where the `published` flag is `true`.
*	If a Product is unpublished (`published=false`), the endpoint returns a [Not Found](/../api/errors#404-not-found) error.
*
*	Required access scopes:
*
*	- To retrieve the current representation of published Products (published data), the `view_published_products:{projectKey}` scope is required.
*
*	- To retrieve the staged representation of Products (draft data) or access unpublished Products, the API Client must have the `view_products:{projectKey}` scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/product-projections/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if the current or staged representations of a Product exists with the provided `id` in the specified [Store](ctp:api:type:Store). Returns a `200` status if the ProductProjection exists, or a `404` status otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/product-projections/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
