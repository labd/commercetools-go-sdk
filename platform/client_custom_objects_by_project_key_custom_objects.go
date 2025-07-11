package platform

// Generated file, please do not change!!!

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
*	For performance reasons, it is highly advisable to query for Custom Objects in a container by using the `container` field in the `where` predicate.
*
 */
func (rb *ByProjectKeyCustomObjectsRequestBuilder) Get() *ByProjectKeyCustomObjectsRequestMethodGet {
	return &ByProjectKeyCustomObjectsRequestMethodGet{
		url:    fmt.Sprintf("/%s/custom-objects", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more CustomObjects exist for the provided query predicate. Returns a `200 OK` status if any CustomObjects match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyCustomObjectsRequestBuilder) Head() *ByProjectKeyCustomObjectsRequestMethodHead {
	return &ByProjectKeyCustomObjectsRequestMethodHead{
		url:    fmt.Sprintf("/%s/custom-objects", rb.projectKey),
		client: rb.client,
	}
}

/**
*	If an object with the given container/key exists, the object will be replaced with the new value and the version is incremented.
*	If the request contains a version and an object with the given container/key, then the version must match the version of the existing object. Concurrent updates to the same Custom Object returns a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error even if the version is not provided.
*
*	Fields within `value` that have `null` values **are not saved**.
*
 */
func (rb *ByProjectKeyCustomObjectsRequestBuilder) Post(body CustomObjectDraft) *ByProjectKeyCustomObjectsRequestMethodPost {
	return &ByProjectKeyCustomObjectsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/custom-objects", rb.projectKey),
		client: rb.client,
	}
}
