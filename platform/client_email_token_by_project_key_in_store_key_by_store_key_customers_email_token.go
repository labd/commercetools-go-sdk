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
*	If the Customer exists in the Project but the `stores` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenRequestBuilder) Post(body CustomerCreateEmailToken) *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/email-token", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
