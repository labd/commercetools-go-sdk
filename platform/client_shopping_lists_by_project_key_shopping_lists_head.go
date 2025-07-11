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

type ByProjectKeyShoppingListsRequestMethodHead struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyShoppingListsRequestMethodHeadInput
}

func (r *ByProjectKeyShoppingListsRequestMethodHead) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyShoppingListsRequestMethodHeadInput struct {
	Where []string
}

func (input *ByProjectKeyShoppingListsRequestMethodHeadInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyShoppingListsRequestMethodHead) Where(v []string) *ByProjectKeyShoppingListsRequestMethodHead {
	if rb.params == nil {
		rb.params = &ByProjectKeyShoppingListsRequestMethodHeadInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyShoppingListsRequestMethodHead) WithQueryParams(input ByProjectKeyShoppingListsRequestMethodHeadInput) *ByProjectKeyShoppingListsRequestMethodHead {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyShoppingListsRequestMethodHead) WithHeaders(headers http.Header) *ByProjectKeyShoppingListsRequestMethodHead {
	rb.headers = headers
	return rb
}

/**
*	Checks if one or more ShoppingLists exist for the provided query predicate. Returns a `200 OK` status if any ShoppingLists match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyShoppingListsRequestMethodHead) Execute(ctx context.Context) error {
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
