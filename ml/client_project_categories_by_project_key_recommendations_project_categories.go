package ml

// Generated file, please do not change!!!

import ()

type ByProjectKeyRecommendationsProjectCategoriesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyRecommendationsProjectCategoriesRequestBuilder) WithProductId(productId string) *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestBuilder {
	return &ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestBuilder{
		productId:  productId,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
