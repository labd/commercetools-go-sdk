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

type ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost struct {
	body    OrderUpdate
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPostInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPostInput) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Updates an Order in a [Store](ctp:api:type:Store) using one or more [update actions](/../api/projects/orders#update-actions).
*	If the Order exists in the Project but does not have a `store` specified, or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodPost) Execute(ctx context.Context) (result *Order, err error) {
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
