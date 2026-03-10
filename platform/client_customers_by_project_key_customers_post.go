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

type ByProjectKeyCustomersRequestMethodPost struct {
	body    CustomerDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCustomersRequestMethodPostInput
}

func (r *ByProjectKeyCustomersRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCustomersRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyCustomersRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCustomersRequestMethodPost) Expand(v []string) *ByProjectKeyCustomersRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomersRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCustomersRequestMethodPost) WithQueryParams(input ByProjectKeyCustomersRequestMethodPostInput) *ByProjectKeyCustomersRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCustomersRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyCustomersRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Allows converting an anonymous Cart to the active Cart of a Customer with [cart merge](/../api/customers-overview#cart-merge-during-sign-in-and-sign-up).
*
*	Creating a Customer produces the [CustomerCreated](ctp:api:type:CustomerCreatedMessage) Message. Simultaneously creating two Customers with the same email address can return a [LockedField](ctp:api:type:LockedFieldError) error.
*
 */
func (rb *ByProjectKeyCustomersRequestMethodPost) Execute(ctx context.Context) (result *CustomerSignInResult, err error) {
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
