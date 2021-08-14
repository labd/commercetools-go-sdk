// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Returns a customer by its ID from a specific Store. The {storeKey} path parameter maps to a Store's key.
*	It also considers customers that do not have the stores field.
*	If the customer exists in the commercetools project but the stores field references different stores,
*	this method returns a ResourceNotFound error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a customer in the store specified by {storeKey}. The {storeKey} path parameter maps to a Store's key.
*	If the customer exists in the commercetools project but the stores field references a different store,
*	this method returns a ResourceNotFound error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestBuilder) Post(body CustomerUpdate) *ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete Customer by ID
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
