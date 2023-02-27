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

type ByProjectKeyStoresRequestMethodPost struct {
	body    StoreDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyStoresRequestMethodPostInput
}

func (r *ByProjectKeyStoresRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyStoresRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyStoresRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyStoresRequestMethodPost) Expand(v []string) *ByProjectKeyStoresRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyStoresRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyStoresRequestMethodPost) WithQueryParams(input ByProjectKeyStoresRequestMethodPostInput) *ByProjectKeyStoresRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyStoresRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyStoresRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyStoresRequestMethodPost) Execute(ctx context.Context) (result *Store, err error) {
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
