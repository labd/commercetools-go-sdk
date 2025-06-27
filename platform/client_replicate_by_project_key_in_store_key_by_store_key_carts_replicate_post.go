package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyInStoreKeyByStoreKeyCartsReplicateRequestMethodPost struct {
	body    ReplicaCartDraft
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyInStoreKeyByStoreKeyCartsReplicateRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsReplicateRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyCartsReplicateRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Creates a new Cart in a [Store](ctp:api:type:Store) by replicating an existing Cart or Order.
*
*	The following applies to the new Cart:
*
*	- It contains the same Customer information, Line Items and Custom Line Items, Custom Fields, Discount Codes, and other settings of the originating Cart or Order.
*	- If a Line Item becomes invalid, it is removed from the new Cart. A common reason for this is removed Products or Prices.
*	- Line items and Custom Line Items are reset to their initial [state](/projects/carts#itemstate).
*	- It contains no payments or delivery information.
*	- It contains up-to-date Tax Rates, Prices, and Line Item product data.
*	- The [CartState](/projects/carts#cartstate) is `Active`.
*	- If using the `customerGroup` field (for a single Customer Group) and the referenced Customer switched to another Customer Group, the new Cart is automatically updated to reflect the new group and corresponding prices.
*	- If using the `customerGroupAssignments` field (for multiple Customer Groups), the Cart no longer keeps a direct reference to a Customer Group. If a Customer’s group assignments change, the Cart and its Line Item prices are not updated automatically. Prices are only updated when the Cart is changed via a [direct update action](/projects/carts#update-actions).
*
*	Specific Error Codes:
*
*	- [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError)
*	- [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError)
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsReplicateRequestMethodPost) Execute(ctx context.Context) (result *Cart, err error) {
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
