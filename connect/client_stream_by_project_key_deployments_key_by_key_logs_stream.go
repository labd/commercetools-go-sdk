package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Stream logs for the given deployment.
 */
func (rb *ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestBuilder) Get() *ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestMethodGet {
	return &ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestMethodGet{
		url:    fmt.Sprintf("/%s/deployments/key=%s/logs/stream", rb.projectKey, rb.key),
		client: rb.client,
	}
}
