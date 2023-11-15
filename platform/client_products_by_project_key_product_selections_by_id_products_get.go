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

type ByProjectKeyProductSelectionsByIDProductsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductSelectionsByIDProductsRequestMethodGetInput
}

func (r *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductSelectionsByIDProductsRequestMethodGetInput struct {
	Expand    []string
	Limit     *int
	Offset    *int
	WithTotal *bool
	Sort      []string
}

func (input *ByProjectKeyProductSelectionsByIDProductsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
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
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet) Expand(v []string) *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsByIDProductsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet) Limit(v int) *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsByIDProductsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet) Offset(v int) *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsByIDProductsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet) WithTotal(v bool) *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsByIDProductsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet) Sort(v []string) *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsByIDProductsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet) WithQueryParams(input ByProjectKeyProductSelectionsByIDProductsRequestMethodGetInput) *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet) Execute(ctx context.Context) (result *ProductSelectionProductPagedQueryResponse, err error) {
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
