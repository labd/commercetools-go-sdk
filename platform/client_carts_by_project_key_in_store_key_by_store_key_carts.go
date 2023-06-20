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
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCartsRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	The `store` field in the created [Cart](ctp:api:type:Cart) is set to the Store specified by the `storeKey` path parameter.
*
*	Specific Error Codes: [CountryNotConfiguredInStore](ctp:api:type:CountryNotConfiguredInStoreError)
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder) Post(body CartDraft) *ByProjectKeyInStoreKeyByStoreKeyCartsRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
