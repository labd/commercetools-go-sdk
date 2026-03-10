package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestBuilder struct {
	projectKey string
	storeKey   string
	key        string
	client     *Client
}

/**
*	Retrieves a Quote with the provided `key` in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Quote exists with the provided `key` in a [Store](ctp:api:type:Store). Returns a `200` status if the Quote exists, or a `404` status otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a Quote in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestBuilder) Post(body QuoteUpdate) *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a Quote in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
