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

type ByProjectKeyMeQuotesRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyMeQuotesRequestMethodGetInput
}

func (r *ByProjectKeyMeQuotesRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyMeQuotesRequestMethodGetInput struct {
	Expand       []string
	Sort         []string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyMeQuotesRequestMethodGetInput) Values() url.Values {
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

func (rb *ByProjectKeyMeQuotesRequestMethodGet) Expand(v []string) *ByProjectKeyMeQuotesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeQuotesRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyMeQuotesRequestMethodGet) Sort(v []string) *ByProjectKeyMeQuotesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeQuotesRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyMeQuotesRequestMethodGet) Limit(v int) *ByProjectKeyMeQuotesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeQuotesRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyMeQuotesRequestMethodGet) Offset(v int) *ByProjectKeyMeQuotesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeQuotesRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyMeQuotesRequestMethodGet) WithTotal(v bool) *ByProjectKeyMeQuotesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeQuotesRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyMeQuotesRequestMethodGet) Where(v []string) *ByProjectKeyMeQuotesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeQuotesRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyMeQuotesRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyMeQuotesRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeQuotesRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyMeQuotesRequestMethodGet) WithQueryParams(input ByProjectKeyMeQuotesRequestMethodGetInput) *ByProjectKeyMeQuotesRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyMeQuotesRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyMeQuotesRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyMeQuotesRequestMethodGet) Execute(ctx context.Context) (result *QuotePagedQueryResponse, err error) {
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

	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
