// Generated file, please do not change!!!
package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyCustomObjectsByContainerRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCustomObjectsByContainerRequestMethodGetInput
}

func (r *ByProjectKeyCustomObjectsByContainerRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCustomObjectsByContainerRequestMethodGetInput struct {
	Where        []string
	PredicateVar map[string][]string
	Expand       []string
}

func (input *ByProjectKeyCustomObjectsByContainerRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.Where {
		values.Add("where", fmt.Sprintf("%v", v))
	}
	for k, v := range input.PredicateVar {
		for _, x := range v {
			values.Set(k, x)
		}
	}
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) Where(v []string) *ByProjectKeyCustomObjectsByContainerRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsByContainerRequestMethodGetInput{}
	}
	rb.params.Where = v
	return rb
}

func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) PredicateVar(v map[string][]string) *ByProjectKeyCustomObjectsByContainerRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsByContainerRequestMethodGetInput{}
	}
	rb.params.PredicateVar = v
	return rb
}

func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) Expand(v []string) *ByProjectKeyCustomObjectsByContainerRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomObjectsByContainerRequestMethodGetInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) WithQueryParams(input ByProjectKeyCustomObjectsByContainerRequestMethodGetInput) *ByProjectKeyCustomObjectsByContainerRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyCustomObjectsByContainerRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyCustomObjectsByContainerRequestMethodGet) Execute(ctx context.Context) (result *CustomObjectPagedQueryResponse, err error) {
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
	case 400, 401, 403, 500, 502, 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
		}
		return nil, result
	default:
		return nil, fmt.Errorf("Unhandled StatusCode: %d", resp.StatusCode)
	}

}
