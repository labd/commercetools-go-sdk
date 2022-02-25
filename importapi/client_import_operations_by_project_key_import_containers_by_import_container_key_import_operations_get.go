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

type ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGetInput
}

func (r *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGetInput struct {
	Limit       *float64
	Offset      *float64
	Sort        []string
	ResourceKey *string
	State       *ProcessingState
	Debug       *bool
}

func (input *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGetInput) Values() url.Values {
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

func (rb *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet) Limit(v float64) *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet) Offset(v float64) *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet) Sort(v []string) *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet) ResourceKey(v string) *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.ResourceKey = &v
	return rb
}

func (rb *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet) State(v ProcessingState) *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.State = &v
	return rb
}

func (rb *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet) Debug(v bool) *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGetInput{}
	}
	rb.params.Debug = &v
	return rb
}

func (rb *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet) WithQueryParams(input ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGetInput) *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves all [ImportOperations](ctp:import:type:ImportOperation) of a given ImportContainer key.
*
 */
func (rb *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet) Execute(ctx context.Context) (result *ImportOperationPagedResponse, err error) {
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
