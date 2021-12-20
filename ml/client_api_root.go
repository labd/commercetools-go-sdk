// Generated file, please do not change!!!
package ml

func (c *Client) WithProjectKey(projectKey string) *ByProjectKeyRequestBuilder {
	return &ByProjectKeyRequestBuilder{
		projectKey: projectKey,
		client:     c,
	}
}
