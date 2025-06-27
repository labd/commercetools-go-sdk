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

func (rb *ByProjectKeyStagedQuotesByIDRequestBuilder) Get() *ByProjectKeyStagedQuotesByIDRequestMethodGet {
	return &ByProjectKeyStagedQuotesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/staged-quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a StagedQuote exists with the provided `id`. Returns a `200 OK` status if the StagedQuote exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyStagedQuotesByIDRequestBuilder) Head() *ByProjectKeyStagedQuotesByIDRequestMethodHead {
	return &ByProjectKeyStagedQuotesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/staged-quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStagedQuotesByIDRequestBuilder) Post(body StagedQuoteUpdate) *ByProjectKeyStagedQuotesByIDRequestMethodPost {
	return &ByProjectKeyStagedQuotesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/staged-quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStagedQuotesByIDRequestBuilder) Delete() *ByProjectKeyStagedQuotesByIDRequestMethodDelete {
	return &ByProjectKeyStagedQuotesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/staged-quotes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
