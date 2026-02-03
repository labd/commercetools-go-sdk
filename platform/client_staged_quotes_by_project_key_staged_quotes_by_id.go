package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyStagedQuotesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves a StagedQuote with the provided `id`.
 */
func (rb *ByProjectKeyStagedQuotesByIDRequestBuilder) Get() *ByProjectKeyStagedQuotesByIDRequestMethodGet {
	return &ByProjectKeyStagedQuotesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/staged-quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a StagedQuote exists with the provided `id`. Returns a `200 OK` status if the StagedQuote exists, or a [Not Found](/../api/errors#404-not-found) error otherwise.
 */
func (rb *ByProjectKeyStagedQuotesByIDRequestBuilder) Head() *ByProjectKeyStagedQuotesByIDRequestMethodHead {
	return &ByProjectKeyStagedQuotesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/staged-quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a StagedQuote in the Project using one or more [update actions](/../api/projects/staged-quotes#update-actions).
 */
func (rb *ByProjectKeyStagedQuotesByIDRequestBuilder) Post(body StagedQuoteUpdate) *ByProjectKeyStagedQuotesByIDRequestMethodPost {
	return &ByProjectKeyStagedQuotesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/staged-quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a StagedQuote in the Project.
 */
func (rb *ByProjectKeyStagedQuotesByIDRequestBuilder) Delete() *ByProjectKeyStagedQuotesByIDRequestMethodDelete {
	return &ByProjectKeyStagedQuotesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/staged-quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
