package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyShippingMethodsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyShippingMethodsRequestBuilder) WithKey(key string) *ByProjectKeyShippingMethodsKeyByKeyRequestBuilder {
	return &ByProjectKeyShippingMethodsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Get ShippingMethods for a Cart
 */
func (rb *ByProjectKeyShippingMethodsRequestBuilder) MatchingCart() *ByProjectKeyShippingMethodsMatchingCartRequestBuilder {
	return &ByProjectKeyShippingMethodsMatchingCartRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Get ShippingMethods for a Cart and Location
 */
func (rb *ByProjectKeyShippingMethodsRequestBuilder) MatchingCartLocation() *ByProjectKeyShippingMethodsMatchingCartLocationRequestBuilder {
	return &ByProjectKeyShippingMethodsMatchingCartLocationRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Get ShippingMethods for an Order Edit
 */
func (rb *ByProjectKeyShippingMethodsRequestBuilder) MatchingOrderedit() *ByProjectKeyShippingMethodsMatchingOrdereditRequestBuilder {
	return &ByProjectKeyShippingMethodsMatchingOrdereditRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Get ShippingMethods for a Location
 */
func (rb *ByProjectKeyShippingMethodsRequestBuilder) MatchingLocation() *ByProjectKeyShippingMethodsMatchingLocationRequestBuilder {
	return &ByProjectKeyShippingMethodsMatchingLocationRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyShippingMethodsRequestBuilder) WithId(id string) *ByProjectKeyShippingMethodsByIDRequestBuilder {
	return &ByProjectKeyShippingMethodsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyShippingMethodsRequestBuilder) Get() *ByProjectKeyShippingMethodsRequestMethodGet {
	return &ByProjectKeyShippingMethodsRequestMethodGet{
		url:    fmt.Sprintf("/%s/shipping-methods", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more ShippingMethods exist for the provided query predicate. Returns a `200 OK` status if any ShippingMethods match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyShippingMethodsRequestBuilder) Head() *ByProjectKeyShippingMethodsRequestMethodHead {
	return &ByProjectKeyShippingMethodsRequestMethodHead{
		url:    fmt.Sprintf("/%s/shipping-methods", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyShippingMethodsRequestBuilder) Post(body ShippingMethodDraft) *ByProjectKeyShippingMethodsRequestMethodPost {
	return &ByProjectKeyShippingMethodsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/shipping-methods", rb.projectKey),
		client: rb.client,
	}
}
