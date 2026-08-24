package insights

// Generated file, please do not change!!!

type ByProjectKeyRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyRequestBuilder) InsightsConfiguration() *ByProjectKeyInsightsConfigurationRequestBuilder {
	return &ByProjectKeyInsightsConfigurationRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
