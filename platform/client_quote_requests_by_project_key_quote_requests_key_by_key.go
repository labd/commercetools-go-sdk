package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyQuoteRequestsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyQuoteRequestsKeyByKeyRequestBuilder) Get() *ByProjectKeyQuoteRequestsKeyByKeyRequestMethodGet {
	return &ByProjectKeyQuoteRequestsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/quote-requests/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a QuoteRequest exists with the provided `key`. Returns a `200 OK` status if the QuoteRequest exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyQuoteRequestsKeyByKeyRequestBuilder) Head() *ByProjectKeyQuoteRequestsKeyByKeyRequestMethodHead {
	return &ByProjectKeyQuoteRequestsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/quote-requests/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyQuoteRequestsKeyByKeyRequestBuilder) Post(body QuoteRequestUpdate) *ByProjectKeyQuoteRequestsKeyByKeyRequestMethodPost {
	return &ByProjectKeyQuoteRequestsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/quote-requests/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyQuoteRequestsKeyByKeyRequestBuilder) Delete() *ByProjectKeyQuoteRequestsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyQuoteRequestsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/quote-requests/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
