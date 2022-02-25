package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeSignupRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeSignupRequestBuilder) Post(body MyCustomerDraft) *ByProjectKeyMeSignupRequestMethodPost {
	return &ByProjectKeyMeSignupRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/signup", rb.projectKey),
		client: rb.client,
	}
}
