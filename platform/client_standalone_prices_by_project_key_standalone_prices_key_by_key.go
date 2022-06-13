package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyStandalonePricesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyStandalonePricesKeyByKeyRequestBuilder) Get() *ByProjectKeyStandalonePricesKeyByKeyRequestMethodGet {
	return &ByProjectKeyStandalonePricesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/standalone-prices/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStandalonePricesKeyByKeyRequestBuilder) Post(body StandalonePriceUpdate) *ByProjectKeyStandalonePricesKeyByKeyRequestMethodPost {
	return &ByProjectKeyStandalonePricesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/standalone-prices/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Produces the [StandalonePriceDeletedMessage](ctp:api:type:StandalonePriceDeletedMessage).
*
 */
func (rb *ByProjectKeyStandalonePricesKeyByKeyRequestBuilder) Delete() *ByProjectKeyStandalonePricesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyStandalonePricesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/standalone-prices/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
