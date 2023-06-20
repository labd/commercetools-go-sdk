package connect

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestMethodGetInput
}

func (r *ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestMethodGetInput struct {
	ApplicationName *string
}

func (input *ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.ApplicationName != nil {
		values.Add("applicationName", fmt.Sprintf("%v", *input.ApplicationName))
	}
	return values
}

func (rb *ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestMethodGet) ApplicationName(v string) *ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestMethodGetInput{}
	}
	rb.params.ApplicationName = &v
	return rb
}

func (rb *ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestMethodGet) WithQueryParams(input ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestMethodGetInput) *ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Stream logs for the given deployment.
 */
func (rb *ByProjectKeyDeploymentsKeyByKeyLogsStreamRequestMethodGet) Execute(ctx context.Context) (result *interface{}, err error) {
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
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
