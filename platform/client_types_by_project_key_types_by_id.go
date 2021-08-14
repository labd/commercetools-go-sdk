// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyTypesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyTypesByIDRequestBuilder) Get() *ByProjectKeyTypesByIDRequestMethodGet {
	return &ByProjectKeyTypesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyTypesByIDRequestBuilder) Post(body TypeUpdate) *ByProjectKeyTypesByIDRequestMethodPost {
	return &ByProjectKeyTypesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyTypesByIDRequestBuilder) Delete() *ByProjectKeyTypesByIDRequestMethodDelete {
	return &ByProjectKeyTypesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
