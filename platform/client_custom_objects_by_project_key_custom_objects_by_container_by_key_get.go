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

type ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGetInput
}

func (r *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGetInput struct {
	Expand []string
}

func (input *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGet) Expand(v []string) *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGet) WithQueryParams(input ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGetInput) *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyCustomObjectsByContainerByKeyRequestMethodGet) Execute(ctx context.Context) (result *CustomObject, err error) {
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
	case 400, 401, 403, 500, 502, 503:
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
