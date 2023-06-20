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
*	Retrieves logs for the given deployment.
 */
func (rb *ByProjectKeyDeploymentsKeyByKeyLogsRequestBuilder) Get() *ByProjectKeyDeploymentsKeyByKeyLogsRequestMethodGet {
	return &ByProjectKeyDeploymentsKeyByKeyLogsRequestMethodGet{
		url:    fmt.Sprintf("/%s/deployments/key=%s/logs", rb.projectKey, rb.key),
		client: rb.client,
	}
}
