package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Retrieves a StagedQuote with the provided `id` in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a StagedQuote exists with the provided `id` in a [Store](ctp:api:type:Store). Returns a `200` status if the StagedQuote exists, or a `404` status otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a StagedQuote in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestBuilder) Post(body StagedQuoteUpdate) *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a StagedQuote in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
