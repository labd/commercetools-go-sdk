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

func (rb *ByProjectKeyQuoteRequestsByIDRequestBuilder) Get() *ByProjectKeyQuoteRequestsByIDRequestMethodGet {
	return &ByProjectKeyQuoteRequestsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/quote-requests/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a QuoteRequest exists for a given `id`. Returns a `200 OK` status if the QuoteRequest exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyQuoteRequestsByIDRequestBuilder) Head() *ByProjectKeyQuoteRequestsByIDRequestMethodHead {
	return &ByProjectKeyQuoteRequestsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/quote-requests/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyQuoteRequestsByIDRequestBuilder) Post(body QuoteRequestUpdate) *ByProjectKeyQuoteRequestsByIDRequestMethodPost {
	return &ByProjectKeyQuoteRequestsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/quote-requests/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyQuoteRequestsByIDRequestBuilder) Delete() *ByProjectKeyQuoteRequestsByIDRequestMethodDelete {
	return &ByProjectKeyQuoteRequestsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/quote-requests/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
