package importapi

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyImportSinksByImportSinkKeyRequestMethodPut struct {
	body    ImportSinkUpdateDraft
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyImportSinksByImportSinkKeyRequestMethodPut) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyImportSinksByImportSinkKeyRequestMethodPut) WithHeaders(headers http.Header) *ByProjectKeyImportSinksByImportSinkKeyRequestMethodPut {
	rb.headers = headers
	return rb
}

/**
*	Updates the import sink given by the key.
 */
func (rb *ByProjectKeyImportSinksByImportSinkKeyRequestMethodPut) Execute(ctx context.Context) (result *ImportSink, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	queryParams := url.Values{}
	resp, err := rb.client.put(
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
	case 409:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	default:
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
