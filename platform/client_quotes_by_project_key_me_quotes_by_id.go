package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeQuotesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyMeQuotesByIDRequestBuilder) Get() *ByProjectKeyMeQuotesByIDRequestMethodGet {
	return &ByProjectKeyMeQuotesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeQuotesByIDRequestBuilder) Post(body MyQuoteUpdate) *ByProjectKeyMeQuotesByIDRequestMethodPost {
	return &ByProjectKeyMeQuotesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
