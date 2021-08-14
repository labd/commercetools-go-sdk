// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyTypesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyTypesRequestBuilder) WithKey(key string) *ByProjectKeyTypesKeyByKeyRequestBuilder {
	return &ByProjectKeyTypesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyTypesRequestBuilder) WithId(id string) *ByProjectKeyTypesByIdRequestBuilder {
	return &ByProjectKeyTypesByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Query types
 */
func (rb *ByProjectKeyTypesRequestBuilder) Get() *ByProjectKeyTypesRequestMethodGet {
	return &ByProjectKeyTypesRequestMethodGet{
		url:    fmt.Sprintf("/%s/types", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Create Type
 */
func (rb *ByProjectKeyTypesRequestBuilder) Post(body TypeDraft) *ByProjectKeyTypesRequestMethodPost {
	return &ByProjectKeyTypesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/types", rb.projectKey),
		client: rb.client,
	}
}
