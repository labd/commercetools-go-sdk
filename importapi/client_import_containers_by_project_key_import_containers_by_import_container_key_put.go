package importapi

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyImportContainersByImportContainerKeyRequestMethodPut struct {
	body    ImportContainerUpdateDraft
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyImportContainersByImportContainerKeyRequestMethodPut) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyImportContainersByImportContainerKeyRequestMethodPut) WithHeaders(headers http.Header) *ByProjectKeyImportContainersByImportContainerKeyRequestMethodPut {
	rb.headers = headers
	return rb
}

/**
*	Updates the import container given by the key.
 */
func (rb *ByProjectKeyImportContainersByImportContainerKeyRequestMethodPut) Execute(ctx context.Context) (result *ImportContainer, err error) {
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
		if err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
