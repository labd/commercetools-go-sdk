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

type ByProjectKeyQuotesRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyQuotesRequestMethodGetInput
}

func (r *ByProjectKeyQuotesRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyQuotesRequestMethodGetInput struct {
	Expand       []string
	Sort         []string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyQuotesRequestMethodGetInput) Values() url.Values {
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

func (rb *ByProjectKeyQuotesRequestMethodGet) Expand(v []string) *ByProjectKeyQuotesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyQuotesRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyQuotesRequestMethodGet) Sort(v []string) *ByProjectKeyQuotesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyQuotesRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyQuotesRequestMethodGet) Limit(v int) *ByProjectKeyQuotesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyQuotesRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyQuotesRequestMethodGet) Offset(v int) *ByProjectKeyQuotesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyQuotesRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyQuotesRequestMethodGet) WithTotal(v bool) *ByProjectKeyQuotesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyQuotesRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyQuotesRequestMethodGet) Where(v []string) *ByProjectKeyQuotesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyQuotesRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyQuotesRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyQuotesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyQuotesRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyQuotesRequestMethodGet) WithQueryParams(input ByProjectKeyQuotesRequestMethodGetInput) *ByProjectKeyQuotesRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyQuotesRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyQuotesRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyQuotesRequestMethodGet) Execute(ctx context.Context) (result *QuotePagedQueryResponse, err error) {
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
