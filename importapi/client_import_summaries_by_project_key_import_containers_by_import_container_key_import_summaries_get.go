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
*	Retrieves an [ImportSummary](ctp:import:type:ImportSummary) for the given import container. An [ImportSummary](ctp:import:type:ImportSummary) is calculated on demand.
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
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
