// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyReviewsKeyByKeyRequestMethodPost struct {
	body   ReviewUpdate
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyReviewsKeyByKeyRequestMethodPostInput
}

func (r *ByProjectKeyReviewsKeyByKeyRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyReviewsKeyByKeyRequestMethodPostInput struct {
	Expand *[]string
}

func (input *ByProjectKeyReviewsKeyByKeyRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
		}
	}
	return values
}

func (rb *ByProjectKeyReviewsKeyByKeyRequestMethodPost) WithQueryParams(input *ByProjectKeyReviewsKeyByKeyRequestMethodPostInput) *ByProjectKeyReviewsKeyByKeyRequestMethodPost {
	rb.query = input.Values()
	return rb
}

/**
*	Update Review by key
 */
func (rb *ByProjectKeyReviewsKeyByKeyRequestMethodPost) Execute() (result *Review, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	resp, err := rb.client.post(
		context.Background(),
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
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 409, 400, 401, 403, 500, 503:
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
