package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	Queries Variant Projections in the specified [Store](ctp:api:type:Store).
*
*	Only Variants belonging to Products distributed through the Store's configured channels are returned.
*	By default, this endpoint returns the `current` representation of Variants where the `published` flag is `true`.
*
*	If a [ProductTailoring](ctp:api:type:ProductTailoring) exists for a Product and the given Store, the returned Variant Projections include the tailored Product name, slug, and description, along with the tailored Variant images, assets, and attributes. The `staged` query parameter selects either the current or staged tailored data. When no Product Tailoring exists for a Product and Store, the Variant Projection is returned without Product Tailoring data.
*
*	Required access scopes:
*
*	- To retrieve the current representation of published Variants, the `view_published_products:{projectKey}` scope is required.
*
*	- To retrieve the staged representation, the API Client must have the `view_products:{projectKey}` scope.
*
*	- To access Variant Projections in the context of a Store, the `view_products:{projectKey}:{storeKey}` scope is required.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/variant-projections", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Variant Projections exist for the provided query predicate in the specified [Store](ctp:api:type:Store).
*	Returns a `200` status if any Variant Projections match, or a `404` status otherwise.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyVariantProjectionsRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/variant-projections", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
