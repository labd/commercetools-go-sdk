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

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a QuoteRequest exists for a given `id`. Returns a `200 OK` status if the QuoteRequest exists or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestBuilder) Post(body QuoteRequestUpdate) *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
