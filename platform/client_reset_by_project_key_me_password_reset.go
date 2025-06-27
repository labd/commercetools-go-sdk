package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMePasswordResetRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	This is the last step in the [password reset process of a Customer](/../api/projects/customers#password-reset-of-customer).
*
*	Resetting a password of the Customer produces the [CustomerPasswordUpdated](ctp:api:type:CustomerPasswordUpdatedMessage) Message with `reset=true`.
*
*	After the password is reset, all password tokens issued previously through the [password reset flow](/../api/projects/customers#password-reset-of-customer) are invalidated. In addition, any access and refresh tokens issued previously through the [password flow](/../api/authorization#password-flow) and [refresh token flow](/../api/authorization#refresh-token-flow) are invalidated. This invalidation of tokens is [eventually consistent](/../api/general-concepts#eventual-consistency).
*
 */
func (rb *ByProjectKeyMePasswordResetRequestBuilder) Post(body MyCustomerResetPassword) *ByProjectKeyMePasswordResetRequestMethodPost {
	return &ByProjectKeyMePasswordResetRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/password/reset", rb.projectKey),
		client: rb.client,
	}
}
