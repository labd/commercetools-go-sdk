// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIdRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	Query orders
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Create Order
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder) Post(body MyOrderFromCartDraft) *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
