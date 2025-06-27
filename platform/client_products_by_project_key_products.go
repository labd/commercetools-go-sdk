package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductsRequestBuilder) WithKey(key string) *ByProjectKeyProductsKeyByKeyRequestBuilder {
	return &ByProjectKeyProductsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductsRequestBuilder) WithId(id string) *ByProjectKeyProductsByIDRequestBuilder {
	return &ByProjectKeyProductsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	This endpoint provides high-performance search queries over Products. Product Search allows searching through all products with a current projection in your Project.
*
 */
func (rb *ByProjectKeyProductsRequestBuilder) Search() *ByProjectKeyProductsSearchRequestBuilder {
	return &ByProjectKeyProductsSearchRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	If [Product price selection query parameters](/../api/pricing-and-discounts-overview#product-price-selection) are provided, the selected Prices are added to the response.
 */
func (rb *ByProjectKeyProductsRequestBuilder) Get() *ByProjectKeyProductsRequestMethodGet {
	return &ByProjectKeyProductsRequestMethodGet{
		url:    fmt.Sprintf("/%s/products", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Products exist for the provided query predicate. Returns a `200 OK` status if any Products match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyProductsRequestBuilder) Head() *ByProjectKeyProductsRequestMethodHead {
	return &ByProjectKeyProductsRequestMethodHead{
		url:    fmt.Sprintf("/%s/products", rb.projectKey),
		client: rb.client,
	}
}

/**
*	To create a new Product, send a representation that is going to become the initial _staged_ and _current_ representation of the new Product in the catalog.
*	If [Product price selection query parameters](/../api/pricing-and-discounts-overview#product-price-selection) are provided, selected Prices will be added to the response.
*	Produces the [ProductCreated](/projects/messages/product-catalog-messages#product-created) Message.
*
 */
func (rb *ByProjectKeyProductsRequestBuilder) Post(body ProductDraft) *ByProjectKeyProductsRequestMethodPost {
	return &ByProjectKeyProductsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products", rb.projectKey),
		client: rb.client,
	}
}
