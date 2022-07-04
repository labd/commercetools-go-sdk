package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyQuoteRequestsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyQuoteRequestsRequestBuilder) WithKey(key string) *ByProjectKeyQuoteRequestsKeyByKeyRequestBuilder {
	return &ByProjectKeyQuoteRequestsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyQuoteRequestsRequestBuilder) WithId(id string) *ByProjectKeyQuoteRequestsByIDRequestBuilder {
	return &ByProjectKeyQuoteRequestsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyQuoteRequestsRequestBuilder) Get() *ByProjectKeyQuoteRequestsRequestMethodGet {
	return &ByProjectKeyQuoteRequestsRequestMethodGet{
		url:    fmt.Sprintf("/%s/quote-requests", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyQuoteRequestsRequestBuilder) Post(body QuoteRequestDraft) *ByProjectKeyQuoteRequestsRequestMethodPost {
	return &ByProjectKeyQuoteRequestsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/quote-requests", rb.projectKey),
		client: rb.client,
	}
}
