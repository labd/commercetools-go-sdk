package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeActiveCartRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeActiveCartRequestBuilder) Get() *ByProjectKeyMeActiveCartRequestMethodGet {
	return &ByProjectKeyMeActiveCartRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/active-cart", rb.projectKey),
		client: rb.client,
	}
}
