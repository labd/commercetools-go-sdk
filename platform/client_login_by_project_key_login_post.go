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

type ByProjectKeyLoginRequestMethodPost struct {
	body    CustomerSignin
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyLoginRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyLoginRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyLoginRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Authenticate Customer (Sign In). Retrieves the authenticated
*	customer (a customer that matches the given email/password pair).
*	If used with an access token for Anonymous Sessions,
*	all orders and carts belonging to the anonymousId will be assigned to the newly created customer.
*	If a cart is is returned as part of the CustomerSignInResult,
*	it has been recalculated (It will have up-to-date prices, taxes and discounts,
*	and invalid line items have been removed.).
*
 */
func (rb *ByProjectKeyLoginRequestMethodPost) Execute(ctx context.Context) (result *CustomerSignInResult, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	queryParams := url.Values{}
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
