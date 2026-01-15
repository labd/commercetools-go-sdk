package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	If used with an optional [access token for an anonymous session](ctp:api:type:AnonymousSession), all Orders and Carts that belong to the `anonymousId` are assigned to the newly created Customer.
*
*	If omitted in the request body, the [Customer](ctp:api:type:Customer) `stores` field is set to the [Store](ctp:api:type:Store) specified in the path parameter.
*
*	If the Customer has multiple active Carts, the anonymous Cart is [merged](/../api/customers-overview#cart-merge-during-sign-in-and-sign-up) into the most recently modified active Cart.
*
*	Creating a Customer produces the [CustomerCreated](ctp:api:type:CustomerCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestBuilder) Post(body MyCustomerDraft) *ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/signup", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
