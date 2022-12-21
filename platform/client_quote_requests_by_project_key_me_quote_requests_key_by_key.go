package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeQuoteRequestsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyMeQuoteRequestsKeyByKeyRequestBuilder) Post(body MyQuoteRequestUpdate) *ByProjectKeyMeQuoteRequestsKeyByKeyRequestMethodPost {
	return &ByProjectKeyMeQuoteRequestsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/quote-requests/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeQuoteRequestsKeyByKeyRequestBuilder) Delete() *ByProjectKeyMeQuoteRequestsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyMeQuoteRequestsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/quote-requests/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeQuoteRequestsKeyByKeyRequestBuilder) Get() *ByProjectKeyMeQuoteRequestsKeyByKeyRequestMethodGet {
	return &ByProjectKeyMeQuoteRequestsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/quote-requests/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
