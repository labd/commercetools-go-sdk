package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersEmailConfirmRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Verifying the email of the Customer produces the [CustomerEmailVerified](ctp:api:type:CustomerEmailVerifiedMessage) Message.
*
 */
func (rb *ByProjectKeyCustomersEmailConfirmRequestBuilder) Post(body CustomerEmailVerify) *ByProjectKeyCustomersEmailConfirmRequestMethodPost {
	return &ByProjectKeyCustomersEmailConfirmRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/email/confirm", rb.projectKey),
		client: rb.client,
	}
}
