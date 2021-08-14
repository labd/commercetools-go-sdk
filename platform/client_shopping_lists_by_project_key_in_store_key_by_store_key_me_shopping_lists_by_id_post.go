// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestMethodPost struct {
	body   MyShoppingListUpdate
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestMethodPostInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestMethodPostInput struct {
	Expand *[]string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
		}
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestMethodPost) WithQueryParams(input *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestMethodPostInput) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestMethodPost {
	rb.query = input.Values()
	return rb
}

/**
*	Update ShoppingList by ID
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestMethodPost) Execute() (result *ShoppingList, err error) {
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
