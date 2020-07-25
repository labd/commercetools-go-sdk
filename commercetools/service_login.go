// Automatically generated, do not edit

package commercetools

import (
	"context"
	"net/url"
)

/*
Login Authenticate Customer (Sign In). Retrieves the authenticated
customer (a customer that matches the given email/password pair).
If used with an access token for Anonymous Sessions,
all orders and carts belonging to the anonymousId will be assigned to the newly created customer.
If a cart is is returned as part of the CustomerSignInResult,
it has been recalculated (It will have up-to-date prices, taxes and discounts,
and invalid line items have been removed.).
*/
func (client *Client) Login(ctx context.Context, value *CustomerSignin, opts ...RequestOption) (result *CustomerSignInResult, err error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(&params)
	}

	endpoint := "login"
	err = client.create(ctx, endpoint, params, value, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
