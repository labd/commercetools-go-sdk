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

type ByProjectKeyCategoriesByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCategoriesByIDRequestMethodGetInput
}

func (r *ByProjectKeyCategoriesByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCategoriesByIDRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyCategoriesByIDRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCategoriesByIDRequestMethodGet) Expand(v []string) *ByProjectKeyCategoriesByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCategoriesByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCategoriesByIDRequestMethodGet) WithQueryParams(input ByProjectKeyCategoriesByIDRequestMethodGetInput) *ByProjectKeyCategoriesByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCategoriesByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyCategoriesByIDRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyCategoriesByIDRequestMethodGet) Execute(ctx context.Context) (result *Category, err error) {
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
