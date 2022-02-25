package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCartsReplicateRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsReplicateRequestBuilder) Post(body ReplicaCartDraft) *ByProjectKeyInStoreKeyByStoreKeyCartsReplicateRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsReplicateRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts/replicate", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
