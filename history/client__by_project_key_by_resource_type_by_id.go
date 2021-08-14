// Generated file, please do not change!!!
package history

import (
	"fmt"
)

type ByProjectKeyByResourceTypeByIdRequestBuilder struct {
	projectKey   string
	resourceType string
	id           string
	client       *Client
}

func (rb *ByProjectKeyByResourceTypeByIdRequestBuilder) Get() *ByProjectKeyByResourceTypeByIdRequestMethodGet {
	return &ByProjectKeyByResourceTypeByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/%s/%s", rb.projectKey, rb.resourceType, rb.id),
		client: rb.client,
	}
}
