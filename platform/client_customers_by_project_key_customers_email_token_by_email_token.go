// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCustomersEmailTokenByEmailTokenRequestBuilder struct {
	projectKey string
	emailToken string
	client     *Client
}

/**
*	Get Customer by emailToken
 */
func (rb *ByProjectKeyCustomersEmailTokenByEmailTokenRequestBuilder) Get() *ByProjectKeyCustomersEmailTokenByEmailTokenRequestMethodGet {
	return &ByProjectKeyCustomersEmailTokenByEmailTokenRequestMethodGet{
		url:    fmt.Sprintf("/%s/customers/email-token=%s", rb.projectKey, rb.emailToken),
		client: rb.client,
	}
}
