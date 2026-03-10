package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyLoginRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Authenticates a global Customer.
*
*	Allows [merging](/../api/customers-overview#cart-merge-during-sign-in-and-sign-up) items from an anonymous Cart into the most recently modified active Cart of a Customer.
*	If no active Cart exists, the anonymous Cart becomes the Customer's active Cart.
*	If the Customer has multiple active Carts, the anonymous Cart is merged into the most recently modified active Cart.
*
*	If an account with the given credentials is not found, an [InvalidCredentials](ctp:api:type:InvalidCredentialsError) error is returned.
*
 */
func (rb *ByProjectKeyLoginRequestBuilder) Post(body CustomerSignin) *ByProjectKeyLoginRequestMethodPost {
	return &ByProjectKeyLoginRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/login", rb.projectKey),
		client: rb.client,
	}
}
