package ml

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGetInput
}

func (r *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
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

func (rb *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet) ProductImageUrl(v string) *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGetInput{}
	}
	rb.params.ProductImageUrl = &v
	return rb
}

func (rb *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet) ProductName(v string) *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGetInput{}
	}
	rb.params.ProductName = v
	return rb
}

func (rb *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet) Limit(v int) *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet) Offset(v int) *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet) ConfidenceMin(v float64) *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGetInput{}
	}
	rb.params.ConfidenceMin = &v
	return rb
}

func (rb *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet) ConfidenceMax(v float64) *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGetInput{}
	}
	rb.params.ConfidenceMax = &v
	return rb
}

func (rb *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet) WithQueryParams(input ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGetInput) *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	This endpoint takes arbitrary product names or image URLs and generates recommendations from a general set of categories, which cover a broad range of industries. The full list of supported categories can be found [here](https://docs.commercetools.com/category_recommendations_supported_categories.txt). These are independent of the categories that are actually defined in your project. The main  purpose of this API is to provide a quick way to test the behavior of the category recommendations engine for different names and images. In contrast to the [project-specific endpoint](https://docs.commercetools.com/http-api-projects-categoryrecommendations#project-specific-category-recommendations), this endpoint does not have [activation criteria](https://docs.commercetools.com/http-api-projects-categoryrecommendations#activating-the-api) and is enabled for all projects.
*
 */
func (rb *ByProjectKeyRecommendationsGeneralCategoriesRequestMethodGet) Execute(ctx context.Context) (result *GeneralCategoryRecommendationPagedQueryResponse, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.get(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	default:
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
