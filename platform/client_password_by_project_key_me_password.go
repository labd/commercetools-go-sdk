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
func (rb *ByProjectKeyMePasswordRequestBuilder) Post() *ByProjectKeyMePasswordRequestMethodPost {
	return &ByProjectKeyMePasswordRequestMethodPost{
		url:    fmt.Sprintf("/%s/me/password", rb.projectKey),
		client: rb.client,
	}
}
