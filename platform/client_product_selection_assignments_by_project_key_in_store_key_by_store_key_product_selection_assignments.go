package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Queries Product Selection assignments in a specific Store.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/product-selection-assignments", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
