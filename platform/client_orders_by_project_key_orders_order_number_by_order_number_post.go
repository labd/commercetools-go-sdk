// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPost struct {
	body    OrderUpdate
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPostInput
}

func (r *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPost) Expand(v []string) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPost) WithQueryParams(input ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPostInput) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPost) Execute(ctx context.Context) (result *Order, err error) {
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
	case 404:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
		}
		return nil, result
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
