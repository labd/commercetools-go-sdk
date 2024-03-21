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

type ByProjectKeyMeShoppingListsByIDRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyMeShoppingListsByIDRequestMethodDeleteInput
}

func (r *ByProjectKeyMeShoppingListsByIDRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyMeShoppingListsByIDRequestMethodDeleteInput struct {
	Expand  []string
	Version int
}

func (input *ByProjectKeyMeShoppingListsByIDRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	values.Add("version", strconv.Itoa(input.Version))
	return values
}

func (rb *ByProjectKeyMeShoppingListsByIDRequestMethodDelete) Expand(v []string) *ByProjectKeyMeShoppingListsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeShoppingListsByIDRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyMeShoppingListsByIDRequestMethodDelete) Version(v int) *ByProjectKeyMeShoppingListsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeShoppingListsByIDRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyMeShoppingListsByIDRequestMethodDelete) WithQueryParams(input ByProjectKeyMeShoppingListsByIDRequestMethodDeleteInput) *ByProjectKeyMeShoppingListsByIDRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyMeShoppingListsByIDRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyMeShoppingListsByIDRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyMeShoppingListsByIDRequestMethodDelete) Execute(ctx context.Context) (result *ShoppingList, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.delete(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		nil,
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
