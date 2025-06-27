package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Use this method to create an email token for a Store-specific Customer during their [email verification process](/../api/customers-overview#customer-email-verification).
*
*	If the Customer exists in the Project but the `stores` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
*	Creating an email token for the Customer produces the [CustomerEmailTokenCreated](ctp:api:type:CustomerEmailTokenCreatedMessage) Message.
*	The Message will include the token's value, if the token's validity is 60 minutes or less.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenRequestBuilder) Post(body CustomerCreateEmailToken) *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/email-token", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
