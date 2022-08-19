package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersQuotesRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Create an Order from a Quote
 */
func (rb *ByProjectKeyOrdersQuotesRequestBuilder) Post(body OrderFromQuoteDraft) *ByProjectKeyOrdersQuotesRequestMethodPost {
	return &ByProjectKeyOrdersQuotesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/quotes", rb.projectKey),
		client: rb.client,
	}
}
