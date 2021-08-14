// Generated file, please do not change!!!
package history

import (
	"fmt"
)

type ByProjectKeyRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyRequestBuilder) WithResourceTypeValue(resourceType string) *ByProjectKeyByResourceTypeRequestBuilder {
	return &ByProjectKeyByResourceTypeRequestBuilder{
		resourceType: resourceType,
		projectKey:   rb.projectKey,
		client:       rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) Get() *ByProjectKeyRequestMethodGet {
	return &ByProjectKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s", rb.projectKey),
		client: rb.client,
	}
}
