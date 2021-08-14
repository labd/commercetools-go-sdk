// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
)

type ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestMethodDelete struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestMethodDeleteInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestMethodDeleteInput struct {
	DataErasure *bool
	Version     int
	Expand      *[]string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	if input.DataErasure != nil {
		if *input.DataErasure == true {
			values.Add("dataErasure", "true")
		} else {
			values.Add("dataErasure", "false")
		}
	}
	values.Add("version", strconv.Itoa(input.Version))
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
		}
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestMethodDelete) WithQueryParams(input *ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestMethodDeleteInput) *ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestMethodDelete {
	rb.query = input.Values()
	return rb
}

/**
*	Delete ShoppingList by key
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestMethodDelete) Execute() (result *ShoppingList, err error) {
	resp, err := rb.client.delete(
		context.Background(),
		rb.url,
		rb.query,
		nil,
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
