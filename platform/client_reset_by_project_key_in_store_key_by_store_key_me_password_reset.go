package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMePasswordResetRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMePasswordResetRequestBuilder) Post(body MyCustomerResetPassword) *ByProjectKeyInStoreKeyByStoreKeyMePasswordResetRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMePasswordResetRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/password/reset", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
