package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDeploymentsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyDeploymentsByIDRequestBuilder) Logs() *ByProjectKeyDeploymentsByIDLogsRequestBuilder {
	return &ByProjectKeyDeploymentsByIDLogsRequestBuilder{
		projectKey: rb.projectKey,
		id:         rb.id,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyDeploymentsByIDRequestBuilder) LogsStream() *ByProjectKeyDeploymentsByIDLogsStreamRequestBuilder {
	return &ByProjectKeyDeploymentsByIDLogsStreamRequestBuilder{
		projectKey: rb.projectKey,
		id:         rb.id,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyDeploymentsByIDRequestBuilder) Get() *ByProjectKeyDeploymentsByIDRequestMethodGet {
	return &ByProjectKeyDeploymentsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/deployments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyDeploymentsByIDRequestBuilder) Post(body DeploymentUpdate) *ByProjectKeyDeploymentsByIDRequestMethodPost {
	return &ByProjectKeyDeploymentsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/deployments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyDeploymentsByIDRequestBuilder) Delete() *ByProjectKeyDeploymentsByIDRequestMethodDelete {
	return &ByProjectKeyDeploymentsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/deployments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
