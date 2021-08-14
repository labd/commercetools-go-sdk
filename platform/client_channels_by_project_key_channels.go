// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyChannelsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyChannelsRequestBuilder) WithId(id string) *ByProjectKeyChannelsByIdRequestBuilder {
	return &ByProjectKeyChannelsByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Query channels
 */
func (rb *ByProjectKeyChannelsRequestBuilder) Get() *ByProjectKeyChannelsRequestMethodGet {
	return &ByProjectKeyChannelsRequestMethodGet{
		url:    fmt.Sprintf("/%s/channels", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Create Channel
 */
func (rb *ByProjectKeyChannelsRequestBuilder) Post(body ChannelDraft) *ByProjectKeyChannelsRequestMethodPost {
	return &ByProjectKeyChannelsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/channels", rb.projectKey),
		client: rb.client,
	}
}
