package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
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
*	Authenticates a global Customer not associated with a Store.
*	For more information, see [Global versus Store-specific Customers](/../api/customers-overview#global-versus-store-specific-customers).
*	If the Customer is registered in a Store, use the [Authenticate (sign in) Customer in Store](ctp:api:endpoint:/{projectKey}/in-store/key={storeKey}/login:POST) method.
*
*	Triggers [Cart merge during sign-in](/../api/customers-overview#cart-merge-during-sign-in).
*
*	A Cart returned in the [CustomerSignInResult](ctp:api:type:CustomerSignInResult) has any invalid Line Items removed and is [updated](/api/carts-orders-overview#update-a-cart) with the latest prices, taxes, and discounts. During these updates, the following errors can be returned: [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError) and [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError).
*
*	If an account with the given credentials is not found, an [InvalidCredentials](ctp:api:type:InvalidCredentialsError) error is returned.
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
