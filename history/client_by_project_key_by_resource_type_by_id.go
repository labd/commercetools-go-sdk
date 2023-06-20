package history

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyByResourceTypeByIDRequestBuilder struct {
	projectKey   string
	resourceType string
	id           string
	client       *Client
}

/**
*	The `view_audit_log:{projectKey}` scope is required, and depending on the [resource type](ctp:history:type:ChangeHistoryResourceType) queried, their respective scopes must be granted.
 */
func (rb *ByProjectKeyByResourceTypeByIDRequestBuilder) Get() *ByProjectKeyByResourceTypeByIDRequestMethodGet {
	return &ByProjectKeyByResourceTypeByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/%s/%s", rb.projectKey, rb.resourceType, rb.id),
		client: rb.client,
	}
}
