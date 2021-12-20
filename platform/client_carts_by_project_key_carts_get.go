// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyCartsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCartsRequestMethodGetInput
}

func (r *ByProjectKeyCartsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCartsRequestMethodGetInput struct {
	CustomerId   *string
	Expand       []string
	Sort         []string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyCartsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.CustomerId != nil {
		values.Add("customerId", fmt.Sprintf("%v", *input.CustomerId))
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
		if *input.WithTotal == true {
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

func (rb *ByProjectKeyCartsRequestMethodGet) CustomerId(v string) *ByProjectKeyCartsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCartsRequestMethodGetInput{}
	}
	rb.params.CustomerId = &v
	return rb
}

func (rb *ByProjectKeyCartsRequestMethodGet) Expand(v []string) *ByProjectKeyCartsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCartsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCartsRequestMethodGet) Sort(v []string) *ByProjectKeyCartsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCartsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyCartsRequestMethodGet) Limit(v int) *ByProjectKeyCartsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCartsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyCartsRequestMethodGet) Offset(v int) *ByProjectKeyCartsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCartsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyCartsRequestMethodGet) WithTotal(v bool) *ByProjectKeyCartsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCartsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyCartsRequestMethodGet) Where(v []string) *ByProjectKeyCartsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCartsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyCartsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyCartsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCartsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyCartsRequestMethodGet) WithQueryParams(input ByProjectKeyCartsRequestMethodGetInput) *ByProjectKeyCartsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCartsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyCartsRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyCartsRequestMethodGet) Execute(ctx context.Context) (result *CartPagedQueryResponse, err error) {
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
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
