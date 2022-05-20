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

type ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput
}

func (r *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput struct {
	Limit       *float64
	Offset      *float64
	Sort        []string
	ResourceKey *string
	State       *ProcessingState
	Debug       *bool
}

func (input *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput) Values() url.Values {
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

func (rb *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Limit(v float64) *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Offset(v float64) *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Sort(v []string) *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) ResourceKey(v string) *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.ResourceKey = &v
	return rb
}

func (rb *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) State(v ProcessingState) *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.State = &v
	return rb
}

func (rb *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Debug(v bool) *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.Debug = &v
	return rb
}

func (rb *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) WithQueryParams(input ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGetInput) *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves all import operations of an import sink key.
 */
func (rb *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet) Execute(ctx context.Context) (result *ImportOperationPagedResponse, err error) {
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
	case 400:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
		}
		return nil, result
	default:
		return nil, fmt.Errorf("unhandled StatusCode: %d", resp.StatusCode)
	}

}
