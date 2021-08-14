// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyChannelsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get Channel by ID
 */
func (rb *ByProjectKeyChannelsByIdRequestBuilder) Get() *ByProjectKeyChannelsByIdRequestMethodGet {
	return &ByProjectKeyChannelsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/channels/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update Channel by ID
 */
func (rb *ByProjectKeyChannelsByIdRequestBuilder) Post(body ChannelUpdate) *ByProjectKeyChannelsByIdRequestMethodPost {
	return &ByProjectKeyChannelsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/channels/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete Channel by ID
 */
func (rb *ByProjectKeyChannelsByIdRequestBuilder) Delete() *ByProjectKeyChannelsByIdRequestMethodDelete {
	return &ByProjectKeyChannelsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/channels/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
