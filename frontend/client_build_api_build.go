package frontend

// Generated file, please do not change!!!

type ApiBuildRequestBuilder struct {
	client *Client
}

func (rb *ApiBuildRequestBuilder) Upload() *ApiBuildUploadRequestBuilder {
	return &ApiBuildUploadRequestBuilder{
		client: rb.client,
	}
}
