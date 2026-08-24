package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMcpServersKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyMcpServersKeyByKeyRequestBuilder) Get() *ByProjectKeyMcpServersKeyByKeyRequestMethodGet {
	return &ByProjectKeyMcpServersKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/mcp-servers/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Specific Error Codes:
*	- [InvalidOperation](ctp:api:type:InvalidOperationError)
*
 */
func (rb *ByProjectKeyMcpServersKeyByKeyRequestBuilder) Post(body McpServerUpdate) *ByProjectKeyMcpServersKeyByKeyRequestMethodPost {
	return &ByProjectKeyMcpServersKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/mcp-servers/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMcpServersKeyByKeyRequestBuilder) Delete() *ByProjectKeyMcpServersKeyByKeyRequestMethodDelete {
	return &ByProjectKeyMcpServersKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/mcp-servers/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
