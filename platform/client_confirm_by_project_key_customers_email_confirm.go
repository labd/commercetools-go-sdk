// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCustomersEmailConfirmRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Verifies customer's email using a token.
 */
func (rb *ByProjectKeyCustomersEmailConfirmRequestBuilder) Post(body CustomerEmailVerify) *ByProjectKeyCustomersEmailConfirmRequestMethodPost {
	return &ByProjectKeyCustomersEmailConfirmRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/email/confirm", rb.projectKey),
		client: rb.client,
	}
}
