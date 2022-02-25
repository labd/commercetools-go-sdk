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

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
