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
*
*	The response will include duplicate Products whenever more than one active Product Selection of the Store
*	includes a Product. To make clear through which Product Selection a Product is available in the Store
*	the response contains assignments including both the Product and the Product Selection.
*	Only Products of Product Selections that are activated in Store will be returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/product-selection-assignments", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
