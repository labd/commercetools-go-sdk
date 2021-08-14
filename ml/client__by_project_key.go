// Generated file, please do not change!!!
package ml

import ()

type ByProjectKeyRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Search for similar products using an image as search input.
*
 */
func (rb *ByProjectKeyRequestBuilder) ImageSearch() *ByProjectKeyImageSearchRequestBuilder {
	return &ByProjectKeyImageSearchRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) Recommendations() *ByProjectKeyRecommendationsRequestBuilder {
	return &ByProjectKeyRecommendationsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) MissingData() *ByProjectKeyMissingDataRequestBuilder {
	return &ByProjectKeyMissingDataRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) Similarities() *ByProjectKeySimilaritiesRequestBuilder {
	return &ByProjectKeySimilaritiesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
