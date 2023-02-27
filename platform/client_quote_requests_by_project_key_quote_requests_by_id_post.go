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

type ByProjectKeyQuoteRequestsByIDRequestMethodPost struct {
	body    QuoteRequestUpdate
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyQuoteRequestsByIDRequestMethodPostInput
}

func (r *ByProjectKeyQuoteRequestsByIDRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyQuoteRequestsByIDRequestMethodPostInput struct {
	Expand []string
}

func (input *ByProjectKeyQuoteRequestsByIDRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyQuoteRequestsByIDRequestMethodPost) Expand(v []string) *ByProjectKeyQuoteRequestsByIDRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyQuoteRequestsByIDRequestMethodPostInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyQuoteRequestsByIDRequestMethodPost) WithQueryParams(input ByProjectKeyQuoteRequestsByIDRequestMethodPostInput) *ByProjectKeyQuoteRequestsByIDRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyQuoteRequestsByIDRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyQuoteRequestsByIDRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyQuoteRequestsByIDRequestMethodPost) Execute(ctx context.Context) (result *QuoteRequest, err error) {
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
	case 409, 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, ErrNotFound
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
