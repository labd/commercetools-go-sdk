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

type ByProjectKeyMeShoppingListsRequestMethodPost struct {
	body    MyShoppingListDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyMeShoppingListsRequestMethodPostInput
}

func (r *ByProjectKeyMeShoppingListsRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyMeShoppingListsRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyMeShoppingListsRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyMeShoppingListsRequestMethodPost) Expand(v []string) *ByProjectKeyMeShoppingListsRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeShoppingListsRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyMeShoppingListsRequestMethodPost) WithQueryParams(input ByProjectKeyMeShoppingListsRequestMethodPostInput) *ByProjectKeyMeShoppingListsRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyMeShoppingListsRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyMeShoppingListsRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Creates a ShoppingList for the Customer or anonymous user. The `customerId` or `anonymousId` on the ShoppingList is automatically set based on the given [customer:{id}](/scopes#customer_idid) or [anonymous_id:{id}](/scopes#anonymous_idid) scope.
*
 */
func (rb *ByProjectKeyMeShoppingListsRequestMethodPost) Execute(ctx context.Context) (result *ShoppingList, err error) {
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
	case 201:
		err = json.Unmarshal(content, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
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
