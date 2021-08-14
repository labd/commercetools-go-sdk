// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCartsReplicateRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCartsReplicateRequestBuilder) Post(body ReplicaCartDraft) *ByProjectKeyCartsReplicateRequestMethodPost {
	return &ByProjectKeyCartsReplicateRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/carts/replicate", rb.projectKey),
		client: rb.client,
	}
}
