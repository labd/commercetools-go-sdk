package importapi

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyProductTailoringsImportContainersByImportContainerKeyRequestMethodPost struct {
	body    ProductTailoringImportRequest
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyProductTailoringsImportContainersByImportContainerKeyRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyProductTailoringsImportContainersByImportContainerKeyRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyProductTailoringsImportContainersByImportContainerKeyRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Creates an Import Request for Product Tailoring.
 */
func (rb *ByProjectKeyProductTailoringsImportContainersByImportContainerKeyRequestMethodPost) Execute(ctx context.Context) (result *ImportResponse, err error) {
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
		if err != nil {
			return nil, err
		}
		return result, nil
	case 400:
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
