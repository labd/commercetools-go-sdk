package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestBuilder struct {
	projectKey string
	storeKey   string
	key        string
	client     *Client
}

/**
*
*	Retrieves a Cart with the provided `key` in a [Store](ctp:api:type:Store).
*
*	If the Cart exists in the Project but does not have a `store` specified, or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
*	To ensure the Cart is up-to-date with current values (such as Prices and Discounts), use the [Recalculate](ctp:api:type:CartRecalculateAction) update action.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Cart exists with the provided `key` in a [Store](ctp:api:type:Store). Returns a `200 OK` status if the Cart exists or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a Cart in a [Store](ctp:api:type:Store) using one or more [update actions](/../api/projects/carts#update-actions).
*
*	If the Cart exists in the Project but does not have a `store` specified, or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestBuilder) Post(body CartUpdate) *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a Cart in a [Store](ctp:api:type:Store).
*
*	If the Cart exists in the Project but does not have a `store` specified, or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
