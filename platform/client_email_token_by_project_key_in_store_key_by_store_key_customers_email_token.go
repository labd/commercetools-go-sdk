// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Create a Token for verifying the Customer's Email in store
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenRequestBuilder) Post(body CustomerCreateEmailToken) *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/email-token", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
