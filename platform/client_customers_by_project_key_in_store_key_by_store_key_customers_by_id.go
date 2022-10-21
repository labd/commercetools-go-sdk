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
*	If the Customer exists in the Project but the `stores` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	If the Customer exists in the Project but the `stores` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestBuilder) Post(body CustomerUpdate) *ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deleting a Customer produces the [CustomerDeleted](ctp:api:type:CustomerDeletedMessage) Message.
*
*	If the Customer exists in the Project but the `stores` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
