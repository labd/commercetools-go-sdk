// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdRequestBuilder struct {
	projectKey string
	storeKey   string
	customerId string
	client     *Client
}

/**
*	Retrieves the active cart of the customer that has been modified most recently in a specific Store.
*	The {storeKey} path parameter maps to a Store's key.
*
*	If the cart exists in the commercetools project but does not have the store field, or the store field
*	references a different store, this method returns a ResourceNotFound error.
*
*	The cart may not contain up-to-date prices, discounts etc. If you want to ensure they're up-to-date,
*	send an Update request with the Recalculate update action instead.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/customer-id=%s", rb.projectKey, rb.storeKey, rb.customerId),
		client: rb.client,
	}
}
