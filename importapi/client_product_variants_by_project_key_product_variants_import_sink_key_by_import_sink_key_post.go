// Generated file, please do not change!!!
package importapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyRequestMethodPost struct {
	body    ProductVariantImportRequest
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Creates import request for creating new product variants or updating existing ones.
 */
func (rb *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyRequestMethodPost) Execute(ctx context.Context) (result *ImportResponse, err error) {
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
	case 201:
		err = json.Unmarshal(content, &result)
		return result, nil
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
