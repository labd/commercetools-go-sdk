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

type ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGetInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGet) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGet) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGetInput) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Returns an order by its order number from a specific Store.
*
*	If the order exists in the commercetools project but does not have the store field,
*	or the store field references a different store, this method returns a ResourceNotFound error.
*	In case the orderNumber does not match the regular expression [a-zA-Z0-9_\-]+,
*	it should be provided in URL-encoded format.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestMethodGet) Execute(ctx context.Context) (result *Order, err error) {
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
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
