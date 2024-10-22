package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMessagesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMessagesRequestBuilder) WithId(id string) *ByProjectKeyMessagesByIDRequestBuilder {
	return &ByProjectKeyMessagesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Deprecated OAuth 2.0 Scope: `view_orders:{projectKey}`
 */
func (rb *ByProjectKeyMessagesRequestBuilder) Get() *ByProjectKeyMessagesRequestMethodGet {
	return &ByProjectKeyMessagesRequestMethodGet{
		url:    fmt.Sprintf("/%s/messages", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a Message exists for a given Query Predicate. Returns a `200 OK` status if any Messages match the Query Predicate or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyMessagesRequestBuilder) Head() *ByProjectKeyMessagesRequestMethodHead {
	return &ByProjectKeyMessagesRequestMethodHead{
		url:    fmt.Sprintf("/%s/messages", rb.projectKey),
		client: rb.client,
	}
}
