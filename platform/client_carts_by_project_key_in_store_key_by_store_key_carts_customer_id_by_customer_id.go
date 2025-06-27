package platform

// Generated file, please do not change!!!

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
*	Retrieves the most recently modified [active Cart](ctp:api:type:CartState) of a Customer with [CartOrigin](ctp:api:type:CartOrigin) `Customer`. If no active Cart exists, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
*	If the Cart exists in the Project but does not have a `store` specified, or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
*	To ensure the Cart is up-to-date with current values (such as Prices and Discounts), use the [Recalculate](ctp:api:type:CartRecalculateAction) update action.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/customer-id=%s", rb.projectKey, rb.storeKey, rb.customerId),
		client: rb.client,
	}
}

/**
*	Checks if one or more Carts exist for a Customer in a [Store](ctp:api:type:Store). Returns a `200 OK` status if the Cart exists or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/customer-id=%s", rb.projectKey, rb.storeKey, rb.customerId),
		client: rb.client,
	}
}
