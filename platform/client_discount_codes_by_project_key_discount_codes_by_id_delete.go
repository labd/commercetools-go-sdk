// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyDiscountCodesByIDRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyDiscountCodesByIDRequestMethodDeleteInput
}

func (r *ByProjectKeyDiscountCodesByIDRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyDiscountCodesByIDRequestMethodDeleteInput struct {
	DataErasure *bool
	Version     int
	Expand      []string
}

func (input *ByProjectKeyDiscountCodesByIDRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	if input.DataErasure != nil {
		if *input.DataErasure {
			values.Add("dataErasure", "true")
		} else {
			values.Add("dataErasure", "false")
		}
	}
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyDiscountCodesByIDRequestMethodDelete) DataErasure(v bool) *ByProjectKeyDiscountCodesByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyDiscountCodesByIDRequestMethodDeleteInput{}
	}
	rb.params.DataErasure = &v
	return rb
}

func (rb *ByProjectKeyDiscountCodesByIDRequestMethodDelete) Version(v int) *ByProjectKeyDiscountCodesByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyDiscountCodesByIDRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyDiscountCodesByIDRequestMethodDelete) Expand(v []string) *ByProjectKeyDiscountCodesByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyDiscountCodesByIDRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyDiscountCodesByIDRequestMethodDelete) WithQueryParams(input ByProjectKeyDiscountCodesByIDRequestMethodDeleteInput) *ByProjectKeyDiscountCodesByIDRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyDiscountCodesByIDRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyDiscountCodesByIDRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyDiscountCodesByIDRequestMethodDelete) Execute(ctx context.Context) (result *DiscountCode, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.delete(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		nil,
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
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
