package history

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyByResourceTypeRequestBuilder struct {
	projectKey   string
	resourceType string
	client       *Client
}

func (rb *ByProjectKeyByResourceTypeRequestBuilder) WithIdValue(id string) *ByProjectKeyByResourceTypeByIDRequestBuilder {
	return &ByProjectKeyByResourceTypeByIDRequestBuilder{
		id:           id,
		projectKey:   rb.projectKey,
		resourceType: rb.resourceType,
		client:       rb.client,
	}
}

/**
*	The `view_audit_log:{projectKey}` scope is required, and depending on the [resource type](ctp:history:type:ChangeHistoryResourceType) queried, their respective scopes must be granted.
 */
func (rb *ByProjectKeyByResourceTypeRequestBuilder) Get() *ByProjectKeyByResourceTypeRequestMethodGet {
	return &ByProjectKeyByResourceTypeRequestMethodGet{
		url:    fmt.Sprintf("/%s/%s", rb.projectKey, rb.resourceType),
		client: rb.client,
	}
}
