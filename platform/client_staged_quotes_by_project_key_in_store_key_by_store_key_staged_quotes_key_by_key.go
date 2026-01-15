package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestBuilder struct {
	projectKey string
	storeKey   string
	key        string
	client     *Client
}

/**
*	Retrieves a StagedQuote with the provided `key` in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a StagedQuote exists with the provided `key` in a [Store](ctp:api:type:Store). Returns a `200` status if the StagedQuote exists, or a `404` status otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a StagedQuote in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestBuilder) Post(body StagedQuoteUpdate) *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a StagedQuote in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
