package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

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
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeRequestBuilder) EmailConfirm() *ByProjectKeyInStoreKeyByStoreKeyMeEmailConfirmRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeEmailConfirmRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeRequestBuilder) Password() *ByProjectKeyInStoreKeyByStoreKeyMePasswordRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMePasswordRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeRequestBuilder) Signup() *ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeRequestBuilder) Login() *ByProjectKeyInStoreKeyByStoreKeyMeLoginRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeLoginRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeRequestBuilder) Post(body MyCustomerUpdate) *ByProjectKeyInStoreKeyByStoreKeyMeRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyMeRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyMeRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
