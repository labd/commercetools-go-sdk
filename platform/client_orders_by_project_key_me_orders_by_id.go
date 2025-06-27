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
*	Retrieves an Order with the provided `id` for the authenticated Customer or anonymous user. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Order exists for the provided `id`.
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
*	Checks if an Order exists with the provided `id` for the authenticated Customer or anonymous user. Returns a `200 OK` status if successful.
*
*	A [Not Found](/../api/errors#404-not-found) error is returned in the following scenarios:
*
*	- If no Order exists for the provided `id`.
*	- If the Order exists but does not have either a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeOrdersByIDRequestBuilder) Head() *ByProjectKeyMeOrdersByIDRequestMethodHead {
	return &ByProjectKeyMeOrdersByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
