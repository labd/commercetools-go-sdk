package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyMeOrdersQuotesRequestMethodPost struct {
	body    MyOrderFromQuoteDraft
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyMeOrdersQuotesRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyMeOrdersQuotesRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyMeOrdersQuotesRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	When creating [B2B Orders](/associates-overview#b2b-resources), the Customer must have the `CreateMyOrdersFromMyQuotes` [Permission](ctp:api:type:Permission).
*
*	Creating an Order produces the [OrderCreated](ctp:api:type:OrderCreatedMessage) Message.
*
*	Specific Error Codes:
*
*	- [OutOfStock](ctp:api:type:OutOfStockError)
*	- [PriceChanged](ctp:api:type:PriceChangedError)
*	- [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError)
*	- [CountryNotConfiguredInStore](ctp:api:type:CountryNotConfiguredInStoreError)
*	- [AssociateMissingPermission](ctp:api:type:AssociateMissingPermissionError)
*
 */
func (rb *ByProjectKeyMeOrdersQuotesRequestMethodPost) Execute(ctx context.Context) (result *Order, err error) {
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
