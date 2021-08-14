// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInventoryByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get InventoryEntry by ID
 */
func (rb *ByProjectKeyInventoryByIdRequestBuilder) Get() *ByProjectKeyInventoryByIdRequestMethodGet {
	return &ByProjectKeyInventoryByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/inventory/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update InventoryEntry by ID
 */
func (rb *ByProjectKeyInventoryByIdRequestBuilder) Post(body InventoryEntryUpdate) *ByProjectKeyInventoryByIdRequestMethodPost {
	return &ByProjectKeyInventoryByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/inventory/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete InventoryEntry by ID
 */
func (rb *ByProjectKeyInventoryByIdRequestBuilder) Delete() *ByProjectKeyInventoryByIdRequestMethodDelete {
	return &ByProjectKeyInventoryByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/inventory/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
