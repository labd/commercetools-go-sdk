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

type ByProjectKeyOrdersEditsRequestMethodPost struct {
	body    OrderEditDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyOrdersEditsRequestMethodPostInput
}

func (r *ByProjectKeyOrdersEditsRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyOrdersEditsRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyOrdersEditsRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyOrdersEditsRequestMethodPost) Expand(v []string) *ByProjectKeyOrdersEditsRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyOrdersEditsRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyOrdersEditsRequestMethodPost) WithQueryParams(input ByProjectKeyOrdersEditsRequestMethodPostInput) *ByProjectKeyOrdersEditsRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyOrdersEditsRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyOrdersEditsRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	You can either create multiple Order Edits for an Order and apply them sequentially to an Order, or create multiple Order Edits parallelly (as alternatives to each other) and apply one of them to the Order.
*
 */
func (rb *ByProjectKeyOrdersEditsRequestMethodPost) Execute(ctx context.Context) (result *OrderEdit, err error) {
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
