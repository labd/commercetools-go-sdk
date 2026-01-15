package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestBuilder struct {
	projectKey string
	storeKey   string
	key        string
	client     *Client
}

/**
*	Retrieves a QuoteRequest with the provided `key` in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a QuoteRequest exists with the provided `key` in a [Store](ctp:api:type:Store). Returns a `200` status if the QuoteRequest exists, or a `404` status otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a QuoteRequest in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestBuilder) Post(body QuoteRequestUpdate) *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a QuoteRequest in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
