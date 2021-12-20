// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyLoginRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Authenticate Customer (Sign In) in store
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyLoginRequestBuilder) Post(body CustomerSignin) *ByProjectKeyInStoreKeyByStoreKeyLoginRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyLoginRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/login", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
