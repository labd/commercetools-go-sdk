// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCustomObjectsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCustomObjectsRequestBuilder) WithContainerAndKey(container string, key string) *ByProjectKeyCustomObjectsByContainerByKeyRequestBuilder {
	return &ByProjectKeyCustomObjectsByContainerByKeyRequestBuilder{
		container:  container,
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomObjectsRequestBuilder) WithContainer(container string) *ByProjectKeyCustomObjectsByContainerRequestBuilder {
	return &ByProjectKeyCustomObjectsByContainerRequestBuilder{
		container:  container,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	The query endpoint allows to retrieve custom objects in a specific container or all custom objects.
*	For performance reasons, it is highly advisable to query only for custom objects in a container by using
*	the container field in the where predicate.
*
 */
func (rb *ByProjectKeyCustomObjectsRequestBuilder) Get() *ByProjectKeyCustomObjectsRequestMethodGet {
	return &ByProjectKeyCustomObjectsRequestMethodGet{
		url:    fmt.Sprintf("/%s/custom-objects", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a new custom object or updates an existing custom object.
*	If an object with the given container/key exists,
*	the object will be replaced with the new value and the version is incremented.
*	If the request contains a version and an object with the given container/key exists then the version
*	must match the version of the existing object. Concurrent updates for the same custom object still can result
*	in a Conflict (409) even if the version is not provided.
*	Fields with null values will not be saved.
*
 */
func (rb *ByProjectKeyCustomObjectsRequestBuilder) Post(body CustomObjectDraft) *ByProjectKeyCustomObjectsRequestMethodPost {
	return &ByProjectKeyCustomObjectsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/custom-objects", rb.projectKey),
		client: rb.client,
	}
}
