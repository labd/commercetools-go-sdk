package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductProjectionsKeyByKeyVariantAttributesRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Returns a lightweight representation of all Variants for a [Product](ctp:api:type:Product), including only the requested variant-level [Attributes](ctp:api:type:Attribute) and minimal availability data.
*
*	Designed for building attribute selectors on product detail pages (PDPs) with large numbers of variants.
*	Only available for Projects with [productCatalogModel](ctp:api:type:ProductCatalogModel) set to `Modular`.
*
*	Product-level Attributes are omitted from the response.
*
*	Required access scopes:
*
*	- To retrieve the current representation of published Products (published data), the `view_published_products:{projectKey}` or `view_products:{projectKey}` scope is required.
*	- To retrieve the staged representation of Products (draft data), the API Client must have the `view_products:{projectKey}` scope.
*
 */
func (rb *ByProjectKeyProductProjectionsKeyByKeyVariantAttributesRequestBuilder) Get() *ByProjectKeyProductProjectionsKeyByKeyVariantAttributesRequestMethodGet {
	return &ByProjectKeyProductProjectionsKeyByKeyVariantAttributesRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-projections/key=%s/variant-attributes", rb.projectKey, rb.key),
		client: rb.client,
	}
}
