package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCartsCustomerIdByCustomerIdRequestBuilder struct {
	projectKey string
	customerId string
	client     *Client
}

/**
*	Retrieves the active cart of the customer that has been modified most recently.
*	It does not consider carts with CartOrigin Merchant. If no active cart exists, a 404 Not Found error is returned.
*
*	The cart may not contain up-to-date prices, discounts etc. If you want to ensure they're up-to-date,
*	send an Update request with the Recalculate update action instead.
*
 */
func (rb *ByProjectKeyCartsCustomerIdByCustomerIdRequestBuilder) Get() *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet {
	return &ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/carts/customer-id=%s", rb.projectKey, rb.customerId),
		client: rb.client,
	}
}
