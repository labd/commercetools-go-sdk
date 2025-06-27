package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestBuilder struct {
	projectKey  string
	storeKey    string
	orderNumber string
	client      *Client
}

/**
*	Retrieves an Order with the provided `orderNumber` in a [Store](ctp:api:type:Store).
*	If the Order exists in the Project but does not have a `store` specified, or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders/order-number=%s", rb.projectKey, rb.storeKey, rb.orderNumber),
		client: rb.client,
	}
}

/**
*	Checks if an Order exists with the provided `orderNumber` in a [Store](ctp:api:type:Store). Returns a `200 OK` status if the Order exists or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders/order-number=%s", rb.projectKey, rb.storeKey, rb.orderNumber),
		client: rb.client,
	}
}

/**
*	Updates an Order in a [Store](ctp:api:type:Store) using one or more [update actions](/../api/projects/orders#update-actions).
*	If the Order exists in the Project but does not have a `store` specified, or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestBuilder) Post(body OrderUpdate) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders/order-number=%s", rb.projectKey, rb.storeKey, rb.orderNumber),
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
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders/order-number=%s", rb.projectKey, rb.storeKey, rb.orderNumber),
		client: rb.client,
	}
}
