package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyQuotesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyQuotesKeyByKeyRequestBuilder) Get() *ByProjectKeyQuotesKeyByKeyRequestMethodGet {
	return &ByProjectKeyQuotesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyQuotesKeyByKeyRequestBuilder) Post(body QuoteUpdate) *ByProjectKeyQuotesKeyByKeyRequestMethodPost {
	return &ByProjectKeyQuotesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyQuotesKeyByKeyRequestBuilder) Delete() *ByProjectKeyQuotesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyQuotesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
