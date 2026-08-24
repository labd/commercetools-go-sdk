package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMcpServersTypesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMcpServersTypesRequestBuilder) WithMcpServerType(mcpServerType string) *ByProjectKeyMcpServersTypesByMcpServerTypeRequestBuilder {
	return &ByProjectKeyMcpServersTypesByMcpServerTypeRequestBuilder{
		mcpServerType: mcpServerType,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}

/**
*	Returns the catalog of available [MCP Server types](ctp:api:type:McpServerType), each with its major versions and the tools you can enable for them.
*
 */
func (rb *ByProjectKeyMcpServersTypesRequestBuilder) Get() *ByProjectKeyMcpServersTypesRequestMethodGet {
	return &ByProjectKeyMcpServersTypesRequestMethodGet{
		url:    fmt.Sprintf("/%s/mcp-servers/types", rb.projectKey),
		client: rb.client,
	}
}
