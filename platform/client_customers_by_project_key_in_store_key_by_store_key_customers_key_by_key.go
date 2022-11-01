package platform

// Generated file, please do not change!!!

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
*	If the Customer exists in the Project but the `stores` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	If the Customer exists in the Project but the `stores` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestBuilder) Post(body CustomerUpdate) *ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deleting a Customer produces the [CustomerDeleted](ctp:api:type:CustomerDeletedMessage) Message.
*
*	If the Customer exists in the Project but the `stores` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
