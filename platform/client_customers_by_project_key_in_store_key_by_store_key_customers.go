package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) WithPasswordToken(passwordToken string) *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestBuilder{
		passwordToken: passwordToken,
		projectKey:    rb.projectKey,
		storeKey:      rb.storeKey,
		client:        rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) WithEmailToken(emailToken string) *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenByEmailTokenRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenByEmailTokenRequestBuilder{
		emailToken: emailToken,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) EmailToken() *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) EmailConfirm() *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailConfirmRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersEmailConfirmRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) Password() *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) PasswordReset() *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordResetRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) PasswordToken() *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if a Customer exists for a given Query Predicate. Returns a `200 OK` status if any Customers match the Query Predicate or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	When using this endpoint, if omitted, the Customer `stores` field is set to the [Store](ctp:api:type:Store) specified in the path parameter.
*
*	If the `anonymousCart` field is set on the [CustomerDraft](ctp:api:type:CustomerDraft), then the newly created Customer will be assigned to that [Cart](ctp:api:type:Cart).
*	Similarly, if the `anonymousId` field is set, the Customer will be set on all [Carts](ctp:api:type:Cart), [Orders](ctp:api:type:Order), [ShoppingLists](ctp:api:type:ShoppingList) and [Payments](ctp:api:type:Payment) with the same `anonymousId`.
*	If a Cart with a `store` field specified, the `store` field must reference the same [Store](ctp:api:type:Store) specified in the `{storeKey}` path parameter.
*	Creating a Customer produces the [CustomerCreated](ctp:api:type:CustomerCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) Post(body CustomerDraft) *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
