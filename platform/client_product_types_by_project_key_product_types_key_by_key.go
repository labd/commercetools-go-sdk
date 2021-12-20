// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyProductTypesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyProductTypesKeyByKeyRequestBuilder) Get() *ByProjectKeyProductTypesKeyByKeyRequestMethodGet {
	return &ByProjectKeyProductTypesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-types/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductTypesKeyByKeyRequestBuilder) Post(body ProductTypeUpdate) *ByProjectKeyProductTypesKeyByKeyRequestMethodPost {
	return &ByProjectKeyProductTypesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-types/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductTypesKeyByKeyRequestBuilder) Delete() *ByProjectKeyProductTypesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyProductTypesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/product-types/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
