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
*	Checks if one or more Customers exist for the provided query predicate. Returns a `200` status if any Customers match the query predicate, or a `404` status otherwise.
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
*	If a Cart with a `store` field specified, the `store` field must reference the same [Store](ctp:api:type:Store) specified in the `{storeKey}` path parameter.
*
*	Allows converting an anonymous Cart to the active Cart of a Customer with [cart merge](/../api/customers-overview#cart-merge-during-sign-in-and-sign-up).
*	If the Customer has multiple active Carts, the anonymous Cart is merged into the most recently modified active Cart.
*
*	Creating a Customer produces the [CustomerCreated](ctp:api:type:CustomerCreatedMessage) Message. Simultaneously creating two Customers with the same email address can return a [LockedField](ctp:api:type:LockedFieldError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestBuilder) Post(body CustomerDraft) *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
