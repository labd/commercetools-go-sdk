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

type ByProjectKeyMeBusinessUnitsRequestMethodPost struct {
	body    MyBusinessUnitDraft
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyMeBusinessUnitsRequestMethodPostInput
}

func (r *ByProjectKeyMeBusinessUnitsRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyMeBusinessUnitsRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyMeBusinessUnitsRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyMeBusinessUnitsRequestMethodPost) Expand(v []string) *ByProjectKeyMeBusinessUnitsRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyMeBusinessUnitsRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyMeBusinessUnitsRequestMethodPost) WithQueryParams(input ByProjectKeyMeBusinessUnitsRequestMethodPostInput) *ByProjectKeyMeBusinessUnitsRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyMeBusinessUnitsRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyMeBusinessUnitsRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyMeBusinessUnitsRequestMethodPost) Execute(ctx context.Context) (result *BusinessUnit, err error) {
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

	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
