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

type ByProjectKeySimilaritiesProductsRequestMethodPost struct {
	body    SimilarProductSearchRequest
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeySimilaritiesProductsRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeySimilaritiesProductsRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeySimilaritiesProductsRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeySimilaritiesProductsRequestMethodPost) Execute(ctx context.Context) (result *TaskToken, err error) {
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
	case 400:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
		}
		return nil, result
	default:
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
