package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeQuoteRequestsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyMeQuoteRequestsKeyByKeyRequestBuilder) Get() *ByProjectKeyMeQuoteRequestsKeyByKeyRequestMethodGet {
	return &ByProjectKeyMeQuoteRequestsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/quote-requests/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a QuoteRequest exists with the provided `key`. Returns a `200 OK` status if the QuoteRequest exists or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
 */
func (rb *ByProjectKeyMeQuoteRequestsKeyByKeyRequestBuilder) Head() *ByProjectKeyMeQuoteRequestsKeyByKeyRequestMethodHead {
	return &ByProjectKeyMeQuoteRequestsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/quote-requests/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeQuoteRequestsKeyByKeyRequestBuilder) Post(body MyQuoteRequestUpdate) *ByProjectKeyMeQuoteRequestsKeyByKeyRequestMethodPost {
	return &ByProjectKeyMeQuoteRequestsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/quote-requests/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
