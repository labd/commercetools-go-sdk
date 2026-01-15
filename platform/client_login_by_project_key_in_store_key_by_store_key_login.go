package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyLoginRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Authenticates a Customer associated with a [Store](ctp:api:type:Store).
*
*	Allows [merging](/../api/customers-overview#cart-merge-during-sign-in-and-sign-up) items from an anonymous Cart into the most recently modified active Cart of a Customer.
*	If no active Cart exists, the anonymous Cart becomes the Customer's active Cart.
*	If the Customer has multiple active Carts, the anonymous Cart is merged into the most recently modified active Cart.
*
*	If the Customer exists in the Project but the `stores` field references a different [Store](ctp:api:type:Store), this method returns an [InvalidCredentials](ctp:api:type:InvalidCredentialsError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyLoginRequestBuilder) Post(body CustomerSignin) *ByProjectKeyInStoreKeyByStoreKeyLoginRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyLoginRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/login", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
