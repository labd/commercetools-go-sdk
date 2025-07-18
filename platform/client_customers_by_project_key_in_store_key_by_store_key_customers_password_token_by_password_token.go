package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestBuilder struct {
	projectKey    string
	storeKey      string
	passwordToken string
	client        *Client
}

/**
*	Use this method to retrieve a Store-specific Customer's details by using the password reset token during their [password reset process](/../api/customers-overview#customer-password-reset).
*
*	If the Customer exists in the Project but the `stores` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/password-token=%s", rb.projectKey, rb.storeKey, rb.passwordToken),
		client: rb.client,
	}
}
