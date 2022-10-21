package platform

// Generated file, please do not change!!!

type ByProjectKeyInStoreKeyByStoreKeyRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	A shopping cart holds product variants and can be ordered.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyRequestBuilder) Carts() *ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	An order can be created from a cart, usually after a checkout process has been completed.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyRequestBuilder) Orders() *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyRequestBuilder) Me() *ByProjectKeyInStoreKeyByStoreKeyMeRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	A Customer is a person purchasing products. Carts, Orders,
*	Comments and Reviews can be associated to a Customer.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyRequestBuilder) Customers() *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	Retrieves the authenticated customer.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyRequestBuilder) Login() *ByProjectKeyInStoreKeyByStoreKeyLoginRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyLoginRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyRequestBuilder) ShippingMethods() *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyShippingMethodsRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	shopping-lists e.g. for wishlist support
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyRequestBuilder) ShoppingLists() *ByProjectKeyInStoreKeyByStoreKeyShoppingListsRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyShoppingListsRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyRequestBuilder) ProductProjections() *ByProjectKeyInStoreKeyByStoreKeyProductProjectionsRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyProductProjectionsRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyRequestBuilder) ProductSelectionAssignments() *ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyProductSelectionAssignmentsRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
