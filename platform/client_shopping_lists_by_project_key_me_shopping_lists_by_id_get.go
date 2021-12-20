// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyMeShoppingListsByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyMeShoppingListsByIDRequestMethodGetInput
}

func (r *ByProjectKeyMeShoppingListsByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyMeShoppingListsByIDRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyMeShoppingListsByIDRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyMeShoppingListsByIDRequestMethodGet) Expand(v []string) *ByProjectKeyMeShoppingListsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeShoppingListsByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyMeShoppingListsByIDRequestMethodGet) WithQueryParams(input ByProjectKeyMeShoppingListsByIDRequestMethodGetInput) *ByProjectKeyMeShoppingListsByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyMeShoppingListsByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyMeShoppingListsByIDRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyMeShoppingListsByIDRequestMethodGet) Execute(ctx context.Context) (result *ShoppingList, err error) {
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
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
		}
		return nil, result
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
