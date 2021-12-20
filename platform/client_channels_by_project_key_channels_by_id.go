// Generated file, please do not change!!!
package platform

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
