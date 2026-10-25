package integrationlayer

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMainExtensionConfigRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMainExtensionConfigRequestBuilder) Get() *ByProjectKeyMainExtensionConfigRequestMethodGet {
	return &ByProjectKeyMainExtensionConfigRequestMethodGet{
		url:    fmt.Sprintf("/%s/main/extension/config", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMainExtensionConfigRequestBuilder) Patch(body ExtensionConfigPatch) *ByProjectKeyMainExtensionConfigRequestMethodPatch {
	return &ByProjectKeyMainExtensionConfigRequestMethodPatch{
		body:   body,
		url:    fmt.Sprintf("/%s/main/extension/config", rb.projectKey),
		client: rb.client,
	}
}
