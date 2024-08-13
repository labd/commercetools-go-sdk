package connect

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

type ByProjectKeyDeploymentsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyDeploymentsRequestMethodGetInput
}

func (r *ByProjectKeyDeploymentsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyDeploymentsRequestMethodGetInput struct {
	IntegrationTypes []IntegrationType
	DeploymentType   *DeploymentType
	Limit            *int
	Offset           *int
	Sort             []string
}

func (input *ByProjectKeyDeploymentsRequestMethodGetInput) Values() url.Values {
	values := url.Values{}
	for _, v := range input.IntegrationTypes {
		values.Add("integrationTypes", fmt.Sprintf("%v", v))
	}
	if input.DeploymentType != nil {
		values.Add("deploymentType", fmt.Sprintf("%v", *input.DeploymentType))
	}
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

func (rb *ByProjectKeyDeploymentsRequestMethodGet) IntegrationTypes(v []IntegrationType) *ByProjectKeyDeploymentsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyDeploymentsRequestMethodGetInput{}
	}
	rb.params.IntegrationTypes = v
	return rb
}

func (rb *ByProjectKeyDeploymentsRequestMethodGet) DeploymentType(v DeploymentType) *ByProjectKeyDeploymentsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyDeploymentsRequestMethodGetInput{}
	}
	rb.params.DeploymentType = &v
	return rb
}

func (rb *ByProjectKeyDeploymentsRequestMethodGet) Limit(v int) *ByProjectKeyDeploymentsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyDeploymentsRequestMethodGetInput{}
	}
	rb.params.Limit = &v
	return rb
}

func (rb *ByProjectKeyDeploymentsRequestMethodGet) Offset(v int) *ByProjectKeyDeploymentsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyDeploymentsRequestMethodGetInput{}
	}
	rb.params.Offset = &v
	return rb
}

func (rb *ByProjectKeyDeploymentsRequestMethodGet) Sort(v []string) *ByProjectKeyDeploymentsRequestMethodGet {
	if rb.params == nil {
		rb.params = &ByProjectKeyDeploymentsRequestMethodGetInput{}
	}
	rb.params.Sort = v
	return rb
}

func (rb *ByProjectKeyDeploymentsRequestMethodGet) WithQueryParams(input ByProjectKeyDeploymentsRequestMethodGetInput) *ByProjectKeyDeploymentsRequestMethodGet {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyDeploymentsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyDeploymentsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves all deployments of a project key.
 */
func (rb *ByProjectKeyDeploymentsRequestMethodGet) Execute(ctx context.Context) (result *DeploymentPagedQueryResponse, err error) {
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
