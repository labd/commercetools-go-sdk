// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyStoresKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Get Store by key
 */
func (rb *ByProjectKeyStoresKeyByKeyRequestBuilder) Get() *ByProjectKeyStoresKeyByKeyRequestMethodGet {
	return &ByProjectKeyStoresKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/stores/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Update Store by key
 */
func (rb *ByProjectKeyStoresKeyByKeyRequestBuilder) Post(body StoreUpdate) *ByProjectKeyStoresKeyByKeyRequestMethodPost {
	return &ByProjectKeyStoresKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/stores/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Delete Store by key
 */
func (rb *ByProjectKeyStoresKeyByKeyRequestBuilder) Delete() *ByProjectKeyStoresKeyByKeyRequestMethodDelete {
	return &ByProjectKeyStoresKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/stores/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
