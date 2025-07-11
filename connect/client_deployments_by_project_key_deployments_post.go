package connect

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyDeploymentsRequestMethodPost struct {
	body    DeploymentDraft
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyDeploymentsRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyDeploymentsRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyDeploymentsRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Specific Error Codes:
*	- [ConnectorStagedNotPreviewable](ctp:connect:type:ConnectorStagedNotPreviewableError)
*	- [DeploymentUnsupportedRegion](ctp:connect:type:DeploymentUnsupportedRegionError)
*	- [DeploymentUnknownApplicationConfiguration](ctp:connect:type:DeploymentUnknownApplicationConfigurationError)
*	- [DeploymentUnknownApplicationConfigurationKey](ctp:connect:type:DeploymentUnknownApplicationConfigurationKeyError)
*	- [DeploymentInvalidType](ctp:connect:type:DeploymentInvalidTypeError)
*	- [DeploymentProductionDeactivated](ctp:connect:type:DeploymentProductionDeactivatedError)
*	|
*	The `manage_api_clients:{projectKey}` scope is required if you use automatically generated API Client credentials.
*
 */
func (rb *ByProjectKeyDeploymentsRequestMethodPost) Execute(ctx context.Context) (result *Deployment, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	queryParams := url.Values{}
	resp, err := rb.client.post(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		data,
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
	case 201:
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
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
