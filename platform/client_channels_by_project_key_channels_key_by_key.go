package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyChannelsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyChannelsKeyByKeyRequestBuilder) Get() *ByProjectKeyChannelsKeyByKeyRequestMethodGet {
	return &ByProjectKeyChannelsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/channels/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Channel exists with the provided `key`. Returns a `200 OK` status if the Channel exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyChannelsKeyByKeyRequestBuilder) Head() *ByProjectKeyChannelsKeyByKeyRequestMethodHead {
	return &ByProjectKeyChannelsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/channels/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyChannelsKeyByKeyRequestBuilder) Post(body ChannelUpdate) *ByProjectKeyChannelsKeyByKeyRequestMethodPost {
	return &ByProjectKeyChannelsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/channels/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Returns a [ReferenceExists](ctp:api:type:ReferenceExistsError) error if other resources reference the Channel to be deleted.
*
 */
func (rb *ByProjectKeyChannelsKeyByKeyRequestBuilder) Delete() *ByProjectKeyChannelsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyChannelsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/channels/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
