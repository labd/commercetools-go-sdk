package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInventoryRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyInventoryRequestBuilder) WithId(id string) *ByProjectKeyInventoryByIDRequestBuilder {
	return &ByProjectKeyInventoryByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInventoryRequestBuilder) WithKey(key string) *ByProjectKeyInventoryKeyByKeyRequestBuilder {
	return &ByProjectKeyInventoryKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInventoryRequestBuilder) Get() *ByProjectKeyInventoryRequestMethodGet {
	return &ByProjectKeyInventoryRequestMethodGet{
		url:    fmt.Sprintf("/%s/inventory", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Produces the [InventoryEntryCreatedMessage](ctp:api:type:InventoryEntryCreatedMessage).
 */
func (rb *ByProjectKeyInventoryRequestBuilder) Post(body InventoryEntryDraft) *ByProjectKeyInventoryRequestMethodPost {
	return &ByProjectKeyInventoryRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/inventory", rb.projectKey),
		client: rb.client,
	}
}
