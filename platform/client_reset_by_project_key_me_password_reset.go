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
 */
func (rb *ByProjectKeyMePasswordResetRequestBuilder) Post(body MyCustomerResetPassword) *ByProjectKeyMePasswordResetRequestMethodPost {
	return &ByProjectKeyMePasswordResetRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/password/reset", rb.projectKey),
		client: rb.client,
	}
}
