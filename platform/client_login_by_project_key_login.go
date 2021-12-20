// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyLoginRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Authenticate Customer (Sign In). Retrieves the authenticated
*	customer (a customer that matches the given email/password pair).
*	If used with an access token for Anonymous Sessions,
*	all orders and carts belonging to the anonymousId will be assigned to the newly created customer.
*	If a cart is is returned as part of the CustomerSignInResult,
*	it has been recalculated (It will have up-to-date prices, taxes and discounts,
*	and invalid line items have been removed.).
*
 */
func (rb *ByProjectKeyLoginRequestBuilder) Post(body CustomerSignin) *ByProjectKeyLoginRequestMethodPost {
	return &ByProjectKeyLoginRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/login", rb.projectKey),
		client: rb.client,
	}
}
