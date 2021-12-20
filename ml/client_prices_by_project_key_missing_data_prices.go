// Generated file, please do not change!!!
package ml

import (
	"fmt"
)

type ByProjectKeyMissingDataPricesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMissingDataPricesRequestBuilder) Status() *ByProjectKeyMissingDataPricesStatusRequestBuilder {
	return &ByProjectKeyMissingDataPricesStatusRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMissingDataPricesRequestBuilder) Post(body MissingPricesSearchRequest) *ByProjectKeyMissingDataPricesRequestMethodPost {
	return &ByProjectKeyMissingDataPricesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/missing-data/prices", rb.projectKey),
		client: rb.client,
	}
}
