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

type ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput
}

func (r *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput struct {
	Limit       *float64
	Offset      *float64
	Sort        []string
	ResourceKey *string
	State       *ProcessingState
	Debug       *bool
}

func (input *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Limit != nil {
		values.Add("limit", fmt.Sprintf("%f", *input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", fmt.Sprintf("%f", *input.Offset))
	}
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	if input.ResourceKey != nil {
		values.Add("resourceKey", fmt.Sprintf("%v", *input.ResourceKey))
	}
	if input.State != nil {
		values.Add("state", fmt.Sprintf("%v", *input.State))
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

func (rb *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Limit(v float64) *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Offset(v float64) *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Sort(v []string) *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) ResourceKey(v string) *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.ResourceKey = &v
	return rb
}

func (rb *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) State(v ProcessingState) *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.State = &v
	return rb
}

func (rb *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Debug(v bool) *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.Debug = &v
	return rb
}

func (rb *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) WithQueryParams(input ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput) *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves all product-variant-patches import operations of an import sink key.
 */
func (rb *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Execute(ctx context.Context) (result *ImportOperationPagedResponse, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
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
