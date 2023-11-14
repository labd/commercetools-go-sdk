package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductProjectionsRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	This endpoint provides high performance search queries over ProductProjections. The query result contains the
*	ProductProjections for which at least one ProductVariant matches the search query. This means that variants can
*	be included in the result also for which the search query does not match. To determine which ProductVariants match
*	the search query, the returned ProductProjections include the additional field isMatchingVariant.
*
 */
func (rb *ByProjectKeyProductProjectionsRequestBuilder) Search() *ByProjectKeyProductProjectionsSearchRequestBuilder {
	return &ByProjectKeyProductProjectionsSearchRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	The source of data for suggestions is the searchKeyword field in a product
 */
func (rb *ByProjectKeyProductProjectionsRequestBuilder) Suggest() *ByProjectKeyProductProjectionsSuggestRequestBuilder {
	return &ByProjectKeyProductProjectionsSuggestRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductProjectionsRequestBuilder) WithKey(key string) *ByProjectKeyProductProjectionsKeyByKeyRequestBuilder {
	return &ByProjectKeyProductProjectionsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductProjectionsRequestBuilder) WithId(id string) *ByProjectKeyProductProjectionsByIDRequestBuilder {
	return &ByProjectKeyProductProjectionsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Use the Product Projections query endpoint to get the current or staged representations of Products.
*	When used with an API Client that has the `view_published_products:{projectKey}` scope,
*	this endpoint only returns published (current) Product Projections.
*
 */
func (rb *ByProjectKeyProductProjectionsRequestBuilder) Get() *ByProjectKeyProductProjectionsRequestMethodGet {
	return &ByProjectKeyProductProjectionsRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-projections", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if the current or staged representation of a Product exists for a given Query Predicate. Returns a `200 OK` status if any ProductProjections match the Query Predicate or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyProductProjectionsRequestBuilder) Head() *ByProjectKeyProductProjectionsRequestMethodHead {
	return &ByProjectKeyProductProjectionsRequestMethodHead{
		url:    fmt.Sprintf("/%s/product-projections", rb.projectKey),
		client: rb.client,
	}
}
