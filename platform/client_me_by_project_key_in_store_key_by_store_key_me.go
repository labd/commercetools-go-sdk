package platform

// Generated file, please do not change!!!

import ()

type ByProjectKeyInStoreKeyByStoreKeyMeRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	A shopping cart holds product variants and can be ordered.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeRequestBuilder) Carts() *ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	An order can be created from a order, usually after a checkout process has been completed.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeRequestBuilder) Orders() *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeRequestBuilder) ActiveCart() *ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeActiveCartRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	shopping-lists e.g. for wishlist support
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeRequestBuilder) ShoppingLists() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
