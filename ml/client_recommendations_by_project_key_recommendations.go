package ml

// Generated file, please do not change!!!

type ByProjectKeyRecommendationsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyRecommendationsRequestBuilder) ProjectCategories() *ByProjectKeyRecommendationsProjectCategoriesRequestBuilder {
	return &ByProjectKeyRecommendationsProjectCategoriesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRecommendationsRequestBuilder) GeneralCategories() *ByProjectKeyRecommendationsGeneralCategoriesRequestBuilder {
	return &ByProjectKeyRecommendationsGeneralCategoriesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
