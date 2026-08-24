package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdMergeRequestBuilder struct {
	projectKey string
	storeKey   string
	customerId string
	client     *Client
}

/**
*	Merges items from an anonymous Cart into the most recently modified active Cart of a Customer. If no active Cart exists, the anonymous Cart becomes the Customer's active Cart.
*
*	If the Cart exists in the Project but does not have a `store` specified, or the `store` field references a different Store, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
*	- [InvalidOperation](ctp:api:type:InvalidOperationError) is returned in several cases, including the following:
*	    - The referenced Shipping Method is not active.
*	    - The referenced Shipping Method is scoped to a Store that differs from the Cart's Store.
*	    - The referenced Shipping Method is scoped to a Store, but the Cart does not belong to a Store.
*
*	For more information about merge mode behaviors, merge rules, and tax recalculation, see [Merge a Cart](/api/carts-orders-overview#merge-a-cart).
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdMergeRequestBuilder) Post(body MergeCartDraft) *ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdMergeRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdMergeRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/customer-id=%s/merge", rb.projectKey, rb.storeKey, rb.customerId),
		client: rb.client,
	}
}
