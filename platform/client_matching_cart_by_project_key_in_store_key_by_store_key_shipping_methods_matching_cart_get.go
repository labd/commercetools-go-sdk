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

type ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGetInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGetInput struct {
	CartId string
	Expand []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	values.Add("cartId", fmt.Sprintf("%v", input.CartId))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet) CartId(v string) *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGetInput{}
	}
	rb.params.CartId = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGetInput) *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyShippingMethodsMatchingCartRequestMethodGet) Execute(ctx context.Context) (result *ShippingMethodPagedQueryResponse, err error) {
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
