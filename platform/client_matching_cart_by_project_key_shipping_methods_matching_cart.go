package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyShippingMethodsMatchingCartRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Retrieves the active ShippingMethods that can ship to the shipping address of the given Cart.
*	Each ShippingMethod contains exactly one ShippingRate with the flag `isMatching` set to `true`.
*	This ShippingRate is used when the ShippingMethod is [added to the Cart](ctp:api:type:CartSetShippingMethodAction).
*	If a matching ShippingMethod has `isDefault` set to `true`, it is returned as the first item in the array.
*
 */
func (rb *ByProjectKeyShippingMethodsMatchingCartRequestBuilder) Get() *ByProjectKeyShippingMethodsMatchingCartRequestMethodGet {
	return &ByProjectKeyShippingMethodsMatchingCartRequestMethodGet{
		url:    fmt.Sprintf("/%s/shipping-methods/matching-cart", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if an active ShippingMethod exists for the given Cart. If a matching ShippingMethod has `isDefault` set to `true`, it is returned as the first item in the array. Returns a `200 OK` status if the ShippingMethod exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyShippingMethodsMatchingCartRequestBuilder) Head() *ByProjectKeyShippingMethodsMatchingCartRequestMethodHead {
	return &ByProjectKeyShippingMethodsMatchingCartRequestMethodHead{
		url:    fmt.Sprintf("/%s/shipping-methods/matching-cart", rb.projectKey),
		client: rb.client,
	}
}
