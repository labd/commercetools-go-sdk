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

type ByProjectKeyImportContainersRequestMethodPost struct {
	body    ImportContainerDraft
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyImportContainersRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyImportContainersRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyImportContainersRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Creates a new import container.
 */
func (rb *ByProjectKeyImportContainersRequestMethodPost) Execute(ctx context.Context) (result *ImportContainer, err error) {
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
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
