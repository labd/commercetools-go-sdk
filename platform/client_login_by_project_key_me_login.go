// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyMeLoginRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeLoginRequestBuilder) Post() *ByProjectKeyMeLoginRequestMethodPost {
	return &ByProjectKeyMeLoginRequestMethodPost{
		url:    fmt.Sprintf("/%s/me/login", rb.projectKey),
		client: rb.client,
	}
}
