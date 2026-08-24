package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyDiscountCodesRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyDiscountCodesRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyDiscountCodesKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyDiscountCodesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyDiscountCodesRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyDiscountCodesByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyDiscountCodesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyDiscountCodesRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyDiscountCodesRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyDiscountCodesRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/discount-codes", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more DiscountCodes exist for the provided query predicate in a Store. Returns a `200` status if any DiscountCodes match the query predicate, or a `404` status otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyDiscountCodesRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyDiscountCodesRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyDiscountCodesRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/discount-codes", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
