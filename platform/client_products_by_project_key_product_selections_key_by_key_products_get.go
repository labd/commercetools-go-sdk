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

type ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput
}

func (r *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput struct {
	Where     []string
	WithTotal *bool
	Expand    []string
	Limit     *int
	Offset    *int
	Sort      []string
}

func (input *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
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
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) Where(v []string) *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) WithTotal(v bool) *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) Expand(v []string) *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) Limit(v int) *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) Offset(v int) *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) Sort(v []string) *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) WithQueryParams(input ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGetInput) *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet) Execute(ctx context.Context) (result *ProductSelectionProductPagedQueryResponse, err error) {
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
