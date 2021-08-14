// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCustomersPasswordResetRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Set a new password using a token.
 */
func (rb *ByProjectKeyCustomersPasswordResetRequestBuilder) Post(body CustomerResetPassword) *ByProjectKeyCustomersPasswordResetRequestMethodPost {
	return &ByProjectKeyCustomersPasswordResetRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/password/reset", rb.projectKey),
		client: rb.client,
	}
}
