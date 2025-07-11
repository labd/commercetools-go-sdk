package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDiscountCodesImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates an Import Request for Discount Codes.
 */
func (rb *ByProjectKeyDiscountCodesImportContainersByImportContainerKeyRequestBuilder) Post(body DiscountCodeImportRequest) *ByProjectKeyDiscountCodesImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyDiscountCodesImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/discount-codes/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
