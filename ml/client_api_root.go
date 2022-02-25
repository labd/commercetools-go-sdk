package ml

// Generated file, please do not change!!!

func (c *Client) WithProjectKey(projectKey string) *ByProjectKeyRequestBuilder {
	return &ByProjectKeyRequestBuilder{
		projectKey: projectKey,
		client:     c,
	}
}
