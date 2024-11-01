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

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Quote exists for a given `id`. Returns a `200 OK` status if the Quote exists or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestBuilder) Post(body QuoteUpdate) *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyQuotesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quotes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}