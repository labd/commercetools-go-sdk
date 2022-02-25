package history

// Generated file, please do not change!!!

func (c *Client) WithProjectKeyValue(projectKey string) *ByProjectKeyRequestBuilder {
	return &ByProjectKeyRequestBuilder{
		projectKey: projectKey,
		client:     c,
	}
}
