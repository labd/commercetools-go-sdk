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

/**
*	Returns all Quote Requests that match a given Query Predicate. Returns a `200 OK` status if successful.
*
 */
func (rb *ByProjectKeyMeQuoteRequestsRequestBuilder) Get() *ByProjectKeyMeQuoteRequestsRequestMethodGet {
	return &ByProjectKeyMeQuoteRequestsRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/quote-requests", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a QuoteRequest exists for a given Query Predicate. Returns a `200 OK` status if any QuoteRequests match the Query Predicate or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
*
 */
func (rb *ByProjectKeyMeQuoteRequestsRequestBuilder) Head() *ByProjectKeyMeQuoteRequestsRequestMethodHead {
	return &ByProjectKeyMeQuoteRequestsRequestMethodHead{
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
