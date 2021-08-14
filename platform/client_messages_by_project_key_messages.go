// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyMessagesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMessagesRequestBuilder) WithId(id string) *ByProjectKeyMessagesByIdRequestBuilder {
	return &ByProjectKeyMessagesByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Query messages
 */
func (rb *ByProjectKeyMessagesRequestBuilder) Get() *ByProjectKeyMessagesRequestMethodGet {
	return &ByProjectKeyMessagesRequestMethodGet{
		url:    fmt.Sprintf("/%s/messages", rb.projectKey),
		client: rb.client,
	}
}
