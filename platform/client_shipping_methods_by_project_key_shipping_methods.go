// Generated file, please do not change!!!
package platform

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
*	Get ShippingMethods for a cart
 */
func (rb *ByProjectKeyShippingMethodsRequestBuilder) MatchingCart() *ByProjectKeyShippingMethodsMatchingCartRequestBuilder {
	return &ByProjectKeyShippingMethodsMatchingCartRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Get ShippingMethods for an order edit
 */
func (rb *ByProjectKeyShippingMethodsRequestBuilder) MatchingOrderedit() *ByProjectKeyShippingMethodsMatchingOrdereditRequestBuilder {
	return &ByProjectKeyShippingMethodsMatchingOrdereditRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Get ShippingMethods for a location
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

func (rb *ByProjectKeyShippingMethodsRequestBuilder) Post(body ShippingMethodDraft) *ByProjectKeyShippingMethodsRequestMethodPost {
	return &ByProjectKeyShippingMethodsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/shipping-methods", rb.projectKey),
		client: rb.client,
	}
}
