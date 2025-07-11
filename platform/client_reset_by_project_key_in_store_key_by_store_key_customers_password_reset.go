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
*	Use this method to reset a Store-specific Customer's password during their [password reset process](/../api/customers-overview#customer-password-reset).
*
*	Resetting the password of the Customer produces the [CustomerPasswordUpdated](ctp:api:type:CustomerPasswordUpdatedMessage) Message with `reset=true`.
*
*	After the password is reset, all password tokens issued previously through the [password reset flow](/../api/projects/customers#password-reset-of-customer) are invalidated. In addition, any access and refresh tokens issued previously through the [password flow](/../api/authorization#password-flow) and [refresh token flow](/../api/authorization#refresh-token-flow) are invalidated. This invalidation of tokens is [eventually consistent](/../api/general-concepts#eventual-consistency).
*
*	If the Customer exists in the Project but the `stores` field references a different [Store](ctp:api:type:Store), then this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestBuilder) Post(body CustomerResetPassword) *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/password/reset", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
