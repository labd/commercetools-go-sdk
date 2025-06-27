package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Use this method to create a password reset token for a Store-specific Customer during their [password reset process](/../api/customers-overview#customer-password-reset).
*
*	If the Customer exists in the Project but the `stores` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
*	Creating a password reset token for the Customer produces the [CustomerPasswordTokenCreated](ctp:api:type:CustomerPasswordTokenCreatedMessage) Message.
*	The Message will include the token's value, if the token's validity is 60 minutes or less.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenRequestBuilder) Post(body CustomerCreatePasswordResetToken) *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/password-token", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
