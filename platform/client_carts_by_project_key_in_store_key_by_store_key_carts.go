package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder) WithCustomerId(customerId string) *ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdRequestBuilder{
		customerId: customerId,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder) Replicate() *ByProjectKeyInStoreKeyByStoreKeyCartsReplicateRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsReplicateRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	Queries carts in a specific Store.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCartsRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Creates a [Cart](ctp:api:type:Cart) in the Store specified by `storeKey`.
*	When using this endpoint the Cart's `store` field is always set to the store specified in the path parameter.
*	Creating a Cart can fail with an [InvalidOperationError](ctp:api:type:InvalidOperationError) if the referenced [ShippingMethod](ctp:api:type:ShippingMethod)
*	in the [CartDraft](ctp:api:type:CartDraft) has a predicate which does not match the Cart.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder) Post(body CartDraft) *ByProjectKeyInStoreKeyByStoreKeyCartsRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
