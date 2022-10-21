package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMessagesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Deprecated scope: `view_orders:{projectKey}`
 */
func (rb *ByProjectKeyMessagesByIDRequestBuilder) Get() *ByProjectKeyMessagesByIDRequestMethodGet {
	return &ByProjectKeyMessagesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/messages/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
