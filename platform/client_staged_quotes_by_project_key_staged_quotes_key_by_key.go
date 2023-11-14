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

func (rb *ByProjectKeyStagedQuotesKeyByKeyRequestBuilder) Get() *ByProjectKeyStagedQuotesKeyByKeyRequestMethodGet {
	return &ByProjectKeyStagedQuotesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/staged-quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a StagedQuote exists for a given `key`. Returns a `200 OK` status if the StagedQuote exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyStagedQuotesKeyByKeyRequestBuilder) Head() *ByProjectKeyStagedQuotesKeyByKeyRequestMethodHead {
	return &ByProjectKeyStagedQuotesKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/staged-quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStagedQuotesKeyByKeyRequestBuilder) Post(body StagedQuoteUpdate) *ByProjectKeyStagedQuotesKeyByKeyRequestMethodPost {
	return &ByProjectKeyStagedQuotesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/staged-quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStagedQuotesKeyByKeyRequestBuilder) Delete() *ByProjectKeyStagedQuotesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyStagedQuotesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/staged-quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
