package platform

// Generated file, please do not change!!!

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
func (rb *ByProjectKeyStoresRequestBuilder) WithId(id string) *ByProjectKeyStoresByIDRequestBuilder {
	return &ByProjectKeyStoresByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyStoresRequestBuilder) Get() *ByProjectKeyStoresRequestMethodGet {
	return &ByProjectKeyStoresRequestMethodGet{
		url:    fmt.Sprintf("/%s/stores", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Stores exist for the provided query predicate. Returns a `200 OK` status if any Stores match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyStoresRequestBuilder) Head() *ByProjectKeyStoresRequestMethodHead {
	return &ByProjectKeyStoresRequestMethodHead{
		url:    fmt.Sprintf("/%s/stores", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStoresRequestBuilder) Post(body StoreDraft) *ByProjectKeyStoresRequestMethodPost {
	return &ByProjectKeyStoresRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/stores", rb.projectKey),
		client: rb.client,
	}
}
