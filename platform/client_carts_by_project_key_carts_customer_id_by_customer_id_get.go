// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGetInput
}

func (r *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGetInput struct {
	Expand *[]string
}

func (input *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
		}
	}
	return values
}

func (rb *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet) WithQueryParams(input *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGetInput) *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet {
	rb.query = input.Values()
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
func (rb *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet) Execute() (result *Cart, err error) {
	resp, err := rb.client.get(
		context.Background(),
		rb.url,
		rb.query,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 400, 401, 403, 500, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, fmt.Errorf("StatusCode %d returend", resp.StatusCode)
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
