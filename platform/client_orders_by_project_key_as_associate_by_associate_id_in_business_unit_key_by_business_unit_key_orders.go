package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersRequestBuilder struct {
	projectKey      string
	associateId     string
	businessUnitKey string
	client          *Client
}

func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersRequestBuilder) OrderQuote() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersQuotesRequestBuilder {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersQuotesRequestBuilder{
		projectKey:      rb.projectKey,
		associateId:     rb.associateId,
		businessUnitKey: rb.businessUnitKey,
		client:          rb.client,
	}
}
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersRequestBuilder) WithOrderNumber(orderNumber string) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersOrderNumberByOrderNumberRequestBuilder {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersOrderNumberByOrderNumberRequestBuilder{
		orderNumber:     orderNumber,
		projectKey:      rb.projectKey,
		associateId:     rb.associateId,
		businessUnitKey: rb.businessUnitKey,
		client:          rb.client,
	}
}
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersRequestBuilder) WithId(id string) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersByIDRequestBuilder {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersByIDRequestBuilder{
		id:              id,
		projectKey:      rb.projectKey,
		associateId:     rb.associateId,
		businessUnitKey: rb.businessUnitKey,
		client:          rb.client,
	}
}

/**
*	Retrieves Orders in a [BusinessUnit](ctp:api:type:BusinessUnit).
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersRequestBuilder) Get() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersRequestMethodGet {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersRequestMethodGet{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/orders", rb.projectKey, rb.associateId, rb.businessUnitKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Orders exist with provided query predicate in a [BusinessUnit](ctp:api:type:BusinessUnit). Returns a `200 OK` status if any Orders match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersRequestBuilder) Head() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersRequestMethodHead {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersRequestMethodHead{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/orders", rb.projectKey, rb.associateId, rb.businessUnitKey),
		client: rb.client,
	}
}

/**
*	Creates an Order from a [Cart](ctp:api:type:Cart) in a [BusinessUnit](ctp:api:type:BusinessUnit).
*
*	The Cart must have a shipping address and an active Shipping Method set.
*
*	If the Cart does not reference the same BusinessUnit as the `businessUnitKey` path parameter, an [InvalidOperation](ctp:api:type:InvalidOperationError) is returned.
*
*	Specific Error Codes:
*
*	- [DiscountCodeNonApplicable](ctp:api:type:DiscountCodeNonApplicableError)
*	- [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError)
*	- [OutOfStock](ctp:api:type:OutOfStockError)
*	- [PriceChanged](ctp:api:type:PriceChangedError)
*	- [ShippingMethodDoesNotMatchCart](ctp:api:type:ShippingMethodDoesNotMatchCartError)
*	- [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError)
*	- [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError)
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersRequestBuilder) Post(body OrderFromCartDraft) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersRequestMethodPost {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/orders", rb.projectKey, rb.associateId, rb.businessUnitKey),
		client: rb.client,
	}
}
