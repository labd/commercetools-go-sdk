package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersEmailConfirmRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Use this method to verify a Store-specific Customer's email during their [email verification process](/../api/customers-overview#customer-email-verification).
*
*	Verifying the email of the Customer produces the [CustomerEmailVerified](ctp:api:type:CustomerEmailVerifiedMessage) Message.
*
*	If the Customer exists in the Project but the `stores` field references a different [Store](ctp:api:type:Store), this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
*	After the email is verified, all email tokens issued previously through the [email verification flow](/../api/projects/customers#email-verification-of-customer) are invalidated. This invalidation of tokens is [eventually consistent](/../api/general-concepts#eventual-consistency).
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailConfirmRequestBuilder) Post(body CustomerEmailVerify) *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailConfirmRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersEmailConfirmRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/email/confirm", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
