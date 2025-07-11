package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDeploymentsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyDeploymentsRequestBuilder) WithId(id string) *ByProjectKeyDeploymentsByIDRequestBuilder {
	return &ByProjectKeyDeploymentsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyDeploymentsRequestBuilder) WithKey(key string) *ByProjectKeyDeploymentsKeyByKeyRequestBuilder {
	return &ByProjectKeyDeploymentsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Specific Error Codes:
*	- [ConnectorStagedNotPreviewable](ctp:connect:type:ConnectorStagedNotPreviewableError)
*	- [DeploymentUnsupportedRegion](ctp:connect:type:DeploymentUnsupportedRegionError)
*	- [DeploymentUnknownApplicationConfiguration](ctp:connect:type:DeploymentUnknownApplicationConfigurationError)
*	- [DeploymentUnknownApplicationConfigurationKey](ctp:connect:type:DeploymentUnknownApplicationConfigurationKeyError)
*	- [DeploymentInvalidType](ctp:connect:type:DeploymentInvalidTypeError)
*	- [DeploymentProductionDeactivated](ctp:connect:type:DeploymentProductionDeactivatedError)
*	|
*	The `manage_api_clients:{projectKey}` scope is required if you use automatically generated API Client credentials.
*
 */
func (rb *ByProjectKeyDeploymentsRequestBuilder) Post(body DeploymentDraft) *ByProjectKeyDeploymentsRequestMethodPost {
	return &ByProjectKeyDeploymentsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/deployments", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Retrieves all deployments of a project key.
 */
func (rb *ByProjectKeyDeploymentsRequestBuilder) Get() *ByProjectKeyDeploymentsRequestMethodGet {
	return &ByProjectKeyDeploymentsRequestMethodGet{
		url:    fmt.Sprintf("/%s/deployments", rb.projectKey),
		client: rb.client,
	}
}
