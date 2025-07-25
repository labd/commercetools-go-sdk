package importapi

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyImportContainersByImportContainerKeyImportSummariesRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyImportContainersByImportContainerKeyImportSummariesRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyImportContainersByImportContainerKeyImportSummariesRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyImportContainersByImportContainerKeyImportSummariesRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves an [ImportSummary](ctp:import:type:ImportSummary) for the [ImportContainer](ctp:import:type:ImportContainer) with the provided `importContainerKey`.
*
 */
func (rb *ByProjectKeyImportContainersByImportContainerKeyImportSummariesRequestMethodGet) Execute(ctx context.Context) (result *ImportSummary, err error) {
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
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
