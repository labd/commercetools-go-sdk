package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRequestBuilder struct {
	projectKey string
	key        string
	revision   string
	client     *Client
}

func (rb *ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRequestBuilder) Rollback() *ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRollbackRequestBuilder {
	return &ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRollbackRequestBuilder{
		projectKey: rb.projectKey,
		key:        rb.key,
		revision:   rb.revision,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRequestBuilder) Get() *ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRequestMethodGet {
	return &ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRequestMethodGet{
		url:    fmt.Sprintf("/%s/deployments/key=%s/revisions/%s", rb.projectKey, rb.key, rb.revision),
		client: rb.client,
	}
}
