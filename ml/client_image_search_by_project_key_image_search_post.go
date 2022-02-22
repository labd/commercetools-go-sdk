// Generated file, please do not change!!!
package ml

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyImageSearchRequestMethodPost struct {
	body    io.Reader
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyImageSearchRequestMethodPostInput
}

func (r *ByProjectKeyImageSearchRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyImageSearchRequestMethodPostInput struct {
	Limit  *int
	Offset *int
}

func (input *ByProjectKeyImageSearchRequestMethodPostInput) Values() url.Values {
	values := url.Values{}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	return values
}

func (rb *ByProjectKeyImageSearchRequestMethodPost) Limit(v int) *ByProjectKeyImageSearchRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyImageSearchRequestMethodPostInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyImageSearchRequestMethodPost) Offset(v int) *ByProjectKeyImageSearchRequestMethodPost {
	if rb.params == nil {
		rb.params = &ByProjectKeyImageSearchRequestMethodPostInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyImageSearchRequestMethodPost) WithQueryParams(input ByProjectKeyImageSearchRequestMethodPostInput) *ByProjectKeyImageSearchRequestMethodPost {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyImageSearchRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyImageSearchRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Accepts an image file and returns similar products from product catalogue.
*
 */
func (rb *ByProjectKeyImageSearchRequestMethodPost) Execute(ctx context.Context) (result *ImageSearchResponse, err error) {
	data := rb.body
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
	default:
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
