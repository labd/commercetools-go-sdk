// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyStoresByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyStoresByIDRequestBuilder) Get() *ByProjectKeyStoresByIDRequestMethodGet {
	return &ByProjectKeyStoresByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/stores/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStoresByIDRequestBuilder) Post(body StoreUpdate) *ByProjectKeyStoresByIDRequestMethodPost {
	return &ByProjectKeyStoresByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/stores/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStoresByIDRequestBuilder) Delete() *ByProjectKeyStoresByIDRequestMethodDelete {
	return &ByProjectKeyStoresByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/stores/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
