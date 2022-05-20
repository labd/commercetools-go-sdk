package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyStandalonePricesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyStandalonePricesRequestBuilder) WithKey(key string) *ByProjectKeyStandalonePricesKeyByKeyRequestBuilder {
	return &ByProjectKeyStandalonePricesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyStandalonePricesRequestBuilder) WithId(id string) *ByProjectKeyStandalonePricesByIDRequestBuilder {
	return &ByProjectKeyStandalonePricesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyStandalonePricesRequestBuilder) Get() *ByProjectKeyStandalonePricesRequestMethodGet {
	return &ByProjectKeyStandalonePricesRequestMethodGet{
		url:    fmt.Sprintf("/%s/standalone-prices", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStandalonePricesRequestBuilder) Post(body StandalonePriceDraft) *ByProjectKeyStandalonePricesRequestMethodPost {
	return &ByProjectKeyStandalonePricesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/standalone-prices", rb.projectKey),
		client: rb.client,
	}
}
