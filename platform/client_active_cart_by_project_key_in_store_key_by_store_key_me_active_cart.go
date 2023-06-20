package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Retrieves the Customer's most recently modified active Cart in the Store specified by the `storeKey` path parameter.
*
*	Carts with `Merchant` or `Quote` [CartOrigin](ctp:api:type:CartOrigin) are ignored.
*
*	If no active Cart exists, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/active-cart", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
