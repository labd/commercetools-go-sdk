package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeQuoteRequestsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyMeQuoteRequestsByIDRequestBuilder) Get() *ByProjectKeyMeQuoteRequestsByIDRequestMethodGet {
	return &ByProjectKeyMeQuoteRequestsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/quote-requests/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if my QuoteRequest exists for a given `id`. Returns a `200 OK` status if the QuoteRequest exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyMeQuoteRequestsByIDRequestBuilder) Head() *ByProjectKeyMeQuoteRequestsByIDRequestMethodHead {
	return &ByProjectKeyMeQuoteRequestsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/quote-requests/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeQuoteRequestsByIDRequestBuilder) Post(body MyQuoteRequestUpdate) *ByProjectKeyMeQuoteRequestsByIDRequestMethodPost {
	return &ByProjectKeyMeQuoteRequestsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/quote-requests/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeQuoteRequestsByIDRequestBuilder) Delete() *ByProjectKeyMeQuoteRequestsByIDRequestMethodDelete {
	return &ByProjectKeyMeQuoteRequestsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/quote-requests/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
