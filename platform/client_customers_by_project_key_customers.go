package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCustomersRequestBuilder) WithPasswordToken(passwordToken string) *ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestBuilder {
	return &ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestBuilder{
		passwordToken: passwordToken,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) WithEmailToken(emailToken string) *ByProjectKeyCustomersEmailTokenByEmailTokenRequestBuilder {
	return &ByProjectKeyCustomersEmailTokenByEmailTokenRequestBuilder{
		emailToken: emailToken,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) EmailToken() *ByProjectKeyCustomersEmailTokenRequestBuilder {
	return &ByProjectKeyCustomersEmailTokenRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) EmailConfirm() *ByProjectKeyCustomersEmailConfirmRequestBuilder {
	return &ByProjectKeyCustomersEmailConfirmRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) Password() *ByProjectKeyCustomersPasswordRequestBuilder {
	return &ByProjectKeyCustomersPasswordRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) PasswordReset() *ByProjectKeyCustomersPasswordResetRequestBuilder {
	return &ByProjectKeyCustomersPasswordResetRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Produces the [CustomerPasswordTokenCreated](ctp:api:type:CustomerPasswordTokenCreatedMessage) Message.
 */
func (rb *ByProjectKeyCustomersRequestBuilder) PasswordToken() *ByProjectKeyCustomersPasswordTokenRequestBuilder {
	return &ByProjectKeyCustomersPasswordTokenRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) WithKey(key string) *ByProjectKeyCustomersKeyByKeyRequestBuilder {
	return &ByProjectKeyCustomersKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) WithId(id string) *ByProjectKeyCustomersByIDRequestBuilder {
	return &ByProjectKeyCustomersByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	This endpoint provides high-performance search queries over Customers.
*
 */
func (rb *ByProjectKeyCustomersRequestBuilder) Search() *ByProjectKeyCustomersSearchRequestBuilder {
	return &ByProjectKeyCustomersSearchRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) SearchIndexingStatus() *ByProjectKeyCustomersSearchIndexingStatusRequestBuilder {
	return &ByProjectKeyCustomersSearchIndexingStatusRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) Get() *ByProjectKeyCustomersRequestMethodGet {
	return &ByProjectKeyCustomersRequestMethodGet{
		url:    fmt.Sprintf("/%s/customers", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Customers exist for the provided query predicate. Returns a `200 OK` status if any Customers match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyCustomersRequestBuilder) Head() *ByProjectKeyCustomersRequestMethodHead {
	return &ByProjectKeyCustomersRequestMethodHead{
		url:    fmt.Sprintf("/%s/customers", rb.projectKey),
		client: rb.client,
	}
}

/**
*	If the `anonymousCart` field is set on the [CustomerDraft](ctp:api:type:CustomerDraft), then the newly created Customer will be assigned to that [Cart](ctp:api:type:Cart).
*	Similarly, if the `anonymousId` field is set, the Customer will be set on all [Carts](ctp:api:type:Cart), [Orders](ctp:api:type:Order), [ShoppingLists](ctp:api:type:ShoppingList) and [Payments](ctp:api:type:Payment) with the same `anonymousId`.
*
*	Creating a Customer produces the [CustomerCreated](ctp:api:type:CustomerCreatedMessage) Message. Simultaneously creating two Customers with the same email address can return a [LockedField](ctp:api:type:LockedFieldError) error.
*
 */
func (rb *ByProjectKeyCustomersRequestBuilder) Post(body CustomerDraft) *ByProjectKeyCustomersRequestMethodPost {
	return &ByProjectKeyCustomersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers", rb.projectKey),
		client: rb.client,
	}
}
