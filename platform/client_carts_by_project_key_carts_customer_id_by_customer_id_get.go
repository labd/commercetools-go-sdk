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

type ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGetInput
}

func (r *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet) Expand(v []string) *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet) WithQueryParams(input ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGetInput) *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves the active cart of the customer that has been modified most recently.
*	It does not consider carts with CartOrigin Merchant. If no active cart exists, a 404 Not Found error is returned.
*
*	The cart may not contain up-to-date prices, discounts etc. If you want to ensure they're up-to-date,
*	send an Update request with the Recalculate update action instead.
*
 */
func (rb *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet) Execute(ctx context.Context) (result *Cart, err error) {
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

	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
