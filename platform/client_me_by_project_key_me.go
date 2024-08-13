package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeRequestBuilder) EmailConfirm() *ByProjectKeyMeEmailConfirmRequestBuilder {
	return &ByProjectKeyMeEmailConfirmRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeRequestBuilder) Password() *ByProjectKeyMePasswordRequestBuilder {
	return &ByProjectKeyMePasswordRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeRequestBuilder) Signup() *ByProjectKeyMeSignupRequestBuilder {
	return &ByProjectKeyMeSignupRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeRequestBuilder) Login() *ByProjectKeyMeLoginRequestBuilder {
	return &ByProjectKeyMeLoginRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeRequestBuilder) ActiveCart() *ByProjectKeyMeActiveCartRequestBuilder {
	return &ByProjectKeyMeActiveCartRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	MyBusinessUnit creates and provides access to Business Units scoped to a specific user.
 */
func (rb *ByProjectKeyMeRequestBuilder) BusinessUnits() *ByProjectKeyMeBusinessUnitsRequestBuilder {
	return &ByProjectKeyMeBusinessUnitsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	A shopping cart holds product variants and can be ordered.
 */
func (rb *ByProjectKeyMeRequestBuilder) Carts() *ByProjectKeyMeCartsRequestBuilder {
	return &ByProjectKeyMeCartsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	An order can be created from a cart, usually after a checkout process has been completed.
 */
func (rb *ByProjectKeyMeRequestBuilder) Orders() *ByProjectKeyMeOrdersRequestBuilder {
	return &ByProjectKeyMeOrdersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	The My Payments endpoint creates and provides access to payments scoped to a specific user.
 */
func (rb *ByProjectKeyMeRequestBuilder) Payments() *ByProjectKeyMePaymentsRequestBuilder {
	return &ByProjectKeyMePaymentsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	The My Quote Requests endpoint creates and provides access to Quote Requests scoped to a specific user.
 */
func (rb *ByProjectKeyMeRequestBuilder) QuoteRequests() *ByProjectKeyMeQuoteRequestsRequestBuilder {
	return &ByProjectKeyMeQuoteRequestsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	The My Quote endpoint provides access to Quotes scoped to a specific user.
 */
func (rb *ByProjectKeyMeRequestBuilder) Quotes() *ByProjectKeyMeQuotesRequestBuilder {
	return &ByProjectKeyMeQuotesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	The My Shopping Lists endpoint creates and provides access to Shopping Lists scoped to a specific user.
 */
func (rb *ByProjectKeyMeRequestBuilder) ShoppingLists() *ByProjectKeyMeShoppingListsRequestBuilder {
	return &ByProjectKeyMeShoppingListsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Returns a Customer for a given Query Predicate. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Customer exists for the given Query Predicate.
*	- If a Customer exists for the given Query Predicate, but does not have an `id` value that matches the [customer:{id}](/scopes#customer_idid) scope.
*
 */
func (rb *ByProjectKeyMeRequestBuilder) Get() *ByProjectKeyMeRequestMethodGet {
	return &ByProjectKeyMeRequestMethodGet{
		url:    fmt.Sprintf("/%s/me", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Updates the Customer specified in the [customer:{id}](/scopes#customer_idid) scope. Returns a `200 OK` status if successful.
*
 */
func (rb *ByProjectKeyMeRequestBuilder) Post(body MyCustomerUpdate) *ByProjectKeyMeRequestMethodPost {
	return &ByProjectKeyMeRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Deletes the Customer specified in the [customer:{id}](/scopes#customer_idid) scope. Returns a `200 OK` status if successful.
 */
func (rb *ByProjectKeyMeRequestBuilder) Delete() *ByProjectKeyMeRequestMethodDelete {
	return &ByProjectKeyMeRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me", rb.projectKey),
		client: rb.client,
	}
}
