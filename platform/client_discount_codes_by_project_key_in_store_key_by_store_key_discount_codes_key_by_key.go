package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyDiscountCodesKeyByKeyRequestBuilder struct {
	projectKey string
	storeKey   string
	key        string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyDiscountCodesKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyDiscountCodesKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyDiscountCodesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/discount-codes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a DiscountCode exists with the provided `key`. Returns a `200` status if the DiscountCode exists, or a `404` status otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyDiscountCodesKeyByKeyRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyDiscountCodesKeyByKeyRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyDiscountCodesKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/discount-codes/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
