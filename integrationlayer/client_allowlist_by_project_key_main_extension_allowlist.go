package integrationlayer

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMainExtensionAllowlistRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMainExtensionAllowlistRequestBuilder) Get() *ByProjectKeyMainExtensionAllowlistRequestMethodGet {
	return &ByProjectKeyMainExtensionAllowlistRequestMethodGet{
		url:    fmt.Sprintf("/%s/main/extension/allowlist", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMainExtensionAllowlistRequestBuilder) Put(body []string) *ByProjectKeyMainExtensionAllowlistRequestMethodPut {
	return &ByProjectKeyMainExtensionAllowlistRequestMethodPut{
		body:   body,
		url:    fmt.Sprintf("/%s/main/extension/allowlist", rb.projectKey),
		client: rb.client,
	}
}
