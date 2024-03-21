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
*	Authenticates a Customer associated with a [Store](ctp:api:type:Store). For more information, see [Global versus Store-specific Customers](/../api/customers-overview#global-versus-store-specific-customers).
*
*	Triggers [Cart merge during sign-in](/../api/customers-overview#cart-merge-during-sign-in).
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
