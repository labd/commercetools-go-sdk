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

type ByProjectKeyMissingDataImagesStatusByTaskIdRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyMissingDataImagesStatusByTaskIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyMissingDataImagesStatusByTaskIdRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyMissingDataImagesStatusByTaskIdRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyMissingDataImagesStatusByTaskIdRequestMethodGet) Execute(ctx context.Context) (result *MissingImagesTaskStatus, err error) {
	queryParams := url.Values{}
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
	default:
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
