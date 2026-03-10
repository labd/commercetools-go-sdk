package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Retrieves a QuoteRequest with the provided `id` in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a QuoteRequest exists with the provided `id` in a [Store](ctp:api:type:Store). Returns a `200` status if the QuoteRequest exists, or a `404` status otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a QuoteRequest in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestBuilder) Post(body QuoteRequestUpdate) *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a QuoteRequest in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
