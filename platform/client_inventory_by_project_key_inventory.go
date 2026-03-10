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
*	Checks if one or more InventoryEntries exist for the provided query predicate. Returns a `200` status if any Inventory Entries match the Query Predicate, or a `404` status otherwise.
 */
func (rb *ByProjectKeyInventoryRequestBuilder) Head() *ByProjectKeyInventoryRequestMethodHead {
	return &ByProjectKeyInventoryRequestMethodHead{
		url:    fmt.Sprintf("/%s/inventory", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates an InventoryEntry in the Project.
*
*	If quantity limits are provided, existing Line Items that reference a Product Variant with an SKU that matches the Inventory Entry can be affected. For more information, see [Quantity limits](/../api/carts-orders-overview#quantity-limits).
*
*	Produces the [InventoryEntryCreated](ctp:api:type:InventoryEntryCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyInventoryRequestBuilder) Post(body InventoryEntryDraft) *ByProjectKeyInventoryRequestMethodPost {
	return &ByProjectKeyInventoryRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/inventory", rb.projectKey),
		client: rb.client,
	}
}
