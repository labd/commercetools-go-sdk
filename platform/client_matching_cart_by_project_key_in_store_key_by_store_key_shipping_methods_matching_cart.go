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
*	Retrieves all the ShippingMethods that can ship to the shipping address of the given Cart in a given Store.
*	Each ShippingMethod contains exactly one ShippingRate with the flag `isMatching` set to `true`.
*	This ShippingRate is used when the ShippingMethod is [added to the Cart](ctp:api:type:CartSetShippingMethodAction).
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/shipping-methods/matching-cart", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
