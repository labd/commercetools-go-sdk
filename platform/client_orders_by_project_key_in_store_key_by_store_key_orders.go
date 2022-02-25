package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyOrdersRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestBuilder) WithOrderNumber(orderNumber string) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestBuilder{
		orderNumber: orderNumber,
		projectKey:  rb.projectKey,
		storeKey:    rb.storeKey,
		client:      rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	Queries orders in a specific Store.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Creates an order from a Cart from a specific Store.
*	When using this endpoint the orders's store field is always set to the store specified in the path parameter.
*	The cart must have a shipping address set before creating an order. When using the Platform TaxMode,
*	the shipping address is used for tax calculation.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestBuilder) Post(body OrderFromCartDraft) *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
