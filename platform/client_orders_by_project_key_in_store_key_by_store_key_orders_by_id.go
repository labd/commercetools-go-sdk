// Generated file, please do not change!!!
package platform

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
*	Returns an order by its ID from a specific Store. The {storeKey} path parameter maps to a Store's key.
*	If the order exists in the commercetools project but does not have the store field,
*	or the store field references a different store, this method returns a ResourceNotFound error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates an order in the store specified by {storeKey}. The {storeKey} path parameter maps to a Store's key.
*	If the order exists in the commercetools project but does not have the store field,
*	or the store field references a different store, this method returns a ResourceNotFound error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestBuilder) Post(body OrderUpdate) *ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
