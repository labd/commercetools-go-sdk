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
*	Deprecated scope: `view_orders:{projectKey}`
 */
func (rb *ByProjectKeyMessagesRequestBuilder) Get() *ByProjectKeyMessagesRequestMethodGet {
	return &ByProjectKeyMessagesRequestMethodGet{
		url:    fmt.Sprintf("/%s/messages", rb.projectKey),
		client: rb.client,
	}
}
