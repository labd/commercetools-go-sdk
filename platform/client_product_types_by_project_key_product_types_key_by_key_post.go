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

type ByProjectKeyProductTypesKeyByKeyRequestMethodPost struct {
	body    ProductTypeUpdate
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductTypesKeyByKeyRequestMethodPostInput
}

func (r *ByProjectKeyProductTypesKeyByKeyRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductTypesKeyByKeyRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyProductTypesKeyByKeyRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyProductTypesKeyByKeyRequestMethodPost) Expand(v []string) *ByProjectKeyProductTypesKeyByKeyRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductTypesKeyByKeyRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductTypesKeyByKeyRequestMethodPost) WithQueryParams(input ByProjectKeyProductTypesKeyByKeyRequestMethodPostInput) *ByProjectKeyProductTypesKeyByKeyRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductTypesKeyByKeyRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyProductTypesKeyByKeyRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyProductTypesKeyByKeyRequestMethodPost) Execute(ctx context.Context) (result *ProductType, err error) {
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
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 409, 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj

	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
