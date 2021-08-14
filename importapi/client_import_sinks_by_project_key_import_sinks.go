// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyImportSinksRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyImportSinksRequestBuilder) WithImportSinkKeyValue(importSinkKey string) *ByProjectKeyImportSinksByImportSinkKeyRequestBuilder {
	return &ByProjectKeyImportSinksByImportSinkKeyRequestBuilder{
		importSinkKey: importSinkKey,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}

/**
*	Creates a new import sink.
 */
func (rb *ByProjectKeyImportSinksRequestBuilder) Post(body ImportSinkDraft) *ByProjectKeyImportSinksRequestMethodPost {
	return &ByProjectKeyImportSinksRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/import-sinks", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Retrieves all import sinks of a project key.
 */
func (rb *ByProjectKeyImportSinksRequestBuilder) Get() *ByProjectKeyImportSinksRequestMethodGet {
	return &ByProjectKeyImportSinksRequestMethodGet{
		url:    fmt.Sprintf("/%s/import-sinks", rb.projectKey),
		client: rb.client,
	}
}
