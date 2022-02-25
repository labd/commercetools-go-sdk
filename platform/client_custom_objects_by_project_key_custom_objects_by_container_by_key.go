package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomObjectsByContainerByKeyRequestBuilder struct {
	projectKey string
	container  string
	key        string
	client     *Client
}

/**
*	Get CustomObject by container and key
 */
func (rb *ByProjectKeyCustomObjectsByContainerByKeyRequestBuilder) Get() *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGet {
	return &ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/custom-objects/%s/%s", rb.projectKey, rb.container, rb.key),
		client: rb.client,
	}
}

/**
*	Delete CustomObject by container and key
 */
func (rb *ByProjectKeyCustomObjectsByContainerByKeyRequestBuilder) Delete() *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete {
	return &ByProjectKeyCustomObjectsByContainerByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/custom-objects/%s/%s", rb.projectKey, rb.container, rb.key),
		client: rb.client,
	}
}
