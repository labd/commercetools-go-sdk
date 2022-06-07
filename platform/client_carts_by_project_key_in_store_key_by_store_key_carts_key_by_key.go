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
*	Returns a cart by its key from a specific Store.
*	If the cart exists in the project but does not have the store field,
*	or the store field references a different store, this method returns a ResourceNotFound error.
*	The cart may not contain up-to-date prices, discounts etc.
*	If you want to ensure they're up-to-date, send an Update request with the Recalculate update action instead.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a cart in the store specified by {storeKey}.
*	If the cart exists in the project but does not have the store field,
*	or the store field references a different store, this method returns a ResourceNotFound error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestBuilder) Post(body CartUpdate) *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
