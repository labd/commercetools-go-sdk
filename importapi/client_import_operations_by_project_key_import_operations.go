package importapi

// Generated file, please do not change!!!

import ()

type ByProjectKeyImportOperationsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyImportOperationsRequestBuilder) WithIdValue(id string) *ByProjectKeyImportOperationsByIdRequestBuilder {
	return &ByProjectKeyImportOperationsByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
