package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDeploymentsByIDRevisionsRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyDeploymentsByIDRevisionsRequestBuilder) WithRevision(revision string) *ByProjectKeyDeploymentsByIDRevisionsByRevisionRequestBuilder {
	return &ByProjectKeyDeploymentsByIDRevisionsByRevisionRequestBuilder{
		revision:   revision,
		projectKey: rb.projectKey,
		id:         rb.id,
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
func (rb *ByProjectKeyDeploymentsByIDRevisionsRequestBuilder) Get() *ByProjectKeyDeploymentsByIDRevisionsRequestMethodGet {
	return &ByProjectKeyDeploymentsByIDRevisionsRequestMethodGet{
		url:    fmt.Sprintf("/%s/deployments/%s/revisions", rb.projectKey, rb.id),
		client: rb.client,
	}
}
