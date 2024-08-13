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

type ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGetInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGetInput) *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Returns an Order for a given `id` in a Store. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Order exists in the Store for the given `id`.
*	- If the Order exists but does not have a `store` specified, or the `store` field references a different Store.
*	- If the Order exists but does not have a `customerId` that matches the [customer:{id}](/scopes#customer_idid) scope, or `anonymousId` that matches the [anonymous_id:{id}](/scopes#anonymous_idid) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet) Execute(ctx context.Context) (result *Order, err error) {
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
