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

type ByProjectKeyCustomersByIDRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCustomersByIDRequestMethodDeleteInput
}

func (r *ByProjectKeyCustomersByIDRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCustomersByIDRequestMethodDeleteInput struct {
	DataErasure *bool
	Version     int
	Expand      []string
}

func (input *ByProjectKeyCustomersByIDRequestMethodDeleteInput) Values() url.Values {
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

func (rb *ByProjectKeyCustomersByIDRequestMethodDelete) DataErasure(v bool) *ByProjectKeyCustomersByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomersByIDRequestMethodDeleteInput{}
	}
	rb.params.DataErasure = &v
	return rb
}

func (rb *ByProjectKeyCustomersByIDRequestMethodDelete) Version(v int) *ByProjectKeyCustomersByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomersByIDRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyCustomersByIDRequestMethodDelete) Expand(v []string) *ByProjectKeyCustomersByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomersByIDRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCustomersByIDRequestMethodDelete) WithQueryParams(input ByProjectKeyCustomersByIDRequestMethodDeleteInput) *ByProjectKeyCustomersByIDRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCustomersByIDRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyCustomersByIDRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyCustomersByIDRequestMethodDelete) Execute(ctx context.Context) (result *Customer, err error) {
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
