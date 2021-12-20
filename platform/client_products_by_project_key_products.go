// Generated file, please do not change!!!
package platform

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
*	You can use the query endpoint to get the full representations of products.
*	REMARK: We suggest to use the performance optimized search endpoint which has a bunch functionalities,
*	the query API lacks like sorting on custom attributes, etc.
*
 */
func (rb *ByProjectKeyProductsRequestBuilder) Get() *ByProjectKeyProductsRequestMethodGet {
	return &ByProjectKeyProductsRequestMethodGet{
		url:    fmt.Sprintf("/%s/products", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if products exist.
 */
func (rb *ByProjectKeyProductsRequestBuilder) Head() *ByProjectKeyProductsRequestMethodHead {
	return &ByProjectKeyProductsRequestMethodHead{
		url:    fmt.Sprintf("/%s/products", rb.projectKey),
		client: rb.client,
	}
}

/**
*	To create a new product, send a representation that is going to become the initial staged representation
*	of the new product in the master catalog. If price selection query parameters are provided,
*	the selected prices will be added to the response.
*
 */
func (rb *ByProjectKeyProductsRequestBuilder) Post(body ProductDraft) *ByProjectKeyProductsRequestMethodPost {
	return &ByProjectKeyProductsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products", rb.projectKey),
		client: rb.client,
	}
}
