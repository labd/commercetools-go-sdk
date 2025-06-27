package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyQuotesRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyQuotesRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Quotes exist for the provided query predicate. Returns a `200 OK` status if any Quotes match the query predicate, or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyQuotesRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesRequestBuilder) Post(body QuoteDraft) *ByProjectKeyInStoreKeyByStoreKeyQuotesRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
