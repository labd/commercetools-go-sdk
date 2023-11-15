package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersImportRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Importing an Order produces the [Order Imported](ctp:api:type:OrderImportedMessage) Message.
*
*	Specific Error Codes:
*
*	- [OutOfStock](ctp:api:type:OutOfStockError)
*	- [CountryNotConfiguredInStore](ctp:api:type:CountryNotConfiguredInStoreError)
*
 */
func (rb *ByProjectKeyOrdersImportRequestBuilder) Post(body OrderImportDraft) *ByProjectKeyOrdersImportRequestMethodPost {
	return &ByProjectKeyOrdersImportRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/import", rb.projectKey),
		client: rb.client,
	}
}
