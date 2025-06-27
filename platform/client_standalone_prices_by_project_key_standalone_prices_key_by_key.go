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

/**
*	Checks if a StandalonePrice exists with the provided `key`. Returns a `200 OK` status if the StandalonePrice exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyStandalonePricesKeyByKeyRequestBuilder) Head() *ByProjectKeyStandalonePricesKeyByKeyRequestMethodHead {
	return &ByProjectKeyStandalonePricesKeyByKeyRequestMethodHead{
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
*	Produces the [StandalonePriceDeleted](ctp:api:type:StandalonePriceDeletedMessage) Message.
*
 */
func (rb *ByProjectKeyStandalonePricesKeyByKeyRequestBuilder) Delete() *ByProjectKeyStandalonePricesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyStandalonePricesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/standalone-prices/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
