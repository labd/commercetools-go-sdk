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

type ByProjectKeyImportSinksRequestMethodPost struct {
	body    ImportSinkDraft
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyImportSinksRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyImportSinksRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyImportSinksRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Creates a new import sink.
 */
func (rb *ByProjectKeyImportSinksRequestMethodPost) Execute(ctx context.Context) (result *ImportSink, err error) {
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
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
