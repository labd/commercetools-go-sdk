package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeCartsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyMeCartsByIDRequestBuilder) Get() *ByProjectKeyMeCartsByIDRequestMethodGet {
	return &ByProjectKeyMeCartsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeCartsByIDRequestBuilder) Post(body MyCartUpdate) *ByProjectKeyMeCartsByIDRequestMethodPost {
	return &ByProjectKeyMeCartsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeCartsByIDRequestBuilder) Delete() *ByProjectKeyMeCartsByIDRequestMethodDelete {
	return &ByProjectKeyMeCartsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
