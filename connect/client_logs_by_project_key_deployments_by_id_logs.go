package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDeploymentsByIDLogsRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Specific Error Codes:
*	- [DeploymentApplicationDoesNotExist](ctp:connect:type:DeploymentApplicationDoesNotExistError)
*	- [DeploymentLogInvalidDate](ctp:connect:type:DeploymentLogInvalidDateError)
*	- [DeploymentLogInvalidPageToken](ctp:connect:type:DeploymentLogInvalidPageTokenError)
*
 */
func (rb *ByProjectKeyDeploymentsByIDLogsRequestBuilder) Get() *ByProjectKeyDeploymentsByIDLogsRequestMethodGet {
	return &ByProjectKeyDeploymentsByIDLogsRequestMethodGet{
		url:    fmt.Sprintf("/%s/deployments/%s/logs", rb.projectKey, rb.id),
		client: rb.client,
	}
}
