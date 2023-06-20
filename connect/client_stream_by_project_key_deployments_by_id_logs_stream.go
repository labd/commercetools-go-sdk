package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDeploymentsByIDLogsStreamRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Stream logs for the given deployment.
 */
func (rb *ByProjectKeyDeploymentsByIDLogsStreamRequestBuilder) Get() *ByProjectKeyDeploymentsByIDLogsStreamRequestMethodGet {
	return &ByProjectKeyDeploymentsByIDLogsStreamRequestMethodGet{
		url:    fmt.Sprintf("/%s/deployments/%s/logs/stream", rb.projectKey, rb.id),
		client: rb.client,
	}
}
