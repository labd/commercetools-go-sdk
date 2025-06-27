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

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Quote exists with the provided `key`. Returns a `200 OK` status if the Quote exists or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestBuilder) Post(body QuoteUpdate) *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
