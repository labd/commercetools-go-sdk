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

type ByProjectKeyShippingMethodsMatchingCartRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyShippingMethodsMatchingCartRequestMethodHeadInput
}

func (r *ByProjectKeyShippingMethodsMatchingCartRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyShippingMethodsMatchingCartRequestMethodHeadInput struct {
	CartId string
}

func (input *ByProjectKeyShippingMethodsMatchingCartRequestMethodHeadInput) Values() url.Values {
	values := url.Values{}
	values.Add("cartId", fmt.Sprintf("%v", input.CartId))
	return values
}

func (rb *ByProjectKeyShippingMethodsMatchingCartRequestMethodHead) CartId(v string) *ByProjectKeyShippingMethodsMatchingCartRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingCartRequestMethodHeadInput{}
	}
	rb.params.CartId = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingCartRequestMethodHead) WithQueryParams(input ByProjectKeyShippingMethodsMatchingCartRequestMethodHeadInput) *ByProjectKeyShippingMethodsMatchingCartRequestMethodHead {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyShippingMethodsMatchingCartRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyShippingMethodsMatchingCartRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if a ShippingMethod exists for the given Cart. Returns a `200 OK` status if the ShippingMethod exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyShippingMethodsMatchingCartRequestMethodHead) Execute(ctx context.Context) error {
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
