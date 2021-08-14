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

type ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGetInput
}

func (r *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
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
		if *input.Staged == true {
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

func (rb *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet) WithQueryParams(input *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGetInput) *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet {
	rb.query = input.Values()
	return rb
}

/**
*	Response Representation: PagedQueryResult with a results array of ProjectCategoryrecommendation, sorted by confidence scores in descending order and the meta information of ProjectCategoryrecommendationMeta.
*
 */
func (rb *ByProjectKeyRecommendationsProjectCategoriesByProductIdRequestMethodGet) Execute() (result *ProjectCategoryRecommendationPagedQueryResponse, err error) {
	resp, err := rb.client.get(
		context.Background(),
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
