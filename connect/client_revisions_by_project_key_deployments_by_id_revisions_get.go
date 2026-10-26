package connect

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyDeploymentsByIDRevisionsRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyDeploymentsByIDRevisionsRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyDeploymentsByIDRevisionsRequestMethodGet) WithHeaders(headers http.Header) *ByProjectKeyDeploymentsByIDRevisionsRequestMethodGet {
	rb.headers = headers
	return rb
}

/**
*	Retrieves the revisions of a Deployment, most recent first.
*
*	Only versions that were deployed successfully are returned, so every result is a valid rollback target.
*	At most 20 revisions are returned.
*
 */
func (rb *ByProjectKeyDeploymentsByIDRevisionsRequestMethodGet) Execute(ctx context.Context) (result *[]DeploymentRevision, err error) {
	queryParams := url.Values{}
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
