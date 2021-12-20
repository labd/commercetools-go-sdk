// Generated file, please do not change!!!
package history

import (
	"fmt"
)

type ByProjectKeyByResourceTypeRequestBuilder struct {
	projectKey   string
	resourceType string
	client       *Client
}

func (rb *ByProjectKeyByResourceTypeRequestBuilder) WithIdValue(id string) *ByProjectKeyByResourceTypeByIDRequestBuilder {
	return &ByProjectKeyByResourceTypeByIDRequestBuilder{
		id:           id,
		projectKey:   rb.projectKey,
		resourceType: rb.resourceType,
		client:       rb.client,
	}
}
func (rb *ByProjectKeyByResourceTypeRequestBuilder) Get() *ByProjectKeyByResourceTypeRequestMethodGet {
	return &ByProjectKeyByResourceTypeRequestMethodGet{
		url:    fmt.Sprintf("/%s/%s", rb.projectKey, rb.resourceType),
		client: rb.client,
	}
}
