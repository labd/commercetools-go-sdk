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
*	Retrieves the active ShippingMethods that can ship to the shipping address of the provided Cart in a [Store](ctp:api:type:Store).
*
*	The Cart must belong to the Store specified in the path. If no Cart exists for the given `cartId` in the specified Store, either because the Cart does not exist in the Project or because it exists but does not belong to that Store, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
*
*	The results include globally scoped ShippingMethods (those with an empty `stores` field) and ShippingMethods scoped to the Store specified in the path.
*
*	Each ShippingMethod contains exactly one ShippingRate with the flag `isMatching` set to `true`. This ShippingRate is used when the ShippingMethod is [added to the Cart](ctp:api:type:CartSetShippingMethodAction).
*
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
*	Checks if an active ShippingMethod that can ship to the shipping address of the provided Cart exists in a [Store](ctp:api:type:Store). Returns a `200 OK` status if the ShippingMethod exists or a [Not Found](/api/errors#404-not-found) error otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/shipping-methods/matching-cart", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
