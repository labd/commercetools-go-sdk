package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyQuotesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves a Quote with the provided `key`.
 */
func (rb *ByProjectKeyQuotesKeyByKeyRequestBuilder) Get() *ByProjectKeyQuotesKeyByKeyRequestMethodGet {
	return &ByProjectKeyQuotesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Quote exists with the provided `key`. Returns a `200 OK` status if the Quote exists, or a [Not Found](/../api/errors#404-not-found) error otherwise.
 */
func (rb *ByProjectKeyQuotesKeyByKeyRequestBuilder) Head() *ByProjectKeyQuotesKeyByKeyRequestMethodHead {
	return &ByProjectKeyQuotesKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a Quote in the Project using one or more [update actions](/../api/projects/quotes#update-actions).
 */
func (rb *ByProjectKeyQuotesKeyByKeyRequestBuilder) Post(body QuoteUpdate) *ByProjectKeyQuotesKeyByKeyRequestMethodPost {
	return &ByProjectKeyQuotesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a Quote in the Project.
 */
func (rb *ByProjectKeyQuotesKeyByKeyRequestBuilder) Delete() *ByProjectKeyQuotesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyQuotesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
