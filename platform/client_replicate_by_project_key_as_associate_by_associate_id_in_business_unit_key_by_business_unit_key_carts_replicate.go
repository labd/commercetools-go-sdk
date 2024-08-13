package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsReplicateRequestBuilder struct {
	projectKey      string
	associateId     string
	businessUnitKey string
	client          *Client
}

/**
*	Creates a new Cart by replicating an existing Cart or Order. Can be useful in cases where a customer wants to cancel a recent order to make some changes or reorder a previous order.
*
*	The replicated Cart preserves Customer information, Line Items and Custom Line Items, Custom Fields, Discount Codes, and other settings of the Cart or Order. If the Line Items become invalid, for example, due to removed Products or Prices, they are removed from the new Cart. If the Customer switches to another Customer Group, the new Cart is updated with the new value. It has up-to-date Tax Rates, Prices, and Line Item product data and is in `Active` [CartState](ctp:api:type:CartState).
*
*	The new Cart does not contain Payments or Deliveries. The [State](ctp:api:type:ItemState) of Line Items and Custom Line Items is reset to `initial`.
*
*	If the Cart exists in the [Project](ctp:api:type:Project) but does not reference the requested [BusinessUnit](ctp:api:type:BusinessUnit), this method returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
*
*	Specific Error Codes:
*
*	- [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError)
*	- [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError)
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsReplicateRequestBuilder) Post(body ReplicaCartDraft) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsReplicateRequestMethodPost {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsReplicateRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/carts/replicate", rb.projectKey, rb.associateId, rb.businessUnitKey),
		client: rb.client,
	}
}
