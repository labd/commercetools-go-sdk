package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyStatesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyStatesByIDRequestBuilder) Get() *ByProjectKeyStatesByIDRequestMethodGet {
	return &ByProjectKeyStatesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/states/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStatesByIDRequestBuilder) Post(body StateUpdate) *ByProjectKeyStatesByIDRequestMethodPost {
	return &ByProjectKeyStatesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/states/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStatesByIDRequestBuilder) Delete() *ByProjectKeyStatesByIDRequestMethodDelete {
	return &ByProjectKeyStatesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/states/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
