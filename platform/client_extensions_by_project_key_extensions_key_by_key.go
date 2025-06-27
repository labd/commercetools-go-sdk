package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyExtensionsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyExtensionsKeyByKeyRequestBuilder) Get() *ByProjectKeyExtensionsKeyByKeyRequestMethodGet {
	return &ByProjectKeyExtensionsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/extensions/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if an Extension exists with the provided `key`. Returns a `200 OK` status if the Extension exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyExtensionsKeyByKeyRequestBuilder) Head() *ByProjectKeyExtensionsKeyByKeyRequestMethodHead {
	return &ByProjectKeyExtensionsKeyByKeyRequestMethodHead{
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
