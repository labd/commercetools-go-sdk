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

type ByProjectKeyCustomerGroupsByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCustomerGroupsByIDRequestMethodGetInput
}

func (r *ByProjectKeyCustomerGroupsByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCustomerGroupsByIDRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyCustomerGroupsByIDRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCustomerGroupsByIDRequestMethodGet) Expand(v []string) *ByProjectKeyCustomerGroupsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomerGroupsByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCustomerGroupsByIDRequestMethodGet) WithQueryParams(input ByProjectKeyCustomerGroupsByIDRequestMethodGetInput) *ByProjectKeyCustomerGroupsByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCustomerGroupsByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyCustomerGroupsByIDRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyCustomerGroupsByIDRequestMethodGet) Execute(ctx context.Context) (result *CustomerGroup, err error) {
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
