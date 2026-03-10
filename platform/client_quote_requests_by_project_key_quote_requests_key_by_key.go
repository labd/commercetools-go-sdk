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

/**
*	Retrieves a QuoteRequest with the provided `key`.
 */
func (rb *ByProjectKeyQuoteRequestsKeyByKeyRequestBuilder) Get() *ByProjectKeyQuoteRequestsKeyByKeyRequestMethodGet {
	return &ByProjectKeyQuoteRequestsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/quote-requests/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a QuoteRequest exists with the provided `key`. Returns a `200 OK` status if the QuoteRequest exists, or a [Not Found](/../api/errors#404-not-found) error otherwise.
 */
func (rb *ByProjectKeyQuoteRequestsKeyByKeyRequestBuilder) Head() *ByProjectKeyQuoteRequestsKeyByKeyRequestMethodHead {
	return &ByProjectKeyQuoteRequestsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/quote-requests/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a QuoteRequest in the Project using one or more [update actions](/../api/projects/quote-requests#update-actions).
 */
func (rb *ByProjectKeyQuoteRequestsKeyByKeyRequestBuilder) Post(body QuoteRequestUpdate) *ByProjectKeyQuoteRequestsKeyByKeyRequestMethodPost {
	return &ByProjectKeyQuoteRequestsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/quote-requests/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a QuoteRequest in the Project.
 */
func (rb *ByProjectKeyQuoteRequestsKeyByKeyRequestBuilder) Delete() *ByProjectKeyQuoteRequestsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyQuoteRequestsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/quote-requests/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
