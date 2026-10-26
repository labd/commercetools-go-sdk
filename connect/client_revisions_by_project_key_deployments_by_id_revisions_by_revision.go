package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDeploymentsByIDRevisionsByRevisionRequestBuilder struct {
	projectKey string
	id         string
	revision   string
	client     *Client
}

func (rb *ByProjectKeyDeploymentsByIDRevisionsByRevisionRequestBuilder) Rollback() *ByProjectKeyDeploymentsByIDRevisionsByRevisionRollbackRequestBuilder {
	return &ByProjectKeyDeploymentsByIDRevisionsByRevisionRollbackRequestBuilder{
		projectKey: rb.projectKey,
		id:         rb.id,
		revision:   rb.revision,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyDeploymentsByIDRevisionsByRevisionRequestBuilder) Get() *ByProjectKeyDeploymentsByIDRevisionsByRevisionRequestMethodGet {
	return &ByProjectKeyDeploymentsByIDRevisionsByRevisionRequestMethodGet{
		url:    fmt.Sprintf("/%s/deployments/%s/revisions/%s", rb.projectKey, rb.id, rb.revision),
		client: rb.client,
	}
}
