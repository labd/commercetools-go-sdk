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

type ByProjectKeyProductDiscountsMatchingRequestMethodPost struct {
	body    ProductDiscountMatchQuery
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyProductDiscountsMatchingRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyProductDiscountsMatchingRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyProductDiscountsMatchingRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	This endpoint can be used to simulate which Product Discounts would be applied if a specified Product Variant had a specified Price.
*	Given Product and Product Variant IDs and a Price, this endpoint will return the [ProductDiscount](ctp:api:type:ProductDiscount) that would have been applied to that Price.
*
 */
func (rb *ByProjectKeyProductDiscountsMatchingRequestMethodPost) Execute(ctx context.Context) (result *ProductDiscount, err error) {
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
	case 404:
		errorObj := NoMatchingProductDiscountFoundError{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	default:
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
