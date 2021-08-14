// Generated file, please do not change!!!
package importapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyImportSinksRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyImportSinksRequestMethodGetInput
}

func (r *ByProjectKeyImportSinksRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyImportSinksRequestMethodGetInput struct {
	Limit  *float64
	Offset *float64
	Sort   []string
}

func (input *ByProjectKeyImportSinksRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Limit != nil {
		values.Add("limit", fmt.Sprintf("%f", *input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", fmt.Sprintf("%f", *input.Offset))
	}
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyImportSinksRequestMethodGet) Limit(v float64) *ByProjectKeyImportSinksRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyImportSinksRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyImportSinksRequestMethodGet) Offset(v float64) *ByProjectKeyImportSinksRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyImportSinksRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyImportSinksRequestMethodGet) Sort(v []string) *ByProjectKeyImportSinksRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyImportSinksRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyImportSinksRequestMethodGet) WithQueryParams(input ByProjectKeyImportSinksRequestMethodGetInput) *ByProjectKeyImportSinksRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyImportSinksRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyImportSinksRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves all import sinks of a project key.
 */
func (rb *ByProjectKeyImportSinksRequestMethodGet) Execute(ctx context.Context) (result *ImportSinkPagedResponse, err error) {
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
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
