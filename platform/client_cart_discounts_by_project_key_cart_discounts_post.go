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

type ByProjectKeyCartDiscountsRequestMethodPost struct {
	body    CartDiscountDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCartDiscountsRequestMethodPostInput
}

func (r *ByProjectKeyCartDiscountsRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCartDiscountsRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyCartDiscountsRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCartDiscountsRequestMethodPost) Expand(v []string) *ByProjectKeyCartDiscountsRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyCartDiscountsRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCartDiscountsRequestMethodPost) WithQueryParams(input ByProjectKeyCartDiscountsRequestMethodPostInput) *ByProjectKeyCartDiscountsRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCartDiscountsRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyCartDiscountsRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyCartDiscountsRequestMethodPost) Execute(ctx context.Context) (result *CartDiscount, err error) {
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
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
		}
		return nil, result
	default:
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
