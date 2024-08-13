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

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a QuoteRequest exists for a given `key`. Returns a `200 OK` status if the QuoteRequest exists or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestBuilder) Post(body QuoteRequestUpdate) *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
