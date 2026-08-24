package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyMcpServersTypesByMcpServerTypeRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyMcpServersTypesByMcpServerTypeRequestMethodGetInput
}

func (r *ByProjectKeyMcpServersTypesByMcpServerTypeRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyMcpServersTypesByMcpServerTypeRequestMethodGetInput struct {
	MajorVersion *string
}

func (input *ByProjectKeyMcpServersTypesByMcpServerTypeRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.MajorVersion != nil {
		values.Add("majorVersion", fmt.Sprintf("%v", *input.MajorVersion))
	}
	return values
}

func (rb *ByProjectKeyMcpServersTypesByMcpServerTypeRequestMethodGet) MajorVersion(v string) *ByProjectKeyMcpServersTypesByMcpServerTypeRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyMcpServersTypesByMcpServerTypeRequestMethodGetInput{}
	}
	rb.params.MajorVersion = &v
	return rb
}

func (rb *ByProjectKeyMcpServersTypesByMcpServerTypeRequestMethodGet) WithQueryParams(input ByProjectKeyMcpServersTypesByMcpServerTypeRequestMethodGetInput) *ByProjectKeyMcpServersTypesByMcpServerTypeRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyMcpServersTypesByMcpServerTypeRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyMcpServersTypesByMcpServerTypeRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Returns a single [MCP Server type](ctp:api:type:McpServerType). When `majorVersion` is omitted, the latest major version is returned.
*
 */
func (rb *ByProjectKeyMcpServersTypesByMcpServerTypeRequestMethodGet) Execute(ctx context.Context) (result *McpServerType, err error) {
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
