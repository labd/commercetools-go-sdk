package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeEmailConfirmRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeEmailConfirmRequestBuilder) Post(body CustomerEmailVerify) *ByProjectKeyMeEmailConfirmRequestMethodPost {
	return &ByProjectKeyMeEmailConfirmRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/email/confirm", rb.projectKey),
		client: rb.client,
	}
}
