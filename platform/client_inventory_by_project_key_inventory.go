// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInventoryRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyInventoryRequestBuilder) WithId(id string) *ByProjectKeyInventoryByIdRequestBuilder {
	return &ByProjectKeyInventoryByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Query inventory
 */
func (rb *ByProjectKeyInventoryRequestBuilder) Get() *ByProjectKeyInventoryRequestMethodGet {
	return &ByProjectKeyInventoryRequestMethodGet{
		url:    fmt.Sprintf("/%s/inventory", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Create InventoryEntry
 */
func (rb *ByProjectKeyInventoryRequestBuilder) Post(body InventoryEntryDraft) *ByProjectKeyInventoryRequestMethodPost {
	return &ByProjectKeyInventoryRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/inventory", rb.projectKey),
		client: rb.client,
	}
}
