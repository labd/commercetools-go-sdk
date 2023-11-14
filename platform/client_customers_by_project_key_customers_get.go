package platform

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

type ByProjectKeyCustomersRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCustomersRequestMethodGetInput
}

func (r *ByProjectKeyCustomersRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCustomersRequestMethodGetInput struct {
	Expand       []string
	Sort         []string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyCustomersRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	if input.WithTotal != nil {
		if *input.WithTotal {
			values.Add("withTotal", "true")
		} else {
			values.Add("withTotal", "false")
		}
	}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	for k, v := range input.PredicateVar {
		for _, x := range v {
			values.Set(k, x)
		}
	}
	return values
}

func (rb *ByProjectKeyCustomersRequestMethodGet) Expand(v []string) *ByProjectKeyCustomersRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomersRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCustomersRequestMethodGet) Sort(v []string) *ByProjectKeyCustomersRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomersRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyCustomersRequestMethodGet) Limit(v int) *ByProjectKeyCustomersRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomersRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyCustomersRequestMethodGet) Offset(v int) *ByProjectKeyCustomersRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomersRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyCustomersRequestMethodGet) WithTotal(v bool) *ByProjectKeyCustomersRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomersRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyCustomersRequestMethodGet) Where(v []string) *ByProjectKeyCustomersRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomersRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyCustomersRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyCustomersRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomersRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyCustomersRequestMethodGet) WithQueryParams(input ByProjectKeyCustomersRequestMethodGetInput) *ByProjectKeyCustomersRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCustomersRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyCustomersRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyCustomersRequestMethodGet) Execute(ctx context.Context) (result *CustomerPagedQueryResponse, err error) {
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
