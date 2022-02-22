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

type ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput
}

func (r *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput struct {
	Limit       *float64
	Offset      *float64
	Sort        []string
	ResourceKey *string
	State       *ProcessingState
	Debug       *bool
}

func (input *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput) Values() url.Values {
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
		if *input.Debug {
			values.Add("debug", "true")
		} else {
			values.Add("debug", "false")
		}
	}
	return values
}

func (rb *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Limit(v float64) *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Offset(v float64) *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Sort(v []string) *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) ResourceKey(v string) *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.ResourceKey = &v
	return rb
}

func (rb *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) State(v ProcessingState) *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.State = &v
	return rb
}

func (rb *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Debug(v bool) *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.Debug = &v
	return rb
}

func (rb *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) WithQueryParams(input ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput) *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves all product-variant import operations of an import sink key.
 */
func (rb *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Execute(ctx context.Context) (result *ImportOperationPagedResponse, err error) {
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
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	default:
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
