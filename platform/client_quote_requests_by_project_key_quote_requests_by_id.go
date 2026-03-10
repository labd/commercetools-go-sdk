package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyQuoteRequestsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves a QuoteRequest with the provided `id`.
 */
func (rb *ByProjectKeyQuoteRequestsByIDRequestBuilder) Get() *ByProjectKeyQuoteRequestsByIDRequestMethodGet {
	return &ByProjectKeyQuoteRequestsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/quote-requests/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a QuoteRequest exists with the provided `id`. Returns a `200 OK` status if the QuoteRequest exists, or a [Not Found](/../api/errors#404-not-found) error otherwise.
 */
func (rb *ByProjectKeyQuoteRequestsByIDRequestBuilder) Head() *ByProjectKeyQuoteRequestsByIDRequestMethodHead {
	return &ByProjectKeyQuoteRequestsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/quote-requests/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a QuoteRequest in the Project using one or more [update actions](/../api/projects/quote-requests#update-actions).
 */
func (rb *ByProjectKeyQuoteRequestsByIDRequestBuilder) Post(body QuoteRequestUpdate) *ByProjectKeyQuoteRequestsByIDRequestMethodPost {
	return &ByProjectKeyQuoteRequestsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/quote-requests/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a QuoteRequest in the Project.
 */
func (rb *ByProjectKeyQuoteRequestsByIDRequestBuilder) Delete() *ByProjectKeyQuoteRequestsByIDRequestMethodDelete {
	return &ByProjectKeyQuoteRequestsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/quote-requests/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
