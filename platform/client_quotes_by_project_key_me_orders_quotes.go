package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeOrdersQuotesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeOrdersQuotesRequestBuilder) Post(body MyOrderFromQuoteDraft) *ByProjectKeyMeOrdersQuotesRequestMethodPost {
	return &ByProjectKeyMeOrdersQuotesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/orders/quotes", rb.projectKey),
		client: rb.client,
	}
}
