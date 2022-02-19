// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestBuilder struct {
	projectKey string
	storeKey   string
	key        string
	client     *Client
}

/**
*	Returns a customer by its Key from a specific Store.
*	It also considers customers that do not have the stores field.
*	If the customer exists in the commercetools project but the stores field references different stores,
*	this method returns a ResourceNotFound error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	If the customer exists in the commercetools project but the stores field references a different store,
*	this method returns a ResourceNotFound error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestBuilder) Post(body CustomerUpdate) *ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
