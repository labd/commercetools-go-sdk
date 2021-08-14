// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyStoresByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get Store by ID
 */
func (rb *ByProjectKeyStoresByIdRequestBuilder) Get() *ByProjectKeyStoresByIdRequestMethodGet {
	return &ByProjectKeyStoresByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/stores/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update Store by ID
 */
func (rb *ByProjectKeyStoresByIdRequestBuilder) Post(body StoreUpdate) *ByProjectKeyStoresByIdRequestMethodPost {
	return &ByProjectKeyStoresByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/stores/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete Store by ID
 */
func (rb *ByProjectKeyStoresByIdRequestBuilder) Delete() *ByProjectKeyStoresByIdRequestMethodDelete {
	return &ByProjectKeyStoresByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/stores/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
