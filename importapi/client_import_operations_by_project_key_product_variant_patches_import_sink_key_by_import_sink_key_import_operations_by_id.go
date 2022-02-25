package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder struct {
	projectKey    string
	importSinkKey string
	id            string
	client        *Client
}

/**
*	Retrieves the import operation with the given id.
*
 */
func (rb *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder) Get() *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet {
	return &ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-variant-patches/importSinkKey=%s/import-operations/%s", rb.projectKey, rb.importSinkKey, rb.id),
		client: rb.client,
	}
}
