// Generated file, please do not change!!!
package ml

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
)

type ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGetInput
}

func (r *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGetInput struct {
	ProductImageUrl *string
	ProductName     string
	Limit           *int
	Offset          *int
	ConfidenceMin   *float64
	ConfidenceMax   *float64
}

func (input *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.ProductImageUrl != nil {
		values.Add("productImageUrl", fmt.Sprintf("%v", *input.ProductImageUrl))
	}
	values.Add("productName", fmt.Sprintf("%v", input.ProductName))
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	if input.ConfidenceMin != nil {
		values.Add("confidenceMin", fmt.Sprintf("%f", *input.ConfidenceMin))
	}
	if input.ConfidenceMax != nil {
		values.Add("confidenceMax", fmt.Sprintf("%f", *input.ConfidenceMax))
	}
	return values
}

func (rb *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet) WithQueryParams(input *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGetInput) *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet {
	rb.query = input.Values()
	return rb
}

/**
*	This endpoint takes arbitrary product names or image URLs and generates recommendations from a general set of categories, which cover a broad range of industries. The full list of supported categories can be found [here](https://docs.commercetools.com/category_recommendations_supported_categories.txt). These are independent of the categories that are actually defined in your project. The main  purpose of this API is to provide a quick way to test the behavior of the category recommendations engine for different names and images. In contrast to the [project-specific endpoint](https://docs.commercetools.com/http-api-projects-categoryrecommendations#project-specific-category-recommendations), this endpoint does not have [activation criteria](https://docs.commercetools.com/http-api-projects-categoryrecommendations#activating-the-api) and is enabled for all projects.
*
 */
func (rb *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet) Execute(ctx context.Context) (result *GeneralCategoryRecommendationPagedQueryResponse, err error) {
	resp, err := rb.client.get(
		ctx,
		rb.url,
		rb.query,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
