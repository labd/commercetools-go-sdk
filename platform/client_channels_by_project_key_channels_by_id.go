package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyChannelsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyChannelsByIDRequestBuilder) Get() *ByProjectKeyChannelsByIDRequestMethodGet {
	return &ByProjectKeyChannelsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/channels/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Channel exists for a given `id`. Returns a `200 OK` status if the Channel exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyChannelsByIDRequestBuilder) Head() *ByProjectKeyChannelsByIDRequestMethodHead {
	return &ByProjectKeyChannelsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/channels/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyChannelsByIDRequestBuilder) Post(body ChannelUpdate) *ByProjectKeyChannelsByIDRequestMethodPost {
	return &ByProjectKeyChannelsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/channels/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyChannelsByIDRequestBuilder) Delete() *ByProjectKeyChannelsByIDRequestMethodDelete {
	return &ByProjectKeyChannelsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/channels/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
