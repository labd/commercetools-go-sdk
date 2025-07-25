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

type ByProjectKeyOrdersRequestMethodPost struct {
	body    OrderFromCartDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyOrdersRequestMethodPostInput
}

func (r *ByProjectKeyOrdersRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyOrdersRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyOrdersRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyOrdersRequestMethodPost) Expand(v []string) *ByProjectKeyOrdersRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyOrdersRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyOrdersRequestMethodPost) WithQueryParams(input ByProjectKeyOrdersRequestMethodPostInput) *ByProjectKeyOrdersRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyOrdersRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyOrdersRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Creates an Order from a Cart.
*
*	The Cart must have a shipping address and an active Shipping Method set.
*
*	The shipping address is used for tax calculation for a Cart with `Platform` [TaxMode](ctp:api:type:TaxMode).
*
*	Creating an Order produces the [OrderCreated](ctp:api:type:OrderCreatedMessage) Message. If the Order is created from a Recurring Order schedule, the [OrderCreatedFromRecurringOrder](ctp:api:type:OrderCreatedFromRecurringOrderMessage) Message is generated.
*
*	If a server-side problem occurs, indicated by a 500 Internal Server Error HTTP response, the Order creation may still successfully complete after the error is returned.
*	If you receive this error, you should verify the status of the Order by querying a unique identifier supplied during the creation request, such as the Order number.
*
*	Specific Error Codes:
*
*	- [OutOfStock](ctp:api:type:OutOfStockError)
*	- [PriceChanged](ctp:api:type:PriceChangedError)
*	- [DiscountCodeNonApplicable](ctp:api:type:DiscountCodeNonApplicableError)
*	- [ShippingMethodDoesNotMatchCart](ctp:api:type:ShippingMethodDoesNotMatchCartError)
*	- [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError)
*	- [InvalidOperation](ctp:api:type:InvalidOperationError)
*	- [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError)
*	- [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError)
*
 */
func (rb *ByProjectKeyOrdersRequestMethodPost) Execute(ctx context.Context) (result *Order, err error) {
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
