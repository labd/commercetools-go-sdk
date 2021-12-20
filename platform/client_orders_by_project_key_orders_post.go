// Generated file, please do not change!!!
package platform

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
*	Creates an order from a Cart.
*	The cart must have a shipping address set before creating an order.
*	When using the Platform TaxMode, the shipping address is used for tax calculation.
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
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 201:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 409, 400, 401, 403, 500, 502, 503:
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
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
