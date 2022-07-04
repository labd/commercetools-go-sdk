package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyQuotesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyQuotesByIDRequestBuilder) Get() *ByProjectKeyQuotesByIDRequestMethodGet {
	return &ByProjectKeyQuotesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyQuotesByIDRequestBuilder) Post(body QuoteUpdate) *ByProjectKeyQuotesByIDRequestMethodPost {
	return &ByProjectKeyQuotesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyQuotesByIDRequestBuilder) Delete() *ByProjectKeyQuotesByIDRequestMethodDelete {
	return &ByProjectKeyQuotesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
