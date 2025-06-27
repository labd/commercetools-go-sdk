package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a StagedQuote exists with the provided `id`. Returns a `200 OK` status if the StagedQuote exists or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestBuilder) Post(body StagedQuoteUpdate) *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
