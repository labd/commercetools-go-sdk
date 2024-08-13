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

type ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHeadInput
}

func (r *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHeadInput struct {
	Country string
	State   *string
	CartId  string
}

func (input *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHeadInput) Values() url.Values {
	values := url.Values{}
	values.Add("country", fmt.Sprintf("%v", input.Country))
	if input.State != nil {
		values.Add("state", fmt.Sprintf("%v", *input.State))
	}
	values.Add("cartId", fmt.Sprintf("%v", input.CartId))
	return values
}

func (rb *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHead) Country(v string) *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHeadInput{}
	}
	rb.params.Country = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHead) State(v string) *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHeadInput{}
	}
	rb.params.State = &v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHead) CartId(v string) *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHeadInput{}
	}
	rb.params.CartId = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHead) WithQueryParams(input ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHeadInput) *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHead {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if an active ShippingMethod that can ship to the given [Location](ctp:api:type:Location) exists for the given Cart. Returns a `200 OK` status if the ShippingMethod exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyShippingMethodsMatchingCartLocationRequestMethodHead) Execute(ctx context.Context) error {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.head(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
	)

	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 404:
		return ErrNotFound
	case 500:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 502:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 200:
		return nil
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return result
	}

}
