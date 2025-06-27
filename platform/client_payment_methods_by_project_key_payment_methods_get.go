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

type ByProjectKeyPaymentMethodsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyPaymentMethodsRequestMethodGetInput
}

func (r *ByProjectKeyPaymentMethodsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyPaymentMethodsRequestMethodGetInput struct {
	Expand       []string
	Sort         []string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyPaymentMethodsRequestMethodGetInput) Values() url.Values {
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
			values.Add(k, x)
		}
	}
	return values
}

func (rb *ByProjectKeyPaymentMethodsRequestMethodGet) Expand(v []string) *ByProjectKeyPaymentMethodsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentMethodsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyPaymentMethodsRequestMethodGet) Sort(v []string) *ByProjectKeyPaymentMethodsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentMethodsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyPaymentMethodsRequestMethodGet) Limit(v int) *ByProjectKeyPaymentMethodsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentMethodsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyPaymentMethodsRequestMethodGet) Offset(v int) *ByProjectKeyPaymentMethodsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentMethodsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyPaymentMethodsRequestMethodGet) WithTotal(v bool) *ByProjectKeyPaymentMethodsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentMethodsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyPaymentMethodsRequestMethodGet) Where(v []string) *ByProjectKeyPaymentMethodsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentMethodsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyPaymentMethodsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyPaymentMethodsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentMethodsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyPaymentMethodsRequestMethodGet) WithQueryParams(input ByProjectKeyPaymentMethodsRequestMethodGetInput) *ByProjectKeyPaymentMethodsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyPaymentMethodsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyPaymentMethodsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves all PaymentMethods in the Project.
*
 */
func (rb *ByProjectKeyPaymentMethodsRequestMethodGet) Execute(ctx context.Context) (result *PaymentMethodPagedQueryResponse, err error) {
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
