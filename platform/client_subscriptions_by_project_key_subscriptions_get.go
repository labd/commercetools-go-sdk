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

type ByProjectKeySubscriptionsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeySubscriptionsRequestMethodGetInput
}

func (r *ByProjectKeySubscriptionsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeySubscriptionsRequestMethodGetInput struct {
	Sort         []string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeySubscriptionsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
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

func (rb *ByProjectKeySubscriptionsRequestMethodGet) Sort(v []string) *ByProjectKeySubscriptionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeySubscriptionsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeySubscriptionsRequestMethodGet) Limit(v int) *ByProjectKeySubscriptionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeySubscriptionsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeySubscriptionsRequestMethodGet) Offset(v int) *ByProjectKeySubscriptionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeySubscriptionsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeySubscriptionsRequestMethodGet) WithTotal(v bool) *ByProjectKeySubscriptionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeySubscriptionsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeySubscriptionsRequestMethodGet) Where(v []string) *ByProjectKeySubscriptionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeySubscriptionsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeySubscriptionsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeySubscriptionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeySubscriptionsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeySubscriptionsRequestMethodGet) WithQueryParams(input ByProjectKeySubscriptionsRequestMethodGetInput) *ByProjectKeySubscriptionsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeySubscriptionsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeySubscriptionsRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeySubscriptionsRequestMethodGet) Execute(ctx context.Context) (result *SubscriptionPagedQueryResponse, err error) {
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
