package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDeploymentsKeyByKeyLogsRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Specific Error Codes:
*	- [DeploymentApplicationDoesNotExist](ctp:connect:type:DeploymentApplicationDoesNotExistError)
*	- [DeploymentLogInvalidDate](ctp:connect:type:DeploymentLogInvalidDateError)
*	- [DeploymentLogInvalidPageToken](ctp:connect:type:DeploymentLogInvalidPageTokenError)
*
 */
func (rb *ByProjectKeyDeploymentsKeyByKeyLogsRequestBuilder) Get() *ByProjectKeyDeploymentsKeyByKeyLogsRequestMethodGet {
	return &ByProjectKeyDeploymentsKeyByKeyLogsRequestMethodGet{
		url:    fmt.Sprintf("/%s/deployments/key=%s/logs", rb.projectKey, rb.key),
		client: rb.client,
	}
}
