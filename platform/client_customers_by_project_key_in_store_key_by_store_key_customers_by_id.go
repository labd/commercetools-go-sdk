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
*	Checks if a Customer exists with the provided `id`. Returns a `200 OK` status if the Customer exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	If the Customer exists in the Project but the `stores` field references a different [Store](ctp:api:type:Store), this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
*	Simultaneously updating two Customers with the same email address can return a [LockedField](ctp:api:type:LockedFieldError) error.
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
