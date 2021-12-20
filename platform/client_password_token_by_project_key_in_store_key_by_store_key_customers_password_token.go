// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	The token value is used to reset the password of the customer with the given email. The token is
*	valid only for 10 minutes.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenRequestBuilder) Post(body CustomerCreatePasswordResetToken) *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/password-token", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
