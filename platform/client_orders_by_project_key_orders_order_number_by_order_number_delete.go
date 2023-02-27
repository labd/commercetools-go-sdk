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

type ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDeleteInput
}

func (r *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDeleteInput struct {
	DataErasure *bool
	Version     int
	Expand      []string
}

func (input *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDeleteInput) Values() url.Values {
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

func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete) DataErasure(v bool) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDeleteInput{}
	}
	rb.params.DataErasure = &v
	return rb
}

func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete) Version(v int) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete) Expand(v []string) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete) WithQueryParams(input ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDeleteInput) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete) Execute(ctx context.Context) (result *Order, err error) {
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
