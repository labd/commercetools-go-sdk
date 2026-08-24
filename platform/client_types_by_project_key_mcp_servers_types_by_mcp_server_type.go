package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMcpServersTypesByMcpServerTypeRequestBuilder struct {
	projectKey    string
	mcpServerType string
	client        *Client
}

/**
*	Returns a single [MCP Server type](ctp:api:type:McpServerType). When `majorVersion` is omitted, the latest major version is returned.
*
 */
func (rb *ByProjectKeyMcpServersTypesByMcpServerTypeRequestBuilder) Get() *ByProjectKeyMcpServersTypesByMcpServerTypeRequestMethodGet {
	return &ByProjectKeyMcpServersTypesByMcpServerTypeRequestMethodGet{
		url:    fmt.Sprintf("/%s/mcp-servers/types/%s", rb.projectKey, rb.mcpServerType),
		client: rb.client,
	}
}
