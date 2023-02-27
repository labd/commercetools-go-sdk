package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyReviewsRequestMethodPost struct {
	body    ReviewDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyReviewsRequestMethodPostInput
}

func (r *ByProjectKeyReviewsRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
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

func (rb *ByProjectKeyReviewsRequestMethodPost) Expand(v []string) *ByProjectKeyReviewsRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyReviewsRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyReviewsRequestMethodPost) WithQueryParams(input ByProjectKeyReviewsRequestMethodPostInput) *ByProjectKeyReviewsRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyReviewsRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyReviewsRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyReviewsRequestMethodPost) Execute(ctx context.Context) (result *Review, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.post(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		data,
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
	case 201:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, ErrNotFound
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
