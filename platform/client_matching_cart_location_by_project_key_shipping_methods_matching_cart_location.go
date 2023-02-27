package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyShippingMethodsMatchingCartLocationRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Retrieves all the ShippingMethods that can ship to the given [Location](/projects/zones#location)
*	with a `predicate` that matches the given Cart.
*	Each ShippingMethod contains exactly one ShippingRate with the flag `isMatching` set to `true`.
*	This ShippingRate is used when the ShippingMethod is [added to the Cart](ctp:api:type:CartSetShippingMethodAction).
*
 */
func (rb *ByProjectKeyShippingMethodsMatchingCartLocationRequestBuilder) Get() *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet {
	return &ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet{
		url:    fmt.Sprintf("/%s/shipping-methods/matching-cart-location", rb.projectKey),
		client: rb.client,
	}
}
