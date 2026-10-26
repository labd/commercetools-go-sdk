package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDeploymentsKeyByKeyRevisionsRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyDeploymentsKeyByKeyRevisionsRequestBuilder) WithRevision(revision string) *ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRequestBuilder {
	return &ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRequestBuilder{
		revision:   revision,
		projectKey: rb.projectKey,
		key:        rb.key,
		client:     rb.client,
	}
}

/**
*	Retrieves the revisions of a Deployment, most recent first.
*
*	Only versions that were deployed successfully are returned, so every result is a valid rollback target.
*	At most 20 revisions are returned.
*
 */
func (rb *ByProjectKeyDeploymentsKeyByKeyRevisionsRequestBuilder) Get() *ByProjectKeyDeploymentsKeyByKeyRevisionsRequestMethodGet {
	return &ByProjectKeyDeploymentsKeyByKeyRevisionsRequestMethodGet{
		url:    fmt.Sprintf("/%s/deployments/key=%s/revisions", rb.projectKey, rb.key),
		client: rb.client,
	}
}
