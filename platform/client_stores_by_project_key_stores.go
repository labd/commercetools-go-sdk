// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyStoresRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyStoresRequestBuilder) WithKey(key string) *ByProjectKeyStoresKeyByKeyRequestBuilder {
	return &ByProjectKeyStoresKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyStoresRequestBuilder) WithId(id string) *ByProjectKeyStoresByIdRequestBuilder {
	return &ByProjectKeyStoresByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Query stores
 */
func (rb *ByProjectKeyStoresRequestBuilder) Get() *ByProjectKeyStoresRequestMethodGet {
	return &ByProjectKeyStoresRequestMethodGet{
		url:    fmt.Sprintf("/%s/stores", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Create Store
 */
func (rb *ByProjectKeyStoresRequestBuilder) Post(body StoreDraft) *ByProjectKeyStoresRequestMethodPost {
	return &ByProjectKeyStoresRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/stores", rb.projectKey),
		client: rb.client,
	}
}
