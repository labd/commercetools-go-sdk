package ml

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestBuilder struct {
	projectKey string
	productId  string
	client     *Client
}

/**
*	Response Representation: PagedQueryResult with a results array of ProjectCategoryrecommendation, sorted by confidence scores in descending order and the meta information of ProjectCategoryrecommendationMeta.
*
 */
func (rb *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestBuilder) Get() *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet {
	return &ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/recommendations/project-categories/%s", rb.projectKey, rb.productId),
		client: rb.client,
	}
}
