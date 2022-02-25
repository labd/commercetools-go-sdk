package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyShippingMethodsMatchingCartRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyShippingMethodsMatchingCartRequestMethodGetInput
}

func (r *ByProjectKeyShippingMethodsMatchingCartRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyShippingMethodsMatchingCartRequestMethodGetInput struct {
	CartId string
	Expand []string
}

func (input *ByProjectKeyShippingMethodsMatchingCartRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	values.Add("cartId", fmt.Sprintf("%v", input.CartId))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyShippingMethodsMatchingCartRequestMethodGet) CartId(v string) *ByProjectKeyShippingMethodsMatchingCartRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingCartRequestMethodGetInput{}
	}
	rb.params.CartId = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingCartRequestMethodGet) Expand(v []string) *ByProjectKeyShippingMethodsMatchingCartRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingCartRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingCartRequestMethodGet) WithQueryParams(input ByProjectKeyShippingMethodsMatchingCartRequestMethodGetInput) *ByProjectKeyShippingMethodsMatchingCartRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyShippingMethodsMatchingCartRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyShippingMethodsMatchingCartRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyShippingMethodsMatchingCartRequestMethodGet) Execute(ctx context.Context) (result *ShippingMethodPagedQueryResponse, err error) {
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
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
		}
		return nil, result
	default:
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
