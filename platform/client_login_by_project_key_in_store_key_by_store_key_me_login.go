package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeLoginRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeLoginRequestBuilder) Post(body MyCustomerSignin) *ByProjectKeyInStoreKeyByStoreKeyMeLoginRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeLoginRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/login", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
