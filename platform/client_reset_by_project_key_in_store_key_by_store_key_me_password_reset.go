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
*	Resetting a password produces the Customer [CustomerPasswordUpdated](ctp:api:type:CustomerPasswordUpdatedMessage) Message with `reset=true`.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Customer exists with the `id` specified in the [customer:{id}](/scopes#customer_idid) scope.
*	- If the Customer exists but is associated with a different Store than what is specified in the `manage_my_profile:{projectKey}:{storeKey}` scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMePasswordResetRequestBuilder) Post(body MyCustomerResetPassword) *ByProjectKeyInStoreKeyByStoreKeyMePasswordResetRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMePasswordResetRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/password/reset", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
