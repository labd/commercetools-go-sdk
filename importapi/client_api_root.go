// Generated file, please do not change!!!
package importapi

func (c *Client) WithProjectKeyValue(projectKey string) *ByProjectKeyRequestBuilder {
	return &ByProjectKeyRequestBuilder{
		projectKey: projectKey,
		client:     c,
	}
}
