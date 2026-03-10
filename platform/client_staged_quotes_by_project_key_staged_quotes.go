package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyStagedQuotesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyStagedQuotesRequestBuilder) WithKey(key string) *ByProjectKeyStagedQuotesKeyByKeyRequestBuilder {
	return &ByProjectKeyStagedQuotesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyStagedQuotesRequestBuilder) WithId(id string) *ByProjectKeyStagedQuotesByIDRequestBuilder {
	return &ByProjectKeyStagedQuotesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves all StagedQuotes in the Project.
 */
func (rb *ByProjectKeyStagedQuotesRequestBuilder) Get() *ByProjectKeyStagedQuotesRequestMethodGet {
	return &ByProjectKeyStagedQuotesRequestMethodGet{
		url:    fmt.Sprintf("/%s/staged-quotes", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more StagedQuotes exist for the provided query predicate. Returns a `200 OK` status if any StagedQuotes match the query predicate or a [Not Found](/../api/errors#404-not-found) error otherwise.
 */
func (rb *ByProjectKeyStagedQuotesRequestBuilder) Head() *ByProjectKeyStagedQuotesRequestMethodHead {
	return &ByProjectKeyStagedQuotesRequestMethodHead{
		url:    fmt.Sprintf("/%s/staged-quotes", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a StagedQuote in the Project.
 */
func (rb *ByProjectKeyStagedQuotesRequestBuilder) Post(body StagedQuoteDraft) *ByProjectKeyStagedQuotesRequestMethodPost {
	return &ByProjectKeyStagedQuotesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/staged-quotes", rb.projectKey),
		client: rb.client,
	}
}
