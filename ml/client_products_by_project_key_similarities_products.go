// Generated file, please do not change!!!
package ml

import (
	"fmt"
)

type ByProjectKeySimilaritiesProductsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeySimilaritiesProductsRequestBuilder) Status() *ByProjectKeySimilaritiesProductsStatusRequestBuilder {
	return &ByProjectKeySimilaritiesProductsStatusRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeySimilaritiesProductsRequestBuilder) Post(body SimilarProductSearchRequest) *ByProjectKeySimilaritiesProductsRequestMethodPost {
	return &ByProjectKeySimilaritiesProductsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/similarities/products", rb.projectKey),
		client: rb.client,
	}
}
