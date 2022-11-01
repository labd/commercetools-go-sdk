package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyMeLoginRequestMethodPost struct {
	body    MyCustomerSignin
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyMeLoginRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyMeLoginRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyMeLoginRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Retrieves the authenticated customer (that matches the given email/password pair).
*
*	If used with [an access token for an anonymous session](/../api/authorization#tokens-for-anonymous-sessions), all Orders and Carts that belong to the `anonymousId` are assigned to the newly logged-in Customer.
*
*	- If the Customer does not have a Cart yet, the most recently modified anonymous cart becomes the Customer's Cart.
*	- If the Customer already has a Cart, the most recently modified anonymous cart is handled in accordance with [AnonymousCartSignInMode](ctp:api:type:AnonymousCartSignInMode).
*
*	A Cart returned as part of the [CustomerSignInResult](ctp:api:type:CustomerSignInResult) is [recalculated](ctp:api:type:Recalculate) with up-to-date prices, taxes, discounts, and invalid line items removed.
*
*	If an account with the given credentials is not found, an [InvalidCredentials](ctp:api:type:InvalidCredentialsError) error is returned.
*
 */
func (rb *ByProjectKeyMeLoginRequestMethodPost) Execute(ctx context.Context) (result *CustomerSignInResult, err error) {
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

	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
