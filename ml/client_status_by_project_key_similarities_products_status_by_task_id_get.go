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

type ByProjectKeySimilaritiesProductsStatusByTaskIdRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeySimilaritiesProductsStatusByTaskIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeySimilaritiesProductsStatusByTaskIdRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeySimilaritiesProductsStatusByTaskIdRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeySimilaritiesProductsStatusByTaskIdRequestMethodGet) Execute(ctx context.Context) (result *SimilarProductsTaskStatus, err error) {
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
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
