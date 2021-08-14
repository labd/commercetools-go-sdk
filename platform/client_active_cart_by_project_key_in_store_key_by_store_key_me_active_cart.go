// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/active-cart", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
