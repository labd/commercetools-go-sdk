package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyDiscountCodesByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyDiscountCodesByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyDiscountCodesByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyDiscountCodesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/discount-codes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a DiscountCode exists with the provided `id`. Returns a `200` status if the DiscountCode exists, or a `404` status otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyDiscountCodesByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyDiscountCodesByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyDiscountCodesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/discount-codes/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
