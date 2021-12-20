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

type ByProjectKeyDiscountCodesByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyDiscountCodesByIDRequestMethodGetInput
}

func (r *ByProjectKeyDiscountCodesByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyDiscountCodesByIDRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyDiscountCodesByIDRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyDiscountCodesByIDRequestMethodGet) Expand(v []string) *ByProjectKeyDiscountCodesByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyDiscountCodesByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyDiscountCodesByIDRequestMethodGet) WithQueryParams(input ByProjectKeyDiscountCodesByIDRequestMethodGetInput) *ByProjectKeyDiscountCodesByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyDiscountCodesByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyDiscountCodesByIDRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyDiscountCodesByIDRequestMethodGet) Execute(ctx context.Context) (result *DiscountCode, err error) {
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
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
