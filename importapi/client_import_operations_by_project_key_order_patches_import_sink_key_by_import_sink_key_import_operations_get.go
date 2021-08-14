// Generated file, please do not change!!!
package importapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
)

type ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet struct {
	url    string
	client *Client
	query  url.Values
	params *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput
}

func (r *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":   r.url,
		"query": r.query,
	}
}

type ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput struct {
	Limit       *float64
	Offset      *float64
	Sort        *[]string
	ResourceKey *string
	State       *ProcessingState
	Debug       *bool
}

func (input *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Limit != nil {
		values.Add("limit", fmt.Sprintf("%f", *input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", fmt.Sprintf("%f", *input.Offset))
	}
	if input.Sort != nil {
		for _, v := range *input.Sort {
			values.Add("sort", v)
		}
	}
	if input.ResourceKey != nil {
		values.Add("resourceKey", *input.ResourceKey)
	}
	if input.State != nil {
		values.Add("state", *input.State)
	}
	if input.Debug != nil {
		if *input.Debug == true {
			values.Add("debug", "true")
		} else {
			values.Add("debug", "false")
		}
	}
	return values
}

func (rb *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) WithQueryParams(input *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput) *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	rb.query = input.Values()
	return rb
}

/**
*	Retrieves all OrderPatch ImportOperations of a given ImportSink key.
*
 */
func (rb *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Execute() (result *ImportOperationPagedResponse, err error) {
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
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
