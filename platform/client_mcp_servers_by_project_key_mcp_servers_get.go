package platform

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

type ByProjectKeyMcpServersRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyMcpServersRequestMethodGetInput
}

func (r *ByProjectKeyMcpServersRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyMcpServersRequestMethodGetInput struct {
	Sort      []string
	Limit     *int
	Offset    *int
	WithTotal *bool
}

func (input *ByProjectKeyMcpServersRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Sort {
		values.Add("sort", fmt.Sprintf("%v", v))
	}
	if input.Limit != nil {
		values.Add("limit", strconv.Itoa(*input.Limit))
	}
	if input.Offset != nil {
		values.Add("offset", strconv.Itoa(*input.Offset))
	}
	if input.WithTotal != nil {
		if *input.WithTotal {
			values.Add("withTotal", "true")
		} else {
			values.Add("withTotal", "false")
		}
	}
	return values
}

func (rb *ByProjectKeyMcpServersRequestMethodGet) Sort(v []string) *ByProjectKeyMcpServersRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMcpServersRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyMcpServersRequestMethodGet) Limit(v int) *ByProjectKeyMcpServersRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMcpServersRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyMcpServersRequestMethodGet) Offset(v int) *ByProjectKeyMcpServersRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMcpServersRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyMcpServersRequestMethodGet) WithTotal(v bool) *ByProjectKeyMcpServersRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMcpServersRequestMethodGetInput{}
	}
	rb.params.WithTotal = &v
	return rb
}

func (rb *ByProjectKeyMcpServersRequestMethodGet) WithQueryParams(input ByProjectKeyMcpServersRequestMethodGetInput) *ByProjectKeyMcpServersRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyMcpServersRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyMcpServersRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyMcpServersRequestMethodGet) Execute(ctx context.Context) (result *McpServerPagedQueryResponse, err error) {
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
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, ErrNotFound
	case 500:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 502:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
