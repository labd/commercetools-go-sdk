package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersPasswordResetRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Resetting the password of the Customer produces the [CustomerPasswordUpdated](ctp:api:type:CustomerPasswordUpdatedMessage) Message with `reset=true`.
*
 */
func (rb *ByProjectKeyCustomersPasswordResetRequestBuilder) Post(body CustomerResetPassword) *ByProjectKeyCustomersPasswordResetRequestMethodPost {
	return &ByProjectKeyCustomersPasswordResetRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/password/reset", rb.projectKey),
		client: rb.client,
	}
}
