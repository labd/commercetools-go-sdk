package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyOrdersByIDRequestBuilder) Get() *ByProjectKeyOrdersByIDRequestMethodGet {
	return &ByProjectKeyOrdersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyOrdersByIDRequestBuilder) Post(body OrderUpdate) *ByProjectKeyOrdersByIDRequestMethodPost {
	return &ByProjectKeyOrdersByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyOrdersByIDRequestBuilder) Delete() *ByProjectKeyOrdersByIDRequestMethodDelete {
	return &ByProjectKeyOrdersByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
