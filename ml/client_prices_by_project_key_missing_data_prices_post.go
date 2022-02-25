package ml

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyMissingDataPricesRequestMethodPost struct {
	body    MissingPricesSearchRequest
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyMissingDataPricesRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyMissingDataPricesRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyMissingDataPricesRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyMissingDataPricesRequestMethodPost) Execute(ctx context.Context) (result *TaskToken, err error) {
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
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 202:
		err = json.Unmarshal(content, &result)
		return result, nil
	default:
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
