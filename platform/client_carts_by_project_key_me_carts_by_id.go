package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeCartsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Returns a Cart for a given `id`. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Cart exists for a given `id`.
*	- If the Cart exists but does not have a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeCartsByIDRequestBuilder) Get() *ByProjectKeyMeCartsByIDRequestMethodGet {
	return &ByProjectKeyMeCartsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Cart exists for a given `id`. Returns a `200 OK` status if the Cart exists.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Cart exists for a given `id`.
*	- If the Cart exists but does not have a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeCartsByIDRequestBuilder) Head() *ByProjectKeyMeCartsByIDRequestMethodHead {
	return &ByProjectKeyMeCartsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates the Cart for a given `id`. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Cart exists for a given `id`.
*	- If the Cart exists but does not have a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeCartsByIDRequestBuilder) Post(body MyCartUpdate) *ByProjectKeyMeCartsByIDRequestMethodPost {
	return &ByProjectKeyMeCartsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes the Cart for a given `id`. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Cart exists for a given `id`.
*	- If the Cart exists but does not have a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeCartsByIDRequestBuilder) Delete() *ByProjectKeyMeCartsByIDRequestMethodDelete {
	return &ByProjectKeyMeCartsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
