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

type ByProjectKeyProductSelectionsByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductSelectionsByIDRequestMethodGetInput
}

func (r *ByProjectKeyProductSelectionsByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductSelectionsByIDRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyProductSelectionsByIDRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyProductSelectionsByIDRequestMethodGet) Expand(v []string) *ByProjectKeyProductSelectionsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductSelectionsByIDRequestMethodGet) WithQueryParams(input ByProjectKeyProductSelectionsByIDRequestMethodGetInput) *ByProjectKeyProductSelectionsByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductSelectionsByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductSelectionsByIDRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyProductSelectionsByIDRequestMethodGet) Execute(ctx context.Context) (result *ProductSelection, err error) {
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

	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
