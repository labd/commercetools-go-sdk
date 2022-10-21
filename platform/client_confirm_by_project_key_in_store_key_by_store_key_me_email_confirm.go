package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeEmailConfirmRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	This is the last step in the [email verification process of a Customer](/../api/projects/customers#email-verification-of-customer-in-store).
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeEmailConfirmRequestBuilder) Post(body MyCustomerEmailVerify) *ByProjectKeyInStoreKeyByStoreKeyMeEmailConfirmRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeEmailConfirmRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/email/confirm", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
