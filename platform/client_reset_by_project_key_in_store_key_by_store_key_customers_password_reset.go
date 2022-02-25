package platform

// Generated file, please do not change!!!

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
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestBuilder) Post(body MyCustomerResetPassword) *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/password/reset", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
