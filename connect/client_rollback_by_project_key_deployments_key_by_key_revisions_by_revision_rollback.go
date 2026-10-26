package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRollbackRequestBuilder struct {
	projectKey string
	key        string
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
func (rb *ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRollbackRequestBuilder) Post(body DeploymentRollback) *ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRollbackRequestMethodPost {
	return &ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRollbackRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/deployments/key=%s/revisions/%s/rollback", rb.projectKey, rb.key, rb.revision),
		client: rb.client,
	}
}
