// Generated file, please do not change!!!
package importapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyRequestMethodPost struct {
	body   ProductVariantPatchRequest
	url    string
	client *Client
	query  url.Values
}

func (r *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

/**
*	Creates a new import request for product variant patches
 */
func (rb *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyRequestMethodPost) Execute() (result *ImportResponse, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	resp, err := rb.client.post(
		context.Background(),
		rb.url,
		rb.query,
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
