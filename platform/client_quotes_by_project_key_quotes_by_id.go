package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyQuotesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves a Quote with the provided `id`.
 */
func (rb *ByProjectKeyQuotesByIDRequestBuilder) Get() *ByProjectKeyQuotesByIDRequestMethodGet {
	return &ByProjectKeyQuotesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Quote exists with the provided `id`. Returns a `200 OK` status if the Quote exists, or a [Not Found](/../api/errors#404-not-found) error otherwise.
 */
func (rb *ByProjectKeyQuotesByIDRequestBuilder) Head() *ByProjectKeyQuotesByIDRequestMethodHead {
	return &ByProjectKeyQuotesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a Quote in the Project using one or more [update actions](/../api/projects/quotes#update-actions).
 */
func (rb *ByProjectKeyQuotesByIDRequestBuilder) Post(body QuoteUpdate) *ByProjectKeyQuotesByIDRequestMethodPost {
	return &ByProjectKeyQuotesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a Quote in the Project.
 */
func (rb *ByProjectKeyQuotesByIDRequestBuilder) Delete() *ByProjectKeyQuotesByIDRequestMethodDelete {
	return &ByProjectKeyQuotesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
