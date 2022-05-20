package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeCartsReplicateRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeCartsReplicateRequestBuilder) Post(body ReplicaMyCartDraft) *ByProjectKeyMeCartsReplicateRequestMethodPost {
	return &ByProjectKeyMeCartsReplicateRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/carts/replicate", rb.projectKey),
		client: rb.client,
	}
}
