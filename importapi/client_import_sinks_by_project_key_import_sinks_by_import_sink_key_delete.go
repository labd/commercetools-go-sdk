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

type ByProjectKeyImportSinksByImportSinkKeyRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyImportSinksByImportSinkKeyRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyImportSinksByImportSinkKeyRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyImportSinksByImportSinkKeyRequestMethodDelete {
	rb.headers = headers
	return rb
}

/**
*	Deletes the import sink given by the key.
 */
func (rb *ByProjectKeyImportSinksByImportSinkKeyRequestMethodDelete) Execute(ctx context.Context) (result *ImportSink, err error) {
	queryParams := url.Values{}
	resp, err := rb.client.delete(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		nil,
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
	case 404:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
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
