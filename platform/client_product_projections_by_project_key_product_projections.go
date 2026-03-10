package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductProjectionsRequestBuilder struct {
	projectKey string
	client     *Client
}

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
*	Retrieves the [projected](/../api/projects/productProjections#projection-dimensions) representation of [Products](ctp:api:type:Product) by [query predicates](/../api/predicates/query).
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
func (rb *ByProjectKeyProductProjectionsRequestBuilder) Get() *ByProjectKeyProductProjectionsRequestMethodGet {
	return &ByProjectKeyProductProjectionsRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-projections", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if the current or staged representation of a Product exists for the provided query predicate. Returns a `200` status if any ProductProjections match the query predicate, or a `404` status otherwise.
 */
func (rb *ByProjectKeyProductProjectionsRequestBuilder) Head() *ByProjectKeyProductProjectionsRequestMethodHead {
	return &ByProjectKeyProductProjectionsRequestMethodHead{
		url:    fmt.Sprintf("/%s/product-projections", rb.projectKey),
		client: rb.client,
	}
}
