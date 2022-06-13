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

type ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGetInput
}

func (r *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGetInput struct {
	Limit         *int
	Offset        *int
	Staged        *bool
	ConfidenceMin *float64
	ConfidenceMax *float64
}

func (input *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	if input.Staged != nil {
		if *input.Staged {
			values.Add("staged", "true")
		} else {
			values.Add("staged", "false")
		}
	}
	if input.ConfidenceMin != nil {
		values.Add("confidenceMin", fmt.Sprintf("%f", *input.ConfidenceMin))
	}
	if input.ConfidenceMax != nil {
		values.Add("confidenceMax", fmt.Sprintf("%f", *input.ConfidenceMax))
	}
	return values
}

func (rb *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet) Limit(v int) *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet) Offset(v int) *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet) Staged(v bool) *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGetInput{}
	}
	rb.params.Staged = &v
	return rb
}

func (rb *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet) ConfidenceMin(v float64) *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGetInput{}
	}
	rb.params.ConfidenceMin = &v
	return rb
}

func (rb *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet) ConfidenceMax(v float64) *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGetInput{}
	}
	rb.params.ConfidenceMax = &v
	return rb
}

func (rb *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet) WithQueryParams(input ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGetInput) *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Response Representation: PagedQueryResult with a results array of ProjectCategoryrecommendation, sorted by confidence scores in descending order and the meta information of ProjectCategoryrecommendationMeta.
*
 */
func (rb *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet) Execute(ctx context.Context) (result *ProjectCategoryRecommendationPagedQueryResponse, err error) {
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
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
