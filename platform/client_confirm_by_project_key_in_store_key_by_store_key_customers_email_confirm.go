// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersEmailConfirmRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Verifies customer's email using a token.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailConfirmRequestBuilder) Post(body CustomerEmailVerify) *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailConfirmRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersEmailConfirmRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/email/confirm", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
