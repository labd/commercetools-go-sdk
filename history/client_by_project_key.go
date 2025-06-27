package history

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyRequestBuilder) Graphql() *ByProjectKeyGraphqlRequestBuilder {
	return &ByProjectKeyGraphqlRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) WithResourceTypeValue(resourceType string) *ByProjectKeyByResourceTypeRequestBuilder {
	return &ByProjectKeyByResourceTypeRequestBuilder{
		resourceType: resourceType,
		projectKey:   rb.projectKey,
		client:       rb.client,
	}
}

/**
*	The `view_audit_log:{projectKey}` scope is required, and depending on the [resource type](ctp:history:type:ChangeHistoryResourceType) queried, their respective scopes must be granted.
 */
func (rb *ByProjectKeyRequestBuilder) Get() *ByProjectKeyRequestMethodGet {
	return &ByProjectKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s", rb.projectKey),
		client: rb.client,
	}
}
