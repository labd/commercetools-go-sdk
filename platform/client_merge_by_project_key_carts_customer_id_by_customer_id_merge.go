package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCartsCustomerIdByCustomerIdMergeRequestBuilder struct {
	projectKey string
	customerId string
	client     *Client
}

/**
*	Merges items from an anonymous Cart into the most recently modified active Cart of a Customer. If no active Cart exists, the anonymous Cart becomes the Customer's active Cart.
*	For more information about merge mode behaviors, merge rules, and tax recalculation, see [Merge a Cart](/../api/carts-orders-overview#merge-a-cart).
*
 */
func (rb *ByProjectKeyCartsCustomerIdByCustomerIdMergeRequestBuilder) Post(body MergeCartDraft) *ByProjectKeyCartsCustomerIdByCustomerIdMergeRequestMethodPost {
	return &ByProjectKeyCartsCustomerIdByCustomerIdMergeRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/carts/customer-id=%s/merge", rb.projectKey, rb.customerId),
		client: rb.client,
	}
}
