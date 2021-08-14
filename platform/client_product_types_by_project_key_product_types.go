// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyProductTypesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductTypesRequestBuilder) WithKey(key string) *ByProjectKeyProductTypesKeyByKeyRequestBuilder {
	return &ByProjectKeyProductTypesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductTypesRequestBuilder) WithId(id string) *ByProjectKeyProductTypesByIDRequestBuilder {
	return &ByProjectKeyProductTypesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductTypesRequestBuilder) Get() *ByProjectKeyProductTypesRequestMethodGet {
	return &ByProjectKeyProductTypesRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-types", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductTypesRequestBuilder) Post(body ProductTypeDraft) *ByProjectKeyProductTypesRequestMethodPost {
	return &ByProjectKeyProductTypesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-types", rb.projectKey),
		client: rb.client,
	}
}
