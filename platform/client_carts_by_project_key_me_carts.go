package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeCartsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeCartsRequestBuilder) WithKey(key string) *ByProjectKeyMeCartsKeyByKeyRequestBuilder {
	return &ByProjectKeyMeCartsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeCartsRequestBuilder) WithId(id string) *ByProjectKeyMeCartsByIDRequestBuilder {
	return &ByProjectKeyMeCartsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeCartsRequestBuilder) Replicate() *ByProjectKeyMeCartsReplicateRequestBuilder {
	return &ByProjectKeyMeCartsReplicateRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeCartsRequestBuilder) Get() *ByProjectKeyMeCartsRequestMethodGet {
	return &ByProjectKeyMeCartsRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/carts", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeCartsRequestBuilder) Post(body MyCartDraft) *ByProjectKeyMeCartsRequestMethodPost {
	return &ByProjectKeyMeCartsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/carts", rb.projectKey),
		client: rb.client,
	}
}
