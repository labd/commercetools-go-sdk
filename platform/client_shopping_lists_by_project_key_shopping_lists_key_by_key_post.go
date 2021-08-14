// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyShoppingListsKeyByKeyRequestMethodPost struct {
	body   ShoppingListUpdate
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyShoppingListsKeyByKeyRequestMethodPostInput
}

func (r *ByProjectKeyShoppingListsKeyByKeyRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyShoppingListsKeyByKeyRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyShoppingListsKeyByKeyRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyShoppingListsKeyByKeyRequestMethodPost) WithQueryParams(input *ByProjectKeyShoppingListsKeyByKeyRequestMethodPostInput) *ByProjectKeyShoppingListsKeyByKeyRequestMethodPost {
	rb.query = input.Values()
	return rb
}

/**
*	Update a shopping list found by its Key.
 */
func (rb *ByProjectKeyShoppingListsKeyByKeyRequestMethodPost) Execute(ctx context.Context) (result *ShoppingList, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	resp, err := rb.client.post(
		ctx,
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
