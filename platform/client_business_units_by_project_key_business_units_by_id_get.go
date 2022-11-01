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

type ByProjectKeyBusinessUnitsByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyBusinessUnitsByIDRequestMethodGetInput
}

func (r *ByProjectKeyBusinessUnitsByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyBusinessUnitsByIDRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyBusinessUnitsByIDRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyBusinessUnitsByIDRequestMethodGet) Expand(v []string) *ByProjectKeyBusinessUnitsByIDRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyBusinessUnitsByIDRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyBusinessUnitsByIDRequestMethodGet) WithQueryParams(input ByProjectKeyBusinessUnitsByIDRequestMethodGetInput) *ByProjectKeyBusinessUnitsByIDRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyBusinessUnitsByIDRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyBusinessUnitsByIDRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyBusinessUnitsByIDRequestMethodGet) Execute(ctx context.Context) (result *BusinessUnit, err error) {
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
