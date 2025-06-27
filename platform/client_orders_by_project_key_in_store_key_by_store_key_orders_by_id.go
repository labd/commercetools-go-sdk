package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Retrieves an Order with the provided `id` in a [Store](ctp:api:type:Store).
*	If the Order exists in the Project but does not have a `store` specified, or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if an Order exists with the provided `id` in a [Store](ctp:api:type:Store). Returns a `200 OK` status if the Order exists or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates an Order in a [Store](ctp:api:type:Store) using one or more [update actions](/../api/projects/orders#update-actions).
*	If the Order exists in the Project but does not have a `store` specified, or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestBuilder) Post(body OrderUpdate) *ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes an Order in a [Store](ctp:api:type:Store).
*	If the Order exists in the Project but does not have a `store` specified, or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
*	Deleting an Order produces the [OrderDeleted](ctp:api:type:OrderDeletedMessage) Message.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
