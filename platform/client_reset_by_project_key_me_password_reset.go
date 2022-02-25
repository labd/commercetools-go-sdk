package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMePasswordResetRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMePasswordResetRequestBuilder) Post(body MyCustomerResetPassword) *ByProjectKeyMePasswordResetRequestMethodPost {
	return &ByProjectKeyMePasswordResetRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/password/reset", rb.projectKey),
		client: rb.client,
	}
}
