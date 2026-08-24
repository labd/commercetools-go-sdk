package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	Retrieves StagedQuotes in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more StagedQuotes exist for the provided query predicate in a [Store](ctp:api:type:Store). Returns a `200` status if any StagedQuotes match the query predicate, or a `404` status otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Creates a StagedQuote in a [Store](ctp:api:type:Store).
*
*	- [InvalidOperation](ctp:api:type:InvalidOperationError) is returned in several cases, including the following:
*	    - The referenced Shipping Method is not active.
*	    - The referenced Shipping Method is scoped to a Store that differs from the Store referenced by the [Quote Request](ctp:api:type:QuoteRequest).
*	    - The referenced Shipping Method is scoped to a Store, but the [Quote Request](ctp:api:type:QuoteRequest) does not belong to a Store.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestBuilder) Post(body StagedQuoteDraft) *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
