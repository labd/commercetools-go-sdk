package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeQuoteRequestsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeQuoteRequestsRequestBuilder) WithId(id string) *ByProjectKeyMeQuoteRequestsByIDRequestBuilder {
	return &ByProjectKeyMeQuoteRequestsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeQuoteRequestsRequestBuilder) WithKey(key string) *ByProjectKeyMeQuoteRequestsKeyByKeyRequestBuilder {
	return &ByProjectKeyMeQuoteRequestsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeQuoteRequestsRequestBuilder) Get() *ByProjectKeyMeQuoteRequestsRequestMethodGet {
	return &ByProjectKeyMeQuoteRequestsRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/quote-requests", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeQuoteRequestsRequestBuilder) Post(body MyQuoteRequestDraft) *ByProjectKeyMeQuoteRequestsRequestMethodPost {
	return &ByProjectKeyMeQuoteRequestsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/quote-requests", rb.projectKey),
		client: rb.client,
	}
}
