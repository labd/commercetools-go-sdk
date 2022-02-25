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
*	Create an Order by Import
 */
func (rb *ByProjectKeyOrdersImportRequestBuilder) Post(body OrderImportDraft) *ByProjectKeyOrdersImportRequestMethodPost {
	return &ByProjectKeyOrdersImportRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/import", rb.projectKey),
		client: rb.client,
	}
}
