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

type ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGetInput
}

func (r *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet) Expand(v []string) *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet) WithQueryParams(input ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGetInput) *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves the recently modified active Cart of a Customer with [CartOrigin](ctp:api:type:CartOrigin) `Customer`. If no active Cart exists, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
*	To ensure the Cart is up-to-date with current values (such as Prices and Discounts), use the [Recalculate](ctp:api:type:CartRecalculateAction) update action.
*
 */
func (rb *ByProjectKeyCartsCustomerIdByCustomerIdRequestMethodGet) Execute(ctx context.Context) (result *Cart, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.get(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
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
