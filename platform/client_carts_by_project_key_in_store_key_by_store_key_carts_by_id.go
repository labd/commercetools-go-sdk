package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	If the Cart exists in the Project but does not have the `store` field, or the `store` field references a different Store, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
*	To ensure the Cart is up-to-date with current values (such as Prices and Discounts), use the [Recalculate](ctp:api:type:CartRecalculateAction) update action.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a [Cart](ctp:api:type:Cart) in the Store specified by `storeKey`.
*	If the Cart exists in the Project but does not have the `store` field, or the `store` field references a different Store, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestBuilder) Post(body CartUpdate) *ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	If the Cart exists in the Project but does not have the `store` field, or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
