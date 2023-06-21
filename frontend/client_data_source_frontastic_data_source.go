package frontend

// Generated file, please do not change!!!

type FrontasticDataSourceRequestBuilder struct {
	client *Client
}

func (rb *FrontasticDataSourceRequestBuilder) WithIdentifierValue(identifier string) *FrontasticDataSourceByIdentifierRequestBuilder {
	return &FrontasticDataSourceByIdentifierRequestBuilder{
		identifier: identifier,
		client:     rb.client,
	}
}
