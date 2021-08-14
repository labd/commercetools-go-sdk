// Generated file, please do not change!!!
package ml

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
)

type ByProjectKeyImageSearchRequestMethodPost struct {
	body   io.Reader
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyImageSearchRequestMethodPostInput
}

func (r *ByProjectKeyImageSearchRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
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

func (rb *ByProjectKeyImageSearchRequestMethodPost) WithQueryParams(input *ByProjectKeyImageSearchRequestMethodPostInput) *ByProjectKeyImageSearchRequestMethodPost {
	rb.query = input.Values()
	return rb
}

/**
*	Accepts an image file and returns similar products from product catalogue.
*
 */
func (rb *ByProjectKeyImageSearchRequestMethodPost) Execute(ctx context.Context) (result *ImageSearchResponse, err error) {
	data := rb.body
	resp, err := rb.client.post(
		ctx,
		rb.url,
		rb.query,
		data,
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
