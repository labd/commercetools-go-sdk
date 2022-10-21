package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeQuotesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeQuotesRequestBuilder) WithId(id string) *ByProjectKeyMeQuotesByIDRequestBuilder {
	return &ByProjectKeyMeQuotesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeQuotesRequestBuilder) WithKey(key string) *ByProjectKeyMeQuotesKeyByKeyRequestBuilder {
	return &ByProjectKeyMeQuotesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeQuotesRequestBuilder) Get() *ByProjectKeyMeQuotesRequestMethodGet {
	return &ByProjectKeyMeQuotesRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/quotes", rb.projectKey),
		client: rb.client,
	}
}
