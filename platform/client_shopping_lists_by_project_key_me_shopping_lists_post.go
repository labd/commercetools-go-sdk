// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyMeShoppingListsRequestMethodPost struct {
	body   MyShoppingListDraft
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyMeShoppingListsRequestMethodPostInput
}

func (r *ByProjectKeyMeShoppingListsRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyMeShoppingListsRequestMethodPostInput struct {
	Expand *[]string
}

func (input *ByProjectKeyMeShoppingListsRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
		}
	}
	return values
}

func (rb *ByProjectKeyMeShoppingListsRequestMethodPost) WithQueryParams(input *ByProjectKeyMeShoppingListsRequestMethodPostInput) *ByProjectKeyMeShoppingListsRequestMethodPost {
	rb.query = input.Values()
	return rb
}

/**
*	Create ShoppingList
 */
func (rb *ByProjectKeyMeShoppingListsRequestMethodPost) Execute() (result *ShoppingList, err error) {
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
