// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyInStoreKeyByStoreKeyOrdersByIdRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyInStoreKeyByStoreKeyOrdersByIdRequestMethodGetInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyOrdersByIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyOrdersByIdRequestMethodGetInput struct {
	Expand *[]string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyOrdersByIdRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
		}
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersByIdRequestMethodGet) WithQueryParams(input *ByProjectKeyInStoreKeyByStoreKeyOrdersByIdRequestMethodGetInput) *ByProjectKeyInStoreKeyByStoreKeyOrdersByIdRequestMethodGet {
	rb.query = input.Values()
	return rb
}

/**
*	Returns an order by its ID from a specific Store. The {storeKey} path parameter maps to a Store's key.
*	If the order exists in the commercetools project but does not have the store field,
*	or the store field references a different store, this method returns a ResourceNotFound error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersByIdRequestMethodGet) Execute() (result *Order, err error) {
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
