package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Retrieves the active ShippingMethods that can ship to the shipping address of the given Cart in a given [Store](ctp:api:type:Store).
*	Each ShippingMethod contains exactly one ShippingRate with the flag `isMatching` set to `true`.
*	This ShippingRate is used when the ShippingMethod is [added to the Cart](ctp:api:type:CartSetShippingMethodAction).
*	If a matching ShippingMethod has `isDefault` set to `true`, it is returned as the first item in the array.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/shipping-methods/matching-cart", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if an active ShippingMethod that can ship to the shipping address of the given Cart exists in the given [Store](ctp:api:type:Store). Returns a `200 OK` status if the ShippingMethod exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/shipping-methods/matching-cart", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
