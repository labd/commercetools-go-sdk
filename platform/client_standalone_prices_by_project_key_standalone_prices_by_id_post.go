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

type ByProjectKeyStandalonePricesByIDRequestMethodPost struct {
	body    StandalonePriceUpdate
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyStandalonePricesByIDRequestMethodPostInput
}

func (r *ByProjectKeyStandalonePricesByIDRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyStandalonePricesByIDRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyStandalonePricesByIDRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyStandalonePricesByIDRequestMethodPost) Expand(v []string) *ByProjectKeyStandalonePricesByIDRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyStandalonePricesByIDRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyStandalonePricesByIDRequestMethodPost) WithQueryParams(input ByProjectKeyStandalonePricesByIDRequestMethodPostInput) *ByProjectKeyStandalonePricesByIDRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyStandalonePricesByIDRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyStandalonePricesByIDRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyStandalonePricesByIDRequestMethodPost) Execute(ctx context.Context) (result *StandalonePrice, err error) {
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
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 409:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj

	case 500:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 502:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 503:
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
