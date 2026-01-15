package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyStagedQuotesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves a StagedQuote with the provided `key`.
 */
func (rb *ByProjectKeyStagedQuotesKeyByKeyRequestBuilder) Get() *ByProjectKeyStagedQuotesKeyByKeyRequestMethodGet {
	return &ByProjectKeyStagedQuotesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/staged-quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a StagedQuote exists with the provided `key`. Returns a `200 OK` status if the StagedQuote exists, or a [Not Found](/../api/errors#404-not-found) error otherwise.
 */
func (rb *ByProjectKeyStagedQuotesKeyByKeyRequestBuilder) Head() *ByProjectKeyStagedQuotesKeyByKeyRequestMethodHead {
	return &ByProjectKeyStagedQuotesKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/staged-quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a StagedQuote in the Project using one or more [update actions](/../api/projects/staged-quotes#update-actions).
 */
func (rb *ByProjectKeyStagedQuotesKeyByKeyRequestBuilder) Post(body StagedQuoteUpdate) *ByProjectKeyStagedQuotesKeyByKeyRequestMethodPost {
	return &ByProjectKeyStagedQuotesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/staged-quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a StagedQuote in the Project.
 */
func (rb *ByProjectKeyStagedQuotesKeyByKeyRequestBuilder) Delete() *ByProjectKeyStagedQuotesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyStagedQuotesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/staged-quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
