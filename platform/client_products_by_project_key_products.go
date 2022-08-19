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
*	If [Price selection](ctp:api:type:ProductPriceSelection) query parameters are provided, the selected Prices are added to the response.
 */
func (rb *ByProjectKeyProductsRequestBuilder) Get() *ByProjectKeyProductsRequestMethodGet {
	return &ByProjectKeyProductsRequestMethodGet{
		url:    fmt.Sprintf("/%s/products", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Check if Products exist. Responds with a `200 OK` status if any Products match the Query Predicate, or `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyProductsRequestBuilder) Head() *ByProjectKeyProductsRequestMethodHead {
	return &ByProjectKeyProductsRequestMethodHead{
		url:    fmt.Sprintf("/%s/products", rb.projectKey),
		client: rb.client,
	}
}

/**
*	To create a new Product, send a representation that is going to become the initial _staged_ representation of the new Product in the master catalog.
*	If [Price Selection](ctp:api:type:ProductPriceSelection) query parameters are provided, selected Prices will be added to the response.
*	Produces the [ProductCreatedMessage](/message-types#productcreatedmessage).
*
 */
func (rb *ByProjectKeyProductsRequestBuilder) Post(body ProductDraft) *ByProjectKeyProductsRequestMethodPost {
	return &ByProjectKeyProductsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products", rb.projectKey),
		client: rb.client,
	}
}
