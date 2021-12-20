// Generated file, please do not change!!!
package ml

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyImageSearchConfigRequestMethodPost struct {
	body    ImageSearchConfigRequest
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyImageSearchConfigRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyImageSearchConfigRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyImageSearchConfigRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Endpoint to update the image search config.
 */
func (rb *ByProjectKeyImageSearchConfigRequestMethodPost) Execute(ctx context.Context) (result *ImageSearchConfigResponse, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	queryParams := url.Values{}
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
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
