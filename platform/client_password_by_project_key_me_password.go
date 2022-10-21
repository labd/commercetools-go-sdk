package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMePasswordRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMePasswordRequestBuilder) Reset() *ByProjectKeyMePasswordResetRequestBuilder {
	return &ByProjectKeyMePasswordResetRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Changing the password of the Customer produces the [CustomerPasswordUpdated](ctp:api:type:CustomerPasswordUpdatedMessage) Message with `reset=false`.
*
*	If the current password does not match, an [InvalidCurrentPassword](ctp:api:type:InvalidCurrentPasswordError) error is returned.
*
 */
func (rb *ByProjectKeyMePasswordRequestBuilder) Post(body MyCustomerChangePassword) *ByProjectKeyMePasswordRequestMethodPost {
	return &ByProjectKeyMePasswordRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/password", rb.projectKey),
		client: rb.client,
	}
}
