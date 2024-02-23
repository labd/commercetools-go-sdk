package frontend

// Generated file, please do not change!!!

type ApiRequestBuilder struct {
	client *Client
}

func (rb *ApiRequestBuilder) Build() *ApiBuildRequestBuilder {
	return &ApiBuildRequestBuilder{
		client: rb.client,
	}
}
