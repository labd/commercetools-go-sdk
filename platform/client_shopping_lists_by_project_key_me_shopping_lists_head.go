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

type ByProjectKeyMeShoppingListsRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyMeShoppingListsRequestMethodHeadInput
}

func (r *ByProjectKeyMeShoppingListsRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyMeShoppingListsRequestMethodHeadInput struct {
	Where []string
}

func (input *ByProjectKeyMeShoppingListsRequestMethodHeadInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyMeShoppingListsRequestMethodHead) Where(v []string) *ByProjectKeyMeShoppingListsRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeShoppingListsRequestMethodHeadInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyMeShoppingListsRequestMethodHead) WithQueryParams(input ByProjectKeyMeShoppingListsRequestMethodHeadInput) *ByProjectKeyMeShoppingListsRequestMethodHead {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyMeShoppingListsRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyMeShoppingListsRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if one or more ShoppingLists exist for the provided query predicate for the authenticated Customer or anonymous user. Returns a `200 OK` status if successful.
*
*	A [Not Found](/../api/errors#404-not-found) error is returned in the following scenarios:
*
*	- If no ShoppingList exists for the provided query predicate.
*	- If a ShoppingList exists but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeShoppingListsRequestMethodHead) Execute(ctx context.Context) error {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.head(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
	)

	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		return nil
	case 404:
		return ErrNotFound
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 500:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 502:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	case 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return err
		}
		return errorObj
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return result
	}

}
