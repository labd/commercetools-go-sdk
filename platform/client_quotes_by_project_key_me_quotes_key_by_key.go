package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeQuotesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyMeQuotesKeyByKeyRequestBuilder) Get() *ByProjectKeyMeQuotesKeyByKeyRequestMethodGet {
	return &ByProjectKeyMeQuotesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Quote exists with the provided `key`. Returns a `200 OK` status if the Quote exists or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
 */
func (rb *ByProjectKeyMeQuotesKeyByKeyRequestBuilder) Head() *ByProjectKeyMeQuotesKeyByKeyRequestMethodHead {
	return &ByProjectKeyMeQuotesKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeQuotesKeyByKeyRequestBuilder) Post(body MyQuoteUpdate) *ByProjectKeyMeQuotesKeyByKeyRequestMethodPost {
	return &ByProjectKeyMeQuotesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/quotes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
