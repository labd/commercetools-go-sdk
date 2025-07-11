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
func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more StagedQuotes exist for the provided query predicate. Returns a `200 OK` status if any StagedQuotes match the query predicate, or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestBuilder) Post(body StagedQuoteDraft) *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
