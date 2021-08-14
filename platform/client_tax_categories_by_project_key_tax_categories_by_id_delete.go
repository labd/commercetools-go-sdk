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

type ByProjectKeyTaxCategoriesByIDRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyTaxCategoriesByIDRequestMethodDeleteInput
}

func (r *ByProjectKeyTaxCategoriesByIDRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyTaxCategoriesByIDRequestMethodDeleteInput struct {
	Version int
	Expand  []string
}

func (input *ByProjectKeyTaxCategoriesByIDRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyTaxCategoriesByIDRequestMethodDelete) Version(v int) *ByProjectKeyTaxCategoriesByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyTaxCategoriesByIDRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyTaxCategoriesByIDRequestMethodDelete) Expand(v []string) *ByProjectKeyTaxCategoriesByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyTaxCategoriesByIDRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyTaxCategoriesByIDRequestMethodDelete) WithQueryParams(input ByProjectKeyTaxCategoriesByIDRequestMethodDeleteInput) *ByProjectKeyTaxCategoriesByIDRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyTaxCategoriesByIDRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyTaxCategoriesByIDRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyTaxCategoriesByIDRequestMethodDelete) Execute(ctx context.Context) (result *TaxCategory, err error) {
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
