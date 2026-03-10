package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyShippingMethodsMatchingLocationRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Retrieves the active ShippingMethods that can ship to the provided [Location](/projects/zones#location).
*
*	The following applies:
*
*	- ShippingMethods that have a `predicate` defined are automatically disqualified.
*	- If the `currency` parameter is provided, then the ShippingMethods must also have a rate defined in the specified currency.
*	- Each ShippingMethod contains at least one ShippingRate with the flag `isMatching` set to `true`.
*	- If the `currency` parameter is provided, exactly one ShippingRate will contain it.
*	- If a matching ShippingMethod has `isDefault` set to `true`, it is returned as the first item in the array.
*
 */
func (rb *ByProjectKeyShippingMethodsMatchingLocationRequestBuilder) Get() *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet {
	return &ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet{
		url:    fmt.Sprintf("/%s/shipping-methods/matching-location", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if an active ShippingMethod that can ship to the provided [Location](ctp:api:type:Location) exists. Returns a `200 OK` status if the ShippingMethod exists or a [Not Found](/../api/errors#404-not-found) error otherwise.
 */
func (rb *ByProjectKeyShippingMethodsMatchingLocationRequestBuilder) Head() *ByProjectKeyShippingMethodsMatchingLocationRequestMethodHead {
	return &ByProjectKeyShippingMethodsMatchingLocationRequestMethodHead{
		url:    fmt.Sprintf("/%s/shipping-methods/matching-location", rb.projectKey),
		client: rb.client,
	}
}
