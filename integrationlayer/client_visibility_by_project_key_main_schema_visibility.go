package integrationlayer

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMainSchemaVisibilityRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMainSchemaVisibilityRequestBuilder) Get() *ByProjectKeyMainSchemaVisibilityRequestMethodGet {
	return &ByProjectKeyMainSchemaVisibilityRequestMethodGet{
		url:    fmt.Sprintf("/%s/main/schema/visibility", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMainSchemaVisibilityRequestBuilder) Put(body []SourceCoordinate) *ByProjectKeyMainSchemaVisibilityRequestMethodPut {
	return &ByProjectKeyMainSchemaVisibilityRequestMethodPut{
		body:   body,
		url:    fmt.Sprintf("/%s/main/schema/visibility", rb.projectKey),
		client: rb.client,
	}
}
