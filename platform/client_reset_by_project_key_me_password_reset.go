// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyMePasswordResetRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMePasswordResetRequestBuilder) Post() *ByProjectKeyMePasswordResetRequestMethodPost {
	return &ByProjectKeyMePasswordResetRequestMethodPost{
		url:    fmt.Sprintf("/%s/me/password/reset", rb.projectKey),
		client: rb.client,
	}
}
