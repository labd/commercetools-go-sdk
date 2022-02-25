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

type ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGetInput
}

func (r *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGetInput struct {
	OrderEditId string
	Country     string
	State       *string
}

func (input *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	values.Add("orderEditId", fmt.Sprintf("%v", input.OrderEditId))
	values.Add("country", fmt.Sprintf("%v", input.Country))
	if input.State != nil {
		values.Add("state", fmt.Sprintf("%v", *input.State))
	}
	return values
}

func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet) OrderEditId(v string) *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGetInput{}
	}
	rb.params.OrderEditId = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet) Country(v string) *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGetInput{}
	}
	rb.params.Country = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet) State(v string) *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGetInput{}
	}
	rb.params.State = &v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet) WithQueryParams(input ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGetInput) *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodGet) Execute(ctx context.Context) (result *ShippingMethodPagedQueryResponse, err error) {
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
