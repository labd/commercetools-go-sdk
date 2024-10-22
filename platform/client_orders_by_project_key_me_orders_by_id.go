package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeOrdersByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Returns an Order for a given `id`. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Order exists for the given `id`.
*	- If the Order exists but does not have either a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeOrdersByIDRequestBuilder) Get() *ByProjectKeyMeOrdersByIDRequestMethodGet {
	return &ByProjectKeyMeOrdersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if an Order exists for a given `id`. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Order exists for the given `id`.
*	- If the Order exists but does not have either a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeOrdersByIDRequestBuilder) Head() *ByProjectKeyMeOrdersByIDRequestMethodHead {
	return &ByProjectKeyMeOrdersByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
