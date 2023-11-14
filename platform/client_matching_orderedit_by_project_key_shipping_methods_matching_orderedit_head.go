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

type ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHeadInput
}

func (r *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHeadInput struct {
	OrderEditId string
	Country     string
	State       *string
}

func (input *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHeadInput) Values() url.Values {
	values := url.Values{}
	values.Add("orderEditId", fmt.Sprintf("%v", input.OrderEditId))
	values.Add("country", fmt.Sprintf("%v", input.Country))
	if input.State != nil {
		values.Add("state", fmt.Sprintf("%v", *input.State))
	}
	return values
}

func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHead) OrderEditId(v string) *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHeadInput{}
	}
	rb.params.OrderEditId = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHead) Country(v string) *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHeadInput{}
	}
	rb.params.Country = v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHead) State(v string) *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHeadInput{}
	}
	rb.params.State = &v
	return rb
}

func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHead) WithQueryParams(input ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHeadInput) *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHead {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if a ShippingMethod that can ship to the given [Location](ctp:api:type:Location) exists for the given [OrderEdit](ctp:api:type:OrderEdit). Returns a `200 OK` status if the ShippingMethod exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyShippingMethodsMatchingOrdereditRequestMethodHead) Execute(ctx context.Context) error {
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
