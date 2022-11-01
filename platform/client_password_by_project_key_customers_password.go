package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersPasswordRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Changing the password produces the [CustomerPasswordUpdated](ctp:api:type:CustomerPasswordUpdatedMessage) Message with `reset=false`.
*
 */
func (rb *ByProjectKeyCustomersPasswordRequestBuilder) Post(body CustomerChangePassword) *ByProjectKeyCustomersPasswordRequestMethodPost {
	return &ByProjectKeyCustomersPasswordRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/password", rb.projectKey),
		client: rb.client,
	}
}
