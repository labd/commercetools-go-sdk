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

type ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPost struct {
	body    CustomerDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPostInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPost) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPost) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPostInput) *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	When using this endpoint, if omitted, the Customer `stores` field is set to the Store specified in the path parameter.
*
*	If the `anonymousCart` field is set on the [CustomerDraft](ctp:api:type:CustomerDraft), then the newly created Customer will be assigned to that [Cart](ctp:api:type:Cart).
*	Similarly, if the `anonymousId` field is set, the Customer will be set on all [Carts](ctp:api:type:Cart), [Orders](ctp:api:type:Order), [ShoppingLists](ctp:api:type:ShoppingList) and [Payments](ctp:api:type:Payment) with the same `anonymousId`.
*	If a Cart with a `store` field specified, the `store` field must reference the same Store specified in the `{storeKey}` path parameter.
*	Creating a Customer produces the [CustomerCreated](ctp:api:type:CustomerCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCustomersRequestMethodPost) Execute(ctx context.Context) (result *CustomerSignInResult, err error) {
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
		return result, nil
	case 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, ErrNotFound
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
