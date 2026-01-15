package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	Retrieves QuoteRequests in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more QuoteRequests exist for the provided query predicate in a [Store](ctp:api:type:Store). Returns a `200` status if any QuoteRequests match the query predicate, or a `404` status otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Creates a QuoteRequest in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestBuilder) Post(body QuoteRequestDraft) *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
