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

type ByProjectKeyApiClientsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyApiClientsRequestMethodGetInput
}

func (r *ByProjectKeyApiClientsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyApiClientsRequestMethodGetInput struct {
	Expand       []string
	Sort         []string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyApiClientsRequestMethodGetInput) Values() url.Values {
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

func (rb *ByProjectKeyApiClientsRequestMethodGet) Expand(v []string) *ByProjectKeyApiClientsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyApiClientsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyApiClientsRequestMethodGet) Sort(v []string) *ByProjectKeyApiClientsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyApiClientsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyApiClientsRequestMethodGet) Limit(v int) *ByProjectKeyApiClientsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyApiClientsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyApiClientsRequestMethodGet) Offset(v int) *ByProjectKeyApiClientsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyApiClientsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyApiClientsRequestMethodGet) WithTotal(v bool) *ByProjectKeyApiClientsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyApiClientsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyApiClientsRequestMethodGet) Where(v []string) *ByProjectKeyApiClientsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyApiClientsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyApiClientsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyApiClientsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyApiClientsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyApiClientsRequestMethodGet) WithQueryParams(input ByProjectKeyApiClientsRequestMethodGetInput) *ByProjectKeyApiClientsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyApiClientsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyApiClientsRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyApiClientsRequestMethodGet) Execute(ctx context.Context) (result *ApiClientPagedQueryResponse, err error) {
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
