// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyExtensionsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves the representation of an extension by its key.
 */
func (rb *ByProjectKeyExtensionsKeyByKeyRequestBuilder) Get() *ByProjectKeyExtensionsKeyByKeyRequestMethodGet {
	return &ByProjectKeyExtensionsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/extensions/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyExtensionsKeyByKeyRequestBuilder) Post(body ExtensionUpdate) *ByProjectKeyExtensionsKeyByKeyRequestMethodPost {
	return &ByProjectKeyExtensionsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/extensions/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyExtensionsKeyByKeyRequestBuilder) Delete() *ByProjectKeyExtensionsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyExtensionsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/extensions/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
