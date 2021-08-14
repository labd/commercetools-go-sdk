// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyOrdersRequestMethodPost struct {
	body   OrderFromCartDraft
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyOrdersRequestMethodPostInput
}

func (r *ByProjectKeyOrdersRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyOrdersRequestMethodPostInput struct {
	Expand *[]string
}

func (input *ByProjectKeyOrdersRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
		}
	}
	return values
}

func (rb *ByProjectKeyOrdersRequestMethodPost) WithQueryParams(input *ByProjectKeyOrdersRequestMethodPostInput) *ByProjectKeyOrdersRequestMethodPost {
	rb.query = input.Values()
	return rb
}

/**
*	Creates an order from a Cart.
*	The cart must have a shipping address set before creating an order.
*	When using the Platform TaxMode, the shipping address is used for tax calculation.
*
 */
func (rb *ByProjectKeyOrdersRequestMethodPost) Execute() (result *Order, err error) {
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
