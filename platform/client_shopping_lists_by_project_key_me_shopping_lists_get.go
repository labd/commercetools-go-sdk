package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyMeShoppingListsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyMeShoppingListsRequestMethodGetInput
}

func (r *ByProjectKeyMeShoppingListsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyMeShoppingListsRequestMethodGetInput struct {
	Expand       []string
	Sort         []string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Where        []string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyMeShoppingListsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	if input.WithTotal != nil {
		if *input.WithTotal {
			values.Add("withTotal", "true")
		} else {
			values.Add("withTotal", "false")
		}
	}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	for k, v := range input.PredicateVar {
		for _, x := range v {
			values.Add(k, x)
		}
	}
	return values
}

func (rb *ByProjectKeyMeShoppingListsRequestMethodGet) Expand(v []string) *ByProjectKeyMeShoppingListsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeShoppingListsRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyMeShoppingListsRequestMethodGet) Sort(v []string) *ByProjectKeyMeShoppingListsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeShoppingListsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyMeShoppingListsRequestMethodGet) Limit(v int) *ByProjectKeyMeShoppingListsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeShoppingListsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyMeShoppingListsRequestMethodGet) Offset(v int) *ByProjectKeyMeShoppingListsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeShoppingListsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyMeShoppingListsRequestMethodGet) WithTotal(v bool) *ByProjectKeyMeShoppingListsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeShoppingListsRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyMeShoppingListsRequestMethodGet) Where(v []string) *ByProjectKeyMeShoppingListsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeShoppingListsRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyMeShoppingListsRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyMeShoppingListsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeShoppingListsRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyMeShoppingListsRequestMethodGet) WithQueryParams(input ByProjectKeyMeShoppingListsRequestMethodGetInput) *ByProjectKeyMeShoppingListsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyMeShoppingListsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyMeShoppingListsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves ShoppingLists for the authenticated Customer or anonymous user. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingList exists for the provided query predicate.
*	- If a ShoppingList exists but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeShoppingListsRequestMethodGet) Execute(ctx context.Context) (result *ShoppingListPagedQueryResponse, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.get(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
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
