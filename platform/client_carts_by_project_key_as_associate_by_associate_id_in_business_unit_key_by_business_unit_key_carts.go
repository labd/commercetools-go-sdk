package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsRequestBuilder struct {
	projectKey      string
	associateId     string
	businessUnitKey string
	client          *Client
}

func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsRequestBuilder) WithKey(key string) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsKeyByKeyRequestBuilder {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsKeyByKeyRequestBuilder{
		key:             key,
		projectKey:      rb.projectKey,
		associateId:     rb.associateId,
		businessUnitKey: rb.businessUnitKey,
		client:          rb.client,
	}
}
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsRequestBuilder) WithId(id string) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsByIDRequestBuilder {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsByIDRequestBuilder{
		id:              id,
		projectKey:      rb.projectKey,
		associateId:     rb.associateId,
		businessUnitKey: rb.businessUnitKey,
		client:          rb.client,
	}
}
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsRequestBuilder) Replicate() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsReplicateRequestBuilder {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsReplicateRequestBuilder{
		projectKey:      rb.projectKey,
		associateId:     rb.associateId,
		businessUnitKey: rb.businessUnitKey,
		client:          rb.client,
	}
}
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsRequestBuilder) Get() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsRequestMethodGet {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsRequestMethodGet{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/carts", rb.projectKey, rb.associateId, rb.businessUnitKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Carts exist for the provided query predicate. Returns a `200 OK` status if any Carts match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsRequestBuilder) Head() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsRequestMethodHead {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsRequestMethodHead{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/carts", rb.projectKey, rb.associateId, rb.businessUnitKey),
		client: rb.client,
	}
}

/**
*	Creates a Cart in the [BusinessUnit](ctp:api:type:BusinessUnit) referenced by `businessUnitKey`. As such, the `businessUnit` field on [CartDraft](ctp:api:type:CartDraft) is ignored for this request.
*	Creating a Cart can fail with an [InvalidOperation](ctp:api:type:InvalidOperationError) if the referenced [ShippingMethod](ctp:api:type:ShippingMethod) in the [CartDraft](ctp:api:type:CartDraft) has a predicate that does not match the Cart.
*
*	Specific Error Codes:
*
*	- [DiscountCodeNonApplicable](ctp:api:type:DiscountCodeNonApplicableError)
*	- [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError)
*	- [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError)
*	- [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError)
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsRequestBuilder) Post(body CartDraft) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsRequestMethodPost {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyCartsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/carts", rb.projectKey, rb.associateId, rb.businessUnitKey),
		client: rb.client,
	}
}
