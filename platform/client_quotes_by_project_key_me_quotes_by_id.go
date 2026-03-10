package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeQuotesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves a Quote with the provided `id` for the authenticated Customer.
*
 */
func (rb *ByProjectKeyMeQuotesByIDRequestBuilder) Get() *ByProjectKeyMeQuotesByIDRequestMethodGet {
	return &ByProjectKeyMeQuotesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Quote exists with the provided `id` for the authenticated Customer.
*	Returns a `200 OK` status if the Quote exists or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
*
 */
func (rb *ByProjectKeyMeQuotesByIDRequestBuilder) Head() *ByProjectKeyMeQuotesByIDRequestMethodHead {
	return &ByProjectKeyMeQuotesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a Quote for the authenticated Customer using one or more [update actions](/../api/projects/quotes#update-actions).
*
 */
func (rb *ByProjectKeyMeQuotesByIDRequestBuilder) Post(body MyQuoteUpdate) *ByProjectKeyMeQuotesByIDRequestMethodPost {
	return &ByProjectKeyMeQuotesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
