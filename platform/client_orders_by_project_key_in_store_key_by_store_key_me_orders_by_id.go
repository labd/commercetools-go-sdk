package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	If the Order exists in the Project but does not have the `store` field, or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if an Order exists for a given `id`. Returns a `200 OK` status if the My Order exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
