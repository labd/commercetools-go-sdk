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

type ByProjectKeyProductsByIDProductSelectionsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductsByIDProductSelectionsRequestMethodGetInput
}

func (r *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductsByIDProductSelectionsRequestMethodGetInput struct {
	WithTotal    *bool
	Expand       []string
	Sort         []string
	Limit        *int
	Offset       *int
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyProductsByIDProductSelectionsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.WithTotal != nil {
		if *input.WithTotal {
			values.Add("withTotal", "true")
		} else {
			values.Add("withTotal", "false")
		}
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

func (rb *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet) WithTotal(v bool) *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDProductSelectionsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet) Expand(v []string) *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDProductSelectionsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet) Sort(v []string) *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDProductSelectionsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet) Limit(v int) *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDProductSelectionsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet) Offset(v int) *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDProductSelectionsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet) Where(v []string) *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDProductSelectionsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductsByIDProductSelectionsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet) WithQueryParams(input ByProjectKeyProductsByIDProductSelectionsRequestMethodGetInput) *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet) Execute(ctx context.Context) (result *AssignedProductSelectionPagedQueryResponse, err error) {
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
