package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeCartsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeCartsRequestBuilder) WithId(id string) *ByProjectKeyMeCartsByIDRequestBuilder {
	return &ByProjectKeyMeCartsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeCartsRequestBuilder) Replicate() *ByProjectKeyMeCartsReplicateRequestBuilder {
	return &ByProjectKeyMeCartsReplicateRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Returns all Carts that match a given Query Predicate and contain either a matching `customerId` or `anonymousId`.
*
 */
func (rb *ByProjectKeyMeCartsRequestBuilder) Get() *ByProjectKeyMeCartsRequestMethodGet {
	return &ByProjectKeyMeCartsRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/carts", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a Cart exists that matches a given Query Predicate and contains either a matching `customerId` or `anonymousId`. Returns a `200 OK` status if the Cart exists, or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
*
 */
func (rb *ByProjectKeyMeCartsRequestBuilder) Head() *ByProjectKeyMeCartsRequestMethodHead {
	return &ByProjectKeyMeCartsRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/carts", rb.projectKey),
		client: rb.client,
	}
}

/**
*
*	Creates a Cart for the Customer or anonymous user. The `customerId` or `anonymousId` field on the Cart is automatically set based on the [customer:{id}](/scopes#composable-commerce-oauth) or [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
*	Specific Error Codes:
*
*	- [DiscountCodeNonApplicable](ctp:api:type:DiscountCodeNonApplicableError)
*	- [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError)
*	- [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError)
*	- [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError)
*
 */
func (rb *ByProjectKeyMeCartsRequestBuilder) Post(body MyCartDraft) *ByProjectKeyMeCartsRequestMethodPost {
	return &ByProjectKeyMeCartsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/carts", rb.projectKey),
		client: rb.client,
	}
}
