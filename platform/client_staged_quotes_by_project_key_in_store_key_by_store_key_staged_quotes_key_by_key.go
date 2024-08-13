package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestBuilder struct {
	projectKey string
	storeKey   string
	key        string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a StagedQuote exists for a given `key`. Returns a `200 OK` status if the StagedQuote exists or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestBuilder) Post(body StagedQuoteUpdate) *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyStagedQuotesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/staged-quotes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
