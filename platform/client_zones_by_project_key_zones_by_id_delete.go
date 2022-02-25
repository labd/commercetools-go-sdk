package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyZonesByIDRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyZonesByIDRequestMethodDeleteInput
}

func (r *ByProjectKeyZonesByIDRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyZonesByIDRequestMethodDeleteInput struct {
	Version int
	Expand  []string
}

func (input *ByProjectKeyZonesByIDRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyZonesByIDRequestMethodDelete) Version(v int) *ByProjectKeyZonesByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyZonesByIDRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyZonesByIDRequestMethodDelete) Expand(v []string) *ByProjectKeyZonesByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyZonesByIDRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyZonesByIDRequestMethodDelete) WithQueryParams(input ByProjectKeyZonesByIDRequestMethodDeleteInput) *ByProjectKeyZonesByIDRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyZonesByIDRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyZonesByIDRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyZonesByIDRequestMethodDelete) Execute(ctx context.Context) (result *Zone, err error) {
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
