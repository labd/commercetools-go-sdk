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

type ByProjectKeyMeCartsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyMeCartsRequestMethodGetInput
}

func (r *ByProjectKeyMeCartsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyMeCartsRequestMethodGetInput struct {
	Expand       []string
	Sort         []string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyMeCartsRequestMethodGetInput) Values() url.Values {
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

func (rb *ByProjectKeyMeCartsRequestMethodGet) Expand(v []string) *ByProjectKeyMeCartsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeCartsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyMeCartsRequestMethodGet) Sort(v []string) *ByProjectKeyMeCartsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeCartsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyMeCartsRequestMethodGet) Limit(v int) *ByProjectKeyMeCartsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeCartsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyMeCartsRequestMethodGet) Offset(v int) *ByProjectKeyMeCartsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeCartsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyMeCartsRequestMethodGet) WithTotal(v bool) *ByProjectKeyMeCartsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeCartsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyMeCartsRequestMethodGet) Where(v []string) *ByProjectKeyMeCartsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeCartsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyMeCartsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyMeCartsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeCartsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyMeCartsRequestMethodGet) WithQueryParams(input ByProjectKeyMeCartsRequestMethodGetInput) *ByProjectKeyMeCartsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyMeCartsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyMeCartsRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyMeCartsRequestMethodGet) Execute(ctx context.Context) (result *CartPagedQueryResponse, err error) {
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
