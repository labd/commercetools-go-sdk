// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyCustomersRequestMethodPost struct {
	body   CustomerDraft
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyCustomersRequestMethodPostInput
}

func (r *ByProjectKeyCustomersRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyCustomersRequestMethodPostInput struct {
	Expand *[]string
}

func (input *ByProjectKeyCustomersRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
		}
	}
	return values
}

func (rb *ByProjectKeyCustomersRequestMethodPost) WithQueryParams(input *ByProjectKeyCustomersRequestMethodPostInput) *ByProjectKeyCustomersRequestMethodPost {
	rb.query = input.Values()
	return rb
}

/**
*	Creates a customer. If an anonymous cart is passed in,
*	then the cart is assigned to the created customer and the version number of the Cart will increase.
*	If the ID of an anonymous session is given, all carts and orders will be assigned to the created customer.
*
 */
func (rb *ByProjectKeyCustomersRequestMethodPost) Execute() (result *CustomerSignInResult, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	resp, err := rb.client.post(
		context.Background(),
		rb.url,
		rb.query,
		data,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 201:
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
