package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestBuilder) Post(body MyCustomerDraft) *ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/signup", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
