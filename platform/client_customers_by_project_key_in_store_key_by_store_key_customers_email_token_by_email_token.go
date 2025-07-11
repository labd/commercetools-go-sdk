package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenByEmailTokenRequestBuilder struct {
	projectKey string
	storeKey   string
	emailToken string
	client     *Client
}

/**
*	Use this method to retrieve a Store-specific Customer's details by using the email token during their [email verification process](/../api/customers-overview#customer-email-verification).
*
*	If the Customer exists in the Project but the `stores` field references a different [Store](ctp:api:type:Store), this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenByEmailTokenRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenByEmailTokenRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenByEmailTokenRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/email-token=%s", rb.projectKey, rb.storeKey, rb.emailToken),
		client: rb.client,
	}
}
