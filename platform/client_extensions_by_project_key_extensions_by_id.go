// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyExtensionsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves the representation of an extension by its id.
 */
func (rb *ByProjectKeyExtensionsByIdRequestBuilder) Get() *ByProjectKeyExtensionsByIdRequestMethodGet {
	return &ByProjectKeyExtensionsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/extensions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update Extension by ID
 */
func (rb *ByProjectKeyExtensionsByIdRequestBuilder) Post(body ExtensionUpdate) *ByProjectKeyExtensionsByIdRequestMethodPost {
	return &ByProjectKeyExtensionsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/extensions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete Extension by ID
 */
func (rb *ByProjectKeyExtensionsByIdRequestBuilder) Delete() *ByProjectKeyExtensionsByIdRequestMethodDelete {
	return &ByProjectKeyExtensionsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/extensions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
