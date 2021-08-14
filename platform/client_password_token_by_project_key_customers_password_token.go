// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCustomersPasswordTokenRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	The token value is used to reset the password of the customer with the given email. The token is
*	valid only for 10 minutes.
*
 */
func (rb *ByProjectKeyCustomersPasswordTokenRequestBuilder) Post(body CustomerCreatePasswordResetToken) *ByProjectKeyCustomersPasswordTokenRequestMethodPost {
	return &ByProjectKeyCustomersPasswordTokenRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/password-token", rb.projectKey),
		client: rb.client,
	}
}
