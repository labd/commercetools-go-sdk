package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyShippingMethodsMatchingOrdereditRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Retrieves the active ShippingMethods that can ship to the provided [Location](ctp:api:type:Location) for an [OrderEdit](ctp:api:type:OrderEdit).
*
*	If a matching ShippingMethod has `isDefault` set to `true`, it is returned as the first item in the array.
*
*	If the OrderEdit preview cannot be generated, an [EditPreviewFailed](ctp:api:type:EditPreviewFailedError) error is returned.
*
 */
func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestBuilder) Get() *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet {
	return &ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet{
		url:    fmt.Sprintf("/%s/shipping-methods/matching-orderedit", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if an active ShippingMethod that can ship to the provided [Location](ctp:api:type:Location) exists for the provided [OrderEdit](ctp:api:type:OrderEdit). Returns a `200 OK` status if the ShippingMethod exists or a [Not Found](/../api/errors#404-not-found) error otherwise.
 */
func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestBuilder) Head() *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHead {
	return &ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHead{
		url:    fmt.Sprintf("/%s/shipping-methods/matching-orderedit", rb.projectKey),
		client: rb.client,
	}
}
