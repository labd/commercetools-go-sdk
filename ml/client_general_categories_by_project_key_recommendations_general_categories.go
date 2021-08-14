// Generated file, please do not change!!!
package ml

import (
	"fmt"
)

type ByProjectKeyRecommendationsGeneralCategoriesRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	This endpoint takes arbitrary product names or image URLs and generates recommendations from a general set of categories, which cover a broad range of industries. The full list of supported categories can be found [here](https://docs.commercetools.com/category_recommendations_supported_categories.txt). These are independent of the categories that are actually defined in your project. The main  purpose of this API is to provide a quick way to test the behavior of the category recommendations engine for different names and images. In contrast to the [project-specific endpoint](https://docs.commercetools.com/http-api-projects-categoryrecommendations#project-specific-category-recommendations), this endpoint does not have [activation criteria](https://docs.commercetools.com/http-api-projects-categoryrecommendations#activating-the-api) and is enabled for all projects.
*
 */
func (rb *ByProjectKeyRecommendationsGeneralCategoriesRequestBuilder) Get() *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet {
	return &ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet{
		url:    fmt.Sprintf("/%s/recommendations/general-categories", rb.projectKey),
		client: rb.client,
	}
}
