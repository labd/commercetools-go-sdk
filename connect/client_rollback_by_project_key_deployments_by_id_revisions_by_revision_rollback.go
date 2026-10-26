package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDeploymentsByIDRevisionsByRevisionRollbackRequestBuilder struct {
	projectKey string
	id         string
	revision   string
	client     *Client
}

/**
*	Restores the configuration of the revision onto the Deployment.
*	A new Deployment version is created from the revision and deployed.
*
*	The `manage_api_clients:{projectKey}` scope is required if the connector version uses commercetools managed API Client credentials.
*
*	Specific Error Codes:
*	- [DeploymentRollbackNotAllowed](ctp:connect:type:DeploymentRollbackNotAllowedError)
*	- [DeploymentInvalidStatusTransition](ctp:connect:type:DeploymentInvalidStatusTransitionError)
*	- [ConcurrentModification](ctp:connect:type:ConcurrentModificationError)
*
 */
func (rb *ByProjectKeyDeploymentsByIDRevisionsByRevisionRollbackRequestBuilder) Post(body DeploymentRollback) *ByProjectKeyDeploymentsByIDRevisionsByRevisionRollbackRequestMethodPost {
	return &ByProjectKeyDeploymentsByIDRevisionsByRevisionRollbackRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/deployments/%s/revisions/%s/rollback", rb.projectKey, rb.id, rb.revision),
		client: rb.client,
	}
}
