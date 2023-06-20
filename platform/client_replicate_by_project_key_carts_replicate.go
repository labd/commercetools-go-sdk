package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCartsReplicateRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Creates a new Cart by replicating an existing Cart or Order. Can be useful in cases where a customer wants to cancel a recent order to make some changes or reorder a previous order.
*
*	The replicated Cart preserves Customer information, Line Items and Custom Line Items, Custom Fields, Discount Codes, and other settings of the Cart or Order. If the Line Items become invalid, for example, due to removed Products or Prices, they are removed from the new Cart. If the Customer switches to another Customer Group, the new Cart is updated with the new value. It has up-to-date Tax Rates, Prices, and Line Item product data and is in `Active` [CartState](ctp:api:type:CartState).
*
*	The new Cart does not contain Payments or Deliveries. The [State](ctp:api:type:ItemState) of Line Items and Custom Line Items is reset to `initial`.
*
 */
func (rb *ByProjectKeyCartsReplicateRequestBuilder) Post(body ReplicaCartDraft) *ByProjectKeyCartsReplicateRequestMethodPost {
	return &ByProjectKeyCartsReplicateRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/carts/replicate", rb.projectKey),
		client: rb.client,
	}
}
