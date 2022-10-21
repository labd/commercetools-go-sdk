package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Changing the password of the Customer produces the [CustomerPasswordUpdated](ctp:api:type:CustomerPasswordUpdatedMessage) Message with `reset=false`.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordRequestBuilder) Post(body CustomerChangePassword) *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/password", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
