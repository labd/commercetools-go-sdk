package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestBuilder struct {
	projectKey  string
	storeKey    string
	orderNumber string
	client      *Client
}

/**
*	Returns an order by its order number from a specific Store.
*
*	If the order exists in the project but does not have the store field,
*	or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*	In case the orderNumber does not match the regular expression [a-zA-Z0-9_\-]+,
*	it should be provided in URL-encoded format.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders/order-number=%s", rb.projectKey, rb.storeKey, rb.orderNumber),
		client: rb.client,
	}
}

/**
*	Updates an order in the store specified by {storeKey}.
*	If the order exists in the project but does not have the store field,
*	or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*	In case the orderNumber does not match the regular expression [a-zA-Z0-9_\-]+,
*	it should be provided in URL-encoded format.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestBuilder) Post(body OrderUpdate) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders/order-number=%s", rb.projectKey, rb.storeKey, rb.orderNumber),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders/order-number=%s", rb.projectKey, rb.storeKey, rb.orderNumber),
		client: rb.client,
	}
}
