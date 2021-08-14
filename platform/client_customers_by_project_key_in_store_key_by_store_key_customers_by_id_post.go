// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestMethodPost struct {
	body   CustomerUpdate
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestMethodPostInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestMethodPostInput struct {
	Expand *[]string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
		}
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestMethodPost) WithQueryParams(input *ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestMethodPostInput) *ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestMethodPost {
	rb.query = input.Values()
	return rb
}

/**
*	Updates a customer in the store specified by {storeKey}. The {storeKey} path parameter maps to a Store's key.
*	If the customer exists in the commercetools project but the stores field references a different store,
*	this method returns a ResourceNotFound error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersByIdRequestMethodPost) Execute() (result *Customer, err error) {
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
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 409, 400, 401, 403, 500, 503:
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
