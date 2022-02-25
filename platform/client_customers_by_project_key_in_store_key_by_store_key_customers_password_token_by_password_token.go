package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestBuilder struct {
	projectKey    string
	storeKey      string
	passwordToken string
	client        *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/customers/password-token=%s", rb.projectKey, rb.storeKey, rb.passwordToken),
		client: rb.client,
	}
}
