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

type ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGetInput
}

func (r *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGetInput struct {
	Country string
	State   *string
	CartId  string
	Expand  []string
}

func (input *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	values.Add("country", fmt.Sprintf("%v", input.Country))
	if input.State != nil {
		values.Add("state", fmt.Sprintf("%v", *input.State))
	}
	values.Add("cartId", fmt.Sprintf("%v", input.CartId))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet) Country(v string) *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGetInput{}
	}
	rb.params.Country = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet) State(v string) *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGetInput{}
	}
	rb.params.State = &v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet) CartId(v string) *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGetInput{}
	}
	rb.params.CartId = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet) Expand(v []string) *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet) WithQueryParams(input ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGetInput) *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves all the ShippingMethods that can ship to the given [Location](/projects/zones#location)
*	with a `predicate` that matches the given Cart.
*	Each ShippingMethod contains exactly one ShippingRate with the flag `isMatching` set to `true`.
*	This ShippingRate is used when the ShippingMethod is [added to the Cart](ctp:api:type:CartSetShippingMethodAction).
*
 */
func (rb *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodGet) Execute(ctx context.Context) (result *ShippingMethodPagedQueryResponse, err error) {
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
		return nil, ErrNotFound
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
