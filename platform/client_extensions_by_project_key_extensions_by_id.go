package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyExtensionsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyExtensionsByIDRequestBuilder) Get() *ByProjectKeyExtensionsByIDRequestMethodGet {
	return &ByProjectKeyExtensionsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/extensions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyExtensionsByIDRequestBuilder) Post(body ExtensionUpdate) *ByProjectKeyExtensionsByIDRequestMethodPost {
	return &ByProjectKeyExtensionsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/extensions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyExtensionsByIDRequestBuilder) Delete() *ByProjectKeyExtensionsByIDRequestMethodDelete {
	return &ByProjectKeyExtensionsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/extensions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
