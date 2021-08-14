// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCustomObjectsByContainerRequestBuilder struct {
	projectKey string
	container  string
	client     *Client
}

func (rb *ByProjectKeyCustomObjectsByContainerRequestBuilder) Get() *ByProjectKeyCustomObjectsByContainerRequestMethodGet {
	return &ByProjectKeyCustomObjectsByContainerRequestMethodGet{
		url:    fmt.Sprintf("/%s/custom-objects/%s", rb.projectKey, rb.container),
		client: rb.client,
	}
}
