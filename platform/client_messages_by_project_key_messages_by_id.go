// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyMessagesByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get Message by ID
 */
func (rb *ByProjectKeyMessagesByIdRequestBuilder) Get() *ByProjectKeyMessagesByIdRequestMethodGet {
	return &ByProjectKeyMessagesByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/messages/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
