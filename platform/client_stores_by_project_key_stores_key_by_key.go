package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyStoresKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyStoresKeyByKeyRequestBuilder) Get() *ByProjectKeyStoresKeyByKeyRequestMethodGet {
	return &ByProjectKeyStoresKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/stores/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStoresKeyByKeyRequestBuilder) Post(body StoreUpdate) *ByProjectKeyStoresKeyByKeyRequestMethodPost {
	return &ByProjectKeyStoresKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/stores/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStoresKeyByKeyRequestBuilder) Delete() *ByProjectKeyStoresKeyByKeyRequestMethodDelete {
	return &ByProjectKeyStoresKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/stores/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
