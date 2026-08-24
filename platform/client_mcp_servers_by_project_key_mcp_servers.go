package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMcpServersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMcpServersRequestBuilder) Types() *ByProjectKeyMcpServersTypesRequestBuilder {
	return &ByProjectKeyMcpServersTypesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMcpServersRequestBuilder) WithKey(key string) *ByProjectKeyMcpServersKeyByKeyRequestBuilder {
	return &ByProjectKeyMcpServersKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMcpServersRequestBuilder) WithId(id string) *ByProjectKeyMcpServersByIDRequestBuilder {
	return &ByProjectKeyMcpServersByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMcpServersRequestBuilder) Get() *ByProjectKeyMcpServersRequestMethodGet {
	return &ByProjectKeyMcpServersRequestMethodGet{
		url:    fmt.Sprintf("/%s/mcp-servers", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a new MCP Server for the Project.
*
*	Specific Error Codes:
*	- [DuplicateField](ctp:api:type:DuplicateFieldError)
*
 */
func (rb *ByProjectKeyMcpServersRequestBuilder) Post(body McpServerDraft) *ByProjectKeyMcpServersRequestMethodPost {
	return &ByProjectKeyMcpServersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/mcp-servers", rb.projectKey),
		client: rb.client,
	}
}
