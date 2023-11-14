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
*	If the Customer exists in the Project but the `stores` field references a different [Store](ctp:api:type:Store), this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenByEmailTokenRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenByEmailTokenRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenByEmailTokenRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/email-token=%s", rb.projectKey, rb.storeKey, rb.emailToken),
		client: rb.client,
	}
}
