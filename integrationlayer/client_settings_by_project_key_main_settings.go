package integrationlayer

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMainSettingsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMainSettingsRequestBuilder) Get() *ByProjectKeyMainSettingsRequestMethodGet {
	return &ByProjectKeyMainSettingsRequestMethodGet{
		url:    fmt.Sprintf("/%s/main/settings", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMainSettingsRequestBuilder) Put(body ProjectSettings) *ByProjectKeyMainSettingsRequestMethodPut {
	return &ByProjectKeyMainSettingsRequestMethodPut{
		body:   body,
		url:    fmt.Sprintf("/%s/main/settings", rb.projectKey),
		client: rb.client,
	}
}
