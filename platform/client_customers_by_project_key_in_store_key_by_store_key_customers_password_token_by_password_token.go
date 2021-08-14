// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestBuilder struct {
	projectKey    string
	storeKey      string
	passwordToken string
	client        *Client
}

/**
*	Get Customer by passwordToken
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/password-token=%s", rb.projectKey, rb.storeKey, rb.passwordToken),
		client: rb.client,
	}
}
