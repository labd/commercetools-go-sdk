package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMcpServersByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyMcpServersByIDRequestBuilder) Get() *ByProjectKeyMcpServersByIDRequestMethodGet {
	return &ByProjectKeyMcpServersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/mcp-servers/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Specific Error Codes:
*	- [InvalidOperation](ctp:api:type:InvalidOperationError)
*
 */
func (rb *ByProjectKeyMcpServersByIDRequestBuilder) Post(body McpServerUpdate) *ByProjectKeyMcpServersByIDRequestMethodPost {
	return &ByProjectKeyMcpServersByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/mcp-servers/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMcpServersByIDRequestBuilder) Delete() *ByProjectKeyMcpServersByIDRequestMethodDelete {
	return &ByProjectKeyMcpServersByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/mcp-servers/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
