// Generated file, please do not change!!!
package importapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
}

func (r *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

/**
*	Retrieves the ImportOperation of a given ID.
*
 */
func (rb *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet) Execute() (result *ImportOperation, err error) {
	resp, err := rb.client.get(
		context.Background(),
		rb.url,
		rb.query,
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
	case 404, 503:
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
