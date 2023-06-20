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
*	Retrieves logs for the given deployment.
 */
func (rb *ByProjectKeyDeploymentsByIDLogsRequestBuilder) Get() *ByProjectKeyDeploymentsByIDLogsRequestMethodGet {
	return &ByProjectKeyDeploymentsByIDLogsRequestMethodGet{
		url:    fmt.Sprintf("/%s/deployments/%s/logs", rb.projectKey, rb.id),
		client: rb.client,
	}
}
