package connect

// Generated file, please do not change!!!

type ConnectorsRequestBuilder struct {
	client *Client
}

func (rb *ConnectorsRequestBuilder) WithId(id string) *ConnectorsByIDRequestBuilder {
	return &ConnectorsByIDRequestBuilder{
		id:     id,
		client: rb.client,
	}
}
func (rb *ConnectorsRequestBuilder) WithKey(key string) *ConnectorsKeyByKeyRequestBuilder {
	return &ConnectorsKeyByKeyRequestBuilder{
		key:    key,
		client: rb.client,
	}
}
func (rb *ConnectorsRequestBuilder) Search() *ConnectorsSearchRequestBuilder {
	return &ConnectorsSearchRequestBuilder{
		client: rb.client,
	}
}
func (rb *ConnectorsRequestBuilder) Drafts() *ConnectorsDraftsRequestBuilder {
	return &ConnectorsDraftsRequestBuilder{
		client: rb.client,
	}
}
