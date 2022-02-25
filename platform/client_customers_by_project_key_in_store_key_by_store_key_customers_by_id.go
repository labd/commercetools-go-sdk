package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Returns a customer by its ID from a specific Store.
*	It also considers customers that do not have the stores field.
*	If the customer exists in the commercetools project but the stores field references different stores,
*	this method returns a ResourceNotFound error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a customer in the store specified by {storeKey}.
*	If the customer exists in the commercetools project but the stores field references a different store,
*	this method returns a ResourceNotFound error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestBuilder) Post(body CustomerUpdate) *ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
