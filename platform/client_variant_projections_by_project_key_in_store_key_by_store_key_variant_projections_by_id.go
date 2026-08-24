package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Retrieves a Variant Projection by its ID in the specified [Store](ctp:api:type:Store).
*
*	Only returns Variants belonging to Products distributed through the Store's configured channels.
*
*	If a [ProductTailoring](ctp:api:type:ProductTailoring) exists for the Product and the given Store, the returned Variant Projection includes the tailored Product name, slug, and description, along with the tailored Variant images, assets, and attributes. The `staged` query parameter selects either the current or staged tailored data. When no Product Tailoring exists for the Product and Store, the Variant Projection is returned without Product Tailoring data.
*
*	Required access scopes:
*
*	- To retrieve the current representation, the `view_published_products:{projectKey}` scope is required.
*
*	- To retrieve the staged representation, the API Client must have the `view_products:{projectKey}` scope.
*
*	- To access Variant Projections in the context of a Store, the `view_products:{projectKey}:{storeKey}` scope is required.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/variant-projections/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Variant Projection exists with the provided `id` in the specified [Store](ctp:api:type:Store).
*	Returns a `200` status if the Variant Projection exists, or a `404` status otherwise.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/variant-projections/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
