package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyChannelsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyChannelsRequestBuilder) WithId(id string) *ByProjectKeyChannelsByIDRequestBuilder {
	return &ByProjectKeyChannelsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyChannelsRequestBuilder) Get() *ByProjectKeyChannelsRequestMethodGet {
	return &ByProjectKeyChannelsRequestMethodGet{
		url:    fmt.Sprintf("/%s/channels", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyChannelsRequestBuilder) Post(body ChannelDraft) *ByProjectKeyChannelsRequestMethodPost {
	return &ByProjectKeyChannelsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/channels", rb.projectKey),
		client: rb.client,
	}
}
