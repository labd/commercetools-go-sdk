// Generated file, please do not change!!!
package platform

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
func (rb *ByProjectKeyMePasswordRequestBuilder) Post(body MyCustomerChangePassword) *ByProjectKeyMePasswordRequestMethodPost {
	return &ByProjectKeyMePasswordRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/password", rb.projectKey),
		client: rb.client,
	}
}
