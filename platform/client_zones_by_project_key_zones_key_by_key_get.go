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

type ByProjectKeyZonesKeyByKeyRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyZonesKeyByKeyRequestMethodGetInput
}

func (r *ByProjectKeyZonesKeyByKeyRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyZonesKeyByKeyRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyZonesKeyByKeyRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyZonesKeyByKeyRequestMethodGet) Expand(v []string) *ByProjectKeyZonesKeyByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyZonesKeyByKeyRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyZonesKeyByKeyRequestMethodGet) WithQueryParams(input ByProjectKeyZonesKeyByKeyRequestMethodGetInput) *ByProjectKeyZonesKeyByKeyRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyZonesKeyByKeyRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyZonesKeyByKeyRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyZonesKeyByKeyRequestMethodGet) Execute(ctx context.Context) (result *Zone, err error) {
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
