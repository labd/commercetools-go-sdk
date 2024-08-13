package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestMethodPost struct {
	body    MyCustomerDraft
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*
*	If used with an optional [access token for an anonymous session](ctp:api:type:AnonymousSession), all Orders and Carts that belong to the `anonymousId` are assigned to the newly created Customer.
*
*	If omitted in the request body, the [Customer](ctp:api:type:Customer) `stores` field is set to the [Store](ctp:api:type:Store) specified in the path parameter.
*
*	A Cart returned in the [CustomerSignInResult](ctp:api:type:CustomerSignInResult) has any invalid Line Items removed and is [updated](/api/carts-orders-overview#cart-updates) with the latest prices, taxes, and discounts. During these updates, the following errors can be returned: [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError) and [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError).
*
*	Creating a Customer produces the [CustomerCreated](ctp:api:type:CustomerCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestMethodPost) Execute(ctx context.Context) (result *CustomerSignInResult, err error) {
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
