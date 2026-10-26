package connect

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRollbackRequestMethodPost struct {
	body    DeploymentRollback
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRollbackRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRollbackRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRollbackRequestMethodPost {
	rb.headers = headers
	return rb
}

/**
*	Restores the configuration of the revision onto the Deployment.
*	A new Deployment version is created from the revision and deployed.
*
*	The `manage_api_clients:{projectKey}` scope is required if the connector version uses commercetools managed API Client credentials.
*
*	Specific Error Codes:
*	- [DeploymentRollbackNotAllowed](ctp:connect:type:DeploymentRollbackNotAllowedError)
*	- [DeploymentInvalidStatusTransition](ctp:connect:type:DeploymentInvalidStatusTransitionError)
*	- [ConcurrentModification](ctp:connect:type:ConcurrentModificationError)
*
 */
func (rb *ByProjectKeyDeploymentsKeyByKeyRevisionsByRevisionRollbackRequestMethodPost) Execute(ctx context.Context) (result *Deployment, err error) {
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
	case 202:
		err = json.Unmarshal(content, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		return nil, ErrNotFound
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
