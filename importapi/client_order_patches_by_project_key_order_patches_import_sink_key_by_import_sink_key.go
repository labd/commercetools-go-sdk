package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyRequestBuilder) ImportOperations() *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder {
	return &ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder{
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Creates a new import request for order patches
 */
func (rb *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyRequestBuilder) Post(body OrderPatchImportRequest) *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyRequestMethodPost {
	return &ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/order-patches/importSinkKey=%s", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
