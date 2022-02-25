package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyPaymentsByIDRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyPaymentsByIDRequestMethodDeleteInput
}

func (r *ByProjectKeyPaymentsByIDRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyPaymentsByIDRequestMethodDeleteInput struct {
	DataErasure *bool
	Version     int
	Expand      []string
}

func (input *ByProjectKeyPaymentsByIDRequestMethodDeleteInput) Values() url.Values {
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

func (rb *ByProjectKeyPaymentsByIDRequestMethodDelete) DataErasure(v bool) *ByProjectKeyPaymentsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentsByIDRequestMethodDeleteInput{}
	}
	rb.params.DataErasure = &v
	return rb
}

func (rb *ByProjectKeyPaymentsByIDRequestMethodDelete) Version(v int) *ByProjectKeyPaymentsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentsByIDRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyPaymentsByIDRequestMethodDelete) Expand(v []string) *ByProjectKeyPaymentsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyPaymentsByIDRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyPaymentsByIDRequestMethodDelete) WithQueryParams(input ByProjectKeyPaymentsByIDRequestMethodDeleteInput) *ByProjectKeyPaymentsByIDRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyPaymentsByIDRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyPaymentsByIDRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyPaymentsByIDRequestMethodDelete) Execute(ctx context.Context) (result *Payment, err error) {
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
