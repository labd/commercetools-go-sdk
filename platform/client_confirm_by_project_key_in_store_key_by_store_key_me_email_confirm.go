package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeEmailConfirmRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	This is the last step in the [email verification process of a Customer](/../api/projects/customers#email-verification-of-customer-in-store). Returns a `200 OK` status if successful.
*
*	After the email is verified, all email tokens issued previously through the [email verification flow](/../api/projects/customers#email-verification-of-customer) are invalidated. This invalidation of tokens is [eventually consistent](/../api/general-concepts#eventual-consistency).
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Customer exists with the `id` specified in the [customer:{id}](/scopes#composable-commerce-oauth) scope.
*	- If the Customer exists but is associated with a different Store than what is specified in the `manage_my_profile:{projectKey}:{storeKey}` scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeEmailConfirmRequestBuilder) Post(body MyCustomerEmailVerify) *ByProjectKeyInStoreKeyByStoreKeyMeEmailConfirmRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeEmailConfirmRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/email/confirm", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
