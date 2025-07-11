package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersPasswordResetRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Use this method to reset a global Customer's password during their [password reset process](/../api/customers-overview#customer-password-reset).
*
*	Resetting the password of the Customer produces the [CustomerPasswordUpdated](ctp:api:type:CustomerPasswordUpdatedMessage) Message with `reset=true`.
*
*	After the password is reset, all password tokens issued previously through the [password reset flow](/../api/projects/customers#password-reset-of-customer) are invalidated. In addition, any access and refresh tokens issued previously through the [password flow](/../api/authorization#password-flow) and [refresh token flow](/../api/authorization#refresh-token-flow) are invalidated. This invalidation of tokens is [eventually consistent](/../api/general-concepts#eventual-consistency).
*
 */
func (rb *ByProjectKeyCustomersPasswordResetRequestBuilder) Post(body CustomerResetPassword) *ByProjectKeyCustomersPasswordResetRequestMethodPost {
	return &ByProjectKeyCustomersPasswordResetRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/password/reset", rb.projectKey),
		client: rb.client,
	}
}
