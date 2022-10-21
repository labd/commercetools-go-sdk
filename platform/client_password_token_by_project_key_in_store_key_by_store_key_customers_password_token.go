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
*	If the Customer exists in the Project but the `stores` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenRequestBuilder) Post(body CustomerCreatePasswordResetToken) *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/password-token", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
