package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyPricesImportSinkKeyByImportSinkKeyRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyPricesImportSinkKeyByImportSinkKeyRequestBuilder) ImportOperations() *ByProjectKeyPricesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder {
	return &ByProjectKeyPricesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder{
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Creates import request for creating new prices or updating existing ones.
 */
func (rb *ByProjectKeyPricesImportSinkKeyByImportSinkKeyRequestBuilder) Post(body PriceImportRequest) *ByProjectKeyPricesImportSinkKeyByImportSinkKeyRequestMethodPost {
	return &ByProjectKeyPricesImportSinkKeyByImportSinkKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/prices/importSinkKey=%s", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
