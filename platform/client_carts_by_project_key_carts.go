package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCartsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCartsRequestBuilder) Replicate() *ByProjectKeyCartsReplicateRequestBuilder {
	return &ByProjectKeyCartsReplicateRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCartsRequestBuilder) WithCustomerId(customerId string) *ByProjectKeyCartsCustomerIdByCustomerIdRequestBuilder {
	return &ByProjectKeyCartsCustomerIdByCustomerIdRequestBuilder{
		customerId: customerId,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCartsRequestBuilder) WithKey(key string) *ByProjectKeyCartsKeyByKeyRequestBuilder {
	return &ByProjectKeyCartsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCartsRequestBuilder) WithId(id string) *ByProjectKeyCartsByIDRequestBuilder {
	return &ByProjectKeyCartsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves Carts in the Project.
 */
func (rb *ByProjectKeyCartsRequestBuilder) Get() *ByProjectKeyCartsRequestMethodGet {
	return &ByProjectKeyCartsRequestMethodGet{
		url:    fmt.Sprintf("/%s/carts", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Carts exist for the provided query predicate. Returns a `200 OK` status if any Carts match the query predicate, or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyCartsRequestBuilder) Head() *ByProjectKeyCartsRequestMethodHead {
	return &ByProjectKeyCartsRequestMethodHead{
		url:    fmt.Sprintf("/%s/carts", rb.projectKey),
		client: rb.client,
	}
}

/**
*
*	Creates a Cart in the Project.
*
*	If the referenced [ShippingMethod](ctp:api:type:ShippingMethod) in the [CartDraft](ctp:api:type:CartDraft) has a predicate that does not match, or if the Shipping Method is not active, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
*
*	Specific Error Codes:
*
*	- [DiscountCodeNonApplicable](ctp:api:type:DiscountCodeNonApplicableError)
*	- [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError)
*	- [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError)
*	- [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError)
*
 */
func (rb *ByProjectKeyCartsRequestBuilder) Post(body CartDraft) *ByProjectKeyCartsRequestMethodPost {
	return &ByProjectKeyCartsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/carts", rb.projectKey),
		client: rb.client,
	}
}
