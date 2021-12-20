// Generated file, please do not change!!!
package history

import (
	"fmt"
)

type ByProjectKeyByResourceTypeByIDRequestBuilder struct {
	projectKey   string
	resourceType string
	id           string
	client       *Client
}

func (rb *ByProjectKeyByResourceTypeByIDRequestBuilder) Get() *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	return &ByProjectKeyByResourceTypeByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/%s/%s", rb.projectKey, rb.resourceType, rb.id),
		client: rb.client,
	}
}
