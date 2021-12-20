// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Set a new password using a token.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestBuilder) Post(body CustomerResetPassword) *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/password/reset", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
