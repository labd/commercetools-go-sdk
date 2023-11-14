package connect

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDeploymentsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyDeploymentsKeyByKeyRequestBuilder) Logs() *ByProjectKeyDeploymentsKeyByKeyLogsRequestBuilder {
	return &ByProjectKeyDeploymentsKeyByKeyLogsRequestBuilder{
		projectKey: rb.projectKey,
		key:        rb.key,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyDeploymentsKeyByKeyRequestBuilder) Get() *ByProjectKeyDeploymentsKeyByKeyRequestMethodGet {
	return &ByProjectKeyDeploymentsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/deployments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyDeploymentsKeyByKeyRequestBuilder) Post(body DeploymentUpdate) *ByProjectKeyDeploymentsKeyByKeyRequestMethodPost {
	return &ByProjectKeyDeploymentsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/deployments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyDeploymentsKeyByKeyRequestBuilder) Delete() *ByProjectKeyDeploymentsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyDeploymentsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/deployments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
