// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyTypesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Get Type by key
 */
func (rb *ByProjectKeyTypesKeyByKeyRequestBuilder) Get() *ByProjectKeyTypesKeyByKeyRequestMethodGet {
	return &ByProjectKeyTypesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/types/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Update Type by key
 */
func (rb *ByProjectKeyTypesKeyByKeyRequestBuilder) Post(body TypeUpdate) *ByProjectKeyTypesKeyByKeyRequestMethodPost {
	return &ByProjectKeyTypesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/types/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Delete Type by key
 */
func (rb *ByProjectKeyTypesKeyByKeyRequestBuilder) Delete() *ByProjectKeyTypesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyTypesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/types/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
