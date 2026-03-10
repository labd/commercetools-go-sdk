package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductProjectionsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves the [projected](/../api/projects/productProjections#projection-dimensions) representation of a [Product](ctp:api:type:Product) by its ID.
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
func (rb *ByProjectKeyProductProjectionsByIDRequestBuilder) Get() *ByProjectKeyProductProjectionsByIDRequestMethodGet {
	return &ByProjectKeyProductProjectionsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-projections/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if the current or staged representation of a Product exists with the provided `id`. Returns a `200` status if the ProductProjection exists, or a `404` status otherwise.
 */
func (rb *ByProjectKeyProductProjectionsByIDRequestBuilder) Head() *ByProjectKeyProductProjectionsByIDRequestMethodHead {
	return &ByProjectKeyProductProjectionsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/product-projections/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
