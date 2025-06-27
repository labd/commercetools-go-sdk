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
*	Creates a new Cart by replicating an existing Cart or Order.
*
*	The following applies to the new Cart:
*
*	- It contains the same Customer information, Line Items and Custom Line Items, Custom Fields, Discount Codes, and other settings of the originating Cart or Order.
*	- If a Line Item becomes invalid, it is removed from the new Cart. A common reason for this is removed Products or Prices.
*	- Line items and Custom Line Items are reset to their initial [state](/projects/carts#itemstate).
*	- It contains no payments or delivery information.
*	- It contains up-to-date Tax Rates, Prices, and Line Item product data.
*	- The [CartState](/projects/carts#cartstate) is `Active`.
*	- If using the `customerGroup` field (for a single Customer Group) and the referenced Customer switched to another Customer Group, the new Cart is automatically updated to reflect the new group and corresponding prices.
*	- If using the `customerGroupAssignments` field (for multiple Customer Groups), the Cart no longer keeps a direct reference to a Customer Group. If a Customer’s group assignments change, the Cart and its Line Item prices are not updated automatically. Prices are only updated when the Cart is changed via a [direct update action](/projects/carts#update-actions).
*
*	Specific Error Codes:
*
*	- [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError)
*	- [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError)
*
 */
func (rb *ByProjectKeyCartsReplicateRequestBuilder) Post(body ReplicaCartDraft) *ByProjectKeyCartsReplicateRequestMethodPost {
	return &ByProjectKeyCartsReplicateRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/carts/replicate", rb.projectKey),
		client: rb.client,
	}
}
