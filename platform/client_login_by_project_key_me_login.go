// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyMeLoginRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeLoginRequestBuilder) Post(body CustomerSignin) *ByProjectKeyMeLoginRequestMethodPost {
	return &ByProjectKeyMeLoginRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/login", rb.projectKey),
		client: rb.client,
	}
}
