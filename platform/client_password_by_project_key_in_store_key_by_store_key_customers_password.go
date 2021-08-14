// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Change a customers password
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordRequestBuilder) Post(body CustomerChangePassword) *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/password", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
