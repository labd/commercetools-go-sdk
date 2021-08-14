// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyReviewsRequestMethodPost struct {
	body   ReviewDraft
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyReviewsRequestMethodPostInput
}

func (r *ByProjectKeyReviewsRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyReviewsRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyReviewsRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyReviewsRequestMethodPost) WithQueryParams(input *ByProjectKeyReviewsRequestMethodPostInput) *ByProjectKeyReviewsRequestMethodPost {
	rb.query = input.Values()
	return rb
}

/**
*	Create Review
 */
func (rb *ByProjectKeyReviewsRequestMethodPost) Execute(ctx context.Context) (result *Review, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	resp, err := rb.client.post(
		ctx,
		rb.url,
		rb.query,
		data,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 201:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 400, 401, 403, 500, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, fmt.Errorf("StatusCode %d returend", resp.StatusCode)
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
