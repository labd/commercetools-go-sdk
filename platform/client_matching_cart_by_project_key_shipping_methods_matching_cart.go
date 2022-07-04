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
*	Retrieves all the ShippingMethods that can ship to the shipping address of the given Cart.
*	Each ShippingMethod contains exactly one ShippingRate with the flag `isMatching` set to `true`.
*	This ShippingRate is used when the ShippingMethod is [added to the Cart](ctp:api:type:CartSetShippingMethodAction).
*
 */
func (rb *ByProjectKeyShippingMethodsMatchingCartRequestBuilder) Get() *ByProjectKeyShippingMethodsMatchingCartRequestMethodGet {
	return &ByProjectKeyShippingMethodsMatchingCartRequestMethodGet{
		url:    fmt.Sprintf("/%s/shipping-methods/matching-cart", rb.projectKey),
		client: rb.client,
	}
}
