package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInventoryKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyInventoryKeyByKeyRequestBuilder) Get() *ByProjectKeyInventoryKeyByKeyRequestMethodGet {
	return &ByProjectKeyInventoryKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/inventory/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if an InventoryEntry exists with the provided `key`. Returns a `200 OK` status if the InventoryEntry exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInventoryKeyByKeyRequestBuilder) Head() *ByProjectKeyInventoryKeyByKeyRequestMethodHead {
	return &ByProjectKeyInventoryKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/inventory/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInventoryKeyByKeyRequestBuilder) Post(body InventoryEntryUpdate) *ByProjectKeyInventoryKeyByKeyRequestMethodPost {
	return &ByProjectKeyInventoryKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/inventory/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Produces the [InventoryEntryDeleted](ctp:api:type:InventoryEntryDeletedMessage) Message.
 */
func (rb *ByProjectKeyInventoryKeyByKeyRequestBuilder) Delete() *ByProjectKeyInventoryKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInventoryKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/inventory/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
