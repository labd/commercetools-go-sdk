package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeEmailConfirmRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeEmailConfirmRequestBuilder) Post(body CustomerEmailVerify) *ByProjectKeyInStoreKeyByStoreKeyMeEmailConfirmRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeEmailConfirmRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/email/confirm", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
