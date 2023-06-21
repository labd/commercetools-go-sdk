package frontend

// Generated file, please do not change!!!

type FrontasticRequestBuilder struct {
	client *Client
}

func (rb *FrontasticRequestBuilder) Context() *FrontasticContextRequestBuilder {
	return &FrontasticContextRequestBuilder{
		client: rb.client,
	}
}
func (rb *FrontasticRequestBuilder) Page() *FrontasticPageRequestBuilder {
	return &FrontasticPageRequestBuilder{
		client: rb.client,
	}
}
func (rb *FrontasticRequestBuilder) Preview() *FrontasticPreviewRequestBuilder {
	return &FrontasticPreviewRequestBuilder{
		client: rb.client,
	}
}
func (rb *FrontasticRequestBuilder) Action() *FrontasticActionRequestBuilder {
	return &FrontasticActionRequestBuilder{
		client: rb.client,
	}
}
func (rb *FrontasticRequestBuilder) DataSource() *FrontasticDataSourceRequestBuilder {
	return &FrontasticDataSourceRequestBuilder{
		client: rb.client,
	}
}
