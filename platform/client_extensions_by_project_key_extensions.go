// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyExtensionsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyExtensionsRequestBuilder) WithKey(key string) *ByProjectKeyExtensionsKeyByKeyRequestBuilder {
	return &ByProjectKeyExtensionsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyExtensionsRequestBuilder) WithId(id string) *ByProjectKeyExtensionsByIdRequestBuilder {
	return &ByProjectKeyExtensionsByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Query extensions
 */
func (rb *ByProjectKeyExtensionsRequestBuilder) Get() *ByProjectKeyExtensionsRequestMethodGet {
	return &ByProjectKeyExtensionsRequestMethodGet{
		url:    fmt.Sprintf("/%s/extensions", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Currently, a maximum of 25 extensions can be created per project.
 */
func (rb *ByProjectKeyExtensionsRequestBuilder) Post(body ExtensionDraft) *ByProjectKeyExtensionsRequestMethodPost {
	return &ByProjectKeyExtensionsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/extensions", rb.projectKey),
		client: rb.client,
	}
}
