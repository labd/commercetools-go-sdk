// Generated file, please do not change!!!
package history

func (c *Client) WithProjectKeyValue(projectKey string) *ByProjectKeyRequestBuilder {
	return &ByProjectKeyRequestBuilder{
		projectKey: projectKey,
		client:     c,
	}
}
