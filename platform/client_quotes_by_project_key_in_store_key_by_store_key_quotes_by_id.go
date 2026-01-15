package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Retrieves a Quote with the provided `id` in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Quote exists with the provided `id` in a [Store](ctp:api:type:Store). Returns a `200` status if the Quote exists, or a `404` status otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a Quote in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestBuilder) Post(body QuoteUpdate) *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a Quote in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
