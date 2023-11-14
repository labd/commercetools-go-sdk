package connect

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type ByProjectKeyDeploymentsByIDLogsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyDeploymentsByIDLogsRequestMethodGetInput
}

func (r *ByProjectKeyDeploymentsByIDLogsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyDeploymentsByIDLogsRequestMethodGetInput struct {
	PageToken       *string
	ApplicationName *string
	StartDate       *time.Time
	EndDate         *time.Time
}

func (input *ByProjectKeyDeploymentsByIDLogsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	if input.PageToken != nil {
		values.Add("pageToken", fmt.Sprintf("%v", *input.PageToken))
	}
	if input.ApplicationName != nil {
		values.Add("applicationName", fmt.Sprintf("%v", *input.ApplicationName))
	}
	if input.StartDate != nil {
		values.Add("startDate", fmt.Sprintf("%v", *input.StartDate))
	}
	if input.EndDate != nil {
		values.Add("endDate", fmt.Sprintf("%v", *input.EndDate))
	}
	return values
}

func (rb *ByProjectKeyDeploymentsByIDLogsRequestMethodGet) PageToken(v string) *ByProjectKeyDeploymentsByIDLogsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyDeploymentsByIDLogsRequestMethodGetInput{}
	}
	rb.params.PageToken = &v
	return rb
}

func (rb *ByProjectKeyDeploymentsByIDLogsRequestMethodGet) ApplicationName(v string) *ByProjectKeyDeploymentsByIDLogsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyDeploymentsByIDLogsRequestMethodGetInput{}
	}
	rb.params.ApplicationName = &v
	return rb
}

func (rb *ByProjectKeyDeploymentsByIDLogsRequestMethodGet) StartDate(v time.Time) *ByProjectKeyDeploymentsByIDLogsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyDeploymentsByIDLogsRequestMethodGetInput{}
	}
	rb.params.StartDate = &v
	return rb
}

func (rb *ByProjectKeyDeploymentsByIDLogsRequestMethodGet) EndDate(v time.Time) *ByProjectKeyDeploymentsByIDLogsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyDeploymentsByIDLogsRequestMethodGetInput{}
	}
	rb.params.EndDate = &v
	return rb
}

func (rb *ByProjectKeyDeploymentsByIDLogsRequestMethodGet) WithQueryParams(input ByProjectKeyDeploymentsByIDLogsRequestMethodGetInput) *ByProjectKeyDeploymentsByIDLogsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyDeploymentsByIDLogsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyDeploymentsByIDLogsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves logs for the given deployment.
 */
func (rb *ByProjectKeyDeploymentsByIDLogsRequestMethodGet) Execute(ctx context.Context) (result *DeploymentLogCursorPagedQueryResponse, err error) {
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
