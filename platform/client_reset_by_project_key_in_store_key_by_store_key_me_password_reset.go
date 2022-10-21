package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMePasswordResetRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	This is the last step in the [password reset process of the authenticated Customer](/../api/projects/customers#password-reset-of-customer-in-store).
*
*	Resetting a password produces the of the Customer [CustomerPasswordUpdated](ctp:api:type:CustomerPasswordUpdatedMessage) Message with `reset=true`.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMePasswordResetRequestBuilder) Post(body MyCustomerResetPassword) *ByProjectKeyInStoreKeyByStoreKeyMePasswordResetRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMePasswordResetRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/password/reset", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
