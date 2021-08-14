// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyMeEmailConfirmRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeEmailConfirmRequestBuilder) Post() *ByProjectKeyMeEmailConfirmRequestMethodPost {
	return &ByProjectKeyMeEmailConfirmRequestMethodPost{
		url:    fmt.Sprintf("/%s/me/email/confirm", rb.projectKey),
		client: rb.client,
	}
}
