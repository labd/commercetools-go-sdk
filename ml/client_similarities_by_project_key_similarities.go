package ml

// Generated file, please do not change!!!

type ByProjectKeySimilaritiesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeySimilaritiesRequestBuilder) Products() *ByProjectKeySimilaritiesProductsRequestBuilder {
	return &ByProjectKeySimilaritiesProductsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
