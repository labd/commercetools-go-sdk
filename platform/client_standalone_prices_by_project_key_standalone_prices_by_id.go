package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyStandalonePricesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyStandalonePricesByIDRequestBuilder) Get() *ByProjectKeyStandalonePricesByIDRequestMethodGet {
	return &ByProjectKeyStandalonePricesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/standalone-prices/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStandalonePricesByIDRequestBuilder) Post(body StandalonePriceUpdate) *ByProjectKeyStandalonePricesByIDRequestMethodPost {
	return &ByProjectKeyStandalonePricesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/standalone-prices/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStandalonePricesByIDRequestBuilder) Delete() *ByProjectKeyStandalonePricesByIDRequestMethodDelete {
	return &ByProjectKeyStandalonePricesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/standalone-prices/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
