package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Retrieves the Customer's most recently modified [active Cart](ctp:api:type:CartState) in a Store. Returns a `200 OK` status if successful.
*
*	Carts with `Merchant` or `Quote` [CartOrigin](ctp:api:type:CartOrigin) are ignored.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no active Cart exists.
*	- If an active Cart exists but does not have a `store` specified, or the `store` field references a different Store.
*	- If an active Cart exists but does not have a `customerId` that matches the [customer:{id}](/scopes#customer_idid) scope, or `anonymousId` that matches the [anonymous_id:{id}](/scopes#anonymous_idid) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/active-cart", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if an active Cart exists in a Store. Returns `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no active Cart exists in a Store.
*	- If an active Cart exists but does not have a `store` specified, or the `store` field references a different Store.
*	- If an active Cart exists but does not contain a `customerId` that matches the [customer:{id}](/scopes#customer_idid) scope, or `anonymousId` that matches the [anonymous_id:{id}](/scopes#anonymous_idid) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/active-cart", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
