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
func (rb *ByProjectKeyChannelsRequestBuilder) WithKey(key string) *ByProjectKeyChannelsKeyByKeyRequestBuilder {
	return &ByProjectKeyChannelsKeyByKeyRequestBuilder{
		key:        key,
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

/**
*	Checks if one or more Channels exist for the provided query predicate. Returns a `200 OK` status if any Channels match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyChannelsRequestBuilder) Head() *ByProjectKeyChannelsRequestMethodHead {
	return &ByProjectKeyChannelsRequestMethodHead{
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
