// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIdRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Get Order by ID
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIdRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIdRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
