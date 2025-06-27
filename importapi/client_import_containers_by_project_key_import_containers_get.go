package importapi

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyImportContainersRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyImportContainersRequestMethodGetInput
}

func (r *ByProjectKeyImportContainersRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyImportContainersRequestMethodGetInput struct {
	Limit  *int
	Offset *int
	Sort   []string
}

func (input *ByProjectKeyImportContainersRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyImportContainersRequestMethodGet) Limit(v int) *ByProjectKeyImportContainersRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyImportContainersRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyImportContainersRequestMethodGet) Offset(v int) *ByProjectKeyImportContainersRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyImportContainersRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyImportContainersRequestMethodGet) Sort(v []string) *ByProjectKeyImportContainersRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyImportContainersRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyImportContainersRequestMethodGet) WithQueryParams(input ByProjectKeyImportContainersRequestMethodGetInput) *ByProjectKeyImportContainersRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyImportContainersRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyImportContainersRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves all ImportContainers in the Project.
 */
func (rb *ByProjectKeyImportContainersRequestMethodGet) Execute(ctx context.Context) (result *ImportContainerPagedResponse, err error) {
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
