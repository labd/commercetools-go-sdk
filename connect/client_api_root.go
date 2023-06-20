package connect

// Generated file, please do not change!!!

/**
*	The Project endpoint is used to retrieve certain information from a project.
 */
func (c *Client) WithProjectKey(projectKey string) *ByProjectKeyRequestBuilder {
	return &ByProjectKeyRequestBuilder{
		projectKey: projectKey,
		client:     c,
	}
}
func (c *Client) Connectors() *ConnectorsRequestBuilder {
	return &ConnectorsRequestBuilder{
		client: c,
	}
}
