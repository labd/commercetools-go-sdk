package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/carts", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestBuilder) Post(body MyCartDraft) *ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/carts", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
