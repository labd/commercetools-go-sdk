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

type ByProjectKeyExtensionsByIDRequestMethodPost struct {
	body    ExtensionUpdate
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyExtensionsByIDRequestMethodPostInput
}

func (r *ByProjectKeyExtensionsByIDRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyExtensionsByIDRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyExtensionsByIDRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyExtensionsByIDRequestMethodPost) Expand(v []string) *ByProjectKeyExtensionsByIDRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyExtensionsByIDRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyExtensionsByIDRequestMethodPost) WithQueryParams(input ByProjectKeyExtensionsByIDRequestMethodPostInput) *ByProjectKeyExtensionsByIDRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyExtensionsByIDRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyExtensionsByIDRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyExtensionsByIDRequestMethodPost) Execute(ctx context.Context) (result *Extension, err error) {
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
