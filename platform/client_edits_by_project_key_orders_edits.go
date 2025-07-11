package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersEditsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyOrdersEditsRequestBuilder) WithKey(key string) *ByProjectKeyOrdersEditsKeyByKeyRequestBuilder {
	return &ByProjectKeyOrdersEditsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyOrdersEditsRequestBuilder) WithId(id string) *ByProjectKeyOrdersEditsByIDRequestBuilder {
	return &ByProjectKeyOrdersEditsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves OrderEdits in the Project.
 */
func (rb *ByProjectKeyOrdersEditsRequestBuilder) Get() *ByProjectKeyOrdersEditsRequestMethodGet {
	return &ByProjectKeyOrdersEditsRequestMethodGet{
		url:    fmt.Sprintf("/%s/orders/edits", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more OrderEdits exist for the provided query predicate. Returns a `200 OK` status if any OrderEdits match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyOrdersEditsRequestBuilder) Head() *ByProjectKeyOrdersEditsRequestMethodHead {
	return &ByProjectKeyOrdersEditsRequestMethodHead{
		url:    fmt.Sprintf("/%s/orders/edits", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates an OrderEdit in the Project.
*	You can either create multiple Order Edits for an Order and apply them sequentially to an Order, or create multiple Order Edits parallelly (as alternatives to each other) and apply one of them to the Order.
*
*	You can only create an Order Edit if the [InventoryMode](/projects/carts#inventorymode) of the Order and its [LineItems](/projects/carts#lineitem) is `None`.
*
 */
func (rb *ByProjectKeyOrdersEditsRequestBuilder) Post(body OrderEditDraft) *ByProjectKeyOrdersEditsRequestMethodPost {
	return &ByProjectKeyOrdersEditsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/edits", rb.projectKey),
		client: rb.client,
	}
}
