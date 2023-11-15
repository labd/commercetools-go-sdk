package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInventoryByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyInventoryByIDRequestBuilder) Get() *ByProjectKeyInventoryByIDRequestMethodGet {
	return &ByProjectKeyInventoryByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/inventory/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if an InventoryEntry exists for a given `id`. Returns a `200 OK` status if the InventoryEntry exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInventoryByIDRequestBuilder) Head() *ByProjectKeyInventoryByIDRequestMethodHead {
	return &ByProjectKeyInventoryByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/inventory/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInventoryByIDRequestBuilder) Post(body InventoryEntryUpdate) *ByProjectKeyInventoryByIDRequestMethodPost {
	return &ByProjectKeyInventoryByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/inventory/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Produces the [InventoryEntryDeleted](ctp:api:type:InventoryEntryDeletedMessage) Message.
 */
func (rb *ByProjectKeyInventoryByIDRequestBuilder) Delete() *ByProjectKeyInventoryByIDRequestMethodDelete {
	return &ByProjectKeyInventoryByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/inventory/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
