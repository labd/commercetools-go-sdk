package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPost struct {
	body    MyShoppingListUpdate
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPostInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPost) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPost) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPostInput) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Updates a ShoppingList for a given `id` in a Store. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingList matches the given `id` in a Store.
*	- If a ShoppingList matches the given `id` but does not have a `store` specified, or the `store` field references a different Store.
*	- If a ShoppingList matches the given `id` in a Store but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#anonymous_idid) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#customer_idid) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPost) Execute(ctx context.Context) (result *ShoppingList, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.post(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		data,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, ErrNotFound
	case 500:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 502:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
