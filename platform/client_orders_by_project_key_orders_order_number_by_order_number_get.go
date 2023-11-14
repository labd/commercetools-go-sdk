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

type ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGetInput
}

func (r *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet) Expand(v []string) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet) WithQueryParams(input ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGetInput) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet) Execute(ctx context.Context) (result *Order, err error) {
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
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj

	case 500:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 502:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 503:
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
