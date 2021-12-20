// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenByEmailTokenRequestBuilder struct {
	projectKey string
	storeKey   string
	emailToken string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenByEmailTokenRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenByEmailTokenRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersEmailTokenByEmailTokenRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/email-token=%s", rb.projectKey, rb.storeKey, rb.emailToken),
		client: rb.client,
	}
}
