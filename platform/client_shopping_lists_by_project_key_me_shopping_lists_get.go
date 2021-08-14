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

type ByProjectKeyMeShoppingListsRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyMeShoppingListsRequestMethodGetInput
}

func (r *ByProjectKeyMeShoppingListsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyMeShoppingListsRequestMethodGetInput struct {
	Expand       *[]string
	Sort         *[]string
	Limit        *int
	Offset       *int
	WithTotal    *bool
	Where        *[]string
	PredicateVar map[string][]string
}

func (input *ByProjectKeyMeShoppingListsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Expand != nil {
		for _, v := range *input.Expand {
			values.Add("expand", v)
		}
	}
	if input.Sort != nil {
		for _, v := range *input.Sort {
			values.Add("sort", v)
		}
	}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	if input.WithTotal != nil {
		if *input.WithTotal == true {
			values.Add("withTotal", "true")
		} else {
			values.Add("withTotal", "false")
		}
	}
	if input.Where != nil {
		for _, v := range *input.Where {
			values.Add("where", v)
		}
	}

	return values
}

func (rb *ByProjectKeyMeShoppingListsRequestMethodGet) WithQueryParams(input *ByProjectKeyMeShoppingListsRequestMethodGetInput) *ByProjectKeyMeShoppingListsRequestMethodGet {
	rb.query = input.Values()
	return rb
}

/**
*	Query shopping-lists
 */
func (rb *ByProjectKeyMeShoppingListsRequestMethodGet) Execute() (result *ShoppingListPagedQueryResponse, err error) {
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
