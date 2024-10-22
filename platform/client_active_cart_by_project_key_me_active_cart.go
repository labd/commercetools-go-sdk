package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeActiveCartRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Retrieves the Customer's most recently modified [active Cart](ctp:api:type:CartState). Returns a `200 OK` status if successful.
*
*	Carts with `Merchant` or `Quote` [CartOrigin](ctp:api:type:CartOrigin) are ignored.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no active Cart exists.
*	- If an active Cart exists but does not have a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeActiveCartRequestBuilder) Get() *ByProjectKeyMeActiveCartRequestMethodGet {
	return &ByProjectKeyMeActiveCartRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/active-cart", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if an active Cart exists. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no active Cart exists.
*	- If an active Cart exists but does not have a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeActiveCartRequestBuilder) Head() *ByProjectKeyMeActiveCartRequestMethodHead {
	return &ByProjectKeyMeActiveCartRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/active-cart", rb.projectKey),
		client: rb.client,
	}
}
