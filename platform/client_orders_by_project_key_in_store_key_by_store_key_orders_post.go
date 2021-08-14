// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodPost struct {
	body   OrderFromCartDraft
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodPostInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodPostInput struct {
	Expand *[]string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
		}
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodPost) WithQueryParams(input *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodPostInput) *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodPost {
	rb.query = input.Values()
	return rb
}

/**
*	Creates an order from a Cart from a specific Store. The {storeKey} path parameter maps to a Store's key.
*	When using this endpoint the orders's store field is always set to the store specified in the path parameter.
*	The cart must have a shipping address set before creating an order. When using the Platform TaxMode,
*	the shipping address is used for tax calculation.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodPost) Execute() (result *Order, err error) {
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
