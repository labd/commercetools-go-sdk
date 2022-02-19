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

type ByProjectKeyProductSelectionsByIDRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductSelectionsByIDRequestMethodDeleteInput
}

func (r *ByProjectKeyProductSelectionsByIDRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductSelectionsByIDRequestMethodDeleteInput struct {
	Version int
	Expand  []string
}

func (input *ByProjectKeyProductSelectionsByIDRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyProductSelectionsByIDRequestMethodDelete) Version(v int) *ByProjectKeyProductSelectionsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsByIDRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyProductSelectionsByIDRequestMethodDelete) Expand(v []string) *ByProjectKeyProductSelectionsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductSelectionsByIDRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyProductSelectionsByIDRequestMethodDelete) WithQueryParams(input ByProjectKeyProductSelectionsByIDRequestMethodDeleteInput) *ByProjectKeyProductSelectionsByIDRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductSelectionsByIDRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyProductSelectionsByIDRequestMethodDelete {
	rb.headers = headers
	return rb
}

/**
*	Deletion will only succeed if the Product Selection is not assigned to any [Store](/../api/projects/stores#store).
 */
func (rb *ByProjectKeyProductSelectionsByIDRequestMethodDelete) Execute(ctx context.Context) (result *ProductSelection, err error) {
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
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
