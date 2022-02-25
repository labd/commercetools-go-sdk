package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeOrdersByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyMeOrdersByIDRequestBuilder) Get() *ByProjectKeyMeOrdersByIDRequestMethodGet {
	return &ByProjectKeyMeOrdersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
