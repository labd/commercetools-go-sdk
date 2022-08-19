package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyQuotesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyQuotesRequestBuilder) WithKey(key string) *ByProjectKeyQuotesKeyByKeyRequestBuilder {
	return &ByProjectKeyQuotesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyQuotesRequestBuilder) WithId(id string) *ByProjectKeyQuotesByIDRequestBuilder {
	return &ByProjectKeyQuotesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyQuotesRequestBuilder) Get() *ByProjectKeyQuotesRequestMethodGet {
	return &ByProjectKeyQuotesRequestMethodGet{
		url:    fmt.Sprintf("/%s/quotes", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyQuotesRequestBuilder) Post(body QuoteDraft) *ByProjectKeyQuotesRequestMethodPost {
	return &ByProjectKeyQuotesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/quotes", rb.projectKey),
		client: rb.client,
	}
}