package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMePasswordRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMePasswordRequestBuilder) Reset() *ByProjectKeyInStoreKeyByStoreKeyMePasswordResetRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMePasswordResetRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	Changing the password of the Customer produces the [CustomerPasswordUpdated](ctp:api:type:CustomerPasswordUpdatedMessage) Message with `reset=false`.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMePasswordRequestBuilder) Post(body MyCustomerChangePassword) *ByProjectKeyInStoreKeyByStoreKeyMePasswordRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMePasswordRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/password", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
