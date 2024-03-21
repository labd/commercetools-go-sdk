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

type ByProjectKeyCustomerGroupsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCustomerGroupsRequestMethodGetInput
}

func (r *ByProjectKeyCustomerGroupsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCustomerGroupsRequestMethodGetInput struct {
	Where        []string
	Expand       []string
	Sort         []string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	PredicateVar map[string][]string
}

func (input *ByProjectKeyCustomerGroupsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
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
	for k, v := range input.PredicateVar {
		for _, x := range v {
			values.Add(k, x)
		}
	}
	return values
}

func (rb *ByProjectKeyCustomerGroupsRequestMethodGet) Where(v []string) *ByProjectKeyCustomerGroupsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomerGroupsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyCustomerGroupsRequestMethodGet) Expand(v []string) *ByProjectKeyCustomerGroupsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomerGroupsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCustomerGroupsRequestMethodGet) Sort(v []string) *ByProjectKeyCustomerGroupsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomerGroupsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyCustomerGroupsRequestMethodGet) Limit(v int) *ByProjectKeyCustomerGroupsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomerGroupsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyCustomerGroupsRequestMethodGet) Offset(v int) *ByProjectKeyCustomerGroupsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomerGroupsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyCustomerGroupsRequestMethodGet) WithTotal(v bool) *ByProjectKeyCustomerGroupsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomerGroupsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyCustomerGroupsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyCustomerGroupsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomerGroupsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyCustomerGroupsRequestMethodGet) WithQueryParams(input ByProjectKeyCustomerGroupsRequestMethodGetInput) *ByProjectKeyCustomerGroupsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCustomerGroupsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyCustomerGroupsRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyCustomerGroupsRequestMethodGet) Execute(ctx context.Context) (result *CustomerGroupPagedQueryResponse, err error) {
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
		if err != nil {
			return nil, err
		}
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
	case 404:
		return nil, ErrNotFound
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
