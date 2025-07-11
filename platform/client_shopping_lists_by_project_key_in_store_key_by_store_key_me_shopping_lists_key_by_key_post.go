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

type ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPost struct {
	body    MyShoppingListUpdate
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPostInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPost) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPost) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPostInput) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Updates a ShoppingList for the authenticated Customer or anonymous user in a [Store](ctp:api:type:Store) using one or more [update actions](/../api/projects/me-shoppingLists#update-actions). Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingList matches the provided `key` in a Store.
*	- If a ShoppingList matches the provided `key` but does not have a `store` specified, or the `store` field references a different Store.
*	- If a ShoppingList matches the provided `key` in a Store but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPost) Execute(ctx context.Context) (result *ShoppingList, err error) {
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
