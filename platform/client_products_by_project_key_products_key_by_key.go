package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) ProductSelections() *ByProjectKeyProductsKeyByKeyProductSelectionsRequestBuilder {
	return &ByProjectKeyProductsKeyByKeyProductSelectionsRequestBuilder{
		projectKey: rb.projectKey,
		key:        rb.key,
		client:     rb.client,
	}
}

/**
*	If [Price selection](ctp:api:type:ProductPriceSelection) query parameters are provided, the selected Prices are added to the response.
 */
func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) Get() *ByProjectKeyProductsKeyByKeyRequestMethodGet {
	return &ByProjectKeyProductsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/products/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Check if a Product exists with a specified `key`. Responds with a `200 OK` status if the Product exists or `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) Head() *ByProjectKeyProductsKeyByKeyRequestMethodHead {
	return &ByProjectKeyProductsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/products/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) Post(body ProductUpdate) *ByProjectKeyProductsKeyByKeyRequestMethodPost {
	return &ByProjectKeyProductsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	If [Price selection](ctp:api:type:ProductPriceSelection) query parameters are provided, the selected Prices are added to the response.
*	Produces the [ProductDeleted](/projects/messages#product-deleted) Message.
 */
func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) Delete() *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyProductsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/products/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
