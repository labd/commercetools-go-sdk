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

type ByProjectKeyTypesRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyTypesRequestMethodGetInput
}

func (r *ByProjectKeyTypesRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyTypesRequestMethodGetInput struct {
	Expand       []string
	Sort         []string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyTypesRequestMethodGetInput) Values() url.Values {
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

func (rb *ByProjectKeyTypesRequestMethodGet) Expand(v []string) *ByProjectKeyTypesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyTypesRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyTypesRequestMethodGet) Sort(v []string) *ByProjectKeyTypesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyTypesRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyTypesRequestMethodGet) Limit(v int) *ByProjectKeyTypesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyTypesRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyTypesRequestMethodGet) Offset(v int) *ByProjectKeyTypesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyTypesRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyTypesRequestMethodGet) WithTotal(v bool) *ByProjectKeyTypesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyTypesRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyTypesRequestMethodGet) Where(v []string) *ByProjectKeyTypesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyTypesRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyTypesRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyTypesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyTypesRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyTypesRequestMethodGet) WithQueryParams(input ByProjectKeyTypesRequestMethodGetInput) *ByProjectKeyTypesRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyTypesRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyTypesRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyTypesRequestMethodGet) Execute(ctx context.Context) (result *TypePagedQueryResponse, err error) {
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
